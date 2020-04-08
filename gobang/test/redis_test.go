package test

import (
	"fmt"
	"gobang/entity"
	"gobang/redis"
	"log"
	"testing"
)

func TestAddDialogMsg(t *testing.T) {
	msg := &entity.DialogMsg{
		Time:    "23:35",
		From:    "A",
		Content: "12",
	}
	redis.AddDialogMsg(msg)
	dialog, err := redis.GetDialog()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(*dialog)
}
func TestGetPlayers(t *testing.T) {
	redis.GetPlayers()
}

func TestAddRoom(t *testing.T) {
	pd := entity.PlayerDetails{
		Player: entity.Player{
			Id:     "p125",
			Name:   "Tom",
			Status: "leisure",
		},
		Role:  "host",
		Color: 0,
		Ready: false,
	}
	room := &entity.Room{
		Id:      "r123",
		Dialog:  nil,
		Steps:   nil,
		Started: false,
		Host:    pd,
		//Challenger: nil,
		//Spectators: nil,
	}
	redis.SetRoom(room)
	rooms, err := redis.GetRooms()
	if err != nil {
		log.Println(err)
	}
	log.Println(*rooms)
}
