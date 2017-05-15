package openrtb

//表示一个Native类型的展示。Native广告单元需要无缝的插入其周围的内容中（例如， 一个对Twitter或Facebook赞助）。
//因此，响应必须具有良好的结构， 让展示者能够在细粒度上控制渲染过程。
//Native小组委员会为OpenRTB开发了一个组合规范，名为Native Ad规范。 定义了Native广告的请求参数以及响应结构。
//这个对象以字符串的形式提供请求参数， 这样的话具体的请求参数就可以按照Native Ad规范独立演进。同样的， 广告markup也会按照该文档指定的结构提供。
//Native作为Imp的子对象出现表示它是一个具有native类型的展示对象。
//同样的展示也可以是一个Banner或者Video广告， 只要包含Banner对象或者Video对象。
//然而， 任何为展示给定的竞价请求必须符合提供类型中的一个。
type Native struct {
	Request Extension `json:"request"`         //遵守Native Ad规范的请求体
	Ver     string    `json:"ver,omitempty"`   //Native Ad规范的版本， 为了高效解析强烈推荐
	API     []int     `json:"api,omitempty"`   //本次展示支持的API框架列表， 参考表5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	BAttr   []int     `json:"battr,omitempty"` //限制的物料属性，参考表5.3
	Ext     Extension `json:"ext,omitempty"`   //特定交易的OpenRTB协议的扩展信息占位符
}
