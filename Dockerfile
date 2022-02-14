FROM golang:1.17-alpine

WORKDIR $GOPATH/github.com/airbenders/auth

COPY . .

RUN go get -d -v ./...

EXPOSE 8080

CMD ["go", "run", "main.go"]
