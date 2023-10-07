module github.com/PestovOleg/mini-bank/backend/pkg/middleware

go 1.20

require (
	github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0
	github.com/Unleash/unleash-client-go/v3 v3.8.0
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/twmb/murmur3 v1.1.5 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//nolint:gomoddirectives
replace github.com/PestovOleg/mini-bank/backend/pkg/logger => ../../../backend/pkg/logger
