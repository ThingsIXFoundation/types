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
	"encoding/json"
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

func (grh GatewayRewardHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"gatewayId":            grh.GatewayID,
		"date":                 grh.Date,
		"assumedCoverageUnits": "0x" + grh.AssumedCoverageShareUnits.Text(16),
		"rewards":              "0x" + grh.Rewards.Text(16),
	})
}
