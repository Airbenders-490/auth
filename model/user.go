package model

type User struct {
	Id 		 	uint	`json:"id"`
	FirstName 	string 	`json:"first_name"`
	LastName	string 	`json:"last_name"`
	Email 	 	string 	`json:"email" gorm:"uniqueIndex"`
	Password 	[]byte	`json:"-"` // password is not shown in json format
}
