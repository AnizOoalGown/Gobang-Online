import store from "../store";

export default {
    fail(data) {
        alert(data)
    },
    hallChat(dialogMsg) {
        store.dispatch('setHallDialogMsg', dialogMsg)
    },
    async getHallDialog(dialog) {
        for (let dialogMsg of dialog) {
             await this.hallChat(dialogMsg)
        }
    },
    getRooms(rooms) {
        store.dispatch('setRooms', rooms)
    },
    createRoom(room) {
        store.dispatch('addRoom', room)
        this.enterRoom(room)
    },
    enterRoom(room) {
        store.dispatch('addTab', room).then(() => {
            store.dispatch('changeTab', room.id)
            store.dispatch(('setPlayerTable'), {
                roomId: room.id,
                players: room.spectators
            })
            store.dispatch('setMatchDetails', {
                roomId: room.id,
                host: room.host,
                challenger: room.challenger
            })
            store.dispatch('setChessboard', {
                roomId: room.id,
                steps: room.steps
            })
        })
    },
    leaveRoom(room) {
        store.dispatch('addTab', room).then(() => {
            store.dispatch('setPlayerTable', {
                roomId: room.id,
                players: room.spectators
            })
            store.dispatch('setMatchDetails', {
                roomId: room.id,
                host: room.host,
                challenger: room.challenger
            })

            store.dispatch('setChessboard', {
                roomId: room.id,
                steps: room.steps
            })
        })
    },
    delRoom(roomId) {
        store.dispatch('removeTab', roomId)
        store.dispatch('delRoomById', roomId)
    },
    roomChat(roomChatDTO) {
        store.dispatch('setRoomChatDTO', roomChatDTO)
    },

    getPlayer(player) {
        store.dispatch('setPlayer', player)
    },
    getPlayers(players) {
        let playerTable = {
            roomId: 'hall',
            players
        }
        store.dispatch('setPlayerTable', playerTable)
    },
    setReady(room) {
        store.dispatch('setMatchDetails', {
            roomId: room.id,
            host: room.host,
            challenger: room.challenger
        })
    },
    makeStep(step) {
        store.dispatch('setStep', step)
    },
    gameOver(gameOverDTO) {
        store.dispatch('setGameOverDTO', gameOverDTO)
    },
    askDraw(drawDTO) {
        store.dispatch('setDrawDTO', drawDTO)
    },
    retractStep(retractDTO) {
        store.dispatch('setRetractDTO', retractDTO)
    }
}

// export function