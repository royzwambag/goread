package goodreads

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	goodreads "github.com/royzwambag/go-goodreads/goodreads/types"
)

// PaginatedAuthorList returns an xml response with a paginated list of an authors books
// Possible parameters are:
// id: Goodreads author id
// page: Integer
func PaginatedAuthorList(id int, page int) (goodreads.AuthorList, error) {
	var authorList goodreads.AuthorList

	url := "https://www.goodreads.com/author/list/"
	parameters := map[string]string{
		"id":   strconv.Itoa(id),
		"page": strconv.Itoa(page),
	}

	response, err := APIGet(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return authorList, err
	}
	xml.Unmarshal(response, &authorList)

	fmt.Println(authorList)
	return authorList, nil
}
