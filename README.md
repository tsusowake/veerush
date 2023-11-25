# veerush

Work!

## Setup

```shell
$  docker-compose up -d
```

## Tools

```shell
brew install sqldef/sqldef/psqldef
```

## Create migration file

```shell
$ touch db/sql/$(date "+%s")_create_hoge.sql
```

## ...

```shell
export POSTGRES_DATABASE=veerush
export POSTGRES_USER=veerush
export POSTGRES_PASSWORD=password
export POSTGRES_PORT=5432
export POSTGRES_HOST=localhost

```

## sqlc

### dump schema

```shell
make dump-schema
```

### generate code

```shell
make gen-sqlc

```