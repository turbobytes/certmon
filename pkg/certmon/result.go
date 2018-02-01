package certmon

import (
	"crypto/tls"
	"time"
)

//Results is the final output
type Results struct {
	Results   []Result  `json:"results"`
	Timestamp time.Time `json:"timestamp"`
}

//Result is the output for individual check
type Result struct {
	Hostname  string                      `json:"hostname"` //The value in SNI and certificate will be validated against it.
	Endpoints map[string]IndividualResult `json:"endpoints"`
	Timestamp time.Time                   `json:"timestamp"`
}

//IndividualResult is the output of checking individual endpoint
type IndividualResult struct {
	ConnectionState *tls.ConnectionState `json:"connectionstate"`
	Expiry          time.Time            `json:"expiry"`
	OK              bool                 `json:"ok"`
	Err             string               `json:"err"`
	Timestamp       time.Time            `json:"timestamp"`
}
