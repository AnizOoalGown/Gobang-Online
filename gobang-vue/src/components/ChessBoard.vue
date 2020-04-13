<template>
    <el-container>
        <el-header style="height: 20px; display: block; text-align: center">
            <span v-if="!started" style="font-size: large; font-weight: bold">{{title}}</span>
            <el-button size="mini" style="float: right" @click="onExit()">{{$t('lang.chessboard.button.exit')}}</el-button>
        </el-header>
        <el-main>
            <canvas :id="roomId" @click="onClick">Your browser doesn't support canvas</canvas>
        </el-main>
        <el-footer align="center" style="height: 20px">
            <el-button size="mini" @click="onRetract()" :disabled="buttonDisabled">{{$t('lang.chessboard.button.retract')}}</el-button>
            <el-button size="mini" @click="onSurrender()" :disabled="buttonDisabled">{{$t('lang.chessboard.button.surrender')}}</el-button>
            <el-button size="mini" @click="onDraw()" :disabled="buttonDisabled">{{$t('lang.chessboard.button.draw')}}</el-button>
        </el-footer>
    </el-container>
</template>

<script>
    import constant from "../constants/color";
    import {askDraw, leaveRoom, makeStep, retractStep, surrender} from "../websocket/send-api";
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
                title: this.$t('lang.chessboard.message.clickReady'),
                started: false,
                waitResponse: false
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
            drawChess(i, j, color) {
                let context = this.context
                context.beginPath()
                context.arc(m +i * d, m + j * d, r, 0, 2 * Math.PI)
                context.closePath()
                if (color === constant.black) {
                    context.fillStyle = '#000000'
                    context.fill()
                }
                else if (color === constant.white) {
                    context.stroke()
                    context.fillStyle = '#FFFFFF'
                    context.fill()
                }
            },
            drawLabel(i, j) {
                let context = this.context
                context.beginPath()
                context.arc(m +i * d, m + j * d, r / 3, 0, 2 * Math.PI)
                context.closePath()
                context.fillStyle = '#FFB90F'
                context.fill()
            },
            labelLastStep() {
                let length = this.steps.length
                if (length > 0) {
                    let lastStep = this.steps[length - 1]
                    this.drawLabel(lastStep.i, lastStep.j)
                    if (length > 1) {
                        let index = length - 2
                        let c = this.steps[index]
                        this.drawChess(c.i, c.j, index % 2)
                    }
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
                this.steps.push({i, j})
                this.drawChess(i, j, 1 - this.turn)
                this.labelLastStep()
            },
            onClick(e) {
                if (this.chessboardDisabled) {
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
                this.waitResponse = true
                this.$message.info(this.$t('lang.chessboard.message.retractSent'))
                retractStep(this.roomId, 1)
            },
            onExit() {
                this.$store.dispatch('removeTab', this.roomId)
                leaveRoom(this.roomId)
                setPlayerStatus("leisure")
            },
            onSurrender() {
                surrender(this.roomId)
            },
            onDraw() {
                this.waitResponse = true
                this.$message.info(this.$t('lang.chessboard.message.drawSent'))
                askDraw(this.roomId, 1)
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
            chessboardDisabled() {
                return this.myColor !== this.turn || this.waitResponse
            },
            buttonDisabled() {
                return this.myColor === -1 || this.waitResponse
            },
            step() {
                return this.$store.getters.step
            },
            gameOverDTO() {
                return this.$store.getters.gameOverDTO
            },
            chessboard() {
                return this.$store.getters.chessboard
            },
            drawDTO() {
                return this.$store.getters.drawDTO
            },
            retractDTO() {
                return this.$store.getters.retractDTO
            }
        },
        watch: {
            matchDetails(details) {
                if (details.roomId === this.roomId) {
                    if (details.host.ready && details.challenger.ready) {
                        this.started = true
                        this.steps = []
                        this.context.clearRect(0, 0, side, side)
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
                    this.started = false
                    if (gameOverDTO.cause === 'five') {
                        this.title = this.$t('lang.chessboard.message.over.five') + gameOverDTO.winner.name
                    }
                    else if (gameOverDTO.cause === 'escape') {
                        this.title = gameOverDTO.loser.name + this.$t('lang.chessboard.message.over.escape') + gameOverDTO.winner.name
                    }
                    else if (gameOverDTO.cause === 'surrender') {
                        this.title = gameOverDTO.loser.name + this.$t('lang.chessboard.message.over.surrender') + gameOverDTO.winner.name
                    }
                    else if (gameOverDTO.cause === 'draw') {
                        this.waitResponse = false
                        this.title = this.$t('lang.chessboard.message.over.draw')
                    }
                    else {
                        this.title = ''
                    }
                    this.$alert(this.title)
                    this.myColor = -1
                }
            },
            chessboard(chessboard) {
                if (chessboard.roomId === this.roomId) {
                    this.steps = chessboard.steps
                    this.steps.forEach((step, index) => {
                        this.drawChess(step.i, step.j, index % 2)
                    })
                    this.labelLastStep()
                }
            },
            drawDTO(drawDTO) {
                if (drawDTO.rid === this.roomId) {
                    if (drawDTO.consent === 1) {
                        this.$confirm(this.$t('lang.chessboard.message.askDraw.info'), this.$t('lang.chessboard.message.askDraw.title'), {
                            confirmButtonText: this.$t('lang.pop.yes'),
                            cancelButtonText: this.$t('lang.pop.no'),
                            type: 'info'
                        }).then(() => {
                            askDraw(this.roomId, 2)
                        }).catch(() => {
                            askDraw(this.roomId, 0)
                        })
                    }
                    else if (drawDTO.consent === 0) {
                        this.waitResponse = false
                        this.$alert(this.$t('lang.chessboard.message.rejectDraw.info'), this.$t('lang.chessboard.message.rejectDraw.title'))
                    }
                }
            },
            retractDTO(retractDTO) {
                if (retractDTO.rid === this.roomId) {
                    if (retractDTO.consent === 1) {
                        this.$confirm(this.$t('lang.chessboard.message.askRetract.info'), this.$t('lang.chessboard.message.askRetract.title'), {
                            confirmButtonText: this.$t('lang.pop.yes'),
                            cancelButtonText: this.$t('lang.pop.no'),
                            type: 'info'
                        }).then(() => {
                            retractStep(this.roomId, 2)
                        }).catch(() => {
                            retractStep(this.roomId, 0)
                        })
                    }
                    else if (retractDTO.consent === 0) {
                        this.waitResponse = false
                        this.$alert(this.$t('lang.chessboard.message.rejectRetract.info'), this.$t('lang.chessboard.message.rejectRetract.title'))
                    }
                    else if (retractDTO.consent === 2) {
                        this.$message.info(this.$t('lang.chessboard.message.agreeRetract.info'))
                        for (let i = 0; i < retractDTO.count; i++) {
                            let lastIndex = this.steps.length - 1
                            let step = this.steps[lastIndex]
                            this.removeChess(step.i, step.j)
                            this.steps.splice(lastIndex, 1)
                        }
                        this.labelLastStep()
                        this.waitResponse = false
                    }
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
        pointer-events: auto;
        background-image: url("../assets/images/chessboard.jpg");
        background-size: 100% 100%;
    }
</style>