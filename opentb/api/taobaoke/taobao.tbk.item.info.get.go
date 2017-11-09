// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 淘宝客商品详情查询（简版） */
type TbkItemInfoGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 需返回的字段列表 */
func (r *TbkItemInfoGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品ID串，用,分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个 */
func (r *TbkItemInfoGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemInfoGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

func (r *TbkItemInfoGetRequest) GetResponse(accessToken string) (*TbkItemInfoGetResponse, []byte, error) {
	var resp TbkItemInfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.info.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemInfoGetResponse struct {
	Results []*NTbkItem `json:"results"`
}

type TbkItemInfoGetResponseResult struct {
	Response *TbkItemInfoGetResponse `json:"tbk_item_info_get_response"`
}
