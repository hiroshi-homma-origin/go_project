package useCase

import (
	"database/sql"
	. "fmt"
	"github.com/gomodule/redigo/redis"
	"go_project/model"
	"go_project/psqlMod/pConnect"
	"strconv"
	"time"
)

func SelectAndPub(rc redis.Conn) {
	now := time.Now()
	Println("*** 開始1 ***")

	db, err := pConnect.RedisConnection()
	defer db.Close()

	rows, err := db.Query("SELECT user_id, user_code FROM employees")
	rows1, _ := db.Query("SELECT COUNT(*) as count FROM employees")
	c := strconv.Itoa(checkCount(rows1))
	Printf("Total count:%s\n", c)
	for rows.Next() {
		var e model.EMPLOYEE
		rows.Scan(
			&e.USERID,
			&e.USERCODE,
		)
		//updateUserCode := "update_" + e.USERCODE
		updateUserCode := e.USERCODE
		redis.Int(rc.Do(
			"PUBLISH",
			"channel_1",
			e.USERID+","+updateUserCode+","+c))
	}
	Println("*** 終了 ***")
	Printf("経過1: %vms\n", time.Since(now).Milliseconds())
	if err != nil {
		Errorf("")
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	return count
}
