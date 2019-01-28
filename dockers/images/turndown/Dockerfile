# Version: 1.0.0
FROM node:11.8.0
MAINTAINER "Liqiang Liu" "liqianglau@outlook.com"

WORKDIR /usr/src/turndown
COPY . /usr/src/turndown

RUN npm install && npm run build
RUN npm install http-server -g

EXPOSE 80

ENTRYPOINT ["http-server", "/usr/src/turndown"]
CMD ["-p", "80"]
