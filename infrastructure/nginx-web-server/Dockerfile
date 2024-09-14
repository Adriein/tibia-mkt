FROM nginx:alpine

WORKDIR /etc/nginx/conf.d

RUN rm -rf default.conf

COPY ./*.conf .

EXPOSE 80