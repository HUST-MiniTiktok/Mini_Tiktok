package redis

import (
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	LockSuffix  = ":lock"
	RetryTime   = 30 * time.Millisecond
	LockTimeout = 2 * time.Second
)

func (c *RDClient) AcquireLock(Id int64, field string) bool {
	key := strconv.FormatInt(Id, 10) + field + LockSuffix
	// SetNX 2s
	result, err := c.Client.SetNX(key, 1, LockTimeout).Result()
	if err != nil {
		klog.Errorf("Acquire Lock failed: %v", err)
		return false
	}
	return result
}

func (c *RDClient) ReleaseLock(Id int64, field string) {
	key := strconv.FormatInt(Id, 10) + field + LockSuffix
	c.Client.Del(key)
}
