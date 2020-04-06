<template>
    <div class="container">
        <div class="header">
            <span class="title">Match Details</span>
        </div>
        <div class="scrollbar">
            <div class="role">
                {{host.name}}
                <el-button style="margin-left: 10%" size="medium" v-if="host.roomStatus==='unready'">ready</el-button>
                <span style="margin-left: 10%" v-else-if="host.roomStatus==='ready'">ready!</span>
            </div>
            <el-form>
                <el-form-item label="color">
                    <div :class="getChessClass(host.color)"/>
                </el-form-item>
                <el-form-item label="role">host</el-form-item>
                <el-form-item label="turn">{{host.turn}}</el-form-item>
            </el-form>
            <el-divider/>
            <div class="role">
                {{challenger.name}}
                <el-button style="margin-left: 10%" size="medium" v-if="challenger.roomStatus==='unready'">ready</el-button>
                <span style="margin-left: 10%" v-else-if="challenger.roomStatus==='ready'">ready!</span>
            </div>
            <el-form>
                <el-form-item label="color">
                    <div :class="getChessClass(challenger.color)"/>
                </el-form-item>
                <el-form-item label="role">challenger</el-form-item>
                <el-form-item label="turn">{{challenger.turn}}</el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
    import color from "@/constants/color";

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
            }
        },
        computed: {
            matchDetails() {
                return this.$store.getters.matchDetails
            }
        },
        watch: {
            matchDetails(details) {
                if (details.roomId === this.roomId) {
                    this.host = details.host
                    this.challenger = details.challenger
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