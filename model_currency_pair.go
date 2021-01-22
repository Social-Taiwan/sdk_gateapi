/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Spot currency pair
type CurrencyPair struct {
	// Currency pair
	Id string `json:"id,omitempty"`
	// Base currency
	Base string `json:"base,omitempty"`
	// Quote currency
	Quote string `json:"quote,omitempty"`
	// Trading fee
	Fee string `json:"fee,omitempty"`
	// Minimum amount of base currency to trade, `null` means no limit
	MinBaseAmount string `json:"min_base_amount,omitempty"`
	// Minimum amount of quote currency to trade, `null` means no limit
	MinQuoteAmount string `json:"min_quote_amount,omitempty"`
	// Amount scale
	AmountPrecision int32 `json:"amount_precision,omitempty"`
	// Price scale
	Precision int32 `json:"precision,omitempty"`
	// How currency pair can be traded  - untradable: cannot be bought or sold - buyable: can be bought - sellable: can be sold - tradable: can be bought or sold
	TradeStatus string `json:"trade_status,omitempty"`
	// ETF net value
	EtfNetValue string `json:"etf_net_value,omitempty"`
	// ETF previous net value at re-balancing time
	EtfPreNetValue string `json:"etf_pre_net_value,omitempty"`
	// ETF previous re-balancing time
	EtfPreTimestamp int64 `json:"etf_pre_timestamp,omitempty"`
	// ETF current leverage
	EtfLeverage string `json:"etf_leverage,omitempty"`
}
