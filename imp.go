package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidImpNoID        = errors.New("openrtb: impression ID missing")
	ErrInvalidImpMultiAssets = errors.New("openrtb: impression has multiple assets") // at least two out of Banner, Video, Native
)

// 这个对象描述了一个广告位或者将要参与竞拍的展现。一个竞价请求可以包含多个Imp对象，
// 这种状况的一个示例是一个交易平台支持售卖一个页面的所有广告位。 为了便于竞拍者区分， 每一个Imp对象都要有一个唯一标识（ID).
// Banner, Video以及Native对象都属于Imp对象，只是标明了各自的展示类型。 展示者可以选择其中的一种类型或者混合使用多种类型。
// 但是，对于展示的任何给定的竞价请求必须属于提供的类型之一。
type Impression struct {
	ID                string         `json:"id"`                          // 在当前竞价请求上下文中唯一标识本次展示的标识（通常从1开始并以此递增)
	Metric            []Metric       `json:"metric,omitempty"`            // An array of Metric object (Section 3.2.5).
	Banner            *Banner        `json:"banner,omitempty"`            // Banner对象，如果展示需要以banner的形式提供则需要填充
	Video             *Video         `json:"video,omitempty"`             // Video对象，如果展示需要以视频的形式提供则需要填充
	Audio             *Audio         `json:"audio,omitempty"`             // Audio对象
	Native            *Native        `json:"native,omitempty"`            // Native对象，如果展示需要以Native广告的形式提供则需要填充
	Pmp               *Pmp           `json:"pmp,omitempty"`               // Pmp对象 包含对本次展示生效的任何私有市场交易
	DisplayManager    string         `json:"displaymanager,omitempty"`    // 广告媒体合作伙伴的名字，用于渲染广告的SDK技术或者播放器（通常是视频或者移动广告）某些广告服务需要根据合作伙伴定制广告代码， 推荐在视频广告或应用广告中填充
	DisplayManagerVer string         `json:"displaymanagerver,omitempty"` // 广告媒体合作伙伴， 用于渲染广告的SDK技术或者播放器（通常是视频或者移动广告）的版本号。 某些广告服务需要根据合作伙伴定制广告代码， 推荐在视频广告或应用广告中填充
	Instl             int            `json:"instl,omitempty"`             // 1标识广告是插屏或者全屏广告，0表示不是插屏广告
	TagID             string         `json:"tagid,omitempty"`             // 某个特定广告位或者广告标签的标识，用于发起竞拍。为了方便调试问题或者进行买方优化
	BidFloor          float64        `json:"bidfloor,omitempty"`          // 本次展示的最低竞拍价格，以CPM表示
	BidFloorCurrency  string         `json:"bidfloorcur,omitempty"`       // ISO-4217规定的字母码表标识的货币类型。如果交易平台允许，可能与从竞拍者返回的竞价货币不同
	Secure            NumberOrString `json:"secure,omitempty"`            // 标识展示请求是否需要使用HTTPS加密物料信息以及markup以保证安全，0标识不需要使用安全链路，1标识需要使用安全链路，如果不填充，则表示未知，可以认为是不需要使用安全链路
	Exp               int            `json:"exp,omitempty"`               // 特定交易的OpenRTB协议的扩展信息占位符
	IFrameBuster      []string       `json:"iframebuster,omitempty"`      // 特定交易支持的iframe buster的名字数组
	Ext               Extension      `json:"ext,omitempty"`               // 特定交易的OpenRTB协议的扩展信息占位符
}

func (imp *Impression) assetCount() int {
	n := 0
	if imp.Banner != nil {
		n++
	}
	if imp.Video != nil {
		n++
	}
	if imp.Native != nil {
		n++
	}
	return n
}

// Validates the `imp` object
func (imp *Impression) Validate() error {
	if imp.ID == "" {
		return ErrInvalidImpNoID
	}

	if count := imp.assetCount(); count > 1 {
		return ErrInvalidImpMultiAssets
	}

	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}

	return nil
}
