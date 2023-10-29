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

// UpdateDock implements usecase.DockRepository.
func (repo *dockRepository) UpdateDock(d *model.Dock) error {
	dock := map[string]interface{}{
		"name":      d.Name,
		"latitude":  d.Latitude,
		"longitude": d.Longitude,
	}
	err := repo.db.Clauses(clause.Returning{}).Table(dockTableName).Where("dock_id = ?", d.DockId).Updates(&dock).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateDock implements usecase.DockRepository.
func (repo *dockRepository) CreateDock(dock *model.Dock) (int, error) {
	err := repo.db.Clauses(clause.Returning{}).Table(dockTableName).Select("name", "active", "latitude", "longitude").Create(&dock).Error
	if err != nil {
		return 0, err
	}
	return dock.DockId, nil
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

func (repo *dockRepository) GetNearestDocks(lat float64, lon float64, count int) ([]model.Dock, error) {
	var docks []model.Dock
	err := repo.db.Raw(
		"SELECT * FROM gog_demo.dock ORDER BY gog_demo.calculate_distance(latitude, longitude, ?, ?) LIMIT ?", lat, lon, count).
		Scan(&docks).Error
	if err != nil {
		return nil, err
	}
	return docks, nil
}
