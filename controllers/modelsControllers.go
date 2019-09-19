package controllers

import (
	"net/http"
	"go-cars/models"
	"github.com/gorilla/mux"
	"encoding/json"
	u "go-cars/utils"
)

var CreateModel = func(w http.ResponseWriter, r *http.Request) {

	model := &models.Model{}
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}	
	resp := model.Create()
	if(model.Mileage > 0){
		resp["is new"] = false
	} else {
		resp["is new"] = true 
	}
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

   resp := model.GetModels(params["id"])
  
   u.Respond(w, resp)
}
