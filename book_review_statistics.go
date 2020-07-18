package goread

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// BookReviewStatistic is a struct that information about the reviews
type BookReviewStatistic struct {
	ID               int     `json:"id"`
	ISBN             string  `json:"isbn"`
	ISBN13           string  `json:"isbn13"`
	AverageRating    float32 `json:"average_rating"`
	RatingsCount     int     `json:"ratings_count"`
	ReviewsCount     int     `json:"reviews_count"`
	TextReviewsCount int     `json:"text_reviews_count"`
	WorkRatingsCount int     `json:"work_ratings_count"`
	WorkReviewsCount int     `json:"work_reviews_count"`
}

type nestedStatistics struct {
	BookReviewStatistics []BookReviewStatistic `json:"books"`
}

// BookReviewStatistics returns the statistics about given books
// This function supports both ISBN and ISBN13
func (client *Client) BookReviewStatistics(isbn ...string) ([]BookReviewStatistic, error) {
	var statistics nestedStatistics

	url := "https://www.goodreads.com/book/review_counts/"
	parameters := map[string]string{
		"isbns":  strings.Join(isbn[:], ","),
		"format": "json",
	}

	response, err := client.get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return statistics.BookReviewStatistics, err
	}

	json.Unmarshal(response, &statistics)
	fmt.Println(statistics)

	return statistics.BookReviewStatistics, nil
}
