<template>
    <div class="container">
        <el-image :src="src" fit="fill"/>
    </div>
</template>

<script>
    export default {
        name: "Advertisement",
        data() {
            return {
                src: require('../assets/images/ad2.jpg')
            }
        }
    }
</script>

<style scoped>
    .container {
        margin-top: 3%;
        height: calc(40vh);
        min-height: 180px;
    }
</style>