package openrtb

type Inventory struct {
	ID            string     `json:"id,omitempty"`            // ID on the exchange
	Name          string     `json:"name,omitempty"`          // Site name (may be aliased at the publisher’s request).
	Domain        string     `json:"domain,omitempty"`        // Domain of the site (e.g., “mysite.foo.com”).
	Cat           []string   `json:"cat,omitempty"`           // Array of IAB content categories
	SectionCat    []string   `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	PageCat       []string   `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	PrivacyPolicy *int       `json:"privacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Publisher     *Publisher `json:"publisher,omitempty"`     // Details about the Publisher
	Content       *Content   `json:"content,omitempty"`       // Details about the Content
	Keywords      string     `json:"keywords,omitempty"`      // Comma separated list of keywords about the site.
	Ext           Extension  `json:"ext,omitempty"`           // 特定交易的OpenRTB协议的扩展信息占位符
}

// GetPrivacyPolicy returns the privacy policy value
func (a *Inventory) GetPrivacyPolicy() int {
	if a.PrivacyPolicy != nil {
		return *a.PrivacyPolicy
	}
	return 1
}

// 如果广告载体是非浏览器应用（通常是移动设备）时应该包含该对象， 网站则不需要包含。
// 一个竞价请求一定不能同时包含Site对象和App对象。 提供一个App标识或者bundle是很有用的， 但是不是严格必须的。
type App struct {
	Inventory
	Bundle   string `json:"bundle,omitempty"`   // 应用信息或者包名（例如， com.foo.mygame); 需要是在整个交易过程中唯一的标识
	StoreURL string `json:"storeurl,omitempty"` // 应用的商店地址， 遵循AQG 1.5
	Ver      string `json:"ver,omitempty"`      // 应用版本号
	Paid     int    `json:"paid,omitempty"`     // 应用是否需要付费， 0表示免费， 1表示付费
}

// 如果广告载体是一个网站时应该包含这个对象，如果是非浏览器应用时则不需要。
// 一个竞价请求一定不能同时包含Site对象和App对象。
// 提供一个站点标识或者页面地址是很有用的， 但是不是严格必须的
type Site struct {
	Inventory
	Page   string `json:"page,omitempty"`   // 展示广告将要被展示的页面地址
	Ref    string `json:"ref,omitempty"`    // 引导到当前页面的referrer地址
	Search string `json:"search,omitempty"` // 引导到当前页面的搜索字符串
	Mobile int    `json:"mobile,omitempty"` // 移动优化标志， 0表示否，1表示是
}

// This object describes the nature and behavior of the entity that is the source of the bid request
// upstream from the exchange. The primary purpose of this object is to define post-auction or upstream
// decisioning when the exchange itself does not control the final decision. A common example of this is
// header bidding, but it can also apply to upstream server entities such as another RTB exchange, a
// mediation platform, or an ad server combines direct campaigns with 3rd party demand in decisioning
type Source struct {
	FD     int       `json:"id,omitempty"`     // Entity responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	TID    string    `json:"tid,omitempty"`    // Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	PChain string    `json:"pchain,omitempty"` // Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol v1.0.
	Ext    Extension `json:"ext,omitempty"`    // 特定交易的OpenRTB协议的扩展信息占位符
}

// This object is associated with an impression as an array of metrics. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
type Metric struct {
	Type   string    `json:"type,omitempty"`   // Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Value  string    `json:"value,omitempty"`  // Number representing the value of the metric. Probabilities must be in the range 0.0 – 1.0.
	Vendor string    `json:"vendor,omitempty"` // Source of the value using exchange curated string names which should be published to bidders a priori. If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	Ext    Extension `json:"ext,omitempty"`    // 特定交易的OpenRTB协议的扩展信息占位符
}
