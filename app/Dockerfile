FROM golang:1.13.5-alpine3.10 as builder

RUN apk add --update --no-cache ca-certificates git gcc

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY handler handler

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
