package dao

import (
	"fmt"
	"testing"
)

func TestQueryOrderDay(t *testing.T) {
	count := QueryOrderDay("2019-06-28")
	// count := QueryOrderDay("now()")
	fmt.Println(count)
}
