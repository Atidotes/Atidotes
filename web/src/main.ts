import { createApp } from "vue";
import App from "@/App.vue";
import store from "@/store/index";
import { ElMessage, ElMessageBox } from "element-plus";
import "element-plus/theme-chalk/dark/css-vars.css";
import "element-plus/es/components/message/style/css";
import "element-plus/es/components/message-box/style/css";
import "@/style/style.less";

const app = createApp(App);
app.use(store);
app.mount("#app");

// 挂载全局UI库
app.config.globalProperties.$message = ElMessage;
app.config.globalProperties.$messageBox = ElMessageBox;
