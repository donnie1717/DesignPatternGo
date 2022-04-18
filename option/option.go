package option

import "time"

const (
	defaultTimeout = 0
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// 可选参数
type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(option *options)
}

type optionFuc func(*options)

func (f optionFuc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFuc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	var fuc optionFuc = func(o *options) {
		o.caching = cache
	}
	return fuc
}

func Connect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}
	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}
