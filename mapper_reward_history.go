package types

import (
	"math/big"
	"time"
)

// MapperRewardHistory stores the historical data of rewards as received by a mapper
type MapperRewardHistory struct {

	// ID of the mapper
	MapperID ID `json:"mapperId"`

	// Date these rewards where issued
	Date time.Time `json:"date"`

	// The total amount of MappingUnits the mapper got rewards for
	MappingUnits *big.Int `json:"mappingUnits"`

	// The reward in THIX "gweis" for this mapper
	Rewards *big.Int `json:"rewards"`
}
