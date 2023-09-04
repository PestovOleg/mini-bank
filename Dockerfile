FROM golang:1.20-alpine

WORKDIR /app
COPY build/minibank ./
ENTRYPOINT ["./minibank"]