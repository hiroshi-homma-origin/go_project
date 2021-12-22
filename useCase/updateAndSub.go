package useCase

import (
	"database/sql"
	. "fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var count1 = 0
var count2 = 0

var count11 = 0
var count12 = 0

var whenText1 = ""
var wheretext11 = ""
var wheretext12 = ""

var whenText2 = ""
var wheretext21 = ""
var wheretext22 = ""

var countStr1 = ""
var countStr2 = ""

var now1 = time.Now()

func UpdateAndSub(rc redis.Conn, pc *sql.DB) {
	psc := redis.PubSubConn{Conn: rc}
	psc.Subscribe("channel_1", "channel_2")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			arr1 := strings.Split(string(v.Data), ",")
			bulkUpdate(arr1, v.Channel, pc)
		case redis.Subscription:
			Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}

func bulkUpdate(a []string, channel string, pc *sql.DB) {
	switch channel {
	case "channel_1":
		count1++
		whenText1 = whenText1 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count1 {
		case 1:
			wheretext11 = wheretext11 + "('" + a[0] + "'"
			wheretext12 = wheretext12 + "(" + a[3] + ""
		case 500:
			count11++
			countStr1 = strconv.Itoa(count11)
			wheretext11 = wheretext11 + ", '" + a[0] + "')"
			wheretext12 = wheretext12 + ", " + a[3] + ")"
			transaction(whenText1, wheretext11, wheretext12, pc)
			//Println("check_data1:"+whereText1)
			Printf("経過1_"+countStr1+": %vms\n", time.Since(now1).Milliseconds())
			whenText1 = ""
			wheretext11 = ""
			wheretext12 = ""
			if count11 == 5000 {
				count11 = 0
			}
			if count1 == 500 {
				count1 = 0
			}
		default:
			wheretext11 = wheretext11 + ", '" + a[0] + "'"
			wheretext12 = wheretext12 + ", " + a[3]
		}
	case "channel_2":
		count2++
		whenText2 = whenText2 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count2 {
		case 1:
			wheretext21 = wheretext21 + "('" + a[0] + "'"
			wheretext22 = wheretext22 + "(" + a[3] + ""
		case 500:
			count12++
			countStr2 = strconv.Itoa(count12)
			wheretext21 = wheretext21 + ", '" + a[0] + "')"
			wheretext22 = wheretext22 + ", " + a[3] + ")"
			transaction(whenText2, wheretext21, wheretext22, pc)
			//Println("check_data1:"+whereText1)
			Printf("経過2_"+countStr2+": %vms\n", time.Since(now1).Milliseconds())
			whenText2 = ""
			wheretext21 = ""
			wheretext22 = ""
			if count12 == 5000 {
				count12 = 0
			}
			if count2 == 500 {
				count2 = 0
			}
		default:
			wheretext21 = wheretext21 + ", '" + a[0] + "'"
			wheretext22 = wheretext22 + ", " + a[3]
		}
	}
}

func transaction(whenStr string, whereStr1 string, whereStr2 string, pc *sql.DB) {
	sqlStr := "update public.employees \n set user_code = case user_id \n" +
		whenStr + " end\n where user_id in \n" +
		whereStr1 + " AND id in " +
		whereStr2 + ";"
	pc.Exec(sqlStr)
}
