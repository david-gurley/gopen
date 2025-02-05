/*
 * Tokenauth API reference
 *
 * APIs to generate node auth tokens  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// TokenauthV1ApiService TokenauthV1Api service
type TokenauthV1ApiService service

type ApiTokenauthGetGenerateNodeTokenRequest struct {
	ctx _context.Context
	ApiService *TokenauthV1ApiService
	audience *[]string
	validityStart *time.Time
}

func (r ApiTokenauthGetGenerateNodeTokenRequest) Audience(audience []string) ApiTokenauthGetGenerateNodeTokenRequest {
	r.audience = &audience
	return r
}
func (r ApiTokenauthGetGenerateNodeTokenRequest) ValidityStart(validityStart time.Time) ApiTokenauthGetGenerateNodeTokenRequest {
	r.validityStart = &validityStart
	return r
}

func (r ApiTokenauthGetGenerateNodeTokenRequest) Execute() (TokenauthNodeTokenResponse, *_nethttp.Response, error) {
	return r.ApiService.GetGenerateNodeTokenExecute(r)
}

/*
 * GetGenerateNodeToken Method for GetGenerateNodeToken
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTokenauthGetGenerateNodeTokenRequest
 */
func (a *TokenauthV1ApiService) GetGenerateNodeToken(ctx _context.Context) ApiTokenauthGetGenerateNodeTokenRequest {
	return ApiTokenauthGetGenerateNodeTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TokenauthNodeTokenResponse
 */
func (a *TokenauthV1ApiService) GetGenerateNodeTokenExecute(r ApiTokenauthGetGenerateNodeTokenRequest) (TokenauthNodeTokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenauthNodeTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenauthV1ApiService.GetGenerateNodeToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenauth/v1/node"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.audience != nil {
		localVarQueryParams.Add("audience", parameterToString(*r.audience, "csv"))
	}
	if r.validityStart != nil {
		localVarQueryParams.Add("validity-start", parameterToString(*r.validityStart, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.GetGenerateNodeTokenExecute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
