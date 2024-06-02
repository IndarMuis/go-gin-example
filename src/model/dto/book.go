package dto

type BookResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author,omitempty"`
	AuthorId      uint   `json:"author_id"`
	Category      string `json:"category,omitempty"`
	CategoryId    uint   `json:"category_id"`
	PublishedYear string `json:"published_year,omitempty"`
}

type BookRequest struct {
	Title         string `json:"title" binding:"required"`
	CategoryId    uint   `json:"category_id" binding:"required"`
	AuthorId      uint   `json:"author_id" binding:"required"`
	PublishedYear string `json:"published_year" binding:"required"`
}
