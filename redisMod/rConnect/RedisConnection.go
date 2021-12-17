package rConnect

import "github.com/gomodule/redigo/redis"

func RedisConnection() (redis.Conn, error) {
	// 接続
	return connection()
}

// Connection Test
func connection() (redis.Conn, error) {
	const Addr = "localhost:6379"
	return redis.Dial(
		"tcp",
		Addr,
		redis.DialPassword("dip04634_2"),
		redis.DialDatabase(0),
	)
}
