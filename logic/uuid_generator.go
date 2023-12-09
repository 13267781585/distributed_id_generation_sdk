package logic

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
	"uuid_client/rpc"
	"uuid_client/tools"
	"uuid_client/utils"
)

type UUIDGenerator struct {
	config   *Config
	qpsCount *tools.QPSCounter
	uuidChan chan int64
	isClose  int32
}

func NewUUIDGenerator(opts ...Option) *UUIDGenerator {
	config := NewDefaultConfig()
	for _, opt := range opts {
		opt(config)
	}

	if config.BizCode == nil {
		panic("invalid biz code")
	}

	generator := &UUIDGenerator{
		qpsCount: tools.NewQPSCounter(*config.QPSWindow),
		config:   config,
		uuidChan: make(chan int64, *config.BufferSize),
	}

	utils.GoFunc(generator.cycleCheck)
	return generator
}

func (u *UUIDGenerator) Close() {
	atomic.StoreInt32(&u.isClose, 1)
}

func (u *UUIDGenerator) cycleCheck() {
	for {
		if atomic.LoadInt32(&u.isClose) == 1 {
			fmt.Print("cycle check return")
			return
		}

		dataSize := int64(len(u.uuidChan))
		if dataSize < utils.Int64(u.config.BufferSize)*utils.Int64(u.config.Factor)/100 {
			startTime := time.Now().UnixMilli()
			u.fetchUUIDsFromServer()
			endTime := time.Now().UnixMilli()

			spendTime := endTime - startTime
			if spendTime > 1000 {
				fmt.Printf("fetch uuids spend time:%v \n", spendTime)
			}
		}

		time.Sleep(500 * time.Microsecond)
	}
}

func (u *UUIDGenerator) fetchUUIDsFromServer() {
	fetchCount := u.qpsCount.SumCount() * utils.Int64(u.config.QPSBuffer)
	if fetchCount > utils.Int64(u.config.MaxFetchCount) {
		fetchCount = utils.Int64(u.config.MaxFetchCount)
	} else if fetchCount < utils.Int64(u.config.MinFetchCount) {
		fetchCount = utils.Int64(u.config.MinFetchCount)
	}

	bounds, ok := rpc.GetUUIDBounds(context.Background(), fetchCount, utils.Int64(u.config.BizCode))
	if !ok {
		fmt.Println("rpc GetUUIDBounds get error")
		return
	}

	for _, bound := range bounds {
		start := utils.Int64(bound.Start)
		for ; start < utils.Int64(bound.End); start++ {
			u.uuidChan <- start
		}
	}
}

func (u *UUIDGenerator) GetUUID() (int64, bool) {
	u.qpsCount.AddCount(1)
	select {
	case nextID := <-u.uuidChan:
		return nextID, true
	case <-time.After(u.config.GetUUIDTimeout):
		fmt.Println("get uuid time out")
		return 0, false
	}
}

func (u *UUIDGenerator) MGetUUIDs(count int64) ([]int64, bool) {
	u.qpsCount.AddCount(count)
	uuids := make([]int64, count)
	for i := int64(0); i < count; i++ {
		select {
		case nextID := <-u.uuidChan:
			uuids[i] = nextID
		case <-time.After(u.config.GetUUIDTimeout):
			fmt.Println("mget uuid time out")
			return uuids, false
		}
	}
	return uuids, true
}
