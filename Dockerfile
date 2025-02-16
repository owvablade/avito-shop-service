FROM golang:alpine

COPY . /go/src/app

WORKDIR /go/src/app

RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]