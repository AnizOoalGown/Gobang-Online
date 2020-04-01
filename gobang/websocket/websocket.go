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
)

func InitMelody() *melody.Melody {
	m = melody.New()
	m.HandleMessage(Receive)
	m.HandleConnect(Connect)
	m.HandleDisconnect(Disconnect)
	return m
}

func Connect(s *melody.Session) {
	id := uuid.NewV4().String()
	_, err := service.NewPlayerConnect(id)
	if err != nil {
		return
	}
	idSessionMap.Store(id, s)
	s.Set("id", id)
}

func Disconnect(s *melody.Session) {
	idObject, ok := s.Get("id")
	if !ok {
		log.Println("session with no 'id' key")
		return
	}
	id := idObject.(string)
	idSessionMap.Delete(id)
	service.PlayerDisconnect(id)
}

func Send(s *melody.Session, msg *dto.Message) {
	msgByte, _ := json.Marshal(msg)
	s.Write(msgByte)
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
	Send2PId(r.Host.Id, msg)
	Send2PId(r.Challenger.Id, msg)
	for _, spectator := range r.Spectators {
		Send2PId(spectator.Id, msg)
	}
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
	}
}

func Broadcast(msg *dto.Message) {
	msgByte, _ := json.Marshal(*msg)
	m.Broadcast(msgByte)
}
