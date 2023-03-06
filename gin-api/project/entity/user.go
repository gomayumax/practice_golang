package entity

type User struct {
	ID uint `json:"id"`
}

func (User) TableName() string {
	return "users"
}
