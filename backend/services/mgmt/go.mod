module github.com/PestovOleg/mini-bank/backend/services/mgmt

go 1.20

require (
	github.com/PestovOleg/mini-bank/backend/pkg/config v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/database v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/logger v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/middleware v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/server v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/signal v0.0.0
	github.com/PestovOleg/mini-bank/backend/pkg/unleash v0.0.0
)

replace github.com/PestovOleg/mini-bank/backend/pkg/logger => ../../../backend/pkg/logger

replace github.com/PestovOleg/mini-bank/backend/pkg/config => ../../../backend/pkg/config

replace github.com/PestovOleg/mini-bank/backend/pkg/database => ../../../backend/pkg/database

replace github.com/PestovOleg/mini-bank/backend/pkg/middleware => ../../../backend/pkg/middleware

replace github.com/PestovOleg/mini-bank/backend/pkg/server => ../../../backend/pkg/server

replace github.com/PestovOleg/mini-bank/backend/pkg/signal => ../../../backend/pkg/signal

replace github.com/PestovOleg/mini-bank/backend/pkg/unleash => ../../../backend/pkg/unleash