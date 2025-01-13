## About

Small script to automate tweeting links to youtube videos to my twitter account.

## Get started

Make sure you are in go version==1.23.0.

run to verify

```shell
go version
```

clone the project

```shell
https://github.com/SarkiMudboy/automate-youtube-tweet.git
```

Build the binary

```
go build -o tweettube.exe
```

or
run

```shell
go run main.go oauth.go tweet.go
```

### Usage

- myRating -> retrieves all liked videos i.e. `tweettube -myRating`
- maxResults -> The maximum number of videos to include in the API response i.e. `tweettube -maxResults 5` (5 videos)
- part -> Comma-separated list of playlist resource parts that API response will include i.e. `tweettube -part snippet`
