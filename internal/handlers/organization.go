package handlers

import (
	"backend/internal/db"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func OrganizationPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organization" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/organizations.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	organizations, err := db.GetOrganizations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, organizations)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organization/register" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		// tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/register.html")
		// if err != nil {
		// 	ErrorPage(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		tmpl, err := template.ParseFiles("./ui/templates/register.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bin := r.FormValue("bin")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Println(bin, name, address, r.FormValue("number"), "--")
		workers, err := strconv.Atoi(r.FormValue("number"))
		if err != nil {
			fmt.Println("error atoi")
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = db.AddOrganization(bin, name, address, workers)
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/organization", http.StatusFound)

	default:
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
