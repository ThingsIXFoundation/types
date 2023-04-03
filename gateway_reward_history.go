package types

import (
	"math/big"
	"time"
)

// GatewayRewardHistory stores the historical data of rewards as received by a gateway
type GatewayRewardHistory struct {

	// ID of the gateway
	GatewayID ID `json:"gatewayId"`

	// Date these rewards where issued
	Date time.Time `json:"date"`

	// The total amount of Coverage Share Units this gateway has a the date.
	AssumedCoverageShareUnits *big.Int `json:"assumedCoverageUnits"`

	// The reward in THIX "gweis" for this gateway
	Rewards *big.Int `json:"rewards"`
}
