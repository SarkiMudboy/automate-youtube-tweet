package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"google.golang.org/api/youtube/v3"
)

var (
	maxResults = flag.Int64("maxResults", 5, "The maximum number of playlist resources to include in the API response.")
	part = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	myRating = flag.String("rating", "", "Retrieve all my liked videos")
)

func playlistsList(service *youtube.Service, part []string, maxResults int64, rating string) *youtube.VideoListResponse {

	call := service.Videos.List(part)

	if rating != "" {
		call = call.MyRating(rating)
	}

	if maxResults != 0 {
		call = call.MaxResults(maxResults)
	}

	response, err := call.Do()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	return response
}

// TODO: Add a json file storage to ensure I don't duplicate a tweet -> store {'id': "vid title"}

func main() {

	var parts []string
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(1)

	flag.Parse()
	wg.Add(1)

	client := getClient(youtube.YoutubeReadonlyScope)
	service, err := youtube.New(client)

	if err != nil {
			log.Fatalf("Error creating YouTube client: %v", err)
	}

	if *part != "" {
		parts = strings.Split(*part, ",")
	}

	response := playlistsList(service, parts, *maxResults, *myRating)

	for _, playlist := range response.Items {

		playlistId := playlist.Id
		playlistTitle := playlist.Snippet.Title
		link := fmt.Sprintf("https://www.youtube.com/watch?v=%s", playlistId)
		tweet := fmt.Sprintf("TIL\n%s\n%s", playlistTitle, link)

		fmt.Printf("%s:%s link -> %s\n", playlistId, playlistTitle, link)
		go makeTweet(tweet, &wg)

	}
	wg.Wait()
	fmt.Println("\nTerminating Program")

}