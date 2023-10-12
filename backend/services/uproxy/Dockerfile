# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /workspace
COPY ./backend/services/uproxy/go.mod ./backend/services/uproxy/go.sum ./backend/services/uproxy/
COPY ./backend/services/uproxy ./backend/services/uproxy
COPY ./backend/pkg ./backend/pkg

RUN cd ./backend/services/uproxy && go mod download
RUN cd ./backend/services/uproxy && go build -o /minibank ./cmd/main.go

# Stage 2: Run 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /minibank .

CMD ["./minibank"]
