package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidBidNoID    = errors.New("openrtb: bid is missing ID")
	ErrInvalidBidNoImpID = errors.New("openrtb: bid is missing impression ID")
)

// 一个SeatBid对象包括一个或者多个Bid对象， 每一个Bid对象通过impid关联竞价请求中的一个展示, 由一个对该展示出价组成。
type Bid struct {
	ID             string         `json:"id"`
	ImpID          string         `json:"impid"`                    // 竞拍者生成的竞价ID，用于记录日志或行为追踪
	Price          float64        `json:"price"`                    // 关联的竞价请求中的Imp对象的ID
	AdID           string         `json:"adid,omitempty"`           // 虽然本次只是对某一个展示的出价，但是竞拍价格是以CPM表示。需要注意数据类型是float,所以在处理货币的时候强烈建议使用相关的数学处理对象（比如，Java中的BigDecimal)
	NURL           string         `json:"nurl,omitempty"`           // 预加载的广告ID, 可以在交易胜出的时候使用
	BURL           string         `json:"burl,omitempty"`           // 胜出通知地址， 如果竞价胜出的时候由交易平台调用； 可选标识serving ad markup
	LURL           string         `json:"lurl,omitempty"`           // 落败通知地址
	AdMarkup       string         `json:"adm,omitempty"`            // Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	AdvDomain      []string       `json:"adomain,omitempty"`        // Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	Bundle         string         `json:"bundle,omitempty"`         // A platform-specific application identifier intended to be unique to the app and independent of the exchange.
	IURL           string         `json:"iurl,omitempty"`           // Sample image URL.
	CampaignID     StringOrNumber `json:"cid,omitempty"`            // Campaign ID that appears with the Ad markup.
	CreativeID     string         `json:"crid,omitempty"`           // Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Tactic         string         `json:"tactic,omitempty"`         // Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted.
	Cat            []string       `json:"cat,omitempty"`            // creative的IAB内容类型，参考表5.1
	Attr           []int          `json:"attr,omitempty"`           // 描述creative的属性集合，参考表5.3
	API            int            `json:"api,omitempty"`            // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	Protocol       int            `json:"protocol,omitempty"`       // 支持的视频竞价响应协议。参考5.8.至少一个支持的协议必须在protocol或者protocols属性中被指定
	QAGMediaRating int            `json:"qagmediarating,omitempty"` // Creative media rating per IQG guidelines.
	Language       string         `json:"language,omitempty"`       // Language of the creative using ISO-639-1-alpha-2.
	DealID         string         `json:"dealid,omitempty"`         // 如果出价从属与某个私有市场直接交易规则， 则指向竞价请求中该交易规则的deal.id
	H              int            `json:"h,omitempty"`              // creative 的高度， 以像素为单位
	W              int            `json:"w,omitempty"`              // creative 的宽度， 以像素为单位
	WRatio         int            `json:"wratio,omitempty"`         // Relative width of the creative when expressing size as a ratio.
	HRatio         int            `json:"hratio,omitempty"`         // Relative height of the creative when expressing size as a ratio.
	Exp            int            `json:"exp,omitempty"`            // Advisory as to the number of seconds the bidder is willing to wait between the auction and the actual impression.
	Ext            Extension      `json:"ext,omitempty"`            // 特定交易的OpenRTB协议的扩展信息占位符
}

// Validate required attributes
func (bid *Bid) Validate() error {
	if bid.ID == "" {
		return ErrInvalidBidNoID
	} else if bid.ImpID == "" {
		return ErrInvalidBidNoImpID
	}

	return nil
}
