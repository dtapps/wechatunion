package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PromotionAddResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Pid     string `json:"pid"`     // 推广位ID，PID
}

type PromotionAddResult struct {
	Result PromotionAddResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func newPromotionAddResult(result PromotionAddResponse, body []byte, http gorequest.Response, err error) *PromotionAddResult {
	return &PromotionAddResult{Result: result, Body: body, Http: http, Err: err}
}

// PromotionAdd 添加推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_1-%E6%B7%BB%E5%8A%A0%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionAdd(promotionSourceName string) *PromotionAddResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("promotionSourceName", promotionSourceName) // 推广位名称
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/promoter/promotion/add?access_token%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response PromotionAddResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPromotionAddResult(response, request.ResponseBody, request, err)
}
