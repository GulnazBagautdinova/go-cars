package controllers

import (
	"net/http"
	"github.com/GulnazBagautdinova/go-cars/models"
	"github.com/gorilla/mux"
	"encoding/json"
	u "github.com/GulnazBagautdinova/go-cars/utils"
)

var CreateModel = func(w http.ResponseWriter, r *http.Request) {

	model := &models.Model{}
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}	
	resp := model.Create()
	u.Respond(w, resp)
}

var UpdateModel = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	model := &models.Model{}
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}	
	
	resp := model.Update(params["id"])
    u.Respond(w, resp)
}

var DeleteModel = func (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	model := &models.Model{}

    resp := model.Delete(params["id"])
    u.Respond(w, resp)
}

var GetModel = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	model := &models.Model{}

	resp := model.Get(params["id"])
    u.Respond(w, resp)
}

