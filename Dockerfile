FROM golang:1.10-alpine

ENV DISTRIBUTION_DIR /go/src/github.com/pmint93/default-backend

WORKDIR $DISTRIBUTION_DIR

COPY . $DISTRIBUTION_DIR

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/default-backend

FROM scratch

COPY --from=0 /go/bin/default-backend /usr/bin/default-backend

ENTRYPOINT ["/usr/bin/default-backend"]
