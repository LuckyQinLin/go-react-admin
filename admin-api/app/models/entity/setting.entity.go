package entity

import "time"

// Setting 参数配置表
type Setting struct {
	ConfigId    int64  `gorm:"column:config_id;primaryKey;not null;autoIncrement;comment:参数主键" json:"configId"`
	ConfigName  string `gorm:"column:config_name;size:100;default:'';comment:参数名称" json:"configName"`
	ConfigKey   string `gorm:"column:config_key;size:100;default:'';comment:参数键名" json:"configKey"`
	ConfigValue string `gorm:"column:config_value;size:500;default:'';comment:参数键值" json:"configValue"`
	ConfigType  int    `gorm:"column:config_type;default:0;comment:系统内置(1是 0否)" json:"configType"`
	Remark      string `gorm:"size:500;default:null;comment:备注" json:"remark"`
	BaseField
}

// Visit 系统访问表
type Visit struct {
	VisitId       int64     `gorm:"column:visit_id;primaryKey;not null;autoIncrement;comment:访问ID" json:"visitId"`
	UserName      string    `gorm:"column:user_name;size:50;default:'';comment:用户账号" json:"userName"`
	IpAddr        string    `gorm:"column:ip_addr;size:128;default:'';comment:登录IP地址" json:"ipAddr"`
	LoginLocation string    `gorm:"column:login_location;size:255;default:'';comment:登录地点" json:"loginLocation"`
	Browser       string    `gorm:"column:browser;size:50;default:'';comment:浏览器类型" json:"browser"`
	Os            string    `gorm:"column:os;size:50;default:'';comment:操作系统" json:"os"`
	Msg           string    `gorm:"column:msg;size:255;default:'';comment:提示消息" json:"msg"`
	LoginTime     time.Time `gorm:"column:login_time;comment:访问时间" json:"loginTime"`
	Status        int       `gorm:"column:status;default:1;comment:登录状态(1成功 0失败)" json:"status"`
}

// Notice 系统通知表
type Notice struct {
	NoticeId      int64  `gorm:"column:notice_id;primaryKey;not null;autoIncrement;comment:公告ID" json:"noticeId"`
	NoticeTitle   string `gorm:"column:notice_title;size:50;not null;comment:公告标题" json:"noticeTitle"`
	NoticeType    int    `gorm:"column:notice_type;type:char;size:1;not null;comment:公告类型(1通知 2公告)" json:"noticeType"`
	NoticeContent string `gorm:"column:notice_content;type:text;default:null;comment:公告内容" json:"noticeContent"`
	Status        int    `gorm:"column:status;default:1;comment:公告状态(1正常 0停用)" json:"status"`
	Remark        string `gorm:"size:500;default:null;comment:备注" json:"remark"`
	BaseField
}
