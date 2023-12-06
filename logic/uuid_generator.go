package logic

import (
	"uuid_client/tools"
	"uuid_client/utils"
)

type UUIDGenerator struct {
	config   *Config
	qpsCount *tools.QPSCounter
	uuidChan chan int64
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

	utils.GoFunc(generator.init)
	return generator
}

func (u *UUIDGenerator) init() {

}

func (u *UUIDGenerator) GetUUID() (int64, error) {
	u.qpsCount.AddCount(1)
	return 0, nil
}

func (u *UUIDGenerator) MGetUUIDs(count int64) ([]int64, error) {
	u.qpsCount.AddCount(count)
	return []int64{}, nil
}
