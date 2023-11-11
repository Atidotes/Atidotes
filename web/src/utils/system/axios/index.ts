import axios from "axios";

/** 基本配置 */
const request = axios.create({
  baseURL: "/api",
  timeout: 600 * 1000,
});

// 添加请求拦截器
request.interceptors.request.use(
  (config) => {
    // 携带token信息
    const token = localStorage.getItem('token')
    if(token){
      config.headers.Authorization = `Bearer ${token}`
    }
    return config;
  },
  (error) => {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
request.interceptors.response.use(
  (response) => {
    // 获取token信息，存储本地
    const { authorization } = response.headers;
    authorization && localStorage.setItem("token", authorization);
    return response.data;
  },
  (error) => {
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

export default request;
