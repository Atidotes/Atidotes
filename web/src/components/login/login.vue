<template>
    <el-dialog title="登录" v-model="isDialog" width="35vw">
    
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const isDialog = ref(false)



/** 暴露属性 */
defineExpose({
    isDialog
})
</script>

<style lang="less" scoped>
</style>