package openrtb

type Native struct {
	Request Extension `json:"request"`
	Ver     string    `json:"ver,omitempty"`
	API     []int     `json:"api,omitempty"`
	BAttr   []int     `json:"battr,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}
