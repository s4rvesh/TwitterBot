package main

import (
	"TwitterBot/handler"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Go-Twitter Bot v0.01")

	creds := handler.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := handler.GetClient(&creds)

	if err != nil {
		log.Println("Error getting twitter client")
		log.Println(err)
	}

	fmt.Printf("%+v\n", client)

	tweet, resp, err := client.Statuses.Update("Hey, if you see this Tweet, it's tweeted from a tweetBot I'm building. Yay its working. Version 0.1. PEACE OUT", nil)

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

}
