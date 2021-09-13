FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /build
ENV GOPROXY https://goproxy.io
COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o ./enorith cmd/app/main.go

FROM alpine:3.9 
RUN apk add ca-certificates

WORKDIR /app

COPY config config
COPY --from=build_base /build/enorith enorith
RUN chmod +x enorith

EXPOSE 8000

CMD ["./enorith"]