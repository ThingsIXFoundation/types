// Copyright 2023 Stichting ThingsIX Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

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
