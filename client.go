package wechatunion

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// 缓存前缀
// wechat_union:wechat_access_token:
type redisCachePrefixFun func() (wechatAccessToken string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId               string              // 小程序唯一凭证，即 appId
	AppSecret           string              // 小程序唯一凭证密钥，即 appSecret
	Pid                 string              // 推广位PID
	RedisClient         *dorm.RedisClient   // 缓存数据库
	GormClientFun       gormClientFun       // 日志配置
	MongoClientFun      mongoClientFun      // 日志配置
	Debug               bool                // 日志开关
	ZapLog              *golog.ZapLog       // 日志服务
	RedisCachePrefixFun redisCachePrefixFun // 缓存前缀
	CurrentIp           string              // 当前ip
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	currentIp     string         // 当前ip
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
		gorm           bool              // 日志开关
		gormClient     *dorm.GormClient  // 日志数据库
		logGormClient  *golog.ApiClient  // 日志服务
		mongo          bool              // 日志开关
		mongoClient    *dorm.MongoClient // 日志数据库
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.zapLog = config.ZapLog

	c.currentIp = config.CurrentIp

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.pid = config.Pid

	c.requestClient = gorequest.NewHttp()

	gormClient := config.GormClientFun()
	if gormClient != nil && gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(&golog.ApiGormClientConfig{
			GormClientFun: func() (*dorm.GormClient, string) {
				return gormClient, logTable
			},
			Debug:     config.Debug,
			ZapLog:    c.zapLog,
			CurrentIp: c.currentIp,
		})
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
		c.log.gormClient = gormClient
	}

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient != nil && mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(&golog.ApiMongoClientConfig{
			MongoClientFun: func() (*dorm.MongoClient, string, string) {
				return mongoClient, databaseName, logTable
			},
			Debug:     config.Debug,
			ZapLog:    c.zapLog,
			CurrentIp: c.currentIp,
		})
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
		c.log.mongoClient = mongoClient
	}

	c.cache.redisClient = config.RedisClient

	c.cache.wechatAccessTokenPrefix = config.RedisCachePrefixFun()
	if c.cache.wechatAccessTokenPrefix == "" {
		return nil, redisCachePrefixNoConfig
	}

	return c, nil
}
