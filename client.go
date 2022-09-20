package wechatunion

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// 缓存前缀
// wechat_union:wechat_access_token:
type redisCachePrefixFun func() (wechatAccessToken string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId               string              // 小程序唯一凭证，即 appId
	AppSecret           string              // 小程序唯一凭证密钥，即 appSecret
	Pid                 string              // 推广位PID
	RedisClient         *dorm.RedisClient   // 缓存数据库
	ApiGormClientFun    golog.ApiClientFun  // 日志配置
	Debug               bool                // 日志开关
	ZapLog              *golog.ZapLog       // 日志服务
	RedisCachePrefixFun redisCachePrefixFun // 缓存前缀
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	config        struct {
		appId       string // 小程序唯一凭证，即 appId
		appSecret   string // 小程序唯一凭证密钥，即 appSecret
		accessToken string // 接口调用凭证
		pid         string // 推广位PID
	}
	cache struct {
		redisClient             *dorm.RedisClient // 缓存数据库
		wechatAccessTokenPrefix string            // AccessToken
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.pid = config.Pid

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.status = true
	}

	c.cache.redisClient = config.RedisClient

	c.cache.wechatAccessTokenPrefix = config.RedisCachePrefixFun()
	if c.cache.wechatAccessTokenPrefix == "" {
		return nil, redisCachePrefixNoConfig
	}

	return c, nil
}
