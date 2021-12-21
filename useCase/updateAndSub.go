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
var count3 = 0
var count4 = 0
var count5 = 0
var count6 = 0

var count11 = 0

var whenText1 = ""
var whereText1 = ""

var whenText2 = ""
var whereText2 = ""

var whenText3 = ""
var whereText3 = ""

var whenText4 = ""
var whereText4 = ""

var whenText5 = ""
var whereText5 = ""

var whenText6 = ""
var whereText6 = ""

var countStr1 = ""

var now1 = time.Now()

func UpdateAndSub(rc redis.Conn, pc *sql.DB) {
	psc := redis.PubSubConn{Conn: rc}
	//psc.Subscribe("channel_1", "channel_2", "channel_3")
	psc.Subscribe("channel_1", "channel_2", "channel_3", "channel_4", "channel_5", "channel_6")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			arr1 := strings.Split(string(v.Data), ",")
			bulkUpdate(arr1, pc, v.Channel)
			//Println(sqlStr)
		case redis.Subscription:
			Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}

func bulkUpdate(a []string, pc *sql.DB, channel string) {
	switch channel {
	case "channel_2":
		count2++
		whenText2 = whenText2 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count2 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText2 = whereText2 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			whereText2 = whereText2 + ", '" + a[0] + "')"
			sqlStr2 := "update public.employees \n set user_code = case user_id \n" +
				whenText2 + " end\n where user_id in \n" +
				whereText2 + ";"
			pc.Exec(sqlStr2)
			whenText2 = ""
			whereText2 = ""
		default:
			whereText2 = whereText2 + ", '" + a[0] + "'"
		}
	case "channel_3":
		count3++
		whenText3 = whenText3 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count3 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText3 = whereText3 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			whereText3 = whereText3 + ", '" + a[0] + "')"
			sqlStr3 := "update public.employees \n set user_code = case user_id \n" +
				whenText3 + " end\n where user_id in \n" +
				whereText3 + ";"
			pc.Exec(sqlStr3)
			whenText3 = ""
			whereText3 = ""
		default:
			whereText3 = whereText3 + ", '" + a[0] + "'"
		}
	case "channel_4":
		count4++
		whenText4 = whenText4 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count4 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText4 = whereText4 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			whereText4 = whereText4 + ", '" + a[0] + "')"
			sqlStr4 := "update public.employees \n set user_code = case user_id \n" +
				whenText4 + " end\n where user_id in \n" +
				whereText4 + ";"
			pc.Exec(sqlStr4)
			whenText4 = ""
			whereText4 = ""
		default:
			whereText4 = whereText4 + ", '" + a[0] + "'"
		}
	case "channel_5":
		count5++
		whenText5 = whenText5 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count5 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText5 = whereText5 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			whereText5 = whereText5 + ", '" + a[0] + "')"
			sqlStr5 := "update public.employees \n set user_code = case user_id \n" +
				whenText5 + " end\n where user_id in \n" +
				whereText5 + ";"
			pc.Exec(sqlStr5)
			whenText5 = ""
			whereText5 = ""
		default:
			whereText5 = whereText5 + ", '" + a[0] + "'"
		}
	case "channel_6":
		count6++
		whenText6 = whenText6 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count6 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText6 = whereText6 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			whereText6 = whereText6 + ", '" + a[0] + "')"
			sqlStr6 := "update public.employees \n set user_code = case user_id \n" +
				whenText6 + " end\n where user_id in \n" +
				whereText6 + ";"
			pc.Exec(sqlStr6)
			whenText6 = ""
			whereText6 = ""
		default:
			whereText6 = whereText6 + ", '" + a[0] + "'"
		}
	default:
		count1++
		whenText1 = whenText1 + "when '" + a[0] + "' then '" + a[1] + "' \n"
		switch count1 {
		case 1, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001:
			whereText1 = whereText1 + "('" + a[0] + "'"
		case 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			count11++
			countStr1 = strconv.Itoa(count11)
			whereText1 = whereText1 + ", '" + a[0] + "')"
			sqlStr := "update public.employees \n set user_code = case user_id \n" +
				whenText1 + " end\n where user_id in \n" +
				whereText1 + ";"
			pc.Exec(sqlStr)
			whenText1 = ""
			whereText1 = ""
			Printf("経過_"+countStr1+": %vms\n", time.Since(now1).Milliseconds())
		default:
			whereText1 = whereText1 + ", '" + a[0] + "'"
		}
	}
}
