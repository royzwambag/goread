# goread
A simple goodreads api consumer written in go. 
Note: This does not contain all api calls (yet). The goodreads api has a mixture of both XML and JSON responses.

# Installation

```sh
go get github.com/royzwambag/goread
```

# Usage

Set your goodreads developer key in the `GOODREADS_DEVELOPER_KEY` environment variable for easiest use

```go
package main

import (
    "github.com/royzwambag/goread"
)

func main() {
    // Create a Client with NewClient() to create a Client struct with
    // the developer key from your `GOODREADS_DEVELOPER_KEY` environment variable
    client := goread.NewClient()
    // Now you can use this client to make calls to the api

    // Alternatively, you can create your own Client by supplying a http.client
    // as wel as the developer key
    httpClient := http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
	client := Client{
		http: &httpClient,
		key:  "ABCDEF",
	}

    // Get information from an author with given goodreads id
    // This method will return an `Author` struct which contains
    // all the data about the author
    client.AuthorInfo(18541)

    // Get the goodreads ID of one or more books with given isbns
    // This function accepts both isbn and isbn13
    client.ISBNToID("0062565710", "9780596802813")
    // returns int[]{34017076, 6356381}

    // Get review statistics of one or multiple books with given isbns
    // This function accepts both isbn and isbn13
    client.BookReviewStatistics("0062565710", "9780596802813")
    // returns BookReviewStatistics[]{}
```