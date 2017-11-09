// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 淘宝客店铺关联推荐查询 */
type TbkShopRecommendGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 返回数量，默认20，最大值40 */
func (r *TbkShopRecommendGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 需返回的字段列表 */
func (r *TbkShopRecommendGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkShopRecommendGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 卖家Id */
func (r *TbkShopRecommendGetRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}

func (r *TbkShopRecommendGetRequest) GetResponse(accessToken string) (*TbkShopRecommendGetResponse, []byte, error) {
	var resp TbkShopRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shop.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopRecommendGetResponse struct {
	Results []*NTbkShop `json:"results"`
}

type TbkShopRecommendGetResponseResult struct {
	Response *TbkShopRecommendGetResponse `json:"tbk_shop_recommend_get_response"`
}
