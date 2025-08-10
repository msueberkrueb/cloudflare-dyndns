FROM golang:alpine AS build

RUN apk --update add ca-certificates && update-ca-certificates

WORKDIR /tmp/cloudflare-dyndns

COPY ./cmd/cloudflare-dyndns ./
COPY ./go.mod ./go.sum ./

RUN go get -v
RUN go build -o /usr/local/bin/cloudflare-dyndns
RUN chown 1000:1000 /usr/local/bin/cloudflare-dyndns

FROM scratch

COPY --from=build /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=build /usr/local/bin/cloudflare-dyndns /usr/local/bin/cloudflare-dyndns

USER 1000:1000

LABEL org.opencontainers.image.source=https://github.com/msueberkrueb/cloudflare-dyndns \
  org.opencontainers.image.description="CloudFlare DynDNS is used for dynamic and configurable A record updates for networks without a static ip address." \
  org.opencontainers.image.licenses=UNLICENSE

ENTRYPOINT ["/usr/local/bin/cloudflare-dyndns"]
