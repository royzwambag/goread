package goread

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	goread "github.com/royzwambag/go-read/types"
)

// AuthorInfo returns an AuthorList struct with a lot of information about the author, including their books
func AuthorInfo(id int) (goread.AuthorList, error) {
	var authorList goread.AuthorList

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

	fmt.Println(authorList)
	return authorList, nil
}
