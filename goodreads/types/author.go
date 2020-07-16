package goodreads

// AuthorList is a struct that contains all data of an author
type AuthorList struct {
	Author Author `xml:"author"`
}

// Author is a struct that contains all data of an author
type Author struct {
	ID       int      `xml:"id"`
	Name     string   `xml:"name"`
	Link     string   `xml:"link"`
	BookList BookList `xml:"books"`
}
