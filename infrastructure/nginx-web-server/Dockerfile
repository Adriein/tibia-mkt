FROM nginx:alpine

WORKDIR /etc/nginx/conf.d

RUN rm -rf default.conf

COPY config/api.conf .

EXPOSE 80