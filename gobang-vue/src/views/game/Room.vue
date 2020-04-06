<template>
    <el-row>
        <el-col :span="5">
            <PlayerTable :roomId="roomId"/>
            <Dialog :roomId="roomId"/>
        </el-col>
        <el-col :span="14">
            <ChessBoard :roomId="roomId"/>
        </el-col>
        <el-col :span="5">
            <MatchDetails :roomId="roomId"/>
        </el-col>
    </el-row>
</template>

<script>
    import PlayerTable from "@/components/PlayerTable"
    import Dialog from "@/components/Dialog"
    import MatchDetails from "@/components/MatchDetails"
    import ChessBoard from "@/components/ChessBoard"

    export default {
        name: "Room",
        components: {ChessBoard, MatchDetails, Dialog, PlayerTable},
        props: ['roomId'],
        data() {
            return {
                // roomId: '123'
            }
        },
        methods: {
            test() {
                console.log('test')
            }
        }
    }
</script>

<style scoped>

</style>