package websocket

import (
	"fmt"
	"gobang/constants"
	"gobang/dto"
	"gobang/entity"
	"gobang/service"
	"gopkg.in/olahol/melody.v1"
	"time"
)

func HallChat(s *melody.Session, msg *dto.Message) {
	id, _ := s.Get("id")
	content, ok := msg.Data.(string)
	if !ok {
		err := fmt.Errorf("interface conversion: data is not string")
		Send(s, dto.NewErrMsg(err))
		return
	}

	p, err := service.GetPlayer(id.(string))
	if err != nil {
		Send(s, dto.NewErrMsg(err))
		return
	}

	dialogMsg := &entity.DialogMsg{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		From:    p.Name,
		Content: content,
	}
	if err := service.HallChat(dialogMsg); err != nil {
		Send(s, dto.NewErrMsg(err))
		return
	}
	msg = &dto.Message{
		Code: constants.HallChat,
		Data: *dialogMsg,
	}
	Broadcast(msg)
}

func GetHallDialog(s *melody.Session, msg *dto.Message) {
	dialog, err := service.GetHallDialog()
	if err != nil {
		Send(s, dto.NewErrMsg(err))
		return
	}
	msg.Data = dialog
	Send(s, msg)
}

func GetRooms(s *melody.Session, msg *dto.Message) {
	rooms, err := service.GetRooms()
	if err != nil {
		Send(s, dto.NewErrMsg(err))
		return
	}
	msg.Data = rooms
	Send(s, msg)
}

func CreateRoom(s *melody.Session, msg *dto.Message) {
	pidObj, _ := s.Get("id")
	pid, _ := pidObj.(string)
	//
	color, ok := msg.Data.(int)
	if !ok {
		err := fmt.Errorf("interface conversion: data is not int")
		Send(s, dto.NewErrMsg(err))
		return
	}

	room, err := service.CreateRoom(pid, int8(msg.Data))
	if err != nil {
		Send(s, dto.NewErrMsg(err))
		return
	}

	msg.Data = room
	Send(s, msg)
}

func EnterRoom(s *melody.Session, msg *dto.Message) {
	//pidObj, _ := s.Get("id")
	//pid, _ := pidObj.(string)
	//
	//rid, ok := msg.Data.(string)
	//if !ok {
	//	err := fmt.Errorf("interface conversion: data is not string")
	//	Send(s, dto.NewErrMsg(err))
	//	return
	//}
	//
	//if err := service.EnterRoom(pid, rid, msg.Info); err != nil {
	//	Send(s, dto.NewErrMsg(err))
	//	return
	//}
	//Broadcast(msg)
}
