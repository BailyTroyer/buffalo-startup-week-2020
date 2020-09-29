FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/main/bflobox-api
COPY go.* ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/
RUN GOOS=linux CGO_ENABLED=0 go build -o /go/bin/bflobox-api ./cmd/main.go

FROM scratch
COPY --from=builder /go/bin/bflobox-api /go/bin/bflobox-api
ENTRYPOINT ["/go/bin/bflobox-api"]
