package dao

import (
	"fmt"
	"testing"
)

// func TestQueryOrderDay(t *testing.T) {
// 	count := QueryOrderDay("2019-06-28")
// 	// count := QueryOrderDay("now()")
// 	fmt.Println("TestQueryOrderDay", count)
// }

func TestQuerySellMoneyPriceDay(t *testing.T) {
	count := QuerySellMoneyPriceDay("2019-06-28")
	// count := QueryOrderDay("now()")
	fmt.Println("TestQuerySellMoneyPriceDay", count)
}

func TestQuerySellMoneyRecvDay(t *testing.T) {
	count := QuerySellMoneyRecvDay("2019-06-28")
	// count := QueryOrderDay("now()")
	fmt.Println("TestQuerySellMoneyRecvDay", count)
}

// func TestQuerySellOrderDay(t *testing.T) {
// 	count := QuerySellOrderDay("2019-06-28")
// 	// count := QueryOrderDay("now()")
// 	fmt.Println("TestQuerySellOrderDay", count)
// }

// func TestQuerySellNumToday(t *testing.T) {
// 	count := QuerySellNumDay("2019-06-28")
// 	// count := QueryOrderDay("now()")
// 	fmt.Println("TestQuerySellNumToday", count)
// }
