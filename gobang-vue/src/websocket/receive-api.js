import store from "../store";

export default {
    hallChat(dialogMsg) {
        store.dispatch('setHallDialogMsg', dialogMsg)
    },
    getPlayer(player) {
        store.dispatch('setPlayer', player)
    }
}

// export function