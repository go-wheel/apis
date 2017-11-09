package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 淘宝客店铺查询 */
type TbkShopGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 累计推广商品上限 */
func (r *TbkShopGetRequest) SetEndAuctionCount(value string) {
	r.SetValue("end_auction_count", value)
}

/* 淘客佣金比率上限，1~10000 */
func (r *TbkShopGetRequest) SetEndCommissionRate(value string) {
	r.SetValue("end_commission_rate", value)
}

/* 信用等级上限，1~20 */
func (r *TbkShopGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 店铺商品总数上限 */
func (r *TbkShopGetRequest) SetEndTotalAction(value string) {
	r.SetValue("end_total_action", value)
}

/* 需返回的字段列表 */
func (r *TbkShopGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性 */
func (r *TbkShopGetRequest) SetIsTmall(value string) {
	r.SetValue("is_tmall", value)
}

/* 第几页，默认1，1~100 */
func (r *TbkShopGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkShopGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkShopGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkShopGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction） */
func (r *TbkShopGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 累计推广商品下限 */
func (r *TbkShopGetRequest) SetStartAuctionCount(value string) {
	r.SetValue("start_auction_count", value)
}

/* 淘客佣金比率下限，1~10000 */
func (r *TbkShopGetRequest) SetStartCommissionRate(value string) {
	r.SetValue("start_commission_rate", value)
}

/* 信用等级下限，1~20 */
func (r *TbkShopGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 店铺商品总数下限 */
func (r *TbkShopGetRequest) SetStartTotalAction(value string) {
	r.SetValue("start_total_action", value)
}

func (r *TbkShopGetRequest) GetResponse(accessToken string) (*TbkShopGetResponse, []byte, error) {
	var resp TbkShopGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shop.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopGetResponse struct {
	Results      []*NTbkShop `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkShopGetResponseResult struct {
	Response *TbkShopGetResponse `json:"tbk_shop_get_response"`
}
