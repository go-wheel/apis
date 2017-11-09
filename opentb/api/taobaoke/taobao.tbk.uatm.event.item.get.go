package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 通过指定定向招商活动id，获取该活动id下的宝贝信息；
宝贝信息中包含了适用于定向招商活动的高佣金淘客点击串; 注意，只能获取已经开始的定向招商id下面的宝贝信息，当天新开始的定向招商活动在凌晨2点生效； */
type TbkUatmEventItemGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 推广位id，需要在淘宝联盟后台创建；且属于appkey对应的备案媒体id（siteid），如何获取adzoneid，请参考：http://club.alimama.com/read-htm-tid-6333967.html?spm=0.0.0.0.msZnx5 */
func (r *TbkUatmEventItemGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 招商活动id */
func (r *TbkUatmEventItemGetRequest) SetEventId(value string) {
	r.SetValue("event_id", value)
}

/* 需要输出则字段列表，逗号分隔 */
func (r *TbkUatmEventItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认：１，从1开始计数 */
func (r *TbkUatmEventItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkUatmEventItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkUatmEventItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道 */
func (r *TbkUatmEventItemGetRequest) SetUnid(value string) {
	r.SetValue("unid", value)
}

func (r *TbkUatmEventItemGetRequest) GetResponse(accessToken string) (*TbkUatmEventItemGetResponse, []byte, error) {
	var resp TbkUatmEventItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.event.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmEventItemGetResponse struct {
	Results      []*UatmTbkItem `json:"results"`
	TotalResults int            `json:"total_results"`
}

type TbkUatmEventItemGetResponseResult struct {
	Response *TbkUatmEventItemGetResponse `json:"tbk_uatm_event_item_get_response"`
}
