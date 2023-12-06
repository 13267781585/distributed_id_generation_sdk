package main

import (
	"context"
	"time"
	"uuid_client/rpc"
)

func main() {
	for {
		rpc.GetUUIDBounds(context.Background(), 111, 1)
		time.Sleep(5 * time.Second)
	}
}
