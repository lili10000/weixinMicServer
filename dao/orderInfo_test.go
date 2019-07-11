package dao

import (
	"fmt"
	"testing"
)

// func TestQueryOrderDay(t *testing.T) {
// 	count := QueryOrderDay("2019-07-10")
// 	fmt.Println("TestQueryOrderDay", count)
// }

// func TestQuerySellNumDay(t *testing.T) {
// 	count := QuerySellNumDay("2019-07-10")
// 	fmt.Println("TestQuerySellNumDay", count)
// }

// func TestQuerySellMoneyPriceDay(t *testing.T) {
// 	count := QuerySellMoneyPriceDay("2019-07-10")
// 	fmt.Println("TestQuerySellMoneyPriceDay", count)
// }

// func TestQuerySellMoneyRecvDay(t *testing.T) {
// 	count := QuerySellMoneyRecvDay("2019-07-10")
// 	fmt.Println("TestQuerySellMoneyRecvDay", count)
// }
// func TestQuerySession(t *testing.T) {
// 	count := QuerySession("2019-07-10")
// 	fmt.Println("TestQuerySellMoneyRecvDay", count)
// }

func TestQuerySellOrderSession(t *testing.T) {
	count := QuerySellOrderSession("2019-07-10 12:00:00")
	fmt.Println("QuerySellOrderSession", count)
}

func TestQuerySellNumSession(t *testing.T) {
	count := QuerySellNumSession("2019-07-10 12:00:00")
	fmt.Println("QuerySellNumSession", count)
}

func TestQuerySellMoneyPriceSession(t *testing.T) {
	count := QuerySellMoneyPriceSession("2019-07-10 12:00:00")
	fmt.Println("QuerySellMoneyPriceSession", count)
}

func TestQuerySellMoneyRecvSession(t *testing.T) {
	count := QuerySellMoneyRecvSession("2019-07-10 12:00:00")
	fmt.Println("QuerySellMoneyRecvSession", count)
}

// func TestQuerySellOrderDay(t *testing.T) {
// 	count := QuerySellOrderDay("2019-07-10")
// 	fmt.Println("TestQuerySellOrderDay", count)
// }

// func TestQuerySellNumToday(t *testing.T) {
// 	count := QuerySellNumDay("2019-06-28")
// 	// count := QueryOrderDay("now()")
// 	fmt.Println("TestQuerySellNumToday", count)
// }
