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
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
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
				volumesOptionsModel.Limit = core.Int64Ptr(int64(10))
				volumesOptionsModel.Name = core.StringPtr("myhost1")
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 5, "next": {"href": "Href"}, "total_count": 10, "volumes": [{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}]}`)
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
				volumesOptionsModel.Limit = core.Int64Ptr(int64(10))
				volumesOptionsModel.Name = core.StringPtr("myhost1")
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 5, "next": {"href": "Href"}, "total_count": 10, "volumes": [{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}]}`)
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
				volumesOptionsModel.Limit = core.Int64Ptr(int64(10))
				volumesOptionsModel.Name = core.StringPtr("myhost1")
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
				volumesOptionsModel.Limit = core.Int64Ptr(int64(10))
				volumesOptionsModel.Name = core.StringPtr("myhost1")
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
				volumesOptionsModel.Limit = core.Int64Ptr(int64(10))
				volumesOptionsModel.Name = core.StringPtr("myhost1")
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
					Expect(req.URL.Query()["hostnqnstring"]).To(Equal([]string{"nqn.2024-07.org:1234"}))
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

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.Hostnqnstring = core.StringPtr("nqn.2024-07.org:1234")
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

					Expect(req.URL.Query()["hostnqnstring"]).To(Equal([]string{"nqn.2024-07.org:1234"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.Hostnqnstring = core.StringPtr("nqn.2024-07.org:1234")
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

					Expect(req.URL.Query()["hostnqnstring"]).To(Equal([]string{"nqn.2024-07.org:1234"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.Hostnqnstring = core.StringPtr("nqn.2024-07.org:1234")
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

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.Hostnqnstring = core.StringPtr("nqn.2024-07.org:1234")
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

				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsModel := new(sdsaasv1.VolumeCreateOptions)
				volumeCreateOptionsModel.Capacity = core.Int64Ptr(int64(10))
				volumeCreateOptionsModel.Name = core.StringPtr("my-volume")
				volumeCreateOptionsModel.Hostnqnstring = core.StringPtr("nqn.2024-07.org:1234")
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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
				volumePatchModel.Capacity = core.Int64Ptr(int64(38))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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
				volumePatchModel.Capacity = core.Int64Ptr(int64(38))
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
					fmt.Fprintf(res, "%s", `{"bandwidth": 1000, "capacity": 8, "created_at": "CreatedAt", "hosts": [{"host_id": "HostID", "host_name": "HostName", "host_nqn": "HostNqn"}], "id": "ID", "iops": 10000, "name": "Name", "resource_type": "ResourceType", "status": "Status", "status_reasons": ["StatusReasons"]}`)
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
				volumePatchModel.Capacity = core.Int64Ptr(int64(38))
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
				volumePatchModel.Capacity = core.Int64Ptr(int64(38))
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
				volumePatchModel.Capacity = core.Int64Ptr(int64(38))
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
	Describe(`Creds(credsOptions *CredsOptions) - Operation response error`, func() {
		credsPath := "/v1/object/workspace/credentials"
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
		credsPath := "/v1/object/workspace/credentials"
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
					fmt.Fprintf(res, "%s", `{"access_keys": ["AccessKeys"]}`)
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
					fmt.Fprintf(res, "%s", `{"access_keys": ["AccessKeys"]}`)
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
		credCreatePath := "/v1/object/workspace/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credCreatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["access_key"]).To(Equal([]string{"mytestkey"}))
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
		credCreatePath := "/v1/object/workspace/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credCreatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["access_key"]).To(Equal([]string{"mytestkey"}))
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

					Expect(req.URL.Query()["access_key"]).To(Equal([]string{"mytestkey"}))
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
		credDeletePath := "/v1/object/workspace/credentials"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(credDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["access_key"]).To(Equal([]string{"mytestkey"}))
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
	Describe(`Cert(certOptions *CertOptions) - Operation response error`, func() {
		certPath := "/v1/object/certificate/s3"
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
		certPath := "/v1/object/certificate/s3"
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
					fmt.Fprintf(res, "%s", `{"expiration_date": "ExpirationDate", "expired": false}`)
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
					fmt.Fprintf(res, "%s", `{"expiration_date": "ExpirationDate", "expired": false}`)
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
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.Cert(certOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Cert with error: Operation request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
				certOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.Cert(certOptionsModel)
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
			It(`Invoke Cert successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertOptions model
				certOptionsModel := new(sdsaasv1.CertOptions)
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
	Describe(`CertUpload(certUploadOptions *CertUploadOptions) - Operation response error`, func() {
		certUploadPath := "/v1/object/certificate/s3"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certUploadPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CertUpload with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUploadOptions model
				certUploadOptionsModel := new(sdsaasv1.CertUploadOptions)
				certUploadOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.CertUpload(certUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.CertUpload(certUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CertUpload(certUploadOptions *CertUploadOptions)`, func() {
		certUploadPath := "/v1/object/certificate/s3"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(certUploadPath))
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
					fmt.Fprintf(res, "%s", `{"error": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertUpload successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the CertUploadOptions model
				certUploadOptionsModel := new(sdsaasv1.CertUploadOptions)
				certUploadOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.CertUploadWithContext(ctx, certUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.CertUpload(certUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.CertUploadWithContext(ctx, certUploadOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(certUploadPath))
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
					fmt.Fprintf(res, "%s", `{"error": [{"mapKey": "Inner"}], "valid_certificate": true, "valid_key": true}`)
				}))
			})
			It(`Invoke CertUpload successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.CertUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertUploadOptions model
				certUploadOptionsModel := new(sdsaasv1.CertUploadOptions)
				certUploadOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.CertUpload(certUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CertUpload with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUploadOptions model
				certUploadOptionsModel := new(sdsaasv1.CertUploadOptions)
				certUploadOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.CertUpload(certUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CertUploadOptions model with no property values
				certUploadOptionsModelNew := new(sdsaasv1.CertUploadOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.CertUpload(certUploadOptionsModelNew)
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
			It(`Invoke CertUpload successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the CertUploadOptions model
				certUploadOptionsModel := new(sdsaasv1.CertUploadOptions)
				certUploadOptionsModel.Body = CreateMockReader("This is a mock file.")
				certUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.CertUpload(certUploadOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
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
				hostsOptionsModel.Limit = core.Int64Ptr(int64(10))
				hostsOptionsModel.Name = core.StringPtr("myhost1")
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}], "limit": 5, "next": {"href": "Href"}, "total_count": 10}`)
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
				hostsOptionsModel.Limit = core.Int64Ptr(int64(10))
				hostsOptionsModel.Name = core.StringPtr("myhost1")
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"myhost1"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "hosts": [{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}], "limit": 5, "next": {"href": "Href"}, "total_count": 10}`)
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
				hostsOptionsModel.Limit = core.Int64Ptr(int64(10))
				hostsOptionsModel.Name = core.StringPtr("myhost1")
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
				hostsOptionsModel.Limit = core.Int64Ptr(int64(10))
				hostsOptionsModel.Name = core.StringPtr("myhost1")
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
				hostsOptionsModel.Limit = core.Int64Ptr(int64(10))
				hostsOptionsModel.Name = core.StringPtr("myhost1")
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

				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.Volumes = []sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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

				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.Volumes = []sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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

				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.Volumes = []sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}
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

				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.Volumes = []sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}
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

				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsModel := new(sdsaasv1.HostCreateOptions)
				hostCreateOptionsModel.Nqn = core.StringPtr("nqn.2014-06.org:9345")
				hostCreateOptionsModel.Name = core.StringPtr("my-host")
				hostCreateOptionsModel.Volumes = []sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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
				hostPatchModel.Name = core.StringPtr("mytesthost")
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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
				hostPatchModel.Name = core.StringPtr("mytesthost")
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
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
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
				hostPatchModel.Name = core.StringPtr("mytesthost")
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
				hostPatchModel.Name = core.StringPtr("mytesthost")
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
				hostPatchModel.Name = core.StringPtr("mytesthost")
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
	Describe(`HostVolDeleteall(hostVolDeleteallOptions *HostVolDeleteallOptions)`, func() {
		hostVolDeleteallPath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volumes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostVolDeleteallPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke HostVolDeleteall successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.HostVolDeleteall(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HostVolDeleteallOptions model
				hostVolDeleteallOptionsModel := new(sdsaasv1.HostVolDeleteallOptions)
				hostVolDeleteallOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.HostVolDeleteall(hostVolDeleteallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HostVolDeleteall with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostVolDeleteallOptions model
				hostVolDeleteallOptionsModel := new(sdsaasv1.HostVolDeleteallOptions)
				hostVolDeleteallOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.HostVolDeleteall(hostVolDeleteallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HostVolDeleteallOptions model with no property values
				hostVolDeleteallOptionsModelNew := new(sdsaasv1.HostVolDeleteallOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.HostVolDeleteall(hostVolDeleteallOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostVolDelete(hostVolDeleteOptions *HostVolDeleteOptions)`, func() {
		hostVolDeletePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostVolDeletePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke HostVolDelete successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := sdsaasService.HostVolDelete(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HostVolDeleteOptions model
				hostVolDeleteOptionsModel := new(sdsaasv1.HostVolDeleteOptions)
				hostVolDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = sdsaasService.HostVolDelete(hostVolDeleteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HostVolDelete with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostVolDeleteOptions model
				hostVolDeleteOptionsModel := new(sdsaasv1.HostVolDeleteOptions)
				hostVolDeleteOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolDeleteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := sdsaasService.HostVolDelete(hostVolDeleteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HostVolDeleteOptions model with no property values
				hostVolDeleteOptionsModelNew := new(sdsaasv1.HostVolDeleteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = sdsaasService.HostVolDelete(hostVolDeleteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostVolUpdate(hostVolUpdateOptions *HostVolUpdateOptions) - Operation response error`, func() {
		hostVolUpdatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostVolUpdatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke HostVolUpdate with error: Operation response processing error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostVolUpdateOptions model
				hostVolUpdateOptionsModel := new(sdsaasv1.HostVolUpdateOptions)
				hostVolUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sdsaasService.EnableRetries(0, 0)
				result, response, operationErr = sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`HostVolUpdate(hostVolUpdateOptions *HostVolUpdateOptions)`, func() {
		hostVolUpdatePath := "/hosts/r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e/volumes/r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(hostVolUpdatePath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
				}))
			})
			It(`Invoke HostVolUpdate successfully with retries`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())
				sdsaasService.EnableRetries(0, 0)

				// Construct an instance of the HostVolUpdateOptions model
				hostVolUpdateOptionsModel := new(sdsaasv1.HostVolUpdateOptions)
				hostVolUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sdsaasService.HostVolUpdateWithContext(ctx, hostVolUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sdsaasService.DisableRetries()
				result, response, operationErr := sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sdsaasService.HostVolUpdateWithContext(ctx, hostVolUpdateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(hostVolUpdatePath))
					Expect(req.Method).To(Equal("PUT"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"created_at": "CreatedAt", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-host", "nqn": "nqn-abc-1234", "volumes": [{"status": "Status", "volume_id": "VolumeID", "volume_name": "VolumeName", "storage_identifiers": {"id": "ID", "namespace_id": 11, "namespace_uuid": "NamespaceUUID", "network_info": [{"gateway_ip": "GatewayIP", "port": 4}]}}]}`)
				}))
			})
			It(`Invoke HostVolUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sdsaasService.HostVolUpdate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostVolUpdateOptions model
				hostVolUpdateOptionsModel := new(sdsaasv1.HostVolUpdateOptions)
				hostVolUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke HostVolUpdate with error: Operation validation and request error`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostVolUpdateOptions model
				hostVolUpdateOptionsModel := new(sdsaasv1.HostVolUpdateOptions)
				hostVolUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sdsaasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the HostVolUpdateOptions model with no property values
				hostVolUpdateOptionsModelNew := new(sdsaasv1.HostVolUpdateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sdsaasService.HostVolUpdate(hostVolUpdateOptionsModelNew)
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
			It(`Invoke HostVolUpdate successfully`, func() {
				sdsaasService, serviceErr := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sdsaasService).ToNot(BeNil())

				// Construct an instance of the HostVolUpdateOptions model
				hostVolUpdateOptionsModel := new(sdsaasv1.HostVolUpdateOptions)
				hostVolUpdateOptionsModel.HostID = core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.VolumeID = core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sdsaasService.HostVolUpdate(hostVolUpdateOptionsModel)
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
			sdsaasService, _ := sdsaasv1.NewSdsaasV1(&sdsaasv1.SdsaasV1Options{
				URL:           "http://sdsaasv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCertOptions successfully`, func() {
				// Construct an instance of the CertOptions model
				certOptionsModel := sdsaasService.NewCertOptions()
				certOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certOptionsModel).ToNot(BeNil())
				Expect(certOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCertUploadOptions successfully`, func() {
				// Construct an instance of the CertUploadOptions model
				body := CreateMockReader("This is a mock file.")
				certUploadOptionsModel := sdsaasService.NewCertUploadOptions(body)
				certUploadOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				certUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(certUploadOptionsModel).ToNot(BeNil())
				Expect(certUploadOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(certUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				// Construct an instance of the VolumeMappingIdentity model
				volumeMappingIdentityModel := new(sdsaasv1.VolumeMappingIdentity)
				Expect(volumeMappingIdentityModel).ToNot(BeNil())
				volumeMappingIdentityModel.VolumeID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")
				Expect(volumeMappingIdentityModel.VolumeID).To(Equal(core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")))

				// Construct an instance of the HostCreateOptions model
				hostCreateOptionsNqn := "nqn.2014-06.org:9345"
				hostCreateOptionsModel := sdsaasService.NewHostCreateOptions(hostCreateOptionsNqn)
				hostCreateOptionsModel.SetNqn("nqn.2014-06.org:9345")
				hostCreateOptionsModel.SetName("my-host")
				hostCreateOptionsModel.SetVolumes([]sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel})
				hostCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostCreateOptionsModel).ToNot(BeNil())
				Expect(hostCreateOptionsModel.Nqn).To(Equal(core.StringPtr("nqn.2014-06.org:9345")))
				Expect(hostCreateOptionsModel.Name).To(Equal(core.StringPtr("my-host")))
				Expect(hostCreateOptionsModel.Volumes).To(Equal([]sdsaasv1.VolumeMappingIdentity{*volumeMappingIdentityModel}))
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
			It(`Invoke NewHostVolDeleteOptions successfully`, func() {
				// Construct an instance of the HostVolDeleteOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				hostVolDeleteOptionsModel := sdsaasService.NewHostVolDeleteOptions(hostID, volumeID)
				hostVolDeleteOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteOptionsModel.SetVolumeID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolDeleteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostVolDeleteOptionsModel).ToNot(BeNil())
				Expect(hostVolDeleteOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostVolDeleteOptionsModel.VolumeID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(hostVolDeleteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostVolDeleteallOptions successfully`, func() {
				// Construct an instance of the HostVolDeleteallOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				hostVolDeleteallOptionsModel := sdsaasService.NewHostVolDeleteallOptions(hostID)
				hostVolDeleteallOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolDeleteallOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostVolDeleteallOptionsModel).ToNot(BeNil())
				Expect(hostVolDeleteallOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostVolDeleteallOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostVolUpdateOptions successfully`, func() {
				// Construct an instance of the HostVolUpdateOptions model
				hostID := "r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e"
				volumeID := "r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39"
				hostVolUpdateOptionsModel := sdsaasService.NewHostVolUpdateOptions(hostID, volumeID)
				hostVolUpdateOptionsModel.SetHostID("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")
				hostVolUpdateOptionsModel.SetVolumeID("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")
				hostVolUpdateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostVolUpdateOptionsModel).ToNot(BeNil())
				Expect(hostVolUpdateOptionsModel.HostID).To(Equal(core.StringPtr("r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e")))
				Expect(hostVolUpdateOptionsModel.VolumeID).To(Equal(core.StringPtr("r134-f24710c4-d5f4-4881-ab78-7bfXX6281f39")))
				Expect(hostVolUpdateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostsOptions successfully`, func() {
				// Construct an instance of the HostsOptions model
				hostsOptionsModel := sdsaasService.NewHostsOptions()
				hostsOptionsModel.SetLimit(int64(10))
				hostsOptionsModel.SetName("myhost1")
				hostsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(hostsOptionsModel).ToNot(BeNil())
				Expect(hostsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(hostsOptionsModel.Name).To(Equal(core.StringPtr("myhost1")))
				Expect(hostsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeCreateOptions successfully`, func() {
				// Construct an instance of the VolumeCreateOptions model
				volumeCreateOptionsCapacity := int64(10)
				volumeCreateOptionsModel := sdsaasService.NewVolumeCreateOptions(volumeCreateOptionsCapacity)
				volumeCreateOptionsModel.SetCapacity(int64(10))
				volumeCreateOptionsModel.SetName("my-volume")
				volumeCreateOptionsModel.SetHostnqnstring("nqn.2024-07.org:1234")
				volumeCreateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumeCreateOptionsModel).ToNot(BeNil())
				Expect(volumeCreateOptionsModel.Capacity).To(Equal(core.Int64Ptr(int64(10))))
				Expect(volumeCreateOptionsModel.Name).To(Equal(core.StringPtr("my-volume")))
				Expect(volumeCreateOptionsModel.Hostnqnstring).To(Equal(core.StringPtr("nqn.2024-07.org:1234")))
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
			It(`Invoke NewVolumeMappingIdentity successfully`, func() {
				volumeID := "testString"
				_model, err := sdsaasService.NewVolumeMappingIdentity(volumeID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
				volumesOptionsModel.SetLimit(int64(10))
				volumesOptionsModel.SetName("myhost1")
				volumesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(volumesOptionsModel).ToNot(BeNil())
				Expect(volumesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(volumesOptionsModel.Name).To(Equal(core.StringPtr("myhost1")))
				Expect(volumesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalHostPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.HostPatch)
			model.Name = core.StringPtr("mytesthost")

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
		It(`Invoke UnmarshalVolumeMappingIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.VolumeMappingIdentity)
			model.VolumeID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *sdsaasv1.VolumeMappingIdentity
			err = sdsaasv1.UnmarshalVolumeMappingIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVolumePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(sdsaasv1.VolumePatch)
			model.Capacity = core.Int64Ptr(int64(38))
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
