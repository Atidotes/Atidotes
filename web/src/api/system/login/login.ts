import { Http } from "@/utils/system/axios/http";
import type { ILoginForm } from "@/components/system/login/type";

enum Api {
  chartCaptcha = "/v1/system/chartCaptcha", // 图形验证码
  emailCaptcha = "/v1/system/emailCaptcha", // 邮箱验证码
}

interface ICaptcha {
  id: string;
  captcha?: string;
}

/** 获取图形验证码 */
export const chartCaptchaApi = () => {
  return Http.Get<ICaptcha, any>({ url: Api.chartCaptcha });
};

/** 获取邮箱验证码 */
export const emailCaptchaApi = (params: ILoginForm) => {
  return Http.Get<ICaptcha, ILoginForm>({ url: Api.emailCaptcha, params });
};
