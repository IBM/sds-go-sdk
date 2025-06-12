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

package sdsaasv1_test

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
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.ibm.com/SDSaaS/sds-go-sdk/sdsaasv1"
)

var _ = Describe(`SdsaasV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(sdsaasService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(sdsaasService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
				URL: "https://sdsaasv1/api",
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
				"SDSAAS_URL":       "https://sdsaasv1/api",
				"SDSAAS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1UsingExternalConfig(&sdsaasv1.SdsaasV1Options{})
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
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1UsingExternalConfig(&sdsaasv1.SdsaasV1Options{
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
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1UsingExternalConfig(&sdsaasv1.SdsaasV1Options{})
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
				"SDSAAS_URL":       "https://sdsaasv1/api",
				"SDSAAS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1UsingExternalConfig(&sdsaasv1.SdsaasV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(sdsaasService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SDSAAS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1UsingExternalConfig(&sdsaasv1.SdsaasV1Options{
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
			url, err = sdsaasv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := sdsaasv1.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("{url}"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := sdsaasv1.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`Volumes(volumesOptions *VolumesOptions) - Operation response error`, func() {
		volumesPath := "/volumes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Volumes with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := new(sdsaasv1.VolumesOptions)
				volumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumesOptionsModel.Name = core.StringPtr("my-host")
				volumesOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Volumes(volumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Volumes(volumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Volumes(volumesOptions *VolumesOptions)`, func() {
		volumesPath := "/volumes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20, "volumes": [{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}]}`)
				}))
			})
			It(`Invoke Volumes successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := new(sdsaasv1.VolumesOptions)
				volumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumesOptionsModel.Name = core.StringPtr("my-host")
				volumesOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumesWithContext(ctx, volumesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Volumes(volumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumesWithContext(ctx, volumesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20, "volumes": [{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}]}`)
				}))
			})
			It(`Invoke Volumes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Volumes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := new(sdsaasv1.VolumesOptions)
				volumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumesOptionsModel.Name = core.StringPtr("my-host")
				volumesOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Volumes(volumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Volumes with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := new(sdsaasv1.VolumesOptions)
				volumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumesOptionsModel.Name = core.StringPtr("my-host")
				volumesOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Volumes(volumesOptionsModel)
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
			It(`Invoke Volumes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := new(sdsaasv1.VolumesOptions)
				volumesOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumesOptionsModel.Name = core.StringPtr("my-host")
				volumesOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Volumes(volumesOptionsModel)
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
				responseObject := new(sdsaasv1.VolumeCollection)
				nextObject := new(sdsaasv1.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv1.VolumeCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv1.VolumeCollection)
				nextObject := new(sdsaasv1.PageLink)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1,"volumes":[{"bandwidth":1000,"capacity":30,"created_at":"2019-01-01T12:00:00.000Z","href":"Href","id":"ID","iops":10000,"name":"Name","resource_type":"ResourceType","status":"available","status_reasons":[{"code":"volume_not_found","message":"Specified resource not found","more_info":"MoreInfo"}],"volume_mappings":[{"status":"mapped","storage_identifier":{"subsystem_nqn":"nqn.2014-06.org:1234","namespace_id":1,"namespace_uuid":"NamespaceUUID","gateways":[{"ip_address":"IPAddress","port":22}]},"href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"ID","name":"Name"},"host":{"id":"ID","name":"Name","nqn":"Nqn"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"IPAddress","port":22}]}],"snapshot_count":5,"source_snapshot":{"id":"ID"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"volumes":[{"bandwidth":1000,"capacity":30,"created_at":"2019-01-01T12:00:00.000Z","href":"Href","id":"ID","iops":10000,"name":"Name","resource_type":"ResourceType","status":"available","status_reasons":[{"code":"volume_not_found","message":"Specified resource not found","more_info":"MoreInfo"}],"volume_mappings":[{"status":"mapped","storage_identifier":{"subsystem_nqn":"nqn.2014-06.org:1234","namespace_id":1,"namespace_uuid":"NamespaceUUID","gateways":[{"ip_address":"IPAddress","port":22}]},"href":"Href","id":"1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45","volume":{"id":"ID","name":"Name"},"host":{"id":"ID","name":"Name","nqn":"Nqn"},"subsystem_nqn":"nqn.2014-06.org:1234","namespace":{"id":1,"uuid":"UUID"},"gateways":[{"ip_address":"IPAddress","port":22}]}],"snapshot_count":5,"source_snapshot":{"id":"ID"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use VolumesPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				volumesOptionsModel := &sdsaasv1.VolumesOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name:  core.StringPtr("my-host"),
				}

				pager, err := sdsaasService.NewVolumesPager(volumesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv1.Volume
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use VolumesPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				volumesOptionsModel := &sdsaasv1.VolumesOptions{
					Limit: core.Int64Ptr(int64(20)),
					Name:  core.StringPtr("my-host"),
				}

				pager, err := sdsaasService.NewVolumesPager(volumesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`VolumeCreate(volumeCreateOptions *VolumeCreateOptions) - Operation response error`, func() {
		volumeCreatePath := "/volumes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.SourceSnapshot = sourceSnapshotModel
				volumeCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeCreate(volumeCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeCreate(volumeCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeCreate(volumeCreateOptions *VolumeCreateOptions)`, func() {
		volumeCreatePath := "/volumes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeCreatePath))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke VolumeCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.SourceSnapshot = sourceSnapshotModel
				volumeCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeCreateWithContext(ctx, volumeCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeCreate(volumeCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeCreateWithContext(ctx, volumeCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeCreatePath))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke VolumeCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.SourceSnapshot = sourceSnapshotModel
				volumeCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeCreate(volumeCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.SourceSnapshot = sourceSnapshotModel
				volumeCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeCreate(volumeCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeCreateOptions model with no property values
				volumeCreateOptionsModelNew := new(sdsaasv1.VolumeCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.VolumeCreate(volumeCreateOptionsModelNew)
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
			It(`Invoke VolumeCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				sourceSnapshotModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.SourceSnapshot = sourceSnapshotModel
				volumeCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeCreate(volumeCreateOptionsModel)
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
	Describe(`Volume(volumeOptions *VolumeOptions) - Operation response error`, func() {
		volumePath := "/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Volume with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeOptions model
				volumeOptionsModel := new(sdsaasv1.VolumeOptions)
				volumeOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Volume(volumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Volume(volumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Volume(volumeOptions *VolumeOptions)`, func() {
		volumePath := "/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke Volume successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeOptions model
				volumeOptionsModel := new(sdsaasv1.VolumeOptions)
				volumeOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeWithContext(ctx, volumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Volume(volumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeWithContext(ctx, volumeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke Volume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Volume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeOptions model
				volumeOptionsModel := new(sdsaasv1.VolumeOptions)
				volumeOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Volume(volumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Volume with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeOptions model
				volumeOptionsModel := new(sdsaasv1.VolumeOptions)
				volumeOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Volume(volumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeOptions model with no property values
				volumeOptionsModelNew := new(sdsaasv1.VolumeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.Volume(volumeOptionsModelNew)
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
			It(`Invoke Volume successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeOptions model
				volumeOptionsModel := new(sdsaasv1.VolumeOptions)
				volumeOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Volume(volumeOptionsModel)
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
	Describe(`VolumeDelete(volumeDeleteOptions *VolumeDeleteOptions)`, func() {
		volumeDeletePath := "/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke VolumeDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.VolumeDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the VolumeDeleteOptions model
				volumeDeleteOptionsModel := new(sdsaasv1.VolumeDeleteOptions)
				volumeDeleteOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.VolumeDelete(volumeDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke VolumeDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeDeleteOptions model
				volumeDeleteOptionsModel := new(sdsaasv1.VolumeDeleteOptions)
				volumeDeleteOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.VolumeDelete(volumeDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the VolumeDeleteOptions model with no property values
				volumeDeleteOptionsModelNew := new(sdsaasv1.VolumeDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.VolumeDelete(volumeDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeUpdate(volumeUpdateOptions *VolumeUpdateOptions) - Operation response error`, func() {
		volumeUpdatePath := "/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeUpdatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeUpdate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv1.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("testString")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeUpdateOptions model
				volumeUpdateOptionsModel := new(sdsaasv1.VolumeUpdateOptions)
				volumeUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.VolumePatch = volumePatchModelAsPatch
				volumeUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeUpdate(volumeUpdateOptions *VolumeUpdateOptions)`, func() {
		volumeUpdatePath := "/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke VolumeUpdate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv1.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("testString")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeUpdateOptions model
				volumeUpdateOptionsModel := new(sdsaasv1.VolumeUpdateOptions)
				volumeUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.VolumePatch = volumePatchModelAsPatch
				volumeUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeUpdateWithContext(ctx, volumeUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeUpdateWithContext(ctx, volumeUpdateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 30, "created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "available", "status_reasons": [{"code": "volume_not_found", "message": "Specified resource not found", "more_info": "MoreInfo"}], "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "snapshot_count": 5, "source_snapshot": {"id": "ID"}}`)
				}))
			})
			It(`Invoke VolumeUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeUpdate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv1.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("testString")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeUpdateOptions model
				volumeUpdateOptionsModel := new(sdsaasv1.VolumeUpdateOptions)
				volumeUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.VolumePatch = volumePatchModelAsPatch
				volumeUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeUpdate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv1.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("testString")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeUpdateOptions model
				volumeUpdateOptionsModel := new(sdsaasv1.VolumeUpdateOptions)
				volumeUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.VolumePatch = volumePatchModelAsPatch
				volumeUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeUpdateOptions model with no property values
				volumeUpdateOptionsModelNew := new(sdsaasv1.VolumeUpdateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.VolumeUpdate(volumeUpdateOptionsModelNew)
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
			It(`Invoke VolumeUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumePatch model
				volumePatchModel := new(sdsaasv1.VolumePatch)
				volumePatchModel.Capacity = core.Int64Ptr(int64(100))
				volumePatchModel.Name = core.StringPtr("testString")
				volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeUpdateOptions model
				volumeUpdateOptionsModel := new(sdsaasv1.VolumeUpdateOptions)
				volumeUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.VolumePatch = volumePatchModelAsPatch
				volumeUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeUpdate(volumeUpdateOptionsModel)
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
	Describe(`VolumeSnapshots(volumeSnapshotsOptions *VolumeSnapshotsOptions) - Operation response error`, func() {
		volumeSnapshotsPath := "/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeSnapshots with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := new(sdsaasv1.VolumeSnapshotsOptions)
				volumeSnapshotsOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumeSnapshotsOptionsModel.Name = core.StringPtr("my-host")
				volumeSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeSnapshots(volumeSnapshotsOptions *VolumeSnapshotsOptions)`, func() {
		volumeSnapshotsPath := "/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"snapshots": [{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke VolumeSnapshots successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := new(sdsaasv1.VolumeSnapshotsOptions)
				volumeSnapshotsOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumeSnapshotsOptionsModel.Name = core.StringPtr("my-host")
				volumeSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeSnapshotsWithContext(ctx, volumeSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeSnapshotsWithContext(ctx, volumeSnapshotsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"snapshots": [{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}], "first": {"href": "Href"}, "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke VolumeSnapshots successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := new(sdsaasv1.VolumeSnapshotsOptions)
				volumeSnapshotsOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumeSnapshotsOptionsModel.Name = core.StringPtr("my-host")
				volumeSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeSnapshots with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := new(sdsaasv1.VolumeSnapshotsOptions)
				volumeSnapshotsOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumeSnapshotsOptionsModel.Name = core.StringPtr("my-host")
				volumeSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
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
			It(`Invoke VolumeSnapshots successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := new(sdsaasv1.VolumeSnapshotsOptions)
				volumeSnapshotsOptionsModel.Start = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(20))
				volumeSnapshotsOptionsModel.Name = core.StringPtr("my-host")
				volumeSnapshotsOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeSnapshots(volumeSnapshotsOptionsModel)
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
				responseObject := new(sdsaasv1.SnapshotCollection)
				nextObject := new(sdsaasv1.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(sdsaasv1.SnapshotCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(sdsaasv1.SnapshotCollection)
				nextObject := new(sdsaasv1.PageLink)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"snapshots":[{"id":"ID","href":"Href","name":"Name","created_at":"2019-01-01T12:00:00.000Z","resource_type":"volume","lifecycle_state":"stable","size":30,"minimum_capacity":30,"deletable":true,"source_volume":{"id":"ID","href":"Href","name":"Name","resource_type":"volume"}}],"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"snapshots":[{"id":"ID","href":"Href","name":"Name","created_at":"2019-01-01T12:00:00.000Z","resource_type":"volume","lifecycle_state":"stable","size":30,"minimum_capacity":30,"deletable":true,"source_volume":{"id":"ID","href":"Href","name":"Name","resource_type":"volume"}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use VolumeSnapshotsPager.GetNext successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				volumeSnapshotsOptionsModel := &sdsaasv1.VolumeSnapshotsOptions{
					Limit:          core.Int64Ptr(int64(20)),
					Name:           core.StringPtr("my-host"),
					SourceVolumeID: core.StringPtr("testString"),
				}

				pager, err := sdsaasService.NewVolumeSnapshotsPager(volumeSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []sdsaasv1.Snapshot
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use VolumeSnapshotsPager.GetAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				volumeSnapshotsOptionsModel := &sdsaasv1.VolumeSnapshotsOptions{
					Limit:          core.Int64Ptr(int64(20)),
					Name:           core.StringPtr("my-host"),
					SourceVolumeID: core.StringPtr("testString"),
				}

				pager, err := sdsaasService.NewVolumeSnapshotsPager(volumeSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`VolumeSnapshotCreate(volumeSnapshotCreateOptions *VolumeSnapshotCreateOptions) - Operation response error`, func() {
		volumeSnapshotCreatePath := "/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeSnapshotCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeSnapshotCreateOptions model
				volumeSnapshotCreateOptionsModel := new(sdsaasv1.VolumeSnapshotCreateOptions)
				volumeSnapshotCreateOptionsModel.SourceVolume = sourceVolumePrototypeModel
				volumeSnapshotCreateOptionsModel.Name = core.StringPtr("my-snapshot")
				volumeSnapshotCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeSnapshotCreate(volumeSnapshotCreateOptions *VolumeSnapshotCreateOptions)`, func() {
		volumeSnapshotCreatePath := "/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotCreatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshotCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeSnapshotCreateOptions model
				volumeSnapshotCreateOptionsModel := new(sdsaasv1.VolumeSnapshotCreateOptions)
				volumeSnapshotCreateOptionsModel.SourceVolume = sourceVolumePrototypeModel
				volumeSnapshotCreateOptionsModel.Name = core.StringPtr("my-snapshot")
				volumeSnapshotCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeSnapshotCreateWithContext(ctx, volumeSnapshotCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeSnapshotCreateWithContext(ctx, volumeSnapshotCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotCreatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshotCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeSnapshotCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeSnapshotCreateOptions model
				volumeSnapshotCreateOptionsModel := new(sdsaasv1.VolumeSnapshotCreateOptions)
				volumeSnapshotCreateOptionsModel.SourceVolume = sourceVolumePrototypeModel
				volumeSnapshotCreateOptionsModel.Name = core.StringPtr("my-snapshot")
				volumeSnapshotCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeSnapshotCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeSnapshotCreateOptions model
				volumeSnapshotCreateOptionsModel := new(sdsaasv1.VolumeSnapshotCreateOptions)
				volumeSnapshotCreateOptionsModel.SourceVolume = sourceVolumePrototypeModel
				volumeSnapshotCreateOptionsModel.Name = core.StringPtr("my-snapshot")
				volumeSnapshotCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeSnapshotCreateOptions model with no property values
				volumeSnapshotCreateOptionsModelNew := new(sdsaasv1.VolumeSnapshotCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModelNew)
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
			It(`Invoke VolumeSnapshotCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeSnapshotCreateOptions model
				volumeSnapshotCreateOptionsModel := new(sdsaasv1.VolumeSnapshotCreateOptions)
				volumeSnapshotCreateOptionsModel.SourceVolume = sourceVolumePrototypeModel
				volumeSnapshotCreateOptionsModel.Name = core.StringPtr("my-snapshot")
				volumeSnapshotCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptionsModel)
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
	Describe(`VolumeSnapshotsDelete(volumeSnapshotsDeleteOptions *VolumeSnapshotsDeleteOptions)`, func() {
		volumeSnapshotsDeletePath := "/snapshots"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotsDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["source_volume.id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke VolumeSnapshotsDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.VolumeSnapshotsDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the VolumeSnapshotsDeleteOptions model
				volumeSnapshotsDeleteOptionsModel := new(sdsaasv1.VolumeSnapshotsDeleteOptions)
				volumeSnapshotsDeleteOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.VolumeSnapshotsDelete(volumeSnapshotsDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke VolumeSnapshotsDelete with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotsDeleteOptions model
				volumeSnapshotsDeleteOptionsModel := new(sdsaasv1.VolumeSnapshotsDeleteOptions)
				volumeSnapshotsDeleteOptionsModel.SourceVolumeID = core.StringPtr("testString")
				volumeSnapshotsDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.VolumeSnapshotsDelete(volumeSnapshotsDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeSnapshot(volumeSnapshotOptions *VolumeSnapshotOptions) - Operation response error`, func() {
		volumeSnapshotPath := "/snapshots/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeSnapshot with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotOptions model
				volumeSnapshotOptionsModel := new(sdsaasv1.VolumeSnapshotOptions)
				volumeSnapshotOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeSnapshot(volumeSnapshotOptions *VolumeSnapshotOptions)`, func() {
		volumeSnapshotPath := "/snapshots/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshot successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeSnapshotOptions model
				volumeSnapshotOptionsModel := new(sdsaasv1.VolumeSnapshotOptions)
				volumeSnapshotOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeSnapshotWithContext(ctx, volumeSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeSnapshotWithContext(ctx, volumeSnapshotOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeSnapshotOptions model
				volumeSnapshotOptionsModel := new(sdsaasv1.VolumeSnapshotOptions)
				volumeSnapshotOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeSnapshot with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotOptions model
				volumeSnapshotOptionsModel := new(sdsaasv1.VolumeSnapshotOptions)
				volumeSnapshotOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeSnapshotOptions model with no property values
				volumeSnapshotOptionsModelNew := new(sdsaasv1.VolumeSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModelNew)
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
			It(`Invoke VolumeSnapshot successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotOptions model
				volumeSnapshotOptionsModel := new(sdsaasv1.VolumeSnapshotOptions)
				volumeSnapshotOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeSnapshot(volumeSnapshotOptionsModel)
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
	Describe(`VolumeSnapshotUpdate(volumeSnapshotUpdateOptions *VolumeSnapshotUpdateOptions) - Operation response error`, func() {
		volumeSnapshotUpdatePath := "/snapshots/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotUpdatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke VolumeSnapshotUpdate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv1.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("testString")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeSnapshotUpdateOptions model
				volumeSnapshotUpdateOptionsModel := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				volumeSnapshotUpdateOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				volumeSnapshotUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`VolumeSnapshotUpdate(volumeSnapshotUpdateOptions *VolumeSnapshotUpdateOptions)`, func() {
		volumeSnapshotUpdatePath := "/snapshots/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshotUpdate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv1.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("testString")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeSnapshotUpdateOptions model
				volumeSnapshotUpdateOptionsModel := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				volumeSnapshotUpdateOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				volumeSnapshotUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.VolumeSnapshotUpdateWithContext(ctx, volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.VolumeSnapshotUpdateWithContext(ctx, volumeSnapshotUpdateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "name": "Name", "created_at": "2019-01-01T12:00:00.000Z", "resource_type": "volume", "lifecycle_state": "stable", "size": 30, "minimum_capacity": 30, "deletable": true, "source_volume": {"id": "ID", "href": "Href", "name": "Name", "resource_type": "volume"}}`)
				}))
			})
			It(`Invoke VolumeSnapshotUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.VolumeSnapshotUpdate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv1.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("testString")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeSnapshotUpdateOptions model
				volumeSnapshotUpdateOptionsModel := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				volumeSnapshotUpdateOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				volumeSnapshotUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke VolumeSnapshotUpdate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv1.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("testString")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeSnapshotUpdateOptions model
				volumeSnapshotUpdateOptionsModel := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				volumeSnapshotUpdateOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				volumeSnapshotUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the VolumeSnapshotUpdateOptions model with no property values
				volumeSnapshotUpdateOptionsModelNew := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModelNew)
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
			It(`Invoke VolumeSnapshotUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the SnapshotPatch model
				snapshotPatchModel := new(sdsaasv1.SnapshotPatch)
				snapshotPatchModel.Name = core.StringPtr("testString")
				snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the VolumeSnapshotUpdateOptions model
				volumeSnapshotUpdateOptionsModel := new(sdsaasv1.VolumeSnapshotUpdateOptions)
				volumeSnapshotUpdateOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SnapshotPatch = snapshotPatchModelAsPatch
				volumeSnapshotUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptionsModel)
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
	Describe(`VolumeSnapshotDelete(volumeSnapshotDeleteOptions *VolumeSnapshotDeleteOptions)`, func() {
		volumeSnapshotDeletePath := "/snapshots/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(volumeSnapshotDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke VolumeSnapshotDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.VolumeSnapshotDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the VolumeSnapshotDeleteOptions model
				volumeSnapshotDeleteOptionsModel := new(sdsaasv1.VolumeSnapshotDeleteOptions)
				volumeSnapshotDeleteOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.VolumeSnapshotDelete(volumeSnapshotDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke VolumeSnapshotDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeSnapshotDeleteOptions model
				volumeSnapshotDeleteOptionsModel := new(sdsaasv1.VolumeSnapshotDeleteOptions)
				volumeSnapshotDeleteOptionsModel.SnapID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.VolumeSnapshotDelete(volumeSnapshotDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the VolumeSnapshotDeleteOptions model with no property values
				volumeSnapshotDeleteOptionsModelNew := new(sdsaasv1.VolumeSnapshotDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.VolumeSnapshotDelete(volumeSnapshotDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Creds(credsOptions *CredsOptions) - Operation response error`, func() {
		credsPath := "/s3_credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Creds with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredsOptions model
				credsOptionsModel := new(sdsaasv1.CredsOptions)
				credsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Creds(credsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Creds(credsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Creds(credsOptions *CredsOptions)`, func() {
		credsPath := "/s3_credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"s3_credentials": ["S3Credentials"]}`)
				}))
			})
			It(`Invoke Creds successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CredsOptions model
				credsOptionsModel := new(sdsaasv1.CredsOptions)
				credsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CredsWithContext(ctx, credsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Creds(credsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CredsWithContext(ctx, credsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(credsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"s3_credentials": ["S3Credentials"]}`)
				}))
			})
			It(`Invoke Creds successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Creds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CredsOptions model
				credsOptionsModel := new(sdsaasv1.CredsOptions)
				credsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Creds(credsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Creds with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredsOptions model
				credsOptionsModel := new(sdsaasv1.CredsOptions)
				credsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Creds(credsOptionsModel)
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
			It(`Invoke Creds successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredsOptions model
				credsOptionsModel := new(sdsaasv1.CredsOptions)
				credsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Creds(credsOptionsModel)
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
	Describe(`CredCreate(credCreateOptions *CredCreateOptions) - Operation response error`, func() {
		credCreatePath := "/s3_credentials/mytestkey"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CredCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredCreateOptions model
				credCreateOptionsModel := new(sdsaasv1.CredCreateOptions)
				credCreateOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CredCreate(credCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CredCreate(credCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CredCreate(credCreateOptions *CredCreateOptions)`, func() {
		credCreatePath := "/s3_credentials/mytestkey"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credCreatePath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"access_key": "AccessKey", "secret_key": "SecretKey"}`)
				}))
			})
			It(`Invoke CredCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CredCreateOptions model
				credCreateOptionsModel := new(sdsaasv1.CredCreateOptions)
				credCreateOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CredCreateWithContext(ctx, credCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CredCreate(credCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CredCreateWithContext(ctx, credCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(credCreatePath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"access_key": "AccessKey", "secret_key": "SecretKey"}`)
				}))
			})
			It(`Invoke CredCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CredCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CredCreateOptions model
				credCreateOptionsModel := new(sdsaasv1.CredCreateOptions)
				credCreateOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CredCreate(credCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CredCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredCreateOptions model
				credCreateOptionsModel := new(sdsaasv1.CredCreateOptions)
				credCreateOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CredCreate(credCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CredCreateOptions model with no property values
				credCreateOptionsModelNew := new(sdsaasv1.CredCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CredCreate(credCreateOptionsModelNew)
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
			It(`Invoke CredCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredCreateOptions model
				credCreateOptionsModel := new(sdsaasv1.CredCreateOptions)
				credCreateOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CredCreate(credCreateOptionsModel)
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
	Describe(`CredDelete(credDeleteOptions *CredDeleteOptions)`, func() {
		credDeletePath := "/s3_credentials/mytestkey"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CredDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.CredDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CredDeleteOptions model
				credDeleteOptionsModel := new(sdsaasv1.CredDeleteOptions)
				credDeleteOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.CredDelete(credDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CredDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CredDeleteOptions model
				credDeleteOptionsModel := new(sdsaasv1.CredDeleteOptions)
				credDeleteOptionsModel.AccessKey = core.StringPtr("mytestkey")
				credDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.CredDelete(credDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CredDeleteOptions model with no property values
				credDeleteOptionsModelNew := new(sdsaasv1.CredDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.CredDelete(credDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertTypes(certTypesOptions *CertTypesOptions) - Operation response error`, func() {
		certTypesPath := "/certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certTypesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CertTypes with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := new(sdsaasv1.CertTypesOptions)
				certTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CertTypes(certTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CertTypes(certTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertTypes(certTypesOptions *CertTypesOptions)`, func() {
		certTypesPath := "/certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certTypesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificates": ["Certificates"]}`)
				}))
			})
			It(`Invoke CertTypes successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := new(sdsaasv1.CertTypesOptions)
				certTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CertTypesWithContext(ctx, certTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CertTypes(certTypesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CertTypesWithContext(ctx, certTypesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(certTypesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificates": ["Certificates"]}`)
				}))
			})
			It(`Invoke CertTypes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CertTypes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := new(sdsaasv1.CertTypesOptions)
				certTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CertTypes(certTypesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CertTypes with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := new(sdsaasv1.CertTypesOptions)
				certTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CertTypes(certTypesOptionsModel)
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
			It(`Invoke CertTypes successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := new(sdsaasv1.CertTypesOptions)
				certTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CertTypes(certTypesOptionsModel)
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
	Describe(`Cert(certOptions *CertOptions) - Operation response error`, func() {
		certPath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Cert with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Cert = core.StringPtr("s3")
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Cert(certOptions *CertOptions)`, func() {
		certPath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "expiration_date": "2019-01-01T12:00:00.000Z", "expired": false}`)
				}))
			})
			It(`Invoke Cert successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Cert = core.StringPtr("s3")
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CertWithContext(ctx, certOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CertWithContext(ctx, certOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(certPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "expiration_date": "2019-01-01T12:00:00.000Z", "expired": false}`)
				}))
			})
			It(`Invoke Cert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Cert(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Cert = core.StringPtr("s3")
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Cert with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Cert = core.StringPtr("s3")
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CertOptions model with no property values
				certOptionsModelNew := new(sdsaasv1.CertOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.Cert(certOptionsModelNew)
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
			It(`Invoke Cert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Cert = core.StringPtr("s3")
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Cert(certOptionsModel)
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
	Describe(`CertDelete(certDeleteOptions *CertDeleteOptions)`, func() {
		certDeletePath := "/certificates/s3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CertDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.CertDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CertDeleteOptions model
				certDeleteOptionsModel := new(sdsaasv1.CertDeleteOptions)
				certDeleteOptionsModel.Cert = core.StringPtr("s3")
				certDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.CertDelete(certDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CertDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertDeleteOptions model
				certDeleteOptionsModel := new(sdsaasv1.CertDeleteOptions)
				certDeleteOptionsModel.Cert = core.StringPtr("s3")
				certDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.CertDelete(certDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CertDeleteOptions model with no property values
				certDeleteOptionsModelNew := new(sdsaasv1.CertDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.CertDelete(certDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertCreate(certCreateOptions *CertCreateOptions) - Operation response error`, func() {
		certCreatePath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CertCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertCreateOptions model
				certCreateOptionsModel := new(sdsaasv1.CertCreateOptions)
				certCreateOptionsModel.Cert = core.StringPtr("s3")
				certCreateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CertCreate(certCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CertCreate(certCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertCreate(certCreateOptions *CertCreateOptions)`, func() {
		certCreatePath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certCreatePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "trace": "Trace", "errors": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CertCreateOptions model
				certCreateOptionsModel := new(sdsaasv1.CertCreateOptions)
				certCreateOptionsModel.Cert = core.StringPtr("s3")
				certCreateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CertCreateWithContext(ctx, certCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CertCreate(certCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CertCreateWithContext(ctx, certCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(certCreatePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "trace": "Trace", "errors": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CertCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertCreateOptions model
				certCreateOptionsModel := new(sdsaasv1.CertCreateOptions)
				certCreateOptionsModel.Cert = core.StringPtr("s3")
				certCreateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CertCreate(certCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CertCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertCreateOptions model
				certCreateOptionsModel := new(sdsaasv1.CertCreateOptions)
				certCreateOptionsModel.Cert = core.StringPtr("s3")
				certCreateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CertCreate(certCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CertCreateOptions model with no property values
				certCreateOptionsModelNew := new(sdsaasv1.CertCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CertCreate(certCreateOptionsModelNew)
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
			It(`Invoke CertCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertCreateOptions model
				certCreateOptionsModel := new(sdsaasv1.CertCreateOptions)
				certCreateOptionsModel.Cert = core.StringPtr("s3")
				certCreateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CertCreate(certCreateOptionsModel)
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
	Describe(`CertUpdate(certUpdateOptions *CertUpdateOptions) - Operation response error`, func() {
		certUpdatePath := "/certificates/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certUpdatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CertUpdate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUpdateOptions model
				certUpdateOptionsModel := new(sdsaasv1.CertUpdateOptions)
				certUpdateOptionsModel.Cert = core.StringPtr("s3")
				certUpdateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CertUpdate(certUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CertUpdate(certUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertUpdate(certUpdateOptions *CertUpdateOptions)`, func() {
		certUpdatePath := "/certificates/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "trace": "Trace", "errors": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertUpdate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CertUpdateOptions model
				certUpdateOptionsModel := new(sdsaasv1.CertUpdateOptions)
				certUpdateOptionsModel.Cert = core.StringPtr("s3")
				certUpdateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CertUpdateWithContext(ctx, certUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CertUpdate(certUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CertUpdateWithContext(ctx, certUpdateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(certUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "trace": "Trace", "errors": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CertUpdate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertUpdateOptions model
				certUpdateOptionsModel := new(sdsaasv1.CertUpdateOptions)
				certUpdateOptionsModel.Cert = core.StringPtr("s3")
				certUpdateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CertUpdate(certUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CertUpdate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUpdateOptions model
				certUpdateOptionsModel := new(sdsaasv1.CertUpdateOptions)
				certUpdateOptionsModel.Cert = core.StringPtr("s3")
				certUpdateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CertUpdate(certUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CertUpdateOptions model with no property values
				certUpdateOptionsModelNew := new(sdsaasv1.CertUpdateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CertUpdate(certUpdateOptionsModelNew)
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
			It(`Invoke CertUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUpdateOptions model
				certUpdateOptionsModel := new(sdsaasv1.CertUpdateOptions)
				certUpdateOptionsModel.Cert = core.StringPtr("s3")
				certUpdateOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CertUpdate(certUpdateOptionsModel)
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
	Describe(`Hosts(hostsOptions *HostsOptions) - Operation response error`, func() {
		hostsPath := "/hosts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Hosts with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostsOptions model
				hostsOptionsModel := new(sdsaasv1.HostsOptions)
				hostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				hostsOptionsModel.Name = core.StringPtr("my-host")
				hostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Hosts(hostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Hosts(hostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Hosts(hostsOptions *HostsOptions)`, func() {
		hostsPath := "/hosts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke Hosts successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostsOptions model
				hostsOptionsModel := new(sdsaasv1.HostsOptions)
				hostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				hostsOptionsModel.Name = core.StringPtr("my-host")
				hostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostsWithContext(ctx, hostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Hosts(hostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostsWithContext(ctx, hostsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-host"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke Hosts successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Hosts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostsOptions model
				hostsOptionsModel := new(sdsaasv1.HostsOptions)
				hostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				hostsOptionsModel.Name = core.StringPtr("my-host")
				hostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Hosts(hostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Hosts with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostsOptions model
				hostsOptionsModel := new(sdsaasv1.HostsOptions)
				hostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				hostsOptionsModel.Name = core.StringPtr("my-host")
				hostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Hosts(hostsOptionsModel)
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
			It(`Invoke Hosts successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostsOptions model
				hostsOptionsModel := new(sdsaasv1.HostsOptions)
				hostsOptionsModel.Limit = core.Int64Ptr(int64(20))
				hostsOptionsModel.Name = core.StringPtr("my-host")
				hostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Hosts(hostsOptionsModel)
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
	Describe(`HostCreate(hostCreateOptions *HostCreateOptions) - Operation response error`, func() {
		hostCreatePath := "/hosts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.VolumeMappings = []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				hostCreateOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostCreate(hostCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostCreate(hostCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostCreate(hostCreateOptions *HostCreateOptions)`, func() {
		hostCreatePath := "/hosts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostCreatePath))
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
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke HostCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.VolumeMappings = []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				hostCreateOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostCreateWithContext(ctx, hostCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostCreate(hostCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostCreateWithContext(ctx, hostCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostCreatePath))
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
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke HostCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.VolumeMappings = []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				hostCreateOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostCreate(hostCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.VolumeMappings = []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				hostCreateOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostCreate(hostCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostCreateOptions model with no property values
				hostCreateOptionsModelNew := new(sdsaasv1.HostCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostCreate(hostCreateOptionsModelNew)
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
			It(`Invoke HostCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				volumeMappingPrototypeModel.Volume = volumeIdentityModel

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.VolumeMappings = []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}
				hostCreateOptionsModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostCreate(hostCreateOptionsModel)
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
	Describe(`Host(hostOptions *HostOptions) - Operation response error`, func() {
		hostPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Host with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostOptions model
				hostOptionsModel := new(sdsaasv1.HostOptions)
				hostOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.Host(hostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.Host(hostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Host(hostOptions *HostOptions)`, func() {
		hostPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke Host successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostOptions model
				hostOptionsModel := new(sdsaasv1.HostOptions)
				hostOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostWithContext(ctx, hostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.Host(hostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostWithContext(ctx, hostOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke Host successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.Host(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostOptions model
				hostOptionsModel := new(sdsaasv1.HostOptions)
				hostOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Host(hostOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Host with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostOptions model
				hostOptionsModel := new(sdsaasv1.HostOptions)
				hostOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Host(hostOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostOptions model with no property values
				hostOptionsModelNew := new(sdsaasv1.HostOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.Host(hostOptionsModelNew)
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
			It(`Invoke Host successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostOptions model
				hostOptionsModel := new(sdsaasv1.HostOptions)
				hostOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.Host(hostOptionsModel)
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
	Describe(`HostUpdate(hostUpdateOptions *HostUpdateOptions) - Operation response error`, func() {
		hostUpdatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostUpdatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostUpdate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv1.HostPatch)
				hostPatchModel.Name = core.StringPtr("testString")
				hostPatchModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the HostUpdateOptions model
				hostUpdateOptionsModel := new(sdsaasv1.HostUpdateOptions)
				hostUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.HostPatch = hostPatchModelAsPatch
				hostUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostUpdate(hostUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostUpdate(hostUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostUpdate(hostUpdateOptions *HostUpdateOptions)`, func() {
		hostUpdatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke HostUpdate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv1.HostPatch)
				hostPatchModel.Name = core.StringPtr("testString")
				hostPatchModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the HostUpdateOptions model
				hostUpdateOptionsModel := new(sdsaasv1.HostUpdateOptions)
				hostUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.HostPatch = hostPatchModelAsPatch
				hostUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostUpdateWithContext(ctx, hostUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostUpdate(hostUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostUpdateWithContext(ctx, hostUpdateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostUpdatePath))
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
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "href": "Href", "id": "ID", "name": "Name", "nqn": "nqn.2014-06.org:1234", "psk_enabled": true, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}]}`)
				}))
			})
			It(`Invoke HostUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostUpdate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv1.HostPatch)
				hostPatchModel.Name = core.StringPtr("testString")
				hostPatchModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the HostUpdateOptions model
				hostUpdateOptionsModel := new(sdsaasv1.HostUpdateOptions)
				hostUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.HostPatch = hostPatchModelAsPatch
				hostUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostUpdate(hostUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostUpdate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv1.HostPatch)
				hostPatchModel.Name = core.StringPtr("testString")
				hostPatchModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the HostUpdateOptions model
				hostUpdateOptionsModel := new(sdsaasv1.HostUpdateOptions)
				hostUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.HostPatch = hostPatchModelAsPatch
				hostUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostUpdate(hostUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostUpdateOptions model with no property values
				hostUpdateOptionsModelNew := new(sdsaasv1.HostUpdateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostUpdate(hostUpdateOptionsModelNew)
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
			It(`Invoke HostUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostPatch model
				hostPatchModel := new(sdsaasv1.HostPatch)
				hostPatchModel.Name = core.StringPtr("testString")
				hostPatchModel.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the HostUpdateOptions model
				hostUpdateOptionsModel := new(sdsaasv1.HostUpdateOptions)
				hostUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.HostPatch = hostPatchModelAsPatch
				hostUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostUpdate(hostUpdateOptionsModel)
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
	Describe(`HostDelete(hostDeleteOptions *HostDeleteOptions)`, func() {
		hostDeletePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke HostDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.HostDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HostDeleteOptions model
				hostDeleteOptionsModel := new(sdsaasv1.HostDeleteOptions)
				hostDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.HostDelete(hostDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HostDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostDeleteOptions model
				hostDeleteOptionsModel := new(sdsaasv1.HostDeleteOptions)
				hostDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.HostDelete(hostDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HostDeleteOptions model with no property values
				hostDeleteOptionsModelNew := new(sdsaasv1.HostDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.HostDelete(hostDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostMappings(hostMappingsOptions *HostMappingsOptions) - Operation response error`, func() {
		hostMappingsPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostMappings with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingsOptions model
				hostMappingsOptionsModel := new(sdsaasv1.HostMappingsOptions)
				hostMappingsOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostMappings(hostMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostMappings(hostMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostMappings(hostMappingsOptions *HostMappingsOptions)`, func() {
		hostMappingsPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke HostMappings successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostMappingsOptions model
				hostMappingsOptionsModel := new(sdsaasv1.HostMappingsOptions)
				hostMappingsOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostMappingsWithContext(ctx, hostMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostMappings(hostMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostMappingsWithContext(ctx, hostMappingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "volume_mappings": [{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}], "limit": 20, "next": {"href": "Href"}, "total_count": 20}`)
				}))
			})
			It(`Invoke HostMappings successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostMappings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostMappingsOptions model
				hostMappingsOptionsModel := new(sdsaasv1.HostMappingsOptions)
				hostMappingsOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostMappings(hostMappingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostMappings with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingsOptions model
				hostMappingsOptionsModel := new(sdsaasv1.HostMappingsOptions)
				hostMappingsOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostMappings(hostMappingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostMappingsOptions model with no property values
				hostMappingsOptionsModelNew := new(sdsaasv1.HostMappingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostMappings(hostMappingsOptionsModelNew)
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
			It(`Invoke HostMappings successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingsOptions model
				hostMappingsOptionsModel := new(sdsaasv1.HostMappingsOptions)
				hostMappingsOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostMappings(hostMappingsOptionsModel)
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
	Describe(`HostMappingCreate(hostMappingCreateOptions *HostMappingCreateOptions) - Operation response error`, func() {
		hostMappingCreatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingCreatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostMappingCreate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the HostMappingCreateOptions model
				hostMappingCreateOptionsModel := new(sdsaasv1.HostMappingCreateOptions)
				hostMappingCreateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.Volume = volumeIdentityModel
				hostMappingCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostMappingCreate(hostMappingCreateOptions *HostMappingCreateOptions)`, func() {
		hostMappingCreatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingCreatePath))
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
					fmt.Fprintf(res, "%s", `{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}`)
				}))
			})
			It(`Invoke HostMappingCreate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the HostMappingCreateOptions model
				hostMappingCreateOptionsModel := new(sdsaasv1.HostMappingCreateOptions)
				hostMappingCreateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.Volume = volumeIdentityModel
				hostMappingCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostMappingCreateWithContext(ctx, hostMappingCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostMappingCreateWithContext(ctx, hostMappingCreateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingCreatePath))
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
					fmt.Fprintf(res, "%s", `{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}`)
				}))
			})
			It(`Invoke HostMappingCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostMappingCreate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the HostMappingCreateOptions model
				hostMappingCreateOptionsModel := new(sdsaasv1.HostMappingCreateOptions)
				hostMappingCreateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.Volume = volumeIdentityModel
				hostMappingCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostMappingCreate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the HostMappingCreateOptions model
				hostMappingCreateOptionsModel := new(sdsaasv1.HostMappingCreateOptions)
				hostMappingCreateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.Volume = volumeIdentityModel
				hostMappingCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostMappingCreateOptions model with no property values
				hostMappingCreateOptionsModelNew := new(sdsaasv1.HostMappingCreateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostMappingCreate(hostMappingCreateOptionsModelNew)
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
			It(`Invoke HostMappingCreate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				volumeIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the HostMappingCreateOptions model
				hostMappingCreateOptionsModel := new(sdsaasv1.HostMappingCreateOptions)
				hostMappingCreateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.Volume = volumeIdentityModel
				hostMappingCreateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostMappingCreate(hostMappingCreateOptionsModel)
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
	Describe(`HostMappingDeleteAll(hostMappingDeleteAllOptions *HostMappingDeleteAllOptions)`, func() {
		hostMappingDeleteAllPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingDeleteAllPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke HostMappingDeleteAll successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.HostMappingDeleteAll(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HostMappingDeleteAllOptions model
				hostMappingDeleteAllOptionsModel := new(sdsaasv1.HostMappingDeleteAllOptions)
				hostMappingDeleteAllOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.HostMappingDeleteAll(hostMappingDeleteAllOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HostMappingDeleteAll with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingDeleteAllOptions model
				hostMappingDeleteAllOptionsModel := new(sdsaasv1.HostMappingDeleteAllOptions)
				hostMappingDeleteAllOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.HostMappingDeleteAll(hostMappingDeleteAllOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HostMappingDeleteAllOptions model with no property values
				hostMappingDeleteAllOptionsModelNew := new(sdsaasv1.HostMappingDeleteAllOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.HostMappingDeleteAll(hostMappingDeleteAllOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostMapping(hostMappingOptions *HostMappingOptions) - Operation response error`, func() {
		hostMappingPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostMapping with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingOptions model
				hostMappingOptionsModel := new(sdsaasv1.HostMappingOptions)
				hostMappingOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostMapping(hostMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostMapping(hostMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostMapping(hostMappingOptions *HostMappingOptions)`, func() {
		hostMappingPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}`)
				}))
			})
			It(`Invoke HostMapping successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostMappingOptions model
				hostMappingOptionsModel := new(sdsaasv1.HostMappingOptions)
				hostMappingOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostMappingWithContext(ctx, hostMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostMapping(hostMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostMappingWithContext(ctx, hostMappingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "mapped", "storage_identifier": {"subsystem_nqn": "nqn.2014-06.org:1234", "namespace_id": 1, "namespace_uuid": "NamespaceUUID", "gateways": [{"ip_address": "IPAddress", "port": 22}]}, "href": "Href", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45", "volume": {"id": "ID", "name": "Name"}, "host": {"id": "ID", "name": "Name", "nqn": "Nqn"}, "subsystem_nqn": "nqn.2014-06.org:1234", "namespace": {"id": 1, "uuid": "UUID"}, "gateways": [{"ip_address": "IPAddress", "port": 22}]}`)
				}))
			})
			It(`Invoke HostMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostMapping(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostMappingOptions model
				hostMappingOptionsModel := new(sdsaasv1.HostMappingOptions)
				hostMappingOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostMapping(hostMappingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostMapping with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingOptions model
				hostMappingOptionsModel := new(sdsaasv1.HostMappingOptions)
				hostMappingOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostMapping(hostMappingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostMappingOptions model with no property values
				hostMappingOptionsModelNew := new(sdsaasv1.HostMappingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostMapping(hostMappingOptionsModelNew)
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
			It(`Invoke HostMapping successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingOptions model
				hostMappingOptionsModel := new(sdsaasv1.HostMappingOptions)
				hostMappingOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostMapping(hostMappingOptionsModel)
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
	Describe(`HostMappingDelete(hostMappingDeleteOptions *HostMappingDeleteOptions)`, func() {
		hostMappingDeletePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volume_mappings/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostMappingDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke HostMappingDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.HostMappingDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HostMappingDeleteOptions model
				hostMappingDeleteOptionsModel := new(sdsaasv1.HostMappingDeleteOptions)
				hostMappingDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.HostMappingDelete(hostMappingDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HostMappingDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostMappingDeleteOptions model
				hostMappingDeleteOptionsModel := new(sdsaasv1.HostMappingDeleteOptions)
				hostMappingDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteOptionsModel.VolumeMappingID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.HostMappingDelete(hostMappingDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HostMappingDeleteOptions model with no property values
				hostMappingDeleteOptionsModelNew := new(sdsaasv1.HostMappingDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.HostMappingDelete(hostMappingDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			sdsaasService, _ := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
				URL:           "http://sdsaasv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCertCreateOptions successfully`, func() {
				// Construct an instance of the CertCreateOptions model
				cert := "s3"
				body := CreateMockReader("This is a mock file.")
				certCreateOptionsModel := sdsaasService.NewCertCreateOptions(cert, body)
				certCreateOptionsModel.SetCert("s3")
				certCreateOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				certCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certCreateOptionsModel).ToNot(BeNil())
				Expect(certCreateOptionsModel.Cert).To(Equal(core.StringPtr("s3")))
				Expect(certCreateOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(certCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCertDeleteOptions successfully`, func() {
				// Construct an instance of the CertDeleteOptions model
				cert := "s3"
				certDeleteOptionsModel := sdsaasService.NewCertDeleteOptions(cert)
				certDeleteOptionsModel.SetCert("s3")
				certDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certDeleteOptionsModel).ToNot(BeNil())
				Expect(certDeleteOptionsModel.Cert).To(Equal(core.StringPtr("s3")))
				Expect(certDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCertOptions successfully`, func() {
				// Construct an instance of the CertOptions model
				cert := "s3"
				certOptionsModel := sdsaasService.NewCertOptions(cert)
				certOptionsModel.SetCert("s3")
				certOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certOptionsModel).ToNot(BeNil())
				Expect(certOptionsModel.Cert).To(Equal(core.StringPtr("s3")))
				Expect(certOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCertTypesOptions successfully`, func() {
				// Construct an instance of the CertTypesOptions model
				certTypesOptionsModel := sdsaasService.NewCertTypesOptions()
				certTypesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certTypesOptionsModel).ToNot(BeNil())
				Expect(certTypesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCertUpdateOptions successfully`, func() {
				// Construct an instance of the CertUpdateOptions model
				cert := "s3"
				body := CreateMockReader("This is a mock file.")
				certUpdateOptionsModel := sdsaasService.NewCertUpdateOptions(cert, body)
				certUpdateOptionsModel.SetCert("s3")
				certUpdateOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				certUpdateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certUpdateOptionsModel).ToNot(BeNil())
				Expect(certUpdateOptionsModel.Cert).To(Equal(core.StringPtr("s3")))
				Expect(certUpdateOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(certUpdateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCredCreateOptions successfully`, func() {
				// Construct an instance of the CredCreateOptions model
				accessKey := "mytestkey"
				credCreateOptionsModel := sdsaasService.NewCredCreateOptions(accessKey)
				credCreateOptionsModel.SetAccessKey("mytestkey")
				credCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(credCreateOptionsModel).ToNot(BeNil())
				Expect(credCreateOptionsModel.AccessKey).To(Equal(core.StringPtr("mytestkey")))
				Expect(credCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCredDeleteOptions successfully`, func() {
				// Construct an instance of the CredDeleteOptions model
				accessKey := "mytestkey"
				credDeleteOptionsModel := sdsaasService.NewCredDeleteOptions(accessKey)
				credDeleteOptionsModel.SetAccessKey("mytestkey")
				credDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(credDeleteOptionsModel).ToNot(BeNil())
				Expect(credDeleteOptionsModel.AccessKey).To(Equal(core.StringPtr("mytestkey")))
				Expect(credDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCredsOptions successfully`, func() {
				// Construct an instance of the CredsOptions model
				credsOptionsModel := sdsaasService.NewCredsOptions()
				credsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(credsOptionsModel).ToNot(BeNil())
				Expect(credsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostCreateOptions successfully`, func() {
				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				Expect(volumeIdentityModel).ToNot(BeNil())
				volumeIdentityModel.ID = core.StringPtr("testString")
				Expect(volumeIdentityModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VolumeMappingPrototype model
				volumeMappingPrototypeModel := new(sdsaasv1.VolumeMappingPrototype)
				Expect(volumeMappingPrototypeModel).ToNot(BeNil())
				volumeMappingPrototypeModel.Volume = volumeIdentityModel
				Expect(volumeMappingPrototypeModel.Volume).To(Equal(volumeIdentityModel))

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsNqn := "nqn.2014-06.org:9345"
				hostCreateOptionsModel := sdsaasService.NewHostCreateOptions(hostCreateOptionsNqn)
				hostCreateOptionsModel.SetNqn("nqn.2014-06.org:9345")
				hostCreateOptionsModel.SetName("my-host")
				hostCreateOptionsModel.SetVolumeMappings([]sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel})
				hostCreateOptionsModel.SetPsk("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")
				hostCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostCreateOptionsModel).ToNot(BeNil())
				Expect(hostCreateOptionsModel.Nqn).To(Equal(core.StringPtr("nqn.2014-06.org:9345")))
				Expect(hostCreateOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(hostCreateOptionsModel.VolumeMappings).To(Equal([]sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel}))
				Expect(hostCreateOptionsModel.Psk).To(Equal(core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")))
				Expect(hostCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostDeleteOptions successfully`, func() {
				// Construct an instance of the HostDeleteOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostDeleteOptionsModel := sdsaasService.NewHostDeleteOptions(hostID)
				hostDeleteOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostDeleteOptionsModel).ToNot(BeNil())
				Expect(hostDeleteOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostMappingCreateOptions successfully`, func() {
				// Construct an instance of the VolumeIdentity model
				volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
				Expect(volumeIdentityModel).ToNot(BeNil())
				volumeIdentityModel.ID = core.StringPtr("testString")
				Expect(volumeIdentityModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the HostMappingCreateOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				var hostMappingCreateOptionsVolume *sdsaasv1.VolumeIdentity = nil
				hostMappingCreateOptionsModel := sdsaasService.NewHostMappingCreateOptions(hostID, hostMappingCreateOptionsVolume)
				hostMappingCreateOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingCreateOptionsModel.SetVolume(volumeIdentityModel)
				hostMappingCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostMappingCreateOptionsModel).ToNot(BeNil())
				Expect(hostMappingCreateOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostMappingCreateOptionsModel.Volume).To(Equal(volumeIdentityModel))
				Expect(hostMappingCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostMappingDeleteAllOptions successfully`, func() {
				// Construct an instance of the HostMappingDeleteAllOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostMappingDeleteAllOptionsModel := sdsaasService.NewHostMappingDeleteAllOptions(hostID)
				hostMappingDeleteAllOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteAllOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostMappingDeleteAllOptionsModel).ToNot(BeNil())
				Expect(hostMappingDeleteAllOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostMappingDeleteAllOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostMappingDeleteOptions successfully`, func() {
				// Construct an instance of the HostMappingDeleteOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeMappingID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				hostMappingDeleteOptionsModel := sdsaasService.NewHostMappingDeleteOptions(hostID, volumeMappingID)
				hostMappingDeleteOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingDeleteOptionsModel.SetVolumeMappingID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostMappingDeleteOptionsModel).ToNot(BeNil())
				Expect(hostMappingDeleteOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostMappingDeleteOptionsModel.VolumeMappingID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(hostMappingDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostMappingOptions successfully`, func() {
				// Construct an instance of the HostMappingOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeMappingID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				hostMappingOptionsModel := sdsaasService.NewHostMappingOptions(hostID, volumeMappingID)
				hostMappingOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingOptionsModel.SetVolumeMappingID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostMappingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostMappingOptionsModel).ToNot(BeNil())
				Expect(hostMappingOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostMappingOptionsModel.VolumeMappingID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(hostMappingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostMappingsOptions successfully`, func() {
				// Construct an instance of the HostMappingsOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostMappingsOptionsModel := sdsaasService.NewHostMappingsOptions(hostID)
				hostMappingsOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostMappingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostMappingsOptionsModel).ToNot(BeNil())
				Expect(hostMappingsOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostMappingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostOptions successfully`, func() {
				// Construct an instance of the HostOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostOptionsModel := sdsaasService.NewHostOptions(hostID)
				hostOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostOptionsModel).ToNot(BeNil())
				Expect(hostOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostUpdateOptions successfully`, func() {
				// Construct an instance of the HostUpdateOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostUpdateOptionsModel := sdsaasService.NewHostUpdateOptions(hostID)
				hostUpdateOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostUpdateOptionsModel.SetHostPatch(map[string]interface{}{"anyKey": "anyValue"})
				hostUpdateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostUpdateOptionsModel).ToNot(BeNil())
				Expect(hostUpdateOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostUpdateOptionsModel.HostPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(hostUpdateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostsOptions successfully`, func() {
				// Construct an instance of the HostsOptions model
				hostsOptionsModel := sdsaasService.NewHostsOptions()
				hostsOptionsModel.SetLimit(int64(20))
				hostsOptionsModel.SetName("my-host")
				hostsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostsOptionsModel).ToNot(BeNil())
				Expect(hostsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(hostsOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(hostsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSourceSnapshot successfully`, func() {
				id := "testString"
				_model, err := sdsaasService.NewSourceSnapshot(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceVolumePrototype successfully`, func() {
				id := "testString"
				_model, err := sdsaasService.NewSourceVolumePrototype(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeCreateOptions successfully`, func() {
				// Construct an instance of the SourceSnapshot model
				sourceSnapshotModel := new(sdsaasv1.SourceSnapshot)
				Expect(sourceSnapshotModel).ToNot(BeNil())
				sourceSnapshotModel.ID = core.StringPtr("testString")
				Expect(sourceSnapshotModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsCapacity := int64(10)
				volumeCreateOptionsModel := sdsaasService.NewVolumeCreateOptions(volumeCreateOptionsCapacity)
				volumeCreateOptionsModel.SetCapacity(int64(10))
				volumeCreateOptionsModel.SetName("my-volume")
				volumeCreateOptionsModel.SetSourceSnapshot(sourceSnapshotModel)
				volumeCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeCreateOptionsModel).ToNot(BeNil())
				Expect(volumeCreateOptionsModel.Capacity).To(Equal(core.Int64Ptr(int64(10))))
				Expect(volumeCreateOptionsModel.Name).To(Equal(core.StringPtr("my-volume")))
				Expect(volumeCreateOptionsModel.SourceSnapshot).To(Equal(sourceSnapshotModel))
				Expect(volumeCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeDeleteOptions successfully`, func() {
				// Construct an instance of the VolumeDeleteOptions model
				volumeID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				volumeDeleteOptionsModel := sdsaasService.NewVolumeDeleteOptions(volumeID)
				volumeDeleteOptionsModel.SetVolumeID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeDeleteOptionsModel).ToNot(BeNil())
				Expect(volumeDeleteOptionsModel.VolumeID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(volumeDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeIdentity successfully`, func() {
				id := "testString"
				_model, err := sdsaasService.NewVolumeIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeMappingPrototype successfully`, func() {
				var volume *sdsaasv1.VolumeIdentity = nil
				_, err := sdsaasService.NewVolumeMappingPrototype(volume)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVolumeOptions successfully`, func() {
				// Construct an instance of the VolumeOptions model
				volumeID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				volumeOptionsModel := sdsaasService.NewVolumeOptions(volumeID)
				volumeOptionsModel.SetVolumeID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeOptionsModel).ToNot(BeNil())
				Expect(volumeOptionsModel.VolumeID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(volumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotCreateOptions successfully`, func() {
				// Construct an instance of the SourceVolumePrototype model
				sourceVolumePrototypeModel := new(sdsaasv1.SourceVolumePrototype)
				Expect(sourceVolumePrototypeModel).ToNot(BeNil())
				sourceVolumePrototypeModel.ID = core.StringPtr("testString")
				Expect(sourceVolumePrototypeModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VolumeSnapshotCreateOptions model
				var volumeSnapshotCreateOptionsSourceVolume *sdsaasv1.SourceVolumePrototype = nil
				volumeSnapshotCreateOptionsModel := sdsaasService.NewVolumeSnapshotCreateOptions(volumeSnapshotCreateOptionsSourceVolume)
				volumeSnapshotCreateOptionsModel.SetSourceVolume(sourceVolumePrototypeModel)
				volumeSnapshotCreateOptionsModel.SetName("my-snapshot")
				volumeSnapshotCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotCreateOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotCreateOptionsModel.SourceVolume).To(Equal(sourceVolumePrototypeModel))
				Expect(volumeSnapshotCreateOptionsModel.Name).To(Equal(core.StringPtr("my-snapshot")))
				Expect(volumeSnapshotCreateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotDeleteOptions successfully`, func() {
				// Construct an instance of the VolumeSnapshotDeleteOptions model
				snapID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeSnapshotDeleteOptionsModel := sdsaasService.NewVolumeSnapshotDeleteOptions(snapID)
				volumeSnapshotDeleteOptionsModel.SetSnapID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotDeleteOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotDeleteOptionsModel.SnapID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(volumeSnapshotDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotOptions successfully`, func() {
				// Construct an instance of the VolumeSnapshotOptions model
				snapID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeSnapshotOptionsModel := sdsaasService.NewVolumeSnapshotOptions(snapID)
				volumeSnapshotOptionsModel.SetSnapID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotOptionsModel.SnapID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(volumeSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotUpdateOptions successfully`, func() {
				// Construct an instance of the VolumeSnapshotUpdateOptions model
				snapID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				snapshotPatch := map[string]interface{}{"anyKey": "anyValue"}
				volumeSnapshotUpdateOptionsModel := sdsaasService.NewVolumeSnapshotUpdateOptions(snapID, snapshotPatch)
				volumeSnapshotUpdateOptionsModel.SetSnapID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotUpdateOptionsModel.SetSnapshotPatch(map[string]interface{}{"anyKey": "anyValue"})
				volumeSnapshotUpdateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotUpdateOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotUpdateOptionsModel.SnapID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(volumeSnapshotUpdateOptionsModel.SnapshotPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(volumeSnapshotUpdateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotsDeleteOptions successfully`, func() {
				// Construct an instance of the VolumeSnapshotsDeleteOptions model
				volumeSnapshotsDeleteOptionsModel := sdsaasService.NewVolumeSnapshotsDeleteOptions()
				volumeSnapshotsDeleteOptionsModel.SetSourceVolumeID("testString")
				volumeSnapshotsDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotsDeleteOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotsDeleteOptionsModel.SourceVolumeID).To(Equal(core.StringPtr("testString")))
				Expect(volumeSnapshotsDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeSnapshotsOptions successfully`, func() {
				// Construct an instance of the VolumeSnapshotsOptions model
				volumeSnapshotsOptionsModel := sdsaasService.NewVolumeSnapshotsOptions()
				volumeSnapshotsOptionsModel.SetStart("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumeSnapshotsOptionsModel.SetLimit(int64(20))
				volumeSnapshotsOptionsModel.SetName("my-host")
				volumeSnapshotsOptionsModel.SetSourceVolumeID("testString")
				volumeSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeSnapshotsOptionsModel).ToNot(BeNil())
				Expect(volumeSnapshotsOptionsModel.Start).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(volumeSnapshotsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(volumeSnapshotsOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(volumeSnapshotsOptionsModel.SourceVolumeID).To(Equal(core.StringPtr("testString")))
				Expect(volumeSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeUpdateOptions successfully`, func() {
				// Construct an instance of the VolumeUpdateOptions model
				volumeID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				volumeUpdateOptionsModel := sdsaasService.NewVolumeUpdateOptions(volumeID)
				volumeUpdateOptionsModel.SetVolumeID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				volumeUpdateOptionsModel.SetVolumePatch(map[string]interface{}{"anyKey": "anyValue"})
				volumeUpdateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeUpdateOptionsModel).ToNot(BeNil())
				Expect(volumeUpdateOptionsModel.VolumeID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(volumeUpdateOptionsModel.VolumePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(volumeUpdateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumesOptions successfully`, func() {
				// Construct an instance of the VolumesOptions model
				volumesOptionsModel := sdsaasService.NewVolumesOptions()
				volumesOptionsModel.SetLimit(int64(20))
				volumesOptionsModel.SetName("my-host")
				volumesOptionsModel.SetStart("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				volumesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumesOptionsModel).ToNot(BeNil())
				Expect(volumesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(volumesOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(volumesOptionsModel.Start).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(volumesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalHostPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.HostPatch)
			model.Name = core.StringPtr("testString")
			model.Psk = core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.HostPatch
			err = sdsaasv1.UnmarshalHostPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSnapshotPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.SnapshotPatch)
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.SnapshotPatch
			err = sdsaasv1.UnmarshalSnapshotPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceSnapshot successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.SourceSnapshot)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.SourceSnapshot
			err = sdsaasv1.UnmarshalSourceSnapshot(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceVolumePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.SourceVolumePrototype)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.SourceVolumePrototype
			err = sdsaasv1.UnmarshalSourceVolumePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumeIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.VolumeIdentity)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.VolumeIdentity
			err = sdsaasv1.UnmarshalVolumeIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumeMappingPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.VolumeMappingPrototype)
			model.Volume = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.VolumeMappingPrototype
			err = sdsaasv1.UnmarshalVolumeMappingPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.VolumePatch)
			model.Capacity = core.Int64Ptr(int64(100))
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.VolumePatch
			err = sdsaasv1.UnmarshalVolumePatch(raw, &result)
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
