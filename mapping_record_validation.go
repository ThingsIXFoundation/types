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

type MappingRecordValidation string

const (
	// All validations succeeded, note that later validations can still fail
	MappingRecordValidationOk MappingRecordValidation = "ok"

	// No discovery receipt was received
	MappingRecordValidationNoDiscoveryReceipt MappingRecordValidation = "no-discovery-receipt"

	// No downlink was transmitted (in time)
	MappingRecordValidationNoDownlink MappingRecordValidation = "no-downlink"

	// No downlink receipt was received (in time)
	MappingRecordValidationNoDownlinkReceipt MappingRecordValidation = "no-downlink-receipt"

	// The receipt was received on a wrong channel/frequency
	MappingRecordValidationWrongFrequency MappingRecordValidation = "wrong-frequency"

	// The receipt was received with a wrong spreading factor
	MappingRecordValidationWrongSpreadingFactor MappingRecordValidation = "wrong-spreading-factor"

	// There was a mismatch between the frequency-plan and the location of the mapper and/or of the receiving gateway
	MappingRecordValidationWrongFrequencyPlan MappingRecordValidation = "wrong-frequency-plan"

	// The distance between the mapper and the receiving is too large
	MappingRecordValidationTooLargeDistance MappingRecordValidation = "too-large-distance"

	// The RSSI and/or SNR are unlikely for the distance between the mapper and the receiving gateway
	MappingRecordValidationUnlikelyRssiSnr MappingRecordValidation = "unlikely-rssi-snr"

	// The RSSI is out of the acceptable bounds (too high or too low)
	MappingRecordValidationRssiOutOfBounds MappingRecordValidation = "rssi-out-of-bounds"

	// The receiving gateway is (temporarily) denylisted due to previous suspicious mapping-values
	MappingRecordValidationGatewayDenylisted MappingRecordValidation = "gateway-denylisted"

	// The receiving gateway is blocklisted due to previous suspicious mapping-values
	MappingRecordValidationGatewayBlocklisted MappingRecordValidation = "gateway-blocklisted"

	// The mpper is (temporarily) denylisted due to previous suspicious mapping-values
	MappingRecordValidationMapperTemporarilyDenylisted MappingRecordValidation = "mapper-temp-denylisted"

	// The mapper is blocklisted due to previous suspicious mapping-values
	MappingRecordValidationMapperBlocklisted MappingRecordValidation = "mapper-blocklisted"

	// The mapper is blocklisted due to previous suspicious mapping-values
	MappingRecordValidationMapperNotActive MappingRecordValidation = "mapper-not-active"

	// The mapper record was too old once it was received and delivered to the service
	MappingRecordValidationTooOld MappingRecordValidation = "too-old"

	// The location of the mapper or receiving gateway was invalid
	MappingRecordValidationInvalidLocation MappingRecordValidation = "invalid-location"

	// The location of the mapper was too high or low with respect to the earth surface
	MappingRecordValidationInvalidHeight MappingRecordValidation = "invalid-height"

	// The GPS-security validation by the mapper was invalid
	MappingRecordValidationInvalidGpsSecurity MappingRecordValidation = "invalid-gps-security"

	// The OSNMA validation by the mapper was not sufficient
	MappingRecordValidationInsufficientGpsSecurity MappingRecordValidation = "insufficient-gps-security"

	// An internal server error occcured during validation
	MappingRecordValidationInternalServerError MappingRecordValidation = "internal-server-error"
)
