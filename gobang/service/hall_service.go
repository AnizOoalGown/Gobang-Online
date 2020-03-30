package service

import (
	"gobang/entity"
	"gobang/redis"
)

func HallChat(msg *entity.DialogMsg) error {
	return redis.AddDialogMsg(msg)
}

func GetHallDialog() (*[]entity.DialogMsg, error) {
	return redis.GetDialog()
}
