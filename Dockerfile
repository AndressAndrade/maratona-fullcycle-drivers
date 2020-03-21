FROM golang:1.14

WORKDIR /go/src

COPY . .

RUN GOOS=linux go build

EXPOSE 8081

ENTRYPOINT ("./driver")