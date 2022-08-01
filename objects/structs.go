package objects

type Vehicle struct {
	Maker string
	Model string
	Year  string
	Axels int
}

type Garage struct {
	Owner     string
	City      string
	State     string
	VehicleDb []*Vehicle
	Size      int
}
