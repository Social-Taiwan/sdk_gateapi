/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CollateralAdjust struct {
	// Order ID
	OrderId int64 `json:"order_id"`
	// Operation types: append - for adding collateral, redeem - for withdrawing collateral
	Type string `json:"type"`
	// Collateral Currency List
	Collaterals []CollateralCurrency `json:"collaterals,omitempty"`
}
