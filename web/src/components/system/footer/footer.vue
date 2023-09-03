<template>
    <el-footer class="lee_footer">
        <!-- 网站备案信息 -->
        <div class="lee_filings">
            <a href="https://beian.miit.gov.cn/" target="_blank" rel="noopener">备案号：豫ICP备2023017033号-1</a>
        </div>
    </el-footer>
</template>

<script setup lang="ts">
</script>

<style lang="less" scoped>
.lee_footer {
    width: 100%;
    height: 6vh;
    display: flex;
    justify-content: center;
    background-color: @backgroundColor;

    .lee_filings {
        width: @versionNew;
        display: flex;
        justify-content: center;
        align-items: center;

        a {
            text-decoration: none;
            color: @textColor;
        }
    }
}
</style>