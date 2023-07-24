package entity

import "time"

// Menu 菜单权限表
type Menu struct {
	MenuId     int64      `gorm:"column:menu_id;primaryKey;not null;autoIncrement;comment:菜单id" json:"menuId"`
	MenuName   string     `gorm:"column:menu_code;not null;size:50;comment:菜单名称" json:"menuName"`
	ParentId   int64      `gorm:"column:parent_id;default:0;comment:父菜单id" json:"parentId"`
	OrderNum   int        `gorm:"column:order_num;default:0;comment:显示顺序" json:"orderNum"`
	Path       string     `gorm:"size:200;default:'';comment:路由地址" json:"path"`
	Component  string     `gorm:"size:255;default:null;comment:组件路径" json:"component"`
	Query      string     `gorm:"size:255;default:null;comment:路由参数" json:"query"`
	IsFrame    bool       `gorm:"column:is_frame;type:boolean;default:false;comment:是否外链(true:是 false:不是)" json:"isFrame"`
	IsCache    bool       `gorm:"column:is_cache;type:boolean;default:true;comment:是否缓冲(true:是 false:不是)" json:"isCache"`
	MenuType   string     `gorm:"type:character;size:1;column:menu_type;default:'';comment:菜单类型(M目录 C菜单 F按钮)" json:"menuType"`
	Visible    bool       `gorm:"column:visible;type:boolean;default:true;comment:显隐状态(true显示 false隐藏)" json:"visible"`
	Status     int        `gorm:"column:status;default:1;comment:菜单状态(1正常 0停用)" json:"status"`
	Perms      string     `gorm:"size:100;default:null;comment:权限标识" json:"perms"`
	Icon       string     `gorm:"size:100;default:#;comment:菜单图标" json:"icon"`
	Remark     string     `gorm:"size:500;default:null;comment:备注" json:"remark"`
	CreateBy   string     `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime *time.Time `gorm:"column:create_time;default:null;comment:创建时间" json:"createTime"`
	UpdateBy   string     `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null;comment:更新时间" json:"updateTime"`
}
