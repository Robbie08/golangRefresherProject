package objects

import "fmt"

func (v *Vehicle) Init(maker, model, year string, axels int) {
	v.Maker = maker
	v.Model = model
	v.Year = year
	v.Axels = axels
}

func (v *Vehicle) VehicleInfo() {
	fmt.Printf("================VEHICLE INFORMATION================\n")
	fmt.Printf("Make: %s\nModel: %s\nYear: %s\nAxels: %d\n", v.Maker, v.Model, v.Year, v.Axels)
}
