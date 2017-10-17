# IR-Remocon

## Device
- PC-OP-RS1

## Build

Can use Makefile to compile for cross platform.

```sh
$ make build_for_windows    # windows
$ make build_for_mac        # mac
$ make build_for_linux      # linux
```

## Develop

### Setup

Using [Dep](https://github.com/golang/dep) for dependnecy management.

#### How to install Dep

- homebrew

```sh
$ brew install dep
$ brew upgrade dep
```

- go get

```sh
$ go get -u github.com/golang/dep/cmd/dep
```

### Installing dependencies

```sh
$ dep ensure
```

