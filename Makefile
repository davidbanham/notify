.PHONY: test build

name = notify
docker_user_id = davidbanham

include .env

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
		go vet `go list ./... | grep -v /vendor/`

build: vet
		CGO_ENABLED=0 go build -o ./bin/$(name) -a -installsuffix cgo -ldflags '-s' .

test:
		PORT=23232 \
		NOTIFY_EMAIL_PROVIDER=none \
		NOTIFY_SMS_PROVIDER=none \
		go test -cover `go list ./... | grep -v /vendor/`

install:
		glide install

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
		golint `go list ./... | grep -v /vendor/`

publish: docker_image_build
		docker tag $(name)/$(name) $(docker_user_id)/$(name)
		docker push $(docker_user_id)/$(name)

docker_image_build: test build
		docker build -t $(name)/$(name) .
