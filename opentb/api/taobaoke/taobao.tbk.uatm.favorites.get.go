package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 枚举出淘宝客在淘宝联盟超级搜索，特色频道中，创建的选品库列表 */
type TbkUatmFavoritesGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 需要返回的字段列表，不能为空，字段名之间使用逗号分隔 */
func (r *TbkUatmFavoritesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，从1开始计数 */
func (r *TbkUatmFavoritesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 默认20，页大小，即每一页的活动个数 */
func (r *TbkUatmFavoritesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 默认值-1；选品库类型，1：普通选品组，2：高佣选品组，-1，同时输出所有类型的选品组 */
func (r *TbkUatmFavoritesGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

func (r *TbkUatmFavoritesGetRequest) GetResponse(accessToken string) (*TbkUatmFavoritesGetResponse, []byte, error) {
	var resp TbkUatmFavoritesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.favorites.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmFavoritesGetResponse struct {
	Results      []*TbkFavorites `json:"results"`
	TotalResults int             `json:"total_results"`
}

type TbkUatmFavoritesGetResponseResult struct {
	Response *TbkUatmFavoritesGetResponse `json:"tbk_uatm_favorites_get_response"`
}
