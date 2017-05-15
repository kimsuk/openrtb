package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidRespNoID       = errors.New("openrtb: response missing ID")
	ErrInvalidRespNoSeatBids = errors.New("openrtb: response missing seatbids")
)

//顶级的竞价形影对象（即没有名称的最外层JSON对象）。 id属性是竞价请求的ID,用于记录日志。
//同样的bidid是一个可选的响应追踪标识， 如果指定了该标识， 如果之后竞拍者胜出了，则在胜出通知子流程中，需要将该标识填充进去。
//至少一个 seatbid对象是必须的， 表示对该展示的至少一个出价。 其他的属性都是可选的。
//如果要表示不出价， 可以选择使用空的HTTP响应体以及HTTP 204响应码。 如果竞拍者想要向交易平台传递不出价的原因， 可以返回一个只填充nbr属性的BidResponse对象。
type BidResponse struct {
	ID         string    `json:"id"`                   //竞价请求的标识
	SeatBid    []SeatBid `json:"seatbid"`              //一组SeatBid对象， 如果出价，则至少应该填充一个
	BidID      string    `json:"bidid,omitempty"`      //竞拍者生成的响应ID, 辅助日志或者交易追踪
	Currency   string    `json:"cur,omitempty"`        //使用ISO-4217码表标识货币类型
	CustomData string    `json:"customdata,omitempty"` //可选特性，允许出价者以设置cookie的方式向交易平台传递信息。 字符串可以是任何格式， 必须使用base85编码，JSON编码必须包含转义的引号。
	NBR        int       `json:"nbr,omitempty"`        //不出价的原因， 参考表5.19
	Ext        Extension `json:"ext,omitempty"`        //特定交易的OpenRTB协议的扩展信息占位符
}

// Validate required attributes
func (res *BidResponse) Validate() error {
	if res.ID == "" {
		return ErrInvalidRespNoID
	} else if len(res.SeatBid) == 0 {
		return ErrInvalidRespNoSeatBids
	}

	for _, sb := range res.SeatBid {
		if err := sb.Validate(); err != nil {
			return err
		}
	}

	return nil
}
