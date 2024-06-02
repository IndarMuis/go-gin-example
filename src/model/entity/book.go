package entity

type Book struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	AuthorId      uint   `json:"author_id"`
	Category      string `json:"category"`
	CategoryId    uint   `json:"category_id"`
	PublishedYear string `json:"published_year"`
}

func (*Book) TableName() string {
	return "book"
}
