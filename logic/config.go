package logic

const (
	defaultQPSBuffer     int64 = 100
	defaultQPSWindow     int64 = 30
	defaultMinFetchCount int64 = 1000
	defaultMaxFetchCount int64 = 100000
	defaultFactor        int64 = 50
	defaultBufferSize          = 5000
)

type Config struct {
	BizCode       *int64
	Factor        *int64
	BufferSize    *int64
	QPSBuffer     *int64
	QPSWindow     *int64
	MinFetchCount *int64
	MaxFetchCount *int64
}

func NewDefaultConfig() *Config {
	return &Config{
		QPSBuffer:     utils.Int64Ptr(defaultQPSBuffer),
		QPSWindow:     utils.Int64Ptr(defaultQPSWindow),
		MinFetchCount: utils.Int64Ptr(defaultMinFetchCount),
		MaxFetchCount: utils.Int64Ptr(defaultMaxFetchCount),
		Factor:        utils.Int64Ptr(defaultFactor),
		BufferSize:    utils.Int64Ptr(defaultBufferSize),
	}
}

type Option func(*Config)

func WithBizCode(bizCode int64) Option {
	return func(config *Config) {
		config.BizCode = utils.Int64Ptr(bizCode)
	}
}
