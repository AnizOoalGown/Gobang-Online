import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
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
        activeTabKey: state => state.activeTabKey
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
        }
    }
})

export default store