package goodreads

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	goodreads "github.com/royzwambag/go-read/goodreads/types"
)

// AuthorInfo returns an AuthorList struct with a lot of information about the author, including their books
func AuthorInfo(id int) (goodreads.AuthorList, error) {
	var authorList goodreads.AuthorList

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
