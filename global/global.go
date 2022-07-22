package global

import (
	"gitee.com/lee0771/vote-cache/config"
	"github.com/go-redis/redis/v8"
	oplogging "github.com/op/go-logging"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *oplogging.Logger
	GVA_REDIS  *redis.Client
)
