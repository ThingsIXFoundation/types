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

type RouterEventType string

const (
	// RouterRegisteredEvent is raised when a router is registered by its manager
	RouterRegisteredEvent RouterEventType = "register"
	// RegisterEvent is raised when a router details are updated
	RouterUpdatedEvent RouterEventType = "update"
	// RouterRemovedEvent is raised when router details are removed
	RouterRemovedEvent RouterEventType = "removed"

	RouterUnknownEvent RouterEventType = "unknown"
)

func (event *RouterEventType) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		*event = RouterEventType(str)
		return nil
	}
	if b, ok := value.([]byte); ok {
		*event = RouterEventType(b)
		return nil
	}
	return fmt.Errorf("unsupported type: %v", reflect.TypeOf(value))
}

func (event RouterEventType) Value() (driver.Value, error) {
	return string(event), nil
}

func (event RouterEventType) MarshalText() ([]byte, error) {
	return []byte(event), nil
}

func (event *RouterEventType) UnmarshalText(input []byte) error {
	normalized := strings.ToLower(string(input))
	switch RouterEventType(normalized) {
	case RouterRegisteredEvent, RouterUpdatedEvent, RouterRemovedEvent:
		*event = RouterEventType(normalized)
		return nil
	default:
		return fmt.Errorf(`invalid event type "%s"`, input)
	}
}
