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
		klog.Errorf("Get key %s failed: %v", key, err)
	}
	return val
}

func (c *RDClient) GetInt(key string) int64 {
	cmd := c.Client.Get(key)
	val, err := cmd.Result()
	if err != nil {
		klog.Errorf("Get key %s failed: %v", key, err)
	}
	v_int, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		klog.Errorf("Get key %s failed: %v", key, err)
	}
	return v_int
}

func (c *RDClient) Set(key string, value interface{}, exp time.Duration) bool {
	err := c.Client.Set(key, value, exp).Err()
	if err != nil {
		klog.Errorf("Set key %s:%v failed: %v", key, value, err)
		return false
	}
	return true
}

func (c *RDClient) Del(key string) bool {
	err := c.Client.Del(key).Err()
	if err != nil {
		klog.Errorf("Del key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) Exists(key string) bool {
	v, err := c.Client.Exists(key).Result()
	if err != nil {
		klog.Errorf("Exists key %s failed: %v", key, err)
		return false
	}
	return v > 0
}

func (c *RDClient) IncrBy(key string, num int64) bool {
	err := c.Client.IncrBy(key, num).Err()
	if err != nil {
		klog.Errorf("Incr key %s by %d failed: %v", key, num, err)
		return false
	}
	return true
}

func (c *RDClient) DecrBy(key string, num int64) bool {
	err := c.Client.DecrBy(key, num).Err()
	if err != nil {
		klog.Errorf("Decr key %s by %d failed: %v", key, num, err)
		return false
	}
	return true
}

func (c *RDClient) Expire(key string, tm time.Duration) bool {
	err := c.Client.Expire(key, tm).Err()
	if err != nil {
		klog.Errorf("Expire key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) SAdd(key string, value interface{}, exp time.Duration) bool {
	tx := c.Client.TxPipeline()
	tx.SAdd(key, value)
	tx.Expire(key, exp)
	_, err := tx.Exec()
	if err != nil {
		klog.Errorf("SAdd key %s:%v failed: %v", key, value, err)
		return false
	}
	return true
}

func (c *RDClient) SGetIntSlice(key string) []int64 {
	val, err := c.Client.SMembers(key).Result()
	if err != nil {
		klog.Errorf("SGetIntSlice key %s failed: %v", key, err)
		return nil
	}
	ret := make([]int64, 0, len(val))
	for _, v := range val {
		v_int, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			klog.Errorf("SGetIntSlice key %s failed: %v", key, err)
			continue
		}
		ret = append(ret, v_int)
	}
	return ret
}

func (c *RDClient) SGetInterIntSlice(key_1 string, key_2 string) []int64 {
	v, err := c.Client.SInter(key_1, key_2).Result()
	if err != nil {
		klog.Errorf("SInterIntSlice key %s:%s failed: %v", key_1, key_2, err)
		return nil
	}
	ret := make([]int64, 0, len(v))
	for _, val := range v {
		v_int, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			klog.Errorf("SInterIntSlice key %s:%s failed: %v", key_1, key_2, err)
			continue
		}
		ret = append(ret, v_int)
	}
	return ret
}

func (c *RDClient) SRem(key string, value interface{}) bool {
	err := c.Client.SRem(key, value).Err()
	if err != nil {
		klog.Errorf("SRem key %s:%v failed: %v", key, value, err)
		return false
	}
	return true
}

func (c *RDClient) SExistKey(key string) bool {
	v, err := c.Client.Exists(key).Result()
	if err != nil {
		klog.Errorf("SExist key %s failed: %v", key, err)
		return false
	}
	return v > 0
}

func (c *RDClient) SExistValue(key string, value interface{}) bool {
	v, err := c.Client.SIsMember(key, value).Result()
	if err != nil {
		klog.Errorf("SExist key %s:%v failed: %v", key, value, err)
		return false
	}
	return v
}

func (c *RDClient) SCount(key string) int64 {
	v, err := c.Client.SCard(key).Result()
	if err != nil {
		klog.Errorf("SCount key %s failed: %v", key, err)
		return 0
	}
	return v
}

func (c *RDClient) SCountInter(key_1 string, key_2 string) int64 {
	v, err := c.Client.SInter(key_1, key_2).Result()
	if err != nil {
		klog.Errorf("SCountInter key %s:%s failed: %v", key_1, key_2, err)
		return 0
	}
	return int64(len(v))
}

func (c *RDClient) SExpire(key string, exp time.Duration) bool {
	err := c.Client.Expire(key, exp).Err()
	if err != nil {
		klog.Errorf("SExpire key %s failed: %v", key, err)
		return false
	}
	return true
}

func (c *RDClient) HGet(key string, field string) string {
	val, err := c.Client.HGet(key, field).Result()
	if err != nil {
		klog.Errorf("HGet key %s field %s failed: %v", key, field, err)
	}
	return val
}

func (c *RDClient) HGetInt(key string, field string) int64 {
	val, err := c.Client.HGet(key, field).Result()
	if err != nil {
		klog.Errorf("HGet key %s field %s failed: %v", key, field, err)
	}
	count, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		klog.Errorf("HGet key %s field %s failed: %v", key, field, err)
	}
	return count
}

func (c *RDClient) HSet(key string, field string, value interface{}, exp time.Duration) bool {
	tx := c.Client.TxPipeline()
	tx.HSet(key, field, value)
	tx.Expire(key, exp)
	_, err := tx.Exec()
	if err != nil {
		klog.Errorf("HSet key %s field %s:%v failed: %v", key, field, value, err)
		return false
	}
	return true
}

func (c *RDClient) HDel(key string, field string) bool {
	err := c.Client.HDel(key, field).Err()
	if err != nil {
		klog.Errorf("HDel key %s field %s failed: %v", key, field, err)
		return false
	}
	return true
}

func (c *RDClient) HExists(key string, field string) bool {
	v, err := c.Client.HExists(key, field).Result()
	if err != nil {
		klog.Errorf("HExists key %s field %s failed: %v", key, field, err)
		return false
	}
	return v
}

func (c *RDClient) HIncr(key string, field string, num int64) bool {
	err := c.Client.HIncrBy(key, field, num).Err()
	if err != nil {
		klog.Errorf("HIncr key %s field %s by %d failed: %v", key, field, num, err)
		return false
	}
	return true
}

func (c *RDClient) HDecr(key string, field string, num int64) bool {
	err := c.Client.HIncrBy(key, field, -num).Err()
	if err != nil {
		klog.Errorf("HDecr key %s field %s by %d failed: %v", key, field, num, err)
		return false
	}
	return true
}

func (c *RDClient) HExpire(key string, exp time.Duration) bool {
	err := c.Client.Expire(key, exp).Err()
	if err != nil {
		klog.Errorf("HExpire key %s failed: %v", key, err)
		return false
	}
	return true
}
