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
                <a href="https://github.com/Atidotes/chat" class="gitHub">
                    <span class="iconfont lee_github lee_icon"></span>
                </a>
                <span @click="handleFavorite" :class="favorite" class="iconfont lee_icon"></span>
            </div>
        </div>
    </el-header>

    <!-- 登录 -->
    <login ref="loginRef"></login>
</template>

<script setup lang="ts">
import {ref} from "vue"
import login from '@/components/login/login.vue';

/** 定义登录 */
const loginRef = ref<InstanceType<typeof login>>()
const favorite = ref("lee_notFavorite")

/** 登录 */
const handleDialog = () =>{
    if(loginRef.value != null) {
        loginRef.value.isDialog = true
    }
}

/** 收藏网站 */
const handleFavorite =() =>{
    favorite.value = favorite.value === "lee_notFavorite" ? "lee_favorite" : "lee_notFavorite"
}
</script>
        
<style lang="less" scoped>
.lee_header {
    width: 100%;
    height: 6vh;
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
            .login{
                cursor: pointer;
            }
            .gitHub {
                text-decoration: none;
            }
            .lee_favorite{
                color: #4285f4;
            }
        }
    }
}

.lee_icon {
    cursor: pointer;
    color: initial;
    margin-left: 1vw;
    font-size: 1.5vw;
}
</style>