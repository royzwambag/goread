package goread

import (
	"encoding/xml"
	"log"
	"strconv"
)

// AuthorInfo returns an AuthorList struct with a lot of information about the author, including their books
func AuthorInfo(id int) (AuthorList, error) {
	var authorList AuthorList

	url := "https://www.goodreads.com/author/show/"
	parameters := map[string]string{
		"id": strconv.Itoa(id),
	}

	response, err := Get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return authorList, err
	}
	xml.Unmarshal(response, &authorList)

	return authorList, nil
}
