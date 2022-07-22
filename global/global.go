package global

import (
	"github.com/go-redis/redis/v8"
	oplogging "github.com/op/go-logging"
	"vote-cache/config"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *oplogging.Logger
	GVA_REDIS  *redis.Client
)
