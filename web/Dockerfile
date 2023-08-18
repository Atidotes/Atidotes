# 项目所处环境
FROM nginx

# 代表生产环境
ENV PROJECT_ENV production

# 需要把生产代码部署到环境中
COPY dist/ /usr/share/nginx/html/

# 配置nginx
COPY default.conf /etc/nginx/conf.d/default.conf