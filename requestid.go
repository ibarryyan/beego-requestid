package beego_requestid

import (
	"github.com/beego/beego"
	"github.com/beego/beego/context"
	"github.com/google/uuid"
)

const DefaultHeaderReqIdKey = "X-Request-Id"

type Option func(config *Config)

type GenRequestIdFunc func() string

type Config struct {
	genRequestIdFunc               GenRequestIdFunc
	headerReqIdKey, customReqIdKey string
}

func NewFilter(opts ...Option) beego.FilterFunc {
	cnf := &Config{
		genRequestIdFunc: DefaultGenRequestIdFunc,
		headerReqIdKey:   DefaultHeaderReqIdKey,
	}

	for _, opt := range opts {
		opt(cnf)
	}

	return func(c *context.Context) {
		reqId := c.Request.Header.Get(cnf.headerReqIdKey)
		if reqId == "" {
			reqId = cnf.genRequestIdFunc()
			c.Request.Header.Add(cnf.headerReqIdKey, reqId)
		}
		if cnf.customReqIdKey != "" {
			c.Input.SetData(cnf.customReqIdKey, reqId)
		}
	}
}

func WithGenRequestIdFunc(genFunc GenRequestIdFunc) Option {
	return func(config *Config) {
		config.genRequestIdFunc = genFunc
	}
}

func WithHeaderReqIdKey(key string) Option {
	return func(config *Config) {
		config.headerReqIdKey = key
	}
}

func WithCustomReqIdKey(key string) Option {
	return func(config *Config) {
		config.customReqIdKey = key
	}
}

func DefaultGenRequestIdFunc() string {
	return uuid.NewString()
}
