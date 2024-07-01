/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Borrow or repay
type PortfolioLoan struct {
	// Currency
	Currency string `json:"currency"`
	// type: borrow - borrow, repay - repay
	Type string `json:"type"`
	// The amount of lending or repaying
	Amount string `json:"amount"`
	// Full repayment.  Repay operation only.  If the value is `true`, the amount will be ignored and the loan will be repaid in full.
	RepaidAll bool `json:"repaid_all,omitempty"`
	// User defined custom ID
	Text string `json:"text,omitempty"`
}
