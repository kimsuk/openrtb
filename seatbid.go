package openrtb

import "errors"

type SeatBid struct {
	Bid   []Bid     `json:"bid"`             // Array of bid objects; each realtes to an imp, if exchange supported can have many bid objects.
	Seat  string    `json:"seat,omitempty"`  // ID of the bidder seat optional string ID of the bidder seat on whose behalf this bid is made.
	Group int       `json:"group,omitempty"` // '1' means impression must be won-lost as a group; default is '0'.
	Ext   Extension `json:"ext,omitempty"`
}

// Validation errors
var (
	ErrInvalidSeatBidBid = errors.New("openrtb: seatbid is missing bids")
)

// Validate required attributes
func (sb *SeatBid) Validate() error {
	if len(sb.Bid) == 0 {
		return ErrInvalidSeatBidBid
	}

	for _, bid := range sb.Bid {
		if err := bid.Validate(); err != nil {
			return err
		}
	}

	return nil
}
