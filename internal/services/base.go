package services

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/zbd20/fiber-demo/internal/models"
)

var ServiceSet = wire.NewSet(NewBaseService)

type BaseService struct {
	DemoService *demoService
}

func NewBaseService(db *gorm.DB) (*BaseService, error) {
	if err := models.AutoMigrate(db); err != nil {
		return nil, err
	}

	return &BaseService{
		DemoService: newDemoService(db),
	}, nil
}
