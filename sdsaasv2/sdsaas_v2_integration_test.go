//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/v2/sdsaasv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the sdsaasv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SdsaasV2 Integration Tests`, func() {
	const externalConfigFile = "../sdsaas_v2.env"

	var (
		err           error
		sdsaasService *sdsaasv2.SdsaasV2
		serviceURL    string
		config        map[string]string

		// Variables to hold link values
		hostIDLink   string
		volumeIDLink string
		snapIDLink   string

		volumeMappingIDLink string
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
			config, err = core.GetServiceProperties(sdsaasv2.DefaultServiceName)
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
			sdsaasServiceOptions := &sdsaasv2.SdsaasV2Options{}

			sdsaasService, err = sdsaasv2.NewSdsaasV2UsingExternalConfig(sdsaasServiceOptions)
			Expect(err).To(BeNil())
			Expect(sdsaasService).ToNot(BeNil())
			Expect(sdsaasService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			sdsaasService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateVolume - Create a Volume`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVolume(createVolumeOptions *CreateVolumeOptions)`, func() {

			createVolumeOptions := &sdsaasv2.CreateVolumeOptions{
				Capacity: core.Int64Ptr(int64(1)),
				Name:     core.StringPtr("my-volume"),
			}

			volumeSummary, response, err := sdsaasService.CreateVolume(createVolumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeSummary).ToNot(BeNil())

			volumeIDLink = *volumeSummary.ID
			fmt.Fprintf(GinkgoWriter, "Saved volumeIDLink value: %v\n", volumeIDLink)

			// Wait for resource to create
			time.Sleep(5 * time.Second)
		})
	})

	Describe(`CreateHost - Create a host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateHost(createHostOptions *CreateHostOptions)`, func() {
			createHostOptions := &sdsaasv2.CreateHostOptions{
				Nqn:  core.StringPtr("nqn.2014-08.org.nvmexpress:uuid:12345678-4444-1234-1234-123456789abd"),
				Name: core.StringPtr("my-host"),
			}

			hostSummary, response, err := sdsaasService.CreateHost(createHostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hostSummary).ToNot(BeNil())

			hostIDLink = *hostSummary.ID
			fmt.Fprintf(GinkgoWriter, "Saved hostIDLink value: %v\n", hostIDLink)
		})
	})

	Describe(`ListVolumes - List all Volumes`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVolumes(listVolumesOptions *ListVolumesOptions) with pagination`, func() {
			listVolumesOptions := &sdsaasv2.ListVolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume"),
			}

			listVolumesOptions.Start = nil
			listVolumesOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv2.Volume
			for {
				volumeCollection, response, err := sdsaasService.ListVolumes(listVolumesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(volumeCollection).ToNot(BeNil())
				allResults = append(allResults, volumeCollection.Volumes...)

				listVolumesOptions.Start, err = volumeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listVolumesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListVolumes(listVolumesOptions *ListVolumesOptions) using VolumesPager`, func() {
			listVolumesOptions := &sdsaasv2.ListVolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume"),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewVolumesPager(listVolumesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv2.Volume
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewVolumesPager(listVolumesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListVolumes() returned a total of %d item(s) using VolumesPager.\n", len(allResults))
		})
	})

	Describe(`GetVolume - Retrieve a volume`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVolume(getVolumeOptions *GetVolumeOptions)`, func() {
			getVolumeOptions := &sdsaasv2.GetVolumeOptions{
				ID: &volumeIDLink,
			}

			volume, response, err := sdsaasService.GetVolume(getVolumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`UpdateVolume - Update a volume`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVolume(updateVolumeOptions *UpdateVolumeOptions)`, func() {
			volumePatchModel := &sdsaasv2.VolumePatch{
				Name: core.StringPtr("my-volume-updated"),
			}
			volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVolumeOptions := &sdsaasv2.UpdateVolumeOptions{
				ID:          &volumeIDLink,
				VolumePatch: volumePatchModelAsPatch,
			}

			volume, response, err := sdsaasService.UpdateVolume(updateVolumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`ListHosts - List all hosts`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListHosts(listHostsOptions *ListHostsOptions) with pagination`, func() {
			listHostsOptions := &sdsaasv2.ListHostsOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-host"),
			}

			listHostsOptions.Start = nil
			listHostsOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv2.Host
			for {
				hostCollection, response, err := sdsaasService.ListHosts(listHostsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(hostCollection).ToNot(BeNil())
				allResults = append(allResults, hostCollection.Hosts...)

				listHostsOptions.Start, err = hostCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listHostsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListHosts(listHostsOptions *ListHostsOptions) using HostsPager`, func() {
			listHostsOptions := &sdsaasv2.ListHostsOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-host"),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewHostsPager(listHostsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv2.Host
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewHostsPager(listHostsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListHosts() returned a total of %d item(s) using HostsPager.\n", len(allResults))
		})
	})

	Describe(`GetHost - Retrieve a Host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetHost(getHostOptions *GetHostOptions)`, func() {
			getHostOptions := &sdsaasv2.GetHostOptions{
				ID: &hostIDLink,
			}

			host, response, err := sdsaasService.GetHost(getHostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
	})

	Describe(`UpdateHost - Update a given Host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateHost(updateHostOptions *UpdateHostOptions)`, func() {
			hostPatchModel := &sdsaasv2.HostPatch{
				Name: core.StringPtr("my-resource"),
			}
			hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateHostOptions := &sdsaasv2.UpdateHostOptions{
				ID:        &hostIDLink,
				HostPatch: hostPatchModelAsPatch,
			}

			host, response, err := sdsaasService.UpdateHost(updateHostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
	})

	Describe(`CreateVolumeMapping - Create a Volume mapping for a host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVolumeMapping(createVolumeMappingOptions *CreateVolumeMappingOptions)`, func() {
			volumeIdentityModel := &sdsaasv2.VolumeIdentity{
				ID: core.StringPtr(volumeIDLink),
			}

			createVolumeMappingOptions := &sdsaasv2.CreateVolumeMappingOptions{
				ID:     core.StringPtr(hostIDLink),
				Volume: volumeIdentityModel,
			}

			volumeMappingReference, response, err := sdsaasService.CreateVolumeMapping(createVolumeMappingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeMappingReference).ToNot(BeNil())

			volumeMappingIDLink = *volumeMappingReference.ID
			fmt.Fprintf(GinkgoWriter, "Saved volumeMappingIDLink value: %v\n", volumeMappingIDLink)

			// Wait for resource to create
			time.Sleep(5 * time.Second)
		})
	})

	Describe(`ListVolumeMappings - List all volume mappings for a host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVolumeMappings(listVolumeMappingsOptions *ListVolumeMappingsOptions) with pagination`, func() {
			listVolumeMappingsOptions := &sdsaasv2.ListVolumeMappingsOptions{
				ID:    core.StringPtr(hostIDLink),
				Limit: core.Int64Ptr(int64(20)),
			}

			listVolumeMappingsOptions.Start = nil
			listVolumeMappingsOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv2.VolumeMapping
			for {
				volumeMappingCollection, response, err := sdsaasService.ListVolumeMappings(listVolumeMappingsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(volumeMappingCollection).ToNot(BeNil())
				allResults = append(allResults, volumeMappingCollection.VolumeMappings...)

				listVolumeMappingsOptions.Start, err = volumeMappingCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listVolumeMappingsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListVolumeMappings(listVolumeMappingsOptions *ListVolumeMappingsOptions) using VolumeMappingsPager`, func() {
			listVolumeMappingsOptions := &sdsaasv2.ListVolumeMappingsOptions{
				ID:    core.StringPtr(hostIDLink),
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-resource"),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewVolumeMappingsPager(listVolumeMappingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv2.VolumeMapping
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewVolumeMappingsPager(listVolumeMappingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListVolumeMappings() returned a total of %d item(s) using VolumeMappingsPager.\n", len(allResults))
		})
	})

	Describe(`GetVolumeMapping - Retrieve a volume mapping`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVolumeMapping(getVolumeMappingOptions *GetVolumeMappingOptions)`, func() {
			getVolumeMappingOptions := &sdsaasv2.GetVolumeMappingOptions{
				ID:              core.StringPtr(hostIDLink),
				VolumeMappingID: core.StringPtr(volumeMappingIDLink),
			}

			volumeMapping, response, err := sdsaasService.GetVolumeMapping(getVolumeMappingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMapping).ToNot(BeNil())
		})
	})

	Describe(`CreateHmacCredentials - Create HMAC credentials`, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateHmacCredentials(createHmacCredentialsOptions *CreateHmacCredentialsOptions)`, func() {
			createHmacCredentialsOptions := &sdsaasv2.CreateHmacCredentialsOptions{
				AccessKey: core.StringPtr("mytestkey"),
			}

			accessKeyResponse, response, err := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`ListHmacCredentials - List HMAC credentials`, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListHmacCredentials(listHmacCredentialsOptions *ListHmacCredentialsOptions)`, func() {
			listHmacCredentialsOptions := &sdsaasv2.ListHmacCredentialsOptions{}

			storageCredResponse, response, err := sdsaasService.ListHmacCredentials(listHmacCredentialsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageCredResponse).ToNot(BeNil())
		})
	})

	Describe(`ListCertificates - List the configured certificates`, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCertificates(listCertificatesOptions *ListCertificatesOptions)`, func() {
			listCertificatesOptions := &sdsaasv2.ListCertificatesOptions{}

			certListResponse, response, err := sdsaasService.ListCertificates(listCertificatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certListResponse).ToNot(BeNil())
		})
	})

	Describe(`GetS3SslCertStatus - Get SSL cert status`, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetS3SslCertStatus(getS3SslCertStatusOptions *GetS3SslCertStatusOptions)`, func() {
			getS3SslCertStatusOptions := &sdsaasv2.GetS3SslCertStatusOptions{
				CertType: core.StringPtr("s3"),
			}

			statusResponse, response, err := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statusResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSslCert - Create SSL certificates`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSslCert(createSslCertOptions *CreateSslCertOptions)`, func() {
			createSslCertOptions := &sdsaasv2.CreateSslCertOptions{
				CertType: core.StringPtr("s3"),
				Body:     CreateMockReader("This is a mock file."),
			}

			certResponse, response, err := sdsaasService.CreateSslCert(createSslCertOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSslCert - Replace an SSL certificate`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSslCert(replaceSslCertOptions *ReplaceSslCertOptions)`, func() {
			replaceSslCertOptions := &sdsaasv2.ReplaceSslCertOptions{
				CertType: core.StringPtr("s3"),
				Body:     CreateMockReader("This is a mock file."),
			}

			certResponse, response, err := sdsaasService.ReplaceSslCert(replaceSslCertOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSnapshot - Create a Snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSnapshot(createSnapshotOptions *CreateSnapshotOptions)`, func() {
			sourceVolumePrototypeModel := &sdsaasv2.SourceVolumePrototype{
				ID: core.StringPtr(volumeIDLink),
			}

			createSnapshotOptions := &sdsaasv2.CreateSnapshotOptions{
				Name:         core.StringPtr("my-snapshot"),
				SourceVolume: sourceVolumePrototypeModel,
			}

			snapshot, response, err := sdsaasService.CreateSnapshot(createSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())

			snapIDLink = *snapshot.ID
			fmt.Fprintf(GinkgoWriter, "Saved snapIDLink value: %v\n", snapIDLink)

			// Wait for snapshot to fully create
			time.Sleep(5 * time.Second)
		})
	})

	Describe(`ListSnapshots - List all Snapshots`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) with pagination`, func() {
			listSnapshotsOptions := &sdsaasv2.ListSnapshotsOptions{
				Limit: core.Int64Ptr(int64(20)),
			}

			listSnapshotsOptions.Start = nil
			listSnapshotsOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv2.Snapshot
			for {
				snapshotCollection, response, err := sdsaasService.ListSnapshots(listSnapshotsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(snapshotCollection).ToNot(BeNil())
				allResults = append(allResults, snapshotCollection.Snapshots...)

				listSnapshotsOptions.Start, err = snapshotCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listSnapshotsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) using SnapshotsPager`, func() {
			listSnapshotsOptions := &sdsaasv2.ListSnapshotsOptions{
				Limit: core.Int64Ptr(int64(20)),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewSnapshotsPager(listSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv2.Snapshot
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewSnapshotsPager(listSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSnapshots() returned a total of %d item(s) using SnapshotsPager.\n", len(allResults))
		})
	})

	Describe(`GetSnapshot - Retrieve a snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSnapshot(getSnapshotOptions *GetSnapshotOptions)`, func() {
			getSnapshotOptions := &sdsaasv2.GetSnapshotOptions{
				ID: core.StringPtr(snapIDLink),
			}

			snapshot, response, err := sdsaasService.GetSnapshot(getSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`UpdateSnapshot - Update a snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSnapshot(updateSnapshotOptions *UpdateSnapshotOptions)`, func() {
			snapshotPatchModel := &sdsaasv2.SnapshotPatch{
				Name: core.StringPtr("my-snapshot-updated"),
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSnapshotOptions := &sdsaasv2.UpdateSnapshotOptions{
				ID:            core.StringPtr(snapIDLink),
				SnapshotPatch: snapshotPatchModelAsPatch,
			}

			snapshot, response, err := sdsaasService.UpdateSnapshot(updateSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`DeleteVolumeMapping - Deletes the volume mapping`, Pending, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVolumeMapping(deleteVolumeMappingOptions *DeleteVolumeMappingOptions)`, func() {
			deleteVolumeMappingOptions := &sdsaasv2.DeleteVolumeMappingOptions{
				ID:              core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
				VolumeMappingID: core.StringPtr("r134-b274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			response, err := sdsaasService.DeleteVolumeMapping(deleteVolumeMappingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteVolumeMappings - Deletes all the volume mappings for a given host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVolumeMappings(deleteVolumeMappingsOptions *DeleteVolumeMappingsOptions)`, func() {
			deleteVolumeMappingsOptions := &sdsaasv2.DeleteVolumeMappingsOptions{
				ID: core.StringPtr(hostIDLink),
			}

			response, err := sdsaasService.DeleteVolumeMappings(deleteVolumeMappingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Wait for resource to delete
			time.Sleep(5 * time.Second)
		})
	})

	Describe(`DeleteSnapshot - Delete a snapshot`, Pending, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSnapshot(deleteSnapshotOptions *DeleteSnapshotOptions)`, func() {
			deleteSnapshotOptions := &sdsaasv2.DeleteSnapshotOptions{
				ID: core.StringPtr(snapIDLink),
			}

			response, err := sdsaasService.DeleteSnapshot(deleteSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSnapshots - Delete a filtered collection of snapshots`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSnapshots(deleteSnapshotsOptions *DeleteSnapshotsOptions)`, func() {
			deleteSnapshotsOptions := &sdsaasv2.DeleteSnapshotsOptions{
				SourceVolumeID: core.StringPtr(volumeIDLink),
			}

			response, err := sdsaasService.DeleteSnapshots(deleteSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Wait for resources to delete
			time.Sleep(5 * time.Second)
		})
	})

	Describe(`DeleteVolume - Delete a volume`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVolume(deleteVolumeOptions *DeleteVolumeOptions)`, func() {
			deleteVolumeOptions := &sdsaasv2.DeleteVolumeOptions{
				ID: &volumeIDLink,
			}

			response, err := sdsaasService.DeleteVolume(deleteVolumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteHost - Delete a Host`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteHost(deleteHostOptions *DeleteHostOptions)`, func() {
			deleteHostOptions := &sdsaasv2.DeleteHostOptions{
				ID: &hostIDLink,
			}

			response, err := sdsaasService.DeleteHost(deleteHostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteHmacCredentials - Delete HMAC credentials`, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteHmacCredentials(deleteHmacCredentialsOptions *DeleteHmacCredentialsOptions)`, func() {
			deleteHmacCredentialsOptions := &sdsaasv2.DeleteHmacCredentialsOptions{
				AccessKey: core.StringPtr("mytestkey"),
			}

			response, err := sdsaasService.DeleteHmacCredentials(deleteHmacCredentialsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSslCert - Delete SSL certificate`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSslCert(deleteSslCertOptions *DeleteSslCertOptions)`, func() {
			deleteSslCertOptions := &sdsaasv2.DeleteSslCertOptions{
				CertType: core.StringPtr("s3"),
			}

			response, err := sdsaasService.DeleteSslCert(deleteSslCertOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
