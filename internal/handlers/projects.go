package handlers

import (
	"html/template"
	"net/http"
)

func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/projects.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Здесь можно получить список проектов из базы данных и передать их в шаблон

		// if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		// 	ErrorPage(w, "Ошибка сервера", http.StatusInternalServerError)
		// 	return
		// }
		tmpl.Execute(w, nil)

	default:
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects/create" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/create_project.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		// 	ErrorPage(w, "Ошибка сервера", http.StatusInternalServerError)
		// 	return
		// }
		tmpl.Execute(w, nil)

	case "POST":
		// Обработка отправленной формы

	default:
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func SubmitProject(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects/submit" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Обработка отправленной формы

	default:
		w.Header().Set("Allow", http.MethodPost)
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
