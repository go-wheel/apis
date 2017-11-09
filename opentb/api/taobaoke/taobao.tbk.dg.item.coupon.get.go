// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 好券清单API【导购】 */
type TbkDgItemCouponGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgItemCouponGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (r *TbkDgItemCouponGetRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过200；当指定类目或关键词的时候，则最多出100个结果） */
func (r *TbkDgItemCouponGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgItemCouponGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 1：PC，2：无线，默认：1 */
func (r *TbkDgItemCouponGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkDgItemCouponGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

func (r *TbkDgItemCouponGetRequest) GetResponse(accessToken string) (*TbkDgItemCouponGetResponse, []byte, error) {
	var resp TbkDgItemCouponGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.item.coupon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgItemCouponGetResponse struct {
	Results      []*TbkCoupon `json:"results"`
	TotalResults int          `json:"total_results"`
}

type TbkDgItemCouponGetResponseResult struct {
	Response *TbkDgItemCouponGetResponse `json:"tbk_dg_item_coupon_get_response"`
}
