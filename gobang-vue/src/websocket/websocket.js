import code from "../constants/msg-code"
import receive from "./receive-api"

const url = 'ws://localhost:8080/ws'
let ws = {}

export function initWebSocket() {
    ws = new WebSocket(url)
    ws.onmessage = onMessage
    ws.onerror = onError
}

function onMessage(e) {
    let msg = JSON.parse(e.data)
    let data = msg.data
    switch(msg.code) {
        case code.HallChat:
            receive.hallChat(data)
            break
        case code.GetPlayer:
            receive.getPlayer(data)
            break
        default:
            break
    }
}

function onError() {
    alert('WebSocket connection to \'' + url + '\' failed')
}

export function send(code, data) {
    ws.send(JSON.stringify({
        code,
        data
    }))
}

