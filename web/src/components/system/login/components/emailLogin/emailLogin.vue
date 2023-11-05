<template>
  <el-form
    ref="mailboxRef"
    :model="formState"
    status-icon
    :rules="rules"
    label-width="3.75vw"
    hide-required-asterisk
  >
    <el-row>
      <!-- 邮箱号 -->
      <el-col :span="24">
        <el-form-item label="邮箱号" prop="mailbox">
          <el-input
            v-model.trim="formState.mailbox"
            placeholder="请填写邮箱号"
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
              :loading="loading"
              @click="handleCaptcha"
              :disabled="captchaFlag"
              type="primary"
              >{{ captchaFlag ? countdown : "获取邮箱验证码" }}</el-button
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
import { ref, computed, onBeforeUnmount } from "vue";
import type { ILoginForm } from "../../type";
import { useCurrentInstance } from "@/hooks/system/useSystem";
import { chartCaptchaApi, emailCaptchaApi } from "@/api/system/login/login";
import type {
  FormInstance,
  FormRules,
  FormValidateCallback,
} from "element-plus";

const { proxy } = useCurrentInstance();

const captchaFlag = ref(false); // 控制是否点击发送邮箱验证码
const time = ref(60); // 时间
const timer = ref(); // 定时器
const captcha = ref(); // 验证码数据
const chartLoading = ref(false); // 验证码加载loading
const loading = ref(false); // 邮箱验证码加载loading
const mailboxRef = ref<FormInstance>();
const formState = ref<ILoginForm>({
  mailbox: null,
  captcha: null,
  captchaID: null,
  chartCaptcha: null,
  chartCaptchaID: null,
});

/** 验证规则 */
const rules = ref<FormRules>({
  mailbox: [{ required: true, message: "请输入邮箱号", trigger: "blur" }],
  captcha: [{ required: true, message: "请输入邮箱验证码", trigger: "blur" }],
  chartCaptcha: [
    { required: true, message: "请输入图形验证码", trigger: "blur" },
  ],
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

/** 点击获取验证码 */
const handleCaptcha = () => {
  if (!mailboxRef.value) return;
  mailboxRef.value.validateField("mailbox", async (valid) => {
    if (valid) {
      loading.value = true;
      let res = await emailCaptchaApi({ mailbox: formState.value.mailbox });
      if (res.code === 200 && res.success) {
        loading.value = false;
        captchaFlag.value = true;
        formState.value.captchaID = res.result.id
        proxy.$message.success("发送成功");
      } else {
        loading.value = false;
        proxy.$message.success("发送失败");
      }
    } else {
      proxy.$message.error("请填写邮箱号");
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
const validateMailbox = (callback: FormValidateCallback | undefined) => {
  if (!mailboxRef.value) return;
  mailboxRef.value.validate(callback);
};

/** 重置表单 */
const resetFieldsMailbox = () => {
  if (!mailboxRef.value) return;
  mailboxRef.value.resetFields();
  captchaFlag.value = false;
  clearTimeout(timer.value);
  time.value = 60;
};

/** 页面离开清空数据 */
onBeforeUnmount(() => {
  resetFieldsMailbox();
});

defineExpose({
  validateMailbox,
  resetFieldsMailbox,
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
