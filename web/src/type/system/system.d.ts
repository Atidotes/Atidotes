import type { AxiosResponse } from "axios";

/** 接口返回值 */
export interface IResult<T = any> {
  code: number;
  message: string;
  result: T;
  success: boolean;
  timeStamp: number;
}

export interface IParams<T> {
  url: string;
  params?: T;
}

/** Http请求 */
export interface IHttp {
  Get: <T, D>(params: IParams<D>) => Promise<IResult<T>>;
  Post: <T, D>(params: IParams<D>) => Promise<AxiosResponse<IResult<T>>>;
  Delete: <T, D>(params: IParams<D>) => Promise<AxiosResponse<IResult<T>>>;
  Put: <T, D>(params: IParams<D>) => Promise<AxiosResponse<IResult<T>>>;
}
