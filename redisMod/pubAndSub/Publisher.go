package pubAndSub

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_project/redisMod/rConnect"
)

func Publisher() {
	// 接続
	conn, err := rConnect.RedisConnection()
	defer conn.Close()
	// パブリッシュ
	r, err := redis.Int(conn.Do("PUBLISH", "channel_1", "hello"))
	fmt.Println(r)
	if err != nil {
		panic(err)
	}
}
