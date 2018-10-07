package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type EnvConfig struct {
	BotToken          string
	VerificationToken string
	BotID             string
	ChannelID         string
}

func SetEnv() (*EnvConfig, error) {
	channelID, err := lookUpEnv("CHANNEL_ID")
	if err != nil {
		return nil, errors.Wrap(err, "channelID not found")
	}
	botID, err := lookUpEnv("BOT_ID")
	if err != nil {
		return nil, errors.Wrap(err, "bot id not found")
	}
	verificationToken, err := lookUpEnv("VERIFICATION_TOKEN")
	if err != nil {
		return nil, errors.Wrap(err, "verification token not found")
	}
	botToken, err := lookUpEnv("BOT_OAUTH_USER_TOKEN")
	if err != nil {
		return nil, errors.Wrap(err, "bot token not found")
	}

	return &EnvConfig{
		ChannelID:         channelID,
		BotID:             botID,
		VerificationToken: verificationToken,
		BotToken:          botToken,
	}, nil
}

func lookUpEnv(key string) (string, error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s is invalid token key", key)
	}
	return env, nil
}
