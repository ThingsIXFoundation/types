package types

import (
	"time"

	h3light "github.com/ThingsIXFoundation/h3-light"
)

// CoverageHistory stores the history (and therefore also current coverage) provided by the network
type CoverageHistory struct {
	// Res 10 cell of the location of the coverage
	Location h3light.Cell `json:"location"`

	// Date this coverage was (assumed to be) present based on the measurements
	Date time.Time `json:"date"`
	// ID of the gateway that provides this coverage
	GatewayID ID `json:"gatewayId"`

	// ID of the mapper that mapped this coverage
	MapperID ID `json:"mapperId"`

	// ID of the mapping record that was used to base this coverage on
	MappingID ID `json:"mappingId"`

	// The RSSI (signal strength) of coverage at this location
	RSSI int `json:"rssi"`
}
