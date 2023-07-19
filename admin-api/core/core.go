package core

func InitCore() {
	InitConfig()
	InitLogger()
	InitDb()
	InitRedis()
}
