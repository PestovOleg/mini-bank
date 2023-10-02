module github.com/PestovOleg/mini-bank/backend/services/auth

go 1.20

require (
	github.com/google/uuid v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.10.0
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.13.0
	golang.org/x/sys v0.12.0
)

require go.uber.org/multierr v1.10.0 // indirect

replace github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0 => ../..