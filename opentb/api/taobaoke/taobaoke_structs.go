// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

const VersionNo = "20171020"

/* 淘宝客商品 */
type NTbkItem struct {
	CatLeafName     string   `json:"cat_leaf_name"`      // 叶子类目名称
	CatName         string   `json:"cat_name"`           // 一级类目名称
	ClickUrl        string   `json:"click_url"`          // 淘客地址
	CommissionRate  string   `json:"commission_rate"`    // 佣金比例
	CouponAmount    string   `json:"coupon_amount"`      // couponAmount
	CouponPrice     string   `json:"coupon_price"`       // 优惠券额度
	CouponStartFee  string   `json:"coupon_start_fee"`   // 优惠券使用门槛金额
	ItemUrl         string   `json:"item_url"`           // 商品地址
	Nick            string   `json:"nick"`               // 卖家昵称
	NumIid          int      `json:"num_iid"`            // 商品ID
	PictUrl         string   `json:"pict_url"`           // 商品主图
	Provcity        string   `json:"provcity"`           // 宝贝所在地
	ReservePrice    string   `json:"reserve_price"`      // 商品一口价格
	SellerId        int      `json:"seller_id"`          // 卖家id
	ShopTitle       string   `json:"shop_title"`         //
	SmallImages     []string `json:"small_images"`       // 商品小图列表
	Title           string   `json:"title"`              // 商品标题
	TkRate          string   `json:"tk_rate"`            //
	UserType        int      `json:"user_type"`          // 卖家类型，0表示集市，1表示商城
	Volume          int      `json:"volume"`             // 30天销量
	ZkFinalPrice    string   `json:"zk_final_price"`     // 商品折扣价格
	ZkFinalPriceWap string   `json:"zk_final_price_wap"` //

}

/* 淘宝客选品库 */
type TbkFavorites struct {
	FavoritesId    int    `json:"favorites_id"`    // 选品库id
	FavoritesTitle string `json:"favorites_title"` // 选品组名称
	Type           int    `json:"type"`            // 选品库类型，1：普通类型，2高佣金类型

}

/* 淘宝客店铺 */
type NTbkShop struct {
	ClickUrl   string `json:"click_url"`   // 淘客地址
	PictUrl    string `json:"pict_url"`    // 店标图片
	SellerNick string `json:"seller_nick"` // 卖家昵称
	ShopTitle  string `json:"shop_title"`  // 店铺名称
	ShopType   string `json:"shop_type"`   // 店铺类型，B：天猫，C：淘宝
	ShopUrl    string `json:"shop_url"`    // 店铺地址
	UserId     int    `json:"user_id"`     // 卖家ID

}

/* 淘客定向招商活动基本信息 */
type TbkEvent struct {
	EndTime    string `json:"end_time"`    // 定向招商活动结束开始时间
	EventId    int    `json:"event_id"`    // 淘宝联盟定向招商活动id
	EventTitle string `json:"event_title"` // 淘宝联盟定向招商活动名称
	StartTime  string `json:"start_time"`  // 定向招商活动结束开始时间

}

/* 传播形式对象列表 */
type TbkSpread struct {
	Content string `json:"content"` // 传播形式, 目前只支持短链接
	ErrMsg  string `json:"err_msg"` // 调用错误信息；由于是批量接口，请重点关注每条请求返回的结果，如果非OK，则说明该结果对应的content不正常，请酌情处理;

}

