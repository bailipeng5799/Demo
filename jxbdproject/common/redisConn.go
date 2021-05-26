package common

import (
	"github.com/go-redis/redis"
	"log"
)
var (
	Conn *redis.Client
)
func init(){
	Conn = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})
	if _,err := Conn.Ping().Result();err != nil{
		log.Fatalf("Connect to redis client failed,err:%v\n",err)
	}
}
