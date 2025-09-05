# 編譯用 Golang
FROM golang:1.18-alpine AS builder

WORKDIR /src
COPY . .
RUN go build -o main

# 正式映像
FROM alpine:latest
RUN apk add --no-cache ffmpeg tzdata

WORKDIR /app
COPY --from=builder /src/main .
COPY --from=builder /src/config.json .
COPY --from=builder /src/web ./web

EXPOSE 8008
CMD ["./main"]

