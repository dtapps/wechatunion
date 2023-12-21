package wechatunion

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PromotionListResponse struct {
	Errcode             int    `json:"errcode"` // 错误码
	Errmsg              string `json:"errmsg"`  // 错误信息
	PromotionSourceList []struct {
		PromotionSourceName string `json:"promotionSourceName"` // 推广位名称
		PromotionSourcePid  string `json:"promotionSourcePid"`  // 推广位ID，PID
		Status              string `json:"status"`              // 状态
		PidId               string `json:"pidId"`
	} `json:"promotionSourceList"` // 推广位数据
	Total           int `json:"total"`           // 推广位总数
	PromotionMaxCnt int `json:"promotionMaxCnt"` // 允许创建的推广位最大数量
}

type PromotionListResult struct {
	Result PromotionListResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newPromotionListResult(result PromotionListResponse, body []byte, http gorequest.Response) *PromotionListResult {
	return &PromotionListResult{Result: result, Body: body, Http: http}
}

// PromotionList 获取推广位列表
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_4-%E8%8E%B7%E5%8F%96%E6%8E%A8%E5%B9%BF%E4%BD%8D%E5%88%97%E8%A1%A8
func (c *Client) PromotionList(ctx context.Context, start int, limit int, notMustParams ...gorequest.Params) (*PromotionListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("start", start) // 偏移
	params.Set("limit", limit) // 每页条数
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/promoter/promotion/list?access_token%s", c.getAccessToken(ctx)), params, http.MethodGet)
	if err != nil {
		return newPromotionListResult(PromotionListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PromotionListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPromotionListResult(response, request.ResponseBody, request), err
}
