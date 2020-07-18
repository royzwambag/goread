package goread

import (
	"encoding/xml"
	"log"
	"strconv"
)

// Author is a struct that contains all data of an author
type Author struct {
	ID              int    `xml:"id"`
	Name            string `xml:"name"`
	About           string `xml:"about"`
	Link            string `xml:"link"`
	Fans            int    `xml:"fans_count"`
	Followers       int    `xml:"author_followers_count"`
	ImageURL        string `xml:"image_url"`
	LargeImageURL   string `xml:"large_image_url"`
	SmallImageURL   string `xml:"small_image_url"`
	Influences      string `xml:"influences"`
	WorksWritten    int    `xml:"works_count"`
	Gender          string `xml:"gender"`
	Hometown        string `xml:"hometown"`
	BornAt          string `xml:"born_at"`
	DiedAt          string `xml:"died_at"`
	GoodreadsAuthor bool   `xml:"goodreads_author"`
	UserID          int    `xml:"user>id"`
	Books           []Book `xml:"books>book"`
}

type nestedAuthor struct {
	Author Author `xml:"author"`
}

// AuthorInfo returns an Author struct with a lot of information about the author, including their books
func (client *Client) AuthorInfo(id int) (Author, error) {
	var nestedAuthor nestedAuthor
	var author Author

	url := "https://www.goodreads.com/author/show/"
	parameters := map[string]string{
		"id": strconv.Itoa(id),
	}

	response, err := client.get(url, parameters)
	if err != nil {
		log.Fatalln(err)
		return author, err
	}
	xml.Unmarshal(response, &nestedAuthor)

	return nestedAuthor.Author, nil
}
