package openrtb

// 5.2 Banner Ad Types, 广告类型
const (
	BannerTypeXHTMLText = 1 //XHTML 文本广告（通常用于移动手机）
	BannerTypeXHTML     = 2 //XHTML Banner 广告 （通常用于移动手机）
	BannerTypeJS        = 3 //Javascript 广告， 必须是合法的XHTML(即用script标签包裹)
	BannerTypeFrame     = 4 //iFrame广告
)

// 5.3 Creative Attributes, 物料属性
const (
	CreativeAttributeAudioAdAutoPlay                 = 1  //音频广告（自动播放）
	CreativeAttributeAudioAdUserInitiated            = 2  //音频广告（用户触发）
	CreativeAttributeExpandableAuto                  = 3  //可伸展（自动）
	CreativeAttributeExpandableUserInitiatedClick    = 4  //可伸展（用户点击触发）
	CreativeAttributeExpandableUserInitiatedRollover = 5  //可伸展（用户翻转触发）
	CreativeAttributeInBannerVideoAdAutoPlay         = 6  //Banner内视频广告（自动播放）
	CreativeAttributeInBannerVideoAdUserInitiated    = 7  //Banner内视频广告（用户触发）
	CreativeAttributePop                             = 8  //弹出广告（例如，在退出之前，之后或者当退出的时候）
	CreativeAttributeProvocativeOrSuggestiveImagery  = 9  //Provocative or Suggestive Imagery
	CreativeAttributeExtremeAnimation                = 10 //Shaky, Flashing, Flickering, Extreme Animation, Smileys
	CreativeAttributeSurveys                         = 11 //调查问卷
	CreativeAttributeTextOnly                        = 12 //文本广告
	CreativeAttributeUserInitiated                   = 13 //用户互动 （比如，嵌入式游戏）
	CreativeAttributeWindowsDialogOrAlert            = 14 //窗口对话框或弹出提示框
	CreativeAttributeHasAudioWithPlayer              = 15 //具有音频开关按钮
	CreativeAttributeAdProvidesSkipButton            = 16 //广告可以被跳过（例如，片头广告视频的跳过按钮）
	CreativeAttributeAdobeFlash                      = 17 //Flash广告
)

// 5.4 Ad Position, 广告位置
const (
	AdPosUnknown    = 0 //未知
	AdPosAboveFold  = 1
	AdPosBelowFold  = 3
	AdPosHeader     = 4
	AdPosFooter     = 5
	AdPosSidebar    = 6
	AdPosFullscreen = 7
)

// 5.5 Expandable Direction, 伸展方向
const (
	ExpDirLeft       = 1
	ExpDirRight      = 2
	ExpDirUp         = 3
	ExpDirDown       = 4
	ExpDirFullScreen = 5
)

// 5.6 API Frameworks, API 框架
const (
	APIFrameworkVPAID1 = 1 //VPAID 1.0
	APIFrameworkVPAID2 = 2 //VPAID 2.0
	APIFrameworkMRAID1 = 3 //MARID-1
	APIFrameworkORMMA  = 4 //ORMMA
	APIFrameworkMRAID2 = 5 //MRAID-2
)

// 5.7 Video Linearity, 视频线性
const (
	VideoLinearityLinear    = 1 //线性
	VideoLinearityNonLinear = 2 //非线性
)

// 5.8 Video and Audio Bid Response Protocols, 视频竞价响应协议
const (
	VideoProtoVAST1            = 1  //VAST 1.0
	VideoProtoVAST2            = 2  //VAST 2.0
	VideoProtoVAST3            = 3  //VAST 3.0
	VideoProtoVAST1Wrapper     = 4  //VAST 1.0 Wrapper
	VideoProtoVAST2Wrapper     = 5  //VAST 2.0 Wrapper
	VideoProtoVAST3Wrapper     = 6  //VAST 3.0 Wrapper
	VideoProtoVAST4            = 7  //VAST 4.0
	VideoProtoVAST4Wrapper     = 8  //VAST 4.0 Wrapper
	AudioProtocolDAAST1        = 9  //DAAST1.1
	AudioProtocolDAAST1Wrapper = 10 //DAAST1.1 Wrapper
)

// 5.9 Video Playback Methods, 视频播放方式
const (
	VideoPlaybackAutoSoundOn  = 1 //自动播放（有声）
	VideoPlaybackAutoSoundOff = 2 //自动播放（静音）
	VideoPlaybackClickToPlay  = 3 //点击播放
	VideoPlaybackMouseOver    = 4 //鼠标经过播放
)

// 5.9 Video Placement Types (Spec 2.5)
const (
	VideoPlacementInStream     = 1
	VideoPlacementInBanner     = 2
	VideoPlacementInArticle    = 3
	VideoPlacementInFeed       = 4
	VideoPlacementInterstitial = 5
)

// 5.10 Video Start Delay, 视频播放延时
const (
	VideoStartDelayPreRoll         = 0  //视频前广告
	VideoStartDelayGenericMidRoll  = -1 //普通视频中广告
	VideoStartDelayGenericPostRoll = -2 //普通视频后广告
)

// 5.11 Video Quality, 视频质量
const (
	VideoQualityUnknown      = 0 //未知
	VideoQualityProfessional = 1 //Professionally Produced 专业制作
	VideoQualityProsumer     = 2 //Prosumer 专业级
	VideoQualityUGC          = 3 //User Generated (UGC) 用户生成
)

