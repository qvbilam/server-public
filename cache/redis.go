package cache

import (
	"context"
	"public/global"
	"strconv"
	"time"
)

type RedisServer struct {
}

func (s *RedisServer) Exists(key string) int64 {
	result, _ := global.Redis.Exists(context.Background(), key).Result()
	return result
}

func (s *RedisServer) SetNX(key string, value string, expire int) bool {
	result, _ := global.Redis.SetNX(context.Background(), key, value, time.Duration(expire)*time.Second).Result()
	return result
}

func (s *RedisServer) Get(key string) string {
	result, _ := global.Redis.Get(context.Background(), key).Result()
	return result
}

func (s *RedisServer) Delete(key string, value string) int64 {
	if value != "" {
		result, _ := global.Redis.Get(context.Background(), key).Result()
		if result != value {
			return 0
		}
	}
	result, _ := global.Redis.Del(context.Background(), key).Result()
	return result
}

func (s *RedisServer) Clear(keys ...string) int64 {
	result, _ := global.Redis.Del(context.Background(), keys...).Result()
	return result
}

func (s *RedisServer) FuzzyClear(key string) int64 {
	keys := s.keys(key)
	if len(keys) == 0 {
		return 0
	}
	return s.Clear(keys...)
}

func (s *RedisServer) keys(key string) []string {
	result, _ := global.Redis.Keys(context.Background(), key).Result()
	return result
}

func (s *RedisServer) SetMobileSmsLock(m, t string, expire int64) bool {
	redisKey := RedisKey{}
	key := redisKey.GetMobileSmsTypeLockKey(m, t)
	return s.SetNX(key, strconv.Itoa(int(time.Now().Unix())), int(expire))
}

func (s *RedisServer) DelMobileSmsLock(m, t string) int64 {
	redisKey := RedisKey{}
	key := redisKey.GetMobileSmsTypeLockKey(m, t)
	return s.Clear(key)
}
