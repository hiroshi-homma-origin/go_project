package main

import (
	. "fmt"
	_ "github.com/gomodule/redigo/redis"
	"go_project/psqlMod/pConnect"
	"go_project/redisMod/rConnect"
	_ "go_project/redisMod/rConnect"
	"go_project/useCase"
)

func main() {
	Println("Test Application")

	conn, err := rConnect.RedisConnection()
	pc, err := pConnect.RedisConnection()

	defer conn.Close()
	defer pc.Close()

	useCase.UpdateAndSub(conn, pc)

	if err != nil {
		panic(err)
	}
}
