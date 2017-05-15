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

// 用于封装一个地理位置信息的多种不同属性。 当作为Device对象的子节点的时候，标识设备的地理位置或者用户当前的地理位置。
// 当作为User的子节点的时候，标识用户家的位置（也就是说，不必是用户的当前位置）。
// 设备的位置或者用户家的位置，由其父对象决定
type Geo struct {
	Lat           float64   `json:"lat,omitempty"`           // 纬度信息，取值范围-90.0到+90.0， 负值表示南方
	Lon           float64   `json:"lon,omitempty"`           // 经度信息， 取值返回-180.0到+180.0， 负值表示西方
	Type          int       `json:"type,omitempty"`          // 位置信息的源， 当传递lat/lon的时候推荐填充， 参考表5.16
	Accuracy      int       `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int       `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int       `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string    `json:"country,omitempty"`       // 国家码， 使用 ISO-3166-1-alpha-3
	Region        string    `json:"region,omitempty"`        // 区域码， 使用ISO-3166-2; 如果美国则使用2字母区域码
	RegionFIPS104 string    `json:"regionFIPS104,omitempty"` // 国家的区域，使用 FIPS 10-4 表示。 虽然OpenRTB支持这个属性，它已经与2008年被NIST撤销了
	Metro         string    `json:"metro,omitempty"`         // 谷歌metro code; 与Nielsen DMA相似但不完全相同， 参见附录A
	City          string    `json:"city,omitempty"`          // 城市名，使用联合国贸易与运输位置码， 参见附录A
	Zip           string    `json:"zip,omitempty"`           // 邮政编码或者邮递区号
	UTCOffset     int       `json:"utcoffset,omitempty"`     // 使用UTC加或者减分钟数的方式表示的本地时间
	Ext           Extension `json:"ext,omitempty"`           // 特定交易的OpenRTB协议的扩展信息占位符
}

// 描述了解或者持有设备的用户的信息（也就是广告的受众）。
// 用户id是一个exchange artifact, 可能随着屏幕旋转或者其他的隐私策略改变。
// 尽管如此，用户id必须在足够长的一段时间内保持不变，以为目标用户定向和用户访问频率限制提供合理的服务。
// 设备的用户， 广告的受众
type User struct {
	ID         string    `json:"id,omitempty"`         // 交易特定的用户标识， 推荐id和buyeruid中至少提供一个
	BuyerID    string    `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID   string    `json:"buyeruid,omitempty"`   // 买方为用户指定的ID，由交易平台为买方映射。推荐id和buyeruid中至少提供一个.
	YOB        int       `json:"yob,omitempty"`        // 生日年份，使用4位数字表示
	Gender     string    `json:"gender,omitempty"`     // 性别， M表示男性， F表示女性， O标识其他类型，不填充表示未知
	Keywords   string    `json:"keywords,omitempty"`   // 逗号分隔的关键字， 兴趣或者意向列表
	CustomData string    `json:"customdata,omitempty"` // 可选特性， 用于传递给竞拍者信息，在交易平台的cookie中设置。字符串必须使用base85编码的 cookie，可以是任意格式。 JSON加密的时候必须包括转义的引号
	Geo        *Geo      `json:"geo,omitempty"`        // Geo对象， 用户家的位置信息。不必是用户的当前位置
	Data       []Data    `json:"data,omitempty"`       // 附加的用户信息， 每个 Data对象表示一个不同的数据源
	Ext        Extension `json:"ext,omitempty"`        // 特定交易的OpenRTB协议的扩展信息占位符
}

// Data和Segment对象一起允许指定用户附加信息。数据可能来自多个数据源， 可能来自交易平台自身或者第三方提供的信息， 可以使用id属性区分。
// 一个竞价请求可以混合来自多个提供者的数据信息。 交易平台应该优先提供正在使用的数据提供者的信息
type Data struct {
	ID      string    `json:"id,omitempty"`      // 交易特定的数据提供者标识
	Name    string    `json:"name,omitempty"`    // 交易特定的数据提供者名称
	Segment []Segment `json:"segment,omitempty"` // 包含数据内容的一组Segment对象
	Ext     Extension `json:"ext,omitempty"`     // 特定交易的OpenRTB协议的扩展信息占位符
}

// 数据字段， 描述用户信息数据的键值对。 其父对象Data是某个给定数据提供者的数据字段的集合。
// 交易平台必须优先将字段的名称和值传递给竞拍者
type Segment struct {
	ID    string    `json:"id,omitempty"`    // 数据提供者的特定数据段的ID
	Name  string    `json:"name,omitempty"`  // 数据提供者的特定数据段的名称
	Value string    `json:"value,omitempty"` // 表示数据字段值的字符串
	Ext   Extension `json:"ext,omitempty"`   // 特定交易的OpenRTB协议的扩展信息占位符
}

// 描述任何适用于该请求的法律，政府或者工业管控条例。
// coppa(Children’s Online Privacy Protection Act)标志着是否该请求是否符合美国联邦贸易委员会颁布的美国儿童在线隐私保护法案，详情可参考7.1节
type Regulations struct {
	Coppa int       `json:"coppa,omitempty"` // 标志着该请求是否遵从COPPA法案， 0表示不遵从， 1表示遵从
	Ext   Extension `json:"ext,omitempty"`   // 特定交易的OpenRTB协议的扩展信息占位符
}

// This object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	W   int       `json:"w,omitempty"`   // Width in device independent pixels (DIPS).
	H   int       `json:"h,omitempty"`   // Height in device independent pixels (DIPS).
	Ext Extension `json:"ext,omitempty"` //
}
