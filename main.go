package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

const (
	endpoint = "https://hooks.slack.com/services/"
)

type Slack struct {
	client    *slack.Client
	botID     string
	channelID string
}

func run() int {
	env, err := SetEnv()
	if err != nil {
		log.Println(err)
		return 1
	}

	fmt.Println(env)

	api := slack.New(env.BotToken)
	s := &Slack{
		client:    api,
		botID:     env.BotID,
		channelID: env.ChannelID,
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			log.Println("[INFO] get message")
			if err := s.TopicsPostHandler(event); err != nil {
				log.Println("[ERROR] message handler failed: %v\n", err)
			}
		}
	}
	return 0
}

func main() {
	os.Exit(run())
}
