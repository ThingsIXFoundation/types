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

	"github.com/ethereum/go-ethereum/common"
)

type RouterEvent struct {
	ContractAddress  common.Address `json:"contract"`
	Block            common.Hash    `json:"blockHash"`
	Transaction      common.Hash    `json:"transaction"`
	BlockNumber      uint64         `json:"blockNumber"`
	TransactionIndex uint           `json:"transactionIndex"`
	LogIndex         uint           `json:"logIndex"`
	Time             time.Time      `json:"time"`

	Type        RouterEventType `json:"type"`
	ID          ID              `json:"id"`
	Revision    uint16          `json:"revision"`
	Owner       *common.Address `json:"owner"`
	NewNetID    uint32          `json:"newNetid"`
	OldNetID    uint32          `json:"oldNetid"`
	NewPrefix   uint32          `json:"newPrefix"`
	OldPrefix   uint32          `json:"oldPrefix"`
	NewMask     uint8           `json:"newMask"`
	OldMask     uint8           `json:"oldMask"`
	NewEndpoint string          `json:"newEndpoint"`
	OldEndpoint string          `json:"oldEndpoint"`
}
