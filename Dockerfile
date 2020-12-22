FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/cloud-inventory

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cloud-inventory .


FROM alpine:latest

COPY --from=builder /go/src/cloud-inventory/cloud-inventory ./


ENTRYPOINT ["./cloud-inventory"]