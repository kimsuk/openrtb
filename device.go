package openrtb

type Device struct {
	UA         string    `json:"ua,omitempty"`             // 浏览器User-Agent字符串
	Geo        *Geo      `json:"geo,omitempty"`            // Geo对象，用用户当前位置表示设备位置
	DNT        int       `json:"dnt,omitempty"`            // 浏览器在HTTP头中设置的标准的 “Do Not Track"标识， 0表示不限制追踪， 1表示限制（不允许）追踪
	LMT        int       `json:"lmt,omitempty"`            // "1": Limit Ad Tracking
	IP         string    `json:"ip,omitempty"`             // 最接近设备的IPv4地址
	IPv6       string    `json:"ipv6,omitempty"`           // 最接近设备的IPV6地址
	DeviceType int       `json:"devicetype,omitempty"`     // 设备类型，参考5.17
	Make       string    `json:"make,omitempty"`           // 设备制造商，例如 "Apple"
	Model      string    `json:"model,omitempty"`          // 设备型号，例如 "iphone"
	OS         string    `json:"os,omitempty"`             // 设备操作系统， 例如 "ios"
	OSVer      string    `json:"osv,omitempty"`            // 设备操作系统版本号， 例如 “3.1.2”
	HwVer      string    `json:"hwv,omitempty"`            // 设备硬件版本， 例如 “5S” (e.g., "5S" for iPhone 5S).
	H          int       `json:"h,omitempty"`              // 屏幕的物理高度， 以像素为单位
	W          int       `json:"w,omitempty"`              // 屏幕的物理宽度，以像素为单位
	PPI        int       `json:"ppi,omitempty"`            // 以像素每英寸表示的屏幕尺寸
	PxRatio    float64   `json:"pxratio,omitempty"`        // 设备物理像素与设备无关像素的比率
	JS         int       `json:"js,omitempty"`             // 支持javascript, 0表示不支持， 1表示支持
	GeoFetch   int       `json:"geofetch,omitempty"`       // Indicates if the geolocation API will be available to JavaScript code running in the banner,
	FlashVer   string    `json:"flashver,omitempty"`       // 浏览器支持的Flash版本
	Language   string    `json:"language,omitempty"`       // 浏览器语言，使用ISO-639-1-alpha-2
	Carrier    string    `json:"carrier,omitempty"`        // ISP的附带信息（如版本号）。“WIFI"通常在移动设备中表示高带宽。（例如,video freendly vs. cellular).
	MCCMNC     string    `json:"mccmnc,omitempty"`         // Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA).
	ConnType   int       `json:"connectiontype,omitempty"` // 网络连接类型， 参考表5.18
	IFA        string    `json:"ifa,omitempty"`            // 广告主标识， 明文表示
	IDSHA1     string    `json:"didsha1,omitempty"`        // 硬件设备ID(例如 IMEI),使用SHA1哈希算法
	IDMD5      string    `json:"didmd5,omitempty"`         // 硬件设备ID(例如 IMEI),使用md5哈希算法
	PIDSHA1    string    `json:"dpidsha1,omitempty"`       // 设备平台ID(例如 Android ID),使用SHA1哈希算法
	PIDMD5     string    `json:"dpidmd5,omitempty"`        // 设备平台ID(例如 Android ID),使用md5哈希算法
	MacSHA1    string    `json:"macsha1,omitempty"`        // 设备mac地址,使用SHA1哈希算法
	MacMD5     string    `json:"macmd5,omitempty"`         // 设备mac地址,使用md5哈希算法
	Ext        Extension `json:"ext,omitempty"`            // 特定交易的OpenRTB协议的扩展信息占位符
}
