package wechatunion

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Byte   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newGetCallBackIpResult(result GetCallBackIpResponse, byte []byte, http gorequest.Response) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Byte: byte, Http: http}
}

// GetCallBackIp 获取微信callback IP地址
// callback IP即微信调用开发者服务器所使用的出口IP。
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (c *Client) GetCallBackIp(ctx context.Context, notMustParams ...gorequest.Params) (*GetCallBackIpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", c.getAccessToken(ctx)), params, http.MethodGet)
	if err != nil {
		return newGetCallBackIpResult(GetCallBackIpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetCallBackIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetCallBackIpResult(response, request.ResponseBody, request), err
}
