package redis

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	redis "github.com/go-redis/redis/v7"
)

type RDClient struct {
	Client *redis.Client
}

func NewRDClient() *RDClient {
	return &RDClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     conf.GetConf().GetString("db.redis.host")+ ":" + conf.GetConf().GetString("db.redis.port"),
			Username: conf.GetConf().GetString("db.redis.username"), 
			Password: conf.GetConf().GetString("db.redis.password"),
			DB:       conf.Conf.GetInt("db.redis.db"),
		})}
}

func (c *RDClient) Close() {
	c.Client.Close()
}