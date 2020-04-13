<template>
    <div class="container">
        <div class="header">
            <span class="title">{{$t('lang.roomTable.chessRoomList')}}</span>
            <el-button size="mini" @click="onCreateRoom" style="float: right; margin-right: 10px">{{$t('lang.roomTable.createRoom')}}</el-button>
            <el-button size="mini" @click="onRefresh" style="float: right; margin-right: 10px">{{$t('lang.roomTable.refresh')}}</el-button>
        </div>

        <el-dialog :title="dialogTitle" :visible.sync="dialog.visible" width="25%">
            <el-form>
                <el-form-item :label="$t('lang.roomTable.dialog.color')">
                    <el-select v-model="dialog.color">
                        <el-option :label="$t('lang.color.black')" :value="0"></el-option>
                        <el-option :label="$t('lang.color.white')" :value="1"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="onDialogConfirm" size="small">{{$t('lang.pop.confirm')}}</el-button>
                <el-button @click="dialog.visible = false" size="small">{{$t('lang.pop.cancel')}}</el-button>
            </div>
        </el-dialog>

        <el-table :data="rooms" class="scrollbar" :show-header="false">
            <el-table-column label="room">
                <div slot-scope="props" style="display: inline-block">
                    <div :class="getChessClass(props.row.host)"></div>
                    <div style="float: left">{{props.row.host.name}} vs {{props.row.challenger.name}}</div>
                    <div :class="getChessClass(props.row.challenger)"></div>
                </div>
            </el-table-column>
            <el-table-column align="right">
                <template slot-scope="scope">
                    <el-button size="mini" @click="onChallenge(scope.row)" v-if="challengeShow(scope.row)">{{$t('lang.roomTable.challenge')}}</el-button>
                    <el-button size="mini" @click="onSpectate(scope.row)">{{$t('lang.roomTable.spectate')}}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import {createRoom, enterRoom, getRooms} from "@/websocket/send-api"
    import color from "@/constants/color"
    import {setPlayerStatus} from "../websocket/send-api";

    export default {
        name: "RoomTable",
        data() {
            return {
                dialog: {
                    visible: false,
                    color: color.black
                }
            }
        },
        methods: {
            onDialogConfirm() {
                createRoom(this.dialog.color)
                this.dialog.visible = false
            },
            onCreateRoom() {
                this.dialog.visible = true
            },
            onChallenge(room) {
                enterRoom(room.id, "challenger")
            },
            onSpectate(room) {
                enterRoom(room.id, "spectator")
                if (this.$store.getters.player.status === "leisure") {
                    setPlayerStatus("spectating")
                }
            },
            onRefresh() {
                getRooms()
            },
            getChessClass(player) {
                if (player.id !== '' && player.color === color.black) {
                    return 'black'
                }
                else if (player.id !== '' && player.color === color.white) {
                    return 'white'
                }
                return ''
            },
            challengeShow(room) {
                return room.challenger.id === ''
            }
        },
        computed: {
            dialogTitle() {
                return this.$t('lang.roomTable.dialog.title')
            },
            rooms() {
                return this.$store.getters.rooms
            }
        },
        mounted() {
            getRooms()
        }
    }
</script>

<style scoped>
    .container {
        margin-left: 2%;
        margin-right: 2%;
    }
    .scrollbar {
        height: calc(70vh);
        min-height: 415px;
    }
    .header {
        border-bottom: 1px solid lightgrey;
        padding: 2% 0 2% 5%;
    }
    .title {
        padding-top: 100px;
    }
</style>