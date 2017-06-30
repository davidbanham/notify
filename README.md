# Notify

[![GoDoc](https://godoc.org/github.com/davidbanham/notify?status.svg)](https://godoc.org/github.com/davidbanham/notify)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidbanham/notify)](https://goreportcard.com/report/github.com/davidbanham/notify)

[API docs](https://davidbanham.github.io/notify)

Notify is a notifications microservice. It currently supports the following mechanisms and providers:

1. Email
  1. Gmail
  1. Mandrill
1. SMS
  1. Amazon SNS

Adding more types and providers is super simple!

### API

Notify exposes a RESTful API. Details will be forthcoming in a JSON Schema document.

### Quick Start

Required environment variables are documented in config/config.go. If you miss one, don't worry, the program will refuse to start and ask for the missing variable.

Notify is published on Dockerhub. That's probably the simplest way to get going with it.

```
docker pull davidbanham/notify
```

If you'd like to build it yourself, that's easy too:

```
make install
make build
```

The binary will pop out in `./bin/notify`

### Contributing

We assume that you have [glide](https://github.com/Masterminds/glide) installed.

To get developing just:

```
make install
make test
```

A handy dev server is available via:

```
docker-compose up dev
```

Test coverage is minimal, since testing the individual providers would require either a lot of credentials or a lot of mocking. Generally speaking, more tests are better than less tests, but we don't want to make them such a hassle to write and run that they're a net negative.
