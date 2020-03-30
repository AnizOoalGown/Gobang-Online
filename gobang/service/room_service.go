package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gobang/entity"
	"gobang/redis"
	"log"
)

func isInRoom(pid string, room *entity.Room) (inRoom bool, role string, index int) {
	if room.Host.Id == pid {
		inRoom = true
		role = "host"
		return
	} else if room.Challenger.Id == pid {
		inRoom = true
		role = "challenger"
		return
	} else if len(room.Spectators) > 0 {
		for i, p := range room.Spectators {
			if p.Id == pid {
				inRoom = true
				role = "spectator"
				index = i
				return
			}
		}
	}
	return
}

func EnterRoom(pid string, rid string, role string) error {
	p, err := redis.GetPlayer(pid)
	if err != nil {
		log.Println(err)
		return err
	}

	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return err
	}

	if r.Host.Id == "" {
		err = fmt.Errorf("error: Room %v has no host", rid)
		log.Println(err)
		return err
	}

	if inRoom, _, _ := isInRoom(pid, r); inRoom {
		err = fmt.Errorf("error: %v already in room %v", pid, rid)
		log.Println(err)
		return err
	}

	if role == "challenger" {
		r.Challenger = entity.PlayerDetails{
			Player: *p,
			Role:   role,
			Color:  1 - r.Host.Color,
			Turn:   false,
			Ready:  false,
		}
	} else if role == "spectator" {
		if p.Status == "leisure" {
			p.Status = "spectating"
			if err := SetPlayerStatus(pid, "spectating"); err != nil {
				log.Println(err)
				return err
			}
		}
		r.Spectators = append(r.Spectators, *p)
	} else {
		err = fmt.Errorf("error: The role /'%v/' can't enter room", role)
		log.Println(err)
		return err
	}

	if err = redis.SetRoom(r); err != nil {
		return err
	}
	return nil
}

func CreateRoom(pid string, color int8) (*entity.Room, error) {
	p, err := redis.GetPlayer(pid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	h := entity.PlayerDetails{
		Player: *p,
		Role:   "host",
		Color:  color,
		Turn:   false,
		Ready:  false,
	}

	r := &entity.Room{
		Id:         uuid.NewV4().String(),
		Dialog:     make([]entity.DialogMsg, 0),
		Steps:      make([]entity.Chess, 0),
		Started:    false,
		Host:       h,
		Challenger: entity.PlayerDetails{},
		Spectators: make([]entity.Player, 0),
	}

	if err = redis.SetRoom(r); err != nil {
		return nil, err
	}
	return r, nil
}

func LeaveRoom(pid string, rid string) error {
	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return err
	}

	inRoom, role, i := isInRoom(pid, r)
	if !inRoom {
		err = fmt.Errorf("error: Player %v not in room %v", pid, rid)
		log.Println(err)
		return err
	}

	if role == "host" {
		if r.Challenger.Id == "" {
			if err = redis.DelRoom(rid); err != nil {
				log.Println(err)
				return err
			}
		} else {
			r.Host = r.Challenger
		}
		r.Challenger = entity.PlayerDetails{}
	} else if role == "challenger" {
		r.Challenger = entity.PlayerDetails{}
	} else if role == "spectator" {
		r.Spectators = append(r.Spectators[:i], r.Spectators[i+1:]...)
	}

	if err = redis.SetRoom(r); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func RoomChat(rid string, msg *entity.DialogMsg) error {
	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(r.Dialog) >= 10 {
		r.Dialog = r.Dialog[1:]
	}
	r.Dialog = append(r.Dialog, *msg)

	if err = redis.SetRoom(r); err != nil {
		return err
	}

	return nil
}
