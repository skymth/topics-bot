package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"github.com/skymth/topics-bot/topics"
)

func setParams(topic topics.Topic) slack.PostMessageParameters {
	return slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			slack.Attachment{
				Color:     "#5df91b",
				Title:     topic.Title,
				TitleLink: topic.URL,
				Text:      topic.Description,
			},
		},
	}
}

func (s *Slack) send(event, label string, params slack.PostMessageParameters) error {
	if _, _, err := s.client.PostMessage(event, label, params); err != nil {
		return err
	}
	return nil
}

func (s *Slack) TopicsPostHandler(event *slack.MessageEvent) error {
	if event.Channel != s.channelID {
		log.Println("expect: %s\nactual: %s", event.Channel, s.channelID)
		return nil
	}

	if !strings.HasPrefix(event.Msg.Text, fmt.Sprintf("リマインダー : <@%s> ", s.botID)) {
		log.Println(event.Msg.Text)
		log.Println(s.botID)
		return nil
	}

	m := strings.Split(strings.TrimSpace(event.Msg.Text), " ")[1:]

	switch m[2] {
	case "todays":
		topics, err := topics.GetTopics()
		if err != nil {
			return errors.Wrap(err, "get topics err")
		}
		for _, topic := range topics {
			if err := s.send(event.Channel, "", setParams(topic)); err != nil {
				return errors.Wrap(err, "send topics err")
			}
		}

	default:
		return nil
	}

	return nil
}
