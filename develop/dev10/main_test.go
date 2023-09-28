package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSetConnectionError(t *testing.T) {
	err := setConnection("test123", "9999", time.Duration(time.Second*20))
	if err == nil {
		t.Errorf("Ожидалась ошибка, но её нет")
	} else {
		fmt.Println(err.Error())
	}
}
