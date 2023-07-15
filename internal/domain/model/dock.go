// PROBLEMS HERE

package model

// TODO: Rename "working" to "active" in db
// and maybe latitude and longitude???? (lat and lon there)
// in DB it's like this: dock_id name location lat lon working
type Dock struct {
	DockId    int     `json:"dock_id"`
	Name      string  `json:"name"`
	Active    bool    `json:"active"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// location Point ??? entity shouldn't know about this
}

func (d *Dock) CreateDock(id int, name string, active bool, lat float64, lon float64) Dock {
	return Dock{id, name, active, lat, lon}
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
