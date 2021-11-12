package model

import "time"

// The database indicates custom, the default is the plural form of model, for example, the default is users here
func (User) TableName() string {
	return "users"
}

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `gorm:"not null;default:null" json:"username"`
	Email     string    `gorm:"not null;default:null" json:"email"`
	Password  string    `gorm:"not null;default:null" json:"password"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"createdAt"`
}
