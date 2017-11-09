package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 获取淘抢购的数据，淘客商品转淘客链接，非淘客商品输出普通链接 */
type TbkJuTqgGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 推广位id（推广位申请方式：http://club.alimama.com/read.php?spm=0.0.0.0.npQdST&tid=6306396&ds=1&page=1&toread=1） */
func (r *TbkJuTqgGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 最晚开团时间 */
func (r *TbkJuTqgGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 需返回的字段列表 */
func (r *TbkJuTqgGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认1，1~100 */
func (r *TbkJuTqgGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认40，1~40 */
func (r *TbkJuTqgGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 最早开团时间 */
func (r *TbkJuTqgGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TbkJuTqgGetRequest) GetResponse(accessToken string) (*TbkJuTqgGetResponse, []byte, error) {
	var resp TbkJuTqgGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.ju.tqg.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkJuTqgGetResponse struct {
	Results      []*Results `json:"results"`
	TotalResults int        `json:"total_results"`
}

type TbkJuTqgGetResponseResult struct {
	Response *TbkJuTqgGetResponse `json:"tbk_ju_tqg_get_response"`
}
