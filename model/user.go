package model

// User struct
type User struct {
	ID        string `json:"id"` // uuid
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  []byte `json:"-"` // password is not shown in json format
}
