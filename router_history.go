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
	"github.com/ethereum/go-ethereum/common"
)

type RouterHistory struct {
	ContractAddress common.Address `json:"contract"`
	Block           common.Hash    `json:"blockHash"`
	Transaction     common.Hash    `json:"transaction"`
	BlockNumber     uint64         `json:"blockNumber"`
	Time            time.Time      `json:"time"`

	ID            ID                      `json:"id" gorm:"primaryKey;type:bytea"`
	Owner         *common.Address         `json:"owner"`
	NetID         uint32                  `json:"netid"`
	Prefix        uint32                  `json:"prefix"`
	Mask          uint8                   `json:"mask"`
	FrequencyPlan frequency_plan.BandName `json:"frequencyPlan"`
	Endpoint      string                  `json:"endpoint"`
}

func (rh *RouterHistory) Router() *Router {
	return &Router{
		ID:              rh.ID,
		ContractAddress: rh.ContractAddress,
		Owner:           *rh.Owner,
		NetID:           rh.NetID,
		Prefix:          rh.Prefix,
		Mask:            rh.Mask,
		FrequencyPlan:   rh.FrequencyPlan,
		Endpoint:        rh.Endpoint,
	}
}
