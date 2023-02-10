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
	"github.com/ThingsIXFoundation/frequency-plan/go/frequency_plan"
	h3light "github.com/ThingsIXFoundation/h3-light"
	"github.com/ethereum/go-ethereum/common"
)

type Gateway struct {
	// ID is the ThingsIX compressed public key for this gateway
	ID              ID                       `json:"id" gorm:"primaryKey;type:bytea"`
	ContractAddress common.Address           `json:"contract"`
	Version         uint8                    `json:"version"`
	Owner           common.Address           `json:"owner" gorm:"index;type:bytea"`
	AntennaGain     *float32                 `json:"antennaGain,omitempty"`
	FrequencyPlan   *frequency_plan.BandName `json:"frequencyPlan,omitempty"`
	Location        *h3light.Cell            `json:"location,omitempty"`
	Altitude        *uint                    `json:"altitude,omitempty"`
}
