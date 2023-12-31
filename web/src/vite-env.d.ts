/// <reference types="vite/client" />

/** vue类型 */
declare module '*.vue' {
    import { ComponentOptions } from 'vue'
    const componentOptions: ComponentOptions
    export default componentOptions
}

/** 环境变量 */
interface ImportMetaEnv {
    readonly APP_BASE: string; //定义提示信息 数据是只读的无法被修改
    readonly APP_CHAT_BASE: string;
    readonly APP_ENCRYPTION: string;
}