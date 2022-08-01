package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Robbie08/goOOP/objects"
)

var g *objects.Garage

func main() {
	fmt.Printf("===================Starting Vehicle Server====================")

	g = new(objects.Garage)
	g.Init("Robert", "San Diego", "California")

	car1 := new(objects.Vehicle)
	car1.Init("Toyota", "4Runner", "2002", 2)

	car2 := new(objects.Vehicle)
	car2.Init("Mazda", "3", "2014", 2)

	car3 := new(objects.Vehicle)
	car3.Init("Nissan", "R34 Skyline", "2003", 2)

	g.AddVehicleToGarage(car1)
	g.AddVehicleToGarage(car2)
	g.AddVehicleToGarage(car3)

	http.HandleFunc("/shutdown", shutDown)
	http.HandleFunc("/api/addvehicle", addVehicle)
	http.HandleFunc("/api/popvehicle", popVehicle)
	http.HandleFunc("/api/printgarage", printGarage)
	http.ListenAndServe(":8080", nil)

}

func shutDown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func addVehicle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("You Hit the add vehicle page. Still in progress...")
	decoder := json.NewDecoder(req.Body)
	var v *objects.Vehicle
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}

	g.AddVehicleToGarage(v)
}

func popVehicle(res http.ResponseWriter, req *http.Request) {
	resp, err := g.PopVehicleFromGarage()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Removed: %v\n", resp)
}

func printGarage(res http.ResponseWriter, req *http.Request) {
	if g.GetSize() <= 0 {
		fmt.Println("List is empty, no items to print...")
		return
	}
	fmt.Printf("Size: %d\n", g.GetSize())
	g.PrintGarageVehicles()
}
