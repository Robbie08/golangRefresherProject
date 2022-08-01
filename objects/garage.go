package objects

import "fmt"

var size int

func (g *Garage) Init(owner, city, state string) {
	g.Owner = owner
	g.City = city
	g.State = state
	g.Size = 0
}

func (g *Garage) GetSize() int {
	return g.Size
}

func (g *Garage) AddVehicleToGarage(v *Vehicle) bool {
	g.VehicleDb = append(g.VehicleDb, v)
	g.Size += 1
	return true
}

func (g *Garage) PopVehicleFromGarage() (Vehicle, error) {
	if g.GetSize() <= 0 {
		return Vehicle{}, fmt.Errorf("Cannot Pop any vehicles from garage because it's empty...")
	}

	// pop value from stack
	var popedVehicle *Vehicle
	popedVehicle, g.VehicleDb = g.VehicleDb[len(g.VehicleDb)-1], g.VehicleDb[:len(g.VehicleDb)-1]
	g.Size--
	return *popedVehicle, nil
}

func (g *Garage) PrintGarageVehicles() {
	num := 1
	for _, v := range g.VehicleDb {
		fmt.Printf("Vehicle %d: { make:%v\tmodel:%v\tyear:%v\taxels:%d }\n", num, v.Maker, v.Model, v.Year, v.Axels)
		num++
	}
}
