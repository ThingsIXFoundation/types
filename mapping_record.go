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
	"crypto/sha256"
	"time"

	"github.com/ThingsIXFoundation/frequency-plan/go/frequency_plan"
	h3light "github.com/ThingsIXFoundation/h3-light"
)

type MappingRecord struct {
	ID                        ID                               `json:"id"`
	DiscoveryReceiptRecords   []*MappingDiscoveryReceiptRecord `json:"discoveryReceiptRecords"`
	DiscoveryPhy              []byte                           `json:"discoveryPhy"`
	DownlinkReceiptRecords    []*MappingDownlinkReceiptRecord  `json:"downlinkReceiptRecords"`
	DownlinkPhy               []byte                           `json:"downlinkPhy"`
	MeasuredRssi              *int                             `json:"measuredRssi"`
	MeasuredSnr               *int                             `json:"measuredSnr"`
	FrequencyPlan             frequency_plan.BandName          `json:"frequencyPlan"`
	ChallengedGatewayID       *ID                              `json:"challengedGatewayId"`
	ChallengedGatewayLocation *h3light.Cell                    `json:"challengedGatewayLocation"`
	ChallengedTime            *time.Time                       `json:"challengedTime"`
	MapperID                  ID                               `json:"mapperId"`
	MapperLocation            h3light.Cell                     `json:"mapperLocation"`
	MapperLat                 float64                          `json:"mapperLat"`
	MapperLon                 float64                          `json:"mapperLon"`
	MapperHeight              float64                          `json:"mapperHeight"`
	MapperOsnmaAge            uint8                            `json:"mapperOsnmaAge"`
	MapperSpoofing            uint8                            `json:"mapperSpoofing"`
	MapperTow                 uint32                           `json:"mapperTow"`
	MapperBattery             uint8                            `json:"mapperBattery"`
	MapperVersion             uint8                            `json:"mapperVersion"`
	MapperStatus              uint8                            `json:"mapperStatus"`
	ReceivedTime              time.Time                        `json:"receivedTime"`
	ServiceValidation         MappingRecordValidation          `json:"serviceValidation"`
}

// Return the MappingID of this MappingRecord, on the
func (mr *MappingRecord) MappingID() ID {
	return mappingID(mr.MapperID, mr.DownlinkPhy)
}

func mappingID(mapperID ID, phy []byte) ID {
	h := sha256.Sum256(append(mapperID[:], phy...))
	return ID(h)
}
