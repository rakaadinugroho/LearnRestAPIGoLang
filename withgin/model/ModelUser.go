package model
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
