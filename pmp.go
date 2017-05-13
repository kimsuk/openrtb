package openrtb

import "encoding/json"

type Pmp struct {
	Private int       `json:"private_auction,omitempty"`
	Deals   []Deal    `json:"deals,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}

// PMP Deal
type Deal struct {
	ID               string    `json:"id,omitempty"`
	BidFloor         float64   `json:"bidfloor,omitempty"`
	BidFloorCurrency string    `json:"bidfloorcur,omitempty"`
	WSeat            []string  `json:"wseat,omitempty"`
	WAdvDomain       []string  `json:"wadomain,omitempty"`
	AuctionType      int       `json:"at,omitempty"`
	Ext              Extension `json:"ext,omitempty"`

	Seats []string `json:"seats,omitempty"`
	Type  int      `json:"type,omitempty"`
}

type jsonDeal Deal

// MarshalJSON custom marshalling with normalization
func (d *Deal) MarshalJSON() ([]byte, error) {
	d.normalize()
	return json.Marshal((*jsonDeal)(d))
}

// UnmarshalJSON custom unmarshalling with normalization
func (d *Deal) UnmarshalJSON(data []byte) error {
	var h jsonDeal
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*d = (Deal)(h)
	d.normalize()
	return nil
}

func (d *Deal) normalize() {
	if d.AuctionType == 0 {
		d.AuctionType = 2
	}
}
