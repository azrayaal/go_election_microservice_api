package models

type Candidate struct {
	ID             int64  `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(300)" json:"Candidate Name"`
	Image          string `gorm:"type:varchar(300)" json:"Image"`
	Number         int64  `gorm:"type:varchar(300)" json:"Number"`
	Vision_mission string `gorm:"type:varchar(300)" json:"Vision Mission"`
	PartyID        uint   `json:"party_id"`
	Party          Party  `gorm:"foreignKey:PartyID" json:"Party Data"`
}
