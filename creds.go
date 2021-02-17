package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"
)

type creds struct {
	ConsumerKey    string `yaml:"ConsumerKey"`
	ConsumerSecret string `yaml:"ConsumerSecret"`
	AccessToken    string `yaml:"AccessToken"`
	AccessSecret   string `yaml:"AccessSecret"`
}

func getCreds() *creds {
	return &creds{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_SECRET"),
	}
}

func (c creds) tweet(text string) error {
	if c.ConsumerKey == "" ||
		c.ConsumerSecret == "" ||
		c.AccessToken == "" ||
		c.AccessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	tweet, resp, err := client.Statuses.Update(text, nil)
	if err != nil {
		return err
	}
	log.Infof("tweet: %v, resp: %v", tweet, resp.Body)

	return nil
}
