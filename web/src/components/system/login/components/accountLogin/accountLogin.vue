<template>
  <el-form
    ref="accountRef"
    :model="formState"
    status-icon
    :rules="rules"
    label-width="3.75vw"
    hide-required-asterisk
  >
    <el-row>
      <el-col :span="24">
        <el-form-item label="账号" prop="account">
          <el-input
            v-model.trim="formState.account"
            placeholder="请填写账号"
          ></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-form-item label="密码" prop="password">
          <el-input
            v-model.trim="formState.password"
            placeholder="请填写密码"
          ></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-row :gutter="24">
          <el-col :span="12">
            <el-form-item label="验证码" prop="chartCaptcha">
              <el-input
                v-model.trim="formState.chartCaptcha"
                placeholder="请填写验证码"
              ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <div
              v-loading="chartLoading"
              @click="getChartCaptchaData"
              class="chartCaptcha"
            >
              <img :src="captcha" alt="图形验证码" />
            </div>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount } from "vue";
import type { ILoginForm } from "../../type";
import { chartCaptchaApi } from "@/api/system/login/login";
import type {
  FormInstance,
  FormRules,
  FormValidateCallback,
} from "element-plus";

const captcha = ref(); // 验证码数据
const chartLoading = ref(false); // 验证码加载loading
const accountRef = ref<FormInstance>();
const formState = ref<ILoginForm>({
  account: null,
  password: null,
  chartCaptcha: null,
  chartCaptchaID: null,
});

/** 获取图形验证码 */
const getChartCaptchaData = async () => {
  chartLoading.value = true;
  let res = await chartCaptchaApi();
  if (res.code === 200 && res.success) {
    captcha.value = res.result.captcha;
    formState.value.chartCaptchaID = res.result.id;
    chartLoading.value = false;
  } else {
    chartLoading.value = false;
  }
};

/** 验证规则 */
const rules = ref<FormRules>({
  account: [{ required: true, message: "请输入账号", trigger: "change" }],
  password: [{ required: true, message: "请输入密码", trigger: "change" }],
  chartCaptcha: [
    { required: true, message: "请输入图形验证码", trigger: "change" },
  ],
});

/** 表单验证 */
const validateAccount = (callback: FormValidateCallback | undefined) => {
  if (!accountRef.value) return;
  accountRef.value.validate(callback);
};

/** 重置表单 */
const resetFieldsAccount = () => {
  if (!accountRef.value) return;
  accountRef.value.resetFields();
};

/** 页面离开清空数据 */
onBeforeUnmount(() => {
  resetFieldsAccount();
});

defineExpose({
  validateAccount,
  resetFieldsAccount,
  getChartCaptchaData,
  formState,
});
</script>

<style scoped lang="less">
.chartCaptcha {
  width: 100px;
  height: 38px;
  overflow: hidden;
  border-radius: 2px;
  cursor: pointer;
}
</style>
