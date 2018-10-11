FROM golang:latest

WORKDIR /go/src/topics-bot
COPY . .

ENV BOT_ID ""
ENV VERIFICATION_TOKEN ""
ENV CHANNEL_ID ""
ENV BOT_TOKEN ""
ENV BOT_OAUTH_USER_TOKEN ""

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["topics-bot"]
