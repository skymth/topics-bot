FROM golang:latest

WORKDIR /go/src/topics-bot
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["topics-bot"]
