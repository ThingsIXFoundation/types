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

import "time"

type MappingDiscoveryReceiptRecord struct {
	Frequency         uint32                  `json:"frequency,omitempty"`
	Rssi              int32                   `json:"rssi,omitempty"`
	Snr               float64                 `json:"snr,omitempty"`
	SpreadingFactor   uint32                  `json:"spreadingFactor,omitempty"`
	Bandwidth         uint32                  `json:"bandwidth,omitempty"`
	CodeRate          string                  `json:"codeRate,omitempty"`
	Phy               []byte                  `json:"phy,omitempty"`
	ReceivedTime      time.Time               `json:"receivedTime,omitempty"`
	GatewayTime       time.Time               `json:"gatewayTime,omitempty"`
	GatewaySignature  []byte                  `json:"gatewaySignature,omitempty"`
	GatewayID         ID                      `json:"gatewayId,omitempty"`
	MapperID          ID                      `json:"mapperId,omitempty"`
	ServiceValidation MappingRecordValidation `json:"serviceValidation"`
}

// Return the
func (mdrr *MappingDiscoveryReceiptRecord) MappingID() ID {
	return mappingID(mdrr.MapperID, mdrr.Phy)
}
