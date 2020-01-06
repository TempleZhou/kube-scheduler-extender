FROM node:6.14.2
EXPOSE 8080
COPY server.js .
# 为了让启动完成后 15 秒再 ready
CMD sleep 15 && node server.js
