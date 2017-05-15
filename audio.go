package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidAudioNoMimes = errors.New("openrtb: audio has no mimes")
)

type Audio struct {
	Mimes         []string  `json:"mimes"`                   // 支持的内容mime-type
	MinDuration   int       `json:"minduration,omitempty"`   // 最小的音频长度， 以秒为单位
	MaxDuration   int       `json:"maxduration,omitempty"`   // 最大的音频长度， 以秒为单位
	Protocols     []int     `json:"protocols,omitempty"`     // 支持的音频竞价响应协议数组。参考表5.8.至少一个支持的协议必须在protocol或者protocols属性中被指定
	StartDelay    int       `json:"startdelay,omitempty"`    // 音频前，中及之后的广告位中音频广告的启动延时，以秒为单位
	Sequence      int       `json:"sequence,omitempty"`      // 默认: 1
	BAttr         []int     `json:"battr,omitempty"`         // 限制的物料属性
	MaxExtended   int       `json:"maxextended,omitempty"`   // 最大的音频广告延长时间长度（如果支持延长）。如果为空或者0，表示不允许延长， 如果为-1，表示允许延时，且没有时间限制， 如果为大于0的数字， 则表示可以延长的时间长度比maxduration大的值
	MinBitrate    int       `json:"minbitrate,omitempty"`    // 最小的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	MaxBitrate    int       `json:"maxbitrate,omitempty"`    // 最大的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	Delivery      []int     `json:"delivery,omitempty"`      // 支持的传输方式 （例如流式传输，逐步传输），如果没有指定，表示全部支持
	CompanionAd   []Banner  `json:"companionad,omitempty"`   // 如果支持复合广告，表示一组Banner对象
	API           []int     `json:"api,omitempty"`           // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	CompanionType []int     `json:"companiontype,omitempty"` // 支持的VAST companion 广告类型， 参考5.12。 如果在companionad中填充了Banner对象则推荐使用
	MaxSequence   int       `json:"maxseq,omitempty"`        // 在广告位中可以播放的广告数量上限。
	Feed          int       `json:"feed,omitempty"`          // 音频的提供类型
	Stitched      int       `json:"stitched,omitempty"`      // 指示广告是否与音频内容拼接或独立传送
	NVol          int       `json:"nvol,omitempty"`          // 音量归一化模式。
	Ext           Extension `json:"ext,omitempty"`           // 特定交易的OpenRTB协议的扩展信息占位符
}

type jsonAudio Audio

// Validates the object
func (a *Audio) Validate() error {
	if len(a.Mimes) == 0 {
		return ErrInvalidAudioNoMimes
	}
	return nil
}

// MarshalJSON custom marshalling with normalization
func (a *Audio) MarshalJSON() ([]byte, error) {
	a.normalize()
	return json.Marshal((*jsonAudio)(a))
}

// UnmarshalJSON custom unmarshalling with normalization
func (a *Audio) UnmarshalJSON(data []byte) error {
	var h jsonAudio
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*a = (Audio)(h)
	a.normalize()
	return nil
}

func (a *Audio) normalize() {
	if a.Sequence == 0 {
		a.Sequence = 1
	}
}
