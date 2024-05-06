/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// RebateApiService RebateApi service
type RebateApiService service

// AgencyTransactionHistoryOpts Optional parameters for the method 'AgencyTransactionHistory'
type AgencyTransactionHistoryOpts struct {
	CurrencyPair optional.String
	UserId       optional.Int64
	From         optional.Int64
	To           optional.Int64
	Limit        optional.Int32
	Offset       optional.Int32
}

/*
AgencyTransactionHistory The agency obtains the transaction history of the recommended user
Record time range cannot exceed 30 days
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *AgencyTransactionHistoryOpts - Optional Parameters:
  - @param "CurrencyPair" (optional.String) -  Specify the currency pair, if not specified, return all currency pairs
  - @param "UserId" (optional.Int64) -  User ID. If not specified, all user records will be returned
  - @param "From" (optional.Int64) -  Time range beginning, default to 7 days before current time
  - @param "To" (optional.Int64) -  Time range ending, default to current time
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list
  - @param "Offset" (optional.Int32) -  List offset, starting from 0

@return []AgencyTransactionHistory
*/
func (a *RebateApiService) AgencyTransactionHistory(ctx context.Context, localVarOptionals *AgencyTransactionHistoryOpts) ([]AgencyTransactionHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AgencyTransactionHistory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rebate/agency/transaction_history"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CurrencyPair.IsSet() {
		localVarQueryParams.Add("currency_pair", parameterToString(localVarOptionals.CurrencyPair.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.From.IsSet() {
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.To.IsSet() {
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// AgencyCommissionsHistoryOpts Optional parameters for the method 'AgencyCommissionsHistory'
type AgencyCommissionsHistoryOpts struct {
	Currency optional.String
	UserId   optional.Int64
	From     optional.Int64
	To       optional.Int64
	Limit    optional.Int32
	Offset   optional.Int32
}

/*
AgencyCommissionsHistory The agency obtains the commission history of the recommended user
Record time range cannot exceed 30 days
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *AgencyCommissionsHistoryOpts - Optional Parameters:
  - @param "Currency" (optional.String) -  Filter by currency. Return all currency records if not specified
  - @param "UserId" (optional.Int64) -  User ID. If not specified, all user records will be returned
  - @param "From" (optional.Int64) -  Time range beginning, default to 7 days before current time
  - @param "To" (optional.Int64) -  Time range ending, default to current time
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list
  - @param "Offset" (optional.Int32) -  List offset, starting from 0

@return []AgencyCommissionHistory
*/
func (a *RebateApiService) AgencyCommissionsHistory(ctx context.Context, localVarOptionals *AgencyCommissionsHistoryOpts) ([]AgencyCommissionHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AgencyCommissionHistory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rebate/agency/commission_history"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarQueryParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.From.IsSet() {
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.To.IsSet() {
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RebateBrokerCommissionHistoryOpts Optional parameters for the method 'RebateBrokerCommissionHistory'
type RebateBrokerCommissionHistoryOpts struct {
	Limit  optional.Int32
	Offset optional.Int32
	UserId optional.Int64
}

/*
RebateBrokerCommissionHistory The broker obtains the user's commission rebate records
Record time range cannot exceed 30 days
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *RebateBrokerCommissionHistoryOpts - Optional Parameters:
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list
  - @param "Offset" (optional.Int32) -  List offset, starting from 0
  - @param "UserId" (optional.Int64) -  User ID. If not specified, all user records will be returned

@return []BrokerCommission
*/
func (a *RebateApiService) RebateBrokerCommissionHistory(ctx context.Context, localVarOptionals *RebateBrokerCommissionHistoryOpts) ([]BrokerCommission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []BrokerCommission
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rebate/broker/commission_history"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RebateBrokerTransactionHistoryOpts Optional parameters for the method 'RebateBrokerTransactionHistory'
type RebateBrokerTransactionHistoryOpts struct {
	Limit  optional.Int32
	Offset optional.Int32
	UserId optional.Int64
}

/*
RebateBrokerTransactionHistory The broker obtains the user's trading history
Record time range cannot exceed 30 days
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *RebateBrokerTransactionHistoryOpts - Optional Parameters:
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list
  - @param "Offset" (optional.Int32) -  List offset, starting from 0
  - @param "UserId" (optional.Int64) -  User ID. If not specified, all user records will be returned

@return []BrokerTransaction
*/
func (a *RebateApiService) RebateBrokerTransactionHistory(ctx context.Context, localVarOptionals *RebateBrokerTransactionHistoryOpts) ([]BrokerTransaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []BrokerTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rebate/broker/transaction_history"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
RebateUserInfo User retrieves rebate information
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []RebateUserInfo
*/
func (a *RebateApiService) RebateUserInfo(ctx context.Context) ([]RebateUserInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RebateUserInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v4/rebate/user/info"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
