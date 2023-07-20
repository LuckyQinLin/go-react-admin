package dao

import "admin-api/app/models/entity"

var Visit = new(VisitDao)

type VisitDao struct{}

func (v *VisitDao) Save(visit *entity.Visit) error {
	return nil
}
