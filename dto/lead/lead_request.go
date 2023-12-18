package leaddto

type LeadRequest struct {
	ID           int     `json:"id"`
	BranchOffice string  `json:"branch_office" gorm:"type: varchar(255)"`
	Fullname     string  `json:"fullname" gorm:"type: varchar(255)"`
	Email        string  `json:"email" gorm:"type: varchar(255)"`
	Phone        string  `json:"phone" gorm:"type: varchar(20)"`
	Address      string  `json:"address" gorm:"type: text"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	CompanyName  string  `json:"company_name" gorm:"type: varchar(255)"`
	Status       string  `json:"status" gorm:"type: varchar(50)"`
	Probability  float64 `json:"probability"`
	LeadType     string  `json:"lead_type" gorm:"type: varchar(50)"`
	LeadChannel  string  `json:"lead_channel" gorm:"type: varchar(50)"`
	LeadMedia    string  `json:"lead_media" gorm:"type: varchar(50)"`
	LeadSource   string  `json:"lead_source" gorm:"type: varchar(50)"`
	GeneralNotes string  `json:"general_notes" gorm:"type: text"`
	UserID       int     `json:"user_id" form:"user_id"`
}
