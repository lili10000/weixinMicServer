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

// type NullInt64 sql.NullInt64

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

var condition = "(CO_STATE < 11 and CO_STATE >= 0 )"

func QueryOrderDay(day string) string {
	sql := fmt.Sprintf("SELECT count(distinct CO_SERIAL_CODE) FROM view_order_pay where (TO_DAYS(CO_RESERVATION_DATE) = TO_DAYS('%s'))", day)
	return QueryCountBySql(sql)
}

func QuerySellOrderDay(day string) string {
	sql := fmt.Sprintf("SELECT count(distinct CO_SERIAL_CODE) FROM view_order_pay where (TO_DAYS(CO_RESERVATION_DATE) = TO_DAYS('%s') and %s)", day, condition)
	return QueryCountBySql(sql)
}

func QuerySellNumDay(day string) string {
	sql := fmt.Sprintf("select sum(MD_COUNT) from view_order_pay where (TO_DAYS(CO_RESERVATION_DATE)=TO_DAYS('%s') and %s)", day, condition)
	return QueryCountBySql(sql)
}

func QuerySellMoneyPriceDay(day string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ORIGINAL_ACOUNT),2) from view_order_pay where (TO_DAYS(CO_RESERVATION_DATE)=TO_DAYS('%s') and %s)", day, condition)
	return QueryCountBySql(sql)
}

// 实收
func QuerySellMoneyRecvDay(day string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ACOUNT),2) from view_order_pay where (TO_DAYS(CO_RESERVATION_DATE)=TO_DAYS('%s') and %s)", day, condition)
	return QueryCountBySql(sql)
}

func QueryDayInfo(time string) []string {
	sql := fmt.Sprintf("select IFNULL(count(distinct CO_SERIAL_CODE),0),IFNULL(sum(MD_COUNT),0), IFNULL(ROUND(sum(MD_ORIGINAL_ACOUNT),2), 0), IFNULL(ROUND(sum(MD_ACOUNT),2),0)  from view_order_pay where (TO_DAYS(CO_RESERVATION_DATE)=TO_DAYS('%s') and %s)", time, condition)
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	retnList := make([]string, 0)
	for rows.Next() {
		var count, sell, price, recv string
		err = rows.Scan(&count, &sell, &price, &recv)
		checkErr(err)
		retnList = append(retnList, count)
		retnList = append(retnList, sell)
		retnList = append(retnList, price)
		retnList = append(retnList, recv)
	}
	return retnList
}

type Element []string
type SessionList []Element

func QuerySession(day string) SessionList {
	sql := fmt.Sprintf("select CA_ADDRESS,CO_RESERVATION_DATE from view_order_refund_amount where (TO_DAYS(CO_RESERVATION_DATE)=TO_DAYS('%s') );", day)
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	retnList := make(SessionList, 0)
	for rows.Next() {
		var addr string
		var endData string
		err = rows.Scan(&addr, &endData)
		checkErr(err)
		ele := make([]string, 0)
		ele = append(ele, addr)
		ele = append(ele, endData)
		retnList = append(retnList, ele)
	}
	return retnList
}

func QuerySellOrderSession(time string) string {
	sql := fmt.Sprintf("SELECT count(distinct CO_SERIAL_CODE) FROM view_order_pay where (CO_RESERVATION_DATE = '%s' and %s)", time, condition)
	return QueryCountBySql(sql)
}

func QuerySellNumSession(time string) string {
	sql := fmt.Sprintf("select sum(MD_COUNT) from view_order_pay where(CO_RESERVATION_DATE = '%s' and %s)", time, condition)
	return QueryCountBySql(sql)
}

func QuerySellMoneyPriceSession(time string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ORIGINAL_ACOUNT),2) from view_order_pay where (CO_RESERVATION_DATE = '%s' and %s)", time, condition)
	return QueryCountBySql(sql)
}

// 实收
func QuerySellMoneyRecvSession(time string) string {
	sql := fmt.Sprintf("select ROUND(sum(MD_ACOUNT),2) from view_order_pay where (CO_RESERVATION_DATE = '%s' and %s)", time, condition)
	return QueryCountBySql(sql)
}

func QuerySessionInfo(time string) []string {
	sql := fmt.Sprintf("select IFNULL(count(distinct CO_SERIAL_CODE),0),IFNULL(sum(MD_COUNT),0), IFNULL(ROUND(sum(MD_ORIGINAL_ACOUNT),2), 0), IFNULL(ROUND(sum(MD_ACOUNT),2),0) from view_order_pay where (CO_RESERVATION_DATE = '%s' and %s)", time, condition)
	// fmt.Println(sql)
	rows, err := db.Query(sql)
	defer rows.Close()
	checkErr(err)
	retnList := make([]string, 0)
	for rows.Next() {
		var count, sell, price, recv string
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				log.Error(err)
			}
		}()
		err = rows.Scan(&count, &sell, &price, &recv)

		checkErr(err)
		retnList = append(retnList, count)
		retnList = append(retnList, sell)
		retnList = append(retnList, price)
		retnList = append(retnList, recv)
	}
	return retnList
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
