package logger

type LogFile struct {
	log      chan []byte
	basePath string
}
