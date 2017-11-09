package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
现阶段只支持短连接。 */
type TbkSpreadGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 请求列表，内部包含多个url */
func (r *TbkSpreadGetRequest) SetRequests(value string) {
	r.SetValue("requests", value)
}

func (r *TbkSpreadGetRequest) GetResponse(accessToken string) (*TbkSpreadGetResponse, []byte, error) {
	var resp TbkSpreadGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.spread.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkSpreadGetResponse struct {
	Results      []*TbkSpread `json:"results"`
	TotalResults int          `json:"total_results"`
}

type TbkSpreadGetResponseResult struct {
	Response *TbkSpreadGetResponse `json:"tbk_spread_get_response"`
}
