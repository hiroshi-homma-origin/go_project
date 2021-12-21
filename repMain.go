package main

import (
	. "fmt"
	_ "github.com/gomodule/redigo/redis"
	"replication/psqlMod/pConnect"
	"replication/redisMod/rConnect"
	_ "replication/redisMod/rConnect"
	"replication/useCase"
)

func main() {
	Println("Test Application")

	conn, err := rConnect.RedisConnection()
	pc, err := pConnect.RedisConnection()

	defer conn.Close()
	defer pc.Close()

	//useCase.UpdateAndSub(conn, pc)
	useCase.SelectAndPub(conn)

	if err != nil {
		panic(err)
	}
}
