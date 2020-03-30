<template>
    <div class="container">
        <div class="header">
            <span class="title">Dialog</span>
        </div>
        <div class="scrollbar">
            <p class="line" v-for="(item, index) in dialog" :key="index">{{ item.from }}: {{ item.content }}</p>
        </div>
        <div>
            <el-input size="mini" v-model="input" class="input"/>
            <el-button size="mini" @click="onSend" class="send-button">Send</el-button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Dialog",
        data() {
            return {
                input: '',
                dialog: [{
                    from: 'Me',
                    content: '111'
                }, {
                    from: 'Me',
                    content: '222'
                }]
            }
        },
        methods: {
            onSend() {
                if (this.input !== '') {
                    this.dialog.push({
                        from: 'Me',
                        content: this.input
                    })
                    this.input = ''
                }
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