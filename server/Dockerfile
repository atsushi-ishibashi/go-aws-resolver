FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/atsushi-ishibashi/go-aws-resolver
RUN \
  apk update && \
  rm -rf /var/cache/apk/* && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR/server
RUN go build -ldflags "-s -w" -o go-aws-resolver .

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/atsushi-ishibashi/go-aws-resolver/server/go-aws-resolver ./
CMD ["./go-aws-resolver"]
