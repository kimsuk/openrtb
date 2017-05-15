package openrtb

// Banner是最常见的展示类型。 虽然“banner”这个名词在其他的场景有很特别的意思， 在这里它可以是包括静态图像，
// 可扩展的广告单元或者一个在banner中播放的视频在内的很多东西。
// 一组Banner对象可以出现在Video对象中来描述可选择在VAST贵方中定义的附加广告。
// Banner作为Imp的子对象出现表示它是一个具有banner类型的展示对象. 同样的展示也可以是一个视频或者Native广告， 只要包含Video对象或者Native对象。
// 然而，任何为展示给定的竞价请求必须符合提供类型中的一个。

type Banner struct {
	W        int       `json:"w,omitempty"`        // 展示的宽度，以像素为单位，如果没有指定wmin以及wmax, 这个值指的就是需要的展示宽度，否则指的是一个期望宽度
	H        int       `json:"h,omitempty"`        // 展示的高度，以像素为单位，如果没有指定hmin以及hmax, 这个值指的就是需要的展示高度， 否则指的是一个期望高度
	Format   []Format  `json:"format,omitempty"`   // Array of format objects representing the banner sizes permitted.
	WMax     int       `json:"wmax,omitempty"`     // 展示宽度的最大值，以像素为单位， 如果和w一起出现， 则w应该被解释为推荐宽度或者期望宽度
	HMax     int       `json:"hmax,omitempty"`     // 展示高度的最大值，以像素为单位， 如果和h一起出现， 则h应该被解释为推荐宽度或者期望宽度
	WMin     int       `json:"wmin,omitempty"`     // 展示宽度的最小值，以像素为单位， 如果和w一起出现， 则w应该被解释为推荐宽度或者期望宽度
	HMin     int       `json:"hmin,omitempty"`     // 展示高度的最小值，以像素为单位， 如果和h一起出现， 则h应该被解释为推荐宽度或者期望宽度
	ID       string    `json:"id,omitempty"`       // banner对象的唯一标识。在一个Ad中包含Banner与Video的时候使用。值往往从1开始并依次递增，在依次展示中应当是唯一的
	BType    []int     `json:"btype,omitempty"`    // 限制的banner类型 5.2
	BAttr    []int     `json:"battr,omitempty"`    // 限制的物料属性 5.3
	Pos      int       `json:"pos,omitempty"`      // 广告在屏幕上的位置 5.4
	Mimes    []string  `json:"mimes,omitempty"`    // 支持的内容mime-type. 常用的mime-type包括application/x-shockwave-flash, image/jpg以及 image/gif.
	TopFrame int       `json:"topframe,omitempty"` // banner是在顶层frame中而不是iframe中， 0表示不是， 1表示是
	ExpDir   []int     `json:"expdir,omitempty"`   // banner可以扩展的方向 5.5
	Api      []int     `json:"api,omitempty"`      // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	Ext      Extension `json:"ext,omitempty"`      // 特定交易的OpenRTB协议的扩展信息占位符
}
