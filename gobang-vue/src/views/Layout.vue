<template>
    <div class="shadow-container">
        <el-container>
            <el-header>
                <el-menu router mode="horizontal" :default-active="$route.path">
                    <template v-for="item in this.$router.options.routes[2].children">
                        <el-menu-item :index="item.path" :key="item.path">
                            <span slot="title">{{$t('lang.layout.' + item.name)}}</span>
                        </el-menu-item>
                    </template>
                    <span class="brand-name">{{$t('lang.layout.title')}}</span>
                </el-menu>
            </el-header>
            <el-main>
                <keep-alive>
                    <router-view/>
                </keep-alive>
            </el-main>
        </el-container>
    </div>

</template>

<script>
    export default {
        name: "Layout",
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    .shadow-container {
        height: calc(97vh);
        min-height: 600px;
        min-width: 1200px;
    }
    .brand-name {
        font-size: x-large;
        font-style: italic;
        float: right;
        margin-top: 1%;
    }
</style>