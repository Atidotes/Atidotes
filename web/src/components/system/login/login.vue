<template>
  <el-dialog
    destroy-on-close
    @open="handleOpen"
    title="登录"
    v-model="isDialog"
    width="50vw"
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
        <el-row class="login_tab">
          <el-col :span="24">
            <el-radio-group @change="handleChange" v-model="params.loginType">
              <el-radio-button :label="0">手机登录</el-radio-button>
              <el-radio-button :label="1">邮箱登录</el-radio-button>
              <el-radio-button :label="2">账号登录</el-radio-button>
            </el-radio-group>
          </el-col>
        </el-row>

        <!-- 手机验证码登录 -->
        <el-row v-if="params.loginType === 0">
          <mobileLogin ref="mobileLoginRef" />
        </el-row>

        <!-- 邮箱验证码登录 -->
        <el-row v-if="params.loginType === 1">
          <emailLogin ref="emailLoginRef" />
        </el-row>

        <!-- 账号密码登录 -->
        <el-row v-if="params.loginType === 2">
          <accountLogin ref="accountLoginRef" />
        </el-row>

        <el-row>
          <el-col :span="24">
            <div class="login_button">
              <el-button @click="handleLogin" type="primary">登录</el-button>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue";
import mobileLogin from "./components/mobileLogin/mobileLogin.vue";
import emailLogin from "./components/emailLogin/emailLogin.vue";
import accountLogin from "./components/accountLogin/accountLogin.vue";
import { useCurrentInstance } from "@/hooks/system/useSystem";
import { loginApi } from "@/api/system/login/login";

/** 获取组件实例 */
const { proxy } = useCurrentInstance();

const isDialog = ref(false); // 控制对话框显示隐藏
const mobileLoginRef = ref<InstanceType<typeof mobileLogin>>(); // 手机表单DOM
const emailLoginRef = ref<InstanceType<typeof emailLogin>>(); // 邮箱表单DOM
const accountLoginRef = ref<InstanceType<typeof accountLogin>>(); // 账号表单DOM
const params = ref({
  loginType: 0, // 登录类型 0 手机验证码登录 1 邮箱验证码登录 2 账号密码登录
  mobile: null, // 手机号
  mailbox: null, // 邮箱号
  account: null, // 账号
  password: null, // 密码
  captcha: null, // 验证码
  captchaID: null, // 验证码ID
  chartCaptcha: null, // 图形验证码
  chartCaptchaID: null, // 图形验证码ID
});

/** 弹窗打开时回调 */
const handleOpen = () => {
  handleChange(params.value.loginType);
};

/** 点击切换tab */
const handleChange = (index: number) => {
  if (index === 0) {
    mobileLoginRef.value?.getChartCaptchaData();
    emailLoginRef.value?.resetFieldsMailbox();
    accountLoginRef.value?.resetFieldsAccount();
  } else if (index === 1) {
    emailLoginRef.value?.getChartCaptchaData();
    mobileLoginRef.value?.resetFieldsMobile();
    accountLoginRef.value?.resetFieldsAccount();
  } else if (index === 2) {
    accountLoginRef.value?.getChartCaptchaData();
    mobileLoginRef.value?.resetFieldsMobile();
    emailLoginRef.value?.resetFieldsMailbox();
  }
};

/** 点击提交表单并登录 */
const handleLogin = () => {
  if (params.value.loginType === 0) {
    mobileSubmit();
  } else if (params.value.loginType === 1) {
    emailSubmit();
  } else if (params.value.loginType === 2) {
    accountSubmit();
  }
};

/** 提交手机表单 */
const mobileSubmit = () => {
  if (!mobileLoginRef.value) return;
  mobileLoginRef.value.validateMobile((valid) => {
    if (valid) {
      Object.assign(params.value, mobileLoginRef.value?.formState);
      console.log("手机表单获取参数", params.value);
      proxy.$message.success("登录成功");
      mobileLoginRef.value?.resetFieldsMobile();
      Object.assign(params.value, mobileLoginRef.value?.formState);
    } else {
      proxy.$message.error("表单验证失败");
    }
  });
};

/** 提交邮箱表单 */
const emailSubmit = () => {
  if (!emailLoginRef.value) return;
  emailLoginRef.value.validateMailbox(async (valid) => {
    if (valid) {
      Object.assign(params.value, emailLoginRef.value?.formState);
      let res = await loginApi(params.value);
      if (res.code == 200 && res.success) {
        proxy.$message.success("登录成功");
        emailLoginRef.value?.resetFieldsMailbox();
        Object.assign(params.value, emailLoginRef.value?.formState);
        isDialog.value = false
      } else {
        proxy.$message.error("登录失败");
      }
    } else {
      proxy.$message.error("表单验证失败");
    }
  });
};

/** 提交账号表单 */
const accountSubmit = () => {
  if (!accountLoginRef.value) return;
  accountLoginRef.value.validateAccount((valid) => {
    if (valid) {
      Object.assign(params.value, accountLoginRef.value?.formState);
      console.log("账号获取参数", params.value);

      proxy.$message.success("登录成功");
      accountLoginRef.value?.resetFieldsAccount();
      Object.assign(params.value, accountLoginRef.value?.formState);
    } else {
      proxy.$message.error("表单验证失败");
    }
  });
};

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
    .login_tab {
      margin-bottom: 3.5vh;
    }
    .login_button {
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}
</style>
