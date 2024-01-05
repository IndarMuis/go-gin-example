package dto

type BookResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	PublishedYear string `json:"published_year"`
}

type BookRequest struct {
	Title         string `json:"title" binding:"required"`
	Author        string `json:"author" binding:"required"`
	Category      string `json:"category" binding:"required"`
	PublishedYear string `json:"published_year" binding:"required"`
}
