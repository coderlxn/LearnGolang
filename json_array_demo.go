package main

import (
	"encoding/json"
	"log"
)

var body_byte = []byte(`{
 "status": true,
 "data": {
 "ModelList": [{
   "Id": 1,
   "Name": "foo",
   "CarId": 1,
   "EngName": "bar"
  }]
 }
}`)

type AjaxModelsList struct {
	Id      float64 `json:"Id"`
	Name    string  `json:"Name"`
	CarId   float64 `json:"CarId"`
	EngName string  `json:"EngName"`
}

type AjaxModels struct {
	Status bool                      `json:"status"`
	Data   map[string][]AjaxModelsList `json:"data"`
}

func main() {
	c := AjaxModels{}
	if err := json.Unmarshal(body_byte, &c); err != nil {
		log.Fatalf("an error occured: %v", err)
	}
	log.Println(c, c.Data["ModelList"][0].Id)
}
