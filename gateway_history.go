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

type GatewayHistory struct {
	ContractAddress common.Address `json:"contract"`
	Block           common.Hash    `json:"blockHash"`
	Transaction     common.Hash    `json:"transaction"`
	BlockNumber     uint64         `json:"blockNumber"`
	Time            time.Time      `json:"time"`

	ID            ID                       `json:"id"`
	Version       uint8                    `json:"version"`
	Owner         *common.Address          `json:"owner"`
	AntennaGain   *float32                 `json:"antennaGain,omitempty"`
	FrequencyPlan *frequency_plan.BandName `json:"frequencyPlan,omitempty"`
	Location      *h3light.Cell            `json:"location,omitempty"`
	Altitude      *uint                    `json:"altitude,omitempty"`
}

func (gh *GatewayHistory) Gateway() *Gateway {
	return &Gateway{
		ID:              gh.ID,
		ContractAddress: gh.ContractAddress,
		Version:         gh.Version,
		Owner:           *gh.Owner,
		AntennaGain:     gh.AntennaGain,
		FrequencyPlan:   gh.FrequencyPlan,
		Location:        gh.Location,
		Altitude:        gh.Altitude,
	}
}
