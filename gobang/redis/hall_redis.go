package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gobang/entity"
)

func AddDialogMsg(msg *entity.DialogMsg) error {
	conn := pool.Get()
	defer conn.Close()

	str, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshal msg: %v", err)
	}

	_, err = conn.Do("RPUSH", "dialog", str)
	if err != nil {
		return fmt.Errorf("error add dialog msg: %v", err)
	}

	length, err := redis.Int(conn.Do("LLEN", "dialog"))
	if err != nil {
		return fmt.Errorf("error get dialog length: %v", err)
	}

	if length > 10 {
		_, err = conn.Do("LPOP", "dialog")
		if err != nil {
			return fmt.Errorf("error left pop dialog: %v", err)
		}
	}

	return nil
}

func GetDialog() (*[]entity.DialogMsg, error) {
	conn := pool.Get()
	defer conn.Close()

	bs, err := redis.ByteSlices(conn.Do("LRANGE", "dialog", 0, -1))
	if err != nil {
		return nil, fmt.Errorf("error get dialog: %v", err)
	}
	dialog := make([]entity.DialogMsg, 0, 10)
	for _, b := range bs {
		msg := &entity.DialogMsg{}
		if err = json.Unmarshal(b, msg); err != nil {
			return nil, fmt.Errorf("unmarshal: %v", err)
		}
		dialog = append(dialog, *msg)
	}
	return &dialog, nil
}
