/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Interest record
type UniLoanInterestRecord struct {
	// Currency name
	Currency string `json:"currency,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Actual rate
	ActualRate string `json:"actual_rate,omitempty"`
	// Interest
	Interest string `json:"interest,omitempty"`
	// Status: 0 - fail, 1 - success
	Status int32 `json:"status,omitempty"`
	// Type, platform - platform，margin - margin
	Type string `json:"type,omitempty"`
	// Created time
	CreateTime int64 `json:"create_time,omitempty"`
}
