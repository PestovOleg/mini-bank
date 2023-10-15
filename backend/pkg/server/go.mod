module github.com/PestovOleg/mini-bank/backend/pkg/server

go 1.20

require (
	github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0
	go.uber.org/zap v1.26.0
)

require (
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)

//nolint:gomoddirectives
replace github.com/PestovOleg/mini-bank/backend/pkg/logger => ../../../backend/pkg/logger
