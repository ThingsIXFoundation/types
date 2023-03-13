package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type AccountRewardHistory struct {
	// The acount this reward history is for
	Account common.Address `json:"account"`

	// The reward in THIX "gweis" for this account at this moment
	Rewards *big.Int `json:"rewards"`

	// The sum of rewards ever received uptil now
	TotalRewards *big.Int `json:"totalRewards"`

	// The processor that can process the cheque for the total rewards
	Processor common.Address `json:"processor"`

	// The signature by the cheque creator that can be used to claim the rewards
	Signature hexutil.Bytes `json:"signature"`

	// Date these rewards where issued
	Date time.Time `json:"date"`
}
