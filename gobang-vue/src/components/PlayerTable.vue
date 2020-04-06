<template>
    <div class="container">
        <div class="header">
            <span class="title">{{title}}</span>
        </div>
        <div class="scrollbar">
            <el-table :data="players" :show-header="false" size="mini">
                <el-table-column prop="name"/>
                <el-table-column prop="status" align="right"/>
            </el-table>
        </div>
    </div>
</template>

<script>
    import {getPlayers} from "@/websocket/send-api";

    export default {
        name: "PlayerTable",
        props: ['roomId'],
        data() {
            return {
                title: '',
                players: []
            }
        },
        computed: {
            playerTable() {
                return this.$store.getters.playerTable
            }
        },
        watch: {
            playerTable(newTable) {
                if (newTable.roomId === this.roomId) {
                    this.players = newTable.players
                }
            }
        },
        mounted() {
            getPlayers()
            if (this.roomId === 'hall') {
                this.title = 'Player List'
            }
            else {
                this.title = 'Spectator List'
            }
        }
    }
</script>

<style scoped>
    .scrollbar {
        height: calc(33vh);
        min-height: 180px;
    }
    .header {
        border-bottom: 1px solid lightgrey;
        padding: 2% 0 2% 5%;
    }
    .title {
        padding-top: 100px;
    }
</style>