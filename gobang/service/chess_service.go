package service

import (
	"fmt"
	"gobang/entity"
	"gobang/redis"
	"log"
)

func SetReady(rid string, pid string, ready bool) error {
	room, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return err
	}

	inRoom, role, _ := isInRoom(pid, room)

	if !inRoom {
		err = fmt.Errorf("error: Player %v not in room %v", pid, rid)
		log.Println(err)
		return err
	}

	if role == "host" {
		room.Host.Ready = ready
	} else if role == "challenger" {
		room.Challenger.Ready = ready
	} else {
		err = fmt.Errorf("error: Role %v cannot get ready", role)
		log.Println(err)
		return err
	}
	room.Started = room.Host.Ready && room.Challenger.Ready

	if err = redis.SetRoom(room); err != nil {
		return err
	}

	return nil
}

func MakeStep(rid string, c entity.Chess) error {
	room, err := redis.GetRoom(rid)
	if err != nil {
		log.Println(err)
		return err
	}
	if room.Started {
		room.Steps = append(room.Steps, c)
	} else {
		err = fmt.Errorf("error: Can not make step while game is not started")
		log.Println(err)
		return err
	}

	if err = redis.SetRoom(room); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
