FROM golang:1

RUN curl https://glide.sh/get | sh
ADD . /go/src/github.com/davidbanham/notify

WORKDIR /go/src/github.com/davidbanham/notify

RUN glide install
RUN go install github.com/davidbanham/notify
RUN touch .env

CMD /go/bin/notify
