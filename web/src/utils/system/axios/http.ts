import defaultRequest from "./index";
import type { IHttp } from "@/type/system/system";

/** 请求方法 */
export const Http: IHttp = {
  /** GET请求 */
  Get: (params) => {
    return defaultRequest({
      url: params.url,
      method: "GET",
      params: params.params,
    });
  },

  /** POST请求 */
  Post: (params) => {
    return defaultRequest({
      url: params.url,
      method: "POST",
      data: params.params,
    });
  },

  /** PUT请求 */
  Put: (params) => {
    return defaultRequest({
      url: params.url,
      method: "Put",
      data: params.params,
    });
  },

  /** DELETE请求 */
  Delete: (params) => {
    return defaultRequest({
      url: params.url,
      method: "DELETE",
      params: params.params,
    });
  },
};
