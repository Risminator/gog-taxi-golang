package repository

import (
	// "log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const dockTableName = "gog_demo.dock"

type dockRepository struct {
	db *gorm.DB
}

func NewDockRepository(db *gorm.DB) usecase.DockRepository {
	return &dockRepository{db}
}

// CreateDock implements usecase.DockRepository.
func (repo *dockRepository) CreateDock(dock *model.Dock) error {
	err := repo.db.Clauses(clause.Returning{}).Table(dockTableName).Select("name", "latitude", "longitude").Create(&dock).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDockById implements usecase.DockRepository.
func (repo *dockRepository) GetDockById(id int) (*model.Dock, error) {
	var dock model.Dock
	err := repo.db.Table(dockTableName).First(&dock, "dock_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &dock, nil
}

// GetDocks implements usecase.DockRepository.
func (repo *dockRepository) GetDocks() ([]model.Dock, error) {
	var docks []model.Dock
	err := repo.db.Table(dockTableName).Find(&docks).Error
	if err != nil {
		return nil, err
	}
	return docks, nil
}
