package models

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"type:varchar(300)" json:"full name"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	Address  string `gorm:"type:varchar(300)" json:"address"`
	Gender   string `gorm:"type:varchar(300)" json:"gender"`
	UserName string `gorm:"type:varchar(300)" json:"userName"`
}
