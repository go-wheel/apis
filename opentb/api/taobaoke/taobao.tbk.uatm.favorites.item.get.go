// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 指定选品库id，获取该选品库的宝贝信息 */
type TbkUatmFavoritesItemGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 推广位id，需要在淘宝联盟后台创建；且属于appkey备案的媒体id（siteid），如何获取adzoneid，请参考，http://club.alimama.com/read-htm-tid-6333967.html?spm=0.0.0.0.msZnx5 */
func (r *TbkUatmFavoritesItemGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 选品库的id */
func (r *TbkUatmFavoritesItemGetRequest) SetFavoritesId(value string) {
	r.SetValue("favorites_id", value)
}

/* 需要输出则字段列表，逗号分隔 */
func (r *TbkUatmFavoritesItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认：1，从1开始计数 */
func (r *TbkUatmFavoritesItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkUatmFavoritesItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkUatmFavoritesItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道 */
func (r *TbkUatmFavoritesItemGetRequest) SetUnid(value string) {
	r.SetValue("unid", value)
}

func (r *TbkUatmFavoritesItemGetRequest) GetResponse(accessToken string) (*TbkUatmFavoritesItemGetResponse, []byte, error) {
	var resp TbkUatmFavoritesItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.favorites.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmFavoritesItemGetResponse struct {
	Results      []*UatmTbkItem `json:"results"`
	TotalResults int            `json:"total_results"`
}

type TbkUatmFavoritesItemGetResponseResult struct {
	Response *TbkUatmFavoritesItemGetResponse `json:"tbk_uatm_favorites_item_get_response"`
}
