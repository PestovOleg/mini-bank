FROM golang:1.20-alpine

WORKDIR /app
COPY ./ ./
RUN mkdir -p build/; go build -o build/minibank cmd/main.go
RUN apk update && apk add curl
ENTRYPOINT ["./build/minibank"]