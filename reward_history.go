package types

import (
	"math/big"
	"time"
)

type RewardHistory struct {
	// Date these rewards where issued
	Date time.Time `json:"date"`

	// The total amount of Coverage Share Units issued
	TotalAssumedCoverageShareUnits *big.Int `json:"totalAssumedCoverageUnits"`

	// The total amount of MappingUnits the mapper got rewards for
	TotalMappingUnits *big.Int `json:"totalMappingUnits"`

	// The total rewards issued in THIX "gweis"
	TotalRewards *big.Int `json:"totalRewards"`
}
