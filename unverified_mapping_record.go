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

	h3light "github.com/ThingsIXFoundation/h3-light"
)

type UnverifiedMappingRecord struct {
	ID                  ID                                `json:"id"`
	GatewayRecords      []*UnverifiedMappingGatewayRecord `json:"gatewayRecords"`
	MapperID            string                            `json:"mapperId"`
	BestGatewayID       *ID                               `json:"bestGatewayId"`
	BestGatewayLocation *h3light.Cell                     `json:"bestGatewayLocation"`
	BestGatewayRssi     *int32                            `json:"bestGatewayRssi"`
	BestGatewaySnr      *float64                          `json:"bestGatewaySnr"`
	MapperLocation      h3light.Cell                      `json:"mapperLocation"`
	MapperLat           float64                           `json:"mapperLat"`
	MapperLon           float64                           `json:"mapperLon"`
	MapperAccuracy      float64                           `json:"mapperAccuracy"`
	MapperHeight        float64                           `json:"mapperHeight"`
	ReceivedTime        time.Time                         `json:"receivedTime"`
	Frequency           uint32                            `json:"frequency,omitempty"`
	SpreadingFactor     uint32                            `json:"spreadingFactor,omitempty"`
	Bandwidth           uint32                            `json:"bandwidth,omitempty"`
	CodeRate            string                            `json:"codeRate,omitempty"`
}

type UnverifiedMappingGatewayRecord struct {
	MappingID       ID            `json:"mappingId"`
	GatewayID       ID            `json:"gatewayId,omitempty"`
	GatewayLocation *h3light.Cell `json:"gatewayLocation,omitempty"`
	GatewayTime     time.Time     `json:"gatewayTime,omitempty"`
	Rssi            int32         `json:"rssi,omitempty"`
	Snr             float64       `json:"snr,omitempty"`
}
