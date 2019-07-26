# Hello Function

This is the Hello function

Generated with

```
micro new example-fnc --namespace=nebula.core --alias=hello --type=fnc --plugin=registry=etcd:broker=kafka
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: nebula.core.fnc.hello
- Type: fnc
- Alias: hello

## Dependencies

Micro functions depend on service discovery. The default is consul.

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

Run the function once
```
./hello-fnc
```

Build a docker image
```
make docker
```