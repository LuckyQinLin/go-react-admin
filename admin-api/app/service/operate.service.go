package service

import (
	"admin-api/app/models/entity"
	"admin-api/core"
)

var Operate = NewOperateService()

type OperateService struct {
	pool      chan *entity.Operate
	list      []*entity.Operate
	batchSize int
	handler   func()
}

func NewOperateService() *OperateService {
	result := &OperateService{pool: make(chan *entity.Operate, 50), batchSize: 10}
	result.handler = func() {
		for {
			select {
			case data := <-result.pool:
				result.list = append(result.list, data)
				if len(result.list) == result.batchSize {
					if err := core.DB.Create(result.list).Error; err != nil {
						core.Log.Error("批量记录操作日志失败:%s", err.Error())
					}
					result.list = result.list[:0]
				}
			}
		}
	}
	return result
}

func (o *OperateService) Push(data *entity.Operate) {
	o.pool <- data
}
