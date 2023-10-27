package usecase

import "github.com/Risminator/gog-taxi-golang/internal/domain/model"

type Dock interface {
	GetDockById(id int) (*model.Dock, error)
	GetDocks() ([]model.Dock, error)
	GetNearestDocks(lat float64, lon float64, count int) ([]model.Dock, error)

	CreateDock(name string, latitude float64, longitude float64) (*model.Dock, error)
	UpdateDock(dockId int, name string, latitude, longitude float64) (*model.Dock, error)
}
type DockRepository interface {
	GetDockById(id int) (*model.Dock, error)
	GetDocks() ([]model.Dock, error)
	GetNearestDocks(lat float64, lon float64, count int) ([]model.Dock, error)

	CreateDock(*model.Dock) (int, error)
	UpdateDock(*model.Dock) error
}

type dockUsecase struct {
	dockRepo DockRepository
}

func NewDockUsecase(d DockRepository) Dock {
	return &dockUsecase{
		dockRepo: d,
	}
}

// UpdateDock implements Dock.
func (use *dockUsecase) UpdateDock(dockId int, name string, latitude float64, longitude float64) (*model.Dock, error) {
	dock, err := use.GetDockById(dockId)
	if err != nil {
		return nil, err
	}
	dock.SetName(name)
	dock.SetLatitude(latitude)
	dock.SetLongitude(longitude)

	err = use.dockRepo.UpdateDock(dock)
	if err != nil {
		return nil, err
	}
	return dock, nil
}

// CreateDock implements Dock.
func (use *dockUsecase) CreateDock(name string, latitude float64, longitude float64) (*model.Dock, error) {
	dock := model.CreateDock(0, name, true, latitude, longitude)
	id, err := use.dockRepo.CreateDock(&dock)
	if err != nil {
		return nil, err
	}

	dock.SetDockId(id)
	return &dock, nil
}

// GetDockById implements Dock.
func (use *dockUsecase) GetDockById(id int) (*model.Dock, error) {
	dock, err := use.dockRepo.GetDockById(id)
	if err != nil {
		return nil, err
	}
	return dock, nil
}

// GetDocks implements Dock.
func (use *dockUsecase) GetDocks() ([]model.Dock, error) {
	docks, err := use.dockRepo.GetDocks()
	if err != nil {
		return nil, err
	}
	return docks, nil
}

func (use *dockUsecase) GetNearestDocks(lat float64, lon float64, count int) ([]model.Dock, error) {
	docks, err := use.dockRepo.GetNearestDocks(lat, lon, count)
	if err != nil {
		return nil, err
	}
	return docks, nil
}
