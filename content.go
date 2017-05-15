package openrtb

type Content struct {
	ID                 string    `json:"id,omitempty"`                 //内容唯一标识
	Episode            int       `json:"episode,omitempty"`            //情节数目（通常用于视频内容）
	Title              string    `json:"title,omitempty"`              //内容标题。 视频示例： “Search Committee"(电视）， ”A New Hope"(电影), "Endgame"(为网络制作） 非视频示例： “Why an Antarctic Glacier is Melting So Quickly"(时报杂志文章）
	Series             string    `json:"series,omitempty"`             //内容系列。 视频示例：“The Office"(电视）， ”Start Wars"(电影,"Arby 'N' The Chief"(为网络制作） 非视频示例： “Ecocentric"(时报杂志博客）
	Season             string    `json:"season,omitempty"`             //内容季数， 通常用于视频内容（例如，“第三季”）
	Artist             string    `json:"artist,omitempty"`             // Artist credited with the content.
	Genre              string    `json:"genre,omitempty"`              // Genre that best describes the content
	Album              string    `json:"album,omiyempty"`              // Album to which the content belongs; typically for audio.
	ISRC               string    `json:"isrc,omitempty"`               // International Standard Recording Code conforming to ISO - 3901.
	Producer           *Producer `json:"producer,omitempty"`           //内容提供者的详细信息
	URL                string    `json:"url,omitempty"`                //内容的url, 用于买方了解使用的上下文或者审查
	Cat                []string  `json:"cat,omitempty"`                //内容生产者的IAB内容类型数组， 参考表5.1
	ProdQuality        int       `json:"prodq,omitempty"`              // Production quality per IAB's classification.
	VideoQuality       int       `json:"videoquality,omitempty"`       //视频质量，按照IAB的分类，参考表5.11
	Context            int       `json:"context,omitempty"`            //内容类型（游戏，视频，文本等）， 参考表5.14
	ContentRating      string    `json:"contentrating,omitempty"`      //内容分级（例如， MPAA美国电影分级制度)
	UserRating         string    `json:"userrating,omitempty"`         //内容的用户评分（比如，星数，点赞数等）
	QAGMediaRating     int       `json:"qagmediarating,omitempty"`     //媒体评分，按照QAG规范。参考表5.15
	Keywords           string    `json:"keywords,omitempty"`           //逗号分隔的内容的关键字信息
	LiveStream         int       `json:"livestream,omitempty"`         //0表示不是实时，1表示实时
	SourceRelationship int       `json:"sourcerelationship,omitempty"` //0表示间接源， 1表示直接源
	Len                int       `json:"len,omitempty"`                //内容长度， 用于音频或者视频
	Language           string    `json:"language,omitempty"`           //内容语言， 使用ISO-639-1-alpha-2
	Embeddable         int       `json:"embeddable,omitempty"`         //表示内容是否可嵌套（例如一个可嵌套的视频播放器）， 0表示不可以， 1表示可以
	Data               []Data    `json:"data,omitempty"`               // Additional content data.
	Ext                Extension `json:"ext,omitempty"`                //特定交易的OpenRTB协议的扩展信息占位符
}
