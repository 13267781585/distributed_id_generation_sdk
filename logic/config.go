package logic

import (
	"time"

	"uuid_client/utils"
)

const (
	defaultQPSBuffer      int64         = 100
	defaultQPSWindow      int64         = 30
	defaultMinFetchCount  int64         = 1000
	defaultMaxFetchCount  int64         = 100000
	defaultFactor         int64         = 50
	defaultBufferSize     int64         = 5000
	defaultGetUUIDTimeout time.Duration = 2000 * time.Millisecond // 毫秒
)

type Config struct {
	BizCode        *int64
	Factor         *int64
	BufferSize     *int64
	QPSBuffer      *int64
	QPSWindow      *int64
	MinFetchCount  *int64
	MaxFetchCount  *int64
	GetUUIDTimeout time.Duration
}

func NewDefaultConfig() *Config {
	return &Config{
		QPSBuffer:      utils.Int64Ptr(defaultQPSBuffer),
		QPSWindow:      utils.Int64Ptr(defaultQPSWindow),
		MinFetchCount:  utils.Int64Ptr(defaultMinFetchCount),
		MaxFetchCount:  utils.Int64Ptr(defaultMaxFetchCount),
		Factor:         utils.Int64Ptr(defaultFactor),
		BufferSize:     utils.Int64Ptr(defaultBufferSize),
		GetUUIDTimeout: defaultGetUUIDTimeout,
	}
}

type Option func(*Config)

func WithQPSBuffer(qpsBuffer int64) Option {
	return func(config *Config) {
		config.QPSBuffer = utils.Int64Ptr(qpsBuffer)
	}
}

func WithMinFetchCount(minFetchCount int64) Option {
	return func(config *Config) {
		config.MinFetchCount = utils.Int64Ptr(minFetchCount)
	}
}

func WithBizCode(bizCode int64) Option {
	return func(config *Config) {
		config.BizCode = utils.Int64Ptr(bizCode)
	}
}
