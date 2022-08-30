package wechatunion

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.appId = appId
	c.config.appSecret = appSecret
	return c
}

// ConfigPid 配置
func (c *Client) ConfigPid(pid string) *Client {
	c.config.pid = pid
	return c
}
