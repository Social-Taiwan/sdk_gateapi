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

// EarnApiService EarnApi service
type EarnApiService service

/*
SwapETH2 ETH2 swap
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param eth2Swap
*/
func (a *EarnApiService) SwapETH2(ctx context.Context, eth2Swap Eth2Swap) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/staking/eth2/swap"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &eth2Swap
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
ListDualInvestmentPlans Dual Investment product list
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []DualGetPlans
*/
func (a *EarnApiService) ListDualInvestmentPlans(ctx context.Context) ([]DualGetPlans, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DualGetPlans
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/dual/investment_plan"
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
	ctx = context.WithValue(ctx, ContextPublic, true)
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
ListDualOrders Dual Investment order list
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []DualGetOrders
*/
func (a *EarnApiService) ListDualOrders(ctx context.Context) ([]DualGetOrders, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DualGetOrders
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/dual/orders"
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

/*
PlaceDualOrder Place Dual Investment order
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param placeDualInvestmentOrder
*/
func (a *EarnApiService) PlaceDualOrder(ctx context.Context, placeDualInvestmentOrder PlaceDualInvestmentOrder) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/dual/orders"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &placeDualInvestmentOrder
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ListStructuredProductsOpts Optional parameters for the method 'ListStructuredProducts'
type ListStructuredProductsOpts struct {
	Type_ optional.String
	Page  optional.Int32
	Limit optional.Int32
}

/*
ListStructuredProducts Structured Product List
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param status Status (default: all)  `in_process`-processing  `will_begin`-unstarted  `wait_settlement`-unsettled  `done`-finish
  - @param optional nil or *ListStructuredProductsOpts - Optional Parameters:
  - @param "Type_" (optional.String) -  Product Type (default all)  `SharkFin2.0`-SharkFin  `BullishSharkFin`-BullishSharkFin  `BearishSharkFin`-BearishSharkFin `DoubleNoTouch`-DoubleNoTouch `RangeAccrual`-RangeAccrual `SnowBall`-SnowBall
  - @param "Page" (optional.Int32) -  Page number
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list

@return []StructuredGetProjectList
*/
func (a *EarnApiService) ListStructuredProducts(ctx context.Context, status string, localVarOptionals *ListStructuredProductsOpts) ([]StructuredGetProjectList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []StructuredGetProjectList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/structured/products"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	localVarQueryParams.Add("status", parameterToString(status, ""))
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	ctx = context.WithValue(ctx, ContextPublic, true)
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

// ListStructuredOrdersOpts Optional parameters for the method 'ListStructuredOrders'
type ListStructuredOrdersOpts struct {
	From  optional.Int64
	To    optional.Int64
	Page  optional.Int32
	Limit optional.Int32
}

/*
ListStructuredOrders Structured Product Order List
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *ListStructuredOrdersOpts - Optional Parameters:
  - @param "From" (optional.Int64) -  Start timestamp
  - @param "To" (optional.Int64) -  End timestamp
  - @param "Page" (optional.Int32) -  Page number
  - @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list

@return []StructuredOrderList
*/
func (a *EarnApiService) ListStructuredOrders(ctx context.Context, localVarOptionals *ListStructuredOrdersOpts) ([]StructuredOrderList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []StructuredOrderList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/structured/orders"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.From.IsSet() {
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.To.IsSet() {
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
PlaceStructuredOrder Place Structured Product Order
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param structuredBuy
*/
func (a *EarnApiService) PlaceStructuredOrder(ctx context.Context, structuredBuy StructuredBuy) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/earn/structured/orders"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &structuredBuy
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
