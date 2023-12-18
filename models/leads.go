package models

import "time"

type Lead struct {
	ID           int           `json:"id"`
	BranchOffice string        `json:"branch_office" gorm:"type: varchar(255)"`
	Fullname     string        `json:"fullname" gorm:"type: varchar(255)"`
	Email        string        `json:"email" gorm:"type: varchar(255)"`
	Phone        string        `json:"phone" gorm:"type: varchar(20)"`
	Address      string        `json:"address" gorm:"type: text"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	CompanyName  string        `json:"company_name" gorm:"type: varchar(255)"`
	Status       string        `json:"status" gorm:"type: varchar(50)"`
	Probability  float64       `json:"probability"`
	LeadType     string        `json:"lead_type" gorm:"type: varchar(50)"`
	LeadChannel  string        `json:"lead_channel" gorm:"type: varchar(50)"`
	LeadMedia    string        `json:"lead_media" gorm:"type: varchar(50)"`
	LeadSource   string        `json:"lead_source" gorm:"type: varchar(50)"`
	GeneralNotes string        `json:"general_notes" gorm:"type: text"`
	UserID       int           `json:"user_id" form:"user_id"`
	User         UsersResponse `json:"user"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

type LeadResponse struct {
	ID           int           `json:"id"`
	BranchOffice string        `json:"branch_office"`
	Fullname     string        `json:"fullname"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Address      string        `json:"address"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	CompanyName  string        `json:"company_name"`
	Status       string        `json:"status"`
	Probability  float64       `json:"probability"`
	LeadType     string        `json:"lead_type"`
	LeadChannel  string        `json:"lead_channel"`
	LeadMedia    string        `json:"lead_media"`
	LeadSource   string        `json:"lead_source"`
	GeneralNotes string        `json:"general_notes"`
	UserID       int           `json:"-"`
	User         UsersResponse `json:"user"`
}

func (LeadResponse) TableName() string {
	return "leads"
}
