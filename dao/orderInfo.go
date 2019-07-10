package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"weixin/util"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var log util.LogUtil
var config util.Config

func init() {
	var err error
	mysqlConfig := config.Get()
	// fmt.Println(mysqlConfig)

	db, err = sql.Open("mysql", mysqlConfig.Mysql.Server)
	if err != nil {
		log.Error(err)
	}
	db.SetMaxOpenConns(mysqlConfig.Mysql.MaxConn)
	db.SetMaxIdleConns(mysqlConfig.Mysql.MaxIdleConn)
	err = db.Ping()
	if err != nil {
		log.Error(err)
	}
}

func QueryCountBySql(sql string) string {
	defer func() { //统一异常处理
		if err := recover(); err != nil {
			log.Error(err, "sql:", sql)
		}
	}()
	var countStr string
	if db == nil {
		panic(errors.New("db is nil"))
	}
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&countStr)
		checkErr(err)
		break
	}
	return countStr
}

func QueryOrderDay(day string) string {
	sql := fmt.Sprintf("SELECT count(distinct CO_SERIAL_CODE) FROM view_order_pay where (TO_DAYS(CO_ORDER_DATE) = TO_DAYS('%s'))", day)
	return QueryCountBySql(sql)
}

func QuerySellOrderDay(day string) string {
	sql := fmt.Sprintf("SELECT count(distinct CO_SERIAL_CODE) FROM view_order_pay where (TO_DAYS(CO_ORDER_DATE) = TO_DAYS('%s') and CO_STATE = 10 )", day)
	return QueryCountBySql(sql)
}

func QuerySellNumDay(day string) string {
	sql := fmt.Sprintf("select sum(MD_COUNT) from view_order_pay where (TO_DAYS(CO_ORDER_DATE)=TO_DAYS('%s') and CO_STATE = 10);", day)
	return QueryCountBySql(sql)
}

func QuerySellMoneyPriceDay(day string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ORIGINAL_ACOUNT),2) from view_order_pay where TO_DAYS(CO_ORDER_DATE)=TO_DAYS('%s');", day)
	return QueryCountBySql(sql)
}

// 实收
func QuerySellMoneyRecvDay(day string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ACOUNT),2) from view_order_pay where TO_DAYS(CO_ORDER_DATE)=TO_DAYS('%s');", day)
	return QueryCountBySql(sql)
}

// func QuerySellMoneyRecvToday() string {
// 	return QuerySellMoneyPriceDay("now()")
// }

// func QuerySellMoneyPriceToday() string {
// 	return QuerySellMoneyPriceDay("now()")
// }

// func QueryOrderToday() string {
// 	return QueryOrderDay("now()")
// }
// func QuerySellOrderToday() string {
// 	return QuerySellOrderDay("now()")
// }

// func QuerySellNumToday() string {
// 	return QuerySellNumDay("now()")
// }

func checkErr(err error) {
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
