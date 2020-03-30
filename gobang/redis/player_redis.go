package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gobang/entity"
)

func SetPlayer(player *entity.Player) error {
	conn := pool.Get()
	defer conn.Close()

	str, err := json.Marshal(player)
	if err != nil {
		return fmt.Errorf("error marshal player: %v", err)
	}

	_, err = conn.Do("HSET", "player", player.Id, str)
	if err != nil {
		return fmt.Errorf("error add player: %v", err)
	}

	return nil
}

func GetPlayer(id string) (*entity.Player, error) {
	conn := pool.Get()
	defer conn.Close()

	b, err := redis.Bytes(conn.Do("HGET", "player", id))
	if err != nil {
		return nil, err
	}
	p := &entity.Player{}
	err = json.Unmarshal(b, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
func GetPlayers() (*[]entity.Player, error) {
	conn := pool.Get()
	defer conn.Close()

	bs, err := redis.ByteSlices(conn.Do("HVALS", "player"))
	if err != nil {
		return nil, err
	}
	players := make([]entity.Player, 0, 100)
	for _, b := range bs {
		p := &entity.Player{}
		err = json.Unmarshal(b, p)
		if err != nil {
			return nil, fmt.Errorf("unmarshal: %v", err)
		}
		players = append(players, *p)
	}
	return &players, nil
}

func DelPlayer(id string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", "player", id)
	if err != nil {
		return fmt.Errorf("error delete player %v: %v", id, err)
	}

	return nil
}
