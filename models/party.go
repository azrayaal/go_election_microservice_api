package models

type Party struct {
	ID             int64  `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(300)" json:"Party Name"`
	Image          string `gorm:"type:varchar(300)" json:"Image"`
	Chairman       string `gorm:"type:varchar(300)" json:"Chairman"`
	Vision_mission string `gorm:"type:varchar(300)" json:"Vision Mission"`
	Address        string `gorm:"type:varchar(300)" json:"Address"`
}
