package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gobang/entity"
	"gobang/redis"
	"log"
)

func GetRooms() (*[]entity.Room, error) {
	return redis.GetRooms()
}

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

func EnterRoom(pid string, rid string, role string) (*entity.Room, error) {
	p, err := redis.GetPlayer(pid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if r.Host.Id == "" {
		err = fmt.Errorf("error: Room %v has no host", rid)
		log.Println(err)
		return nil, err
	}

	if inRoom, _, _ := isInRoom(pid, r); inRoom {
		err = fmt.Errorf("error: %v already in room %v", pid, rid)
		log.Println(err)
		return nil, err
	}

	if role == "challenger" {
		r.Challenger = entity.PlayerDetails{
			Player: *p,
			Role:   role,
			Color:  1 - r.Host.Color,
			Ready:  false,
		}
	} else if role == "spectator" {
		if p.Status == "leisure" {
			p.Status = "spectating"
			if err := SetPlayerStatus(pid, "spectating"); err != nil {
				log.Println(err)
				return nil, err
			}
		}
		r.Spectators = append(r.Spectators, *p)
	} else {
		err = fmt.Errorf("error: The role /'%v/' can't enter room", role)
		log.Println(err)
		return nil, err
	}

	if err = redis.SetRoom(r); err != nil {
		return nil, err
	}

	//New player enters room can't see the dialog before
	r.Dialog = make([]entity.DialogMsg, 0)
	return r, nil
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
		Ready:  false,
	}

	r := &entity.Room{
		Id:      uuid.NewV4().String(),
		Dialog:  make([]entity.DialogMsg, 0),
		Steps:   make([]entity.Chess, 0),
		Started: false,
		Host:    h,
		Challenger: entity.PlayerDetails{
			Color: 1 - color,
		},
		Spectators: make([]entity.Player, 0),
	}

	if err = redis.SetRoom(r); err != nil {
		return nil, err
	}
	return r, nil
}

func LeaveRoom(pid string, rid string) (*entity.Room, error) {
	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	inRoom, role, i := isInRoom(pid, r)
	if !inRoom {
		return r, nil
	}

	if role == "host" {
		if r.Challenger.Id == "" {
			if err = redis.DelRoom(rid); err != nil {
				log.Println(err)
				return nil, err
			}
			r.Host = entity.PlayerDetails{}
			return r, nil
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
		return nil, err
	}

	return r, nil
}

func RoomChat(rid string, msg *entity.DialogMsg) (*entity.Room, error) {
	r, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(r.Dialog) >= 10 {
		r.Dialog = r.Dialog[1:]
	}
	r.Dialog = append(r.Dialog, *msg)

	if err = redis.SetRoom(r); err != nil {
		return nil, err
	}

	return r, nil
}
