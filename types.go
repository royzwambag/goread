package goread

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

// BookReviewStatistic is a struct that information about the reviews
type BookReviewStatistic struct {
	ID               int     `json:"id"`
	ISBN             string  `json:"isbn"`
	ISBN13           string  `json:"isbn13"`
	AverageRating    float32 `json:"average_rating"`
	RatingsCount     int     `json:"ratings_count"`
	ReviewsCount     int     `json:"reviews_count"`
	TextReviewsCount int     `json:"text_reviews_count"`
	WorkRatingsCount int     `json:"work_ratings_count"`
	WorkReviewsCount int     `json:"work_reviews_count"`
}
