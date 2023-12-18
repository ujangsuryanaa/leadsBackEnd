package repositories

import (
	"leadsmanagementsystem/models"

	"gorm.io/gorm"
)

type LeadRepository interface {
	FindLeads() ([]models.Lead, error)
	GetLead(ID int) (models.Lead, error)
	CreateLead(lead models.Lead) (models.Lead, error)
	UpdateLead(lead models.Lead) (models.Lead, error)
	DeleteLead(lead models.Lead) error
	// SearchLeads(searchText, startDate, endDate string, status, branchOffices []string) ([]models.Lead, error)
	SearchLeads(searchText, startDate, endDate string, status, branchOffices []string, probability float64, leadChannels, leadMedia, leadSources []string) ([]models.Lead, error)
}

type leadRepository struct {
	db *gorm.DB
}

func RepositoryLead(db *gorm.DB) *leadRepository {
	return &leadRepository{db}
}

func (r *leadRepository) FindLeads() ([]models.Lead, error) {
	var leads []models.Lead
	err := r.db.Find(&leads).Error

	return leads, err
}

func (r *leadRepository) GetLead(ID int) (models.Lead, error) {
	var lead models.Lead
	err := r.db.First(&lead, ID).Error

	return lead, err
}

func (r *leadRepository) CreateLead(lead models.Lead) (models.Lead, error) {
	err := r.db.Create(&lead).Error

	return lead, err
}

func (r *leadRepository) UpdateLead(lead models.Lead) (models.Lead, error) {
	err := r.db.Save(&lead).Error

	return lead, err
}

func (r *leadRepository) DeleteLead(lead models.Lead) error {
	err := r.db.Delete(&lead).Error

	return err
}

// func (r *leadRepository) SearchLeads(searchText, startDate, endDate string, status, branchOffices []string) ([]models.Lead, error) {
// 	var leads []models.Lead

// 	query := r.db.Model(&models.Lead{})

// 	if searchText != "" {
// 		query = query.Where("fullname LIKE ? OR email LIKE ?", "%"+searchText+"%", "%"+searchText+"%")
// 	}

// 	if startDate != "" && endDate != "" {
// 		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
// 	}

// 	if len(status) > 0 {
// 		query = query.Where("status IN (?)", status)
// 	}

// 	if len(branchOffices) > 0 {
// 		query = query.Where("branch_office IN (?)", branchOffices)
// 	}

// 	err := query.Find(&leads).Error

// 	return leads, err
// }

func (r *leadRepository) SearchLeads(
	searchText, startDate, endDate string,
	status, branchOffices []string,
	probability float64,
	leadChannels, leadMedia, leadSources []string,
) ([]models.Lead, error) {
	var leads []models.Lead

	query := r.db.Model(&models.Lead{})

	if searchText != "" {
		query = query.Where("fullname LIKE ? OR email LIKE ?", "%"+searchText+"%", "%"+searchText+"%")
	}

	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	if len(status) > 0 {
		query = query.Where("status IN (?)", status)
	}

	if len(branchOffices) > 0 {
		query = query.Where("branch_office IN (?)", branchOffices)
	}

	if probability > 0 {
		query = query.Where("probability = ?", probability)
	}

	if len(leadChannels) > 0 {
		query = query.Where("lead_channel IN (?)", leadChannels)
	}

	if len(leadMedia) > 0 {
		query = query.Where("lead_media IN (?)", leadMedia)
	}

	if len(leadSources) > 0 {
		query = query.Where("lead_source IN (?)", leadSources)
	}

	err := query.Find(&leads).Error

	return leads, err
}

// Di antarmuka LeadRepository
// type LeadRepository interface {
//     // ... fungsi lainnya

//     SearchLeads(
//         searchText, dateFrom, dateTo string,
//         status, branchOffices []string,
//         probability float64,
//         leadChannels, leadMedia, leadSources []string,
//     ) ([]models.Lead, error)
// }

// // Di implementasi leadRepository
// func (r *leadRepository) SearchLeads(
//     searchText, dateFrom, dateTo string,
//     status, branchOffices []string,
//     probability float64,
//     leadChannels, leadMedia, leadSources []string,
// ) ([]models.Lead, error) {
//     var leads []models.Lead

//     query := r.db.Model(&models.Lead{})

//     if searchText != "" {
//         query = query.Where("fullname LIKE ? OR email LIKE ?", "%"+searchText+"%", "%"+searchText+"%")
//     }

//     if dateFrom != "" && dateTo != "" {
//         query = query.Where("created_at BETWEEN ? AND ?", dateFrom, dateTo)
//     }

//     if len(status) > 0 {
//         query = query.Where("status IN (?)", status)
//     }

//     if len(branchOffices) > 0 {
//         query = query.Where("branch_office IN (?)", branchOffices)
//     }

//     if probability > 0 {
//         query = query.Where("probability = ?", probability)
//     }

//     if len(leadChannels) > 0 {
//         query = query.Where("lead_channel IN (?)", leadChannels)
//     }

//     if len(leadMedia) > 0 {
//         query = query.Where("lead_media IN (?)", leadMedia)
//     }

//     if len(leadSources) > 0 {
//         query = query.Where("lead_source IN (?)", leadSources)
//     }

//     err := query.Find(&leads).Error

//     return leads, err
// }
