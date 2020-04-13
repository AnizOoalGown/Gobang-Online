<template>
    <div class="container">
        <div class="header">
            <span class="title">{{title}}</span>
        </div>
        <div class="scrollbar">
            <el-table :data="players" :show-header="false" size="mini">
                <el-table-column prop="name"/>
                <el-table-column align="right">
                    <template slot-scope="scope">
                        {{$t('lang.status.' + scope.row.status)}}
                    </template>
                </el-table-column>
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
                players: []
            }
        },
        computed: {
            playerTable() {
                return this.$store.getters.playerTable
            },
            title() {
                if (this.roomId === 'hall') {
                    return this.$t('lang.playerTable.playerList')
                }
                else {
                    return this.$t('lang.playerTable.spectatorList')
                }
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