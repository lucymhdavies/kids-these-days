package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type RandomWikiArticles struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Rncontinue string `json:"rncontinue"`
		Continue   string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Random []struct {
			ID    int    `json:"id"`
			Ns    int    `json:"ns"`
			Title string `json:"title"`
		} `json:"random"`
	} `json:"query"`
}

func wikiArticles() ([]string, error) {
	min := 2 // >=
	max := 3 // <
	rand.Seed(time.Now().UTC().UnixNano())
	numArticles := rand.Intn(max-min) + min
	log.Infof("Generating %d articles", numArticles)

	apiURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=random&format=json&rnnamespace=0&rnlimit=%d", numArticles)
	resp, err := http.Get(apiURL)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	var randomWikiArticles RandomWikiArticles
	err = json.NewDecoder(resp.Body).Decode(&randomWikiArticles)
	if err != nil {
		return []string{}, err
	}

	var articles []string
	for _, article := range randomWikiArticles.Query.Random {
		articles = append(articles, article.Title)
	}

	return articles, nil
}

func kidsTheseDays() (string, error) {
	articles, err := wikiArticles()
	if err != nil {
		return "", err
	}

	andYour := strings.Join(articles, ", and your ")

	return "You kids these days and your " + andYour, nil
}
