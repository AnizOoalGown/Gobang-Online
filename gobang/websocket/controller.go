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
		SendErr(s, err)
		return
	}

	p, err := service.GetPlayer(id.(string))
	if err != nil {
		SendErr(s, err)
		return
	}

	dialogMsg := &entity.DialogMsg{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		From:    p.Name,
		Content: content,
	}
	if err := service.HallChat(dialogMsg); err != nil {
		SendErr(s, err)
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
		SendErr(s, err)
		return
	}
	msg.Data = dialog
	Send(s, msg)
}

func GetRooms(s *melody.Session, msg *dto.Message) {
	rooms, err := service.GetRooms()
	if err != nil {
		SendErr(s, err)
		return
	}
	msg.Data = rooms
	Send(s, msg)
}

func CreateRoom(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	//
	color, ok := msg.Data.(float64)
	if !ok {
		err := fmt.Errorf("interface conversion: data is not a number")
		SendErr(s, err)
		return
	}

	room, err := service.CreateRoom(pid, int8(color))
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = room
	Send(s, msg)

	rooms, err := service.GetRooms()
	if err != nil {
		SendErr(s, err)
		return
	}
	Broadcast(&dto.Message{
		Code: constants.GetRooms,
		Data: rooms,
	})
}

func EnterRoom(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)

	data := msg.Data.(map[string]interface{})
	rid := data["rid"].(string)
	role := data["role"].(string)

	room, err := service.EnterRoom(pid, rid, role)
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = room
	Send2Room(room, msg)
}

func LeaveRoom(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)

	rid, ok := msg.Data.(string)
	if !ok {
		err := fmt.Errorf("error: data is not string")
		SendErr(s, err)
		return
	}

	SendLeaveRoom(s, pid, rid)
}

func RoomChat(s *melody.Session, msg *dto.Message) {
	data := msg.Data.(map[string]interface{})
	from := data["from"].(string)
	content := data["content"].(string)
	rid := data["rid"].(string)

	dialogMsg := &entity.DialogMsg{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		From:    from,
		Content: content,
	}

	room, err := service.RoomChat(rid, dialogMsg)
	if err != nil {
		SendErr(s, err)
		return
	}

	roomChatDTO := struct {
		RoomId string `json:"rid"`
		entity.DialogMsg
	}{
		rid,
		*dialogMsg,
	}

	msg.Data = roomChatDTO
	Send2Room(room, msg)
}

func GetPlayer(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	player, err := service.GetPlayer(pid)
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = player
	Send(s, msg)
}

func GetPlayers(s *melody.Session, msg *dto.Message) {
	players, err := service.GetPlayers()
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = players
	Send(s, msg)
}

func PlayerRename(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	name, ok := msg.Data.(string)
	if !ok {
		err := fmt.Errorf("interface conversion: data is not string")
		SendErr(s, err)
		return
	}

	if err := service.PlayerRename(pid, name); err != nil {
		SendErr(s, err)
		return
	}

	msg.Code = constants.GetPlayers
	players, err := service.GetPlayers()
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = players
	Broadcast(msg)
}

func SetPlayerStatus(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	status, ok := msg.Data.(string)
	if !ok {
		err := fmt.Errorf("interface conversion: data is not string")
		SendErr(s, err)
		return
	}

	if err := service.SetPlayerStatus(pid, status); err != nil {
		SendErr(s, err)
		return
	}

	players, err := service.GetPlayers()
	if err != nil {
		SendErr(s, err)
		return
	}
	msg.Code = constants.GetPlayers
	msg.Data = players
	Broadcast(msg)
}

func SetReady(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	data := msg.Data.(map[string]interface{})
	rid := data["rid"].(string)
	ready := data["ready"].(bool)

	room, err := service.SetReady(rid, pid, ready)
	if err != nil {
		SendErr(s, err)
		return
	}

	msg.Data = room
	Send2Room(room, msg)
}

func MakeStep(s *melody.Session, msg *dto.Message) {
	data := msg.Data.(map[string]interface{})
	rid := data["rid"].(string)
	i := data["i"].(float64)
	j := data["j"].(float64)
	c := entity.Chess{
		I: int8(i),
		J: int8(j),
	}
	over, gameOverDTO, room, err := service.MakeStep(rid, c)
	if err != nil {
		SendErr(s, err)
		return
	}
	Send2Room(room, msg)
	if over {
		SendGameOver(room, gameOverDTO)
		Send2Room(room, &dto.Message{
			Code: constants.SetReady,
			Data: *room,
		})
	}
}

func RetractStep(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	data := msg.Data.(map[string]interface{})
	rid := data["rid"].(string)
	consent := int(data["consent"].(float64))
	opponentId, room, count, err := service.RetractStep(pid, rid, consent)
	if err != nil {
		SendErr(s, err)
	}
	if consent == 2 {
		data["count"] = count
		msg.Data = data
		Send2Room(room, msg)
	} else {
		Send2PId(opponentId, msg)
	}
}

func Surrender(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)

	rid, ok := msg.Data.(string)
	if !ok {
		err := fmt.Errorf("error: data is not string")
		SendErr(s, err)
		return
	}

	gameOverDTO, room, err := service.Surrender(pid, rid)
	if err != nil {
		SendErr(s, err)
		return
	}

	SendGameOver(room, gameOverDTO)
	Send2Room(room, &dto.Message{
		Code: constants.SetReady,
		Data: *room,
	})
}

func AskDraw(s *melody.Session, msg *dto.Message) {
	pid, _ := GetPId(s)
	data := msg.Data.(map[string]interface{})
	rid := data["rid"].(string)
	consent := int(data["consent"].(float64))
	opponentId, room, err := service.Draw(pid, rid, consent)
	if err != nil {
		SendErr(s, err)
	}

	if consent == 2 {
		SendGameOver(room, &dto.GameOverDTO{
			RId:   rid,
			Cause: "draw",
		})
		Send2Room(room, &dto.Message{
			Code: constants.SetReady,
			Data: *room,
		})
	} else {
		Send2PId(opponentId, msg)
	}
}
