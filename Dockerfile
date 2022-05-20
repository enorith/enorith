FROM alpine:3.9 

ENV ALPINE_REGISTRY=mirrors.aliyun.com \
    TIMEZONE=UTC

RUN sed -i "s/dl-cdn.alpinelinux.org/${ALPINE_REGISTRY}/g" /etc/apk/repositories && \
    apk update && apk add --no-cache ca-certificates tzdata && \
    ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && echo "${TIMEZONE}" > /etc/timezone

WORKDIR /app

COPY config config
COPY build/enorith enorith
RUN chmod +x enorith

EXPOSE 8000

CMD ["./enorith"]