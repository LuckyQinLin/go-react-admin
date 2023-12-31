package core

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml/v2"
	"os"
	"strings"
)

var Config *Configuration

// Configuration 配置表
type Configuration struct {
	Web    Server    `toml:"web" json:"web"`     // web服务器配置
	Db     Database  `toml:"db" json:"db"`       // 数据库配置
	Redis  Redis     `toml:"redis" json:"redis"` // Redis缓冲
	Jwt    JwtConfig `toml:"jwt" json:"jwt"`     // jwt配置
	Logger Logger    `toml:"log" json:"log"`     // 日志配置
}

// Server 服务器相关配置
type Server struct {
	Port           int64    `toml:"port" json:"port"`                         // 端口
	ContextPath    string   `toml:"context-path" json:"context-path"`         // 路径
	ReadTimeout    int      `toml:"read-timeout" json:"read-timeout"`         // 读取时间
	WriteTimeout   int      `toml:"write-timeout" json:"write-timeout"`       // 写时间
	MaxHeaderBytes int      `toml:"max-header-bytes" json:"max-header-bytes"` // 最大头
	RunModel       string   `toml:"dev-model" json:"run-model"`               // 运行模型
	WhiteList      []string `toml:"white-list" json:"white-list"`             // 白名单
}

func (s *Server) Whites() []string {
	result := make([]string, 0)
	for _, item := range s.WhiteList {
		if s.ContextPath == "" {
			result = append(result, item)
		} else {
			if strings.HasPrefix(s.ContextPath, "/") {
				result = append(result, s.ContextPath+item)
			} else {
				result = append(result, "/"+s.ContextPath+item)
			}
		}
	}
	return result
}

// Database 数据库配置
type Database struct {
	Host     string `toml:"host" json:"host"`         // 地址
	Port     int64  `toml:"port" json:"port"`         // 端口
	Username string `toml:"username" json:"username"` // 用户
	Password string `toml:"password" json:"password"` // 密码
	DbName   string `toml:"db-name" json:"dbName"`    // 数据库
	Schema   string `toml:"schema" json:"schema"`     // 模式
}

// Redis redis配置
type Redis struct {
	Host     string `toml:"host" json:"host"`         // 地址
	Port     int64  `toml:"port" json:"port"`         // 端口
	Password string `toml:"password" json:"password"` // 密码
	Db       int    `toml:"db" json:"db"`             // 数据库
}

func (d Database) Link() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=disable TimeZone=Asia/Shanghai",
		d.Host,
		d.Username,
		d.Password,
		d.DbName,
		d.Port,
		d.Schema)
}

// JwtConfig jwt配置
type JwtConfig struct {
	Issuer      string `toml:"issuer" json:"issuer"`             // 作者
	ExpiresTime int64  `toml:"expires-time" json:"expires-time"` // 超时时间
	SecretKey   string `toml:"secret-key" json:"secret-key"`     // 秘钥
}

// Logger 日志配置
type Logger struct {
	Prefix string `toml:"prefix" json:"prefix"` // 前缀
}

// InitConfig 初始化配置文件
func InitConfig() {
	var (
		config Configuration
		path   string
		bytes  []byte
		err    error
	)
	path, _ = homedir.Expand("~/.admin/config.toml")
	if bytes, err = os.ReadFile(path); err != nil {
		panic("读取配置文件失败 => " + err.Error())
	}
	if err = toml.Unmarshal(bytes, &config); err != nil {
		panic("解析配置文件失败 => " + err.Error())
	}
	Config = &config
}
