module github.com/PestovOleg/mini-bank/backend/pkg/database

go 1.20

require (
	github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
)

replace github.com/PestovOleg/mini-bank/backend/pkg/logger => ../../../backend/pkg/logger
