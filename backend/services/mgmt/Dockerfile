# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /workspace
COPY ./backend/services/mgmt/go.mod ./backend/services/mgmt/go.sum ./backend/services/mgmt/
COPY ./backend/services/mgmt ./backend/services/mgmt
COPY ./backend/pkg ./backend/pkg

RUN cd ./backend/services/mgmt && go mod download
RUN cd ./backend/services/mgmt && go build -o /minibank ./cmd/main.go

# Stage 2: Run 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /minibank .

CMD ["./minibank"]
