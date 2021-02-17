package main

import (
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func main() {
	kids, err := kidsTheseDays()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Infof("%v", kids)

	creds := getCreds()
	creds.tweet(kids)
}
