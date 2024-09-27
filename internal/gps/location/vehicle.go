package location

type Vehicle struct {
	identifier string
	Plate      string  `json:"plate"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
