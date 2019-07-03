package dao

import (
	"fmt"
	"testing"
)

func TestQueryOrderDay(t *testing.T) {
	count := QueryOrderDay("2019-06-28")
	// count := QueryOrderDay("now()")
	fmt.Println("TestQueryOrderDay", count)
}

func TestQuerySellOrderDay(t *testing.T) {
	count := QuerySellOrderDay("2019-06-28")
	// count := QueryOrderDay("now()")
	fmt.Println("TestQuerySellOrderDay", count)
}
