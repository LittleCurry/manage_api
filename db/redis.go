package db

import "gopkg.in/redis.v5"

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "wx.iguiyu.com:56379",
		Password: "",
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	return client
}
