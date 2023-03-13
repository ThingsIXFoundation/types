package types

import (
	"time"

	h3light "github.com/ThingsIXFoundation/h3-light"
)

// AssumedCoverageHistory stores the history (and therefore also current coverage) as assumed based in calculated coverage
// As not every location can be mapped, based on the actually mapped locations, a assumed coverage is calculated
// There can be multiple AssumedCoverageHistory's for the same location for different gateways.
type AssumedCoverageHistory struct {
	// Res 8 cell of the location of the coverage
	Location h3light.Cell `json:"location"`

	// The coverage as being provided by one or multiple gateways
	GatewayCoverage []*AssumedGatewayCoverageHistory `json:"gatewayCoverage"`

	// Date this coverage was (assumed to be) present based on the measurements
	Date time.Time `json:"date"`
}

type AssumedGatewayCoverageHistory struct {

	// ID of the gateway that provides this coverage
	GatewayID ID `json:"gatewayId"`

	// The number of (res10) coverage records this gateway actually has within this (res8) cell
	NumCoverage uint64 `json:"numCoverage"`

	// The share of total all coverage records this gateway in this res8 cell, all shares for different gateways together must be 1000.
	Share uint64 `json:"share"`
}
