package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"weixin/util"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var log util.LogUtil

func init() {
	var err error
	db, err = sql.Open("mysql", "gac:123MAGICBOX!@#@tcp(gz-cdb-bx3dc6o1.sql.tencentcdb.com:61283)/gac_mobile_vehicle?charset=utf8")
	if err != nil {
		log.Error(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		log.Error(err)
	}
}

func QueryCountBySql(sql string) int{
	defer func() { //统一异常处理
		if err := recover(); err != nil {
			log.Error(err, "sql:",sql)
		}
	}()
	count := 0
	if db == nil {
		panic(errors.New("db is nil"))
	}
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		var countStr string
		err = rows.Scan(&countStr)
		checkErr(err)
		count, err = strconv.Atoi(countStr)
		checkErr(err)
		break
	}
	return count
}

func QueryBySql(sql string) int{
	defer func() { //统一异常处理
		if err := recover(); err != nil {
			log.Error(err, "sql:",sql)
		}
	}()
	count := 0
	if db == nil {
		panic(errors.New("db is nil"))
	}
	sql = ""
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		var countStr string
		err = rows.Scan(&countStr)
		checkErr(err)
		fmt.Println(countStr)
		// count, err = strconv.Atoi(countStr)
		// checkErr(err)
		break
	}
	return count
}

func QueryOrderDay(day string) int {
	sql := fmt.Sprintf("SELECT count(*) FROM customer_order where (TO_DAYS(CO_ORDER_DATE) = TO_DAYS('%s'))", day)
	return QueryCountBySql(sql)
}

func QuerySellOrderDay(day string) int {
	sql := fmt.Sprintf("SELECT count(*) FROM customer_order where (TO_DAYS(CO_ORDER_DATE) = TO_DAYS('%s') and CO_PAY_STATE = 1 )", day)
	return QueryCountBySql(sql)
}

func QuerySellNumDay(day string) int {
	sql := fmt.Sprintf("select sum(nums) from view_order_pay where TO_DAYS(下单时间)=TO_DAYS('%s');", day)
	return QueryCountBySql(sql)
}

func QueryOrderToday() int {
	return QueryOrderDay("now()")
}
func QuerySellOrderToday() int {
	return QuerySellOrderDay("now()")
}

func QuerySellNumToday() int {
	return QuerySellNumDay("now()")
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
