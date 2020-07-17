package goread

import (
	"encoding/xml"
	"log"
	"strconv"
)

type xmlResponse struct {
	Author Author `xml:"author"`
}

// AuthorInfo returns an Author struct with a lot of information about the author, including their books
func AuthorInfo(id int) (Author, error) {
	var xmlResponse xmlResponse
	var author Author

	url := "https://www.goodreads.com/author/show/"
	parameters := map[string]string{
		"id": strconv.Itoa(id),
	}

	response, err := get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return author, err
	}
	xml.Unmarshal(response, &xmlResponse)

	return xmlResponse.Author, nil
}
