// PROBLEMS HERE

package model

import (
	"errors"
)

// Should vessel id be here???
// 1:1 vessel is not necessary
// Business rule: Driver can not look for requests with VesselId = 0
type Driver struct {
	DriverId     int          `json:"driver_id" gorm:"primaryKey"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	VesselId     int          `json:"vessel_id"`
	Status       DriverStatus `json:"status"`
	Balance      float64      `json:"balance"`
	CertFirstAid int          `json:"cert_first_aid"`
	CertDriving  int          `json:"cert_driving"`
}

type DriverStatus string

const (
	UnknownDriverStatus DriverStatus = ""
	Busy                DriverStatus = "busy"
	Waiting             DriverStatus = "waiting"
	Afw                 DriverStatus = "afw"
)

func DriverStatusFromString(str string) (DriverStatus, error) {
	switch str {
	case "busy":
		return Busy, nil
	case "waiting":
		return Waiting, nil
	case "afw":
		return Afw, nil
	}

	return UnknownDriverStatus, errors.New("Unknown DriverStatus: " + str)
}

// IS BAD! WHAT TO DO WITH VesselId?
// Also add defaults to db?
func CreateDriver(id int, firstName string, lastName string, vesselId int, status DriverStatus, balance float64, certFA int, certD int) Driver {
	return Driver{id, firstName, lastName, vesselId, status, balance, certFA, certD}
}

func (d *Driver) SetFirstName(f string) {
	d.FirstName = f
}

func (d *Driver) SetLastName(l string) {
	d.LastName = l
}

func (d *Driver) SetVesselId(v int) {
	d.VesselId = v
}

func (d *Driver) SetStatus(s DriverStatus) {
	d.Status = s
}

func (d *Driver) SetBalance(b float64) {
	d.Balance = b
}

func (d *Driver) SetCertFirstAid(c int) {
	d.CertFirstAid = c
}

func (d *Driver) SetCertDriving(c int) {
	d.CertDriving = c
}
