package wechatunion

import (
	"context"
	"time"
)

func (c *Client) GetAccessToken(ctx context.Context) string {
	if c.cache.redisClient == nil {
		return c.config.accessToken
	}
	result, _ := c.cache.redisClient.Get(ctx, c.getAccessTokenCacheKeyName()).Result()
	if result != "" {
		return result
	}
	token, _ := c.CgiBinToken(ctx)
	c.cache.redisClient.Set(ctx, c.getAccessTokenCacheKeyName(), token.Result.AccessToken, time.Second*7000)
	return token.Result.AccessToken
}

func (c *Client) getAccessTokenCacheKeyName() string {
	return c.cache.wechatAccessTokenPrefix + c.GetAppId()
}
