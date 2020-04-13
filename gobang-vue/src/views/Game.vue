<template>
    <el-tabs v-model="activeTabKey">
            <el-tab-pane v-for="item in tabs" :key="item.roomId" :label="setTitle(item.title)" :name="item.roomId">
                <Hall v-if="item.type === 'hall'" :roomId="item.roomId"/>
                <Room v-else-if="item.type === 'room'" :roomId="item.roomId"/>
            </el-tab-pane>
    </el-tabs>
</template>

<script>
    import Hall from "@/views/game/Hall"
    import Room from "@/views/game/Room"

    export default {
        name: "Game",
        components: {Hall, Room},
        data() {
            return {
                // activeTabKey: '',
                // tabs: [{
                //     title: 'Hall',
                //     type: 'hall'
                // }, {
                //     title: 'Room A vs B',
                //     type: 'room'
                // }, {
                //     title: 'Room C vs D',
                //     type: 'room'
                // }]
            }
        },
        methods: {
            setTitle(title) {
                if (this.$i18n.locale === 'zh') {
                    title = title.replace(/Room/g,"房间")
                    title = title.replace(/Hall/g,"大厅")
                }
                return title
            }
        },
        computed: {
            tabs() {
                return this.$store.getters.tabs
            },
            activeTabKey: {
                get() {
                    return this.$store.getters.activeTabKey
                },
                set(roomId) {
                    this.$store.dispatch('changeTab', roomId)
                }
            }
        }
    }
</script>

<style scoped>

</style>