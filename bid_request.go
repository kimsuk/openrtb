package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidReqNoID     = errors.New("openrtb: request ID missing")
	ErrInvalidReqNoImps   = errors.New("openrtb: request has no impressions")
	ErrInvalidReqMultiInv = errors.New("openrtb: request has multiple inventory sources") // has site and app
)

// 顶层竞价请求对象， 包含一个全局唯一的竞价请求或拍卖ID。 id属性是必须的因为其中至少包含一个展示对象。
// 在这个对象中的所有其他对象所建立的规则和限制适用于其中的所有展示对象。
// 有一些子对象用于提供潜在买家的详细信息，在这些对象中个，Site和App对象，用于描述展示广告的发布媒体的类型。
// 这些对象是高度推荐使用的， 但是每个竞价请求中只能使用一个，这取决于发起广告请求的是基于浏览器的页面内容或者非浏览器应用。
type BidRequest struct {
	ID          string       `json:"id"`                // 竞价请求的唯一ID, 由广告交易平台提供
	Imp         []Impression `json:"imp,omitempty"`     // 代表提供的展示信息的数组， 要求至少有一个
	Site        *Site        `json:"site,omitempty"`    // Site对象，表示发布者站点相关的详细信息， 仅仅对站点适用且推荐填充
	App         *App         `json:"app,omitempty"`     // App对象，表示发布应用的详细信息。仅对应用适用且推荐填充
	Device      *Device      `json:"device,omitempty"`  // Device对象， 表示展示将要被发送到的用户设备的信息
	User        *User        `json:"user,omitempty"`    // User对象， 表示使用设备的对象， 广告的受众
	Test        int          `json:"test,omitempty"`    // 标识测试模式，拍卖不计价。 0表示实况（非测试）模式，1表示测试模式 where 0 = live mode, 1 = test mode
	AuctionType int          `json:"at"`                // 拍卖类型（胜出策略）1表示第一价格 ，2标识第二价格。交易特定的拍卖类型可以用大于500的值定义
	TMax        int          `json:"tmax,omitempty"`    // 用于在提交竞价请求时避免超时的最大时间，以毫秒为单位，这个值通常是线下沟通的
	WSeat       []string     `json:"wseat,omitempty"`   // 允许在本次展现上进行竞拍的买家白Seat名单。 交易平台和竞拍者必须提前协商好Seat IDs
	BSeat       []string     `json:"bseat,omitempty"`   // Array of buyer seats blocked to bid on this auction
	WLang       []string     `json:"wlang,omitempty"`   // Array of languages for creatives using ISO-639-1-alpha-2
	AllImps     int          `json:"allimps,omitempty"` // 用于标识交易平台是否可以验证当前的展示列表包含了当前上下文中所有展示。（例如，一个页面上的所有广告位，所有的视频广告点（视频前，视频中，时候后））用于支持路由封锁。 0表示不支持或未知， 1表示提供的展示列表代表所有可用的展示。 Default: 0
	Cur         []string     `json:"cur,omitempty"`     // 本次竞价请求允许的货币列表， 使用ISO-4217 字母码。 仅在交易平台接收多种货币的时候推荐填充。
	Bcat        []string     `json:"bcat,omitempty"`    // 被封锁的广告主类别，使用IAB 内容类别，参考5.1。
	BAdv        []string     `json:"badv,omitempty"`    // 域名封锁列表（比如 ford.com)
	BApp        []string     `json:"bapp,omitempty"`    // Block list of applications by their platform-specific exchange-independent application identifiers. On Android, these should be bundle or package names (e.g., com.foo.mygame).  On iOS, these are numeric IDs.
	Source      *Source      `json:"source,omitempty"`  // A Sorce object (Section 3.2.2) that provides data about the inventory source and which entity makes the final decision.
	Regs        *Regulations `json:"regs,omitempty"`    // Reg对象， 指明对本次请求有效的工业，法律或政府条例
	Ext         Extension    `json:"ext,omitempty"`     // 特定交易的OpenRTB协议的扩展信息占位符
	//Pmp         *Pmp         `json:"pmp,omitempty"`     // DEPRECATED: kept for backwards compatibility
}

// Validates the request
func (req *BidRequest) Validate() error {
	if req.ID == "" {
		return ErrInvalidReqNoID
	} else if len(req.Imp) == 0 {
		return ErrInvalidReqNoImps
	} else if req.Site != nil && req.App != nil {
		return ErrInvalidReqMultiInv
	}

	for _, imp := range req.Imp {
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}

	return nil
}
