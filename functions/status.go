package functions

import (
	"assignment_3/structs"

	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CreateJson() {
	for {
		data := structs.DisasterIndex{}
		data.Status.Water = int(RandomNumber(0, 100))
		data.Status.Wind = int(RandomNumber(0, 100))

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal("Err: ", err.Error())
		}
		err = ioutil.WriteFile("./data.json", jsonData, 0644)

		if err != nil {
			log.Fatal("Err: ", err.Error())
		}
		time.Sleep(time.Second * 15)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Err: ", err.Error())
	}

	var status structs.DisasterIndex

	err = json.Unmarshal(data, &status)
	if err != nil {
		log.Fatal("Err: ", err.Error())
	}

	water := status.Status.Water
	wind := status.Status.Wind
	var statusWater string
	var statusWind string

	switch {
	case water < 5:
		statusWater = "Aman"
	case water >= 6 && status.Status.Water <= 8:
		statusWater = "Siaga"
	case water > 8:
		statusWater = "Bahaya"
	default:
		statusWater = "Not defined"
	}

	switch {
	case wind < 6:
		statusWind = "Aman"
	case wind >= 7 && status.Status.Wind <= 15:
		statusWind = "Siaga"
	case wind > 15:
		statusWind = "Bahaya"
	default:
		statusWind = "Not defined"
	}

	dataStatus := map[string]interface{}{
		"waterValue":  water,
		"waterStatus": statusWater,
		"windValue":   wind,
		"windStatus":  statusWind,
	}
	fmt.Println("Water :", dataStatus["waterValue"])
	fmt.Println("Status :", dataStatus["waterStatus"])
	fmt.Println("Wind :", dataStatus["windValue"])
	fmt.Println("Status :", dataStatus["windStatus"])

	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal("Err :", err.Error())
	}
	template.Execute(w, dataStatus)
}
