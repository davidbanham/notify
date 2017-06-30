.PHONY: test build

name = notify

include .env

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
		go vet $(go list ./... | grep -v /vendor/)

build: vet
		CGO_ENABLED=0 go build -o ./bin/$(name) -a -installsuffix cgo -ldflags '-s' .

test:
		PORT=23232 \
		NOTIFY_EMAIL_PROVIDER=none \
		NOTIFY_SMS_PROVIDER=none \
		go test -cover `go list ./... | grep -v /vendor/`
