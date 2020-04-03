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
        }, {
            roomId: 'qwe123',
            title: '【Room】A vs B',
            type: 'room'
        }, {
            roomId: 'xyz123',
            title: '【Room】C vs D',
            type: 'room'
        }]
    },
    getters: {
        tabs: state => state.tabs,
        activeTabKey: state => state.activeTabKey,
        player: state => state.player,
        hallDialogMsg: state => state.hallDialogMsg
    },
    mutations: {
        addTab(state) {
            let roomId = Math.random().toString(36).slice(-8)
            state.tabs.push({
                roomId,
                title: '【Room】X vs X',
                type: 'room'
            })
        },
        removeTab(state, roomId) {
            let tabIndex
            state.tabs.forEach((tab, i) => {
                if (tab.roomId === roomId) {
                    tabIndex = i
                }
            })
            state.tabs.splice(tabIndex, 1)
            state.activeTabKey = 'hall'
        },
        changeTab(state, roomId) {
            state.activeTabKey = roomId
        },
        setPlayer(state, player) {
            state.player = player
        },
        setHallDialogMsg(state, dialogMsg) {
            state.hallDialogMsg = dialogMsg
        }
    },
    actions: {
        addTab({commit}) {
            commit('addTab')
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
        setHallDialogMsg({commit}, dialogMsg) {
            commit('setHallDialogMsg', dialogMsg)
        }
    }
})

export default store