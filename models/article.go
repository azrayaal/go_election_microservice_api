package models

type Article struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"type:varchar(300)" json:"title"`
	Image   string `gorm:"type:varchar(300)" json:"image"`
	Content string `gorm:"type:varchar(300)" json:"content"`
	Date    string `gorm:"type:varchar(300)" json:"date"`
	Author  string `gorm:"type:varchar(300)" json:"author"`
	UserId  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserId" json:"author detail"`
}
