package entity

type User struct {
	ID       uint   `json:"id"`
	Nrp      string `json:"nrp"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
