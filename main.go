package main

import (
	goodreads "github.com/royzwambag/go-read/goodreads/api"
)

func main() {
	goodreads.PaginatedAuthorList(2917920, 1)
}
