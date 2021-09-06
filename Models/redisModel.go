package Models

import (
	"context"
	"errors"
	"time"

	"github.com/bigby/project/Database"
	"github.com/bigby/project/Utils"
)

func PutJSON(context context.Context, val map[string]interface{}, t time.Duration) (resKey string, err error) {
	key, err := Utils.NewKey(16)

	if err != nil {
		return "", err
	}

	resKey = string(key[:])
	err = Database.Redisdb.HMSet(context, resKey, val).Err()

	if err != nil {
		return
	}

	_, err = Database.Redisdb.Expire(context, resKey, t).Result()

	return
}

func GetJSON(context context.Context, key string) (val map[string]string, err error) {
	exist, err := Database.Redisdb.Exists(context, key).Result()
	if err != nil {
		return
	}

	if exist == 0 {
		err = errors.New("用户数据已过期，请重新登录")
		return
	}

	val, err = Database.Redisdb.HGetAll(context, key).Result()

	return
}
