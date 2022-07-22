package api

import (
	"context"
	"fmt"
	"vote-cache/constant"
	"vote-cache/global"
)

type PlayerApi struct {
}

func (this PlayerApi) GetPlayer(id int) map[string]string {
	values, _ := global.GVA_REDIS.HGetAll(context.Background(), this.generateKey(id)).Result()
	return values
}

func (this PlayerApi) IncrBy(id int, field string, value int64) {
	global.GVA_REDIS.HIncrBy(context.Background(), this.generateKey(id), field, value)
	this.expire(id)
}

func (this PlayerApi) expire(id int) {
	global.GVA_REDIS.Expire(context.Background(), this.generateKey(id), constant.EXPIRE_TIME)
}

func (this PlayerApi) generateKey(id int) string {
	return fmt.Sprintf("%s:%d", constant.PLAYER, id)
}
