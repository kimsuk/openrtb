package openrtb

import "encoding/json"

type Pmp struct {
	Private int       `json:"private_auction,omitempty"` // 标识在Deal对象中指明的席位的竞拍合格标准， 0标识接受所有竞拍， 1标识竞拍受deals属性中描述的规则的限制
	Deals   []Deal    `json:"deals,omitempty"`           // 一组Deal对象， 用于传输适用于本次展示的交易信息
	Ext     Extension `json:"ext,omitempty"`             // 特定交易的OpenRTB协议的扩展信息占位符
}

// PMP Deal
type Deal struct {
	ID               string    `json:"id,omitempty"`          // 直接交易的唯一ID
	BidFloor         float64   `json:"bidfloor,omitempty"`    // 本次展示的最低竞价，以CPM为单位
	BidFloorCurrency string    `json:"bidfloorcur,omitempty"` // 使用ISO-4217码表指定的货币。 如果交易平台允许，这可能与竞价者返回的竞价货币类型不一致
	WSeat            []string  `json:"wseat,omitempty"`       // 允许参与本次交易竞价的买方席位白名单。 席位ID需要交易平台和竞拍者提前协商， 忽略本属性标示没有席位限制
	WAdvDomain       []string  `json:"wadomain,omitempty"`    // Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies no advertiser restrictions.
	AuctionType      int       `json:"at,omitempty"`          // 允许参与本次交易竞价的广告主域名列表（例如， advertiser.com). 忽略本属性标示没有广告主限制。
	Ext              Extension `json:"ext,omitempty"`         // 特定交易的OpenRTB协议的扩展信息占位符
	//Seats []string `json:"seats,omitempty"`
	//Type  int      `json:"type,omitempty"`
}

type jsonDeal Deal

// MarshalJSON custom marshalling with normalization
func (d *Deal) MarshalJSON() ([]byte, error) {
	d.normalize()
	return json.Marshal((*jsonDeal)(d))
}

// UnmarshalJSON custom unmarshalling with normalization
func (d *Deal) UnmarshalJSON(data []byte) error {
	var h jsonDeal
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*d = (Deal)(h)
	d.normalize()
	return nil
}

func (d *Deal) normalize() {
	if d.AuctionType == 0 {
		d.AuctionType = 2
	}
}
