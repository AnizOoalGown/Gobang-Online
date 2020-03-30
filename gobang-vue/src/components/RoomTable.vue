<template>
    <div class="container">
        <div class="header">
            <span class="title">Chess Room List</span>
            <el-button size="mini" @click="onCreateRoom" style="float: right; margin-right: 10px">create room</el-button>
        </div>
        <el-table :data="rooms" class="scrollbar" :show-header="false">
            <el-table-column label="room">
                <div slot-scope="props" style="display: inline-block">
                    <div :class="props.row.player1.color"></div>
                    <div style="float: left">{{props.row.player1.name}} vs {{props.row.player2.name}}</div>
                    <div :class="props.row.player2.color"></div>
                </div>
            </el-table-column>
            <el-table-column align="right">
                <template slot-scope="scope">
                    <el-button size="mini" @click="onChallenge(scope.row)">challenge</el-button>
                    <el-button size="mini" @click="onSpectate(scope.row)">spectate</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    export default {
        name: "RoomTable",
        data() {
            return {
                rooms: [{
                    player1: {
                        name: 'Tom',
                        color: 'white'
                    },
                    player2: {
                        name: 'May',
                        color: 'black'
                    },
                }, {
                    player1: {
                        name: 'Meow',
                        color: 'black'
                    },
                    player2: {
                        name: 'Aoo',
                        color: 'white'
                    },
                }]
            }
        },
        methods: {
            onCreateRoom() {

            },
            onChallenge(room) {
                console.log(room)
            },
            onSpectate(room) {
                console.log(room)
            }
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