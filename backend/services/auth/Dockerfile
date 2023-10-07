# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /workspace
COPY ./backend/services/auth/go.mod ./backend/services/auth/go.sum ./backend/services/auth/
COPY ./backend/services/auth ./backend/services/auth
COPY ./backend/pkg ./backend/pkg

RUN cd ./backend/services/auth && go mod download
RUN cd ./backend/services/auth && go build -o /minibank ./cmd/main.go

# Stage 2: Run 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /minibank .

CMD ["./minibank"]
