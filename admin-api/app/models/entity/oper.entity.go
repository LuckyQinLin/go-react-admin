package entity

import "time"

// Operate 操作日志记录
type Operate struct {
	OperId        int64      `gorm:"column:oper_id;primaryKey;not null;autoIncrement;comment:日志id" json:"operId"`
	Title         string     `gorm:"size:50;default:'';comment:模块标题" json:"title"`
	BusinessType  int        `gorm:"column:business_type;default:0;comment:业务类型(0其它 1新增 2修改 3删除)" json:"businessType"`
	Method        string     `gorm:"size:100;default:'';comment:方法名称" json:"method"`
	RequestMethod string     `gorm:"size:10;default:'';comment:请求方式" json:"requestMethod"`
	OperatorType  int        `gorm:"column:operator_type;default:0;comment:操作类别(0其它 1后台用户 2手机端用户)" json:"operatorType"`
	OperName      string     `gorm:"column:oper_name;size:50;default:'';comment:操作人员" json:"operName"`
	DeptName      string     `gorm:"column:dept_name;size:50;default:'';comment:部门名称" json:"deptName"`
	OperUrl       string     `gorm:"column:oper_url;size:255;default:'';comment:请求URL" json:"operUrl"`
	OperIp        string     `gorm:"column:oper_ip;size:128;default:'';comment:主机地址" json:"operIp"`
	OperLocation  string     `gorm:"column:oper_location;size:255;default:'';comment:操作地点" json:"operLocation"`
	OperParam     string     `gorm:"column:oper_param;size:2000;default:'';comment:请求参数" json:"operParam"`
	JsonResult    string     `gorm:"column:json_result;size:2000;default:'';comment:返回参数" json:"jsonResult"`
	ErrorMsg      string     `gorm:"column:error_msg;size:200;default:'';comment:错误消息" json:"errorMsg"`
	OperTime      *time.Time `gorm:"column:oper_time;default:null;comment:操作时间" json:"operTime"`
	CostTime      int64      `gorm:"column:cost_time;size:50;default:0;comment:消耗时间" json:"costTime"`
	Status        int        `gorm:"column:status;default:1;comment:操作状态(1正常 0异常)" json:"status"`
}
