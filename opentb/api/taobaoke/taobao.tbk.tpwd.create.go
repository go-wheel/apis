// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 提供淘客生成淘口令接口，淘客提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播 */
type TbkTpwdCreateRequest struct {
	opentb.TaobaoMethodRequest
}

/* 扩展字段JSON格式 */
func (r *TbkTpwdCreateRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 口令弹框logoURL */
func (r *TbkTpwdCreateRequest) SetLogo(value string) {
	r.SetValue("logo", value)
}

/* 口令弹框内容 */
func (r *TbkTpwdCreateRequest) SetText(value string) {
	r.SetValue("text", value)
}

/* 口令跳转目标页 */
func (r *TbkTpwdCreateRequest) SetUrl(value string) {
	r.SetValue("url", value)
}

/* 生成口令的淘宝用户ID */
func (r *TbkTpwdCreateRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}

func (r *TbkTpwdCreateRequest) GetResponse(accessToken string) (*TbkTpwdCreateResponse, []byte, error) {
	var resp TbkTpwdCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.tpwd.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkTpwdCreateResponse struct {
	Data *MapData `json:"data"`
}

type TbkTpwdCreateResponseResult struct {
	Response *TbkTpwdCreateResponse `json:"tbk_tpwd_create_response"`
}
