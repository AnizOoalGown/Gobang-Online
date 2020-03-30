package websocket

import (
	uuid "github.com/satori/go.uuid"
	"gobang/service"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
)

var (
	IdSessionMap sync.Map
)

func InitMelody() *melody.Melody {
	m := melody.New()

	m.HandleMessage(func(s *melody.Session, msg []byte) {

	})

	m.HandleConnect(func(s *melody.Session) {
		id := uuid.NewV4().String()

		_, err := service.NewPlayerConnect(id)
		if err != nil {
			return
		}
		IdSessionMap.Store(id, s)
		s.Set("id", id)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		idObject, ok := s.Get("id")
		if !ok {
			log.Println("session with no 'id' key")
			return
		}
		id := idObject.(string)
		IdSessionMap.Delete(id)
		service.PlayerDisconnect(id)
	})

	return m
}
