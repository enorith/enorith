FROM alpine:3.9 
RUN apk add ca-certificates

WORKDIR /app

COPY config config
COPY build/enorith enorith
RUN chmod +x enorith

EXPOSE 8000

CMD ["./enorith"]