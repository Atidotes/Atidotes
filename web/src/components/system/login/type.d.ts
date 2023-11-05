/** 登录表单 */
export interface ILoginForm {
  loginType?: number; // 登录类型 0 手机验证码登录 1 邮箱验证码登录 2 账号密码登录
  mobile?: null | number; // 手机号
  mailbox?: null | string; // 邮箱号
  account?: null | string; // 账号
  password?: null | string; // 密码
  captcha?: null | string; // 验证码
  captchaID?: null | string; // 验证码ID
  chartCaptcha?: null | string; // 图形验证码
  chartCaptchaID?:null | string; // 图形验证码ID
}
