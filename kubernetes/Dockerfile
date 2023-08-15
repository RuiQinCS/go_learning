# 基础镜像
FROM alpine
# 把编译后的打包进来这个镜像，放到工作目录 /app (可自定义)
COPY webook /app/webook
WORKDIR /app
# CMD 是执行命令
# 最佳
ENTRYPOINT ["/app/webook"]