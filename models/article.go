package models

type Article struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"type:varchar(300)" json:"title"`
	Content string `gorm:"type:varchar(300)" json:"content"`
}
