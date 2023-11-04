import {
  ComponentInternalInstance,
  ComponentCustomProperties,
  getCurrentInstance,
} from "vue";
import type { Message, IElMessageBox } from "element-plus";

/**
 * proxy 快速使用组件
 */
export const useCurrentInstance = () => {
  interface ICustomElementsUI extends ComponentCustomProperties {
    $message: Message;
    $messageBox: IElMessageBox;
  }

  const { appContext } = getCurrentInstance() as ComponentInternalInstance;
  const proxy = appContext.config.globalProperties as ICustomElementsUI &
    Record<string, any>;
  return {
    proxy,
  };
};
