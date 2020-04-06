<template>
    <div class="container">
        <div class="header">
            <span class="title">Dialog</span>
        </div>
        <div class="scrollbar">
            <div class="line" v-for="(item, index) in dialog" :key="index">
                <span style="font-size: xx-small">{{item.time}}</span>
                <br>
                <span>{{ item.from }}: {{ item.content }}</span>
            </div>
        </div>
        <div>
            <el-input size="mini" v-model="input" class="input"/>
            <el-button size="mini" @click="onSend" class="send-button">Send</el-button>
        </div>
    </div>
</template>

<script>
    import {getHallDialog, hallChat} from "../websocket/send-api";

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
                    // this.dialog.push({
                    //     from: 'Me',
                    //     content: this.input
                    // })
                    if (this.roomId === 'hall') {
                        hallChat(this.input)
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
                return {}
            }
        },
        watch: {
            dialogMsg(newMsg) {
                this.addDialogMsg(newMsg)
            }
        },
        mounted() {
            if (this.roomId === 'hall') {
                getHallDialog()
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