# Hello Service

This is the Hello service

Generated with

```
micro new example-srv --namespace=nebula.core --alias=hello --type=srv --plugin=registry=etcd:broker=kafka
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: nebula.core.srv.hello
- Type: srv
- Alias: hello

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./hello-srv
```

Build a docker image
```
make docker
```