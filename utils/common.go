package utils

import "fmt"

func Int64Ptr(i64 int64) *int64 {
	return &i64
}

func StringPtr(str string) *string {
	return &str
}

func GoFunc(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic err:%v \n", err)
		}
	}()

	go fn()
}
