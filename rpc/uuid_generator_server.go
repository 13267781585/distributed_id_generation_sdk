package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/luci/go-render/render"

	"uuid_client/kitex_gen/uuid/generator/server"
	"uuid_client/kitex_gen/uuid/generator/server/uuidgeneratorserver"
	"uuid_client/utils"
)

var uuidClient uuidgeneratorserver.Client

func init() {
	var err error
	uuidClient, err = uuidgeneratorserver.NewClient("uuid_generator_server", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err.Error())
	}
}

func GetUUIDBounds(ctx context.Context, count, bizCode int64) ([]*server.UUIDBound, bool) {
	resp, err := uuidClient.GetUUIDBounds(ctx, &server.GetUUIDBoundsRequest{
		Count:   utils.Int64Ptr(count),
		BizCode: utils.Int64Ptr(bizCode),
	})
	if err != nil {
		fmt.Printf("GetUUIDBounds err:%v", err)
		return nil, false
	}

	if resp == nil || resp.Base == nil || resp.Base.GetCode() != 0 {
		fmt.Printf("GetUUIDBounds code resp:%v \n", resp)
		return nil, false
	}

	fmt.Printf("resp:%v", render.Render(resp.UuidBounds))
	return resp.UuidBounds, true
}
