package model

import (
	"fmt"
	"strings"
)

type User struct {
	Username string `json:"username" gorm:"username"`
	Email    string `json:"email" gorm:"email"`
}

// alias name of struct
type PostItem struct {
	Title string `json:"title" gorm:"title"`
}

func (item PostItem) TableName() string {
	return "posting"
}

func (u User) ValidationUser() error {
	if len(strings.Trim(u.Username, " ")) < 1 { // checked Trim
		return fmt.Errorf("masukan username")
	}
	if len(strings.Trim(u.Email, " ")) < 1 {
		return fmt.Errorf("masukan email")
	}
	return nil
}