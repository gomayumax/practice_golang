package entity

type User struct {
	ID    uint   `json:"id" query:"id"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "users"
}
