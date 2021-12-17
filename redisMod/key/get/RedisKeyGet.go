package get

import (
	. "fmt"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

func RedisKeyGet(c redis.Conn) {
	for i := 1; i < 300001; i++ {
		k := Sprintf("key%d", i)
		resSet, _ := get(k, c)
		Println("get_" + k + ":" + resSet) // OK
	}
}

// データの取得(Redis: GET key)
func get(key string, c redis.Conn) (string, error) {
	return redis.String(c.Do("GET", key))
}
