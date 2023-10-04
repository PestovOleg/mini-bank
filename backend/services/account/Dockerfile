# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /workspace
COPY ./backend/services/account/go.mod ./backend/services/account/go.sum ./backend/services/account/
COPY ./backend/services/account ./backend/services/account
COPY ./backend/pkg ./backend/pkg

RUN cd ./backend/services/account && go mod download
RUN cd ./backend/services/account && go build -o /minibank ./cmd/main.go

# Stage 2: Run 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /minibank .

CMD ["./minibank"]
