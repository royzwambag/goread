package goread

import (
	"encoding/xml"
	"log"
	"strconv"
)

type nestedBook struct {
	Book Book `xml:"book"`
}

// BookInfo returns a Book struct with a lot of information about the book
func BookInfo(goodreadsID int) (Book, error) {
	var nestedBook nestedBook

	url := "https://www.goodreads.com/book/show/"
	parameters := map[string]string{
		"id": strconv.Itoa(goodreadsID),
	}

	response, err := get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return nestedBook.Book, err
	}
	xml.Unmarshal(response, &nestedBook)

	return nestedBook.Book, nil
}
