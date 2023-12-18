package routes

import (
	"leadsmanagementsystem/handlers"
	"leadsmanagementsystem/pkg/mysql"
	"leadsmanagementsystem/repositories"

	"github.com/gorilla/mux"
)

func leadRoutes(r *mux.Router) {
	leadRepository := repositories.RepositoryLead(mysql.DB)
	h := handlers.HandlerLead(leadRepository)

	r.HandleFunc("/leads", h.FindLeads).Methods("GET")
	r.HandleFunc("/lead/{id}", h.GetLead).Methods("GET")
	r.HandleFunc("/lead", h.CreateLead).Methods("POST")
	r.HandleFunc("/lead/{id}", h.UpdateLead).Methods("PATCH")
	r.HandleFunc("/lead/{id}", h.DeleteLead).Methods("DELETE")
	r.HandleFunc("/search", h.SearchLeads).Methods("GET")

}
