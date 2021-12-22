package pInsert

import (
	. "fmt"
	_ "github.com/lib/pq"
	"go_project/psqlMod/pConnect"
	"time"
	_ "time"
)

const n1 = 0
const n2 = 500000
const n3 = 1000000
const n4 = 1500000
const n5 = 2000000
const n6 = 2500000
const n7 = 3000000
const n8 = 3500000
const n9 = 4000000
const n10 = 4500000
const n11 = 5000000

func Insert() {
	now := time.Now()
	Println("*** 開始 ***")

	dictAa := dataPrepareProc()

	db, err := pConnect.RedisConnection()
	defer db.Close()

	//_, _ = db.Exec("drop table employees")

	//sqlStr := "CREATE TABLE IF NOT EXISTS public.employees(id SERIAL NOT NULL, user_id text NOT NULL, user_code text NOT NULL, test_content1 text NOT NULL, test_content2 text NOT NULL, test_content3 text NOT NULL, test_content4 text NOT NULL, test_content5 text NOT NULL, test_content6 text NOT NULL, test_content7 text NOT NULL, test_content8 text NOT NULL, test_content9 text NOT NULL, created text NOT NULL, CONSTRAINT employees_pkey PRIMARY KEY (id))"
	//Println(sqlStr)

	//_, _ = db.Exec(sqlStr)

	for key := range dictAa {
		userId := dictAa[key]["user_id"].(string)
		userCode := dictAa[key]["user_code"].(string)
		testContent1 := dictAa[key]["test_content1"].(string)
		testContent2 := dictAa[key]["test_content2"].(string)
		testContent3 := dictAa[key]["test_content3"].(string)
		testContent4 := dictAa[key]["test_content4"].(string)
		testContent5 := dictAa[key]["test_content5"].(string)
		testContent6 := dictAa[key]["test_content6"].(string)
		testContent7 := dictAa[key]["test_content7"].(string)
		testContent8 := dictAa[key]["test_content8"].(string)
		testContent9 := dictAa[key]["test_content9"].(string)
		created := dictAa[key]["created"].(string)
		sqlStr := "insert into employees(user_id, user_code, test_content1, test_content2, test_content3, test_content4, test_content5, test_content6, test_content7, test_content8, test_content9, created) values (" +
			"'" + userId + "'," +
			"'" + userCode + "'," +
			"'" + testContent1 + "'," +
			"'" + testContent2 + "'," +
			"'" + testContent3 + "'," +
			"'" + testContent4 + "'," +
			"'" + testContent5 + "'," +
			"'" + testContent6 + "'," +
			"'" + testContent7 + "'," +
			"'" + testContent8 + "'," +
			"'" + testContent9 + "'," +
			"'" + created + "'" +
			")"
		//Println(sqlStr)
		_, _ = db.Exec(sqlStr)
	}

	Println("*** 終了 ***")
	Printf("経過: %vms\n", time.Since(now).Milliseconds())
	if err != nil {
		Errorf("")
	}
}

// ----------------------------------------------------------------
func dataPrepareProc() map[int]map[string]interface{} {
	dictAa := make(map[int]map[string]interface{})
	for i := n1; i < n2; i++ {
		//for i := n2; i < n3; i++ {
		//for i := n3; i < n4; i++ {
		//for i := n4; i < n5; i++ {
		//for i := n5; i < n6; i++ {
		//for i := n6; i < n7; i++ {
		//for i := n7; i < n8; i++ {
		//for i := n8; i < n9; i++ {
		//for i := n9; i < n10; i++ {
		//for i := n10; i < n11; i++ {
		Printf("insert_data:%d\n", i)
		userId := Sprintf("user_id_%d", i)
		userCode := Sprintf("user_code_%d", i)
		testContent1 := Sprintf("test_content1_%d", i)
		testContent2 := Sprintf("test_content2_%d", i)
		testContent3 := Sprintf("test_content3_%d", i)
		testContent4 := Sprintf("test_content4_%d", i)
		testContent5 := Sprintf("test_content5_%d", i)
		testContent6 := Sprintf("test_content6_%d", i)
		testContent7 := Sprintf("test_content7_%d", i)
		testContent8 := Sprintf("test_content8_%d", i)
		testContent9 := Sprintf("test_content19_%d", i)
		dictAa[i] = unitGenProc(
			userId,
			userCode,
			testContent1,
			testContent2,
			testContent3,
			testContent4,
			testContent5,
			testContent6,
			testContent7,
			testContent8,
			testContent9,
			"XXXX-XX-XX")
	}
	return dictAa
}

// ----------------------------------------------------------------

func unitGenProc(
	userId string,
	userCode string,
	testContent1 string,
	testContent2 string,
	testContent3 string,
	testContent4 string,
	testContent5 string,
	testContent6 string,
	testContent7 string,
	testContent8 string,
	testContent9 string,
	created string,
) map[string]interface{} {
	unitAa := make(map[string]interface{})
	unitAa["user_id"] = userId
	unitAa["user_code"] = userCode
	unitAa["test_content1"] = testContent1
	unitAa["test_content2"] = testContent2
	unitAa["test_content3"] = testContent3
	unitAa["test_content4"] = testContent4
	unitAa["test_content5"] = testContent5
	unitAa["test_content6"] = testContent6
	unitAa["test_content7"] = testContent7
	unitAa["test_content8"] = testContent8
	unitAa["test_content9"] = testContent9
	unitAa["created"] = created
	return unitAa
}

// ----------------------------------------------------------------
