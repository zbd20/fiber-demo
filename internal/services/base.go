package services

import (
	"gorm.io/gorm"
)

type BaseService struct {
	DemoService *demoService
}

func NewBaseService(db *gorm.DB) *BaseService {
	return &BaseService{
		DemoService: newDemoService(db),
	}

}
