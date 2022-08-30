package wechatunion

import (
	"context"
	"time"
)

func (c *Client) GetAccessToken(ctx context.Context) string {
	if c.cache.redisClient.Db == nil {
		return c.config.accessToken
	}
	newCache := c.cache.redisClient.NewSimpleStringCache(c.cache.redisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinToken(ctx)
		return token.Result.AccessToken
	}
	return newCache.GetCache(ctx, c.getAccessTokenCacheKeyName())
}

func (c *Client) getAccessTokenCacheKeyName() string {
	return c.cache.wechatAccessTokenPrefix + c.GetAppId()
}
