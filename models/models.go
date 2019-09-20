package models

import (
    u "github.com/GulnazBagautdinova/go-cars/utils"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Brand string `json:"brand"`
	Vehicle string `json:"vehicle"`
	Price uint `json:"price"` 
    Status string `json:"status"`
	Mileage uint `json:"mileage"`
}

var status = []string{ "В пути", "На складе", "Продан", "Снят с продажи" }

func (model *Model) Validate() (map[string] interface{}, bool) {
     var i int
	for _, s := range status {
        if (model.Status == s){
			i = 1    
		}
	}
	if ( i==0 ) {
	return u.Message(false, "Status not correct"), false
	}
	//There are all  required parameters 
	return u.Message(true, "success"), true
}

func (model *Model) Create() (map[string] interface{}) {

	if resp, ok := model.Validate(); 
	!ok {
		return resp
	}
	GetDB().Create(model)
	resp := u.Message(true, "success")
	resp["model"] = model
	if(model.Mileage > 0){
		resp["is new"] = false
	} else {
		resp["is new"] = true 
	}
	return resp
}

func (model *Model) Update(id string) (map[string] interface{}) {
	 
	if resp, ok := model.Validate(); 
	!ok {
		return resp
	}
	res := GetDB().Table("models").Where("id = ?", id).Update(&model)
	if res == nil {
		resp := u.Message(false, "Model not found")
		return resp
	}
	resp := u.Message(true, "success")
	//resp["model"] = model
	return resp
}

func (model *Model) Delete(id string) (map[string] interface{}) {
	 
	res := GetDB().Table("models").Where("id = ?", id).First(&model)
	if res == nil || model.ID == 0  {
		resp := u.Message(false, "Model not found or was deleted")
		return resp
	}
	GetDB().Table("models").Where("id= ?", id).Delete(res)
	resp := u.Message(true, "success")
	return resp
}

func (model *Model) Get(id string) (map[string] interface{}) {
	 
	res := GetDB().Table("models").Where("id = ?", id).First(&model)
	if res == nil || model.ID == 0  {
		resp := u.Message(false, "Model not found or was deleted")
		return resp
	}
	resp := u.Message(true, "success")
	resp["model"] = model
	if(model.Mileage > 0){
		resp["is new"] = false
	} else {
		resp["is new"] = true 
	}
	return resp
}

