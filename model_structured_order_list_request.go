/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Structured Request Parameters
type StructuredOrderListRequest struct {
	// start time
	From int32 `json:"from"`
	// Finished time
	To int32 `json:"to"`
	// Page number
	Page int32 `json:"page,omitempty"`
	// Number of items returned in the list. Default is 100.
	Limit int32 `json:"limit,omitempty"`
}
