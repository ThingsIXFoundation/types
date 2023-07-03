package types

import (
	"time"

	h3light "github.com/ThingsIXFoundation/h3-light"
)

type AssumedUnverifiedCoverage struct {
	// Res 8 cell of the location of the coverage
	Location h3light.Cell `json:"location"`

	LatestUpdate time.Time `json:"latestUpdate"`
}
