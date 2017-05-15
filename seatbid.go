package openrtb

import "errors"

// 一个竞价响应可以包含多个SeatBid对象， 每个代表着不同的出价者且包含一个或多个独立的出价信息。
// 如果请求中有多个展示信息， group属性可以用来指定一个席位对胜出任何展示感兴趣（可以不是全部展示）或者它仅对胜出所有展示感兴趣
type SeatBid struct {
	Bid   []Bid     `json:"bid"`             // 至少一个Bid对象的数组，每个对象关联一个展示。多个出价可以关联同一个展示
	Seat  string    `json:"seat,omitempty"`  // 出价者席位标识， 代表本次出价的出价人
	Group int       `json:"group,omitempty"` // 0 标识展示可以独立胜出， 1标识展示必须整组胜出或失败
	Ext   Extension `json:"ext,omitempty"`   // 特定交易的OpenRTB协议的扩展信息占位符
}

// Validation errors
var (
	ErrInvalidSeatBidBid = errors.New("openrtb: seatbid is missing bids")
)

// Validate required attributes
func (sb *SeatBid) Validate() error {
	if len(sb.Bid) == 0 {
		return ErrInvalidSeatBidBid
	}

	for _, bid := range sb.Bid {
		if err := bid.Validate(); err != nil {
			return err
		}
	}

	return nil
}
