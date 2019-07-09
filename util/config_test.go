package util

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	var config Config;
	data := config.Get()
	data["test"] = "1"
	fmt.Println(data)

	data_1 := config.Get()
	fmt.Println(data_1)
}