# devcamp-caching

Caching hands-on session for Devcamp 2023

## Getting Started

```text
Make sure go version >=1.17 is already installed on your machine.
All command below should be run on project root folder.
```

### Pull dependencies:

```shell
go mod tidy && go mod vendor -v
```

### Run dependencies using:

```shell
make docker-start
```

### Run application using:

Run #1 hands-on binary:

```shell
go run cmd/1_cache_latency/main.go
```

Run #2 hands-on binary:
```shell
go run cmd/2_cache_local/main.go
```

Run #3 hands-on binary:
```shell
go run cmd/3_cache_pattern/main.go
```

Run #4 hands-on binary:
```shell
go run cmd/4_cache_challenge/main.go
```

## PIC

* Evin Cintiawan (ecintiawan)
