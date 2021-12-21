package main

import (
	. "fmt"
	_ "github.com/gomodule/redigo/redis"
	"go_project/psqlMod/pConnect"
	"go_project/redisMod/rConnect"
	"go_project/useCase"
	_ "replication/redisMod/rConnect"
)

func main() {
	Println("Test Application")

	conn, err := rConnect.RedisConnection()
	pc, err := pConnect.RedisConnection()

	defer conn.Close()
	defer pc.Close()

	useCase.UpdateAndSub(conn, pc)
	//useCase.SelectAndPub(conn)

	if err != nil {
		panic(err)
	}
}
