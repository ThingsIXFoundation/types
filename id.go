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
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type ID [32]byte

func (id ID) String() string {
	return fmt.Sprintf("0x%x", id[:])
}

func (id ID) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

func (id *ID) UnmarshalText(input []byte) error {
	if len(input) != 64 && len(input) != 66 {
		return fmt.Errorf("id has invalid length")
	}

	if len(input) == 66 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X') {
		input = input[2:]
	} else if len(input) != 64 {
		return fmt.Errorf("id doesn't start with 0x")
	}

	decoded, err := hex.DecodeString(string(input))
	if err != nil {
		return fmt.Errorf("invalid id")
	}
	_ = copy(id[:], decoded)
	return nil
}

func IDFromString(s string) ID {
	id := ID{}
	_ = id.UnmarshalText([]byte(s))
	return id
}

func IDFromPublicKeyBytes(publicKeyBytes []byte) ID {
	id := ID{}
	_ = copy(id[:], publicKeyBytes[1:33])
	return id
}

func IDFromBytes(b []byte) ID {
	id := ID{}
	_ = copy(id[:], b[0:32])
	return id
}

func IDFromRandom() ID {
	id := ID{}
	rand.Read(id[:])

	return id
}
