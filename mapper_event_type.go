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
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
)

type MapperEventType string

const (
	// RegisterEvent is raised when a mapper is registered registry by ThingsIX
	MapperRegisteredEvent MapperEventType = "register"
	// OnboardEvent is raised when a mapper is onboarded by its owner
	MapperOnboardedEvent MapperEventType = "onboard"
	// ClaimEvent is raised when a mapper is claimed by ThingsIX
	MapperClaimedEvent MapperEventType = "claim"
	// RemovedEvent is raised when a mapper is removed from the registry
	MapperRemovedEvent MapperEventType = "removed"
	// MapperDeactivated is raised when ThingsIX marks a mapper as inactive
	MapperDeactivated MapperEventType = "deactivated"
	// MapperActive is raised when ThingsIX marks a mapper as active
	MapperActivated MapperEventType = "activated"
	// MapperTransfer is raised when the mapper is transferred to a different owner
	MapperTransfered                   = "transfer"
	MapperUnknownEvent MapperEventType = "unknown"
)

func (event *MapperEventType) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		*event = MapperEventType(str)
		return nil
	}
	if b, ok := value.([]byte); ok {
		*event = MapperEventType(b)
		return nil
	}
	return fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
}

func (event MapperEventType) Value() (driver.Value, error) {
	return string(event), nil
}

func (event MapperEventType) MarshalText() ([]byte, error) {
	return []byte(event), nil
}

func (event *MapperEventType) UnmarshalText(input []byte) error {
	normalized := strings.ToLower(string(input))
	switch MapperEventType(normalized) {
	case MapperRegisteredEvent, MapperOnboardedEvent, MapperClaimedEvent, MapperRemovedEvent, MapperDeactivated, MapperActivated, MapperTransfered:
		*event = MapperEventType(normalized)
		return nil
	default:
		return fmt.Errorf(`invalid event type "%s"`, input)
	}
}
