FROM golang:1.22.2 as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN    cd cmd/game-client \
    && CGO_ENABLED=0 \
       go build -o /go/bin/rps-gRPC-client

FROM gcr.io/distroless/base-debian10

       
COPY --from=builder /go/bin/rps-gRPC-client /
ENTRYPOINT ["/rps-gRPC-client"]      