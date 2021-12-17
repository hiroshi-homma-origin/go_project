package pUpdate

import (
	. "fmt"
	_ "github.com/lib/pq"
	"go_project/psqlMod/pConnect"
	"time"
	_ "time"
)

// ----------------------------------------------------------------

func Update() {
	now := time.Now()
	Println("*** 開始 ***")

	dictAa1 := dataPrepareProc1()

	db, err := pConnect.RedisConnection()
	defer db.Close()

	for key := range dictAa1 {
		userId := dictAa1[key]["user_id"].(string)
		testContent1 := dictAa1[key]["test_content1"].(string)
		sqlStr := "update employees set test_content1='" + testContent1 + "' where user_id='" + userId + "';"
		Println(sqlStr)
		_, _ = db.Exec(sqlStr)
	}

	if err != nil {
		Println(err)
	}
	Println("*** 終了 ***")
	Printf("経過: %vms\n", time.Since(now).Milliseconds())
}

// ----------------------------------------------------------------
func dataPrepareProc1() map[int]map[string]interface{} {
	dictAa := make(map[int]map[string]interface{})
	for i := 0; i < 300000; i++ {
		userId := Sprintf("user_id_%d", i)
		testContent1 := Sprintf("test_content1_up_%d", i)
		dictAa[i] = unitGenProc1(
			userId,
			testContent1,
		)
	}
	return dictAa
}

// ----------------------------------------------------------------
func unitGenProc1(
	userId string,
	testContent1 string,
) map[string]interface{} {
	unitAa := make(map[string]interface{})
	unitAa["user_id"] = userId
	unitAa["test_content1"] = testContent1
	return unitAa
}

// ----------------------------------------------------------------
