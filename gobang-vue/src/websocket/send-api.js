import code from "../constants/msg-code"
import {send} from "./websocket"
import store from "@/store";

export function hallChat(content) {
    send(code.HallChat, content)
}

export function getHallDialog() {
    send(code.GetHallDialog, "")
}

export function getRooms() {
    send(code.GetRooms, "")
}

export function createRoom(color) {
    send(code.CreateRoom, color)
}

export function enterRoom(rid, role) {
    send(code.EnterRoom, {
        rid,
        role
    })
}

export function leaveRoom(rid) {
    send(code.LeaveRoom, rid)
}

export function delRoom(rid) {
    send(code.DelRoom, rid)
}

export function roomChat(from, content, rid) {
    send(code.RoomChat, {
        from,
        content,
        rid
    })
}

export function getPlayer() {
    send(code.GetPlayer, "")
}

export function getPlayers() {
    send(code.GetPlayers, "")
}

export function playerRename(name) {
    store.dispatch('playerRename', name)
    send(code.PlayerRename, name)
}

export function setPlayerStatus(status) {
    send(code.SetPlayerStatus, status)
}

export function setReady(rid, ready) {
    send(code.SetReady, {
        rid,
        ready
    })
}

export function makeStep(rid, i, j) {
    send(code.MakeStep, {
        rid,
        i,
        j
    })
}

export function surrender(rid) {
    send(code.Surrender, rid)
}

export function askDraw(rid, consent) {
    send(code.AskDraw, {
        rid,
        consent
    })
}

export function retractStep(rid, consent) {
    send(code.RetractStep, {
        rid,
        consent
    })
}