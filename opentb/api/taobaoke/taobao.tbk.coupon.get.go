// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 阿里妈妈推广券信息查询 */
type TbkCouponGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 带券ID与商品ID的加密串 */
func (r *TbkCouponGetRequest) SetMe(value string) {
	r.SetValue("me", value)
}

func (r *TbkCouponGetRequest) GetResponse(accessToken string) (*TbkCouponGetResponse, []byte, error) {
	var resp TbkCouponGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.coupon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkCouponGetResponse struct {
	Data *MapData `json:"data"`
}

type TbkCouponGetResponseResult struct {
	Response *TbkCouponGetResponse `json:"tbk_coupon_get_response"`
}
