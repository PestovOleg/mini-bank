module github.com/PestovOleg/mini-bank/backend/services/user

go 1.20

require (
	github.com/google/uuid v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.10.0
	go.uber.org/zap v1.26.0
	golang.org/x/sys v0.12.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Unleash/unleash-client-go/v3 v3.8.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/twmb/murmur3 v1.1.5 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

require (
	github.com/PestovOleg/mini-bank/backend/pkg/config v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/database v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/middleware v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/server v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/signal v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/unleash v0.0.0-00010101000000-000000000000
)

replace github.com/PestovOleg/mini-bank/backend/pkg/logger => ../../../backend/pkg/logger

replace github.com/PestovOleg/mini-bank/backend/pkg/config => ../../../backend/pkg/config

replace github.com/PestovOleg/mini-bank/backend/pkg/database => ../../../backend/pkg/database

replace github.com/PestovOleg/mini-bank/backend/pkg/middleware => ../../../backend/pkg/middleware

replace github.com/PestovOleg/mini-bank/backend/pkg/server => ../../../backend/pkg/server

replace github.com/PestovOleg/mini-bank/backend/pkg/signal => ../../../backend/pkg/signal

replace github.com/PestovOleg/mini-bank/backend/pkg/unleash => ../../../backend/pkg/unleash
