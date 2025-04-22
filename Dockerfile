FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o main ./cmd/web/main.go
FROM alpine:latest
RUN apk update && apk add --no-cache tzdata && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && echo "Europe/Moscow" > /etc/timezone
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
CMD ["./main"]
EXPOSE 8080