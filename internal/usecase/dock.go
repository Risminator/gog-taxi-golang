package usecase

import "github.com/Risminator/gog-taxi-golang/internal/domain/model"

type Dock interface {
	GetDockById(id int) (*model.Dock, error)
	GetDocks() ([]model.Dock, error)

	CreateDock(name string, latitude float64, longitude float64) (*model.Dock, error)
}
type DockRepository interface {
	GetDockById(id int) (*model.Dock, error)
	GetDocks() ([]model.Dock, error)

	CreateDock(*model.Dock) error
}

type dockUsecase struct {
	dockRepo DockRepository
}

func NewDockUsecase(d DockRepository) Dock {
	return &dockUsecase{
		dockRepo: d,
	}
}

// CreateDock implements Dock.
func (use *dockUsecase) CreateDock(name string, latitude float64, longitude float64) (*model.Dock, error) {
	dock := model.CreateDock(0, name, true, latitude, longitude)
	err := use.dockRepo.CreateDock(&dock)
	if err != nil {
		return nil, err
	}
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
