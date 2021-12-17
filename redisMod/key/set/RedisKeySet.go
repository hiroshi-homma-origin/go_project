package set

import (
	. "fmt"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

func RedisKeySet(c redis.Conn) {
	for i := 1; i < 10001; i++ {
		k := Sprintf("key%d", i)
		t := Sprintf("sample_value%d", i)
		resSet, _ := set(k, t, c)
		Println("set_" + k + "=" + t + ":" + resSet) // OK
	}
}

// データの登録(Redis: SET key value)
func set(key, value string, c redis.Conn) (string, error) {
	return redis.String(c.Do("SET", key, value))
}
