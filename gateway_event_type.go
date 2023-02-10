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
	"fmt"
	"reflect"
	"strings"
)

type GatewayEventType string

const (
	GatewayOnboardedEvent   GatewayEventType = "onboard"
	GatewayOffboardedEvent  GatewayEventType = "offboard"
	GatewayUpdatedEvent     GatewayEventType = "update"
	GatewayTransferredEvent GatewayEventType = "transfer"
	GatewayUnknownEvent     GatewayEventType = "unknown"
)

func (event *GatewayEventType) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		*event = GatewayEventType(str)
		return nil
	}

	if b, ok := value.([]byte); ok {
		*event = GatewayEventType(b)
		return nil
	}
	return fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
}

func (event GatewayEventType) MarshalText() ([]byte, error) {
	return []byte(event), nil
}

func (event *GatewayEventType) UnmarshalText(input []byte) error {
	normalized := strings.ToLower(string(input))
	switch GatewayEventType(normalized) {
	case GatewayOnboardedEvent, GatewayOffboardedEvent, GatewayUpdatedEvent, GatewayTransferredEvent:
		*event = GatewayEventType(normalized)
		return nil
	default:
		return fmt.Errorf(`invalid event type "%s"`, input)
	}
}
