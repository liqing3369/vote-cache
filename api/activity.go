package api

import (
	"context"
	"fmt"
	"gitee.com/lee0771/vote-cache/constant"
	"gitee.com/lee0771/vote-cache/global"
)

type ActivityApi struct {
}

//func (this ActivityApi) SaveActivity(activity model.ActActivity) {
//
//}

func (this ActivityApi) GetActivity(id int) map[string]string {
	values, _ := global.GVA_REDIS.HGetAll(context.Background(), this.generateKey(id)).Result()
	return values
}

func (this ActivityApi) IncrBy(id int, field string, value int64) {
	global.GVA_REDIS.HIncrBy(context.Background(), this.generateKey(id), field, value)
	this.expire(id)
}

func (this ActivityApi) expire(id int) {
	global.GVA_REDIS.Expire(context.Background(), this.generateKey(id), constant.EXPIRE_TIME)
}

func (this ActivityApi) generateKey(id int) string {
	return fmt.Sprintf("%s:%d", constant.ACTIVITY, id)
}
