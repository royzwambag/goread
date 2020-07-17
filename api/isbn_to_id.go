package goread

import (
	"log"
	"strings"

	"github.com/royzwambag/goread"
)

// ISBNToID returns the goodreads ID(s) of the given book(s) (found by ISBN)
// This function supports both ISBN and ISBN13
func ISBNToID(isbn ...string) ([]int, error) {
	var ids []int

	url := "https://www.goodreads.com/book/isbn_to_id/"
	parameters := map[string]string{
		"isbn": strings.Join(isbn[:], ","),
	}

	response, err := Get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return ids, err
	}

	stringIDs := strings.Split(string(response), ",")
	ids = goread.StringSliceToIntSlice(stringIDs)

	return ids, nil
}
