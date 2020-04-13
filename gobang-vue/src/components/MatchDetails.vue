<template>
    <div class="container">
        <div class="header">
            <span class="title">{{$t('lang.matchDetails.matchDetails')}}</span>
        </div>
        <div class="scrollbar">
            <div class="role">
                {{host.name}}
                <el-button style="margin-left: 10%" size="medium" v-if="readyBtnShow(host)" @click="onReady">{{$t('lang.matchDetails.ready')}}</el-button>
                <span style="margin-left: 10%" v-else-if="host.ready===false">{{$t('lang.matchDetails.unready')}}</span>
                <span style="margin-left: 10%" v-else-if="host.ready===true">{{$t('lang.matchDetails.already')}}</span>
            </div>
            <el-form>
                <el-form-item :label="$t('lang.color.color')">
                    <div :class="getChessClass(host.color)"/>
                </el-form-item>
                <el-form-item label="role">{{$t('lang.role.host')}}</el-form-item>
            </el-form>
            <el-divider/>
            <div class="role">
                {{challenger.name}}
                <el-button style="margin-left: 10%" size="medium" v-if="readyBtnShow(challenger)" @click="onReady">{{$t('lang.matchDetails.ready')}}</el-button>
                <span style="margin-left: 10%" v-else-if="challenger.ready===false">{{$t('lang.matchDetails.unready')}}</span>
                <span style="margin-left: 10%" v-else-if="challenger.ready===true">{{$t('lang.matchDetails.already')}}</span>
            </div>
            <el-form>
                <el-form-item :label="$t('lang.color.color')">
                    <div :class="getChessClass(challenger.color)"/>
                </el-form-item>
                <el-form-item label="role">{{$t('lang.role.challenger')}}</el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
    import color from "@/constants/color";
    import {setPlayerStatus, setReady} from "../websocket/send-api";

    export default {
        name: "MatchDetails",
        props: ['roomId'],
        data() {
            return {
                host: {},
                challenger: {}
            }
        },
        methods: {
            getChessClass(playerColor) {
                if (playerColor === color.black) {
                    return 'black'
                }
                else if (playerColor === color.white) {
                    return 'white'
                }
                return ''
            },
            readyBtnShow(playerDetails) {
                return playerDetails.id === this.$store.getters.player.id && playerDetails.ready === false
            },
            onReady() {
                setReady(this.roomId, true)
            },
            isPlayer() {
                let id = this.$store.getters.player.id
                return id === this.host.id || id === this.challenger.id
            }
        },
        computed: {
            matchDetails() {
                return this.$store.getters.matchDetails
            },
            playerStatus() {
                return this.$store.getters.player.status
            },
            started() {
                return this.host.ready && this.challenger.ready
            }
        },
        watch: {
            matchDetails(details) {
                if (details.roomId === this.roomId) {
                    this.host = details.host
                    this.challenger = details.challenger
                }
            },
            playerStatus(newStatus) {
                if (newStatus !== "chessing" && this.started && this.isPlayer()) {
                    setPlayerStatus("chessing")
                }
                else if (newStatus === "leisure" && !this.isPlayer()) {
                    setPlayerStatus("spectating")
                }
            },
            started(newStarted) {
                if(this.isPlayer()) {
                    if (newStarted) {
                        setPlayerStatus("chessing")
                    }
                    else {
                        setPlayerStatus("leisure")
                    }
                }
            }
        }
    }
</script>

<style scoped>
    .header {
        border-bottom: 1px solid lightgrey;
        padding: 2% 0 2% 5%;
    }
    .title {
        padding-top: 100px;
    }
    .scrollbar {
        height: 73vh;
        min-height: 420px;
    }
    .el-form-item {
        margin-left: 10%;
        margin-bottom: 0px;
    }
    .black {
        margin: 13px 0 0 5px;
    }
    .white {
        margin: 13px 0 0 5px;
    }
    .role {
        margin: 5% 0 5% 10%;
        font-size: large;
    }
</style>