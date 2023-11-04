<template>
  <el-header class="lee_header">
    <!-- 导航区 -->
    <div class="lee_nav">
      <!-- logo -->
      <div class="left">
        <span class="lee_logo">Atidotes</span>
      </div>

      <!-- nav -->
      <ul class="content">
        <li class="center item">首页</li>
        <li class="center item">我的</li>
      </ul>

      <!-- 功能 -->
      <div class="right">
        <span class="login" @click="handleDialog">登录</span>
        <el-switch
          @change="handleChangeSwitch"
          active-color="#4F4F4F"
          inactive-color="#363636"
          v-model="isDark"
          :active-action-icon="Moon"
          :inactive-action-icon="Sunny"
        />
        <a href="https://github.com/Atidotes/chat" class="gitHub">
          <span class="iconfont lee_github lee_icon"></span>
        </a>
      </div>
    </div>
  </el-header>

  <!-- 登录 -->
  <login ref="loginRef"></login>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { login } from "@/components/system";
import { Moon, Sunny } from "@element-plus/icons-vue";
import { setTheme } from "@/style/index";

/** 定义登录 */
let theme = localStorage.getItem("theme");
const isDark = ref(theme === "dark");
const loginRef = ref<InstanceType<typeof login>>();

/** 初始化数据 */
onMounted(() => {
  setTheme(true);
});

/** 登录 */
const handleDialog = () => {
  if (loginRef.value === undefined) return;
  loginRef.value.isDialog = true;
};

/** 改变switch */
const handleChangeSwitch = (isDark: boolean) => {
  isDark ? setTheme(false, "dark") : setTheme(false, "default");
};
</script>

<style lang="less" scoped>
.lee_header {
  width: 100%;
  height: 6.5vh;
  background-color: @backgroundColor;
  display: flex;
  justify-content: center;
  align-items: center;

  .lee_nav {
    width: @versionNew;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .left {
      .lee_logo {
        font-family: "Atidotes";
        font-size: 1.8vw;
      }
    }

    .content {
      width: auto;
      display: flex;
      justify-content: space-between;
      align-items: center;
      list-style: none;
      margin: 0;
      padding: 0;

      .center {
        color: rgb(157, 157, 231);
        cursor: pointer;
      }

      .item {
        margin-left: 1.5vw;

        &:first-child {
          margin-left: 0;
        }
      }
    }

    .right {
      width: 8vw;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: space-between;

      .login {
        cursor: pointer;
      }

      .gitHub {
        text-decoration: none;
      }
    }
  }
}

.lee_icon {
  cursor: pointer;
  color: initial;
  font-size: 1.5vw;
}
</style>
