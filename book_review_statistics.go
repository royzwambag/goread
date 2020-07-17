package goread

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type nestedStatistics struct {
	BookReviewStatistics []BookReviewStatistic `json:"books"`
}

// BookReviewStatistics returns the statistics about given books
// This function supports both ISBN and ISBN13
func BookReviewStatistics(isbn ...string) ([]BookReviewStatistic, error) {
	var statistics nestedStatistics

	url := "https://www.goodreads.com/book/review_counts/"
	parameters := map[string]string{
		"isbns":  strings.Join(isbn[:], ","),
		"format": "json",
	}

	response, err := get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return statistics.BookReviewStatistics, err
	}

	json.Unmarshal(response, &statistics)
	fmt.Println(statistics)

	return statistics.BookReviewStatistics, nil
}
