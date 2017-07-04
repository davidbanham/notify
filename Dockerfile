FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD ./bin/notify /app

ENTRYPOINT ["/app"]
