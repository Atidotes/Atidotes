import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), "APP_");

  return {
    envPrefix: "APP_", // 设置环境变量

    // 导入环境变量
    define: {
      __APP_ENV__: env.APP_ENV,
    },

    // 打包配置
    build: {
      minify: "terser",
      terserOptions: {
        compress: {
          //生产环境时移除console
          drop_console: true,
          drop_debugger: true,
        },
      },
    },

    // 插件
    plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],

    // 服务代理
    server: {
      proxy: {
        "/api": {
          target: env.APP_BASE,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, "/api"),
        },
      },
    },

    // 配置别名
    resolve: {
      alias: {
        "@": resolve(__dirname, "src"),
      },
    },

    //  全局引入less样式
    css: {
      preprocessorOptions: {
        less: {
          javascriptEnabled: true,
          additionalData: `@import "${resolve(
            __dirname,
            "src/style/style.less"
          )}";`,
        },
      },
    },
  };
});
