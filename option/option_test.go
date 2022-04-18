package option

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	cachingOption := WithCaching(true)
	timeoutOption := WithTimeout(2 * time.Hour)

	fmt.Printf("type cacheingOption:%v value:%v\n", reflect.TypeOf(cachingOption), reflect.TypeOf(cachingOption))
	fmt.Printf("type timeoutOption:%v value:%v\n", reflect.TypeOf(timeoutOption), reflect.TypeOf(timeoutOption))

	connect, err := Connect("127.0.0.1", cachingOption, timeoutOption)
	if err != nil {
		fmt.Println("连接出现错误")
	} else {
		fmt.Println(connect)
	}
}
