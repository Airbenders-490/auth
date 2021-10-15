package model

type User struct {
	Id 		 uint
	Name 	 string
	Email 	 string `gorm:"uniqueIndex"`
	Password string
}
