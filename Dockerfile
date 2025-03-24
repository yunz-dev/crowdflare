FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/server/main.go

COPY scripts/start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

RUN apk add --no-cache curl sudo && \
    curl -L --output cloudflared-linux-amd64 https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64 && \
    chmod +x cloudflared-linux-amd64 && \
    mv cloudflared-linux-amd64 /usr/local/bin/cloudflared && \
    apk del curl sudo

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /usr/local/bin/cloudflared /usr/local/bin/cloudflared

EXPOSE 8080

CMD ["/usr/local/bin/start.sh"]
