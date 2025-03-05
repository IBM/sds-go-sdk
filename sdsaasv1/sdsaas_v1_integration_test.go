//go:build integration

/**
 * (C) Copyright IBM Corp. 2024, 2025.
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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the sdsaasv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SdsaasV1 Integration Tests`, func() {
	const externalConfigFile = "../sdsaas_v1.env"

	var (
		err           error
		sdsaasService *sdsaasv1.SdsaasV1
		serviceURL    string
		config        map[string]string

		// Variables to hold link values
		hostIDLink             string
		volumeIDLink           string
		volumeIDLinkTwo        string
		volumeMappingIDLinkOne string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(sdsaasv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			sdsaasServiceOptions := &sdsaasv1.SdsaasV1Options{}

			sdsaasService, err = sdsaasv1.NewSdsaasV1UsingExternalConfig(sdsaasServiceOptions)
			Expect(err).To(BeNil())
			Expect(sdsaasService).ToNot(BeNil())
			Expect(sdsaasService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			sdsaasService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`VolumeCreate - Create a new volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeCreate(volumeCreateOptions *VolumeCreateOptions)`, func() {
			volumeCreateOptions := &sdsaasv1.VolumeCreateOptions{
				Capacity: core.Int64Ptr(int64(10)),
				Name:     core.StringPtr("my-volume-one"),
			}

			volumeCreateOptionsTwo := &sdsaasv1.VolumeCreateOptions{
				Capacity: core.Int64Ptr(int64(10)),
				Name:     core.StringPtr("my-volume-two"),
			}

			volume, response, err := sdsaasService.VolumeCreate(volumeCreateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volume).ToNot(BeNil())

			volumeTwo, responseTwo, errTwo := sdsaasService.VolumeCreate(volumeCreateOptionsTwo)
			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(201))
			Expect(volumeTwo).ToNot(BeNil())

			volumeIDLink = *volume.ID
			volumeIDLinkTwo = *volumeTwo.ID

			fmt.Fprintf(GinkgoWriter, "Saved volumeIDLink value: %v\n", volumeIDLink)

			time.Sleep(5 * time.Second)
		})
	})

	Describe(`HostCreate - Creates a host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostCreate(hostCreateOptions *HostCreateOptions)`, func() {
			volumeIdentityModel := &sdsaasv1.VolumeIdentity{
				ID: core.StringPtr(volumeIDLink),
			}

			volumeMappingPrototypeModel := &sdsaasv1.VolumeMappingPrototype{
				Volume: volumeIdentityModel,
			}

			hostCreateOptions := &sdsaasv1.HostCreateOptions{
				Name:           core.StringPtr("my-host"),
				Nqn:            core.StringPtr("nqn.2014-06.org:9345"),
				VolumeMappings: []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel},
			}

			host, response, err := sdsaasService.HostCreate(hostCreateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(host).ToNot(BeNil())

			hostIDLink = *host.ID
			volumeMappingIDLinkOne = *host.VolumeMappings[0].ID

			fmt.Fprintf(GinkgoWriter, "Saved hostIDLink value: %v\n", hostIDLink)

			time.Sleep(5 * time.Second)
		})
	})

	Describe(`Volumes - This request lists all volumes in the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Volumes(volumesOptions *VolumesOptions)`, func() {
			volumesOptions := &sdsaasv1.VolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume"),
			}

			volumeCollection, response, err := sdsaasService.Volumes(volumesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeCollection).ToNot(BeNil())
		})
	})

	Describe(`Volume - Retrieve a volume profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Volume(volumeOptions *VolumeOptions)`, func() {
			volumeOptions := &sdsaasv1.VolumeOptions{
				VolumeID: &volumeIDLink,
			}

			volume, response, err := sdsaasService.Volume(volumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`VolumeUpdate - Update a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeUpdate(volumeUpdateOptions *VolumeUpdateOptions)`, func() {
			volumePatchModel := &sdsaasv1.VolumePatch{
				Name: core.StringPtr("my-volume-one-updated"),
			}
			volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			volumeUpdateOptions := &sdsaasv1.VolumeUpdateOptions{
				VolumeID:    &volumeIDLink,
				VolumePatch: volumePatchModelAsPatch,
			}

			volume, response, err := sdsaasService.VolumeUpdate(volumeUpdateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

			time.Sleep(5 * time.Second)
		})
	})

	Describe(`Creds - List storage account credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Creds(credsOptions *CredsOptions)`, func() {
			credsOptions := &sdsaasv1.CredsOptions{}

			credentialsFound, response, err := sdsaasService.Creds(credsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialsFound).ToNot(BeNil())
		})
	})

	Describe(`CredCreate - Create or modify storage account credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CredCreate(credCreateOptions *CredCreateOptions)`, func() {
			credCreateOptions := &sdsaasv1.CredCreateOptions{
				AccessKey: core.StringPtr("mytestkey"),
			}

			credentialsUpdated, response, err := sdsaasService.CredCreate(credCreateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credentialsUpdated).ToNot(BeNil())

			time.Sleep(5 * time.Second)
		})
	})

	Describe(`CertTypes - List the allowed certificate types`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CertTypes(certTypesOptions *CertTypesOptions)`, func() {
			certTypesOptions := &sdsaasv1.CertTypesOptions{}

			certificateList, response, err := sdsaasService.CertTypes(certTypesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certificateList).ToNot(BeNil())
		})
	})

	Describe(`Cert - Retrieves the SSL certificate expiration date and status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Cert(certOptions *CertOptions)`, func() {
			certOptions := &sdsaasv1.CertOptions{
				Cert: core.StringPtr("s3"),
			}

			certificateFound, response, err := sdsaasService.Cert(certOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certificateFound).ToNot(BeNil())
		})
	})
	// Describe(`CertCreate - Creates a new SSL Certificate`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`CertCreate(certCreateOptions *CertCreateOptions)`, func() {
	// 		certCreateOptions := &sdsaasv1.CertCreateOptions{
	// 			Cert: core.StringPtr("s3"),
	// 			Body: CreateMockReader("This is a mock file."),
	// 		}

	// 		certificateUpdated, response, err := sdsaasService.CertCreate(certCreateOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(202))
	// 		Expect(certificateUpdated).ToNot(BeNil())
	// 	})
	// })

	// Describe(`CertUpdate - Updates the SSL Certificate`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`CertUpdate(certUpdateOptions *CertUpdateOptions)`, func() {
	// 		certUpdateOptions := &sdsaasv1.CertUpdateOptions{
	// 			Cert: core.StringPtr("s3"),
	// 			Body: CreateMockReader("This is a mock file."),
	// 		}

	// 		certificateUpdated, response, err := sdsaasService.CertUpdate(certUpdateOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(202))
	// 		Expect(certificateUpdated).ToNot(BeNil())
	// 	})
	// })

	Describe(`Hosts - Lists all hosts and all host IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Hosts(hostsOptions *HostsOptions)`, func() {
			hostsOptions := &sdsaasv1.HostsOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-host"),
			}

			hostCollection, response, err := sdsaasService.Hosts(hostsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hostCollection).ToNot(BeNil())
		})
	})

	Describe(`Host - Retrieve a host by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Host(hostOptions *HostOptions)`, func() {
			hostOptions := &sdsaasv1.HostOptions{
				HostID: &hostIDLink,
			}

			host, response, err := sdsaasService.Host(hostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
	})

	Describe(`HostUpdate - Update a host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostUpdate(hostUpdateOptions *HostUpdateOptions)`, func() {
			hostPatchModel := &sdsaasv1.HostPatch{
				Name: core.StringPtr("my-host-updated"),
			}
			hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			hostUpdateOptions := &sdsaasv1.HostUpdateOptions{
				HostID:    &hostIDLink,
				HostPatch: hostPatchModelAsPatch,
			}

			host, response, err := sdsaasService.HostUpdate(hostUpdateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())

			time.Sleep(8 * time.Second)
		})
	})

	Describe(`HostMappings - List all volume mappings for a host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostMappings(hostMappingsOptions *HostMappingsOptions)`, func() {
			hostMappingsOptions := &sdsaasv1.HostMappingsOptions{
				HostID: &hostIDLink,
			}

			volumeMappingCollection, response, err := sdsaasService.HostMappings(hostMappingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMappingCollection).ToNot(BeNil())
		})
	})

	Describe(`HostMappingCreate - Create a Volume mapping for a host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostMappingCreate(hostMappingCreateOptions *HostMappingCreateOptions)`, func() {
			volumeIdentityModel := &sdsaasv1.VolumeIdentity{
				ID: core.StringPtr(volumeIDLinkTwo),
			}

			hostMappingCreateOptions := &sdsaasv1.HostMappingCreateOptions{
				HostID: &hostIDLink,
				Volume: volumeIdentityModel,
			}

			time.Sleep(5 * time.Second)

			volumeMapping, response, err := sdsaasService.HostMappingCreate(hostMappingCreateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeMapping).ToNot(BeNil())
		})
	})

	Describe(`HostMapping - Retrieve a volume mapping`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostMapping(hostMappingOptions *HostMappingOptions)`, func() {
			hostMappingOptions := &sdsaasv1.HostMappingOptions{
				HostID:          &hostIDLink,
				VolumeMappingID: core.StringPtr(volumeMappingIDLinkOne),
			}

			volumeMapping, response, err := sdsaasService.HostMapping(hostMappingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMapping).ToNot(BeNil())
		})
	})

	Describe(`HostMappingDelete - Deletes the given volume mapping for a specific host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostMappingDelete(hostMappingDeleteOptions *HostMappingDeleteOptions)`, func() {
			hostMappingDeleteOptions := &sdsaasv1.HostMappingDeleteOptions{
				HostID:          &hostIDLink,
				VolumeMappingID: core.StringPtr(volumeMappingIDLinkOne),
			}

			response, err := sdsaasService.HostMappingDelete(hostMappingDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`HostMappingDeleteAll - Deletes all the volume mappings for a given host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostMappingDeleteAll(hostMappingDeleteAllOptions *HostMappingDeleteAllOptions)`, func() {
			hostMappingDeleteAllOptions := &sdsaasv1.HostMappingDeleteAllOptions{
				HostID: &hostIDLink,
			}

			response, err := sdsaasService.HostMappingDeleteAll(hostMappingDeleteAllOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			time.Sleep(10 * time.Second)
		})
	})

	Describe(`HostDelete - Delete a specific host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HostDelete(hostDeleteOptions *HostDeleteOptions)`, func() {
			hostDeleteOptions := &sdsaasv1.HostDeleteOptions{
				HostID: &hostIDLink,
			}

			response, err := sdsaasService.HostDelete(hostDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			time.Sleep(5 * time.Second)
		})
	})

	Describe(`VolumeDelete - Delete a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeDelete(volumeDeleteOptions *VolumeDeleteOptions)`, func() {
			volumeDeleteOptions := &sdsaasv1.VolumeDeleteOptions{
				VolumeID: &volumeIDLink,
			}

			response, err := sdsaasService.VolumeDelete(volumeDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			volumeDeleteOptionsTwo := &sdsaasv1.VolumeDeleteOptions{
				VolumeID: &volumeIDLinkTwo,
			}

			response, err = sdsaasService.VolumeDelete(volumeDeleteOptionsTwo)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CredDelete - Delete storage account credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CredDelete(credDeleteOptions *CredDeleteOptions)`, func() {
			credDeleteOptions := &sdsaasv1.CredDeleteOptions{
				AccessKey: core.StringPtr("mytestkey"),
			}

			response, err := sdsaasService.CredDelete(credDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	// Describe(`CertDelete - Delete SSL certificate`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`CertDelete(certDeleteOptions *CertDeleteOptions)`, func() {
	// 		certDeleteOptions := &sdsaasv1.CertDeleteOptions{
	// 			Cert: core.StringPtr("s3"),
	// 		}

	// 		response, err := sdsaasService.CertDelete(certDeleteOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(204))
	// 	})
	// })
})

//
// Utility functions are declared in the unit test file
//
