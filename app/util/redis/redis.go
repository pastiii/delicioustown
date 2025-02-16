package redis

import (
	"DeliciousTown/global"
	"context"
	"time"
)

func GetValue(key string) any {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	val, err := global.GvaRedis.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	return val
}

func ForgetValue(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	err := global.GvaRedis.Del(ctx, key).Err()
	if err != nil {
		return false
	}

	return true
}


