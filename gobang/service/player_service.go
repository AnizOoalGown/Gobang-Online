package service

import (
	"gobang/entity"
	"gobang/redis"
)

func NewPlayerConnect(id string) (*entity.Player, error) {
	p := &entity.Player{
		Id:     id,
		Name:   "unnamed",
		Status: "leisure",
	}
	err := redis.SetPlayer(p)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.WithField("pid", id).Debug("player connects")
	return p, nil
}

func GetPlayer(id string) (*entity.Player, error) {
	return redis.GetPlayer(id)
}

func GetPlayers() (*[]entity.Player, error) {
	return redis.GetPlayers()
}

func PlayerDisconnect(id string) (*[]entity.Room, error) {
	err := redis.DelPlayer(id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.WithField("pid", id).Debug("player disconnects")
	return redis.GetRooms()
}

func PlayerRename(id string, newName string) error {
	p, err := redis.GetPlayer(id)
	if err != nil {
		logger.Error(err)
		return err
	}
	p.Name = newName
	err = redis.SetPlayer(p)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func SetPlayerStatus(id string, status string) error {
	p, err := redis.GetPlayer(id)
	if err != nil {
		logger.Error(err)
		return err
	}
	p.Status = status
	err = redis.SetPlayer(p)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
