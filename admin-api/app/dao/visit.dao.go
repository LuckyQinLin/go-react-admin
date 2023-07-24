package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
)

var Visit = new(VisitDao)

type VisitDao struct{}

func (v *VisitDao) Save(visit *entity.Visit) error {
	return core.DB.Create(visit).Error
}
