package service

import (
	"gobang/entity"
	"gobang/redis"
)

func HallChat(msg *entity.DialogMsg) error {
	err := redis.AddDialogMsg(msg)
	if err != nil {
		logger.Error(err)
	}

	return err
}

func GetHallDialog() (*[]entity.DialogMsg, error) {
	dialog, err := redis.GetDialog()
	if err != nil {
		logger.Error(err)
	}
	return dialog, err
}
