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

package sdsaasv2_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/v2/sdsaasv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SdsaasV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(sdsaasService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(sdsaasService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
				URL: "https://sdsaasv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(sdsaasService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SDSAAS_URL": "https://sdsaasv2/api",
				"SDSAAS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2UsingExternalConfig(&sdsaasv2.SdsaasV2Options{
				})
				Expect(sdsaasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := sdsaasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sdsaasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sdsaasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sdsaasService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2UsingExternalConfig(&sdsaasv2.SdsaasV2Options{
					URL: "https://testService/api",
				})
				Expect(sdsaasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sdsaasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sdsaasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sdsaasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sdsaasService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2UsingExternalConfig(&sdsaasv2.SdsaasV2Options{
				})
				err := sdsaasService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sdsaasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sdsaasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sdsaasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sdsaasService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SDSAAS_URL": "https://sdsaasv2/api",
				"SDSAAS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2UsingExternalConfig(&sdsaasv2.SdsaasV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(sdsaasService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SDSAAS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2UsingExternalConfig(&sdsaasv2.SdsaasV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(sdsaasService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = sdsaasv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := sdsaasv2.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("{url}"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := sdsaasv2.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListVolumes(listVolumesOptions *ListVolumesOptions) - Operation response error`, func() {
		listVolumesPath := "/volumes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVolumes with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(sdsaasv2.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumesOptionsModel.Name = core.StringPtr("my-resource")
				listVolumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVolumes(listVolumesOptions *ListVolumesOptions)`, func() {
		listVolumesPath := "/volumes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"volumes": [{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListVolumes successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(sdsaasv2.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumesOptionsModel.Name = core.StringPtr("my-resource")
				listVolumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListVolumesWithContext(ctx, listVolumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListVolumesWithContext(ctx, listVolumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"volumes": [{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListVolumes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListVolumes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(sdsaasv2.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumesOptionsModel.Name = core.StringPtr("my-resource")
				listVolumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVolumes with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(sdsaasv2.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumesOptionsModel.Name = core.StringPtr("my-resource")
				listVolumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVolumes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(sdsaasv2.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumesOptionsModel.Name = core.StringPtr("my-resource")
				listVolumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(sdsaasv2.VolumeCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv2.VolumeCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv2.VolumeCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1,"volumes":[{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","created_at":"2019-01-01T12:00:00.000Z","resource_type":"ResourceType","capacity":30,"snapshot_count":10,"bandwidth":1,"iops":150,"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}],"status":"available","status_reasons":[{"code":"Code","message":"Specified resource not found","more_info":"MoreInfo"}],"source_snapshot":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5"},"volume_group":"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"volumes":[{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","created_at":"2019-01-01T12:00:00.000Z","resource_type":"ResourceType","capacity":30,"snapshot_count":10,"bandwidth":1,"iops":150,"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}],"status":"available","status_reasons":[{"code":"Code","message":"Specified resource not found","more_info":"MoreInfo"}],"source_snapshot":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5"},"volume_group":"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use VolumesPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listVolumesOptionsModel := &sdsaasv2.ListVolumesOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewVolumesPager(listVolumesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv2.Volume
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use VolumesPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listVolumesOptionsModel := &sdsaasv2.ListVolumesOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewVolumesPager(listVolumesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateVolume(createVolumeOptions *CreateVolumeOptions) - Operation response error`, func() {
		createVolumePath := "/volumes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVolume with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(sdsaasv2.CreateVolumeOptions)
				createVolumeOptionsModel.Capacity = core.Int64Ptr(int64(1))
				createVolumeOptionsModel.Name = core.StringPtr("my-volume")
				createVolumeOptionsModel.SourceSnapshot = sourceSnapshotModel
				createVolumeOptionsModel.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshotModel
				createVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVolume(createVolumeOptions *CreateVolumeOptions)`, func() {
		createVolumePath := "/volumes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "bandwidth": 1, "iops": 150, "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}]}`)
				}))
			})
			It(`Invoke CreateVolume successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(sdsaasv2.CreateVolumeOptions)
				createVolumeOptionsModel.Capacity = core.Int64Ptr(int64(1))
				createVolumeOptionsModel.Name = core.StringPtr("my-volume")
				createVolumeOptionsModel.SourceSnapshot = sourceSnapshotModel
				createVolumeOptionsModel.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshotModel
				createVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateVolumeWithContext(ctx, createVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateVolumeWithContext(ctx, createVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "bandwidth": 1, "iops": 150, "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}]}`)
				}))
			})
			It(`Invoke CreateVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(sdsaasv2.CreateVolumeOptions)
				createVolumeOptionsModel.Capacity = core.Int64Ptr(int64(1))
				createVolumeOptionsModel.Name = core.StringPtr("my-volume")
				createVolumeOptionsModel.SourceSnapshot = sourceSnapshotModel
				createVolumeOptionsModel.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshotModel
				createVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVolume with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(sdsaasv2.CreateVolumeOptions)
				createVolumeOptionsModel.Capacity = core.Int64Ptr(int64(1))
				createVolumeOptionsModel.Name = core.StringPtr("my-volume")
				createVolumeOptionsModel.SourceSnapshot = sourceSnapshotModel
				createVolumeOptionsModel.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshotModel
				createVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVolumeOptions model with no property values
				createVolumeOptionsModelNew := new(sdsaasv2.CreateVolumeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CreateVolume(createVolumeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(sdsaasv2.CreateVolumeOptions)
				createVolumeOptionsModel.Capacity = core.Int64Ptr(int64(1))
				createVolumeOptionsModel.Name = core.StringPtr("my-volume")
				createVolumeOptionsModel.SourceSnapshot = sourceSnapshotModel
				createVolumeOptionsModel.SourceVolumeGroupSnapshot = sourceVolumeGroupSnapshotModel
				createVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVolume(deleteVolumeOptions *DeleteVolumeOptions)`, func() {
		deleteVolumePath := "/volumes/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVolumePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVolumeOptions model
				deleteVolumeOptionsModel := new(sdsaasv2.DeleteVolumeOptions)
				deleteVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteVolume(deleteVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVolume with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteVolumeOptions model
				deleteVolumeOptionsModel := new(sdsaasv2.DeleteVolumeOptions)
				deleteVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteVolume(deleteVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVolumeOptions model with no property values
				deleteVolumeOptionsModelNew := new(sdsaasv2.DeleteVolumeOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteVolume(deleteVolumeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVolume(getVolumeOptions *GetVolumeOptions) - Operation response error`, func() {
		getVolumePath := "/volumes/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVolume with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(sdsaasv2.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVolume(getVolumeOptions *GetVolumeOptions)`, func() {
		getVolumePath := "/volumes/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}`)
				}))
			})
			It(`Invoke GetVolume successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(sdsaasv2.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.GetVolumeWithContext(ctx, getVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.GetVolumeWithContext(ctx, getVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}`)
				}))
			})
			It(`Invoke GetVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.GetVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(sdsaasv2.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVolume with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(sdsaasv2.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVolumeOptions model with no property values
				getVolumeOptionsModelNew := new(sdsaasv2.GetVolumeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.GetVolume(getVolumeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(sdsaasv2.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVolume(updateVolumeOptions *UpdateVolumeOptions) - Operation response error`, func() {
		updateVolumePath := "/volumes/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVolumePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateVolume with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv2.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("my-volume")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(sdsaasv2.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.VolumePatch = volumePatchModelAsPatch
				updateVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVolume(updateVolumeOptions *UpdateVolumeOptions)`, func() {
		updateVolumePath := "/volumes/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVolumePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}`)
				}))
			})
			It(`Invoke UpdateVolume successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv2.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("my-volume")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(sdsaasv2.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.VolumePatch = volumePatchModelAsPatch
				updateVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.UpdateVolumeWithContext(ctx, updateVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.UpdateVolumeWithContext(ctx, updateVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVolumePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "capacity": 30, "snapshot_count": 10, "bandwidth": 1, "iops": 150, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "status": "available", "status_reasons": [{"code": "Code", "message": "Specified resource not found", "more_info": "MoreInfo"}], "source_snapshot": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}, "volume_group": "r134-b274-678d-4dfb-8981-c71dd9d4daa5"}`)
				}))
			})
			It(`Invoke UpdateVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.UpdateVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv2.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("my-volume")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(sdsaasv2.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.VolumePatch = volumePatchModelAsPatch
				updateVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateVolume with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv2.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("my-volume")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(sdsaasv2.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.VolumePatch = volumePatchModelAsPatch
				updateVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateVolumeOptions model with no property values
				updateVolumeOptionsModelNew := new(sdsaasv2.UpdateVolumeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.UpdateVolume(updateVolumeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateVolume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv2.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("my-volume")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(sdsaasv2.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.VolumePatch = volumePatchModelAsPatch
				updateVolumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListHosts(listHostsOptions *ListHostsOptions) - Operation response error`, func() {
		listHostsPath := "/hosts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHostsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListHosts with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := new(sdsaasv2.ListHostsOptions)
				listHostsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listHostsOptionsModel.Name = core.StringPtr("my-resource")
				listHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListHosts(listHostsOptions *ListHostsOptions)`, func() {
		listHostsPath := "/hosts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHostsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListHosts successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := new(sdsaasv2.ListHostsOptions)
				listHostsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listHostsOptionsModel.Name = core.StringPtr("my-resource")
				listHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListHostsWithContext(ctx, listHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListHostsWithContext(ctx, listHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHostsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListHosts successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListHosts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := new(sdsaasv2.ListHostsOptions)
				listHostsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listHostsOptionsModel.Name = core.StringPtr("my-resource")
				listHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListHosts with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := new(sdsaasv2.ListHostsOptions)
				listHostsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listHostsOptionsModel.Name = core.StringPtr("my-resource")
				listHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListHosts successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := new(sdsaasv2.ListHostsOptions)
				listHostsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listHostsOptionsModel.Name = core.StringPtr("my-resource")
				listHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListHosts(listHostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(sdsaasv2.HostCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv2.HostCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv2.HostCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHostsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"hosts":[{"created_at":"2019-01-01T12:00:00.000Z","id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","nqn":"nqn.2014-06.org:1234","psk_enabled":true,"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}]}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"hosts":[{"created_at":"2019-01-01T12:00:00.000Z","id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","nqn":"nqn.2014-06.org:1234","psk_enabled":true,"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}]}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use HostsPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listHostsOptionsModel := &sdsaasv2.ListHostsOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewHostsPager(listHostsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv2.Host
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use HostsPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listHostsOptionsModel := &sdsaasv2.ListHostsOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewHostsPager(listHostsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateHost(createHostOptions *CreateHostOptions) - Operation response error`, func() {
		createHostPath := "/hosts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHostPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateHost with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the CreateHostOptions model
				createHostOptionsModel := new(sdsaasv2.CreateHostOptions)
				createHostOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:1234")
				createHostOptionsModel.Name = core.StringPtr("my-host")
				createHostOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.VolumeMappings = []sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				createHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateHost(createHostOptions *CreateHostOptions)`, func() {
		createHostPath := "/hosts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHostPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "pending", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}}]}`)
				}))
			})
			It(`Invoke CreateHost successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the CreateHostOptions model
				createHostOptionsModel := new(sdsaasv2.CreateHostOptions)
				createHostOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:1234")
				createHostOptionsModel.Name = core.StringPtr("my-host")
				createHostOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.VolumeMappings = []sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				createHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateHostWithContext(ctx, createHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateHostWithContext(ctx, createHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHostPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "pending", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}}]}`)
				}))
			})
			It(`Invoke CreateHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateHost(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the CreateHostOptions model
				createHostOptionsModel := new(sdsaasv2.CreateHostOptions)
				createHostOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:1234")
				createHostOptionsModel.Name = core.StringPtr("my-host")
				createHostOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.VolumeMappings = []sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				createHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateHost with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the CreateHostOptions model
				createHostOptionsModel := new(sdsaasv2.CreateHostOptions)
				createHostOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:1234")
				createHostOptionsModel.Name = core.StringPtr("my-host")
				createHostOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.VolumeMappings = []sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				createHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateHostOptions model with no property values
				createHostOptionsModelNew := new(sdsaasv2.CreateHostOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CreateHost(createHostOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the CreateHostOptions model
				createHostOptionsModel := new(sdsaasv2.CreateHostOptions)
				createHostOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:1234")
				createHostOptionsModel.Name = core.StringPtr("my-host")
				createHostOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.VolumeMappings = []sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				createHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateHost(createHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteHost(deleteHostOptions *DeleteHostOptions)`, func() {
		deleteHostPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteHostPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteHost(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteHostOptions model
				deleteHostOptionsModel := new(sdsaasv2.DeleteHostOptions)
				deleteHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteHost(deleteHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteHost with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteHostOptions model
				deleteHostOptionsModel := new(sdsaasv2.DeleteHostOptions)
				deleteHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteHost(deleteHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteHostOptions model with no property values
				deleteHostOptionsModelNew := new(sdsaasv2.DeleteHostOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteHost(deleteHostOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHost(getHostOptions *GetHostOptions) - Operation response error`, func() {
		getHostPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHost with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetHostOptions model
				getHostOptionsModel := new(sdsaasv2.GetHostOptions)
				getHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHost(getHostOptions *GetHostOptions)`, func() {
		getHostPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}`)
				}))
			})
			It(`Invoke GetHost successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the GetHostOptions model
				getHostOptionsModel := new(sdsaasv2.GetHostOptions)
				getHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.GetHostWithContext(ctx, getHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.GetHostWithContext(ctx, getHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}`)
				}))
			})
			It(`Invoke GetHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.GetHost(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHostOptions model
				getHostOptionsModel := new(sdsaasv2.GetHostOptions)
				getHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHost with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetHostOptions model
				getHostOptionsModel := new(sdsaasv2.GetHostOptions)
				getHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetHostOptions model with no property values
				getHostOptionsModelNew := new(sdsaasv2.GetHostOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.GetHost(getHostOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetHostOptions model
				getHostOptionsModel := new(sdsaasv2.GetHostOptions)
				getHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.GetHost(getHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateHost(updateHostOptions *UpdateHostOptions) - Operation response error`, func() {
		updateHostPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHostPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateHost with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv2.HostPatch)
				hostPatchModel.Name = core.StringPtr("my-resource")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateHostOptions model
				updateHostOptionsModel := new(sdsaasv2.UpdateHostOptions)
				updateHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.HostPatch = hostPatchModelAsPatch
				updateHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateHost(updateHostOptions *UpdateHostOptions)`, func() {
		updateHostPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHostPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}`)
				}))
			})
			It(`Invoke UpdateHost successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv2.HostPatch)
				hostPatchModel.Name = core.StringPtr("my-resource")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateHostOptions model
				updateHostOptionsModel := new(sdsaasv2.UpdateHostOptions)
				updateHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.HostPatch = hostPatchModelAsPatch
				updateHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.UpdateHostWithContext(ctx, updateHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.UpdateHostWithContext(ctx, updateHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHostPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}]}`)
				}))
			})
			It(`Invoke UpdateHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.UpdateHost(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv2.HostPatch)
				hostPatchModel.Name = core.StringPtr("my-resource")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateHostOptions model
				updateHostOptionsModel := new(sdsaasv2.UpdateHostOptions)
				updateHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.HostPatch = hostPatchModelAsPatch
				updateHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateHost with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv2.HostPatch)
				hostPatchModel.Name = core.StringPtr("my-resource")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateHostOptions model
				updateHostOptionsModel := new(sdsaasv2.UpdateHostOptions)
				updateHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.HostPatch = hostPatchModelAsPatch
				updateHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateHostOptions model with no property values
				updateHostOptionsModelNew := new(sdsaasv2.UpdateHostOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.UpdateHost(updateHostOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateHost successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv2.HostPatch)
				hostPatchModel.Name = core.StringPtr("my-resource")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateHostOptions model
				updateHostOptionsModel := new(sdsaasv2.UpdateHostOptions)
				updateHostOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.HostPatch = hostPatchModelAsPatch
				updateHostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.UpdateHost(updateHostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVolumeMappings(deleteVolumeMappingsOptions *DeleteVolumeMappingsOptions)`, func() {
		deleteVolumeMappingsPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVolumeMappingsPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVolumeMappings successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteVolumeMappings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVolumeMappingsOptions model
				deleteVolumeMappingsOptionsModel := new(sdsaasv2.DeleteVolumeMappingsOptions)
				deleteVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteVolumeMappings(deleteVolumeMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVolumeMappings with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteVolumeMappingsOptions model
				deleteVolumeMappingsOptionsModel := new(sdsaasv2.DeleteVolumeMappingsOptions)
				deleteVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteVolumeMappings(deleteVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVolumeMappingsOptions model with no property values
				deleteVolumeMappingsOptionsModelNew := new(sdsaasv2.DeleteVolumeMappingsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteVolumeMappings(deleteVolumeMappingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVolumeMappings(listVolumeMappingsOptions *ListVolumeMappingsOptions) - Operation response error`, func() {
		listVolumeMappingsPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumeMappingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVolumeMappings with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumeMappingsOptions model
				listVolumeMappingsOptionsModel := new(sdsaasv2.ListVolumeMappingsOptions)
				listVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumeMappingsOptionsModel.Name = core.StringPtr("my-resource")
				listVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVolumeMappings(listVolumeMappingsOptions *ListVolumeMappingsOptions)`, func() {
		listVolumeMappingsPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumeMappingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListVolumeMappings successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListVolumeMappingsOptions model
				listVolumeMappingsOptionsModel := new(sdsaasv2.ListVolumeMappingsOptions)
				listVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumeMappingsOptionsModel.Name = core.StringPtr("my-resource")
				listVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListVolumeMappingsWithContext(ctx, listVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListVolumeMappingsWithContext(ctx, listVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumeMappingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "volume_mappings": [{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListVolumeMappings successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListVolumeMappings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVolumeMappingsOptions model
				listVolumeMappingsOptionsModel := new(sdsaasv2.ListVolumeMappingsOptions)
				listVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumeMappingsOptionsModel.Name = core.StringPtr("my-resource")
				listVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVolumeMappings with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumeMappingsOptions model
				listVolumeMappingsOptionsModel := new(sdsaasv2.ListVolumeMappingsOptions)
				listVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumeMappingsOptionsModel.Name = core.StringPtr("my-resource")
				listVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListVolumeMappingsOptions model with no property values
				listVolumeMappingsOptionsModelNew := new(sdsaasv2.ListVolumeMappingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVolumeMappings successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListVolumeMappingsOptions model
				listVolumeMappingsOptionsModel := new(sdsaasv2.ListVolumeMappingsOptions)
				listVolumeMappingsOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listVolumeMappingsOptionsModel.Name = core.StringPtr("my-resource")
				listVolumeMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListVolumeMappings(listVolumeMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(sdsaasv2.VolumeMappingCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv2.VolumeMappingCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv2.VolumeMappingCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVolumeMappingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}],"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"volume_mappings":[{"status":"mapped","href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource"},"host":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","nqn":"nqn.2014-06.org:1234"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"192.168.3.4","port":22}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use VolumeMappingsPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listVolumeMappingsOptionsModel := &sdsaasv2.ListVolumeMappingsOptions{
					ID: core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewVolumeMappingsPager(listVolumeMappingsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv2.VolumeMapping
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use VolumeMappingsPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listVolumeMappingsOptionsModel := &sdsaasv2.ListVolumeMappingsOptions{
					ID: core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
				}

				pager, err := sdsaasService.NewVolumeMappingsPager(listVolumeMappingsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateVolumeMapping(createVolumeMappingOptions *CreateVolumeMappingOptions) - Operation response error`, func() {
		createVolumeMappingPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumeMappingPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVolumeMapping with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeMappingOptions model
				createVolumeMappingOptionsModel := new(sdsaasv2.CreateVolumeMappingOptions)
				createVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.Volume = volumeIdentityModel
				createVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVolumeMapping(createVolumeMappingOptions *CreateVolumeMappingOptions)`, func() {
		createVolumeMappingPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumeMappingPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"status": "pending", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}}`)
				}))
			})
			It(`Invoke CreateVolumeMapping successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeMappingOptions model
				createVolumeMappingOptionsModel := new(sdsaasv2.CreateVolumeMappingOptions)
				createVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.Volume = volumeIdentityModel
				createVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateVolumeMappingWithContext(ctx, createVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateVolumeMappingWithContext(ctx, createVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVolumeMappingPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"status": "pending", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}}`)
				}))
			})
			It(`Invoke CreateVolumeMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateVolumeMapping(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeMappingOptions model
				createVolumeMappingOptionsModel := new(sdsaasv2.CreateVolumeMappingOptions)
				createVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.Volume = volumeIdentityModel
				createVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVolumeMapping with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeMappingOptions model
				createVolumeMappingOptionsModel := new(sdsaasv2.CreateVolumeMappingOptions)
				createVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.Volume = volumeIdentityModel
				createVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVolumeMappingOptions model with no property values
				createVolumeMappingOptionsModelNew := new(sdsaasv2.CreateVolumeMappingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateVolumeMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeMappingOptions model
				createVolumeMappingOptionsModel := new(sdsaasv2.CreateVolumeMappingOptions)
				createVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.Volume = volumeIdentityModel
				createVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateVolumeMapping(createVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVolumeMapping(deleteVolumeMappingOptions *DeleteVolumeMappingOptions)`, func() {
		deleteVolumeMappingPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVolumeMappingPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVolumeMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteVolumeMapping(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVolumeMappingOptions model
				deleteVolumeMappingOptionsModel := new(sdsaasv2.DeleteVolumeMappingOptions)
				deleteVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteVolumeMapping(deleteVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVolumeMapping with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteVolumeMappingOptions model
				deleteVolumeMappingOptionsModel := new(sdsaasv2.DeleteVolumeMappingOptions)
				deleteVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteVolumeMapping(deleteVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVolumeMappingOptions model with no property values
				deleteVolumeMappingOptionsModelNew := new(sdsaasv2.DeleteVolumeMappingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteVolumeMapping(deleteVolumeMappingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVolumeMapping(getVolumeMappingOptions *GetVolumeMappingOptions) - Operation response error`, func() {
		getVolumeMappingPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumeMappingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVolumeMapping with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeMappingOptions model
				getVolumeMappingOptionsModel := new(sdsaasv2.GetVolumeMappingOptions)
				getVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVolumeMapping(getVolumeMappingOptions *GetVolumeMappingOptions)`, func() {
		getVolumeMappingPath := "/hosts/r134-b274-678d-4dfb-8981-c71dd9d4daa5/volume_mappings/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumeMappingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}`)
				}))
			})
			It(`Invoke GetVolumeMapping successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the GetVolumeMappingOptions model
				getVolumeMappingOptionsModel := new(sdsaasv2.GetVolumeMappingOptions)
				getVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.GetVolumeMappingWithContext(ctx, getVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.GetVolumeMappingWithContext(ctx, getVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVolumeMappingPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "mapped", "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource"}, "host": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "nqn": "nqn.2014-06.org:1234"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "192.168.3.4", "port": 22}]}`)
				}))
			})
			It(`Invoke GetVolumeMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.GetVolumeMapping(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVolumeMappingOptions model
				getVolumeMappingOptionsModel := new(sdsaasv2.GetVolumeMappingOptions)
				getVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVolumeMapping with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeMappingOptions model
				getVolumeMappingOptionsModel := new(sdsaasv2.GetVolumeMappingOptions)
				getVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVolumeMappingOptions model with no property values
				getVolumeMappingOptionsModelNew := new(sdsaasv2.GetVolumeMappingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVolumeMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetVolumeMappingOptions model
				getVolumeMappingOptionsModel := new(sdsaasv2.GetVolumeMappingOptions)
				getVolumeMappingOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.GetVolumeMapping(getVolumeMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListHmacCredentials(listHmacCredentialsOptions *ListHmacCredentialsOptions) - Operation response error`, func() {
		listHmacCredentialsPath := "/s3_credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHmacCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListHmacCredentials with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := new(sdsaasv2.ListHmacCredentialsOptions)
				listHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListHmacCredentials(listHmacCredentialsOptions *ListHmacCredentialsOptions)`, func() {
		listHmacCredentialsPath := "/s3_credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHmacCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"s3_credentials": ["S3Credentials"]}`)
				}))
			})
			It(`Invoke ListHmacCredentials successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := new(sdsaasv2.ListHmacCredentialsOptions)
				listHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListHmacCredentialsWithContext(ctx, listHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListHmacCredentialsWithContext(ctx, listHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listHmacCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"s3_credentials": ["S3Credentials"]}`)
				}))
			})
			It(`Invoke ListHmacCredentials successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListHmacCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := new(sdsaasv2.ListHmacCredentialsOptions)
				listHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListHmacCredentials with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := new(sdsaasv2.ListHmacCredentialsOptions)
				listHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListHmacCredentials successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := new(sdsaasv2.ListHmacCredentialsOptions)
				listHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListHmacCredentials(listHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteHmacCredentials(deleteHmacCredentialsOptions *DeleteHmacCredentialsOptions)`, func() {
		deleteHmacCredentialsPath := "/s3_credentials/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteHmacCredentialsPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteHmacCredentials successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteHmacCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteHmacCredentialsOptions model
				deleteHmacCredentialsOptionsModel := new(sdsaasv2.DeleteHmacCredentialsOptions)
				deleteHmacCredentialsOptionsModel.AccessKey = core.StringPtr("testString")
				deleteHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteHmacCredentials(deleteHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteHmacCredentials with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteHmacCredentialsOptions model
				deleteHmacCredentialsOptionsModel := new(sdsaasv2.DeleteHmacCredentialsOptions)
				deleteHmacCredentialsOptionsModel.AccessKey = core.StringPtr("testString")
				deleteHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteHmacCredentials(deleteHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteHmacCredentialsOptions model with no property values
				deleteHmacCredentialsOptionsModelNew := new(sdsaasv2.DeleteHmacCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteHmacCredentials(deleteHmacCredentialsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateHmacCredentials(createHmacCredentialsOptions *CreateHmacCredentialsOptions) - Operation response error`, func() {
		createHmacCredentialsPath := "/s3_credentials/xxx-xx-xxx"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHmacCredentialsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateHmacCredentials with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateHmacCredentialsOptions model
				createHmacCredentialsOptionsModel := new(sdsaasv2.CreateHmacCredentialsOptions)
				createHmacCredentialsOptionsModel.AccessKey = core.StringPtr("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateHmacCredentials(createHmacCredentialsOptions *CreateHmacCredentialsOptions)`, func() {
		createHmacCredentialsPath := "/s3_credentials/xxx-xx-xxx"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHmacCredentialsPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"access_key": "AccessKey", "secret_key": "SecretKey"}`)
				}))
			})
			It(`Invoke CreateHmacCredentials successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CreateHmacCredentialsOptions model
				createHmacCredentialsOptionsModel := new(sdsaasv2.CreateHmacCredentialsOptions)
				createHmacCredentialsOptionsModel.AccessKey = core.StringPtr("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateHmacCredentialsWithContext(ctx, createHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateHmacCredentialsWithContext(ctx, createHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createHmacCredentialsPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"access_key": "AccessKey", "secret_key": "SecretKey"}`)
				}))
			})
			It(`Invoke CreateHmacCredentials successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateHmacCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateHmacCredentialsOptions model
				createHmacCredentialsOptionsModel := new(sdsaasv2.CreateHmacCredentialsOptions)
				createHmacCredentialsOptionsModel.AccessKey = core.StringPtr("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateHmacCredentials with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateHmacCredentialsOptions model
				createHmacCredentialsOptionsModel := new(sdsaasv2.CreateHmacCredentialsOptions)
				createHmacCredentialsOptionsModel.AccessKey = core.StringPtr("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateHmacCredentialsOptions model with no property values
				createHmacCredentialsOptionsModelNew := new(sdsaasv2.CreateHmacCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateHmacCredentials successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateHmacCredentialsOptions model
				createHmacCredentialsOptionsModel := new(sdsaasv2.CreateHmacCredentialsOptions)
				createHmacCredentialsOptionsModel.AccessKey = core.StringPtr("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCertificates(listCertificatesOptions *ListCertificatesOptions) - Operation response error`, func() {
		listCertificatesPath := "/certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCertificates with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sdsaasv2.ListCertificatesOptions)
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCertificates(listCertificatesOptions *ListCertificatesOptions)`, func() {
		listCertificatesPath := "/certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificates": ["Certificates"]}`)
				}))
			})
			It(`Invoke ListCertificates successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sdsaasv2.ListCertificatesOptions)
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListCertificatesWithContext(ctx, listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListCertificatesWithContext(ctx, listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificates": ["Certificates"]}`)
				}))
			})
			It(`Invoke ListCertificates successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sdsaasv2.ListCertificatesOptions)
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCertificates with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sdsaasv2.ListCertificatesOptions)
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCertificates successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sdsaasv2.ListCertificatesOptions)
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSslCert(deleteSslCertOptions *DeleteSslCertOptions)`, func() {
		deleteSslCertPath := "/certificates/s3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSslCertPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSslCert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteSslCert(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSslCertOptions model
				deleteSslCertOptionsModel := new(sdsaasv2.DeleteSslCertOptions)
				deleteSslCertOptionsModel.CertType = core.StringPtr("s3")
				deleteSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteSslCert(deleteSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSslCert with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteSslCertOptions model
				deleteSslCertOptionsModel := new(sdsaasv2.DeleteSslCertOptions)
				deleteSslCertOptionsModel.CertType = core.StringPtr("s3")
				deleteSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteSslCert(deleteSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSslCertOptions model with no property values
				deleteSslCertOptionsModelNew := new(sdsaasv2.DeleteSslCertOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteSslCert(deleteSslCertOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetS3SslCertStatus(getS3SslCertStatusOptions *GetS3SslCertStatusOptions) - Operation response error`, func() {
		getS3SslCertStatusPath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getS3SslCertStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetS3SslCertStatus with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetS3SslCertStatusOptions model
				getS3SslCertStatusOptionsModel := new(sdsaasv2.GetS3SslCertStatusOptions)
				getS3SslCertStatusOptionsModel.CertType = core.StringPtr("s3")
				getS3SslCertStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetS3SslCertStatus(getS3SslCertStatusOptions *GetS3SslCertStatusOptions)`, func() {
		getS3SslCertStatusPath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getS3SslCertStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"expiration_date": "2019-01-01T12:00:00.000Z", "expired": false, "name": "Name"}`)
				}))
			})
			It(`Invoke GetS3SslCertStatus successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the GetS3SslCertStatusOptions model
				getS3SslCertStatusOptionsModel := new(sdsaasv2.GetS3SslCertStatusOptions)
				getS3SslCertStatusOptionsModel.CertType = core.StringPtr("s3")
				getS3SslCertStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.GetS3SslCertStatusWithContext(ctx, getS3SslCertStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.GetS3SslCertStatusWithContext(ctx, getS3SslCertStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getS3SslCertStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"expiration_date": "2019-01-01T12:00:00.000Z", "expired": false, "name": "Name"}`)
				}))
			})
			It(`Invoke GetS3SslCertStatus successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.GetS3SslCertStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetS3SslCertStatusOptions model
				getS3SslCertStatusOptionsModel := new(sdsaasv2.GetS3SslCertStatusOptions)
				getS3SslCertStatusOptionsModel.CertType = core.StringPtr("s3")
				getS3SslCertStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetS3SslCertStatus with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetS3SslCertStatusOptions model
				getS3SslCertStatusOptionsModel := new(sdsaasv2.GetS3SslCertStatusOptions)
				getS3SslCertStatusOptionsModel.CertType = core.StringPtr("s3")
				getS3SslCertStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetS3SslCertStatusOptions model with no property values
				getS3SslCertStatusOptionsModelNew := new(sdsaasv2.GetS3SslCertStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetS3SslCertStatus successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetS3SslCertStatusOptions model
				getS3SslCertStatusOptionsModel := new(sdsaasv2.GetS3SslCertStatusOptions)
				getS3SslCertStatusOptionsModel.CertType = core.StringPtr("s3")
				getS3SslCertStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSslCert(createSslCertOptions *CreateSslCertOptions) - Operation response error`, func() {
		createSslCertPath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSslCertPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSslCert with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateSslCertOptions model
				createSslCertOptionsModel := new(sdsaasv2.CreateSslCertOptions)
				createSslCertOptionsModel.CertType = core.StringPtr("s3")
				createSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				createSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSslCert(createSslCertOptions *CreateSslCertOptions)`, func() {
		createSslCertPath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSslCertPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"errors": [{"code": "invalid_response", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "trace": "Trace", "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CreateSslCert successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CreateSslCertOptions model
				createSslCertOptionsModel := new(sdsaasv2.CreateSslCertOptions)
				createSslCertOptionsModel.CertType = core.StringPtr("s3")
				createSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				createSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateSslCertWithContext(ctx, createSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateSslCertWithContext(ctx, createSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSslCertPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"errors": [{"code": "invalid_response", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "trace": "Trace", "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CreateSslCert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateSslCert(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSslCertOptions model
				createSslCertOptionsModel := new(sdsaasv2.CreateSslCertOptions)
				createSslCertOptionsModel.CertType = core.StringPtr("s3")
				createSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				createSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSslCert with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateSslCertOptions model
				createSslCertOptionsModel := new(sdsaasv2.CreateSslCertOptions)
				createSslCertOptionsModel.CertType = core.StringPtr("s3")
				createSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				createSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSslCertOptions model with no property values
				createSslCertOptionsModelNew := new(sdsaasv2.CreateSslCertOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CreateSslCert(createSslCertOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateSslCert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CreateSslCertOptions model
				createSslCertOptionsModel := new(sdsaasv2.CreateSslCertOptions)
				createSslCertOptionsModel.CertType = core.StringPtr("s3")
				createSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				createSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateSslCert(createSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSslCert(replaceSslCertOptions *ReplaceSslCertOptions) - Operation response error`, func() {
		replaceSslCertPath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSslCertPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceSslCert with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ReplaceSslCertOptions model
				replaceSslCertOptionsModel := new(sdsaasv2.ReplaceSslCertOptions)
				replaceSslCertOptionsModel.CertType = core.StringPtr("s3")
				replaceSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				replaceSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSslCert(replaceSslCertOptions *ReplaceSslCertOptions)`, func() {
		replaceSslCertPath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSslCertPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"errors": [{"code": "invalid_response", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "trace": "Trace", "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke ReplaceSslCert successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceSslCertOptions model
				replaceSslCertOptionsModel := new(sdsaasv2.ReplaceSslCertOptions)
				replaceSslCertOptionsModel.CertType = core.StringPtr("s3")
				replaceSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				replaceSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ReplaceSslCertWithContext(ctx, replaceSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ReplaceSslCertWithContext(ctx, replaceSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSslCertPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"errors": [{"code": "invalid_response", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "trace": "Trace", "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke ReplaceSslCert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ReplaceSslCert(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceSslCertOptions model
				replaceSslCertOptionsModel := new(sdsaasv2.ReplaceSslCertOptions)
				replaceSslCertOptionsModel.CertType = core.StringPtr("s3")
				replaceSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				replaceSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceSslCert with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ReplaceSslCertOptions model
				replaceSslCertOptionsModel := new(sdsaasv2.ReplaceSslCertOptions)
				replaceSslCertOptionsModel.CertType = core.StringPtr("s3")
				replaceSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				replaceSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceSslCertOptions model with no property values
				replaceSslCertOptionsModelNew := new(sdsaasv2.ReplaceSslCertOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.ReplaceSslCert(replaceSslCertOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ReplaceSslCert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ReplaceSslCertOptions model
				replaceSslCertOptionsModel := new(sdsaasv2.ReplaceSslCertOptions)
				replaceSslCertOptionsModel.CertType = core.StringPtr("s3")
				replaceSslCertOptionsModel.Body = CreateMockReader("This is a mock file.")
				replaceSslCertOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ReplaceSslCert(replaceSslCertOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSnapshots(deleteSnapshotsOptions *DeleteSnapshotsOptions)`, func() {
		deleteSnapshotsPath := "/snapshots"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSnapshotsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSnapshots successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSnapshotsOptions model
				deleteSnapshotsOptionsModel := new(sdsaasv2.DeleteSnapshotsOptions)
				deleteSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteSnapshots(deleteSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSnapshots with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteSnapshotsOptions model
				deleteSnapshotsOptionsModel := new(sdsaasv2.DeleteSnapshotsOptions)
				deleteSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteSnapshots(deleteSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSnapshotsOptions model with no property values
				deleteSnapshotsOptionsModelNew := new(sdsaasv2.DeleteSnapshotsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteSnapshots(deleteSnapshotsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) - Operation response error`, func() {
		listSnapshotsPath := "/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSnapshots with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(sdsaasv2.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listSnapshotsOptionsModel.Name = core.StringPtr("my-resource")
				listSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions)`, func() {
		listSnapshotsPath := "/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"snapshots": [{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListSnapshots successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(sdsaasv2.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listSnapshotsOptionsModel.Name = core.StringPtr("my-resource")
				listSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.ListSnapshotsWithContext(ctx, listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.ListSnapshotsWithContext(ctx, listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-resource"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"r134-b274-678d-4dfb-8981-c71dd9d4daa5"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"snapshots": [{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke ListSnapshots successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.ListSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(sdsaasv2.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listSnapshotsOptionsModel.Name = core.StringPtr("my-resource")
				listSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSnapshots with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(sdsaasv2.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listSnapshotsOptionsModel.Name = core.StringPtr("my-resource")
				listSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSnapshots successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(sdsaasv2.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Start = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listSnapshotsOptionsModel.Name = core.StringPtr("my-resource")
				listSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(sdsaasv2.SnapshotCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv2.SnapshotCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv2.SnapshotCollection)
				nextObject := new(sdsaasv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"snapshots":[{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","created_at":"2019-01-01T12:00:00.000Z","resource_type":"ResourceType","lifecycle_state":"stable","size":30,"minimum_capacity":30,"deletable":true,"source_volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","resource_type":"ResourceType"}}],"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"snapshots":[{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","href":"Href","name":"my-resource","created_at":"2019-01-01T12:00:00.000Z","resource_type":"ResourceType","lifecycle_state":"stable","size":30,"minimum_capacity":30,"deletable":true,"source_volume":{"id":"r134-b274-678d-4dfb-8981-c71dd9d4daa5","name":"my-resource","resource_type":"ResourceType"}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SnapshotsPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listSnapshotsOptionsModel := &sdsaasv2.ListSnapshotsOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
					SourceVolumeID: core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
				}

				pager, err := sdsaasService.NewSnapshotsPager(listSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv2.Snapshot
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SnapshotsPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				listSnapshotsOptionsModel := &sdsaasv2.ListSnapshotsOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name: core.StringPtr("my-resource"),
					SourceVolumeID: core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
				}

				pager, err := sdsaasService.NewSnapshotsPager(listSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateSnapshot(createSnapshotOptions *CreateSnapshotOptions) - Operation response error`, func() {
		createSnapshotPath := "/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSnapshotPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSnapshot with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := new(sdsaasv2.CreateSnapshotOptions)
				createSnapshotOptionsModel.Name = core.StringPtr("my-snapshot")
				createSnapshotOptionsModel.SourceVolume = sourceVolumePrototypeModel
				createSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSnapshot(createSnapshotOptions *CreateSnapshotOptions)`, func() {
		createSnapshotPath := "/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSnapshotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke CreateSnapshot successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := new(sdsaasv2.CreateSnapshotOptions)
				createSnapshotOptionsModel.Name = core.StringPtr("my-snapshot")
				createSnapshotOptionsModel.SourceVolume = sourceVolumePrototypeModel
				createSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CreateSnapshotWithContext(ctx, createSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CreateSnapshotWithContext(ctx, createSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSnapshotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke CreateSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CreateSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := new(sdsaasv2.CreateSnapshotOptions)
				createSnapshotOptionsModel.Name = core.StringPtr("my-snapshot")
				createSnapshotOptionsModel.SourceVolume = sourceVolumePrototypeModel
				createSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSnapshot with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := new(sdsaasv2.CreateSnapshotOptions)
				createSnapshotOptionsModel.Name = core.StringPtr("my-snapshot")
				createSnapshotOptionsModel.SourceVolume = sourceVolumePrototypeModel
				createSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := new(sdsaasv2.CreateSnapshotOptions)
				createSnapshotOptionsModel.Name = core.StringPtr("my-snapshot")
				createSnapshotOptionsModel.SourceVolume = sourceVolumePrototypeModel
				createSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CreateSnapshot(createSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSnapshot(deleteSnapshotOptions *DeleteSnapshotOptions)`, func() {
		deleteSnapshotPath := "/snapshots/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSnapshotPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.DeleteSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSnapshotOptions model
				deleteSnapshotOptionsModel := new(sdsaasv2.DeleteSnapshotOptions)
				deleteSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.DeleteSnapshot(deleteSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSnapshot with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the DeleteSnapshotOptions model
				deleteSnapshotOptionsModel := new(sdsaasv2.DeleteSnapshotOptions)
				deleteSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.DeleteSnapshot(deleteSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSnapshotOptions model with no property values
				deleteSnapshotOptionsModelNew := new(sdsaasv2.DeleteSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.DeleteSnapshot(deleteSnapshotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSnapshot(getSnapshotOptions *GetSnapshotOptions) - Operation response error`, func() {
		getSnapshotPath := "/snapshots/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSnapshotPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSnapshot with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetSnapshotOptions model
				getSnapshotOptionsModel := new(sdsaasv2.GetSnapshotOptions)
				getSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSnapshot(getSnapshotOptions *GetSnapshotOptions)`, func() {
		getSnapshotPath := "/snapshots/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke GetSnapshot successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the GetSnapshotOptions model
				getSnapshotOptionsModel := new(sdsaasv2.GetSnapshotOptions)
				getSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.GetSnapshotWithContext(ctx, getSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.GetSnapshotWithContext(ctx, getSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke GetSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.GetSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSnapshotOptions model
				getSnapshotOptionsModel := new(sdsaasv2.GetSnapshotOptions)
				getSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSnapshot with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetSnapshotOptions model
				getSnapshotOptionsModel := new(sdsaasv2.GetSnapshotOptions)
				getSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSnapshotOptions model with no property values
				getSnapshotOptionsModelNew := new(sdsaasv2.GetSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.GetSnapshot(getSnapshotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the GetSnapshotOptions model
				getSnapshotOptionsModel := new(sdsaasv2.GetSnapshotOptions)
				getSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.GetSnapshot(getSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSnapshot(updateSnapshotOptions *UpdateSnapshotOptions) - Operation response error`, func() {
		updateSnapshotPath := "/snapshots/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSnapshotPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSnapshot with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv2.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("my-snapshot")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateSnapshotOptions model
				updateSnapshotOptionsModel := new(sdsaasv2.UpdateSnapshotOptions)
				updateSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				updateSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSnapshot(updateSnapshotOptions *UpdateSnapshotOptions)`, func() {
		updateSnapshotPath := "/snapshots/r134-b274-678d-4dfb-8981-c71dd9d4daa5"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSnapshotPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke UpdateSnapshot successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv2.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("my-snapshot")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateSnapshotOptions model
				updateSnapshotOptionsModel := new(sdsaasv2.UpdateSnapshotOptions)
				updateSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				updateSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.UpdateSnapshotWithContext(ctx, updateSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.UpdateSnapshotWithContext(ctx, updateSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSnapshotPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "href": "Href", "name": "my-resource", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "ResourceType", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "r134-b274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-resource", "resource_type": "ResourceType"}}`)
				}))
			})
			It(`Invoke UpdateSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.UpdateSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv2.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("my-snapshot")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateSnapshotOptions model
				updateSnapshotOptionsModel := new(sdsaasv2.UpdateSnapshotOptions)
				updateSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				updateSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSnapshot with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv2.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("my-snapshot")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateSnapshotOptions model
				updateSnapshotOptionsModel := new(sdsaasv2.UpdateSnapshotOptions)
				updateSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				updateSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSnapshotOptions model with no property values
				updateSnapshotOptionsModelNew := new(sdsaasv2.UpdateSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.UpdateSnapshot(updateSnapshotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv2.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("my-snapshot")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateSnapshotOptions model
				updateSnapshotOptionsModel := new(sdsaasv2.UpdateSnapshotOptions)
				updateSnapshotOptionsModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				updateSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.UpdateSnapshot(updateSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			sdsaasService, _ := sdsaasv2.NewSdsaasV2(&sdsaasv2.SdsaasV2Options{
				URL:           "http://sdsaasv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateHmacCredentialsOptions successfully`, func() {
				// Construct an instance of the CreateHmacCredentialsOptions model
				accessKey := "xxx-xx-xxx"
				createHmacCredentialsOptionsModel := sdsaasService.NewCreateHmacCredentialsOptions(accessKey)
				createHmacCredentialsOptionsModel.SetAccessKey("xxx-xx-xxx")
				createHmacCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createHmacCredentialsOptionsModel).ToNot(BeNil())
				Expect(createHmacCredentialsOptionsModel.AccessKey).To(Equal(core.StringPtr("xxx-xx-xxx")))
				Expect(createHmacCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateHostOptions successfully`, func() {
				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				Expect(volumeIdentityModel).ToNot(BeNil())
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(volumeIdentityModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv2.VolumeMappingPrototype)
				Expect(volumeMappingPrototypeModel).ToNot(BeNil())
				volumeMappingPrototypeModel.Volume = volumeIdentityModel
				Expect(volumeMappingPrototypeModel.Volume).To(Equal(volumeIdentityModel))

				// Construct an instance of the CreateHostOptions model
				createHostOptionsNqn := "nqn.2014-06.org:1234"
				createHostOptionsModel := sdsaasService.NewCreateHostOptions(createHostOptionsNqn)
				createHostOptionsModel.SetNqn("nqn.2014-06.org:1234")
				createHostOptionsModel.SetName("my-host")
				createHostOptionsModel.SetPsk("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				createHostOptionsModel.SetVolumeMappings([]sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel})
				createHostOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createHostOptionsModel).ToNot(BeNil())
				Expect(createHostOptionsModel.Nqn).To(Equal(core.StringPtr("nqn.2014-06.org:1234")))
				Expect(createHostOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(createHostOptionsModel.Psk).To(Equal(core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")))
				Expect(createHostOptionsModel.VolumeMappings).To(Equal([]sdsaasv2.VolumeMappingPrototype{*volumeMappingPrototypeModel}))
				Expect(createHostOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSnapshotOptions successfully`, func() {
				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv2.SourceVolumePrototype)
				Expect(sourceVolumePrototypeModel).ToNot(BeNil())
				sourceVolumePrototypeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(sourceVolumePrototypeModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the CreateSnapshotOptions model
				createSnapshotOptionsModel := sdsaasService.NewCreateSnapshotOptions()
				createSnapshotOptionsModel.SetName("my-snapshot")
				createSnapshotOptionsModel.SetSourceVolume(sourceVolumePrototypeModel)
				createSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSnapshotOptionsModel).ToNot(BeNil())
				Expect(createSnapshotOptionsModel.Name).To(Equal(core.StringPtr("my-snapshot")))
				Expect(createSnapshotOptionsModel.SourceVolume).To(Equal(sourceVolumePrototypeModel))
				Expect(createSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSslCertOptions successfully`, func() {
				// Construct an instance of the CreateSslCertOptions model
				certType := "s3"
				createSslCertOptionsModel := sdsaasService.NewCreateSslCertOptions(certType)
				createSslCertOptionsModel.SetCertType("s3")
				createSslCertOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				createSslCertOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSslCertOptionsModel).ToNot(BeNil())
				Expect(createSslCertOptionsModel.CertType).To(Equal(core.StringPtr("s3")))
				Expect(createSslCertOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createSslCertOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVolumeMappingOptions successfully`, func() {
				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv2.VolumeIdentity)
				Expect(volumeIdentityModel).ToNot(BeNil())
				volumeIdentityModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(volumeIdentityModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the CreateVolumeMappingOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				var createVolumeMappingOptionsVolume *sdsaasv2.VolumeIdentity = nil
				createVolumeMappingOptionsModel := sdsaasService.NewCreateVolumeMappingOptions(id, createVolumeMappingOptionsVolume)
				createVolumeMappingOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				createVolumeMappingOptionsModel.SetVolume(volumeIdentityModel)
				createVolumeMappingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVolumeMappingOptionsModel).ToNot(BeNil())
				Expect(createVolumeMappingOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(createVolumeMappingOptionsModel.Volume).To(Equal(volumeIdentityModel))
				Expect(createVolumeMappingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVolumeOptions successfully`, func() {
				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv2.SourceSnapshot)
				Expect(sourceSnapshotModel).ToNot(BeNil())
				sourceSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(sourceSnapshotModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the SourceVolumeGroupSnapshotVolume model
				sourceVolumeGroupSnapshotVolumeModel := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
				Expect(sourceVolumeGroupSnapshotVolumeModel).ToNot(BeNil())
				sourceVolumeGroupSnapshotVolumeModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(sourceVolumeGroupSnapshotVolumeModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the SourceVolumeGroupSnapshot model
				sourceVolumeGroupSnapshotModel := new(sdsaasv2.SourceVolumeGroupSnapshot)
				Expect(sourceVolumeGroupSnapshotModel).ToNot(BeNil())
				sourceVolumeGroupSnapshotModel.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				sourceVolumeGroupSnapshotModel.Volume = sourceVolumeGroupSnapshotVolumeModel
				Expect(sourceVolumeGroupSnapshotModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(sourceVolumeGroupSnapshotModel.Volume).To(Equal(sourceVolumeGroupSnapshotVolumeModel))

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsCapacity := int64(1)
				createVolumeOptionsModel := sdsaasService.NewCreateVolumeOptions(createVolumeOptionsCapacity)
				createVolumeOptionsModel.SetCapacity(int64(1))
				createVolumeOptionsModel.SetName("my-volume")
				createVolumeOptionsModel.SetSourceSnapshot(sourceSnapshotModel)
				createVolumeOptionsModel.SetSourceVolumeGroupSnapshot(sourceVolumeGroupSnapshotModel)
				createVolumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVolumeOptionsModel).ToNot(BeNil())
				Expect(createVolumeOptionsModel.Capacity).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createVolumeOptionsModel.Name).To(Equal(core.StringPtr("my-volume")))
				Expect(createVolumeOptionsModel.SourceSnapshot).To(Equal(sourceSnapshotModel))
				Expect(createVolumeOptionsModel.SourceVolumeGroupSnapshot).To(Equal(sourceVolumeGroupSnapshotModel))
				Expect(createVolumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteHmacCredentialsOptions successfully`, func() {
				// Construct an instance of the DeleteHmacCredentialsOptions model
				accessKey := "testString"
				deleteHmacCredentialsOptionsModel := sdsaasService.NewDeleteHmacCredentialsOptions(accessKey)
				deleteHmacCredentialsOptionsModel.SetAccessKey("testString")
				deleteHmacCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteHmacCredentialsOptionsModel).ToNot(BeNil())
				Expect(deleteHmacCredentialsOptionsModel.AccessKey).To(Equal(core.StringPtr("testString")))
				Expect(deleteHmacCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteHostOptions successfully`, func() {
				// Construct an instance of the DeleteHostOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteHostOptionsModel := sdsaasService.NewDeleteHostOptions(id)
				deleteHostOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteHostOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteHostOptionsModel).ToNot(BeNil())
				Expect(deleteHostOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteHostOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSnapshotOptions successfully`, func() {
				// Construct an instance of the DeleteSnapshotOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteSnapshotOptionsModel := sdsaasService.NewDeleteSnapshotOptions(id)
				deleteSnapshotOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSnapshotOptionsModel).ToNot(BeNil())
				Expect(deleteSnapshotOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSnapshotsOptions successfully`, func() {
				// Construct an instance of the DeleteSnapshotsOptions model
				sourceVolumeID := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteSnapshotsOptionsModel := sdsaasService.NewDeleteSnapshotsOptions(sourceVolumeID)
				deleteSnapshotsOptionsModel.SetSourceVolumeID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSnapshotsOptionsModel).ToNot(BeNil())
				Expect(deleteSnapshotsOptionsModel.SourceVolumeID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSslCertOptions successfully`, func() {
				// Construct an instance of the DeleteSslCertOptions model
				certType := "s3"
				deleteSslCertOptionsModel := sdsaasService.NewDeleteSslCertOptions(certType)
				deleteSslCertOptionsModel.SetCertType("s3")
				deleteSslCertOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSslCertOptionsModel).ToNot(BeNil())
				Expect(deleteSslCertOptionsModel.CertType).To(Equal(core.StringPtr("s3")))
				Expect(deleteSslCertOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVolumeMappingOptions successfully`, func() {
				// Construct an instance of the DeleteVolumeMappingOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				volumeMappingID := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteVolumeMappingOptionsModel := sdsaasService.NewDeleteVolumeMappingOptions(id, volumeMappingID)
				deleteVolumeMappingOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.SetVolumeMappingID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVolumeMappingOptionsModel).ToNot(BeNil())
				Expect(deleteVolumeMappingOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteVolumeMappingOptionsModel.VolumeMappingID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteVolumeMappingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVolumeMappingsOptions successfully`, func() {
				// Construct an instance of the DeleteVolumeMappingsOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteVolumeMappingsOptionsModel := sdsaasService.NewDeleteVolumeMappingsOptions(id)
				deleteVolumeMappingsOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeMappingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVolumeMappingsOptionsModel).ToNot(BeNil())
				Expect(deleteVolumeMappingsOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteVolumeMappingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVolumeOptions successfully`, func() {
				// Construct an instance of the DeleteVolumeOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				deleteVolumeOptionsModel := sdsaasService.NewDeleteVolumeOptions(id)
				deleteVolumeOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				deleteVolumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVolumeOptionsModel).ToNot(BeNil())
				Expect(deleteVolumeOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(deleteVolumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHostOptions successfully`, func() {
				// Construct an instance of the GetHostOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				getHostOptionsModel := sdsaasService.NewGetHostOptions(id)
				getHostOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getHostOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHostOptionsModel).ToNot(BeNil())
				Expect(getHostOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(getHostOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetS3SslCertStatusOptions successfully`, func() {
				// Construct an instance of the GetS3SslCertStatusOptions model
				certType := "s3"
				getS3SslCertStatusOptionsModel := sdsaasService.NewGetS3SslCertStatusOptions(certType)
				getS3SslCertStatusOptionsModel.SetCertType("s3")
				getS3SslCertStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getS3SslCertStatusOptionsModel).ToNot(BeNil())
				Expect(getS3SslCertStatusOptionsModel.CertType).To(Equal(core.StringPtr("s3")))
				Expect(getS3SslCertStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSnapshotOptions successfully`, func() {
				// Construct an instance of the GetSnapshotOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				getSnapshotOptionsModel := sdsaasService.NewGetSnapshotOptions(id)
				getSnapshotOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSnapshotOptionsModel).ToNot(BeNil())
				Expect(getSnapshotOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(getSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVolumeMappingOptions successfully`, func() {
				// Construct an instance of the GetVolumeMappingOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				volumeMappingID := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				getVolumeMappingOptionsModel := sdsaasService.NewGetVolumeMappingOptions(id, volumeMappingID)
				getVolumeMappingOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.SetVolumeMappingID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeMappingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVolumeMappingOptionsModel).ToNot(BeNil())
				Expect(getVolumeMappingOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(getVolumeMappingOptionsModel.VolumeMappingID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(getVolumeMappingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVolumeOptions successfully`, func() {
				// Construct an instance of the GetVolumeOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				getVolumeOptionsModel := sdsaasService.NewGetVolumeOptions(id)
				getVolumeOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				getVolumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVolumeOptionsModel).ToNot(BeNil())
				Expect(getVolumeOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(getVolumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCertificatesOptions successfully`, func() {
				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := sdsaasService.NewListCertificatesOptions()
				listCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCertificatesOptionsModel).ToNot(BeNil())
				Expect(listCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListHmacCredentialsOptions successfully`, func() {
				// Construct an instance of the ListHmacCredentialsOptions model
				listHmacCredentialsOptionsModel := sdsaasService.NewListHmacCredentialsOptions()
				listHmacCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listHmacCredentialsOptionsModel).ToNot(BeNil())
				Expect(listHmacCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListHostsOptions successfully`, func() {
				// Construct an instance of the ListHostsOptions model
				listHostsOptionsModel := sdsaasService.NewListHostsOptions()
				listHostsOptionsModel.SetStart("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listHostsOptionsModel.SetLimit(int64(20))
				listHostsOptionsModel.SetName("my-resource")
				listHostsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listHostsOptionsModel).ToNot(BeNil())
				Expect(listHostsOptionsModel.Start).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listHostsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listHostsOptionsModel.Name).To(Equal(core.StringPtr("my-resource")))
				Expect(listHostsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSnapshotsOptions successfully`, func() {
				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := sdsaasService.NewListSnapshotsOptions()
				listSnapshotsOptionsModel.SetStart("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.SetLimit(int64(20))
				listSnapshotsOptionsModel.SetName("my-resource")
				listSnapshotsOptionsModel.SetSourceVolumeID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSnapshotsOptionsModel).ToNot(BeNil())
				Expect(listSnapshotsOptionsModel.Start).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listSnapshotsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listSnapshotsOptionsModel.Name).To(Equal(core.StringPtr("my-resource")))
				Expect(listSnapshotsOptionsModel.SourceVolumeID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVolumeMappingsOptions successfully`, func() {
				// Construct an instance of the ListVolumeMappingsOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				listVolumeMappingsOptionsModel := sdsaasService.NewListVolumeMappingsOptions(id)
				listVolumeMappingsOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.SetStart("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumeMappingsOptionsModel.SetLimit(int64(20))
				listVolumeMappingsOptionsModel.SetName("my-resource")
				listVolumeMappingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVolumeMappingsOptionsModel).ToNot(BeNil())
				Expect(listVolumeMappingsOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listVolumeMappingsOptionsModel.Start).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listVolumeMappingsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listVolumeMappingsOptionsModel.Name).To(Equal(core.StringPtr("my-resource")))
				Expect(listVolumeMappingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVolumesOptions successfully`, func() {
				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := sdsaasService.NewListVolumesOptions()
				listVolumesOptionsModel.SetStart("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				listVolumesOptionsModel.SetLimit(int64(20))
				listVolumesOptionsModel.SetName("my-resource")
				listVolumesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVolumesOptionsModel).ToNot(BeNil())
				Expect(listVolumesOptionsModel.Start).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(listVolumesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listVolumesOptionsModel.Name).To(Equal(core.StringPtr("my-resource")))
				Expect(listVolumesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceSslCertOptions successfully`, func() {
				// Construct an instance of the ReplaceSslCertOptions model
				certType := "s3"
				replaceSslCertOptionsModel := sdsaasService.NewReplaceSslCertOptions(certType)
				replaceSslCertOptionsModel.SetCertType("s3")
				replaceSslCertOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				replaceSslCertOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceSslCertOptionsModel).ToNot(BeNil())
				Expect(replaceSslCertOptionsModel.CertType).To(Equal(core.StringPtr("s3")))
				Expect(replaceSslCertOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(replaceSslCertOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSourceSnapshot successfully`, func() {
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				_model, err := sdsaasService.NewSourceSnapshot(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceVolumeGroupSnapshot successfully`, func() {
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				var volume *sdsaasv2.SourceVolumeGroupSnapshotVolume = nil
				_, err := sdsaasService.NewSourceVolumeGroupSnapshot(id, volume)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewSourceVolumeGroupSnapshotVolume successfully`, func() {
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				_model, err := sdsaasService.NewSourceVolumeGroupSnapshotVolume(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceVolumePrototype successfully`, func() {
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				_model, err := sdsaasService.NewSourceVolumePrototype(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateHostOptions successfully`, func() {
				// Construct an instance of the UpdateHostOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				hostPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateHostOptionsModel := sdsaasService.NewUpdateHostOptions(id, hostPatch)
				updateHostOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateHostOptionsModel.SetHostPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateHostOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateHostOptionsModel).ToNot(BeNil())
				Expect(updateHostOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(updateHostOptionsModel.HostPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateHostOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSnapshotOptions successfully`, func() {
				// Construct an instance of the UpdateSnapshotOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				snapshotPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateSnapshotOptionsModel := sdsaasService.NewUpdateSnapshotOptions(id, snapshotPatch)
				updateSnapshotOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateSnapshotOptionsModel.SetSnapshotPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSnapshotOptionsModel).ToNot(BeNil())
				Expect(updateSnapshotOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(updateSnapshotOptionsModel.SnapshotPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVolumeOptions successfully`, func() {
				// Construct an instance of the UpdateVolumeOptions model
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				volumePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateVolumeOptionsModel := sdsaasService.NewUpdateVolumeOptions(id, volumePatch)
				updateVolumeOptionsModel.SetID("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
				updateVolumeOptionsModel.SetVolumePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateVolumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVolumeOptionsModel).ToNot(BeNil())
				Expect(updateVolumeOptionsModel.ID).To(Equal(core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")))
				Expect(updateVolumeOptionsModel.VolumePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateVolumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeIdentity successfully`, func() {
				id := "r134-b274-678d-4dfb-8981-c71dd9d4daa5"
				_model, err := sdsaasService.NewVolumeIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeMappingPrototype successfully`, func() {
				var volume *sdsaasv2.VolumeIdentity = nil
				_, err := sdsaasService.NewVolumeMappingPrototype(volume)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalHostPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.HostPatch)
			model.Name = core.StringPtr("my-resource")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.HostPatch
			err = sdsaasv2.UnmarshalHostPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSnapshotPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.SnapshotPatch)
			model.Name = core.StringPtr("my-snapshot")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.SnapshotPatch
			err = sdsaasv2.UnmarshalSnapshotPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceSnapshot successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.SourceSnapshot)
			model.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.SourceSnapshot
			err = sdsaasv2.UnmarshalSourceSnapshot(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceVolumeGroupSnapshot successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.SourceVolumeGroupSnapshot)
			model.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")
			model.Volume = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.SourceVolumeGroupSnapshot
			err = sdsaasv2.UnmarshalSourceVolumeGroupSnapshot(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceVolumeGroupSnapshotVolume successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.SourceVolumeGroupSnapshotVolume)
			model.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.SourceVolumeGroupSnapshotVolume
			err = sdsaasv2.UnmarshalSourceVolumeGroupSnapshotVolume(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceVolumePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.SourceVolumePrototype)
			model.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.SourceVolumePrototype
			err = sdsaasv2.UnmarshalSourceVolumePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumeIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.VolumeIdentity)
			model.ID = core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.VolumeIdentity
			err = sdsaasv2.UnmarshalVolumeIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumeMappingPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.VolumeMappingPrototype)
			model.Volume = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.VolumeMappingPrototype
			err = sdsaasv2.UnmarshalVolumeMappingPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv2.VolumePatch)
			model.Capacity = core.Int64Ptr(int64(100))
			model.Name = core.StringPtr("my-volume")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv2.VolumePatch
			err = sdsaasv2.UnmarshalVolumePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
