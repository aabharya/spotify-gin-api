# Minimal Spotify Gin API


## Prerequisites

- `Go 1.22`
- `Postgresql 15`


## Development

### `.env` example

```shell
DB_HOST="localhost"
DB_USER="go_practice"
DB_PASS="123456"
DB_NAME="go_practice"
DB_PORT="5314"
REDIS_HOST="localhost"
REDIS_PORT="6378"
```

### Run the gin app

```shell
make dev
```

### Run tests

```shell
make test
```
