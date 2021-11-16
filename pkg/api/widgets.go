package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type widget struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var widgets = []widget{
	{Id: 1, Name: "camera widget", Description: "a camera widget thing"},
	{Id: 2, Name: "other widget", Description: "other widget"},
	{Id: 3, Name: "magic widget", Description: "it does magic"},
	{Id: 4, Name: "beer me widget", Description: "it likes beer"},
}

func widgetHandler(w http.ResponseWriter, r *http.Request) {
	logger.WithFields(logrus.Fields{"api": "widgets"}).Infof("handling request")

	var result interface{} = nil
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		result = widgets
	} else {
		widgetId, _ := strconv.Atoi(id)

		for _, w := range widgets {
			if w.Id == widgetId {
				result = w
			}
		}
	}

	JSONResponse(w, r, result)
}
