import { Http } from "@/utils/system/axios/http";
import type { ILoginForm } from "@/components/system/login/type";

enum Api {
  chartCaptcha = "/v1/system/chartCaptcha", // 图形验证码
  emailCaptcha = "/v1/system/emailCaptcha", // 邮箱验证码
  login = "/v1/system/login", // 登录
  userInfo = "/v1/system/userInfo", // 获取用户信息
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

/** 登录 */
export const loginApi = (params: ILoginForm) => {
  return Http.Post<any, ILoginForm>({ url: Api.login, params });
};

// TODO 测试
export const userInfoApi = () => {
  return Http.Get<any, any>({ url: Api.userInfo });
};
