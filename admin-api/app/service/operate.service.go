package service

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"time"
)

var Operate = NewOperateService()

type OperateService struct {
	pool      chan *entity.Operate
	list      []*entity.Operate
	batchSize int
	handler   func()
	timer     <-chan time.Time
}

func NewOperateService() *OperateService {
	result := &OperateService{pool: make(chan *entity.Operate, 10), batchSize: 3, timer: time.Tick(5 * time.Second)}
	result.handler = func() {
		for {
			select {
			case data := <-result.pool:
				result.list = append(result.list, data)
				if len(result.list) == result.batchSize {
					core.Log.Info("缓冲操作日志达到条数阈值，将写入数据库")
					if err := core.DB.Create(result.list).Error; err != nil {
						core.Log.Error("批量记录操作日志失败:%s", err.Error())
					}
					result.list = result.list[:0]
				}
			case <-result.timer:
				if len(result.list) > 0 {
					core.Log.Info("缓冲操作日志达到时间阈值，将写入数据库")
					if err := core.DB.Create(result.list).Error; err != nil {
						core.Log.Error("批量记录操作日志失败:%s", err.Error())
					}
					result.list = result.list[:0]
				}
			}
		}
	}
	go result.handler()
	return result
}

func (o *OperateService) Push(data *entity.Operate) {
	o.pool <- data
}
