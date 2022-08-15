package wechatunion

import (
	"context"
	"fmt"
	"time"
)

func (c *Client) GetAccessToken(ctx context.Context) string {
	if c.redisClient.Db == nil {
		return c.config.AccessToken
	}
	newCache := c.redisClient.NewSimpleStringCache(c.redisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinToken(ctx)
		return token.Result.AccessToken
	}
	return newCache.GetCache(ctx, c.getAccessTokenCacheKeyName())
}

func (c *Client) getAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_access_token:%v", c.getAppId())
}
