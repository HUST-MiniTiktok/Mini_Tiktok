package redis

import (
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

func (c *RDClient) Get(key string) string {
	cmd := c.Client.Get(key)
	val, err := cmd.Result()
	if err != nil {
		klog.Errorf("get key %s failed: %v", key, err)
	}
	return val
}

func (c *RDClient) GetInt(key string) int64 {
	cmd := c.Client.Get(key)
	val, err := cmd.Result()
	if err != nil {
		klog.Errorf("get key %s failed: %v", key, err)
	}
	v_int, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		klog.Errorf("get key %s failed: %v", key, err)
	}
	return v_int
}

func (c *RDClient) Set(key string, value interface{}, exp time.Duration) bool {
	err := c.Client.Set(key, value, exp).Err()
	if err != nil {
		klog.Errorf("set key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) Del(key string) bool {
	err := c.Client.Del(key).Err()
	if err != nil {
		klog.Errorf("del key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) Exists(key string) bool {
	v, err := c.Client.Exists(key).Result()
	if err != nil {
		klog.Errorf("exists key %s failed: %v", key, err)
		return false
	}
	return v > 0
}

func (c *RDClient) IncrBy(key string, num int64) bool {
	err := c.Client.IncrBy(key, num).Err()
	if err != nil {
		klog.Errorf("incr key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) DecrBy(key string, num int64) bool {
	err := c.Client.DecrBy(key, num).Err()
	if err != nil {
		klog.Errorf("decr key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) HGet(key string, field string) string {
	val, err := c.Client.HGet(key, field).Result()
	if err != nil {
		klog.Errorf("Hget key %s field %s failed: %v", key, field, err)
	}
	return val
}

func (c *RDClient) HGetInt(key string, field string) int64 {
	val, err := c.Client.HGet(key, field).Result()
	if err != nil {
		klog.Errorf("Hget key %s field %s failed: %v", key, field, err)
	}
	count, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		klog.Errorf("Hget key %s field %s failed: %v", key, field, err)
	}
	return count
}

func (c *RDClient) HSet(key string, field string, value interface{}) bool {
	err := c.Client.HSet(key, field, value).Err()
	if err != nil {
		klog.Errorf("Hset key %s field %s failed: %v", key, field, err)
		return false
	}
	return true
}

func (c *RDClient) HDel(key string, field string) bool {
	err := c.Client.HDel(key, field).Err()
	if err != nil {
		klog.Errorf("Hdel key %s field %s failed: %v", key, field, err)
		return false
	}
	return true
}

func (c *RDClient) HExists(key string, field string) bool {
	v, err := c.Client.HExists(key, field).Result()
	if err != nil {
		klog.Errorf("Hexists key %s field %s failed: %v", key, field, err)
		return false
	}
	return v
}

func (c *RDClient) HIncr(key string, field string, num int64) bool {
	err := c.Client.HIncrBy(key, field, num).Err()
	if err != nil {
		klog.Errorf("Hinc key %s field %s failed: %v", key, field, err)
		return false
	}
	return true
}

func (c *RDClient) HExpire(key string, tm time.Duration) bool {
	err := c.Client.Do("expire", key, tm.Seconds())
	if err != nil {
		klog.Errorf("HExpires key %s failed: %v", key, err)
		return false
	}
	return true
}
