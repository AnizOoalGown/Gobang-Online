<template>
    <el-form>
        <el-form-item :label="$t('lang.settings.language')">
            <el-select v-model="lang" @change="onLang">
                <el-option v-for="item in langOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value">
                </el-option>
            </el-select>
        </el-form-item>
    </el-form>
</template>

<script>
    export default {
        name: "Settings",
        data() {
            return {
                lang: 'en',
                langOptions: [{
                    value: 'en',
                    label: 'English'
                }, {
                    value: 'zh',
                    label: '中文'
                }]
            }
        },
        methods: {
            onLang() {
                this.$i18n.locale = this.lang
            }
        }
    }
</script>

<style scoped>

</style>