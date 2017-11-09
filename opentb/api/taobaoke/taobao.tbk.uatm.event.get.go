package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 枚举指定淘客自己发起的，*正在进行中的*定向招商的活动列表；每天新开始的定向招商活动，在凌晨2点后生效；即凌晨2点后可以枚举到当天开始的定向招商活动列表；时间过早不能取到当天开始的定向招商活动； */
type TbkUatmEventGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 需要返回的字段列表，不能为空，字段名之间使用逗号分隔 */
func (r *TbkUatmEventGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 默认1，第几页，从1开始计数 */
func (r *TbkUatmEventGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 默认20,  页大小，即每一页的活动个数 */
func (r *TbkUatmEventGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *TbkUatmEventGetRequest) GetResponse(accessToken string) (*TbkUatmEventGetResponse, []byte, error) {
	var resp TbkUatmEventGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.event.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmEventGetResponse struct {
	Results      []*TbkEvent `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkUatmEventGetResponseResult struct {
	Response *TbkUatmEventGetResponse `json:"tbk_uatm_event_get_response"`
}
