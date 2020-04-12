package websocket

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gobang/constants"
	"gobang/dto"
	"gobang/entity"
	"gobang/service"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
)

var (
	idSessionMap sync.Map
	m            *melody.Melody
	lock         sync.Mutex
)

func InitMelody() *melody.Melody {
	m = melody.New()
	m.HandleMessage(Receive)
	m.HandleConnect(Connect)
	m.HandleDisconnect(Disconnect)
	return m
}

func GetPId(s *melody.Session) (pid string, ok bool) {
	pidObj, ok := s.Get("id")
	if !ok {
		return "", false
	}
	pid, ok = pidObj.(string)
	return
}

func Connect(s *melody.Session) {
	id := uuid.NewV4().String()
	player, err := service.NewPlayerConnect(id)
	if err != nil {
		return
	}
	idSessionMap.Store(id, s)
	s.Set("id", id)
	Send(s, &dto.Message{
		Code: constants.GetPlayer,
		Data: player,
	})
}

func Disconnect(s *melody.Session) {
	idObject, ok := s.Get("id")
	if !ok {
		log.Println("session with no 'id' key")
		return
	}
	id := idObject.(string)

	lock.Lock()
	defer lock.Unlock()

	idSessionMap.Delete(id)
	rooms, err := service.PlayerDisconnect(id)
	if err != nil {
		SendErr(s, err)
	}
	for _, room := range *rooms {
		r, gameOverDTO, err := service.LeaveRoom(id, room.Id)
		if err != nil {
			SendErr(s, err)
		}
		if gameOverDTO != nil {
			SendGameOver(r, gameOverDTO)
		}
	}
}

func Send(s *melody.Session, msg *dto.Message) {
	msgByte, _ := json.Marshal(msg)
	s.Write(msgByte)
}

func SendSuccess(s *melody.Session) {
	Send(s, &dto.Message{
		Code: constants.Success,
		Data: "OK",
	})
}

func SendErr(s *melody.Session, err error) {
	Send(s, dto.NewErrMsg(err))
}

func Send2PId(pid string, msg *dto.Message) {
	sObj, ok := idSessionMap.Load(pid)
	if !ok {
		err := fmt.Errorf("error: can not load the value of %v from idSessionMap", pid)
		log.Println(err)
		return
	}
	s, ok := sObj.(*melody.Session)
	if !ok {
		err := fmt.Errorf("error: sObj is not type of *melody.Session")
		log.Println(err)
		return
	}
	Send(s, msg)
}

func Send2Room(r *entity.Room, msg *dto.Message) {
	if r.Host.Id != "" {
		Send2PId(r.Host.Id, msg)
	}
	if r.Challenger.Id != "" {
		Send2PId(r.Challenger.Id, msg)
	}
	for _, spectator := range r.Spectators {
		Send2PId(spectator.Id, msg)
	}
}

func SendGameOver(room *entity.Room, gameOverDTO *dto.GameOverDTO) {
	msg := &dto.Message{
		Code: constants.GameOver,
		Data: *gameOverDTO,
	}

	Send2Room(room, msg)
}

func Receive(s *melody.Session, msgByte []byte) {
	msg := &dto.Message{}
	if err := json.Unmarshal(msgByte, msg); err != nil {
		Send(s, dto.NewErrMsg(err))
	}

	switch msg.Code {
	case constants.HallChat:
		HallChat(s, msg)
	case constants.GetHallDialog:
		GetHallDialog(s, msg)
	case constants.GetRooms:
		GetRooms(s, msg)
	case constants.CreateRoom:
		CreateRoom(s, msg)
	case constants.EnterRoom:
		EnterRoom(s, msg)
	case constants.LeaveRoom:
		LeaveRoom(s, msg)
	case constants.RoomChat:
		RoomChat(s, msg)
	case constants.GetPlayer:
		GetPlayer(s, msg)
	case constants.GetPlayers:
		GetPlayers(s, msg)
	case constants.PlayerRename:
		PlayerRename(s, msg)
	case constants.SetPlayerStatus:
		SetPlayerStatus(s, msg)
	case constants.SetReady:
		SetReady(s, msg)
	case constants.MakeStep:
		MakeStep(s, msg)
	case constants.RetractStep:
		RetractStep(s, msg)
	case constants.Surrender:
		Surrender(s, msg)
	case constants.AskDraw:
		AskDraw(s, msg)
	}
}

func Broadcast(msg *dto.Message) {
	msgByte, _ := json.Marshal(*msg)
	m.Broadcast(msgByte)
}
