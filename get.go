package wechatunion

import (
	"context"
	"go.dtapp.net/golog"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetPid() string {
	return c.config.pid
}

func (c *Client) getAccessToken(ctx context.Context) string {
	c.config.accessToken = c.GetAccessToken(ctx)
	return c.config.accessToken
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
