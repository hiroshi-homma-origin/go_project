package useCase

import (
	"database/sql"
	. "fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var count = 0
var count1 = 0
var countStr = ""
var whenText = ""
var whereText = ""
var now = time.Now()

func UpdateAndSub(rc redis.Conn, pc *sql.DB) {
	psc := redis.PubSubConn{Conn: rc}
	//psc.Subscribe("channel_1", "channel_2", "channel_3")
	psc.Subscribe("channel_1")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			arr1 := strings.Split(string(v.Data), ",")
			bulkUpdate(arr1, pc)
			//Println(sqlStr)
		case redis.Subscription:
			Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}

func bulkUpdate(a []string, pc *sql.DB) {
	count++
	whenText = whenText + "when '" + a[0] + "' then '" + a[1] + "' \n"
	switch count {
	case 1, 10001, 20001, 30001, 40001, 50001, 60001, 70001, 80001, 90001, 100001, 110001, 120001, 130001, 140001, 150001, 160001, 170001, 180001, 190001, 200001, 210001, 220001, 230001, 240001, 250001, 260001, 270001, 280001, 290001:
		whereText = whereText + "('" + a[0] + "'"
	case 10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000, 110000, 120000, 130000, 140000, 150000, 160000, 170000, 180000, 190000, 200000, 210000, 220000, 230000, 240000, 250000, 260000, 270000, 280000, 290000, 300000:
		whereText = whereText + ", '" + a[0] + "')"
	default:
		whereText = whereText + ", '" + a[0] + "'"
	}
	if count == 10000 || count == 20000 || count == 30000 ||
		count == 40000 || count == 50000 || count == 60000 ||
		count == 70000 || count == 80000 || count == 90000 ||
		count == 100000 || count == 110000 || count == 120000 ||
		count == 130000 || count == 140000 || count == 150000 ||
		count == 160000 || count == 170000 || count == 180000 ||
		count == 190000 || count == 200000 || count == 210000 ||
		count == 220000 || count == 230000 || count == 240000 ||
		count == 250000 || count == 260000 || count == 270000 ||
		count == 280000 || count == 290000 || count == 300000 {
		count1++
		countStr = strconv.Itoa(count1)
		//whenText = ""
		//whereText = ""
		sqlStr := "update public.employees \n set user_code = case user_id \n" +
			whenText + " end\n where user_id in \n" +
			whereText + ";"
		pc.Exec(sqlStr)
		whenText = ""
		whereText = ""
		Printf("経過"+countStr+": %vms\n", time.Since(now).Milliseconds())
	}
}
