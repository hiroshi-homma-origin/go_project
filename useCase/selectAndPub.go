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

var sCount = 0

func SelectAndPub(rc redis.Conn) {
	now := time.Now()
	Println("*** 開始1 ***")

	db, err := pConnect.RedisConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, user_code FROM employees where id < 20001")
	rows1, _ := db.Query("SELECT COUNT(*) as count FROM employees where id < 20001")
	c := strconv.Itoa(checkCount(rows1))
	Printf("Total count:%s\n", c)
	for rows.Next() {
		sCount++
		var e model.EMPLOYEE
		rows.Scan(
			&e.ID,
			&e.USERID,
			&e.USERCODE,
		)
		updateUserCode := "update_" + e.USERCODE
		//updateUserCode := e.USERCODE
		println(sCount)
		if sCount == 1 {
			redis.Int(rc.Do(
				"PUBLISH",
				"channel_1",
				e.USERID+","+updateUserCode+","+c))
		} else {
			redis.Int(rc.Do(
				"PUBLISH",
				"channel_2",
				e.USERID+","+updateUserCode+","+c))
			sCount = 0
		}
	}
	Println("*** 終了 ***")
	Printf("経過0: %vms\n", time.Since(now).Milliseconds())
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
