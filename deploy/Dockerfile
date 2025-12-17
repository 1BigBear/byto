# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git bash openssh

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o byto-bot ./cmd/bot/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

# Install runtime dependencies
RUN apk add --no-cache ffmpeg python3 curl
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp \
    && chmod a+rx /usr/local/bin/yt-dlp

COPY --from=builder /app/byto-bot .

ENTRYPOINT ["./byto-bot"]