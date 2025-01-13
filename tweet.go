package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/dghubble/oauth1"
)

var (
	consumerKey string
	consumerSecret string
	accessToken string
	accessSecret string

	config *oauth1.Config
	token *oauth1.Token
)

func init () {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	consumerKey = os.Getenv("TWIITER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWIITER_CONSUMER_SECRET")
	accessToken = os.Getenv("TWIITER_ACCESS_TOKEN")
	accessSecret = os.Getenv("TWIITER_ACCESS_SECRET")

	config = oauth1.NewConfig(consumerKey, consumerSecret)
	token = oauth1.NewToken(accessToken, accessSecret)
	fmt.Println(consumerKey)
}

func makeTweet(text string, w *sync.WaitGroup) {

	defer w.Done()

	httpClient := config.Client(oauth1.NoContext, token)

	buff := bytes.NewBuffer(nil)
	url := "https://api.twitter.com/2/tweets"

	raw := map[string]string{
		"text": text,
	}
	tweet, err := json.Marshal(raw)
	if err != nil {
        fmt.Fprint(os.Stderr, err)
    }

	reader := bytes.NewReader([]byte(tweet))
	if _, err := io.Copy(buff, reader); err != nil {
		log.Fatal(err)
	}

	response, err := httpClient.Post(url, "application/json", buff)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
    }
    defer response.Body.Close()
    body, err := io.ReadAll(response.Body)
    if err != nil {
		fmt.Fprint(os.Stderr, err)
    }
	response.Body.Close()
    fmt.Println(string(body))

	return
}