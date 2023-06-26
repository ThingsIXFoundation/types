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
)

// CoverageHistory stores the history (and therefore also current coverage) provided by the network
type CoverageHistory struct {
	// Res 10 cell of the location of the coverage
	Location h3light.Cell `json:"location"`

	// Date this coverage was (assumed to be) present based on the measurements
	Date time.Time `json:"date"`
	// ID of the gateway that provides this coverage
	GatewayID ID `json:"gatewayId"`

	// The location of the gateway that provides this coverage
	GatewayLocation h3light.Cell `json:"gatewayLocation"`

	// The frequency plan that coverage is provided for
	FrequencyPlan frequency_plan.BandName `json:"frequencyPlan"`

	// ID of the mapper that mapped this coverage
	MapperID ID `json:"mapperId"`

	// ID of the mapping record that was used to base this coverage on
	MappingID ID `json:"mappingId"`

	// Time of the mapping record that was used to base this coverage on
	MappingTime time.Time `json:"mappingTime"`

	// The RSSI (signal strength) of coverage at this location
	RSSI int `json:"rssi"`
}
