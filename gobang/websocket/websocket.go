package websocket

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"gobang/constants"
	"gobang/dto"
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

func Receive(s *melody.Session, msgByte []byte) {
	msg := &dto.Message{}
	if err := json.Unmarshal(msgByte, msg); err != nil {
		Send(s, dto.NewErrMsg(err))
	}
	switch msg.Code {
	case constants.HallChat:
		HallChat(s, msg)
	case constants.CreateRoom:

	case constants.EnterRoom:
		EnterRoom(s, msg)
	}
}

func Broadcast(msg *dto.Message) {
	msgByte, _ := json.Marshal(*msg)
	m.Broadcast(msgByte)
}
