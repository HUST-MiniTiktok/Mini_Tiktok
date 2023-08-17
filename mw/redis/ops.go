package redis

import (
	"strconv"
	"time"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (c *RDClient) Get(key string) (string) {
	cmd := c.Client.Get(key)
	val, err := cmd.Result()
	if err != nil {
		klog.Errorf("get key %s failed: %v", key, err)
	}
	return val
}

func (c *RDClient) GetInt(key string) (int) {
	cmd := c.Client.Get(key)
	val, err := cmd.Result()
	if err != nil {
		klog.Errorf("get key %s failed: %v", key, err)
	}
	v_int, err := strconv.Atoi(val)
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

func (c *RDClient) Exists(key string) (bool) {
	v, err := c.Client.Exists(key).Result()
	if err != nil {
		klog.Errorf("exists key %s failed: %v", key, err)
		return false
	}
	return v > 0
}