// 5.12 VAST Companion Types, VAST 伴随类型 http://www.iab.net/vast
const (
	VASTCompanionStatic = 1 //静态资源
	VASTCompanionHTML   = 2 //HTML资源
	VASTCompanionIFrame = 3 //iframe资源
)

// 5.13 Content Delivery Methods, 内容传输方式
const (
	ContentDeliveryStreaming   = 1 //Streaming
	ContentDeliveryProgressive = 2 //Progressive
	ContentDeliveryDownload    = 3 //Download
)

// 5.14 Content Context, 内容上下文 该表继承自QAG,实现者应当与QAG的相关描述保持一致
const (
	ContextVideo       = 1 //视频(一个别用户观看的视频文件或者视频流， 包括网络电视广播)
	ContextGame        = 2 //游戏(用户正在玩的游戏)
	ContextMusic       = 3 //音乐(一个用户正在听的音频文件或者音频流， 包括网络无线广播)
	ContextApplication = 4 //应用(用户正在使用的交互式软件)
	ContextText        = 5 //文本(用户正在浏览或者阅读的文本文档， 包括web页面，电子书或者新闻文章)
	ContextOther       = 6 //其他(其他用户正在消费的位置类型的内容)
	ContextUnknown     = 7 //未知
)

// 5.15 QAG Media Ratings, QAG 媒体级别 http://www.iab.net/ne_guidelines
const (
	QAGAll    = 1 //所有受众
	QAGOver12 = 2 //大于12岁
	QAGMature = 3 //成年人
)

// 5.16 Location Type, 位置类型
const (
	LocationTypeGPS  = 1 //GPS/定位服务
	LocationTypeIP   = 2 //IP地址
	LocationTypeUser = 3 //用户提供(例如注册信息)
)

// 5.17 Device Type, 设备类型 该表继承自QAG
const (
	DeviceTypeUnknown   = 0 //未知
	DeviceTypeMobile    = 1 //手机/平板
	DeviceTypePC        = 2 //个人电脑
	DeviceTypeTV        = 3 //联网电视
	DeviceTypePhone     = 4 //手机
	DeviceTypeTablet    = 5 //平板
	DeviceTypeConnected = 6 //联网设备
	DeviceTypeSetTopBox = 7 //机顶盒
)

// 5.18 Connection Type, 连接方式
const (
	ConnTypeUnknown  = 0 //未知
	ConnTypeEthernet = 1 //Ethernet以太网
	ConnTypeWIFI     = 2 //WIFI
	ConnTypeCell     = 3 //移动网络，未知类型
	ConnTypeCell2G   = 4 //移动网络-2G
	ConnTypeCell3G   = 5 //移动网络-3G
	ConnTypeCell4G   = 6 //移动网络-4G
)

// 5.19 No-Bid Reason Codes, 状态返回
const (
	NBRUnknownError      = 0
	NBRTechnicalError    = 1
	NBRInvalidRequest    = 2
	NBRKnownSpider       = 3
	NBRSuspectedNonHuman = 4
	NBRProxyIP           = 5
	NBRUnsupportedDevice = 6
	NBRBlockedSite       = 7
	NBRUnmatchedUser     = 8
)

/*************************************************************************
 * COMMON OBJECT STRUCTS
 *************************************************************************/

// This object may be useful in the situation where syndicated content contains impressions and
// does not necessarily match the publisher's general content.  The exchange might or might not
// have knowledge of the page where the content is running, as a result of the syndication
// method.  (For example, video impressions embedded in an iframe on an unknown web property
// or device.)
// type Content struct {
// }

// Abstract third-party
type ThirdParty struct {
	ID     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Cat    []string  `json:"cat,omitempty"` // Array of IAB content categories
	Domain string    `json:"domain,omitempty"`
	Ext    Extension `json:"ext,omitempty"`
}

// The publisher object itself and all of its parameters are optional, so default values are not
// provided. If an optional parameter is not specified, it should be considered unknown.
type Publisher ThirdParty

// The producer is useful when content where the ad is shown is syndicated, and may appear on a
// completely different publisher. The producer object itself and all of its parameters are optional,
// so default values are not provided. If an optional parameter is not specified, it should be
// considered unknown.
type Producer ThirdParty

// Note that the Geo Object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Lat           float64   `json:"lat,omitempty"`           // Latitude from -90 to 90
	Lon           float64   `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          int       `json:"type,omitempty"`          // Indicate the source of the geo data
	Accuracy      int       `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int       `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int       `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string    `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string    `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string    `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string    `json:"metro,omitempty"`
	City          string    `json:"city,omitempty"`
	Zip           string    `json:"zip,omitempty"`
	UTCOffset     int       `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           Extension `json:"ext,omitempty"`
}

// This object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID         string    `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerID    string    `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID   string    `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	YOB        int       `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender     string    `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords   string    `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData string    `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo        *Geo      `json:"geo,omitempty"`
	Data       []Data    `json:"data,omitempty"`
	Ext        Extension `json:"ext,omitempty"`
}

// The data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Segment []Segment `json:"segment,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Value string    `json:"value,omitempty"`
	Ext   Extension `json:"ext,omitempty"`
}

// This object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	Coppa int       `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Ext   Extension `json:"ext,omitempty"`
}

// This object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	W   int       `json:"w,omitempty"` // Width in device independent pixels (DIPS).
	H   int       `json:"h,omitempty"` //Height in device independent pixels (DIPS).
	Ext Extension `json:"ext,omitempty"`
}
