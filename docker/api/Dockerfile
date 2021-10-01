FROM golang:1.16.3 as api-builder

WORKDIR /go/src/ip-updater/

COPY adapter ./adapter
COPY entity ./entity
COPY infra ./infra
COPY go.mod go.sum main.go ./

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o api.cgi

FROM teamfruit/nginx-fcgiwrap:1.17.9

COPY docker/api/api_conf /cgi/config
COPY docker/api/nginx_conf /etc/nginx/conf.d
COPY --from=api-builder /go/src/ip-updater/api.cgi /cgi/
RUN chmod 755 -R /cgi