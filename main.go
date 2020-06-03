package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/vald-phoenix/xkcd/client"
)

func main() {
	// Read command line arguments
	comicNo := flag.Int(
		"n", int(client.LatestComic), "Comic number to fetch (default latest)",
	)
	clientTimeout := flag.Int64(
		"t",
		int64(client.DefaultClientTimeout.Seconds()),
		"Client timeout in seconds",
	)
	saveImage := flag.Bool(
		"s", false, "Save image to current directory",
	)
	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)
	flag.Parse()

	// Instantiate the XKCDClient
	xkcdClient := client.NewXKCDClient()
	xkcdClient.SeetTimeout(time.Duration(*clientTimeout) * time.Second)

	// Fetch from API using the XKCDClient
	comic, err := xkcdClient.Fetch(client.ComicNumber(*comicNo), *saveImage)
	if err != nil {
		log.Println(err)
	}

	// Print output either in JSON or plain text
	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}
