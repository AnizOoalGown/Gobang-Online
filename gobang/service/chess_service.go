package service

import (
	"fmt"
	"gobang/constants"
	"gobang/dto"
	"gobang/entity"
	"gobang/lock"
	"gobang/redis"
	"gobang/util"
)

func SetReady(rid string, pid string, ready bool) (*entity.Room, error) {
	lock.RoomLock.Lock(rid)
	defer lock.RoomLock.Unlock(rid)

	room, err := redis.GetRoom(rid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	inRoom, role, _ := isInRoom(pid, room)

	if !inRoom {
		err = fmt.Errorf("error: Player %v not in room %v", pid, rid)
		logger.Error(err)
		return nil, err
	}

	if role == "host" {
		room.Host.Ready = ready
	} else if role == "challenger" {
		room.Challenger.Ready = ready
	} else {
		err = fmt.Errorf("error: Role %v cannot get ready", role)
		logger.Error(err)
		return nil, err
	}
	room.Started = room.Host.Ready && room.Challenger.Ready
	if room.Started {
		room.Steps = make([]entity.Chess, 0)
	}

	if err = redis.SetRoom(room); err != nil {
		logger.Error(err)
		return nil, err
	}

	return room, nil
}

func MakeStep(rid string, c entity.Chess) (bool, *dto.GameOverDTO, *entity.Room, error) {
	lock.RoomLock.Lock(rid)
	defer lock.RoomLock.Unlock(rid)

	room, err := redis.GetRoom(rid)
	if err != nil {
		logger.Error(err)
		return false, nil, nil, err
	}
	if room.Started {
		room.Steps = append(room.Steps, c)
	} else {
		err = fmt.Errorf("error: Can not make step while game is not started")
		logger.Error(err)
		return false, nil, nil, err
	}

	over, gameOverDTO, err := CheckFive(room)
	if err != nil {
		logger.Error(err)
		return false, nil, nil, err
	}
	if err = redis.SetRoom(room); err != nil {
		logger.Error(err)
		return false, nil, nil, err
	}
	return over, gameOverDTO, room, nil
}

func PrepareNewGame(room *entity.Room) {
	room.Host.Ready = false
	room.Challenger.Ready = false
	room.Host.Color = 1 - room.Host.Color
	room.Challenger.Color = 1 - room.Challenger.Color
	room.Started = false
}

func CheckFive(room *entity.Room) (bool, *dto.GameOverDTO, error) {
	hasFive, color := util.CheckFiveOfLastStep(&room.Steps)
	if !hasFive {
		return false, nil, nil
	}

	var gameOverDTO *dto.GameOverDTO
	if room.Host.Color == color {
		gameOverDTO = &dto.GameOverDTO{
			RId:    room.Id,
			Winner: room.Host,
			Loser:  room.Challenger,
			Cause:  "five",
		}
	} else {
		gameOverDTO = &dto.GameOverDTO{
			RId:    room.Id,
			Winner: room.Challenger,
			Loser:  room.Host,
			Cause:  "five",
		}
	}

	PrepareNewGame(room)
	return true, gameOverDTO, nil
}

func RetractStep(pid string, rid string, consent int) (string, *entity.Room, int, error) {
	lock.RoomLock.Lock(rid)
	defer lock.RoomLock.Unlock(rid)

	room, err := redis.GetRoom(rid)
	if err != nil {
		logger.Error(err)
		return "", nil, 0, err
	}
	length := len(room.Steps)
	if !room.Started || length == 0 {
		err = fmt.Errorf("error: room %v is not started or there is no step", rid)
		logger.Error(err)
		return "", nil, 0, err
	}

	inRoom, role, _ := isInRoom(pid, room)
	if !inRoom || role == "spectator" {
		err = fmt.Errorf("error: player %v is not playing in room %v", pid, rid)
		logger.Error(err)
		return "", nil, 0, err
	}

	var opponentId string
	var color int8
	if role == "host" {
		opponentId = room.Challenger.Id
		color = room.Challenger.Color
	} else if role == "challenger" {
		opponentId = room.Host.Id
		color = room.Host.Color
	}

	var count int
	if consent == 2 {
		if length == 1 && color == constants.WHITE {
			err = fmt.Errorf("error: there is no white step so white side can't retract")
			logger.Error(err)
			return "", nil, 0, err
		}
		lastColor := int8((length - 1) % 2)
		if lastColor == color {
			count = 1
			room.Steps = room.Steps[:length-1]
		} else {
			count = 2
			room.Steps = room.Steps[:length-2]
		}
	}

	if err = redis.SetRoom(room); err != nil {
		logger.Error(err)
		return "", nil, 0, err
	}
	return opponentId, room, count, err
}

func Surrender(pid string, rid string) (*dto.GameOverDTO, *entity.Room, error) {
	lock.RoomLock.Lock(rid)
	defer lock.RoomLock.Unlock(rid)

	room, err := redis.GetRoom(rid)
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	if !room.Started {
		err = fmt.Errorf("error: room %v is not started", rid)
		logger.Error(err)
		return nil, nil, err
	}

	inRoom, role, _ := isInRoom(pid, room)
	if !inRoom || role == "spectator" {
		err = fmt.Errorf("error: player %v is not playing in room %v", pid, rid)
		logger.Error(err)
		return nil, nil, err
	}

	gameOverDTO := &dto.GameOverDTO{
		RId:   rid,
		Cause: "surrender",
	}

	if role == "host" {
		gameOverDTO.Winner = room.Challenger
		gameOverDTO.Loser = room.Host
	} else if role == "challenger" {
		gameOverDTO.Winner = room.Host
		gameOverDTO.Loser = room.Challenger
	}

	PrepareNewGame(room)
	if err = redis.SetRoom(room); err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	return gameOverDTO, room, nil
}

func Draw(pid string, rid string, consent int) (string, *entity.Room, error) {
	lock.RoomLock.Lock(rid)
	defer lock.RoomLock.Unlock(rid)

	room, err := redis.GetRoom(rid)
	if err != nil {
		logger.Error(err)
		return "", nil, err
	}
	if !room.Started {
		err = fmt.Errorf("error: room %v is not started", rid)
		logger.Error(err)
		return "", nil, err
	}

	inRoom, role, _ := isInRoom(pid, room)
	if !inRoom || role == "spectator" {
		err = fmt.Errorf("error: player %v is not playing in room %v", pid, rid)
		logger.Error(err)
		return "", nil, err
	}

	if consent == 2 {
		PrepareNewGame(room)
		if err = redis.SetRoom(room); err != nil {
			logger.Error(err)
			return "", nil, err
		}
	}

	var opponentId string
	if role == "host" {
		opponentId = room.Challenger.Id
	} else if role == "challenger" {
		opponentId = room.Host.Id
	}

	return opponentId, room, nil
}
