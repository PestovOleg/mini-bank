#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER ${AUTH_MINIBANK_USER} WITH PASSWORD '${AUTH_MINIBANK_PASSWORD}';
    CREATE DATABASE ${AUTH_MINIBANK_DB} OWNER ${AUTH_MINIBANK_USER};
EOSQL