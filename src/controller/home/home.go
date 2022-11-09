package controllerHome

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hikaaam/go-learn/src/helper"
	"github.com/hikaaam/go-learn/src/model"
)

func GetGorillas(w http.ResponseWriter, r *http.Request) {
	var data []model.Gorilla         //declare arrays
	model.DB().Find(&data)           //find data and assign to array, (select * from table)
	jsonData := helper.SToJson(data) //convert data to JSON
	//send data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShowGorilla(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)            //get params map ["keys":"value"]
	id := params["id"]               //get params["keys"] and assign it to id
	var data model.Gorilla           //declare data
	model.DB().Find(&data, id)       //get data by id and assign to data, {select * from table where id = "id" }
	jsonData := helper.SToJson(data) //convert data to json
	//send data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func AddGorrila(w http.ResponseWriter, r *http.Request) {
	var data model.Gorilla                //declare data
	json.NewDecoder(r.Body).Decode(&data) //get request body and assign to data
	model.DB().Create(&data)              //create data in the database
	jsonData := helper.SToJson(data)      //convert data to json
	//send data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func EditGorilla(w http.ResponseWriter, r *http.Request) {
	var data model.Gorilla                //declare data
	json.NewDecoder(r.Body).Decode(&data) //get request body and assign to data
	model.DB().Save(&data)                //update data using body request
	jsonData := helper.SToJson(data)      //convert data to json
	//send data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func DeleteGorilla(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)            //get params map ["keys":"value"]
	id := params["id"]               //get params["keys"] and assign it to id
	var data model.Gorilla           //declare data
	model.DB().Delete(&data, id)     //delete table where id = "id"
	jsonData := helper.SToJson(data) //convert data to json
	//send data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
