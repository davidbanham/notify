FROM scratch

ADD ./bin/notify /app

ENTRYPOINT ["/app"]
