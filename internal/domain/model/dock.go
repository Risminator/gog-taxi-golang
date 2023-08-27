package model

type Dock struct {
	DockId    int     `json:"dockId" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Active    bool    `json:"active"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func CreateDock(id int, name string, active bool, lat float64, lon float64) Dock {
	return Dock{id, name, active, lat, lon}
}

func (d *Dock) SetId(id int) {
	d.DockId = id
}

func (d *Dock) SetName(n string) {
	d.Name = n
}

func (d *Dock) SetActive(a bool) {
	d.Active = a
}

func (d *Dock) SetLatitude(l float64) {
	d.Latitude = l
}

func (d *Dock) SetLongitude(l float64) {
	d.Longitude = l
}
