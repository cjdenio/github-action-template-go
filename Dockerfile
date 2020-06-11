FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -v ./...
RUN go build -o main

ENTRYPOINT ["/go/src/app/main"]