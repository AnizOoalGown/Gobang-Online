package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gobang/entity"
)

func SetRoom(room *entity.Room) error {
	conn := pool.Get()
	defer conn.Close()

	str, err := json.Marshal(room)
	if err != nil {
		return fmt.Errorf("error marshal room: %v", err)
	}

	_, err = conn.Do("HSET", "room", room.Id, str)
	if err != nil {
		return fmt.Errorf("error add player: %v", err)
	}

	return nil
}

func DelRoom(id string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", "room", id)
	if err != nil {
		return fmt.Errorf("error delete room %v: %v", id, err)
	}

	return nil
}

func GetRoom(id string) (*entity.Room, error) {
	conn := pool.Get()
	defer conn.Close()

	b, err := redis.Bytes(conn.Do("HGET", "room", id))
	if err != nil {
		err = fmt.Errorf("redis: room with id '%v' not found", id)
		return nil, err
	}
	r := &entity.Room{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetRooms() (*[]entity.Room, error) {
	conn := pool.Get()
	defer conn.Close()

	bs, err := redis.ByteSlices(conn.Do("HVALS", "room"))
	if err != nil {
		return nil, err
	}
	rooms := make([]entity.Room, 0, 100)
	for _, b := range bs {
		r := &entity.Room{}
		err = json.Unmarshal(b, r)
		if err != nil {
			return nil, fmt.Errorf("unmarshal: %v", err)
		}
		rooms = append(rooms, *r)
	}
	return &rooms, nil
}
