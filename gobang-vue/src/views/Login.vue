<template>
    <div class="shadow-container">
        <h2 class="title">Welcome to Gobang Online!</h2>
        <el-divider/>
        <el-form class="form" @submit.native.prevent>
            <el-form-item>
                <el-input v-model="input" :placeholder="$t('lang.login.placeholder')" @keyup.enter.native="onLogin"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onLogin">{{$t('lang.login.login')}}</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    import {playerRename} from "@/websocket/send-api";

    export default {
        name: "Login",
        data() {
            return {
                input: ''
            }
        },
        methods: {
            onLogin() {
                if (this.input !== '') {
                    playerRename(this.input)
                    this.$router.push('/game')
                }
            }
        }
    }
</script>

<style scoped>
    .shadow-container {
        margin: 180px auto;
        width: 25%;
        text-align: center;
        min-width: 400px;
    }
    .form {
        width: 60%;
        padding: 3% 1% 1% 20%;
    }
</style>