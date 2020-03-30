<template>
    <div class="container">
        <div class="header">
            <span class="title">Player List</span>
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
    export default {
        name: "PlayerTable",
        data() {
            return {
                players: [{
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, {
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, {
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, {
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, {
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, {
                    name: 'Tom',
                    status: 'playing'
                }, {
                    name: 'Bob',
                    status: 'leisure'
                }, {
                    name: 'Ann',
                    status: 'spectating'
                }, ]
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