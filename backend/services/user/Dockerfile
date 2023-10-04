# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /workspace
COPY ./backend/services/user/go.mod ./backend/services/user/go.sum ./backend/services/user/
COPY ./backend/services/user ./backend/services/user
COPY ./backend/pkg ./backend/pkg

RUN cd ./backend/services/user && go mod download
RUN cd ./backend/services/user && go build -o /minibank ./cmd/main.go

# Stage 2: Run 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /minibank .

CMD ["./minibank"]
