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
	"time"

	"github.com/ThingsIXFoundation/frequency-plan/go/frequency_plan"
	h3light "github.com/ThingsIXFoundation/h3-light"
	"github.com/ethereum/go-ethereum/common"
)

// GatewayEvent represents a log emitted by the gateway registry that is related
// to a gateway.
type GatewayEvent struct {
	ContractAddress  common.Address `json:"contract"`
	Block            common.Hash    `json:"blockHash"`
	Transaction      common.Hash    `json:"transaction"`
	BlockNumber      uint64         `json:"blockNumber"`
	TransactionIndex uint           `json:"transactionIndex"`
	LogIndex         uint           `json:"logIndex"`
	Time             time.Time      `json:"time"`

	Type             GatewayEventType         `json:"type"`
	ID               ID                       `json:"id"`
	Version          uint8                    `json:"version"`
	NewOwner         *common.Address          `json:"newOwner"`
	OldOwner         *common.Address          `json:"oldOwner"`
	NewLocation      *h3light.Cell            `json:"newLocation,omitempty"`
	OldLocation      *h3light.Cell            `json:"oldLocation,omitempty"`
	NewAltitude      *uint                    `json:"newAltitude,omitempty"`
	OldAltitude      *uint                    `json:"oldAltitude,omitempty"`
	NewFrequencyPlan *frequency_plan.BandName `json:"newFrequencyPlan,omitempty"`
	OldFrequencyPlan *frequency_plan.BandName `json:"oldFrequencyPlan,omitempty"`
	NewAntennaGain   *float32                 `json:"newAntennaGain,omitempty"`
	OldAntennaGain   *float32                 `json:"oldAntennaGain,omitempty"`
}
