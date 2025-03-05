/**
 * (C) Copyright IBM Corp. 2024-2025.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.99.0-d27cee72-20250129-204831
 */

// Package sdsaasv1 : Operations and models for the SdsaasV1 service
package sdsaasv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/sds-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// SdsaasV1 : OpenAPI definition for SDSaaS
//
// API Version: 1.0.0
type SdsaasV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "sdsaas"

const ParameterizedServiceURL = "{url}"

var defaultUrlVariables = map[string]string{
	"url": "{url}",
}

// SdsaasV1Options : Service options
type SdsaasV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSdsaasV1UsingExternalConfig : constructs an instance of SdsaasV1 with passed in options and external configuration.
func NewSdsaasV1UsingExternalConfig(options *SdsaasV1Options) (sdsaas *SdsaasV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	sdsaas, err = NewSdsaasV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = sdsaas.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = sdsaas.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewSdsaasV1 : constructs an instance of SdsaasV1 with passed in options.
func NewSdsaasV1(options *SdsaasV1Options) (service *SdsaasV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &SdsaasV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "sdsaas" suitable for processing requests.
func (sdsaas *SdsaasV1) Clone() *SdsaasV1 {
	if core.IsNil(sdsaas) {
		return nil
	}
	clone := *sdsaas
	clone.Service = sdsaas.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (sdsaas *SdsaasV1) SetServiceURL(url string) error {
	err := sdsaas.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (sdsaas *SdsaasV1) GetServiceURL() string {
	return sdsaas.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (sdsaas *SdsaasV1) SetDefaultHeaders(headers http.Header) {
	sdsaas.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (sdsaas *SdsaasV1) SetEnableGzipCompression(enableGzip bool) {
	sdsaas.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (sdsaas *SdsaasV1) GetEnableGzipCompression() bool {
	return sdsaas.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (sdsaas *SdsaasV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	sdsaas.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (sdsaas *SdsaasV1) DisableRetries() {
	sdsaas.Service.DisableRetries()
}

// Volumes : This request lists all volumes in the region
// Volumes are network-connected block storage devices that may be attached to one or more instances in the same region.
func (sdsaas *SdsaasV1) Volumes(volumesOptions *VolumesOptions) (result *VolumeCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.VolumesWithContext(context.Background(), volumesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// VolumesWithContext is an alternate form of the Volumes method which supports a Context parameter
func (sdsaas *SdsaasV1) VolumesWithContext(ctx context.Context, volumesOptions *VolumesOptions) (result *VolumeCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(volumesOptions, "volumesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range volumesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Volumes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if volumesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*volumesOptions.Limit))
	}
	if volumesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*volumesOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "volumes", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// VolumeCreate : Create a new volume
// Create a volume.
func (sdsaas *SdsaasV1) VolumeCreate(volumeCreateOptions *VolumeCreateOptions) (result *Volume, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.VolumeCreateWithContext(context.Background(), volumeCreateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// VolumeCreateWithContext is an alternate form of the VolumeCreate method which supports a Context parameter
func (sdsaas *SdsaasV1) VolumeCreateWithContext(ctx context.Context, volumeCreateOptions *VolumeCreateOptions) (result *Volume, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(volumeCreateOptions, "volumeCreateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(volumeCreateOptions, "volumeCreateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range volumeCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "VolumeCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if volumeCreateOptions.Capacity != nil {
		body["capacity"] = volumeCreateOptions.Capacity
	}
	if volumeCreateOptions.Name != nil {
		body["name"] = volumeCreateOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "volume_create", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolume)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// Volume : Retrieve a volume profile
// This request retrieves a single volume profile specified by ID.
func (sdsaas *SdsaasV1) Volume(volumeOptions *VolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.VolumeWithContext(context.Background(), volumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// VolumeWithContext is an alternate form of the Volume method which supports a Context parameter
func (sdsaas *SdsaasV1) VolumeWithContext(ctx context.Context, volumeOptions *VolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(volumeOptions, "volumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(volumeOptions, "volumeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"volume_id": *volumeOptions.VolumeID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{volume_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range volumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Volume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "volume", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolume)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// VolumeDelete : Delete a volume
// This request deletes a single volume profile based on the name.
func (sdsaas *SdsaasV1) VolumeDelete(volumeDeleteOptions *VolumeDeleteOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.VolumeDeleteWithContext(context.Background(), volumeDeleteOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// VolumeDeleteWithContext is an alternate form of the VolumeDelete method which supports a Context parameter
func (sdsaas *SdsaasV1) VolumeDeleteWithContext(ctx context.Context, volumeDeleteOptions *VolumeDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(volumeDeleteOptions, "volumeDeleteOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(volumeDeleteOptions, "volumeDeleteOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"volume_id": *volumeDeleteOptions.VolumeID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{volume_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range volumeDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "VolumeDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "volume_delete", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// VolumeUpdate : Update a volume
// This request updates a volume with the information in a provided volume patch.
func (sdsaas *SdsaasV1) VolumeUpdate(volumeUpdateOptions *VolumeUpdateOptions) (result *Volume, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.VolumeUpdateWithContext(context.Background(), volumeUpdateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// VolumeUpdateWithContext is an alternate form of the VolumeUpdate method which supports a Context parameter
func (sdsaas *SdsaasV1) VolumeUpdateWithContext(ctx context.Context, volumeUpdateOptions *VolumeUpdateOptions) (result *Volume, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(volumeUpdateOptions, "volumeUpdateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(volumeUpdateOptions, "volumeUpdateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"volume_id": *volumeUpdateOptions.VolumeID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{volume_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range volumeUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "VolumeUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	if volumeUpdateOptions.VolumePatch != nil {
		_, err = builder.SetBodyContentJSON(volumeUpdateOptions.VolumePatch)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "volume_update", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolume)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// Creds : List storage account credentials
// This request retrieves credentials for a specific storage account.
func (sdsaas *SdsaasV1) Creds(credsOptions *CredsOptions) (result *CredentialsFound, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CredsWithContext(context.Background(), credsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CredsWithContext is an alternate form of the Creds method which supports a Context parameter
func (sdsaas *SdsaasV1) CredsWithContext(ctx context.Context, credsOptions *CredsOptions) (result *CredentialsFound, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(credsOptions, "credsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/s3_credentials`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range credsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Creds")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "creds", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentialsFound)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CredCreate : Create or modify storage account credentials
// Updates credentials for a storage account or creates them if they do not exist.
func (sdsaas *SdsaasV1) CredCreate(credCreateOptions *CredCreateOptions) (result *CredentialsUpdated, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CredCreateWithContext(context.Background(), credCreateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CredCreateWithContext is an alternate form of the CredCreate method which supports a Context parameter
func (sdsaas *SdsaasV1) CredCreateWithContext(ctx context.Context, credCreateOptions *CredCreateOptions) (result *CredentialsUpdated, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(credCreateOptions, "credCreateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(credCreateOptions, "credCreateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"access_key": *credCreateOptions.AccessKey,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/s3_credentials/{access_key}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range credCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CredCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "cred_create", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentialsUpdated)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CredDelete : Delete storage account credentials
// Deletes specific credentials for a storage account.
func (sdsaas *SdsaasV1) CredDelete(credDeleteOptions *CredDeleteOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.CredDeleteWithContext(context.Background(), credDeleteOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CredDeleteWithContext is an alternate form of the CredDelete method which supports a Context parameter
func (sdsaas *SdsaasV1) CredDeleteWithContext(ctx context.Context, credDeleteOptions *CredDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(credDeleteOptions, "credDeleteOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(credDeleteOptions, "credDeleteOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"access_key": *credDeleteOptions.AccessKey,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/s3_credentials/{access_key}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range credDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CredDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "cred_delete", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CertTypes : List the allowed certificate types
// Retrieves the list configured certificates.
func (sdsaas *SdsaasV1) CertTypes(certTypesOptions *CertTypesOptions) (result *CertificateList, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CertTypesWithContext(context.Background(), certTypesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CertTypesWithContext is an alternate form of the CertTypes method which supports a Context parameter
func (sdsaas *SdsaasV1) CertTypesWithContext(ctx context.Context, certTypesOptions *CertTypesOptions) (result *CertificateList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(certTypesOptions, "certTypesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range certTypesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CertTypes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "cert_types", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertificateList)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// Cert : Retrieves the SSL certificate expiration date and status
// This request retrieves the SSL certificate expiration date and status.
func (sdsaas *SdsaasV1) Cert(certOptions *CertOptions) (result *CertificateFound, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CertWithContext(context.Background(), certOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CertWithContext is an alternate form of the Cert method which supports a Context parameter
func (sdsaas *SdsaasV1) CertWithContext(ctx context.Context, certOptions *CertOptions) (result *CertificateFound, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(certOptions, "certOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(certOptions, "certOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert": *certOptions.Cert,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range certOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Cert")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "cert", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertificateFound)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CertDelete : Delete SSL certificate
// Delete the provided PEM formatted TLS certificate.
func (sdsaas *SdsaasV1) CertDelete(certDeleteOptions *CertDeleteOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.CertDeleteWithContext(context.Background(), certDeleteOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CertDeleteWithContext is an alternate form of the CertDelete method which supports a Context parameter
func (sdsaas *SdsaasV1) CertDeleteWithContext(ctx context.Context, certDeleteOptions *CertDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(certDeleteOptions, "certDeleteOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(certDeleteOptions, "certDeleteOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert": *certDeleteOptions.Cert,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range certDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CertDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "cert_delete", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CertCreate : Creates a new SSL Certificate
// Creates a new SSL Certificate.
func (sdsaas *SdsaasV1) CertCreate(certCreateOptions *CertCreateOptions) (result *CertificateUpdated, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CertCreateWithContext(context.Background(), certCreateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CertCreateWithContext is an alternate form of the CertCreate method which supports a Context parameter
func (sdsaas *SdsaasV1) CertCreateWithContext(ctx context.Context, certCreateOptions *CertCreateOptions) (result *CertificateUpdated, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(certCreateOptions, "certCreateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(certCreateOptions, "certCreateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert": *certCreateOptions.Cert,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range certCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CertCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/octect-stream")

	_, err = builder.SetBodyContent("application/octect-stream", nil, nil, certCreateOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "cert_create", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertificateUpdated)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CertUpdate : Updates the SSL Certificate
// Updates the SSL Certificate.
func (sdsaas *SdsaasV1) CertUpdate(certUpdateOptions *CertUpdateOptions) (result *CertificateUpdated, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CertUpdateWithContext(context.Background(), certUpdateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CertUpdateWithContext is an alternate form of the CertUpdate method which supports a Context parameter
func (sdsaas *SdsaasV1) CertUpdateWithContext(ctx context.Context, certUpdateOptions *CertUpdateOptions) (result *CertificateUpdated, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(certUpdateOptions, "certUpdateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(certUpdateOptions, "certUpdateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert": *certUpdateOptions.Cert,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range certUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "CertUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/octect-stream")

	_, err = builder.SetBodyContent("application/octect-stream", nil, nil, certUpdateOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "cert_update", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertificateUpdated)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// Hosts : Lists all hosts and all host IDs
// This request lists all hosts in the deployment. Hosts are objects representing the NVMe initiators that may be mapped
// to one or more volumes in the same deployment.
func (sdsaas *SdsaasV1) Hosts(hostsOptions *HostsOptions) (result *HostCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostsWithContext(context.Background(), hostsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostsWithContext is an alternate form of the Hosts method which supports a Context parameter
func (sdsaas *SdsaasV1) HostsWithContext(ctx context.Context, hostsOptions *HostsOptions) (result *HostCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(hostsOptions, "hostsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Hosts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if hostsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*hostsOptions.Limit))
	}
	if hostsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*hostsOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "hosts", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHostCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostCreate : Creates a host
// This request creates a new host from a host template object.
func (sdsaas *SdsaasV1) HostCreate(hostCreateOptions *HostCreateOptions) (result *Host, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostCreateWithContext(context.Background(), hostCreateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostCreateWithContext is an alternate form of the HostCreate method which supports a Context parameter
func (sdsaas *SdsaasV1) HostCreateWithContext(ctx context.Context, hostCreateOptions *HostCreateOptions) (result *Host, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostCreateOptions, "hostCreateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostCreateOptions, "hostCreateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if hostCreateOptions.Nqn != nil {
		body["nqn"] = hostCreateOptions.Nqn
	}
	if hostCreateOptions.Name != nil {
		body["name"] = hostCreateOptions.Name
	}
	if hostCreateOptions.VolumeMappings != nil {
		body["volume_mappings"] = hostCreateOptions.VolumeMappings
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_create", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHost)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// Host : Retrieve a host by ID
// This request retrieves a host specified by the host ID.
func (sdsaas *SdsaasV1) Host(hostOptions *HostOptions) (result *Host, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostWithContext(context.Background(), hostOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostWithContext is an alternate form of the Host method which supports a Context parameter
func (sdsaas *SdsaasV1) HostWithContext(ctx context.Context, hostOptions *HostOptions) (result *Host, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostOptions, "hostOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostOptions, "hostOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "Host")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHost)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostUpdate : Update a host
// This request updates a Host with the information in a provided host patch object. The host patch object is structured
// in the same way as a retrieved host and contains only the information to be updated.
func (sdsaas *SdsaasV1) HostUpdate(hostUpdateOptions *HostUpdateOptions) (result *Host, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostUpdateWithContext(context.Background(), hostUpdateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostUpdateWithContext is an alternate form of the HostUpdate method which supports a Context parameter
func (sdsaas *SdsaasV1) HostUpdateWithContext(ctx context.Context, hostUpdateOptions *HostUpdateOptions) (result *Host, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostUpdateOptions, "hostUpdateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostUpdateOptions, "hostUpdateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostUpdateOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	if hostUpdateOptions.HostPatch != nil {
		_, err = builder.SetBodyContentJSON(hostUpdateOptions.HostPatch)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_update", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHost)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostDelete : Delete a specific host
// This request deletes a host. For this request to succeed, the host must not be mapped to any volumes.
func (sdsaas *SdsaasV1) HostDelete(hostDeleteOptions *HostDeleteOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.HostDeleteWithContext(context.Background(), hostDeleteOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostDeleteWithContext is an alternate form of the HostDelete method which supports a Context parameter
func (sdsaas *SdsaasV1) HostDeleteWithContext(ctx context.Context, hostDeleteOptions *HostDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostDeleteOptions, "hostDeleteOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostDeleteOptions, "hostDeleteOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostDeleteOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_delete", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// HostMappings : List all volume mappings for a host
// This request lists volume mappings for a host.
func (sdsaas *SdsaasV1) HostMappings(hostMappingsOptions *HostMappingsOptions) (result *VolumeMappingCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostMappingsWithContext(context.Background(), hostMappingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostMappingsWithContext is an alternate form of the HostMappings method which supports a Context parameter
func (sdsaas *SdsaasV1) HostMappingsWithContext(ctx context.Context, hostMappingsOptions *HostMappingsOptions) (result *VolumeMappingCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostMappingsOptions, "hostMappingsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostMappingsOptions, "hostMappingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostMappingsOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostMappingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostMappings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_mappings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeMappingCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostMappingCreate : Create a Volume mapping for a host
// This request creates a new volume mapping for a given host.
func (sdsaas *SdsaasV1) HostMappingCreate(hostMappingCreateOptions *HostMappingCreateOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostMappingCreateWithContext(context.Background(), hostMappingCreateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostMappingCreateWithContext is an alternate form of the HostMappingCreate method which supports a Context parameter
func (sdsaas *SdsaasV1) HostMappingCreateWithContext(ctx context.Context, hostMappingCreateOptions *HostMappingCreateOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostMappingCreateOptions, "hostMappingCreateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostMappingCreateOptions, "hostMappingCreateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostMappingCreateOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostMappingCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostMappingCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if hostMappingCreateOptions.Volume != nil {
		body["volume"] = hostMappingCreateOptions.Volume
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_mapping_create", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeMapping)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostMappingDeleteAll : Deletes all the volume mappings for a given host
// This request deletes all volume mappings associated with a specific host ID.
func (sdsaas *SdsaasV1) HostMappingDeleteAll(hostMappingDeleteAllOptions *HostMappingDeleteAllOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.HostMappingDeleteAllWithContext(context.Background(), hostMappingDeleteAllOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostMappingDeleteAllWithContext is an alternate form of the HostMappingDeleteAll method which supports a Context parameter
func (sdsaas *SdsaasV1) HostMappingDeleteAllWithContext(ctx context.Context, hostMappingDeleteAllOptions *HostMappingDeleteAllOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostMappingDeleteAllOptions, "hostMappingDeleteAllOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostMappingDeleteAllOptions, "hostMappingDeleteAllOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id": *hostMappingDeleteAllOptions.HostID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostMappingDeleteAllOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostMappingDeleteAll")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_mapping_delete_all", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// HostMapping : Retrieve a volume mapping
// This request retrieves a single volume mapping specified by the identifier in the URL.
func (sdsaas *SdsaasV1) HostMapping(hostMappingOptions *HostMappingOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.HostMappingWithContext(context.Background(), hostMappingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostMappingWithContext is an alternate form of the HostMapping method which supports a Context parameter
func (sdsaas *SdsaasV1) HostMappingWithContext(ctx context.Context, hostMappingOptions *HostMappingOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostMappingOptions, "hostMappingOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostMappingOptions, "hostMappingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id":           *hostMappingOptions.HostID,
		"volume_mapping_id": *hostMappingOptions.VolumeMappingID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}/volume_mappings/{volume_mapping_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostMappingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostMapping")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_mapping", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeMapping)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// HostMappingDelete : Deletes the given volume mapping for a specific host
// This request deletes a particular volume mapped from the host.
func (sdsaas *SdsaasV1) HostMappingDelete(hostMappingDeleteOptions *HostMappingDeleteOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.HostMappingDeleteWithContext(context.Background(), hostMappingDeleteOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// HostMappingDeleteWithContext is an alternate form of the HostMappingDelete method which supports a Context parameter
func (sdsaas *SdsaasV1) HostMappingDeleteWithContext(ctx context.Context, hostMappingDeleteOptions *HostMappingDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(hostMappingDeleteOptions, "hostMappingDeleteOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(hostMappingDeleteOptions, "hostMappingDeleteOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"host_id":           *hostMappingDeleteOptions.HostID,
		"volume_mapping_id": *hostMappingDeleteOptions.VolumeMappingID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{host_id}/volume_mappings/{volume_mapping_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range hostMappingDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V1", "HostMappingDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "host_mapping_delete", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// CertCreateOptions : The CertCreate options.
type CertCreateOptions struct {
	// The certificate type that is to be used in the request. Acceptable values include - s3.
	Cert *string `json:"cert" validate:"required,ne="`

	// The request body containing the S3 TLS certificate. The CLI will accept certificate files of any type, but they must
	// be in proper .pem format.
	Body io.ReadCloser `json:"body" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCertCreateOptions : Instantiate CertCreateOptions
func (*SdsaasV1) NewCertCreateOptions(cert string, body io.ReadCloser) *CertCreateOptions {
	return &CertCreateOptions{
		Cert: core.StringPtr(cert),
		Body: body,
	}
}

// SetCert : Allow user to set Cert
func (_options *CertCreateOptions) SetCert(cert string) *CertCreateOptions {
	_options.Cert = core.StringPtr(cert)
	return _options
}

// SetBody : Allow user to set Body
func (_options *CertCreateOptions) SetBody(body io.ReadCloser) *CertCreateOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CertCreateOptions) SetHeaders(param map[string]string) *CertCreateOptions {
	options.Headers = param
	return options
}

// CertDeleteOptions : The CertDelete options.
type CertDeleteOptions struct {
	// The certificate type that is to be used in the request. Acceptable values include - s3.
	Cert *string `json:"cert" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCertDeleteOptions : Instantiate CertDeleteOptions
func (*SdsaasV1) NewCertDeleteOptions(cert string) *CertDeleteOptions {
	return &CertDeleteOptions{
		Cert: core.StringPtr(cert),
	}
}

// SetCert : Allow user to set Cert
func (_options *CertDeleteOptions) SetCert(cert string) *CertDeleteOptions {
	_options.Cert = core.StringPtr(cert)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CertDeleteOptions) SetHeaders(param map[string]string) *CertDeleteOptions {
	options.Headers = param
	return options
}

// CertOptions : The Cert options.
type CertOptions struct {
	// The certificate type that is to be used in the request. Acceptable values include - s3.
	Cert *string `json:"cert" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCertOptions : Instantiate CertOptions
func (*SdsaasV1) NewCertOptions(cert string) *CertOptions {
	return &CertOptions{
		Cert: core.StringPtr(cert),
	}
}

// SetCert : Allow user to set Cert
func (_options *CertOptions) SetCert(cert string) *CertOptions {
	_options.Cert = core.StringPtr(cert)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CertOptions) SetHeaders(param map[string]string) *CertOptions {
	options.Headers = param
	return options
}

// CertTypesOptions : The CertTypes options.
type CertTypesOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCertTypesOptions : Instantiate CertTypesOptions
func (*SdsaasV1) NewCertTypesOptions() *CertTypesOptions {
	return &CertTypesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CertTypesOptions) SetHeaders(param map[string]string) *CertTypesOptions {
	options.Headers = param
	return options
}

// CertUpdateOptions : The CertUpdate options.
type CertUpdateOptions struct {
	// The certificate type that is to be used in the request. Acceptable values include - s3.
	Cert *string `json:"cert" validate:"required,ne="`

	// The request body containing the S3 TLS certificate. The CLI will accept certificate files of any type, but they must
	// be in proper .pem format.
	Body io.ReadCloser `json:"body" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCertUpdateOptions : Instantiate CertUpdateOptions
func (*SdsaasV1) NewCertUpdateOptions(cert string, body io.ReadCloser) *CertUpdateOptions {
	return &CertUpdateOptions{
		Cert: core.StringPtr(cert),
		Body: body,
	}
}

// SetCert : Allow user to set Cert
func (_options *CertUpdateOptions) SetCert(cert string) *CertUpdateOptions {
	_options.Cert = core.StringPtr(cert)
	return _options
}

// SetBody : Allow user to set Body
func (_options *CertUpdateOptions) SetBody(body io.ReadCloser) *CertUpdateOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CertUpdateOptions) SetHeaders(param map[string]string) *CertUpdateOptions {
	options.Headers = param
	return options
}

// CertificateFound : The responese object for certificate GET operations.
type CertificateFound struct {
	// The name of the certificate.
	Name *string `json:"name,omitempty"`

	// The expiration date of the certificate.
	ExpirationDate *strfmt.DateTime `json:"expiration_date,omitempty"`

	// The boolean value of the expiration status.
	Expired *bool `json:"expired,omitempty"`
}

// UnmarshalCertificateFound unmarshals an instance of CertificateFound from the specified map of raw messages.
func UnmarshalCertificateFound(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CertificateFound)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration_date", &obj.ExpirationDate)
	if err != nil {
		err = core.SDKErrorf(err, "", "expiration_date-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "expired", &obj.Expired)
	if err != nil {
		err = core.SDKErrorf(err, "", "expired-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CertificateList : The list of configured certificates.
type CertificateList struct {
	// The current list of configured certificates.
	Certificates []string `json:"certificates" validate:"required"`
}

// UnmarshalCertificateList unmarshals an instance of CertificateList from the specified map of raw messages.
func UnmarshalCertificateList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CertificateList)
	err = core.UnmarshalPrimitive(m, "certificates", &obj.Certificates)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificates-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CertificateUpdated : The response object for certificate POST operations.
type CertificateUpdated struct {
	// Name of the certificate.
	Name *string `json:"name,omitempty"`

	// A trace string for the request that caused the error, should be a correlation ID that can be used to track down the
	// underlying issue.
	Trace *string `json:"trace,omitempty"`

	// An array of certificate error codes and their descriptions.
	Errors []map[string]string `json:"errors,omitempty"`

	// The boolean valid status of the certificate.
	ValidCertificate *bool `json:"valid_certificate,omitempty"`

	// The boolean valid status of the key.
	ValidKey *bool `json:"valid_key,omitempty"`
}

// UnmarshalCertificateUpdated unmarshals an instance of CertificateUpdated from the specified map of raw messages.
func UnmarshalCertificateUpdated(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CertificateUpdated)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		err = core.SDKErrorf(err, "", "trace-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "valid_certificate", &obj.ValidCertificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "valid_certificate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "valid_key", &obj.ValidKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "valid_key-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CredCreateOptions : The CredCreate options.
type CredCreateOptions struct {
	// Access key to update or set.
	AccessKey *string `json:"access_key" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCredCreateOptions : Instantiate CredCreateOptions
func (*SdsaasV1) NewCredCreateOptions(accessKey string) *CredCreateOptions {
	return &CredCreateOptions{
		AccessKey: core.StringPtr(accessKey),
	}
}

// SetAccessKey : Allow user to set AccessKey
func (_options *CredCreateOptions) SetAccessKey(accessKey string) *CredCreateOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CredCreateOptions) SetHeaders(param map[string]string) *CredCreateOptions {
	options.Headers = param
	return options
}

// CredDeleteOptions : The CredDelete options.
type CredDeleteOptions struct {
	// Access key to update or set.
	AccessKey *string `json:"access_key" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCredDeleteOptions : Instantiate CredDeleteOptions
func (*SdsaasV1) NewCredDeleteOptions(accessKey string) *CredDeleteOptions {
	return &CredDeleteOptions{
		AccessKey: core.StringPtr(accessKey),
	}
}

// SetAccessKey : Allow user to set AccessKey
func (_options *CredDeleteOptions) SetAccessKey(accessKey string) *CredDeleteOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CredDeleteOptions) SetHeaders(param map[string]string) *CredDeleteOptions {
	options.Headers = param
	return options
}

// CredentialsFound : The response object for credential GET operations.
type CredentialsFound struct {
	// Collection of access keys.
	S3Credentials []string `json:"s3_credentials,omitempty"`
}

// UnmarshalCredentialsFound unmarshals an instance of CredentialsFound from the specified map of raw messages.
func UnmarshalCredentialsFound(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CredentialsFound)
	err = core.UnmarshalPrimitive(m, "s3_credentials", &obj.S3Credentials)
	if err != nil {
		err = core.SDKErrorf(err, "", "s3_credentials-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CredentialsUpdated : The response object for credential POST operations.
type CredentialsUpdated struct {
	// The user created access key.
	AccessKey *string `json:"access_key,omitempty"`

	// The key material associated with and access key.
	SecretKey *string `json:"secret_key,omitempty"`
}

// UnmarshalCredentialsUpdated unmarshals an instance of CredentialsUpdated from the specified map of raw messages.
func UnmarshalCredentialsUpdated(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CredentialsUpdated)
	err = core.UnmarshalPrimitive(m, "access_key", &obj.AccessKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_key", &obj.SecretKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "secret_key-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CredsOptions : The Creds options.
type CredsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCredsOptions : Instantiate CredsOptions
func (*SdsaasV1) NewCredsOptions() *CredsOptions {
	return &CredsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CredsOptions) SetHeaders(param map[string]string) *CredsOptions {
	options.Headers = param
	return options
}

// Gateway : Connection properties for the NVME gateways.
type Gateway struct {
	// Network information for volume/host mappings.
	IPAddress *string `json:"ip_address" validate:"required"`

	// Network information for volume/host mappings.
	Port *int64 `json:"port" validate:"required"`
}

// UnmarshalGateway unmarshals an instance of Gateway from the specified map of raw messages.
func UnmarshalGateway(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Gateway)
	err = core.UnmarshalPrimitive(m, "ip_address", &obj.IPAddress)
	if err != nil {
		err = core.SDKErrorf(err, "", "ip_address-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Host : The host object.
type Host struct {
	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// Unique identifer of the host.
	ID *string `json:"id" validate:"required"`

	// Unique name of the host.
	Name *string `json:"name" validate:"required"`

	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`

	// The host-to-volume map.
	VolumeMappings []VolumeMapping `json:"volume_mappings" validate:"required"`
}

// UnmarshalHost unmarshals an instance of Host from the specified map of raw messages.
func UnmarshalHost(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Host)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "nqn", &obj.Nqn)
	if err != nil {
		err = core.SDKErrorf(err, "", "nqn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume_mappings", &obj.VolumeMappings, UnmarshalVolumeMapping)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_mappings-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HostCollection : A collection of hosts at a particular endpoint with the total number found.  Any hosts beyond the return limit are
// found in the Next link.
type HostCollection struct {
	// A link to the first page of resources.
	First *PageLink `json:"first" validate:"required"`

	// Collection of hosts.
	Hosts []Host `json:"hosts" validate:"required"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit" validate:"required"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalHostCollection unmarshals an instance of HostCollection from the specified map of raw messages.
func UnmarshalHostCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostCollection)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "hosts-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HostCreateOptions : The HostCreate options.
type HostCreateOptions struct {
	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`

	// The name for this host. The name must not be used by another host. If unspecified, the name will be a hyphenated
	// list of randomly-selected words.
	Name *string `json:"name,omitempty"`

	// List of volume IDs to be mapped to the host.
	VolumeMappings []VolumeMappingPrototype `json:"volume_mappings,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostCreateOptions : Instantiate HostCreateOptions
func (*SdsaasV1) NewHostCreateOptions(nqn string) *HostCreateOptions {
	return &HostCreateOptions{
		Nqn: core.StringPtr(nqn),
	}
}

// SetNqn : Allow user to set Nqn
func (_options *HostCreateOptions) SetNqn(nqn string) *HostCreateOptions {
	_options.Nqn = core.StringPtr(nqn)
	return _options
}

// SetName : Allow user to set Name
func (_options *HostCreateOptions) SetName(name string) *HostCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetVolumeMappings : Allow user to set VolumeMappings
func (_options *HostCreateOptions) SetVolumeMappings(volumeMappings []VolumeMappingPrototype) *HostCreateOptions {
	_options.VolumeMappings = volumeMappings
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostCreateOptions) SetHeaders(param map[string]string) *HostCreateOptions {
	options.Headers = param
	return options
}

// HostDeleteOptions : The HostDelete options.
type HostDeleteOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostDeleteOptions : Instantiate HostDeleteOptions
func (*SdsaasV1) NewHostDeleteOptions(hostID string) *HostDeleteOptions {
	return &HostDeleteOptions{
		HostID: core.StringPtr(hostID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostDeleteOptions) SetHostID(hostID string) *HostDeleteOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostDeleteOptions) SetHeaders(param map[string]string) *HostDeleteOptions {
	options.Headers = param
	return options
}

// HostMappingCreateOptions : The HostMappingCreate options.
type HostMappingCreateOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// Volume identifier.
	Volume *VolumeIdentity `json:"volume" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostMappingCreateOptions : Instantiate HostMappingCreateOptions
func (*SdsaasV1) NewHostMappingCreateOptions(hostID string, volume *VolumeIdentity) *HostMappingCreateOptions {
	return &HostMappingCreateOptions{
		HostID: core.StringPtr(hostID),
		Volume: volume,
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostMappingCreateOptions) SetHostID(hostID string) *HostMappingCreateOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetVolume : Allow user to set Volume
func (_options *HostMappingCreateOptions) SetVolume(volume *VolumeIdentity) *HostMappingCreateOptions {
	_options.Volume = volume
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostMappingCreateOptions) SetHeaders(param map[string]string) *HostMappingCreateOptions {
	options.Headers = param
	return options
}

// HostMappingDeleteAllOptions : The HostMappingDeleteAll options.
type HostMappingDeleteAllOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostMappingDeleteAllOptions : Instantiate HostMappingDeleteAllOptions
func (*SdsaasV1) NewHostMappingDeleteAllOptions(hostID string) *HostMappingDeleteAllOptions {
	return &HostMappingDeleteAllOptions{
		HostID: core.StringPtr(hostID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostMappingDeleteAllOptions) SetHostID(hostID string) *HostMappingDeleteAllOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostMappingDeleteAllOptions) SetHeaders(param map[string]string) *HostMappingDeleteAllOptions {
	options.Headers = param
	return options
}

// HostMappingDeleteOptions : The HostMappingDelete options.
type HostMappingDeleteOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// A unique volume mapping ID.
	VolumeMappingID *string `json:"volume_mapping_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostMappingDeleteOptions : Instantiate HostMappingDeleteOptions
func (*SdsaasV1) NewHostMappingDeleteOptions(hostID string, volumeMappingID string) *HostMappingDeleteOptions {
	return &HostMappingDeleteOptions{
		HostID:          core.StringPtr(hostID),
		VolumeMappingID: core.StringPtr(volumeMappingID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostMappingDeleteOptions) SetHostID(hostID string) *HostMappingDeleteOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetVolumeMappingID : Allow user to set VolumeMappingID
func (_options *HostMappingDeleteOptions) SetVolumeMappingID(volumeMappingID string) *HostMappingDeleteOptions {
	_options.VolumeMappingID = core.StringPtr(volumeMappingID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostMappingDeleteOptions) SetHeaders(param map[string]string) *HostMappingDeleteOptions {
	options.Headers = param
	return options
}

// HostMappingOptions : The HostMapping options.
type HostMappingOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// A unique volume mapping ID.
	VolumeMappingID *string `json:"volume_mapping_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostMappingOptions : Instantiate HostMappingOptions
func (*SdsaasV1) NewHostMappingOptions(hostID string, volumeMappingID string) *HostMappingOptions {
	return &HostMappingOptions{
		HostID:          core.StringPtr(hostID),
		VolumeMappingID: core.StringPtr(volumeMappingID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostMappingOptions) SetHostID(hostID string) *HostMappingOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetVolumeMappingID : Allow user to set VolumeMappingID
func (_options *HostMappingOptions) SetVolumeMappingID(volumeMappingID string) *HostMappingOptions {
	_options.VolumeMappingID = core.StringPtr(volumeMappingID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostMappingOptions) SetHeaders(param map[string]string) *HostMappingOptions {
	options.Headers = param
	return options
}

// HostMappingsOptions : The HostMappings options.
type HostMappingsOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostMappingsOptions : Instantiate HostMappingsOptions
func (*SdsaasV1) NewHostMappingsOptions(hostID string) *HostMappingsOptions {
	return &HostMappingsOptions{
		HostID: core.StringPtr(hostID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostMappingsOptions) SetHostID(hostID string) *HostMappingsOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostMappingsOptions) SetHeaders(param map[string]string) *HostMappingsOptions {
	options.Headers = param
	return options
}

// HostOptions : The Host options.
type HostOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostOptions : Instantiate HostOptions
func (*SdsaasV1) NewHostOptions(hostID string) *HostOptions {
	return &HostOptions{
		HostID: core.StringPtr(hostID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostOptions) SetHostID(hostID string) *HostOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostOptions) SetHeaders(param map[string]string) *HostOptions {
	options.Headers = param
	return options
}

// HostPatch : The host PATCH request body.
type HostPatch struct {
	// Unique name of the host.
	Name *string `json:"name,omitempty"`
}

// UnmarshalHostPatch unmarshals an instance of HostPatch from the specified map of raw messages.
func UnmarshalHostPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the HostPatch
func (hostPatch *HostPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(hostPatch.Name) {
		_patch["name"] = hostPatch.Name
	}

	return
}

// HostReference : Host mapping schema.
type HostReference struct {
	// Unique identifer of the host.
	ID *string `json:"id" validate:"required"`

	// Unique name of the host.
	Name *string `json:"name" validate:"required"`

	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`
}

// UnmarshalHostReference unmarshals an instance of HostReference from the specified map of raw messages.
func UnmarshalHostReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "nqn", &obj.Nqn)
	if err != nil {
		err = core.SDKErrorf(err, "", "nqn-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HostUpdateOptions : The HostUpdate options.
type HostUpdateOptions struct {
	// A unique host ID.
	HostID *string `json:"host_id" validate:"required,ne="`

	// JSON Merge-Patch content for host_update.
	HostPatch map[string]interface{} `json:"Host_patch,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostUpdateOptions : Instantiate HostUpdateOptions
func (*SdsaasV1) NewHostUpdateOptions(hostID string) *HostUpdateOptions {
	return &HostUpdateOptions{
		HostID: core.StringPtr(hostID),
	}
}

// SetHostID : Allow user to set HostID
func (_options *HostUpdateOptions) SetHostID(hostID string) *HostUpdateOptions {
	_options.HostID = core.StringPtr(hostID)
	return _options
}

// SetHostPatch : Allow user to set HostPatch
func (_options *HostUpdateOptions) SetHostPatch(hostPatch map[string]interface{}) *HostUpdateOptions {
	_options.HostPatch = hostPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostUpdateOptions) SetHeaders(param map[string]string) *HostUpdateOptions {
	options.Headers = param
	return options
}

// HostsOptions : The Hosts options.
type HostsOptions struct {
	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection of resources by name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewHostsOptions : Instantiate HostsOptions
func (*SdsaasV1) NewHostsOptions() *HostsOptions {
	return &HostsOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *HostsOptions) SetLimit(limit int64) *HostsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *HostsOptions) SetName(name string) *HostsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *HostsOptions) SetHeaders(param map[string]string) *HostsOptions {
	options.Headers = param
	return options
}

// Namespace : The NVMe namespace properties for a given volume mapping.
type Namespace struct {
	// NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
	ID *int64 `json:"id,omitempty"`

	// UUID of the NVMe namespace.
	UUID *string `json:"uuid,omitempty"`
}

// UnmarshalNamespace unmarshals an instance of Namespace from the specified map of raw messages.
func UnmarshalNamespace(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Namespace)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "uuid", &obj.UUID)
	if err != nil {
		err = core.SDKErrorf(err, "", "uuid-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PageLink : A link to the help page.
type PageLink struct {
	// The URL for a page of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPageLink unmarshals an instance of PageLink from the specified map of raw messages.
func UnmarshalPageLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PageLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageIdentifier : Storage network and ID information associated with a volume/host mapping.
type StorageIdentifier struct {
	// The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
	SubsystemNqn *string `json:"subsystem_nqn" validate:"required"`

	// NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
	NamespaceID *int64 `json:"namespace_id" validate:"required"`

	// The namespace UUID associated with a volume/host mapping.
	NamespaceUUID *string `json:"namespace_uuid" validate:"required"`

	// List of NVMe gateways.
	Gateways []Gateway `json:"gateways" validate:"required"`
}

// UnmarshalStorageIdentifier unmarshals an instance of StorageIdentifier from the specified map of raw messages.
func UnmarshalStorageIdentifier(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageIdentifier)
	err = core.UnmarshalPrimitive(m, "subsystem_nqn", &obj.SubsystemNqn)
	if err != nil {
		err = core.SDKErrorf(err, "", "subsystem_nqn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "namespace_id", &obj.NamespaceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "namespace_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "namespace_uuid", &obj.NamespaceUUID)
	if err != nil {
		err = core.SDKErrorf(err, "", "namespace_uuid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "gateways", &obj.Gateways, UnmarshalGateway)
	if err != nil {
		err = core.SDKErrorf(err, "", "gateways-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Volume : The volume metadata prototype.
type Volume struct {
	// The maximum bandwidth (in megabits per second) for the volume.
	Bandwidth *int64 `json:"bandwidth" validate:"required"`

	// The capacity of the volume (in gigabytes).
	Capacity *int64 `json:"capacity" validate:"required"`

	// The date and time that the volume was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The volume profile id.
	ID *string `json:"id" validate:"required"`

	// Iops The maximum I/O operations per second (IOPS) for this volume.
	Iops *int64 `json:"iops" validate:"required"`

	// Unique name of the host.
	Name *string `json:"name" validate:"required"`

	// The resource type of the volume.
	ResourceType *string `json:"resource_type" validate:"required"`

	// The status of the volume resource. The enumerated values for this property will expand in the future. When
	// processing this property, check for and log unknown values. Optionally halt processing and surface the error, or
	// bypass the resource on which the unexpected property value was encountered.
	Status *string `json:"status,omitempty"`

	// The reasons for the current status (if any).
	StatusReasons []VolumeStatusReason `json:"status_reasons,omitempty"`

	// List of volume mappings for this volume.
	VolumeMappings []VolumeMapping `json:"volume_mappings" validate:"required"`
}

// Constants associated with the Volume.Status property.
// The status of the volume resource. The enumerated values for this property will expand in the future. When processing
// this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the
// resource on which the unexpected property value was encountered.
const (
	VolumeStatusAvailableConst       = "available"
	VolumeStatusPendingConst         = "pending"
	VolumeStatusPendingDeletionConst = "pending_deletion"
	VolumeStatusUpdatingConst        = "updating"
)

// UnmarshalVolume unmarshals an instance of Volume from the specified map of raw messages.
func UnmarshalVolume(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Volume)
	err = core.UnmarshalPrimitive(m, "bandwidth", &obj.Bandwidth)
	if err != nil {
		err = core.SDKErrorf(err, "", "bandwidth-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "capacity", &obj.Capacity)
	if err != nil {
		err = core.SDKErrorf(err, "", "capacity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iops", &obj.Iops)
	if err != nil {
		err = core.SDKErrorf(err, "", "iops-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "status_reasons", &obj.StatusReasons, UnmarshalVolumeStatusReason)
	if err != nil {
		err = core.SDKErrorf(err, "", "status_reasons-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume_mappings", &obj.VolumeMappings, UnmarshalVolumeMapping)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_mappings-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeCollection : Volume object showing the results of the GET volumes operation.
type VolumeCollection struct {
	// The first page of volume objects.
	First *PageLink `json:"first,omitempty"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit,omitempty"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// List of volumes retrieved.
	Volumes []Volume `json:"volumes,omitempty"`
}

// UnmarshalVolumeCollection unmarshals an instance of VolumeCollection from the specified map of raw messages.
func UnmarshalVolumeCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeCollection)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volumes", &obj.Volumes, UnmarshalVolume)
	if err != nil {
		err = core.SDKErrorf(err, "", "volumes-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeCreateOptions : The VolumeCreate options.
type VolumeCreateOptions struct {
	// The capacity to use for the volume (in gigabytes). The specified value must be within the capacity range of the
	// volume's profile.
	Capacity *int64 `json:"capacity" validate:"required"`

	// The name for this volume. The name must not be used by another volume. If unspecified, the name will be a hyphenated
	// list of randomly-selected words.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewVolumeCreateOptions : Instantiate VolumeCreateOptions
func (*SdsaasV1) NewVolumeCreateOptions(capacity int64) *VolumeCreateOptions {
	return &VolumeCreateOptions{
		Capacity: core.Int64Ptr(capacity),
	}
}

// SetCapacity : Allow user to set Capacity
func (_options *VolumeCreateOptions) SetCapacity(capacity int64) *VolumeCreateOptions {
	_options.Capacity = core.Int64Ptr(capacity)
	return _options
}

// SetName : Allow user to set Name
func (_options *VolumeCreateOptions) SetName(name string) *VolumeCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *VolumeCreateOptions) SetHeaders(param map[string]string) *VolumeCreateOptions {
	options.Headers = param
	return options
}

// VolumeDeleteOptions : The VolumeDelete options.
type VolumeDeleteOptions struct {
	// The volume profile id.
	VolumeID *string `json:"volume_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewVolumeDeleteOptions : Instantiate VolumeDeleteOptions
func (*SdsaasV1) NewVolumeDeleteOptions(volumeID string) *VolumeDeleteOptions {
	return &VolumeDeleteOptions{
		VolumeID: core.StringPtr(volumeID),
	}
}

// SetVolumeID : Allow user to set VolumeID
func (_options *VolumeDeleteOptions) SetVolumeID(volumeID string) *VolumeDeleteOptions {
	_options.VolumeID = core.StringPtr(volumeID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *VolumeDeleteOptions) SetHeaders(param map[string]string) *VolumeDeleteOptions {
	options.Headers = param
	return options
}

// VolumeIdentity : Volume identifier.
type VolumeIdentity struct {
	// Unique identifer of the host.
	ID *string `json:"id" validate:"required"`
}

// NewVolumeIdentity : Instantiate VolumeIdentity (Generic Model Constructor)
func (*SdsaasV1) NewVolumeIdentity(id string) (_model *VolumeIdentity, err error) {
	_model = &VolumeIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalVolumeIdentity unmarshals an instance of VolumeIdentity from the specified map of raw messages.
func UnmarshalVolumeIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeMapping : Schema for a volume mapping to a host.
type VolumeMapping struct {
	// The status of the volume mapping. The enumerated values for this property will expand in the future. When processing
	// this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the
	// resource on which the unexpected property value was encountered.
	Status *string `json:"status" validate:"required"`

	// Storage network and ID information associated with a volume/host mapping.
	StorageIdentifier *StorageIdentifier `json:"storage_identifier,omitempty"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// Unique identifier of the mapping.
	ID *string `json:"id" validate:"required"`

	// The volume reference.
	Volume *VolumeReference `json:"volume,omitempty"`

	// Host mapping schema.
	Host *HostReference `json:"host,omitempty"`

	// The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
	SubsystemNqn *string `json:"subsystem_nqn,omitempty"`

	// The NVMe namespace properties for a given volume mapping.
	Namespace *Namespace `json:"namespace,omitempty"`

	// List of NVMe gateways.
	Gateways []Gateway `json:"gateways,omitempty"`
}

// Constants associated with the VolumeMapping.Status property.
// The status of the volume mapping. The enumerated values for this property will expand in the future. When processing
// this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the
// resource on which the unexpected property value was encountered.
const (
	VolumeMappingStatusMappedConst           = "mapped"
	VolumeMappingStatusMappingFailedConst    = "mapping_failed"
	VolumeMappingStatusPendingConst          = "pending"
	VolumeMappingStatusPendingUnmappingConst = "pending_unmapping"
)

// UnmarshalVolumeMapping unmarshals an instance of VolumeMapping from the specified map of raw messages.
func UnmarshalVolumeMapping(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMapping)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "storage_identifier", &obj.StorageIdentifier, UnmarshalStorageIdentifier)
	if err != nil {
		err = core.SDKErrorf(err, "", "storage_identifier-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume", &obj.Volume, UnmarshalVolumeReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "host", &obj.Host, UnmarshalHostReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "subsystem_nqn", &obj.SubsystemNqn)
	if err != nil {
		err = core.SDKErrorf(err, "", "subsystem_nqn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "namespace", &obj.Namespace, UnmarshalNamespace)
	if err != nil {
		err = core.SDKErrorf(err, "", "namespace-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "gateways", &obj.Gateways, UnmarshalGateway)
	if err != nil {
		err = core.SDKErrorf(err, "", "gateways-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeMappingCollection : Collection of volume mappings for a host.
type VolumeMappingCollection struct {
	// A link to the first page of resources.
	First *PageLink `json:"first" validate:"required"`

	// List of volume_mappings.
	VolumeMappings []VolumeMapping `json:"volume_mappings" validate:"required"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit" validate:"required"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalVolumeMappingCollection unmarshals an instance of VolumeMappingCollection from the specified map of raw messages.
func UnmarshalVolumeMappingCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMappingCollection)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume_mappings", &obj.VolumeMappings, UnmarshalVolumeMapping)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_mappings-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeMappingPrototype : The volume mapping request.
type VolumeMappingPrototype struct {
	// Volume identifier.
	Volume *VolumeIdentity `json:"volume" validate:"required"`
}

// NewVolumeMappingPrototype : Instantiate VolumeMappingPrototype (Generic Model Constructor)
func (*SdsaasV1) NewVolumeMappingPrototype(volume *VolumeIdentity) (_model *VolumeMappingPrototype, err error) {
	_model = &VolumeMappingPrototype{
		Volume: volume,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalVolumeMappingPrototype unmarshals an instance of VolumeMappingPrototype from the specified map of raw messages.
func UnmarshalVolumeMappingPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMappingPrototype)
	err = core.UnmarshalModel(m, "volume", &obj.Volume, UnmarshalVolumeIdentity)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeOptions : The Volume options.
type VolumeOptions struct {
	// The volume profile id.
	VolumeID *string `json:"volume_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewVolumeOptions : Instantiate VolumeOptions
func (*SdsaasV1) NewVolumeOptions(volumeID string) *VolumeOptions {
	return &VolumeOptions{
		VolumeID: core.StringPtr(volumeID),
	}
}

// SetVolumeID : Allow user to set VolumeID
func (_options *VolumeOptions) SetVolumeID(volumeID string) *VolumeOptions {
	_options.VolumeID = core.StringPtr(volumeID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *VolumeOptions) SetHeaders(param map[string]string) *VolumeOptions {
	options.Headers = param
	return options
}

// VolumePatch : The capacity to use for the volume (in gigabytes). Additionally, the specified value must not be less than the
// current capacity.
type VolumePatch struct {
	// The capacity of the volume (in gigabytes).
	Capacity *int64 `json:"capacity,omitempty"`

	// The name for this volume. The name must not be used by another volume.
	Name *string `json:"name,omitempty"`
}

// UnmarshalVolumePatch unmarshals an instance of VolumePatch from the specified map of raw messages.
func UnmarshalVolumePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumePatch)
	err = core.UnmarshalPrimitive(m, "capacity", &obj.Capacity)
	if err != nil {
		err = core.SDKErrorf(err, "", "capacity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the VolumePatch
func (volumePatch *VolumePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(volumePatch.Capacity) {
		_patch["capacity"] = volumePatch.Capacity
	}
	if !core.IsNil(volumePatch.Name) {
		_patch["name"] = volumePatch.Name
	}

	return
}

// VolumeReference : The volume reference.
type VolumeReference struct {
	// Unique identifer of the host.
	ID *string `json:"id" validate:"required"`

	// Unique name of the host.
	Name *string `json:"name" validate:"required"`
}

// UnmarshalVolumeReference unmarshals an instance of VolumeReference from the specified map of raw messages.
func UnmarshalVolumeReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeStatusReason : The reason for the current status (if any).
type VolumeStatusReason struct {
	// A snake case string succinctly identifying the status reason.
	Code *string `json:"code" validate:"required"`

	// An explanation of the status reason.
	Message *string `json:"message" validate:"required"`

	// Link to documentation about this status reason.
	MoreInfo *string `json:"more_info,omitempty"`
}

// UnmarshalVolumeStatusReason unmarshals an instance of VolumeStatusReason from the specified map of raw messages.
func UnmarshalVolumeStatusReason(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeStatusReason)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		err = core.SDKErrorf(err, "", "code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		err = core.SDKErrorf(err, "", "more_info-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeUpdateOptions : The VolumeUpdate options.
type VolumeUpdateOptions struct {
	// The volume profile id.
	VolumeID *string `json:"volume_id" validate:"required,ne="`

	// A JSON object containing volume information.
	VolumePatch map[string]interface{} `json:"Volume_patch,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewVolumeUpdateOptions : Instantiate VolumeUpdateOptions
func (*SdsaasV1) NewVolumeUpdateOptions(volumeID string) *VolumeUpdateOptions {
	return &VolumeUpdateOptions{
		VolumeID: core.StringPtr(volumeID),
	}
}

// SetVolumeID : Allow user to set VolumeID
func (_options *VolumeUpdateOptions) SetVolumeID(volumeID string) *VolumeUpdateOptions {
	_options.VolumeID = core.StringPtr(volumeID)
	return _options
}

// SetVolumePatch : Allow user to set VolumePatch
func (_options *VolumeUpdateOptions) SetVolumePatch(volumePatch map[string]interface{}) *VolumeUpdateOptions {
	_options.VolumePatch = volumePatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *VolumeUpdateOptions) SetHeaders(param map[string]string) *VolumeUpdateOptions {
	options.Headers = param
	return options
}

// VolumesOptions : The Volumes options.
type VolumesOptions struct {
	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection of resources by name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewVolumesOptions : Instantiate VolumesOptions
func (*SdsaasV1) NewVolumesOptions() *VolumesOptions {
	return &VolumesOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *VolumesOptions) SetLimit(limit int64) *VolumesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *VolumesOptions) SetName(name string) *VolumesOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *VolumesOptions) SetHeaders(param map[string]string) *VolumesOptions {
	options.Headers = param
	return options
}
