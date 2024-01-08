package wechatunion

import (
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/golog"
)

// 缓存前缀
// wechat_union:wechat_access_token:
type redisCachePrefixFun func() (wechatAccessToken string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId               string              `json:"app_id"` // 小程序唯一凭证，即 appId
	AppSecret           string              // 小程序唯一凭证密钥，即 appSecret
	Pid                 string              // 推广位PID
	RedisClient         *redis.Client       // 缓存数据库
	RedisCachePrefixFun redisCachePrefixFun // 缓存前缀
}

// Client 实例
type Client struct {
	config struct {
		appId       string // 小程序唯一凭证，即 appId
		appSecret   string // 小程序唯一凭证密钥，即 appSecret
		accessToken string // 接口调用凭证
		pid         string // 推广位PID
	}
	cache struct {
		redisClient             *redis.Client // 缓存数据库
		wechatAccessTokenPrefix string        // AccessToken
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.pid = config.Pid

	c.cache.redisClient = config.RedisClient

	c.cache.wechatAccessTokenPrefix = config.RedisCachePrefixFun()
	if c.cache.wechatAccessTokenPrefix == "" {
		return nil, redisCachePrefixNoConfig
	}

	return c, nil
}
