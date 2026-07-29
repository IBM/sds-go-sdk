/**
 * (C) Copyright IBM Corp. 2024-2026.
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
 * IBM OpenAPI SDK Code Generator Version: 3.110.1-b7c5b1d1-20260205-165953
 */

// Package sdsaasv2 : Operations and models for the SdsaasV2 service
package sdsaasv2

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

// SdsaasV2 : OpenAPI definition for SDSaaS
//
// API Version: 2.0.0
type SdsaasV2 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "sdsaas"

const ParameterizedServiceURL = "{url}"

var defaultUrlVariables = map[string]string{
	"url": "{url}",
}

// SdsaasV2Options : Service options
type SdsaasV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSdsaasV2UsingExternalConfig : constructs an instance of SdsaasV2 with passed in options and external configuration.
func NewSdsaasV2UsingExternalConfig(options *SdsaasV2Options) (sdsaas *SdsaasV2, err error) {
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

	sdsaas, err = NewSdsaasV2(options)
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

// NewSdsaasV2 : constructs an instance of SdsaasV2 with passed in options.
func NewSdsaasV2(options *SdsaasV2Options) (service *SdsaasV2, err error) {
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

	service = &SdsaasV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "sdsaas" suitable for processing requests.
func (sdsaas *SdsaasV2) Clone() *SdsaasV2 {
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
func (sdsaas *SdsaasV2) SetServiceURL(url string) error {
	err := sdsaas.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (sdsaas *SdsaasV2) GetServiceURL() string {
	return sdsaas.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (sdsaas *SdsaasV2) SetDefaultHeaders(headers http.Header) {
	sdsaas.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (sdsaas *SdsaasV2) SetEnableGzipCompression(enableGzip bool) {
	sdsaas.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (sdsaas *SdsaasV2) GetEnableGzipCompression() bool {
	return sdsaas.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (sdsaas *SdsaasV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	sdsaas.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (sdsaas *SdsaasV2) DisableRetries() {
	sdsaas.Service.DisableRetries()
}

// ListVolumes : List all Volumes
// This request lists volumes in the deployment. Volumes are network-connected block storage devices that may be mapped
// to one or more hosts in the same deployment.
func (sdsaas *SdsaasV2) ListVolumes(listVolumesOptions *ListVolumesOptions) (result *VolumeCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListVolumesWithContext(context.Background(), listVolumesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListVolumesWithContext is an alternate form of the ListVolumes method which supports a Context parameter
func (sdsaas *SdsaasV2) ListVolumesWithContext(ctx context.Context, listVolumesOptions *ListVolumesOptions) (result *VolumeCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVolumesOptions, "listVolumesOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListVolumes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listVolumesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listVolumesOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listVolumesOptions.Start))
	}
	if listVolumesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listVolumesOptions.Limit))
	}
	if listVolumesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listVolumesOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_volumes", getServiceComponentInfo())
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

// CreateVolume : Create a Volume
// This request creates a new volume from a volume prototype object. The prototype object is structured in the same way
// as a retrieved volume, and contains the information necessary to create the new volume.<br> To create restored volume
// using snapshot provide source_snapshot id.<br> To create restored volume using VolumeGroupSnapshot provide
// source_volume_group_snapshot id and volume id.
func (sdsaas *SdsaasV2) CreateVolume(createVolumeOptions *CreateVolumeOptions) (result *VolumeSummary, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateVolumeWithContext(context.Background(), createVolumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateVolumeWithContext is an alternate form of the CreateVolume method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateVolumeWithContext(ctx context.Context, createVolumeOptions *CreateVolumeOptions) (result *VolumeSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVolumeOptions, "createVolumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createVolumeOptions, "createVolumeOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateVolume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createVolumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVolumeOptions.Capacity != nil {
		body["capacity"] = createVolumeOptions.Capacity
	}
	if createVolumeOptions.Name != nil {
		body["name"] = createVolumeOptions.Name
	}
	if createVolumeOptions.SourceSnapshot != nil {
		body["source_snapshot"] = createVolumeOptions.SourceSnapshot
	}
	if createVolumeOptions.SourceVolumeGroupSnapshot != nil {
		body["source_volume_group_snapshot"] = createVolumeOptions.SourceVolumeGroupSnapshot
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
		core.EnrichHTTPProblem(err, "create_volume", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteVolume : Delete a volume
// This request deletes a volume. This operation cannot be reversed. For this request to succeed, the volume must not be
// mapped to any hosts.
func (sdsaas *SdsaasV2) DeleteVolume(deleteVolumeOptions *DeleteVolumeOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteVolumeWithContext(context.Background(), deleteVolumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteVolumeWithContext is an alternate form of the DeleteVolume method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteVolumeWithContext(ctx context.Context, deleteVolumeOptions *DeleteVolumeOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVolumeOptions, "deleteVolumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteVolumeOptions, "deleteVolumeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteVolumeOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteVolume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteVolumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_volume", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetVolume : Retrieve a volume
// This request retrieves a single volume specified by the identifier in the URL.
func (sdsaas *SdsaasV2) GetVolume(getVolumeOptions *GetVolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.GetVolumeWithContext(context.Background(), getVolumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetVolumeWithContext is an alternate form of the GetVolume method which supports a Context parameter
func (sdsaas *SdsaasV2) GetVolumeWithContext(ctx context.Context, getVolumeOptions *GetVolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVolumeOptions, "getVolumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getVolumeOptions, "getVolumeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *getVolumeOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "GetVolume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getVolumeOptions.Headers {
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
		core.EnrichHTTPProblem(err, "get_volume", getServiceComponentInfo())
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

// UpdateVolume : Update a volume
// This request updates a volume with the information in a provided volume patch. The volume patch object is structured
// in the same way as a retrieved volume and contains only the information to be updated.
func (sdsaas *SdsaasV2) UpdateVolume(updateVolumeOptions *UpdateVolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.UpdateVolumeWithContext(context.Background(), updateVolumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateVolumeWithContext is an alternate form of the UpdateVolume method which supports a Context parameter
func (sdsaas *SdsaasV2) UpdateVolumeWithContext(ctx context.Context, updateVolumeOptions *UpdateVolumeOptions) (result *Volume, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateVolumeOptions, "updateVolumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateVolumeOptions, "updateVolumeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateVolumeOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/volumes/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "UpdateVolume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateVolumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateVolumeOptions.VolumePatch)
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
		core.EnrichHTTPProblem(err, "update_volume", getServiceComponentInfo())
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

// ListHosts : List all hosts
// This request lists all hosts in the deployment. Hosts are objects representing the NVMe initiators that may be mapped
// to one or more volumes in the same deployment.
func (sdsaas *SdsaasV2) ListHosts(listHostsOptions *ListHostsOptions) (result *HostCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListHostsWithContext(context.Background(), listHostsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListHostsWithContext is an alternate form of the ListHosts method which supports a Context parameter
func (sdsaas *SdsaasV2) ListHostsWithContext(ctx context.Context, listHostsOptions *ListHostsOptions) (result *HostCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listHostsOptions, "listHostsOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListHosts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listHostsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listHostsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listHostsOptions.Start))
	}
	if listHostsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listHostsOptions.Limit))
	}
	if listHostsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listHostsOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_hosts", getServiceComponentInfo())
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

// CreateHost : Create a host
// This request creates a new host from a host prototype object. The prototype object contains the information necessary
// to provision the new host which can also be mapped to one or more volumes.
func (sdsaas *SdsaasV2) CreateHost(createHostOptions *CreateHostOptions) (result *HostSummary, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateHostWithContext(context.Background(), createHostOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateHostWithContext is an alternate form of the CreateHost method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateHostWithContext(ctx context.Context, createHostOptions *CreateHostOptions) (result *HostSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createHostOptions, "createHostOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createHostOptions, "createHostOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateHost")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createHostOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createHostOptions.Nqn != nil {
		body["nqn"] = createHostOptions.Nqn
	}
	if createHostOptions.Name != nil {
		body["name"] = createHostOptions.Name
	}
	if createHostOptions.Psk != nil {
		body["psk"] = createHostOptions.Psk
	}
	if createHostOptions.VolumeMappings != nil {
		body["volume_mappings"] = createHostOptions.VolumeMappings
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
		core.EnrichHTTPProblem(err, "create_host", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHostSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteHost : Delete a Host
// This request deletes a host. For this request to succeed, the host must not be mapped to any volumes.
func (sdsaas *SdsaasV2) DeleteHost(deleteHostOptions *DeleteHostOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteHostWithContext(context.Background(), deleteHostOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteHostWithContext is an alternate form of the DeleteHost method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteHostWithContext(ctx context.Context, deleteHostOptions *DeleteHostOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteHostOptions, "deleteHostOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteHostOptions, "deleteHostOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteHostOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteHost")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteHostOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_host", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetHost : Retrieve a Host
// This request retrieves a host specified by the identifier in the URL.
func (sdsaas *SdsaasV2) GetHost(getHostOptions *GetHostOptions) (result *Host, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.GetHostWithContext(context.Background(), getHostOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetHostWithContext is an alternate form of the GetHost method which supports a Context parameter
func (sdsaas *SdsaasV2) GetHostWithContext(ctx context.Context, getHostOptions *GetHostOptions) (result *Host, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getHostOptions, "getHostOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getHostOptions, "getHostOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *getHostOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "GetHost")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getHostOptions.Headers {
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
		core.EnrichHTTPProblem(err, "get_host", getServiceComponentInfo())
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

// UpdateHost : Update a given Host
// This request updates a Host with the information in a provided host patch object. The host patch object is structured
// in the same way as a retrieved host and contains only the information to be updated.
func (sdsaas *SdsaasV2) UpdateHost(updateHostOptions *UpdateHostOptions) (result *Host, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.UpdateHostWithContext(context.Background(), updateHostOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateHostWithContext is an alternate form of the UpdateHost method which supports a Context parameter
func (sdsaas *SdsaasV2) UpdateHostWithContext(ctx context.Context, updateHostOptions *UpdateHostOptions) (result *Host, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateHostOptions, "updateHostOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateHostOptions, "updateHostOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateHostOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "UpdateHost")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateHostOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateHostOptions.HostPatch)
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
		core.EnrichHTTPProblem(err, "update_host", getServiceComponentInfo())
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

// DeleteVolumeMappings : Deletes all the volume mappings for a given host
// This request deletes all volume mappings for a host.
func (sdsaas *SdsaasV2) DeleteVolumeMappings(deleteVolumeMappingsOptions *DeleteVolumeMappingsOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteVolumeMappingsWithContext(context.Background(), deleteVolumeMappingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteVolumeMappingsWithContext is an alternate form of the DeleteVolumeMappings method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteVolumeMappingsWithContext(ctx context.Context, deleteVolumeMappingsOptions *DeleteVolumeMappingsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVolumeMappingsOptions, "deleteVolumeMappingsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteVolumeMappingsOptions, "deleteVolumeMappingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteVolumeMappingsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteVolumeMappings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteVolumeMappingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_volume_mappings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListVolumeMappings : List all volume mappings for a host
// This request lists volume mappings for a host.
func (sdsaas *SdsaasV2) ListVolumeMappings(listVolumeMappingsOptions *ListVolumeMappingsOptions) (result *VolumeMappingCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListVolumeMappingsWithContext(context.Background(), listVolumeMappingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListVolumeMappingsWithContext is an alternate form of the ListVolumeMappings method which supports a Context parameter
func (sdsaas *SdsaasV2) ListVolumeMappingsWithContext(ctx context.Context, listVolumeMappingsOptions *ListVolumeMappingsOptions) (result *VolumeMappingCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listVolumeMappingsOptions, "listVolumeMappingsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listVolumeMappingsOptions, "listVolumeMappingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *listVolumeMappingsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListVolumeMappings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listVolumeMappingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listVolumeMappingsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listVolumeMappingsOptions.Start))
	}
	if listVolumeMappingsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listVolumeMappingsOptions.Limit))
	}
	if listVolumeMappingsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listVolumeMappingsOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_volume_mappings", getServiceComponentInfo())
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

// CreateVolumeMapping : Create a Volume mapping for a host
// This request creates a new volume mapping for a given host.
func (sdsaas *SdsaasV2) CreateVolumeMapping(createVolumeMappingOptions *CreateVolumeMappingOptions) (result *VolumeMappingReference, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateVolumeMappingWithContext(context.Background(), createVolumeMappingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateVolumeMappingWithContext is an alternate form of the CreateVolumeMapping method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateVolumeMappingWithContext(ctx context.Context, createVolumeMappingOptions *CreateVolumeMappingOptions) (result *VolumeMappingReference, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVolumeMappingOptions, "createVolumeMappingOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createVolumeMappingOptions, "createVolumeMappingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *createVolumeMappingOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}/volume_mappings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateVolumeMapping")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createVolumeMappingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVolumeMappingOptions.Volume != nil {
		body["volume"] = createVolumeMappingOptions.Volume
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
		core.EnrichHTTPProblem(err, "create_volume_mapping", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVolumeMappingReference)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteVolumeMapping : Deletes the volume mapping
// This request deletes the volume mapping.
func (sdsaas *SdsaasV2) DeleteVolumeMapping(deleteVolumeMappingOptions *DeleteVolumeMappingOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteVolumeMappingWithContext(context.Background(), deleteVolumeMappingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteVolumeMappingWithContext is an alternate form of the DeleteVolumeMapping method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteVolumeMappingWithContext(ctx context.Context, deleteVolumeMappingOptions *DeleteVolumeMappingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVolumeMappingOptions, "deleteVolumeMappingOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteVolumeMappingOptions, "deleteVolumeMappingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteVolumeMappingOptions.ID,
		"volume_mapping_id": *deleteVolumeMappingOptions.VolumeMappingID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}/volume_mappings/{volume_mapping_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteVolumeMapping")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteVolumeMappingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_volume_mapping", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetVolumeMapping : Retrieve a volume mapping
// This request retrieves a single volume mapping specified by the identifier in the URL.
func (sdsaas *SdsaasV2) GetVolumeMapping(getVolumeMappingOptions *GetVolumeMappingOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.GetVolumeMappingWithContext(context.Background(), getVolumeMappingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetVolumeMappingWithContext is an alternate form of the GetVolumeMapping method which supports a Context parameter
func (sdsaas *SdsaasV2) GetVolumeMappingWithContext(ctx context.Context, getVolumeMappingOptions *GetVolumeMappingOptions) (result *VolumeMapping, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVolumeMappingOptions, "getVolumeMappingOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getVolumeMappingOptions, "getVolumeMappingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *getVolumeMappingOptions.ID,
		"volume_mapping_id": *getVolumeMappingOptions.VolumeMappingID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/hosts/{id}/volume_mappings/{volume_mapping_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "GetVolumeMapping")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getVolumeMappingOptions.Headers {
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
		core.EnrichHTTPProblem(err, "get_volume_mapping", getServiceComponentInfo())
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

// ListHmacCredentials : List HMAC credentials
// Gets list of HMAC credential access keys for S3 authentication.
func (sdsaas *SdsaasV2) ListHmacCredentials(listHmacCredentialsOptions *ListHmacCredentialsOptions) (result *StorageCredResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListHmacCredentialsWithContext(context.Background(), listHmacCredentialsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListHmacCredentialsWithContext is an alternate form of the ListHmacCredentials method which supports a Context parameter
func (sdsaas *SdsaasV2) ListHmacCredentialsWithContext(ctx context.Context, listHmacCredentialsOptions *ListHmacCredentialsOptions) (result *StorageCredResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listHmacCredentialsOptions, "listHmacCredentialsOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListHmacCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listHmacCredentialsOptions.Headers {
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
		core.EnrichHTTPProblem(err, "list_hmac_credentials", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStorageCredResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteHmacCredentials : Delete HMAC credentials
// Deletes the specific HMAC credential using the provided access key.
func (sdsaas *SdsaasV2) DeleteHmacCredentials(deleteHmacCredentialsOptions *DeleteHmacCredentialsOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteHmacCredentialsWithContext(context.Background(), deleteHmacCredentialsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteHmacCredentialsWithContext is an alternate form of the DeleteHmacCredentials method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteHmacCredentialsWithContext(ctx context.Context, deleteHmacCredentialsOptions *DeleteHmacCredentialsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteHmacCredentialsOptions, "deleteHmacCredentialsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteHmacCredentialsOptions, "deleteHmacCredentialsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"access_key": *deleteHmacCredentialsOptions.AccessKey,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/s3_credentials/{access_key}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteHmacCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteHmacCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_hmac_credentials", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateHmacCredentials : Create HMAC credentials
// Given an access key, creates an HMAC credential and returns the access key and secret key pair.
func (sdsaas *SdsaasV2) CreateHmacCredentials(createHmacCredentialsOptions *CreateHmacCredentialsOptions) (result *AccessKeyResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateHmacCredentialsWithContext(context.Background(), createHmacCredentialsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateHmacCredentialsWithContext is an alternate form of the CreateHmacCredentials method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateHmacCredentialsWithContext(ctx context.Context, createHmacCredentialsOptions *CreateHmacCredentialsOptions) (result *AccessKeyResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createHmacCredentialsOptions, "createHmacCredentialsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createHmacCredentialsOptions, "createHmacCredentialsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"access_key": *createHmacCredentialsOptions.AccessKey,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/s3_credentials/{access_key}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateHmacCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createHmacCredentialsOptions.Headers {
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
		core.EnrichHTTPProblem(err, "create_hmac_credentials", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccessKeyResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListCertificates : List the configured certificates
// Retrieves the list configured certificates.
func (sdsaas *SdsaasV2) ListCertificates(listCertificatesOptions *ListCertificatesOptions) (result *CertListResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListCertificatesWithContext(context.Background(), listCertificatesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListCertificatesWithContext is an alternate form of the ListCertificates method which supports a Context parameter
func (sdsaas *SdsaasV2) ListCertificatesWithContext(ctx context.Context, listCertificatesOptions *ListCertificatesOptions) (result *CertListResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCertificatesOptions, "listCertificatesOptions")
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

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListCertificates")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listCertificatesOptions.Headers {
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
		core.EnrichHTTPProblem(err, "list_certificates", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertListResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSslCert : Delete SSL certificate
// Delete the provided PEM formatted TLS certificate.
func (sdsaas *SdsaasV2) DeleteSslCert(deleteSslCertOptions *DeleteSslCertOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteSslCertWithContext(context.Background(), deleteSslCertOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSslCertWithContext is an alternate form of the DeleteSslCert method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteSslCertWithContext(ctx context.Context, deleteSslCertOptions *DeleteSslCertOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSslCertOptions, "deleteSslCertOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSslCertOptions, "deleteSslCertOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert_type": *deleteSslCertOptions.CertType,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert_type}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteSslCert")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteSslCertOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_ssl_cert", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetS3SslCertStatus : Get SSL cert status
// Retrieves the expiration date and status of the SSL certificate. The expiration status indicates the validity of the
// certificate.
func (sdsaas *SdsaasV2) GetS3SslCertStatus(getS3SslCertStatusOptions *GetS3SslCertStatusOptions) (result *StatusResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.GetS3SslCertStatusWithContext(context.Background(), getS3SslCertStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetS3SslCertStatusWithContext is an alternate form of the GetS3SslCertStatus method which supports a Context parameter
func (sdsaas *SdsaasV2) GetS3SslCertStatusWithContext(ctx context.Context, getS3SslCertStatusOptions *GetS3SslCertStatusOptions) (result *StatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getS3SslCertStatusOptions, "getS3SslCertStatusOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getS3SslCertStatusOptions, "getS3SslCertStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert_type": *getS3SslCertStatusOptions.CertType,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert_type}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "GetS3SslCertStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getS3SslCertStatusOptions.Headers {
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
		core.EnrichHTTPProblem(err, "get_s3_ssl_cert_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatusResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSslCert : Create SSL certificates
// Create the provided PEM formatted TLS certificate(s).
func (sdsaas *SdsaasV2) CreateSslCert(createSslCertOptions *CreateSslCertOptions) (result *CertResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateSslCertWithContext(context.Background(), createSslCertOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSslCertWithContext is an alternate form of the CreateSslCert method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateSslCertWithContext(ctx context.Context, createSslCertOptions *CreateSslCertOptions) (result *CertResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSslCertOptions, "createSslCertOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSslCertOptions, "createSslCertOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert_type": *createSslCertOptions.CertType,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert_type}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateSslCert")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createSslCertOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/octect-stream")

	_, err = builder.SetBodyContent("application/octect-stream", nil, nil, createSslCertOptions.Body)
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
		core.EnrichHTTPProblem(err, "create_ssl_cert", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceSslCert : Replace an SSL certificate
// Replace the provided PEM formatted TLS certificate.
func (sdsaas *SdsaasV2) ReplaceSslCert(replaceSslCertOptions *ReplaceSslCertOptions) (result *CertResponse, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ReplaceSslCertWithContext(context.Background(), replaceSslCertOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceSslCertWithContext is an alternate form of the ReplaceSslCert method which supports a Context parameter
func (sdsaas *SdsaasV2) ReplaceSslCertWithContext(ctx context.Context, replaceSslCertOptions *ReplaceSslCertOptions) (result *CertResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceSslCertOptions, "replaceSslCertOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceSslCertOptions, "replaceSslCertOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"cert_type": *replaceSslCertOptions.CertType,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/certificates/{cert_type}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ReplaceSslCert")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceSslCertOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/octect-stream")

	_, err = builder.SetBodyContent("application/octect-stream", nil, nil, replaceSslCertOptions.Body)
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
		core.EnrichHTTPProblem(err, "replace_ssl_cert", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCertResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSnapshots : Delete a filtered collection of snapshots
// This request deletes snapshots that match the specified filter.<br> To delete snapshots of single volume provide
// source_volume.id.<br> To delete snapshots of multi volume group provide source_volume_group.id.
func (sdsaas *SdsaasV2) DeleteSnapshots(deleteSnapshotsOptions *DeleteSnapshotsOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteSnapshotsWithContext(context.Background(), deleteSnapshotsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSnapshotsWithContext is an alternate form of the DeleteSnapshots method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteSnapshotsWithContext(ctx context.Context, deleteSnapshotsOptions *DeleteSnapshotsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSnapshotsOptions, "deleteSnapshotsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSnapshotsOptions, "deleteSnapshotsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteSnapshots")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteSnapshotsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("source_volume.id", fmt.Sprint(*deleteSnapshotsOptions.SourceVolumeID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_snapshots", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListSnapshots : List all Snapshots
// This request lists snapshots in the deployment.<br> A snapshot preserves the data of a volume at the time the
// snapshot is created.<br> In case of single volume snapshot, the snapshot response will include the source_volume
// details.<br> In case of multi volume group snapshot, the snapshot response will include source_volume_group
// details.<br>.
func (sdsaas *SdsaasV2) ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) (result *SnapshotCollection, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.ListSnapshotsWithContext(context.Background(), listSnapshotsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSnapshotsWithContext is an alternate form of the ListSnapshots method which supports a Context parameter
func (sdsaas *SdsaasV2) ListSnapshotsWithContext(ctx context.Context, listSnapshotsOptions *ListSnapshotsOptions) (result *SnapshotCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSnapshotsOptions, "listSnapshotsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "ListSnapshots")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listSnapshotsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listSnapshotsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listSnapshotsOptions.Start))
	}
	if listSnapshotsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listSnapshotsOptions.Limit))
	}
	if listSnapshotsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listSnapshotsOptions.Name))
	}
	if listSnapshotsOptions.SourceVolumeID != nil {
		builder.AddQuery("source_volume.id", fmt.Sprint(*listSnapshotsOptions.SourceVolumeID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sdsaas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_snapshots", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSnapshotCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSnapshot : Create a Snapshot
// This request creates a new snapshot from a snapshot prototype object. The prototype object is structured in the same
// way as a retrieved snapshot, and contains the information necessary to provision the new snapshot.<br> Note: After
// snapshot creation, the inital size is reported as 1 GiB by default.<br> To create single volume snapshot provide
// source_volume id.<br> To create multi volume group snapshot provide source_volume_group id.<br> Note: To create the
// snapshot, provide at least one of the source_volume or source_volume_group.<br> The parameters source_volume and
// source_volume_group cannot be specified simultaneously. Please provide only one of these parameters.
func (sdsaas *SdsaasV2) CreateSnapshot(createSnapshotOptions *CreateSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.CreateSnapshotWithContext(context.Background(), createSnapshotOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSnapshotWithContext is an alternate form of the CreateSnapshot method which supports a Context parameter
func (sdsaas *SdsaasV2) CreateSnapshotWithContext(ctx context.Context, createSnapshotOptions *CreateSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSnapshotOptions, "createSnapshotOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSnapshotOptions, "createSnapshotOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "CreateSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createSnapshotOptions.Name != nil {
		body["name"] = createSnapshotOptions.Name
	}
	if createSnapshotOptions.SourceVolume != nil {
		body["source_volume"] = createSnapshotOptions.SourceVolume
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
		core.EnrichHTTPProblem(err, "create_snapshot", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSnapshot)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSnapshot : Delete a snapshot
// This request deletes a snapshot.<br> To delete snapshots of single volume snapshot provide Snapshot id.<br> To delete
// snapshots of multi volume group snapshot provide VolumeGroupSnapshot id.
func (sdsaas *SdsaasV2) DeleteSnapshot(deleteSnapshotOptions *DeleteSnapshotOptions) (response *core.DetailedResponse, err error) {
	response, err = sdsaas.DeleteSnapshotWithContext(context.Background(), deleteSnapshotOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSnapshotWithContext is an alternate form of the DeleteSnapshot method which supports a Context parameter
func (sdsaas *SdsaasV2) DeleteSnapshotWithContext(ctx context.Context, deleteSnapshotOptions *DeleteSnapshotOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSnapshotOptions, "deleteSnapshotOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSnapshotOptions, "deleteSnapshotOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteSnapshotOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "DeleteSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = sdsaas.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_snapshot", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetSnapshot : Retrieve a snapshot
// This request retrieves a single snapshot specified by the identifier in the URL.<br> In case of single volume
// snapshot, the snapshot response will include the source_volume details.<br> In case of multi volume group snapshot,
// the snapshot response will include source_volume_group details.
func (sdsaas *SdsaasV2) GetSnapshot(getSnapshotOptions *GetSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.GetSnapshotWithContext(context.Background(), getSnapshotOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSnapshotWithContext is an alternate form of the GetSnapshot method which supports a Context parameter
func (sdsaas *SdsaasV2) GetSnapshotWithContext(ctx context.Context, getSnapshotOptions *GetSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSnapshotOptions, "getSnapshotOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSnapshotOptions, "getSnapshotOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *getSnapshotOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "GetSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getSnapshotOptions.Headers {
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
		core.EnrichHTTPProblem(err, "get_snapshot", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSnapshot)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateSnapshot : Update a snapshot
// This request updates a snapshot with the information in a provided snapshot patch.
func (sdsaas *SdsaasV2) UpdateSnapshot(updateSnapshotOptions *UpdateSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	result, response, err = sdsaas.UpdateSnapshotWithContext(context.Background(), updateSnapshotOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSnapshotWithContext is an alternate form of the UpdateSnapshot method which supports a Context parameter
func (sdsaas *SdsaasV2) UpdateSnapshotWithContext(ctx context.Context, updateSnapshotOptions *UpdateSnapshotOptions) (result *Snapshot, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSnapshotOptions, "updateSnapshotOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateSnapshotOptions, "updateSnapshotOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateSnapshotOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sdsaas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sdsaas.Service.Options.URL, `/snapshots/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("sdsaas", "V2", "UpdateSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateSnapshotOptions.SnapshotPatch)
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
		core.EnrichHTTPProblem(err, "update_snapshot", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSnapshot)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "2.0.0")
}

// AccessKeyResponse : The response for creating a new HMAC credential, containing the access key and secret key for that HMAC credential.
type AccessKeyResponse struct {
	// The provided access key for the newly created HMAC credential.
	AccessKey *string `json:"access_key,omitempty"`

	// The generated secret key for the newly created HMAC credential.
	SecretKey *string `json:"secret_key,omitempty"`
}

// UnmarshalAccessKeyResponse unmarshals an instance of AccessKeyResponse from the specified map of raw messages.
func UnmarshalAccessKeyResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccessKeyResponse)
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

// CertListResponse : The list of configured certificates.
type CertListResponse struct {
	// The current list of configured certificates.
	Certificates []string `json:"certificates" validate:"required"`
}

// UnmarshalCertListResponse unmarshals an instance of CertListResponse from the specified map of raw messages.
func UnmarshalCertListResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CertListResponse)
	err = core.UnmarshalPrimitive(m, "certificates", &obj.Certificates)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificates-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CertResponse : The response for posting a new SSL certificate, including any errors and if the provided key and certificate are
// valid.
type CertResponse struct {
	// Any errors that were encountered when attempting to upload the provided SSL certificate.
	Errors []ErrorObject `json:"errors" validate:"required"`

	// Name of the certificate.
	Name *string `json:"name,omitempty"`

	// A trace string for the request that caused the error, should be a correlation ID that can be used to track down the
	// underlying issue.
	Trace *string `json:"trace,omitempty"`

	// When set to true, indicates that the provided certificate is valid.
	ValidCertificate *bool `json:"valid_certificate,omitempty"`

	// When set to true, indicates that the provided key for the certificate is valid.
	ValidKey *bool `json:"valid_key,omitempty"`
}

// UnmarshalCertResponse unmarshals an instance of CertResponse from the specified map of raw messages.
func UnmarshalCertResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CertResponse)
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorObject)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
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

// CreateHmacCredentialsOptions : The CreateHmacCredentials options.
type CreateHmacCredentialsOptions struct {
	// The access key to create the new HMAC credential, must be less than 2048 characters long and may only contain
	// alphanumeric characters, underscores, and dashes.
	AccessKey *string `json:"access_key" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateHmacCredentialsOptions : Instantiate CreateHmacCredentialsOptions
func (*SdsaasV2) NewCreateHmacCredentialsOptions(accessKey string) *CreateHmacCredentialsOptions {
	return &CreateHmacCredentialsOptions{
		AccessKey: core.StringPtr(accessKey),
	}
}

// SetAccessKey : Allow user to set AccessKey
func (_options *CreateHmacCredentialsOptions) SetAccessKey(accessKey string) *CreateHmacCredentialsOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateHmacCredentialsOptions) SetHeaders(param map[string]string) *CreateHmacCredentialsOptions {
	options.Headers = param
	return options
}

// CreateHostOptions : The CreateHost options.
type CreateHostOptions struct {
	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`

	// The name for this host. The name must not be used by another host.
	//     If unspecified, the name will be a hyphenated list of randomly-selected words.
	Name *string `json:"name,omitempty"`

	// Transport Layer Security pre-shared key ciphersuites (TLS-PSK) is a set of cryptographic protocols that provide
	// secure communication based on pre-shared keys (PSKs).
	Psk *string `json:"psk,omitempty"`

	// List of volume IDs to be mapped to the host.
	VolumeMappings []VolumeMappingPrototype `json:"volume_mappings,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateHostOptions : Instantiate CreateHostOptions
func (*SdsaasV2) NewCreateHostOptions(nqn string) *CreateHostOptions {
	return &CreateHostOptions{
		Nqn: core.StringPtr(nqn),
	}
}

// SetNqn : Allow user to set Nqn
func (_options *CreateHostOptions) SetNqn(nqn string) *CreateHostOptions {
	_options.Nqn = core.StringPtr(nqn)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateHostOptions) SetName(name string) *CreateHostOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPsk : Allow user to set Psk
func (_options *CreateHostOptions) SetPsk(psk string) *CreateHostOptions {
	_options.Psk = core.StringPtr(psk)
	return _options
}

// SetVolumeMappings : Allow user to set VolumeMappings
func (_options *CreateHostOptions) SetVolumeMappings(volumeMappings []VolumeMappingPrototype) *CreateHostOptions {
	_options.VolumeMappings = volumeMappings
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateHostOptions) SetHeaders(param map[string]string) *CreateHostOptions {
	options.Headers = param
	return options
}

// CreateSnapshotOptions : The CreateSnapshot options.
type CreateSnapshotOptions struct {
	// The name for this snapshot. The name must not be used by another snapshot. If unspecified, the name will be a
	// hyphenated list of randomly-selected words.
	Name *string `json:"name,omitempty"`

	// The source volume this snapshot was created from (may be deleted).
	SourceVolume *SourceVolumePrototype `json:"source_volume,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSnapshotOptions : Instantiate CreateSnapshotOptions
func (*SdsaasV2) NewCreateSnapshotOptions() *CreateSnapshotOptions {
	return &CreateSnapshotOptions{}
}

// SetName : Allow user to set Name
func (_options *CreateSnapshotOptions) SetName(name string) *CreateSnapshotOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetSourceVolume : Allow user to set SourceVolume
func (_options *CreateSnapshotOptions) SetSourceVolume(sourceVolume *SourceVolumePrototype) *CreateSnapshotOptions {
	_options.SourceVolume = sourceVolume
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSnapshotOptions) SetHeaders(param map[string]string) *CreateSnapshotOptions {
	options.Headers = param
	return options
}

// CreateSslCertOptions : The CreateSslCert options.
type CreateSslCertOptions struct {
	// The certificate type that is to be used in the POST request. Acceptable values include: s3.
	CertType *string `json:"cert_type" validate:"required,ne="`

	// A string containing at least one PEM formatted TLS Certificate and a PEM formatted TLS Key for the POST request.
	// These should be formatted according to the RFC-7468 standard, with the CERTIFICATE property being used for
	// certificates and any of the following properties for keys: PRIVATE KEY, RSA PRIVATE KEY, EC PRIVATE KEY. If a
	// provided certificate or key is not in the specified formats, the API will reject it.
	Body io.ReadCloser `json:"body,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSslCertOptions : Instantiate CreateSslCertOptions
func (*SdsaasV2) NewCreateSslCertOptions(certType string) *CreateSslCertOptions {
	return &CreateSslCertOptions{
		CertType: core.StringPtr(certType),
	}
}

// SetCertType : Allow user to set CertType
func (_options *CreateSslCertOptions) SetCertType(certType string) *CreateSslCertOptions {
	_options.CertType = core.StringPtr(certType)
	return _options
}

// SetBody : Allow user to set Body
func (_options *CreateSslCertOptions) SetBody(body io.ReadCloser) *CreateSslCertOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSslCertOptions) SetHeaders(param map[string]string) *CreateSslCertOptions {
	options.Headers = param
	return options
}

// CreateVolumeMappingOptions : The CreateVolumeMapping options.
type CreateVolumeMappingOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Volume identifier.
	Volume *VolumeIdentity `json:"volume" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateVolumeMappingOptions : Instantiate CreateVolumeMappingOptions
func (*SdsaasV2) NewCreateVolumeMappingOptions(id string, volume *VolumeIdentity) *CreateVolumeMappingOptions {
	return &CreateVolumeMappingOptions{
		ID: core.StringPtr(id),
		Volume: volume,
	}
}

// SetID : Allow user to set ID
func (_options *CreateVolumeMappingOptions) SetID(id string) *CreateVolumeMappingOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVolume : Allow user to set Volume
func (_options *CreateVolumeMappingOptions) SetVolume(volume *VolumeIdentity) *CreateVolumeMappingOptions {
	_options.Volume = volume
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVolumeMappingOptions) SetHeaders(param map[string]string) *CreateVolumeMappingOptions {
	options.Headers = param
	return options
}

// CreateVolumeOptions : The CreateVolume options.
type CreateVolumeOptions struct {
	// The capacity to use for the volume (in gigabytes).
	//     Example:
	//       40.
	Capacity *int64 `json:"capacity" validate:"required"`

	// The name for this volume. The name must not be used by another volume.
	//     If unspecified, the name will be a hyphenated list of randomly-selected words.
	Name *string `json:"name,omitempty"`

	// Source snapshot object to restore volume.
	SourceSnapshot *SourceSnapshot `json:"source_snapshot,omitempty"`

	// Source VolumeGroupSnapshot object to restore volume from VolumeGroupSnapshot.
	SourceVolumeGroupSnapshot *SourceVolumeGroupSnapshot `json:"source_volume_group_snapshot,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateVolumeOptions : Instantiate CreateVolumeOptions
func (*SdsaasV2) NewCreateVolumeOptions(capacity int64) *CreateVolumeOptions {
	return &CreateVolumeOptions{
		Capacity: core.Int64Ptr(capacity),
	}
}

// SetCapacity : Allow user to set Capacity
func (_options *CreateVolumeOptions) SetCapacity(capacity int64) *CreateVolumeOptions {
	_options.Capacity = core.Int64Ptr(capacity)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateVolumeOptions) SetName(name string) *CreateVolumeOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetSourceSnapshot : Allow user to set SourceSnapshot
func (_options *CreateVolumeOptions) SetSourceSnapshot(sourceSnapshot *SourceSnapshot) *CreateVolumeOptions {
	_options.SourceSnapshot = sourceSnapshot
	return _options
}

// SetSourceVolumeGroupSnapshot : Allow user to set SourceVolumeGroupSnapshot
func (_options *CreateVolumeOptions) SetSourceVolumeGroupSnapshot(sourceVolumeGroupSnapshot *SourceVolumeGroupSnapshot) *CreateVolumeOptions {
	_options.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshot
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVolumeOptions) SetHeaders(param map[string]string) *CreateVolumeOptions {
	options.Headers = param
	return options
}

// DeleteHmacCredentialsOptions : The DeleteHmacCredentials options.
type DeleteHmacCredentialsOptions struct {
	// The access key to identify the HMAC credential that will be deleted in this request.
	AccessKey *string `json:"access_key" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteHmacCredentialsOptions : Instantiate DeleteHmacCredentialsOptions
func (*SdsaasV2) NewDeleteHmacCredentialsOptions(accessKey string) *DeleteHmacCredentialsOptions {
	return &DeleteHmacCredentialsOptions{
		AccessKey: core.StringPtr(accessKey),
	}
}

// SetAccessKey : Allow user to set AccessKey
func (_options *DeleteHmacCredentialsOptions) SetAccessKey(accessKey string) *DeleteHmacCredentialsOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteHmacCredentialsOptions) SetHeaders(param map[string]string) *DeleteHmacCredentialsOptions {
	options.Headers = param
	return options
}

// DeleteHostOptions : The DeleteHost options.
type DeleteHostOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteHostOptions : Instantiate DeleteHostOptions
func (*SdsaasV2) NewDeleteHostOptions(id string) *DeleteHostOptions {
	return &DeleteHostOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteHostOptions) SetID(id string) *DeleteHostOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteHostOptions) SetHeaders(param map[string]string) *DeleteHostOptions {
	options.Headers = param
	return options
}

// DeleteSnapshotOptions : The DeleteSnapshot options.
type DeleteSnapshotOptions struct {
	// The snapshot identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSnapshotOptions : Instantiate DeleteSnapshotOptions
func (*SdsaasV2) NewDeleteSnapshotOptions(id string) *DeleteSnapshotOptions {
	return &DeleteSnapshotOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteSnapshotOptions) SetID(id string) *DeleteSnapshotOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSnapshotOptions) SetHeaders(param map[string]string) *DeleteSnapshotOptions {
	options.Headers = param
	return options
}

// DeleteSnapshotsOptions : The DeleteSnapshots options.
type DeleteSnapshotsOptions struct {
	// Filters the collection to resources with a source_volume.id property matching the specified identifier.
	SourceVolumeID *string `json:"source_volume.id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSnapshotsOptions : Instantiate DeleteSnapshotsOptions
func (*SdsaasV2) NewDeleteSnapshotsOptions(sourceVolumeID string) *DeleteSnapshotsOptions {
	return &DeleteSnapshotsOptions{
		SourceVolumeID: core.StringPtr(sourceVolumeID),
	}
}

// SetSourceVolumeID : Allow user to set SourceVolumeID
func (_options *DeleteSnapshotsOptions) SetSourceVolumeID(sourceVolumeID string) *DeleteSnapshotsOptions {
	_options.SourceVolumeID = core.StringPtr(sourceVolumeID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSnapshotsOptions) SetHeaders(param map[string]string) *DeleteSnapshotsOptions {
	options.Headers = param
	return options
}

// DeleteSslCertOptions : The DeleteSslCert options.
type DeleteSslCertOptions struct {
	// The certificate type that is to be used in the DELETE request. Acceptable values include: s3.
	CertType *string `json:"cert_type" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSslCertOptions : Instantiate DeleteSslCertOptions
func (*SdsaasV2) NewDeleteSslCertOptions(certType string) *DeleteSslCertOptions {
	return &DeleteSslCertOptions{
		CertType: core.StringPtr(certType),
	}
}

// SetCertType : Allow user to set CertType
func (_options *DeleteSslCertOptions) SetCertType(certType string) *DeleteSslCertOptions {
	_options.CertType = core.StringPtr(certType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSslCertOptions) SetHeaders(param map[string]string) *DeleteSslCertOptions {
	options.Headers = param
	return options
}

// DeleteVolumeMappingOptions : The DeleteVolumeMapping options.
type DeleteVolumeMappingOptions struct {
	// The host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// The Volume mapping identifier.
	VolumeMappingID *string `json:"volume_mapping_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteVolumeMappingOptions : Instantiate DeleteVolumeMappingOptions
func (*SdsaasV2) NewDeleteVolumeMappingOptions(id string, volumeMappingID string) *DeleteVolumeMappingOptions {
	return &DeleteVolumeMappingOptions{
		ID: core.StringPtr(id),
		VolumeMappingID: core.StringPtr(volumeMappingID),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteVolumeMappingOptions) SetID(id string) *DeleteVolumeMappingOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVolumeMappingID : Allow user to set VolumeMappingID
func (_options *DeleteVolumeMappingOptions) SetVolumeMappingID(volumeMappingID string) *DeleteVolumeMappingOptions {
	_options.VolumeMappingID = core.StringPtr(volumeMappingID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVolumeMappingOptions) SetHeaders(param map[string]string) *DeleteVolumeMappingOptions {
	options.Headers = param
	return options
}

// DeleteVolumeMappingsOptions : The DeleteVolumeMappings options.
type DeleteVolumeMappingsOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteVolumeMappingsOptions : Instantiate DeleteVolumeMappingsOptions
func (*SdsaasV2) NewDeleteVolumeMappingsOptions(id string) *DeleteVolumeMappingsOptions {
	return &DeleteVolumeMappingsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteVolumeMappingsOptions) SetID(id string) *DeleteVolumeMappingsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVolumeMappingsOptions) SetHeaders(param map[string]string) *DeleteVolumeMappingsOptions {
	options.Headers = param
	return options
}

// DeleteVolumeOptions : The DeleteVolume options.
type DeleteVolumeOptions struct {
	// The volume identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteVolumeOptions : Instantiate DeleteVolumeOptions
func (*SdsaasV2) NewDeleteVolumeOptions(id string) *DeleteVolumeOptions {
	return &DeleteVolumeOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteVolumeOptions) SetID(id string) *DeleteVolumeOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVolumeOptions) SetHeaders(param map[string]string) *DeleteVolumeOptions {
	options.Headers = param
	return options
}

// ErrorObject : A set of properties used to describe an individual error that was encountered in a request including a unique code,
// message, and a link to find more information.
type ErrorObject struct {
	// A short snake_case error code to describe the error.
	Code *string `json:"code" validate:"required"`

	// The message describing the error that was encountered.
	Message *string `json:"message" validate:"required"`

	// A link to the relevant documentation that contains more information about this error.
	MoreInfo *string `json:"more_info" validate:"required"`
}

// Constants associated with the ErrorObject.Code property.
// A short snake_case error code to describe the error.
const (
	ErrorObjectCodeActionFailedConst = "action_failed"
	ErrorObjectCodeAtrackerContextUnavailableConst = "atracker_context_unavailable"
	ErrorObjectCodeCertificateExpiredConst = "certificate_expired"
	ErrorObjectCodeConfigUnavailableConst = "config_unavailable"
	ErrorObjectCodeEndpointUnavailableConst = "endpoint_unavailable"
	ErrorObjectCodeInvalidCertificateConst = "invalid_certificate"
	ErrorObjectCodeInvalidCharacterConst = "invalid_character"
	ErrorObjectCodeInvalidDateConst = "invalid_date"
	ErrorObjectCodeInvalidFormatConst = "invalid_format"
	ErrorObjectCodeInvalidKeyConst = "invalid_key"
	ErrorObjectCodeInvalidResponseConst = "invalid_response"
	ErrorObjectCodeKeyExistsConst = "key_exists"
	ErrorObjectCodeKeySizeLimitConst = "key_size_limit"
	ErrorObjectCodeMissingQueryConst = "missing_query"
	ErrorObjectCodeMultisiteNotConfiguredConst = "multisite_not_configured"
	ErrorObjectCodeUserUnauthorizedConst = "user_unauthorized"
	ErrorObjectCodeZoneNotPrimaryConst = "zone_not_primary"
)

// UnmarshalErrorObject unmarshals an instance of ErrorObject from the specified map of raw messages.
func UnmarshalErrorObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorObject)
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

// Gateway : Connection properties for the NVME gateways.
type Gateway struct {
	// IP address of the NVMe gateway node.
	IPAddress *string `json:"ip_address" validate:"required"`

	// Port number of the NVMe gateway.
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

// GetHostOptions : The GetHost options.
type GetHostOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetHostOptions : Instantiate GetHostOptions
func (*SdsaasV2) NewGetHostOptions(id string) *GetHostOptions {
	return &GetHostOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetHostOptions) SetID(id string) *GetHostOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetHostOptions) SetHeaders(param map[string]string) *GetHostOptions {
	options.Headers = param
	return options
}

// GetS3SslCertStatusOptions : The GetS3SslCertStatus options.
type GetS3SslCertStatusOptions struct {
	// The certificate type that is to be used in the request. Acceptable values include: s3.
	CertType *string `json:"cert_type" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetS3SslCertStatusOptions : Instantiate GetS3SslCertStatusOptions
func (*SdsaasV2) NewGetS3SslCertStatusOptions(certType string) *GetS3SslCertStatusOptions {
	return &GetS3SslCertStatusOptions{
		CertType: core.StringPtr(certType),
	}
}

// SetCertType : Allow user to set CertType
func (_options *GetS3SslCertStatusOptions) SetCertType(certType string) *GetS3SslCertStatusOptions {
	_options.CertType = core.StringPtr(certType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetS3SslCertStatusOptions) SetHeaders(param map[string]string) *GetS3SslCertStatusOptions {
	options.Headers = param
	return options
}

// GetSnapshotOptions : The GetSnapshot options.
type GetSnapshotOptions struct {
	// The snapshot identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSnapshotOptions : Instantiate GetSnapshotOptions
func (*SdsaasV2) NewGetSnapshotOptions(id string) *GetSnapshotOptions {
	return &GetSnapshotOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetSnapshotOptions) SetID(id string) *GetSnapshotOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSnapshotOptions) SetHeaders(param map[string]string) *GetSnapshotOptions {
	options.Headers = param
	return options
}

// GetVolumeMappingOptions : The GetVolumeMapping options.
type GetVolumeMappingOptions struct {
	// The host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// The volume mapping identifier.
	VolumeMappingID *string `json:"volume_mapping_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetVolumeMappingOptions : Instantiate GetVolumeMappingOptions
func (*SdsaasV2) NewGetVolumeMappingOptions(id string, volumeMappingID string) *GetVolumeMappingOptions {
	return &GetVolumeMappingOptions{
		ID: core.StringPtr(id),
		VolumeMappingID: core.StringPtr(volumeMappingID),
	}
}

// SetID : Allow user to set ID
func (_options *GetVolumeMappingOptions) SetID(id string) *GetVolumeMappingOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVolumeMappingID : Allow user to set VolumeMappingID
func (_options *GetVolumeMappingOptions) SetVolumeMappingID(volumeMappingID string) *GetVolumeMappingOptions {
	_options.VolumeMappingID = core.StringPtr(volumeMappingID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVolumeMappingOptions) SetHeaders(param map[string]string) *GetVolumeMappingOptions {
	options.Headers = param
	return options
}

// GetVolumeOptions : The GetVolume options.
type GetVolumeOptions struct {
	// The volume identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetVolumeOptions : Instantiate GetVolumeOptions
func (*SdsaasV2) NewGetVolumeOptions(id string) *GetVolumeOptions {
	return &GetVolumeOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetVolumeOptions) SetID(id string) *GetVolumeOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVolumeOptions) SetHeaders(param map[string]string) *GetVolumeOptions {
	options.Headers = param
	return options
}

// Host : Host Schema.
type Host struct {
	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name" validate:"required"`

	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`

	// If the PSK is specified while creating the host then the value will be true.
	PskEnabled *bool `json:"psk_enabled" validate:"required"`

	// List of volume mappings for this host.
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
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
	err = core.UnmarshalPrimitive(m, "psk_enabled", &obj.PskEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "psk_enabled-error", common.GetComponentInfo())
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

// HostCollection : Collection of host objects.
type HostCollection struct {
	// A link to the first page of resources.
	First *PageLink `json:"first" validate:"required"`

	// List of hosts.
	Hosts []Host `json:"hosts" validate:"required"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit" validate:"required"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages
	//     Example:
	//       132.
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *HostCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil {
		err = core.SDKErrorf(err, "", "read-query-param-error", common.GetComponentInfo())
		return nil, err
	} else if start == nil {
		return nil, nil
	}
	return start, nil
}

// HostPatch : Host PATCH request.
type HostPatch struct {
	// The unique name for this resource.
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
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The unique name for this resource.
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

// HostSummary : Host creation response.
type HostSummary struct {
	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name" validate:"required"`

	// The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
	Nqn *string `json:"nqn" validate:"required"`

	// If the PSK is specified while creating the host then the value will be true.
	PskEnabled *bool `json:"psk_enabled" validate:"required"`

	// List of volume mappings for this host.
	VolumeMappings []VolumeMappingReference `json:"volume_mappings" validate:"required"`
}

// UnmarshalHostSummary unmarshals an instance of HostSummary from the specified map of raw messages.
func UnmarshalHostSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostSummary)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
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
	err = core.UnmarshalPrimitive(m, "psk_enabled", &obj.PskEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "psk_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume_mappings", &obj.VolumeMappings, UnmarshalVolumeMappingReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_mappings-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCertificatesOptions : The ListCertificates options.
type ListCertificatesOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListCertificatesOptions : Instantiate ListCertificatesOptions
func (*SdsaasV2) NewListCertificatesOptions() *ListCertificatesOptions {
	return &ListCertificatesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCertificatesOptions) SetHeaders(param map[string]string) *ListCertificatesOptions {
	options.Headers = param
	return options
}

// ListHmacCredentialsOptions : The ListHmacCredentials options.
type ListHmacCredentialsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListHmacCredentialsOptions : Instantiate ListHmacCredentialsOptions
func (*SdsaasV2) NewListHmacCredentialsOptions() *ListHmacCredentialsOptions {
	return &ListHmacCredentialsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListHmacCredentialsOptions) SetHeaders(param map[string]string) *ListHmacCredentialsOptions {
	options.Headers = param
	return options
}

// ListHostsOptions : The ListHosts options.
type ListHostsOptions struct {
	// A server-provided token determining what resource to start the page on.
	Start *string `json:"start,omitempty"`

	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection to resources with a name property matching the exact specified name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListHostsOptions : Instantiate ListHostsOptions
func (*SdsaasV2) NewListHostsOptions() *ListHostsOptions {
	return &ListHostsOptions{}
}

// SetStart : Allow user to set Start
func (_options *ListHostsOptions) SetStart(start string) *ListHostsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListHostsOptions) SetLimit(limit int64) *ListHostsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListHostsOptions) SetName(name string) *ListHostsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListHostsOptions) SetHeaders(param map[string]string) *ListHostsOptions {
	options.Headers = param
	return options
}

// ListSnapshotsOptions : The ListSnapshots options.
type ListSnapshotsOptions struct {
	// A server-provided token determining what resource to start the page on.
	Start *string `json:"start,omitempty"`

	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection of resources with a name property matching the exact specified name.
	Name *string `json:"name,omitempty"`

	// Filters the collection to resources with a source_volume.id property matching the specified identifier.
	SourceVolumeID *string `json:"source_volume.id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSnapshotsOptions : Instantiate ListSnapshotsOptions
func (*SdsaasV2) NewListSnapshotsOptions() *ListSnapshotsOptions {
	return &ListSnapshotsOptions{}
}

// SetStart : Allow user to set Start
func (_options *ListSnapshotsOptions) SetStart(start string) *ListSnapshotsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListSnapshotsOptions) SetLimit(limit int64) *ListSnapshotsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListSnapshotsOptions) SetName(name string) *ListSnapshotsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetSourceVolumeID : Allow user to set SourceVolumeID
func (_options *ListSnapshotsOptions) SetSourceVolumeID(sourceVolumeID string) *ListSnapshotsOptions {
	_options.SourceVolumeID = core.StringPtr(sourceVolumeID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSnapshotsOptions) SetHeaders(param map[string]string) *ListSnapshotsOptions {
	options.Headers = param
	return options
}

// ListVolumeMappingsOptions : The ListVolumeMappings options.
type ListVolumeMappingsOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// A server-provided token determining what resource to start the page on.
	Start *string `json:"start,omitempty"`

	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection of resources with a name property matching the exact specified name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListVolumeMappingsOptions : Instantiate ListVolumeMappingsOptions
func (*SdsaasV2) NewListVolumeMappingsOptions(id string) *ListVolumeMappingsOptions {
	return &ListVolumeMappingsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListVolumeMappingsOptions) SetID(id string) *ListVolumeMappingsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListVolumeMappingsOptions) SetStart(start string) *ListVolumeMappingsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListVolumeMappingsOptions) SetLimit(limit int64) *ListVolumeMappingsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListVolumeMappingsOptions) SetName(name string) *ListVolumeMappingsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVolumeMappingsOptions) SetHeaders(param map[string]string) *ListVolumeMappingsOptions {
	options.Headers = param
	return options
}

// ListVolumesOptions : The ListVolumes options.
type ListVolumesOptions struct {
	// A server-provided token determining what resource to start the page on.
	Start *string `json:"start,omitempty"`

	// The number of resources to return on a page.
	Limit *int64 `json:"limit,omitempty"`

	// Filters the collection of resources with a name property matching the exact specified name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListVolumesOptions : Instantiate ListVolumesOptions
func (*SdsaasV2) NewListVolumesOptions() *ListVolumesOptions {
	return &ListVolumesOptions{}
}

// SetStart : Allow user to set Start
func (_options *ListVolumesOptions) SetStart(start string) *ListVolumesOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListVolumesOptions) SetLimit(limit int64) *ListVolumesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListVolumesOptions) SetName(name string) *ListVolumesOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVolumesOptions) SetHeaders(param map[string]string) *ListVolumesOptions {
	options.Headers = param
	return options
}

// Namespace : The NVMe namespace properties for a given volume mapping.
type Namespace struct {
	// NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
	ID *int64 `json:"id" validate:"required"`

	// UUID of the NVMe namespace.
	UUID *string `json:"uuid" validate:"required"`
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
	Href *string `json:"href,omitempty"`
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

// ReplaceSslCertOptions : The ReplaceSslCert options.
type ReplaceSslCertOptions struct {
	// The certificate type that is to be used in the PUT request. Acceptable values include: s3.
	CertType *string `json:"cert_type" validate:"required,ne="`

	// A string containing at least one PEM formatted TLS Certificate and a PEM formatted TLS Key for the PUT request.
	// These should be formatted according to the RFC-7468 standard, with the CERTIFICATE property being used for
	// certificates and any of the following properties for keys: PRIVATE KEY, RSA PRIVATE KEY, EC PRIVATE KEY. If a
	// provided certificate or key is not in the specified formats, the API will reject it.
	Body io.ReadCloser `json:"body,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceSslCertOptions : Instantiate ReplaceSslCertOptions
func (*SdsaasV2) NewReplaceSslCertOptions(certType string) *ReplaceSslCertOptions {
	return &ReplaceSslCertOptions{
		CertType: core.StringPtr(certType),
	}
}

// SetCertType : Allow user to set CertType
func (_options *ReplaceSslCertOptions) SetCertType(certType string) *ReplaceSslCertOptions {
	_options.CertType = core.StringPtr(certType)
	return _options
}

// SetBody : Allow user to set Body
func (_options *ReplaceSslCertOptions) SetBody(body io.ReadCloser) *ReplaceSslCertOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceSslCertOptions) SetHeaders(param map[string]string) *ReplaceSslCertOptions {
	options.Headers = param
	return options
}

// Snapshot : The snapshot object.
type Snapshot struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name" validate:"required"`

	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The type of this resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// The lifecycle state of this snapshot.
	LifecycleState *string `json:"lifecycle_state" validate:"required"`

	// The size of the snapshot (in gigabytes).
	Size *int64 `json:"size" validate:"required"`

	// The capacity of the source volume.
	MinimumCapacity *int64 `json:"minimum_capacity" validate:"required"`

	// Boolean value of snapshot can be deletable or not.
	Deletable *bool `json:"deletable" validate:"required"`

	// The source volume object of this snapshot should be created.
	SourceVolume *SourceVolume `json:"source_volume,omitempty"`
}

// UnmarshalSnapshot unmarshals an instance of Snapshot from the specified map of raw messages.
func UnmarshalSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Snapshot)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "lifecycle_state", &obj.LifecycleState)
	if err != nil {
		err = core.SDKErrorf(err, "", "lifecycle_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		err = core.SDKErrorf(err, "", "size-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_capacity", &obj.MinimumCapacity)
	if err != nil {
		err = core.SDKErrorf(err, "", "minimum_capacity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "deletable", &obj.Deletable)
	if err != nil {
		err = core.SDKErrorf(err, "", "deletable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "source_volume", &obj.SourceVolume, UnmarshalSourceVolume)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_volume-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SnapshotCollection : A page of snapshots.<br> In case of single volume snapshot, the snapshot response will include the source_volume
// details.<br> In case of multi volume group snapshot, the snapshot response will include source_volume_group
// details.<br>.
type SnapshotCollection struct {
	// Collection of snapshots.
	Snapshots []Snapshot `json:"snapshots" validate:"required"`

	// A link to the first page of resources.
	First *PageLink `json:"first" validate:"required"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit" validate:"required"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages
	//     Example:
	//       132.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalSnapshotCollection unmarshals an instance of SnapshotCollection from the specified map of raw messages.
func UnmarshalSnapshotCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SnapshotCollection)
	err = core.UnmarshalModel(m, "snapshots", &obj.Snapshots, UnmarshalSnapshot)
	if err != nil {
		err = core.SDKErrorf(err, "", "snapshots-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *SnapshotCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil {
		err = core.SDKErrorf(err, "", "read-query-param-error", common.GetComponentInfo())
		return nil, err
	} else if start == nil {
		return nil, nil
	}
	return start, nil
}

// SnapshotPatch : The snapshot patch.
type SnapshotPatch struct {
	// The name for this snapshot. The name must not be used by another snapshot.
	Name *string `json:"name,omitempty"`
}

// UnmarshalSnapshotPatch unmarshals an instance of SnapshotPatch from the specified map of raw messages.
func UnmarshalSnapshotPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SnapshotPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the SnapshotPatch
func (snapshotPatch *SnapshotPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(snapshotPatch.Name) {
		_patch["name"] = snapshotPatch.Name
	}

	return
}

// SourceSnapshot : Source snapshot object to restore volume.
type SourceSnapshot struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`
}

// NewSourceSnapshot : Instantiate SourceSnapshot (Generic Model Constructor)
func (*SdsaasV2) NewSourceSnapshot(id string) (_model *SourceSnapshot, err error) {
	_model = &SourceSnapshot{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSourceSnapshot unmarshals an instance of SourceSnapshot from the specified map of raw messages.
func UnmarshalSourceSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceSnapshot)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceVolume : The source volume object of this snapshot should be created.
type SourceVolume struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name,omitempty"`

	// The type of this resource.
	ResourceType *string `json:"resource_type,omitempty"`
}

// UnmarshalSourceVolume unmarshals an instance of SourceVolume from the specified map of raw messages.
func UnmarshalSourceVolume(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceVolume)
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
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceVolumeGroupSnapshot : Source VolumeGroupSnapshot object to restore volume from VolumeGroupSnapshot.
type SourceVolumeGroupSnapshot struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The volume to restore from VolumeGroupSnapshot.
	Volume *SourceVolumeGroupSnapshotVolume `json:"volume" validate:"required"`
}

// NewSourceVolumeGroupSnapshot : Instantiate SourceVolumeGroupSnapshot (Generic Model Constructor)
func (*SdsaasV2) NewSourceVolumeGroupSnapshot(id string, volume *SourceVolumeGroupSnapshotVolume) (_model *SourceVolumeGroupSnapshot, err error) {
	_model = &SourceVolumeGroupSnapshot{
		ID: core.StringPtr(id),
		Volume: volume,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSourceVolumeGroupSnapshot unmarshals an instance of SourceVolumeGroupSnapshot from the specified map of raw messages.
func UnmarshalSourceVolumeGroupSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceVolumeGroupSnapshot)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume", &obj.Volume, UnmarshalSourceVolumeGroupSnapshotVolume)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceVolumeGroupSnapshotVolume : The volume to restore from VolumeGroupSnapshot.
type SourceVolumeGroupSnapshotVolume struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`
}

// NewSourceVolumeGroupSnapshotVolume : Instantiate SourceVolumeGroupSnapshotVolume (Generic Model Constructor)
func (*SdsaasV2) NewSourceVolumeGroupSnapshotVolume(id string) (_model *SourceVolumeGroupSnapshotVolume, err error) {
	_model = &SourceVolumeGroupSnapshotVolume{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSourceVolumeGroupSnapshotVolume unmarshals an instance of SourceVolumeGroupSnapshotVolume from the specified map of raw messages.
func UnmarshalSourceVolumeGroupSnapshotVolume(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceVolumeGroupSnapshotVolume)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceVolumePrototype : The source volume this snapshot was created from (may be deleted).
type SourceVolumePrototype struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`
}

// NewSourceVolumePrototype : Instantiate SourceVolumePrototype (Generic Model Constructor)
func (*SdsaasV2) NewSourceVolumePrototype(id string) (_model *SourceVolumePrototype, err error) {
	_model = &SourceVolumePrototype{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSourceVolumePrototype unmarshals an instance of SourceVolumePrototype from the specified map of raw messages.
func UnmarshalSourceVolumePrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceVolumePrototype)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusResponse : The expiration status and expiration date of the current certificate.
type StatusResponse struct {
	// The expiration date for the current certificate.
	ExpirationDate *strfmt.DateTime `json:"expiration_date,omitempty"`

	// When set to true, indicates that the current certificate is expired.
	Expired *bool `json:"expired,omitempty"`

	// Name of the certificate.
	Name *string `json:"name,omitempty"`
}

// UnmarshalStatusResponse unmarshals an instance of StatusResponse from the specified map of raw messages.
func UnmarshalStatusResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusResponse)
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
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageCredResponse : The list of HMAC credential access keys.
type StorageCredResponse struct {
	// An array of HMAC Credential access keys.
	S3Credentials []string `json:"s3_credentials" validate:"required"`
}

// UnmarshalStorageCredResponse unmarshals an instance of StorageCredResponse from the specified map of raw messages.
func UnmarshalStorageCredResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageCredResponse)
	err = core.UnmarshalPrimitive(m, "s3_credentials", &obj.S3Credentials)
	if err != nil {
		err = core.SDKErrorf(err, "", "s3_credentials-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateHostOptions : The UpdateHost options.
type UpdateHostOptions struct {
	// The Host identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Host Patch body.
	HostPatch map[string]interface{} `json:"Host_patch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateHostOptions : Instantiate UpdateHostOptions
func (*SdsaasV2) NewUpdateHostOptions(id string, hostPatch map[string]interface{}) *UpdateHostOptions {
	return &UpdateHostOptions{
		ID: core.StringPtr(id),
		HostPatch: hostPatch,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateHostOptions) SetID(id string) *UpdateHostOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHostPatch : Allow user to set HostPatch
func (_options *UpdateHostOptions) SetHostPatch(hostPatch map[string]interface{}) *UpdateHostOptions {
	_options.HostPatch = hostPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateHostOptions) SetHeaders(param map[string]string) *UpdateHostOptions {
	options.Headers = param
	return options
}

// UpdateSnapshotOptions : The UpdateSnapshot options.
type UpdateSnapshotOptions struct {
	// The snapshot identifier.
	ID *string `json:"id" validate:"required,ne="`

	// The name for this snapshot. The name must not be used by another snapshot.
	SnapshotPatch map[string]interface{} `json:"Snapshot_patch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateSnapshotOptions : Instantiate UpdateSnapshotOptions
func (*SdsaasV2) NewUpdateSnapshotOptions(id string, snapshotPatch map[string]interface{}) *UpdateSnapshotOptions {
	return &UpdateSnapshotOptions{
		ID: core.StringPtr(id),
		SnapshotPatch: snapshotPatch,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateSnapshotOptions) SetID(id string) *UpdateSnapshotOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetSnapshotPatch : Allow user to set SnapshotPatch
func (_options *UpdateSnapshotOptions) SetSnapshotPatch(snapshotPatch map[string]interface{}) *UpdateSnapshotOptions {
	_options.SnapshotPatch = snapshotPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSnapshotOptions) SetHeaders(param map[string]string) *UpdateSnapshotOptions {
	options.Headers = param
	return options
}

// UpdateVolumeOptions : The UpdateVolume options.
type UpdateVolumeOptions struct {
	// The volume identifier.
	ID *string `json:"id" validate:"required,ne="`

	// Volume patch request body.
	VolumePatch map[string]interface{} `json:"Volume_patch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateVolumeOptions : Instantiate UpdateVolumeOptions
func (*SdsaasV2) NewUpdateVolumeOptions(id string, volumePatch map[string]interface{}) *UpdateVolumeOptions {
	return &UpdateVolumeOptions{
		ID: core.StringPtr(id),
		VolumePatch: volumePatch,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateVolumeOptions) SetID(id string) *UpdateVolumeOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVolumePatch : Allow user to set VolumePatch
func (_options *UpdateVolumeOptions) SetVolumePatch(volumePatch map[string]interface{}) *UpdateVolumeOptions {
	_options.VolumePatch = volumePatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateVolumeOptions) SetHeaders(param map[string]string) *UpdateVolumeOptions {
	options.Headers = param
	return options
}

// Volume : The volume object.
type Volume struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name" validate:"required"`

	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The type of this resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// The capacity of the volume (in gigabytes).
	Capacity *int64 `json:"capacity" validate:"required"`

	// The number of snapshots of the volume or volume group.
	SnapshotCount *int64 `json:"snapshot_count,omitempty"`

	// The maximum bandwidth (in megabits per second) for the volume
	//     Example:
	//       1000.
	Bandwidth *int64 `json:"bandwidth" validate:"required"`

	// The maximum I/O operations per second (IOPS) for this volume.
	//     Example:
	//       10000.
	Iops *int64 `json:"iops" validate:"required"`

	// List of volume mappings for this volume.
	VolumeMappings []VolumeMapping `json:"volume_mappings" validate:"required"`

	// The status of the volume resource. The enumerated values for this property will expand in the future. When
	// processing this property, check for and log unknown values. Optionally halt processing and surface the error, or
	// bypass the resource on which the unexpected property value was encountered.
	Status *string `json:"status" validate:"required"`

	// The reasons for the current status (if any).
	StatusReasons []VolumeStatusReason `json:"status_reasons" validate:"required"`

	// Source snapshot object to restore volume.
	SourceSnapshot *SourceSnapshot `json:"source_snapshot" validate:"required"`

	// The unique identifier for this resource.
	VolumeGroup *string `json:"volume_group,omitempty"`
}

// UnmarshalVolume unmarshals an instance of Volume from the specified map of raw messages.
func UnmarshalVolume(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Volume)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "capacity", &obj.Capacity)
	if err != nil {
		err = core.SDKErrorf(err, "", "capacity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "snapshot_count", &obj.SnapshotCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "snapshot_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bandwidth", &obj.Bandwidth)
	if err != nil {
		err = core.SDKErrorf(err, "", "bandwidth-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iops", &obj.Iops)
	if err != nil {
		err = core.SDKErrorf(err, "", "iops-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volume_mappings", &obj.VolumeMappings, UnmarshalVolumeMapping)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_mappings-error", common.GetComponentInfo())
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
	err = core.UnmarshalModel(m, "source_snapshot", &obj.SourceSnapshot, UnmarshalSourceSnapshot)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_snapshot-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "volume_group", &obj.VolumeGroup)
	if err != nil {
		err = core.SDKErrorf(err, "", "volume_group-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeCollection : A page of volumes.
type VolumeCollection struct {
	// Collection of volumes.
	Volumes []Volume `json:"volumes" validate:"required"`

	// A link to the first page of resources.
	First *PageLink `json:"first" validate:"required"`

	// The maximum number of resources that can be returned by the request.
	Limit *int64 `json:"limit" validate:"required"`

	// A link to the next page of resources. This property is present for all pages except the last page.
	Next *PageLink `json:"next,omitempty"`

	// The total number of resources across all pages
	//     Example:
	//       132.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalVolumeCollection unmarshals an instance of VolumeCollection from the specified map of raw messages.
func UnmarshalVolumeCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeCollection)
	err = core.UnmarshalModel(m, "volumes", &obj.Volumes, UnmarshalVolume)
	if err != nil {
		err = core.SDKErrorf(err, "", "volumes-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *VolumeCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil {
		err = core.SDKErrorf(err, "", "read-query-param-error", common.GetComponentInfo())
		return nil, err
	} else if start == nil {
		return nil, nil
	}
	return start, nil
}

// VolumeIdentity : Volume identifier.
type VolumeIdentity struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`
}

// NewVolumeIdentity : Instantiate VolumeIdentity (Generic Model Constructor)
func (*SdsaasV2) NewVolumeIdentity(id string) (_model *VolumeIdentity, err error) {
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

// VolumeMapping : Volume mapping object depicting the mapping between a volume and a host.
type VolumeMapping struct {
	// The status of the volume mapping. The enumerated values for this property will expand in the future. When processing
	// this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the
	// resource on which the unexpected property value was encountered.
	Status *string `json:"status" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// Unique identifier of the mapping.
	ID *string `json:"id" validate:"required"`

	// The volume reference.
	Volume *VolumeReference `json:"volume" validate:"required"`

	// Host mapping schema.
	Host *HostReference `json:"host" validate:"required"`

	// The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
	SubsystemNqn *string `json:"subsystem_nqn,omitempty"`

	// The NVMe namespace properties for a given volume mapping.
	Namespace *Namespace `json:"namespace,omitempty"`

	// List of NVMe gateways.
	Gateways []Gateway `json:"gateways" validate:"required"`
}

// UnmarshalVolumeMapping unmarshals an instance of VolumeMapping from the specified map of raw messages.
func UnmarshalVolumeMapping(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMapping)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
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

	// The total number of resources across all pages
	//     Example:
	//       132.
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *VolumeMappingCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil {
		err = core.SDKErrorf(err, "", "read-query-param-error", common.GetComponentInfo())
		return nil, err
	} else if start == nil {
		return nil, nil
	}
	return start, nil
}

// VolumeMappingPrototype : The volume mapping request.
type VolumeMappingPrototype struct {
	// Volume identifier.
	Volume *VolumeIdentity `json:"volume" validate:"required"`
}

// NewVolumeMappingPrototype : Instantiate VolumeMappingPrototype (Generic Model Constructor)
func (*SdsaasV2) NewVolumeMappingPrototype(volume *VolumeIdentity) (_model *VolumeMappingPrototype, err error) {
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

// VolumeMappingReference : Schema for a volume mapping to a host in pending state.
type VolumeMappingReference struct {
	// The status of the volume mapping. The enumerated values for this property will expand in the future. When processing
	// this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the
	// resource on which the unexpected property value was encountered.
	Status *string `json:"status" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// Unique identifier of the mapping.
	ID *string `json:"id" validate:"required"`

	// The volume reference.
	Volume *VolumeReference `json:"volume" validate:"required"`

	// Host mapping schema.
	Host *HostReference `json:"host" validate:"required"`
}

// UnmarshalVolumeMappingReference unmarshals an instance of VolumeMappingReference from the specified map of raw messages.
func UnmarshalVolumeMappingReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMappingReference)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumePatch : Volume patch body schema.
type VolumePatch struct {
	// The capacity to use for the volume (in gigabytes). Additionally, the specified value must not be less than the
	// current capacity.
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
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The unique name for this resource.
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

// VolumeSummary : The volume object.
type VolumeSummary struct {
	// The unique identifier for this resource.
	ID *string `json:"id" validate:"required"`

	// The URL for this resource.
	Href *string `json:"href" validate:"required"`

	// The unique name for this resource.
	Name *string `json:"name" validate:"required"`

	// The date and time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The type of this resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// The capacity of the volume (in gigabytes).
	Capacity *int64 `json:"capacity" validate:"required"`

	// The maximum bandwidth (in megabits per second) for the volume
	//        Example:
	//          1000.
	Bandwidth *int64 `json:"bandwidth" validate:"required"`

	// The maximum I/O operations per second (IOPS) for this volume.
	//        Example:
	//          10000.
	Iops *int64 `json:"iops" validate:"required"`

	// The status of the volume resource. The enumerated values for this property will expand in the future. When
	// processing this property, check for and log unknown values. Optionally halt processing and surface the error, or
	// bypass the resource on which the unexpected property value was encountered.
	Status *string `json:"status" validate:"required"`

	// The reasons for the current status (if any).
	StatusReasons []VolumeStatusReason `json:"status_reasons" validate:"required"`
}

// UnmarshalVolumeSummary unmarshals an instance of VolumeSummary from the specified map of raw messages.
func UnmarshalVolumeSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "capacity", &obj.Capacity)
	if err != nil {
		err = core.SDKErrorf(err, "", "capacity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bandwidth", &obj.Bandwidth)
	if err != nil {
		err = core.SDKErrorf(err, "", "bandwidth-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iops", &obj.Iops)
	if err != nil {
		err = core.SDKErrorf(err, "", "iops-error", common.GetComponentInfo())
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// VolumesPager can be used to simplify the use of the "ListVolumes" method.
//
type VolumesPager struct {
	hasNext bool
	options *ListVolumesOptions
	client  *SdsaasV2
	pageContext struct {
		next *string
	}
}

// NewVolumesPager returns a new VolumesPager instance.
func (sdsaas *SdsaasV2) NewVolumesPager(options *ListVolumesOptions) (pager *VolumesPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListVolumesOptions = *options
	pager = &VolumesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  sdsaas,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *VolumesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *VolumesPager) GetNextWithContext(ctx context.Context) (page []Volume, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListVolumesWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			err = core.SDKErrorf(err, errMsg, "get-query-error", common.GetComponentInfo())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Volumes

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *VolumesPager) GetAllWithContext(ctx context.Context) (allItems []Volume, err error) {
	for pager.HasNext() {
		var nextPage []Volume
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *VolumesPager) GetNext() (page []Volume, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *VolumesPager) GetAll() (allItems []Volume, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

//
// HostsPager can be used to simplify the use of the "ListHosts" method.
//
type HostsPager struct {
	hasNext bool
	options *ListHostsOptions
	client  *SdsaasV2
	pageContext struct {
		next *string
	}
}

// NewHostsPager returns a new HostsPager instance.
func (sdsaas *SdsaasV2) NewHostsPager(options *ListHostsOptions) (pager *HostsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListHostsOptions = *options
	pager = &HostsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  sdsaas,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *HostsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *HostsPager) GetNextWithContext(ctx context.Context) (page []Host, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListHostsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			err = core.SDKErrorf(err, errMsg, "get-query-error", common.GetComponentInfo())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Hosts

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *HostsPager) GetAllWithContext(ctx context.Context) (allItems []Host, err error) {
	for pager.HasNext() {
		var nextPage []Host
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *HostsPager) GetNext() (page []Host, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *HostsPager) GetAll() (allItems []Host, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

//
// VolumeMappingsPager can be used to simplify the use of the "ListVolumeMappings" method.
//
type VolumeMappingsPager struct {
	hasNext bool
	options *ListVolumeMappingsOptions
	client  *SdsaasV2
	pageContext struct {
		next *string
	}
}

// NewVolumeMappingsPager returns a new VolumeMappingsPager instance.
func (sdsaas *SdsaasV2) NewVolumeMappingsPager(options *ListVolumeMappingsOptions) (pager *VolumeMappingsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListVolumeMappingsOptions = *options
	pager = &VolumeMappingsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  sdsaas,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *VolumeMappingsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *VolumeMappingsPager) GetNextWithContext(ctx context.Context) (page []VolumeMapping, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListVolumeMappingsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			err = core.SDKErrorf(err, errMsg, "get-query-error", common.GetComponentInfo())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.VolumeMappings

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *VolumeMappingsPager) GetAllWithContext(ctx context.Context) (allItems []VolumeMapping, err error) {
	for pager.HasNext() {
		var nextPage []VolumeMapping
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *VolumeMappingsPager) GetNext() (page []VolumeMapping, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *VolumeMappingsPager) GetAll() (allItems []VolumeMapping, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

//
// SnapshotsPager can be used to simplify the use of the "ListSnapshots" method.
//
type SnapshotsPager struct {
	hasNext bool
	options *ListSnapshotsOptions
	client  *SdsaasV2
	pageContext struct {
		next *string
	}
}

// NewSnapshotsPager returns a new SnapshotsPager instance.
func (sdsaas *SdsaasV2) NewSnapshotsPager(options *ListSnapshotsOptions) (pager *SnapshotsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListSnapshotsOptions = *options
	pager = &SnapshotsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  sdsaas,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *SnapshotsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *SnapshotsPager) GetNextWithContext(ctx context.Context) (page []Snapshot, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListSnapshotsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			err = core.SDKErrorf(err, errMsg, "get-query-error", common.GetComponentInfo())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Snapshots

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *SnapshotsPager) GetAllWithContext(ctx context.Context) (allItems []Snapshot, err error) {
	for pager.HasNext() {
		var nextPage []Snapshot
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *SnapshotsPager) GetNext() (page []Snapshot, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *SnapshotsPager) GetAll() (allItems []Snapshot, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
