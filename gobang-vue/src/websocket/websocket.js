const url = 'ws://localhost:8080/ws'
let ws = {}

export function initWebSocket() {
    ws = new WebSocket(url)
    ws.onmessage = onMessage
}

export function onMessage(e) {
    let msg = JSON.parse(e.data)
    console.log(msg)
}

export function send(msg) {
    ws.send(JSON.stringify(msg))
}