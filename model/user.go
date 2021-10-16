package model

type User struct {
	Id 		 uint	`json:"id"`
	Name 	 string	`json:"name"`
	Email 	 string `json:"email" gorm:"uniqueIndex"`
	Password []byte	`json:"-"` // password is not shown in json format
}
