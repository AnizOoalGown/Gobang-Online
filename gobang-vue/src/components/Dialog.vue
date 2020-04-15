<template>
    <div class="container">
        <div class="header">
            <span class="title">{{$t('lang.dialog.dialog')}}</span>
        </div>
        <div class="scrollbar">
            <div class="line" v-for="(item, index) in dialog" :key="index">
                <span style="font-size: xx-small">{{item.time}}</span>
                <br>
                <span>{{ item.from }}: {{ item.content }}</span>
            </div>
        </div>
        <div>
            <el-input size="mini" v-model="input" class="input" @keyup.enter.native="onSend"/>
            <el-button size="mini" @click="onSend" class="send-button">{{$t('lang.dialog.send')}}</el-button>
        </div>
    </div>
</template>

<script>
    import {hallChat, roomChat} from "../websocket/send-api";

    export default {
        name: "Dialog",
        props: ['roomId'],
        data() {
            return {
                input: '',
                dialog: []
            }
        },
        methods: {
            onSend() {
                if (this.input !== '') {
                    if (this.roomId === 'hall') {
                        hallChat(this.input)
                    }
                    else {
                        roomChat(this.$store.getters.player.name, this.input, this.roomId)
                    }
                    this.input = ''
                }
            },
            addDialogMsg(msg) {
                if (this.dialog.length >= 10) {
                    this.dialog.splice(0, 1)
                }
                this.dialog.push(msg)
            }
        },
        computed: {
            dialogMsg() {
                if (this.roomId === 'hall') {
                    return this.$store.getters.hallDialogMsg
                }
                else {
                    return this.$store.getters.roomChatDTO
                }
            }
        },
        watch: {
            dialogMsg(newMsg) {
                if (this.roomId !== 'hall' && newMsg.rid !== this.roomId) {
                    return
                }
                this.addDialogMsg(newMsg)
            }
        },
        mounted() {
            if (this.roomId === 'hall') {
                let date = new Date()
                this.$store.dispatch('setHallDialogMsg', {
                    time: date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDay() + ' '
                    + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds(),
                    from: '系统消息',
                    content: 'Settings里设置语言'
                })
            }
        }
    }
</script>

<style scoped>
    .container {
        margin-top: 3%;
    }
    .scrollbar {
        height: calc(30vh);
        min-height: 180px;
    }
    .header {
        border-bottom: 1px solid lightgrey;
        padding: 2% 0 2% 5%;
    }
    .title {
        padding-top: 100px;
    }
    .line {
        margin-top: 5px;
        margin-left: 3%;
        font-size: small;
    }
    .input {
        width: 80%;
    }
    .send-button {
        width: 20%;
    }
</style>