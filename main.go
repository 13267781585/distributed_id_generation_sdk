package main

import (
	"fmt"
	"time"
	"uuid_client/logic"
	"uuid_client/utils"
)

func main() {
	n := 100
	for i := 0; i < n; i++ {
		utils.GoFunc(MGetUUIDTest)
	}
	time.Sleep(30 * time.Second)
}

func GetUUIDTest() {
	idGen := logic.NewUUIDGenerator(logic.WithQPSBuffer(1), logic.WithBizCode(1), logic.WithMinFetchCount(1))
	time.Sleep(3 * time.Second)
	uuid, ok := idGen.GetUUID()
	if ok {
		fmt.Printf("main get uuid:%v \n", uuid)
	} else {
		fmt.Println("main get uuid not ok")
	}
}

func MGetUUIDTest() {
	idGen := logic.NewUUIDGenerator(logic.WithQPSBuffer(1), logic.WithBizCode(1), logic.WithMinFetchCount(1))
	time.Sleep(3 * time.Second)
	uuids, ok := idGen.MGetUUIDs(100)
	if ok {
		fmt.Printf("main mget uuid:%v \n", uuids)
	} else {
		fmt.Println("main mget uuid not ok")
	}
}
