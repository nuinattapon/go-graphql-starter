# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine AS builder
RUN apk --update add ca-certificates && apk --update add git

ENV SRC_DIR=/go/src/github.com/nuinattapon/go-graphql-starter/
# Add the source code:
COPY . $SRC_DIR
WORKDIR $SRC_DIR

# Build it:
RUN go get -v github.com/jteeuwen/go-bindata/...
RUN export PATH=$PATH:$GOPATH/bin
RUN go generate ./schema
RUN go get -v ./
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/nuinattapon/go-graphql-starter/go-graphql-starter /app
COPY --from=builder /go/src/github.com/nuinattapon/go-graphql-starter/*.html /app
COPY --from=builder /go/src/github.com/nuinattapon/go-graphql-starter/Config.toml /app
EXPOSE 3000
ENTRYPOINT ["/app/go-graphql-starter"]