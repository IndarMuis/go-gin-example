package entity

type Book struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	PublishedYear string `json:"published_year"`
}
