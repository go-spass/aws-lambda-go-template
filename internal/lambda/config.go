package lambda

import "log"

type Config struct {
	Log       log.Logger
	TableName string
	IsLocal   bool
}
