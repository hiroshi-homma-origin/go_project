package pSelect

import (
	. "fmt"
	"go_project/model"
	"go_project/psqlMod/pConnect"
	"time"
)

func Select() {
	now := time.Now()
	Println("*** 開始1 ***")

	db, err := pConnect.RedisConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, user_code FROM employees")

	var es []model.EMPLOYEE
	for rows.Next() {
		var e model.EMPLOYEE
		rows.Scan(
			&e.USERID,
			&e.USERCODE,
		)
		es = append(es, e)
		Println(e)
	}
	Println("*** 終了 ***")
	Printf("経過1: %vms\n", time.Since(now).Milliseconds())
	if err != nil {
		Errorf("")
	}
}
