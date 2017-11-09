package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 淘宝客商品关联推荐查询 */
type TbkItemRecommendGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 返回数量，默认20，最大值40 */
func (r *TbkItemRecommendGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 需返回的字段列表 */
func (r *TbkItemRecommendGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品Id */
func (r *TbkItemRecommendGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemRecommendGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

func (r *TbkItemRecommendGetRequest) GetResponse(accessToken string) (*TbkItemRecommendGetResponse, []byte, error) {
	var resp TbkItemRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemRecommendGetResponse struct {
	Results []*NTbkItem `json:"results"`
}

type TbkItemRecommendGetResponseResult struct {
	Response *TbkItemRecommendGetResponse `json:"tbk_item_recommend_get_response"`
}
