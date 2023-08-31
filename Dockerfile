FROM golang:1.20-alpine

WORKDIR /app
COPY build/minibank ./
COPY ./config/local.yaml ./
ENV CONFIG_PATH=./local.yaml
ENTRYPOINT ["./minibank"]