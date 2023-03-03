# DouYin

## Introduction

A simple note service built with `Kitex` and `Hertz` which is divided into three microservices.

| Service Name | Usage                    | Framework   | protocol | Path         | IDL                 |
|--------------|--------------------------|-------------|----------|--------------|---------------------|
| api          | HTTP interface           | kitex/hertz | http     | cmd/api      | idl/api.thrift      |
| comment      | comment data management  | kitex/gorm  | thrift   | cmd/comment  | idl/comment.thrift  |
| favorite     | favorite data management | kitex/gorm  | thrift   | cmd/favorite | idl/favorite.thrift |
| feed         | feed data management     | kitex/gorm  | thrift   | cmd/feed     | idl/feed.thrift     |
| message      | message data management  | kitex/gorm  | thrift   | cmd/message  | idl/message.thrift  |
| publish      | publish data management  | kitex/gorm  | thrift   | cmd/publish  | idl/publish.thrift  |
| relation     | relation data management | kitex/gorm  | thrift   | cmd/relation | idl/relation.thrift |
| user         | user data management     | kitex/gorm  | thrift   | cmd/user     | idl/user.thrift     |

### Call Relations

### Basic Features

- Hertz
    - Use `thrift` IDL to define HTTP interface
    - Use `hz` to generate code
    - Use `Hertz` binding and validate
    - Use `obs-opentelemetry` and `jarger` for `tracing`, `metrics`, `logging`
    - Middleware
        - Use `requestid`, `jwt`, `recovery`, `pprof`, `gzip`
- Kitex
    - Use `thrift` IDL to define `RPC` interface
    - Use `kitex` to generate code
    - Use `thrift-gen-validator` for validating RPC request
    - Use `obs-opentelemetry` and `jarger` for `tracing`, `metrics`, `logging`
    - Use `registry-etcd` for service discovery and register

### Catalog Introduce

| catalog     | introduce               |
|-------------|-------------------------|
| handler     | HTTP handler            |
| service     | business logic          |
| rpc         | RPC call logic          |
| dal         | DB operation            |
| pack        | data pack               |
| pkg/mw      | RPC middleware          |
| pkg/consts  | constants               |
| pkg/errno   | customized error number |
| pkg/configs | SQL and Tracing configs |

## Quick Start

### Setup Basic Dependence

```shell
docker-compose up
```

### Run Comment RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Favorite RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Feed RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Message RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Publish RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Relation RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run User RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run API Server

```shell
cd cmd/api
go run .
```

### Jaeger

Visit `http://127.0.0.1:16686/` on browser

#### Snapshots


### Grafana

Visit `http://127.0.0.1:3000/` on browser

#### Dashboard Example


## API Requests

The following is a list of API requests and partial responses.
