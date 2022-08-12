package wechatunion

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type CgiBinTokenResult struct {
	Result CgiBinTokenResponse // 结果
	Byte   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newCgiBinTokenResult(result CgiBinTokenResponse, byte []byte, http gorequest.Response, err error) *CgiBinTokenResult {
	return &CgiBinTokenResult{Result: result, Byte: byte, Http: http, Err: err}
}

// CgiBinToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) CgiBinToken(ctx context.Context) *CgiBinTokenResult {
	// 请求
	request, err := c.request(ctx, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", c.getAppId(), c.getAppSecret()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinTokenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinTokenResult(response, request.ResponseBody, request, err)
}
