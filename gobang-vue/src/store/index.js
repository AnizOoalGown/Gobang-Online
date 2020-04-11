import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        player: {
            id: "x",
            name: "x",
            status: "x"
        },
        playerTable: {
            roomId: '',
            players: []
        },
        step: {},
        chessboard: {
            roomId: '',
            steps: []
        },
        matchDetails: {
            roomId: '',
            host: {
                name: 'Tom',
                color: 'black',
                turn: true,
                roomStatus: 'ready'
            },
            challenger: {
                name: 'Ann',
                color: 'white',
                turn: false,
                roomStatus: 'unready'
            }
        },
        roomChatDTO: {},
        gameOverDTO: {},
        drawDTO: {},
        retractDTO: {},
        hallDialogMsg: {
            "time": "2020-03-31 16:43:00",
            "from": "sys",
            "content": "Welcome to Gobang Online!"
        },
        activeTabKey: 'hall',
        tabs: [{
            roomId: 'hall',
            title: '【Hall】',
            type: 'hall'
        }],
        rooms: []
    },
    getters: {
        tabs: state => state.tabs,
        activeTabKey: state => state.activeTabKey,
        player: state => state.player,
        playerTable: state => state.playerTable,
        step: state => state.step,
        chessboard: state => state.chessboard,
        matchDetails: state => state.matchDetails,
        roomChatDTO: state => state.roomChatDTO,
        gameOverDTO: state => state.gameOverDTO,
        drawDTO: state => state.drawDTO,
        retractDTO: state => state.retractDTO,
        hallDialogMsg: state => state.hallDialogMsg,
        rooms: state => state.rooms
    },
    mutations: {
        addTab(state, room) {
            let newTab = {
                roomId: room.id,
                title: '【Room】'+ room.host.name + ' vs ' + room.challenger.name,
                type: 'room',
            }
            for (let i in state.tabs) {
                if (state.tabs[i].roomId === room.id) {
                    state.tabs[i].title = newTab.title
                    return
                }
            }
            state.tabs.push(newTab)
        },
        removeTab(state, roomId) {
            let tabIndex = -1
            state.tabs.forEach((tab, i) => {
                if (tab.roomId === roomId) {
                    tabIndex = i
                }
            })
            if (tabIndex !== -1) {
                state.tabs.splice(tabIndex, 1)
                state.activeTabKey = 'hall'
            }
        },
        changeTab(state, roomId) {
            state.activeTabKey = roomId
        },
        setPlayer(state, player) {
            state.player = player
        },
        playerRename(state, name) {
            state.player.name = name
        },
        setPlayerTable(state, playerTable) {
            state.playerTable = playerTable
        },
        setStep(state, step) {
            state.step = step
        },
        setChessboard(state, chessboard) {
            state.chessboard = chessboard
        },
        setMatchDetails(state, matchDetails) {
            state.matchDetails = matchDetails
        },
        setRoomChatDTO(state, roomChatDTO) {
            state.roomChatDTO = roomChatDTO
        },
        setGameOverDTO(state, gameOverDTO) {
            state.gameOverDTO = gameOverDTO
        },
        setDrawDTO(state, drawDTO) {
            state.drawDTO = drawDTO
        },
        setRetractDTO(state, retractDTO) {
            state.retractDTO = retractDTO
        },
        setHallDialogMsg(state, dialogMsg) {
            state.hallDialogMsg = dialogMsg
        },
        setRooms(state, rooms) {
            state.rooms = rooms
        },
        addRoom(state, room) {
            state.rooms.push(room)
        },
        delRoomById(state, roomId) {
            let i
            for (i in state.rooms) {
                if (state.rooms[i].id === roomId) {
                    break
                }
            }
            state.rooms.splice(i, 1)
        }
    },
    actions: {
        addTab({commit}, room) {
            commit('addTab', room)
        },
        removeTab({commit}, roomId) {
            commit('removeTab', roomId)
        },
        changeTab({commit}, roomId) {
            commit('changeTab', roomId)
        },
        setPlayer({commit}, player) {
            commit('setPlayer', player)
        },
        playerRename({commit}, name) {
            commit('playerRename', name)
        },
        setPlayerTable({commit}, playerTable) {
            commit('setPlayerTable', playerTable)
        },
        setStep({commit}, step) {
            commit('setStep', step)
        },
        setChessboard({commit}, chessboard) {
            commit('setChessboard', chessboard)
        },
        setMatchDetails({commit}, matchDetails) {
            commit('setMatchDetails', matchDetails)
        },
        setRoomChatDTO({commit}, roomChatDTO) {
            commit('setRoomChatDTO', roomChatDTO)
        },
        setGameOverDTO({commit}, gameOverDTO) {
            commit('setGameOverDTO', gameOverDTO)
        },
        setDrawDTO({commit}, drawDTO) {
            commit('setDrawDTO', drawDTO)
        },
        setRetractDTO({commit}, retractDTO) {
            commit('setRetractDTO', retractDTO)
        },
        setHallDialogMsg({commit}, dialogMsg) {
            commit('setHallDialogMsg', dialogMsg)
        },
        setRooms({commit}, rooms) {
            commit('setRooms', rooms)
        },
        addRoom({commit}, room) {
            commit('addRoom', room)
        },
        delRoomById({commit}, roomId) {
            commit('delRoomById', roomId)
        }
    }
})

export default store