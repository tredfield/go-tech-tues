package api

import (
	"net/http"
	"text/template"
	"time"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/layout.html", "static/templates/welcome.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("static/templates/welcome.html" + err.Error()))
		return
	}

	data := struct {
		Name string
		Time string
	}{
		Name: "slalom",
		Time: time.Now().Format(time.Stamp),
	}

	if name := r.FormValue("name"); name != "" {
		data.Name = name
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "static/templates/welcome.html"+err.Error(), http.StatusInternalServerError)
	}
}
