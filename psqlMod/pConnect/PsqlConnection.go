package pConnect

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func RedisConnection() (*sql.DB, error) {
	strConnect := "user=hiroshi-homma dbname=hiro-test password=dip04634 sslmode=disable"
	return sql.Open("postgres", strConnect)
}
