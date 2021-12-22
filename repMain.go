package main

import (
	. "fmt"
	_ "github.com/gomodule/redigo/redis"
	"github.com/pkg/profile"
	"go_project/psqlMod/pConnect"
	"go_project/redisMod/rConnect"
	"go_project/useCase"
)

func main() {
	Println("Test Application")
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	conn, err := rConnect.RedisConnection()
	pc, err := pConnect.RedisConnection()

	useCase.UpdateAndSub(conn, pc)
	//useCase.SelectAndPub(conn)
	//pInsert.Insert()

	if err != nil {
		panic(err)
	}
}
