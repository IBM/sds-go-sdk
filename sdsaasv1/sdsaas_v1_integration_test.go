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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.ibm.com/SDSaaS/sds-go-sdk/sdsaasv1"
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
		snapIDLink             string
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

	Describe(`VolumeCreate - Create a new volume`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeCreate(volumeCreateOptions *VolumeCreateOptions)`, func() {
			volumeCreateOptions := &sdsaasv1.VolumeCreateOptions{
				Capacity: core.Int64Ptr(int64(2)),
				Name:     core.StringPtr("my-volume-one"),
			}

			volumeCreateOptionsTwo := &sdsaasv1.VolumeCreateOptions{
				Capacity: core.Int64Ptr(int64(2)),
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

	Describe(`HostCreate - Creates a host`, Label("block"), func() {
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
				Nqn:            core.StringPtr("nqn.2014-06.org:9345"),
				Name:           core.StringPtr("my-host"),
				VolumeMappings: []sdsaasv1.VolumeMappingPrototype{*volumeMappingPrototypeModel},
				Psk:            core.StringPtr("NVMeTLSkey-1:01:5CBxDU8ejK+PrqIjTau0yDHnBV2CdfvP6hGmqnPdKhJ9tfi2:"),
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

	Describe(`Volumes - This request lists all volumes in the region`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Volumes(volumesOptions *VolumesOptions) with pagination`, func() {
			volumesOptions := &sdsaasv1.VolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume-one"),
			}

			volumesOptions.Start = nil
			volumesOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv1.Volume
			for {
				volumeCollection, response, err := sdsaasService.Volumes(volumesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(volumeCollection).ToNot(BeNil())
				allResults = append(allResults, volumeCollection.Volumes...)

				volumesOptions.Start, err = volumeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if volumesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`Volumes(volumesOptions *VolumesOptions) using VolumesPager`, func() {
			volumesOptions := &sdsaasv1.VolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				// Name:  core.StringPtr("my-volume"),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewVolumesPager(volumesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv1.Volume
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewVolumesPager(volumesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "Volumes() returned a total of %d item(s) using VolumesPager.\n", len(allResults))
		})
	})

	Describe(`Volume - Retrieve a volume profile`, Label("block"), func() {
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

	Describe(`VolumeUpdate - Update a volume`, Label("block"), func() {
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

	Describe(`VolumeSnapshotCreate - Create a snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeSnapshotCreate(volumeSnapshotCreateOptions *VolumeSnapshotCreateOptions)`, func() {
			SourceVolumePrototype := &sdsaasv1.SourceVolumePrototype{
				ID: core.StringPtr(volumeIDLink),
			}

			volumeSnapshotCreateOptions := &sdsaasv1.VolumeSnapshotCreateOptions{
				SourceVolume: SourceVolumePrototype,
				Name:         core.StringPtr("my-snapshot"),
			}

			snapshot, response, err := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptions)

			snapIDLink = *snapshot.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`VolumeSnapshots - List all snapshots`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeSnapshots(volumeSnapshotsOptions *VolumeSnapshotsOptions) with pagination`, func() {
			volumeSnapshotsOptions := &sdsaasv1.VolumeSnapshotsOptions{
				Limit:          core.Int64Ptr(int64(20)),
				Name:           core.StringPtr("my-snapshot"),
				SourceVolumeID: core.StringPtr(volumeIDLink),
			}

			volumeSnapshotsOptions.Start = nil
			volumeSnapshotsOptions.Limit = core.Int64Ptr(1)

			var allResults []sdsaasv1.Snapshot
			for {
				snapshotCollection, response, err := sdsaasService.VolumeSnapshots(volumeSnapshotsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(snapshotCollection).ToNot(BeNil())
				allResults = append(allResults, snapshotCollection.Snapshots...)

				volumeSnapshotsOptions.Start, err = snapshotCollection.GetNextStart()
				Expect(err).To(BeNil())

				if volumeSnapshotsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`VolumeSnapshots(volumeSnapshotsOptions *VolumeSnapshotsOptions) using VolumeSnapshotsPager`, Label("block"), func() {
			volumeSnapshotsOptions := &sdsaasv1.VolumeSnapshotsOptions{
				Limit:          core.Int64Ptr(int64(20)),
				Name:           core.StringPtr("my-snapshot"),
				SourceVolumeID: core.StringPtr(volumeIDLink),
			}

			// Test GetNext().
			pager, err := sdsaasService.NewVolumeSnapshotsPager(volumeSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []sdsaasv1.Snapshot
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = sdsaasService.NewVolumeSnapshotsPager(volumeSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "VolumeSnapshots() returned a total of %d item(s) using VolumeSnapshotsPager.\n", len(allResults))
		})
	})

	Describe(`VolumeSnapshot - Retrieve a single snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeSnapshot(volumeSnapshotOptions *VolumeSnapshotOptions)`, func() {
			volumeSnapshotOptions := &sdsaasv1.VolumeSnapshotOptions{
				SnapID: core.StringPtr(snapIDLink),
			}

			snapshot, response, err := sdsaasService.VolumeSnapshot(volumeSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`VolumeSnapshotUpdate - Update a snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeSnapshotUpdate(volumeSnapshotUpdateOptions *VolumeSnapshotUpdateOptions)`, func() {
			snapshotPatchModel := &sdsaasv1.SnapshotPatch{
				Name: core.StringPtr("my-snapshot-updated"),
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			volumeSnapshotUpdateOptions := &sdsaasv1.VolumeSnapshotUpdateOptions{
				SnapID:        core.StringPtr(snapIDLink),
				SnapshotPatch: snapshotPatchModelAsPatch,
			}

			snapshot, response, err := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`Creds - List storage account credentials`, Label("object"), func() {
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

	Describe(`CredCreate - Create or modify storage account credentials`, Label("object"), func() {
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

	Describe(`CertTypes - List the allowed certificate types`, Label("object"), func() {
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

	Describe(`Cert - Retrieves the SSL certificate expiration date and status`, Pending, Label("object"), func() {
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
	Describe(`CertCreate - Creates a new SSL Certificate`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CertCreate(certCreateOptions *CertCreateOptions)`, func() {
			certCreateOptions := &sdsaasv1.CertCreateOptions{
				Cert: core.StringPtr("s3"),
				Body: CreateMockReader("This is a mock file."),
			}

			certificateUpdated, response, err := sdsaasService.CertCreate(certCreateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certificateUpdated).ToNot(BeNil())
		})
	})

	Describe(`CertUpdate - Updates the SSL Certificate`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CertUpdate(certUpdateOptions *CertUpdateOptions)`, func() {
			certUpdateOptions := &sdsaasv1.CertUpdateOptions{
				Cert: core.StringPtr("s3"),
				Body: CreateMockReader("This is a mock file."),
			}

			certificateUpdated, response, err := sdsaasService.CertUpdate(certUpdateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certificateUpdated).ToNot(BeNil())
		})
	})

	Describe(`Hosts - Lists all hosts and all host IDs`, Label("block"), func() {
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

	Describe(`Host - Retrieve a host by ID`, Label("block"), func() {
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

	Describe(`HostUpdate - Update a host`, Label("block"), func() {
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

	Describe(`HostMappings - List all volume mappings for a host`, Label("block"), func() {
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

	Describe(`HostMappingCreate - Create a Volume mapping for a host`, Label("block"), func() {
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

	Describe(`HostMapping - Retrieve a volume mapping`, Label("block"), func() {
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

	Describe(`HostMappingDelete - Deletes the given volume mapping for a specific host`, Label("block"), func() {
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

	Describe(`HostMappingDeleteAll - Deletes all the volume mappings for a given host`, Label("block"), func() {
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

	Describe(`HostDelete - Delete a specific host`, Label("block"), func() {
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

	Describe(`VolumeSnapshotDelete - Delete a single snapshot`, Label("block"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeSnapshotDelete(volumeSnapshotDeleteOptions *VolumeSnapshotDeleteOptions)`, func() {
			volumeSnapshotDeleteOptions := &sdsaasv1.VolumeSnapshotDeleteOptions{
				SnapID: core.StringPtr(snapIDLink),
			}

			response, err := sdsaasService.VolumeSnapshotDelete(volumeSnapshotDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			time.Sleep(5 * time.Second) // Wait for the snapshot to delete before deleting the volume
		})
	})

	Describe(`VolumeDelete - Delete a volume`, Label("block"), func() {
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

	Describe(`CredDelete - Delete storage account credentials`, Label("object"), func() {
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

	Describe(`CertDelete - Delete SSL certificate`, Pending, Label("object"), func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CertDelete(certDeleteOptions *CertDeleteOptions)`, func() {
			certDeleteOptions := &sdsaasv1.CertDeleteOptions{
				Cert: core.StringPtr("s3"),
			}

			response, err := sdsaasService.CertDelete(certDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
