FROM golang:1.22

COPY . /go/src/

WORKDIR /go/src/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

EXPOSE 8080

CMD ["./main"]
