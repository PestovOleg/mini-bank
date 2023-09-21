FROM golang:1.20-alpine

WORKDIR /app
COPY ./backend ./
RUN go mod download
RUN mkdir -p build/; go build -o build/minibank cmd/main.go
RUN apk update && apk add curl
ENTRYPOINT ["./build/minibank"]