/* 淘宝联盟选品和招商宝贝信息 */
type UatmTbkItem struct {
	Category          int      `json:"category"`            // 后台一级类目
	ClickUrl          string   `json:"click_url"`           // 淘客地址
	CommissionRate    string   `json:"commission_rate"`     // 佣金比率(%)
	CouponClickUrl    string   `json:"coupon_click_url"`    // 商品优惠券推广链接
	CouponEndTime     string   `json:"coupon_end_time"`     // 优惠券结束时间
	CouponInfo        string   `json:"coupon_info"`         // 优惠券面额
	CouponRemainCount int      `json:"coupon_remain_count"` // 优惠券剩余量
	CouponStartTime   string   `json:"coupon_start_time"`   // 优惠券开始时间
	CouponTotalCount  int      `json:"coupon_total_count"`  // 优惠券总量
	EventEndTime      string   `json:"event_end_time"`      // 招行活动的结束时间；如果该宝贝取自普通的选品组，则取值为1970-01-01 00:00:00
	EventStartTime    string   `json:"event_start_time"`    // 招商活动开始时间；如果该宝贝取自普通选品组，则取值为1970-01-01 00:00:00；
	ItemUrl           string   `json:"item_url"`            // 商品地址
	Nick              string   `json:"nick"`                // 卖家昵称
	NumIid            int      `json:"num_iid"`             // 商品ID
	PictUrl           string   `json:"pict_url"`            // 商品主图
	Provcity          string   `json:"provcity"`            // 宝贝所在地
	ReservePrice      string   `json:"reserve_price"`       // 商品一口价格
	SellerId          int      `json:"seller_id"`           // 卖家id
	ShopTitle         string   `json:"shop_title"`          //
	SmallImages       []string `json:"small_images"`        // 商品小图列表
	Status            int      `json:"status"`              // 宝贝状态，0失效，1有效；注：失效可能是宝贝已经下线或者是被处罚不能在进行推广
	Title             string   `json:"title"`               // 商品标题
	TkRate            string   `json:"tk_rate"`             // 收入比例，举例，取值为20.00，表示比例20.00%
	Type              int      `json:"type"`                // 宝贝类型：1 普通商品； 2 鹊桥高佣金商品；3 定向招商商品；4 营销计划商品;
	UserType          int      `json:"user_type"`           // 卖家类型，0表示集市，1表示商城
	Volume            int      `json:"volume"`              // 30天销量
	ZkFinalPrice      string   `json:"zk_final_price"`      // 商品折扣价格
	ZkFinalPriceWap   string   `json:"zk_final_price_wap"`  // 无线折扣价，即宝贝在无线上的实际售卖价格。

}

/* data */
type MapData struct {
	Model string `json:"model"` // password

}

/* TbkCoupon */
type TbkCoupon struct {
	Category          int      `json:"category"`            // 后台一级类目
	CommissionRate    string   `json:"commission_rate"`     // 佣金比率(%)
	CouponClickUrl    string   `json:"coupon_click_url"`    // 商品优惠券推广链接
	CouponEndTime     string   `json:"coupon_end_time"`     // 优惠券结束时间
	CouponInfo        string   `json:"coupon_info"`         // 优惠券面额
	CouponRemainCount int      `json:"coupon_remain_count"` // 优惠券剩余量
	CouponStartTime   string   `json:"coupon_start_time"`   // 优惠券开始时间
	CouponTotalCount  int      `json:"coupon_total_count"`  // 优惠券总量
	ItemDescription   string   `json:"item_description"`    // 宝贝描述（推荐理由）
	ItemUrl           string   `json:"item_url"`            // 商品详情页链接地址
	Nick              string   `json:"nick"`                // 卖家昵称
	NumIid            int      `json:"num_iid"`             // itemId
	PictUrl           string   `json:"pict_url"`            // 商品主图
	SellerId          int      `json:"seller_id"`           // 卖家id
	ShopTitle         string   `json:"shop_title"`          // 店铺名称
	SmallImages       []string `json:"small_images"`        // 商品小图列表
	Title             string   `json:"title"`               // 商品标题
	UserType          int      `json:"user_type"`           // 卖家类型，0表示集市，1表示商城
	Volume            int      `json:"volume"`              // 30天销量
	ZkFinalPrice      string   `json:"zk_final_price"`      // 折扣价

}

/* 淘抢购对象 */
type Results struct {
	CategoryName string `json:"category_name"`  // 类目名称
	ClickUrl     string `json:"click_url"`      // 商品链接（是淘客商品返回淘客链接，非淘客商品返回普通h5链接）
	EndTime      string `json:"end_time"`       // 结束时间
	NumIid       int    `json:"num_iid"`        // 商品ID
	PicUrl       string `json:"pic_url"`        // 商品主图
	ReservePrice string `json:"reserve_price"`  // 商品原价
	SoldNum      int    `json:"sold_num"`       // 已抢购数量
	StartTime    string `json:"start_time"`     // 开团时间
	Title        string `json:"title"`          // 商品标题
	TotalAmount  int    `json:"total_amount"`   // 总库存
	ZkFinalPrice string `json:"zk_final_price"` // 淘抢购活动价

}
