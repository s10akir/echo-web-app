FROM golang:latest

COPY . /go/src/github.com/s10akir/echo-web-app

WORKDIR /go/src/github.com/s10akir/echo-web-app/src

CMD ["go", "run", "main.go"]
