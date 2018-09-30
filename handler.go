package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

func setParams() slack.PostMessageParameters {
	return slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			slack.Attachment{},
		},
	}
}

func (s *Slack) TopicsPostHandler(event *slack.MessageEvent) error {
	if event.Channel != s.channelID {
		log.Println("expect: %s\nactual: %s", event.Channel, s.channelID)
		return nil
	}

	if !strings.HasPrefix(event.Msg.Text, fmt.Sprintf("<@%s> ", s.botID)) {
		return nil
	}

	m := strings.Split(strings.TrimSpace(event.Msg.Text), " ")[1:]
	var label string
	switch m[0] {
	case "hey":
		label = "hello, help?"
	default:
		return nil
	}

	if _, _, err := s.client.PostMessage(event.Channel, label, setParams()); err != nil {
		return fmt.Errorf("[ERROR] post message failed: %s", err)
	}

	return nil
}
