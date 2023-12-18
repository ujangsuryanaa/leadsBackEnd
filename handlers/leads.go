package handlers

import (
	"encoding/json"
	"errors"
	leaddto "leadsmanagementsystem/dto/lead"
	dto "leadsmanagementsystem/dto/result"
	"leadsmanagementsystem/models"
	"net/http"
	"strconv"

	"leadsmanagementsystem/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handlerLead struct {
	LeadRepository repositories.LeadRepository
}

func HandlerLead(LeadRepository repositories.LeadRepository) *handlerLead {
	return &handlerLead{LeadRepository}
}

func (h *handlerLead) FindLeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	leads, err := h.LeadRepository.FindLeads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: leads}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLead) GetLead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var lead models.Lead
	lead, err := h.LeadRepository.GetLead(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseLead(lead)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLead) CreateLead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(leaddto.LeadRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	lead := models.Lead{

		BranchOffice: request.BranchOffice,
		Fullname:     request.Fullname,
		Email:        request.Email,
		Phone:        request.Phone,
		Address:      request.Address,
		Latitude:     request.Latitude,
		Longitude:    request.Longitude,
		CompanyName:  request.CompanyName,
		Status:       request.Status,
		Probability:  request.Probability,
		LeadType:     request.LeadType,
		LeadChannel:  request.LeadChannel,
		LeadMedia:    request.LeadMedia,
		LeadSource:   request.LeadSource,
		GeneralNotes: request.GeneralNotes,
		UserID:       request.UserID,
	}

	lead, err = h.LeadRepository.CreateLead(lead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	lead, _ = h.LeadRepository.GetLead(lead.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: lead}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLead) UpdateLead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(models.Lead)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	lead, err := h.LeadRepository.GetLead(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if (request.BranchOffice) != "" {
		lead.BranchOffice = request.BranchOffice
	}

	if (request.Fullname) != "" {
		lead.Fullname = request.Fullname
	}

	if (request.Email) != "" {
		lead.Email = request.Email
	}

	if (request.Phone) != "" {
		lead.Phone = request.Phone
	}

	if (request.Address) != "" {
		lead.Address = request.Address
	}

	if (request.Latitude) != 0 {
		lead.Latitude = request.Latitude
	}

	if (request.Longitude) != 0 {
		lead.Longitude = request.Longitude
	}

	if (request.CompanyName) != "" {
		lead.CompanyName = request.CompanyName
	}

	if (request.Status) != "" {
		lead.Status = request.Status
	}

	if (request.Probability) != 0 {
		lead.Probability = request.Probability
	}

	if (request.LeadType) != "" {
		lead.LeadType = request.LeadType
	}

	if (request.LeadChannel) != "" {
		lead.LeadChannel = request.LeadChannel
	}

	if (request.LeadMedia) != "" {
		lead.LeadMedia = request.LeadMedia
	}

	if (request.LeadSource) != "" {
		lead.LeadSource = request.LeadSource
	}

	if (request.GeneralNotes) != "" {
		lead.GeneralNotes = request.GeneralNotes
	}

	updatedLead, err := h.LeadRepository.UpdateLead(lead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseLead(updatedLead)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLead) DeleteLead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	lead, err := h.LeadRepository.GetLead(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			response := dto.ErrorResult{Code: http.StatusNotFound, Message: "Lead not found"}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.LeadRepository.DeleteLead(lead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Message: "Lead deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerLead) SearchLeads(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Extract search parameters from the request
// 	searchText := r.URL.Query().Get("search_text")
// 	startDate := r.URL.Query().Get("start_date")
// 	endDate := r.URL.Query().Get("end_date")
// 	status := r.URL.Query()["status"]
// 	branchOffices := r.URL.Query()["branch_office"]

// 	leads, err := h.LeadRepository.SearchLeads(searchText, startDate, endDate, status, branchOffices)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: leads}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerLead) SearchLeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract search parameters from the request
	searchText := r.URL.Query().Get("search_text")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	status := r.URL.Query()["status"]
	branchOffices := r.URL.Query()["branch_office"]
	probabilityStr := r.URL.Query().Get("probability")
	leadChannels := r.URL.Query()["lead_channel"]
	leadMedia := r.URL.Query()["lead_media"]
	leadSources := r.URL.Query()["lead_source"]

	// Convert probability from string to float64
	probability, _ := strconv.ParseFloat(probabilityStr, 64)

	leads, err := h.LeadRepository.SearchLeads(
		searchText, startDate, endDate, status, branchOffices,
		probability, leadChannels, leadMedia, leadSources,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: leads}
	json.NewEncoder(w).Encode(response)
}

func convertResponseLead(u models.Lead) models.LeadResponse {
	return models.LeadResponse{
		BranchOffice: u.BranchOffice,
		Fullname:     u.Fullname,
		Email:        u.Email,
		Phone:        u.Phone,
		Address:      u.Address,
		Latitude:     u.Latitude,
		Longitude:    u.Longitude,
		CompanyName:  u.CompanyName,
		Status:       u.Status,
		Probability:  u.Probability,
		LeadType:     u.LeadType,
		LeadChannel:  u.LeadChannel,
		LeadMedia:    u.LeadMedia,
		LeadSource:   u.LeadSource,
		GeneralNotes: u.GeneralNotes,
		User:         u.User,
	}
}
