<template>
  <el-container direction="vertical">
    <el-affix><lee-header></lee-header></el-affix>
    <el-main class="lee_box"> 此网站正在开发中,请敬请期待... </el-main>
    <lee-footer></lee-footer>
  </el-container>
</template>

<script setup lang="ts">
import { header as leeHeader, footer as leeFooter } from "@/components/system";
</script>

<style lang="less" scoped>
.lee_box {
  width: 100%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
