package main

import (
	"context"
	"time"
	"uuid_client/rpc"
	_ "uuid_client/rpc"
)

func main() {
	for {
		rpc.GetUUIDBounds(context.Background(), 111, 1)
		time.Sleep(10 * time.Second)
	}
}
