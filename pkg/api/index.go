package api

import (
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/layout.html", "static/templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("static/templates/index.html" + err.Error()))
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", ""); err != nil {
		logger.Errorf("error executing templates: %s", err.Error())
		http.Error(w, "static/templates/welcome.html"+err.Error(), http.StatusInternalServerError)
	}
}
