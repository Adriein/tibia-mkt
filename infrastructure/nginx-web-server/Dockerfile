FROM nginx:alpine

WORKDIR /etc/nginx/conf.d

RUN rm -rf default.conf

COPY ../../infrastructure/nginx-web-server/config .

EXPOSE 80