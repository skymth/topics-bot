package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/skymth/topics-bot/topics"
)

const (
	endpoint = "https://hooks.slack.com/services/"
)

var (
	accessToken = os.Getenv("SLACK_TOKEN")
)

func sendTopics() error {
	//TODO ここでdataを添付
	data := `{"text":"` + "" + `"}`
	body := bytes.NewBuffer([]byte(data))

	url := endpoint + accessToken
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	topics, err := topics.GetTopics()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(topics)
}
