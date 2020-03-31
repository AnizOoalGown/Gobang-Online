package service

import (
	"fmt"
	"gobang/entity"
	"gobang/redis"
)

func HallChat(msg *entity.DialogMsg) error {
	err := redis.AddDialogMsg(msg)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetHallDialog() (*[]entity.DialogMsg, error) {
	dialog, err := redis.GetDialog()
	if err != nil {
		fmt.Println(err)
	}
	return dialog, err
}
