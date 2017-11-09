package taobaoke

import (
	"github.com/go-wheel/apis/opentb"
)

/* 淘宝客商品查询 */
type TbkItemGetRequest struct {
	opentb.TaobaoMethodRequest
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (r *TbkItemGetRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 折扣价范围上限，单位：元 */
func (r *TbkItemGetRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 淘客佣金比率下限，如：1234表示12.34% */
func (r *TbkItemGetRequest) SetEndTkRate(value string) {
	r.SetValue("end_tk_rate", value)
}

/* 需返回的字段列表 */
func (r *TbkItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemGetRequest) SetIsOverseas(value string) {
	r.SetValue("is_overseas", value)
}

/* 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemGetRequest) SetIsTmall(value string) {
	r.SetValue("is_tmall", value)
}

/* 所在地 */
func (r *TbkItemGetRequest) SetItemloc(value string) {
	r.SetValue("itemloc", value)
}

/* 第几页，默认：１ */
func (r *TbkItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkItemGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi） */
func (r *TbkItemGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 折扣价范围下限，单位：元 */
func (r *TbkItemGetRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}

/* 淘客佣金比率上限，如：1234表示12.34% */
func (r *TbkItemGetRequest) SetStartTkRate(value string) {
	r.SetValue("start_tk_rate", value)
}

func (r *TbkItemGetRequest) GetResponse(accessToken string) (*TbkItemGetResponse, []byte, error) {
	var resp TbkItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemGetResponse struct {
	Results      []*NTbkItem `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkItemGetResponseResult struct {
	Response *TbkItemGetResponse `json:"tbk_item_get_response"`
}
