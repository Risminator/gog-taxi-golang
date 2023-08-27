package model

type Vessel struct {
	VesselId   int     `json:"vesselId" gorm:"primaryKey"`
	Model      string  `json:"model"`
	Seats      int     `json:"seats"`
	IsApproved bool    `json:"isApproved"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func (v *Vessel) CreateVessel(vesselId int, model string, seats int, isAppr bool, lat float64, lon float64) Vessel {
	return Vessel{vesselId, model, seats, isAppr, lat, lon}
}

func (v *Vessel) SetVesselId(id int) {
	v.VesselId = id
}

func (v *Vessel) SetModel(m string) {
	v.Model = m
}

func (v *Vessel) SetSeats(s int) {
	v.Seats = s
}

func (v *Vessel) SetIsApproved(a bool) {
	v.IsApproved = a
}

func (v *Vessel) SetLatitude(l float64) {
	v.Latitude = l
}

func (v *Vessel) SetLongitude(l float64) {
	v.Longitude = l
}
