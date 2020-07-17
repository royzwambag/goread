package goodreads

// Book is a struct that contains all data of a book
type Book struct {
	ID                 int     `xml:"id"`
	ISBN               string  `xml:"isbn"`
	ISBN13             string  `xml:"isbn13"`
	TextReviewsCount   int     `xml:"text_reviews_count"`
	URI                string  `xml:"uri"`
	Title              string  `xml:"title"`
	TitleWithoutSeries string  `xml:"title_without_series"`
	ImageURL           string  `xml:"image_url"`
	SmallImageURL      string  `xml:"small_image_url"`
	LargeImageURL      string  `xml:"large_image_url"`
	Link               string  `xml:"link"`
	NumPages           int     `xml:"num_pages"`
	Format             string  `xml:"format"`
	EditionInformation string  `xml:"edition_information"`
	Publisher          string  `xml:"publisher"`
	PublicationDay     int     `xml:"publication_day"`
	PublicationMonth   int     `xml:"publication_month"`
	PublicationYear    int     `xml:"publication_year"`
	AverageRating      float32 `xml:"average_rating"`
	RatingsCount       int     `xml:"ratings_count"`
	Description        string  `xml:"description"`
}
