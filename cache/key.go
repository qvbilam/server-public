package cache

type RedisKey struct {
}

const prefix = "qvbilam:"

func (k *RedisKey) GetMobileSmsTypeLockKey(m, t string) string {
	return prefix + "sms:lock:mobile:" + m + ":type:" + t
}
