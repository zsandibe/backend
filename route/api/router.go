package api

import (
	h "backend/internal/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.HomePage)
	mux.HandleFunc("/organization", h.OrganizationPage)
	mux.HandleFunc("/organization/register", h.RegisterPage)
	mux.HandleFunc("/projects", h.ProjectsPage)
	mux.HandleFunc("/projects/create", h.CreateProject)
	mux.Handle("/ui/css/", http.StripPrefix("/ui/css/", http.FileServer(http.Dir("./ui/css/"))))
	mux.HandleFunc("/projects/submit", h.SubmitProject)
	return mux
}
