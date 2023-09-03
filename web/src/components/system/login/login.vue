<template>
  <el-dialog
    title="登录"
    v-model="isDialog"
    width="50vw"
    custom-class="lee_login"
  >
    <div class="login">
      <!-- 关注公众号 -->
      <div class="follow_official_account">
        <img src="@/assets/system/login/officialAccount.jpg" alt="" />
        <div class="text">
          关注公众号<span class="illustrate">（即时提供反馈）</span>
        </div>
      </div>

      <!-- 表单登录 -->
      <div class="form_login">
        <!-- 选项卡 -->
        <el-radio-group v-model="loginType">
          <el-radio-button label="0">手机验证码登录</el-radio-button>
          <el-radio-button label="1">邮箱验证码登录</el-radio-button>
          <el-radio-button label="2">账号密码登录</el-radio-button>
        </el-radio-group>

        <!-- 登录表单 -->
        <el-form
          ref="loginRef"
          :model="loginForm"
          status-icon
          :rules="loginRules"
          class="login_form"
          label-width="3.75vw"
        >
          <!-- 手机验证码登录 -->
          <el-row v-if="loginType === '0'">
            <!-- 手机号 -->
            <el-col :span="24">
              <el-form-item label="手机号">
                <el-input placeholder="请填写手机号"></el-input>
              </el-form-item>
            </el-col>

            <!-- 发送验证码 -->
            <el-col :span="24">
              <el-row :gutter="24">
                <el-col :span="12">
                  <el-form-item label="验证码">
                    <el-input placeholder="请填写验证码"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-button type="primary">获取验证码</el-button>
                </el-col>
              </el-row>
            </el-col>

            <!-- 图形验证码验证 -->
            <el-col :span="24">
              <el-row :gutter="24">
                <el-col :span="12">
                  <el-form-item label="验证码">
                    <el-input placeholder="请填写验证码"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="12"></el-col>
              </el-row>
            </el-col>

            <!-- 登录 -->
            <el-col :span="24">
              <div class="login_button">
                <el-button type="primary">登录</el-button>
              </div>
            </el-col>
          </el-row>

          <!-- 邮箱验证码登录 -->
          <el-row v-if="loginType === '1'">
            <!-- 邮箱号 -->
            <el-col :span="24">
              <el-form-item label="邮箱号">
                <el-input placeholder="请填写邮箱号"></el-input>
              </el-form-item>
            </el-col>

            <!-- 发送验证码 -->
            <el-col :span="24">
              <el-row :gutter="24">
                <el-col :span="12">
                  <el-form-item label="验证码">
                    <el-input placeholder="请填写验证码"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-button type="primary">获取验证码</el-button>
                </el-col>
              </el-row>
            </el-col>

            <!-- 图形验证码验证 -->
            <el-col :span="24">
              <el-row :gutter="24">
                <el-col :span="12">
                  <el-form-item label="验证码">
                    <el-input placeholder="请填写验证码"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="12"></el-col>
              </el-row>
            </el-col>

            <!-- 登录 -->
            <el-col :span="24">
              <div class="login_button">
                <el-button type="primary">登录</el-button>
              </div>
            </el-col>
          </el-row>

          <!-- 账号密码登录 -->
          <el-row v-if="loginType === '2'">
            <el-col :span="24">
              <el-form-item label="账号">
                <el-input placeholder="请填写账号"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="密码">
                <el-input placeholder="请填写密码"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-row :gutter="24">
                <el-col :span="12">
                  <el-form-item label="验证码">
                    <el-input placeholder="请填写验证码"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="12"></el-col>
              </el-row>
            </el-col>
            <el-col :span="24">
              <div class="login_button">
                <el-button type="primary">登录</el-button>
              </div>
            </el-col>
          </el-row>
        </el-form>
      </div>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import type { FormInstance, FormRules } from "element-plus";

const isDialog = ref(false); // 控制对话框显示隐藏
const loginRef = ref<FormInstance>(); // 表单元素
const loginType = ref("0"); // 登录方式
const loginForm = reactive({});

const loginRules = ref<FormRules>({});

/** 暴露属性 */
defineExpose({
  isDialog,
});
</script>

<style lang="less" scoped>
.login {
  width: 100%;
  display: flex;
  .follow_official_account {
    width: 45%;
    display: flex;
    flex-direction: column;
    align-items: center;
    border-right: 1px solid @textColor;
    .img {
      width: 100%;
      height: 100%;
    }
    .text {
      font-size: 1.25vw;
      margin-top: 1.5vh;
      font-weight: 600;
      .illustrate {
        font-size: 1vw;
        font-weight: 400;
      }
    }
  }
  .form_login {
    margin-left: 1.85vw;
    width: 55%;
    .login_form {
      margin-top: 5vh;
      .login_button {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
}
</style>
