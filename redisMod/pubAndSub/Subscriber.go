package pubAndSub

import (
	"fmt"
	"go_project/redisMod/rConnect"

	"github.com/gomodule/redigo/redis"
)

func Subscriber() {
	// 接続
	conn, err := rConnect.RedisConnection()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("channel_1", "channel_2", "channel_3")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}

	if err != nil {
		panic(err)
	}
}
