package main

import (
	"testing"

	"github.com/Robbie08/goOOP/objects"
)

func TestAddVehicle(t *testing.T) {

	g := new(objects.Garage)
	g.Init("Robert", "San Diego", "CA")

	g.AddVehicleToGarage(&objects.Vehicle{})
	g.AddVehicleToGarage(&objects.Vehicle{})

	if v := g.GetSize(); v != 2 {
		t.Errorf("Expected Size to be '2', but got '%d'", v)
	}
}

func TestDeleteVehicle(t *testing.T) {
	g := new(objects.Garage)
	g.Init("Robert", "San Diego", "CA")

	// add vehicles
	g.AddVehicleToGarage(&objects.Vehicle{})
	g.AddVehicleToGarage(&objects.Vehicle{})
	g.AddVehicleToGarage(&objects.Vehicle{})

	g.PopVehicleFromGarage()
	g.PopVehicleFromGarage()

	if v := g.GetSize(); v != 1 {
		t.Errorf("Expected Size to be '1', but got '%d'", v)
	}
}
