<template>
    <el-container>
        <el-header style="height: 20px">
            <el-button size="mini" style="float: right" @click="onExit()">Exit</el-button>
        </el-header>
        <el-main>
            <canvas :id="roomId" @click="onClick">Your browser doesn't support canvas</canvas>
        </el-main>
        <el-footer align="center" style="height: 20px">
            <el-button size="mini" @click="onRetract()">retract</el-button>
            <el-button size="mini">surrender</el-button>
            <el-button size="mini" @click="onDraw()">draw</el-button>

        </el-footer>
    </el-container>

</template>

<script>
    import color from "../constants/color";
    import {leaveRoom, makeStep} from "../websocket/send-api";
    import {setPlayerStatus} from "../websocket/send-api";

    // let canvas;
    // let context
    //canvas边长
    const side = 495
    // 棋盘格子边长
    const d = 33
    //棋盘格子距离棋盘边缘的距离
    const m = d / 2
    //棋子半径
    const r = 15
    // let turn = constants.black
    // let steps = []

    export default {
        name: "ChessBoard",
        props: ['roomId'],
        data() {
            // let canvas = document.getElementById(this.roomId);
            // let context = canvas.getContext('2d');
            return {
                // canvas,
                context: {},
                steps: [],
                myColor: -1,
            }
        },
        methods: {
            initCanvas() {
                let canvas = document.getElementById(this.roomId);
                // let canvas = this.canvas
                canvas.width = side
                canvas.height = side
                this.context = canvas.getContext('2d');
                // let image = new Image()
                // image.src = require('../assets/images/chessboard.jpg')
                // image.onload = () => {
                //     console.log('image loaded success')
                //     context.drawImage(image, 0, 0, side, side)
                //     // context.strokeStyle = '#B9B9B9';
                //     context.lineWidth = 1
                //     this.drawChessboard()
                //     this.drawChess(7, 7, color.black)
                // }
            },
            drawChessboard() {
                for (let i = 0; i < 15; i++) {
                    let context = this.context
                    context.moveTo(m + i * d , m);
                    context.lineTo(m + i * d , side - m);
                    context.stroke();
                    context.moveTo(m , m + i * d);
                    context.lineTo(side - m, m + i * d);
                    context.stroke();
                }
            },
            drawChess(i, j, turn) {
                let context = this.context
                context.beginPath()
                context.arc(m +i * d, m + j * d, r, 0, 2 * Math.PI)
                context.closePath()
                if (turn === color.black) {
                    context.fillStyle = '#000000'
                    context.fill()
                }
                else if (turn === color.white) {
                    context.stroke()
                    context.fillStyle = '#FFFFFF'
                    context.fill()
                }
            },
            removeChess(i, j) {
                this.context.clearRect((i) * d, (j) * d, d, d);
            },
            hasStep(i, j) {
                for (let step of this.steps) {
                    if (step.i === i && step.j === j) {
                        return true
                    }
                }
                return false
            },
            chess(i, j) {
                this.drawChess(i, j, this.turn)
                this.steps.push({i, j})
            },
            onClick(e) {
                if (!this.myTurn) {
                    return
                }
                let x = e.offsetX
                let y = e.offsetY
                let i = Math.floor(x / d)
                let j = Math.floor(y / d)
                if (!this.hasStep(i, j)) {
                    this.chess(i, j)
                    makeStep(this.roomId, i, j)
                }
            },
            onRetract() {
                let lastIndex = this.steps.length - 1
                let step = this.steps[lastIndex]
                this.removeChess(step.i, step.j)
                this.steps.splice(lastIndex, 1)
            },
            onExit() {
                this.$store.dispatch('removeTab', this.roomId)
                leaveRoom(this.roomId)
                setPlayerStatus("leisure")
            },
            onDraw() {
                console.log(this.steps)
            }
        },
        mounted() {
            this.initCanvas()
            // 当调整窗口大小时重绘canvas
            // window.onresize = () => {
            //     this.initCanvas()
            // }
        },
        computed: {
            matchDetails() {
                return this.$store.getters.matchDetails
            },
            turn() {
                return this.steps.length % 2
            },
            myTurn() {
                return this.turn === this.myColor
            },
            step() {
                return this.$store.getters.step
            },
            gameOverDTO() {
                return this.$store.getters.gameOverDTO
            }
        },
        watch: {
            matchDetails(details) {
                if (details.roomId === this.roomId) {
                    if (details.host.ready && details.challenger.ready) {
                        let id = this.$store.getters.player.id
                        if (details.host.id === id) {
                            this.myColor = details.host.color
                            return
                        }
                        else if (details.challenger.id === id) {
                            this.myColor = details.challenger.color
                            return
                        }
                    }
                    this.myColor = -1
                }
            },
            step(step) {
                if (step.rid === this.roomId && !this.hasStep(step.i, step.j)) {
                    this.chess(step.i, step.j)
                }
            },
            gameOverDTO(gameOverDTO) {
                if (gameOverDTO.rid === this.roomId) {
                    this.$message.info(gameOverDTO.winner.name + ' win')
                    this.myColor = -1
                }
            }
        }

    }
</script>

<style scoped>
    canvas {
        display: block;
        margin: 0px auto;
        box-shadow: -2px -2px 2px #EFEFEF, 5px 5px 5px #B9B9B9;
        cursor: pointer;
        background-image: url("../assets/images/chessboard.jpg");
        background-size: 100% 100%;
    }
</style>