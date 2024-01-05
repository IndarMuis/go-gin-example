package dto

type BookResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	PublishedYear string `json:"published_year"`
}

type BookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	PublishedYear string `json:"published_year"`
}
