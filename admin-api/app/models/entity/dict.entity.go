package entity

import "time"

// DictData 字典数据表
type DictData struct {
	DataId     int64      `gorm:"column:data_id;primaryKey;not null;autoIncrement;comment:字典编码" json:"dataId"`
	DataSort   int        `gorm:"column:data_sort;default:0;comment:字典排序" json:"dataSort"`
	DataLabel  string     `gorm:"column:data_label;size:100;default:'';comment:字典标签" json:"dataLabel"`
	DataValue  string     `gorm:"column:data_value;size:100;default:'';comment:字典键值" json:"dataValue"`
	DataType   string     `gorm:"column:data_type;size:100;default:'';comment:字典类型" json:"dataType"`
	CssClass   string     `gorm:"column:css_class;size:100;default:null;comment:样式属性（其他样式扩展）" json:"cssClass"`
	ListClass  string     `gorm:"column:list_class;size:100;default:null;comment:表格回显样式" json:"listClass"`
	IsDefault  bool       `gorm:"column:is_default;default:false;comment:是否默认(true是 false否)" son:"isDefault"`
	Status     int        `gorm:"column:status;default:1;comment:状态(1正常 0停用)" json:"status"`
	Remark     string     `gorm:"size:500;default:null;comment:备注" json:"remark"`
	CreateBy   string     `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime *time.Time `gorm:"column:create_time;default:null;comment:创建时间" json:"createTime"`
	UpdateBy   string     `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null;comment:更新时间" json:"updateTime"`
}

// DictType 字典类型表
type DictType struct {
	DictId     int64      `gorm:"column:dict_id;primaryKey;not null;autoIncrement;comment:字典主键" json:"dictId"`
	DictName   string     `gorm:"column:dict_name;size:100;default:'';comment:字典名称" json:"dictName"`
	DictType   string     `gorm:"column:dict_type;size:100;default:'';unique;comment:字典类型" json:"dictType"`
	Status     int        `gorm:"column:status;default:1;comment:状态(1正常 0停用)" json:"status"`
	Remark     string     `gorm:"size:500;default:null;comment:备注" json:"remark"`
	CreateBy   string     `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime *time.Time `gorm:"column:create_time;default:null;comment:创建时间" json:"createTime"`
	UpdateBy   string     `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null;comment:更新时间" json:"updateTime"`
}
