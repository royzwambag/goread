# goread
A simple goodreads api consumer written in go. 
Note: This does not contain all api calls (yet). The goodreads api has a mixture of both XML and JSON responses.

# Installation

```sh
go get github.com/royzwambag/goread
```

# Usage

```go
package main

import (
    "fmt"

    "github.com/royzwambag/goread"
)

func main() {
    // Get information from an author with given goodreads id
    // This method will return an `Author` struct which contains
    // all the data about the author
    goread.AuthorInfo(18541)

    // Get the goodreads ID of one or more books with given isbns
    // This function accepts both isbn and isbn13
    goread.ISBNToID("0062565710", "9780596802813")
    // returns int[]{34017076, 6356381}

    // Get review statistics of one or multiple books with given isbns
    // This function accepts both isbn and isbn13
    goread.BookReviewStatistics("0062565710", "9780596802813")
    // returns BookReviewStatistics[]{}
```