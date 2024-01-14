package models

type Voter struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	UserId      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserId" json:"Voter Detail"`
	CandidateId uint      `json:"candidate_id"`
	Candidate   Candidate `gorm:"foreignKey:CandidateId" json:"Voted Candidate"`
}
