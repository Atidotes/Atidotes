<template>
  <el-form
    ref="mobileRef"
    :model="formState"
    status-icon
    :rules="rules"
    label-width="3.75vw"
    hide-required-asterisk
  >
    <el-row>
      <!-- 手机号 -->
      <el-col :span="24">
        <el-form-item label="手机号" prop="mobile">
          <el-input
            v-model.trim="formState.mobile"
            placeholder="请填写手机号"
          ></el-input>
        </el-form-item>
      </el-col>

      <!-- 发送验证码 -->
      <el-col :span="24">
        <el-row :gutter="24">
          <el-col :span="12">
            <el-form-item label="验证码" prop="captcha">
              <el-input
                v-model.trim="formState.captcha"
                placeholder="请填写验证码"
              ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-button
              @click="handleCaptcha"
              :disabled="captchaFlag"
              type="primary"
              >{{ captchaFlag ? countdown : "获取手机验证码" }}</el-button
            >
          </el-col>
        </el-row>
      </el-col>

      <!-- 图形验证码验证 -->
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
          <el-col :span="12"></el-col>
        </el-row>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount, computed } from "vue";
import { useCurrentInstance } from "@/hooks/system/useSystem";
import type {
  FormInstance,
  FormRules,
  FormValidateCallback,
} from "element-plus";
import type { ILoginForm } from "../../type";

const { proxy } = useCurrentInstance();

const captchaFlag = ref(false); // 控制是否点击发送邮箱验证码
const time = ref(60); // 时间
const timer = ref(); // 定时器
const mobileRef = ref<FormInstance>();
const formState = ref<ILoginForm>({
  mobile: null,
  captcha: null,
  chartCaptcha: null,
});

/** 验证规则 */
const rules = ref<FormRules>({
  mobile: [{ required: true, message: "请输入手机号", trigger: "blur" }],
  captcha: [{ required: true, message: "请输入手机验证码", trigger: "blur" }],
  chartCaptcha: [
    { required: true, message: "请输入图形验证码", trigger: "blur" },
  ],
});

/** 点击获取验证码 */
const handleCaptcha = () => {
  if (!mobileRef.value) return;
  mobileRef.value.validateField("mobile", (valid) => {
    if (valid) {
      captchaFlag.value = true;
      proxy.$message.success("发送成功");
    } else {
      proxy.$message.error("请填写手机号");
    }
  });
};

/** 倒计时 */
const countdown = computed(() => {
  if (captchaFlag.value) {
    timer.value = setTimeout(() => {
      time.value = time.value - 1;
      if (time.value === 0 || time.value < 0) {
        captchaFlag.value = false;
        time.value = 60;
        clearTimeout(timer.value);
      }
    }, 1000);
    return `${time.value} 秒`;
  } else {
    clearTimeout(timer.value);
    time.value = 60;
    return `${time.value} 秒`;
  }
});

/** 表单验证 */
const validateMobile = (callback: FormValidateCallback | undefined) => {
  if (!mobileRef.value) return;
  return mobileRef.value.validate(callback);
};

/** 重置表单 */
const resetFieldsMobile = () => {
  if (!mobileRef.value) return;
  mobileRef.value.resetFields();
  captchaFlag.value = false;
  clearTimeout(timer.value);
  time.value = 60;
};

/** 页面离开清空数据 */
onBeforeUnmount(() => {
  resetFieldsMobile();
});

defineExpose({
  validateMobile,
  resetFieldsMobile,
  formState,
});
</script>

<style scoped lang="less"></style>
