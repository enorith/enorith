FROM alpine:3.9 

ENV ALPINE_REGISTRY=mirrors.aliyun.com \
    APP_TIMEZONE=UTC

WORKDIR /app
COPY config config
COPY build/enorith enorith
COPY entrypoint.sh entrypoint.sh

RUN sed -i "s/dl-cdn.alpinelinux.org/${ALPINE_REGISTRY}/g" /etc/apk/repositories && \
    apk update && apk add --no-cache ca-certificates tzdata && \
    chmod +x enorith && chmod +x entrypoint.sh

EXPOSE 8000

CMD ["./entrypoint.sh"]