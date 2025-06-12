//go:build examples

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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.ibm.com/SDSaaS/sds-go-sdk/sdsaasv1"
)

// This file provides an example of how to use the sdsaas service.
//
// The following configuration properties are assumed to be defined:
// SDSAAS_URL=<service base url>
// SDSAAS_AUTH_TYPE=iam
// SDSAAS_APIKEY=<IAM apikey>
// SDSAAS_AUTH_URL=<IAM token service base URL - omit this if using the production environment> (in stage, set to https://iam.test.cloud.ibm.com)
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`SdsaasV1 Examples Tests`, func() {

	const externalConfigFile = "../sdsaas_v1.env"

	var (
		sdsaasService *sdsaasv1.SdsaasV1
		config        map[string]string

		// Variables to hold link values
		hostIDLink             string
		volumeIDLinkOne        string
		volumeIDLinkTwo        string
		volumeMappingIDLinkOne string
		snapIDLinkOne          string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(sdsaasv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			sdsaasServiceOptions := &sdsaasv1.SdsaasV1Options{}

			sdsaasService, err = sdsaasv1.NewSdsaasV1UsingExternalConfig(sdsaasServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(sdsaasService).ToNot(BeNil())
		})
	})

	Describe(`SdsaasV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`VolumeCreate request example`, Label("block"), func() {
			fmt.Println("\nVolumeCreate() result:")
			// begin-volume_create

			volumeCreateOptions := sdsaasService.NewVolumeCreateOptions(
				int64(10),
			)
			volumeCreateOptions.SetName("my-volume-one")

			volumeOne, responseOne, errOne := sdsaasService.VolumeCreate(volumeCreateOptions)
			if errOne != nil {
				panic(errOne)
			}
			b, _ := json.MarshalIndent(volumeOne, "", "  ")
			fmt.Println(string(b))

			// Create a second volume for additional host operations
			volumeCreateOptionsTwo := sdsaasService.NewVolumeCreateOptions(
				int64(10),
			)
			volumeCreateOptionsTwo.SetName("my-volume-two")

			volumeTwo, responseTwo, errTwo := sdsaasService.VolumeCreate(volumeCreateOptionsTwo)
			if errTwo != nil {
				panic(errTwo)
			}

			// end-volume_create

			time.Sleep(5 * time.Second)

			Expect(errOne).To(BeNil())
			Expect(responseOne.StatusCode).To(Equal(201))
			Expect(volumeOne).ToNot(BeNil())

			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(201))
			Expect(volumeTwo).ToNot(BeNil())

			volumeIDLinkOne = *volumeOne.ID
			volumeIDLinkTwo = *volumeTwo.ID
			fmt.Fprintf(GinkgoWriter, "Saved volumeIDLinkOne value: %v\n", volumeIDLinkOne)
			fmt.Fprintf(GinkgoWriter, "Saved volumeIDLinkTwo value: %v\n", volumeIDLinkTwo)
		})
		It(`HostCreate request example`, Label("block"), func() {
			fmt.Println("\nHostCreate() result:")
			// begin-host_create

			hostCreateOptions := sdsaasService.NewHostCreateOptions(
				"nqn.2014-06.org:9345",
			)

			host, response, err := sdsaasService.HostCreate(hostCreateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// end-host_create

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(host).ToNot(BeNil())

			hostIDLink = *host.ID
			fmt.Fprintf(GinkgoWriter, "Saved hostIDLink value: %v\n", hostIDLink)

			time.Sleep(5 * time.Second)
		})
		It(`Volumes request example`, Label("block"), func() {
			fmt.Println("\nVolumes() result:")
			// begin-volumes
			volumesOptions := &sdsaasv1.VolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume-one"),
			}

			pager, err := sdsaasService.NewVolumesPager(volumesOptions)
			if err != nil {
				panic(err)
			}
			var allResults []sdsaasv1.Volume
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-volumes
		})
		It(`Volume request example`, Label("block"), func() {
			fmt.Println("\nVolume() result:")
			// begin-volume

			volumeOptions := sdsaasService.NewVolumeOptions(
				volumeIDLinkOne,
			)

			volume, response, err := sdsaasService.Volume(volumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
		It(`VolumeUpdate request example`, Label("block"), func() {
			fmt.Println("\nVolumeUpdate() result:")
			// begin-volume_update

			volumeUpdateOptions := sdsaasService.NewVolumeUpdateOptions(
				volumeIDLinkOne,
			)

			volumePatch := map[string]interface{}{
				"name": "my-volume-updated-three",
			}

			volumeUpdateOptions.SetVolumePatch(volumePatch)

			volume, response, err := sdsaasService.VolumeUpdate(volumeUpdateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-volume_update

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
		It(`VolumeSnapshots request example`, Label("block"), func() {
			fmt.Println("\nVolumeSnapshots() result:")
			// begin-volume_snapshots
			volumeSnapshotsOptions := &sdsaasv1.VolumeSnapshotsOptions{
				Limit: core.Int64Ptr(int64(20)),
			}

			pager, err := sdsaasService.NewVolumeSnapshotsPager(volumeSnapshotsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []sdsaasv1.Snapshot
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-volume_snapshots
		})
		It(`VolumeSnapshotCreate request example`, Label("block"), func() {
			fmt.Println("\nVolumeSnapshotCreate() result:")
			// begin-volume_snapshot_create

			sourceVolumePrototypeModel := &sdsaasv1.SourceVolumePrototype{
				ID: core.StringPtr(volumeIDLinkOne),
			}

			volumeSnapshotCreateOptions := sdsaasService.NewVolumeSnapshotCreateOptions(
				sourceVolumePrototypeModel,
			)

			snapshot, response, err := sdsaasService.VolumeSnapshotCreate(volumeSnapshotCreateOptions)
			if err != nil {
				panic(err)
			}

			snapIDLinkOne = *snapshot.ID

			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-volume_snapshot_create

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())
		})
		It(`VolumeSnapshot request example`, Label("block"), func() {
			fmt.Println("\nVolumeSnapshot() result:")
			// begin-volume_snapshot

			volumeSnapshotOptions := sdsaasService.NewVolumeSnapshotOptions(
				snapIDLinkOne,
			)

			snapshot, response, err := sdsaasService.VolumeSnapshot(volumeSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-volume_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
		It(`VolumeSnapshotUpdate request example`, Label("block"), func() {
			fmt.Println("\nVolumeSnapshotUpdate() result:")
			// begin-volume_snapshot_update

			snapshotPatchModel := &sdsaasv1.SnapshotPatch{
				Name: core.StringPtr("my-snapshot-updated"),
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			volumeSnapshotUpdateOptions := sdsaasService.NewVolumeSnapshotUpdateOptions(
				snapIDLinkOne,
				snapshotPatchModelAsPatch,
			)

			snapshot, response, err := sdsaasService.VolumeSnapshotUpdate(volumeSnapshotUpdateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-volume_snapshot_update

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
		It(`Creds request example`, Label("object"), func() {
			fmt.Println("\nCreds() result:")
			// begin-creds

			credsOptions := sdsaasService.NewCredsOptions()

			credentialsFound, response, err := sdsaasService.Creds(credsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentialsFound, "", "  ")
			fmt.Println(string(b))

			// end-creds

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialsFound).ToNot(BeNil())
		})
		It(`CredCreate request example`, Label("object"), func() {
			fmt.Println("\nCredCreate() result:")
			// begin-cred_create

			credCreateOptions := sdsaasService.NewCredCreateOptions(
				"mytestkey",
			)

			credentialsUpdated, response, err := sdsaasService.CredCreate(credCreateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentialsUpdated, "", "  ")
			fmt.Println(string(b))

			// end-cred_create

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credentialsUpdated).ToNot(BeNil())
		})
		It(`CertTypes request example`, Label("object"), func() {
			fmt.Println("\nCertTypes() result:")
			// begin-cert_types

			certTypesOptions := sdsaasService.NewCertTypesOptions()

			certificateList, response, err := sdsaasService.CertTypes(certTypesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certificateList, "", "  ")
			fmt.Println(string(b))

			// end-cert_types

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certificateList).ToNot(BeNil())
		})
		It(`Cert request example`, Label("object"), func() {
			fmt.Println("\nCert() result:")
			// begin-cert

			certOptions := sdsaasService.NewCertOptions(
				"s3",
			)

			certificateFound, response, err := sdsaasService.Cert(certOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certificateFound, "", "  ")
			fmt.Println(string(b))

			// end-cert

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certificateFound).ToNot(BeNil())
		})
		It(`CertCreate request example`, Pending, Label("object"), func() {
			fmt.Println("\nCertCreate() result:")
			// begin-cert_create

			certCreateOptions := sdsaasService.NewCertCreateOptions(
				"s3",
				CreateMockReader("This is a mock file."),
			)

			certificateUpdated, response, err := sdsaasService.CertCreate(certCreateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certificateUpdated, "", "  ")
			fmt.Println(string(b))

			// end-cert_create

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certificateUpdated).ToNot(BeNil())
		})
		It(`CertUpdate request example`, Pending, Label("object"), func() {
			fmt.Println("\nCertUpdate() result:")
			// begin-cert_update

			certUpdateOptions := sdsaasService.NewCertUpdateOptions(
				"s3",
				CreateMockReader("This is a mock file."),
			)

			certificateUpdated, response, err := sdsaasService.CertUpdate(certUpdateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certificateUpdated, "", "  ")
			fmt.Println(string(b))

			// end-cert_update

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certificateUpdated).ToNot(BeNil())
		})
		It(`Hosts request example`, Label("block"), func() {
			fmt.Println("\nHosts() result:")
			// begin-hosts

			hostsOptions := sdsaasService.NewHostsOptions()
			hostsOptions.SetLimit(int64(20))
			hostsOptions.SetName("my-host")

			hostCollection, response, err := sdsaasService.Hosts(hostsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hostCollection, "", "  ")
			fmt.Println(string(b))

			// end-hosts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hostCollection).ToNot(BeNil())
		})
		It(`Host request example`, Label("block"), func() {
			fmt.Println("\nHost() result:")
			// begin-host

			hostOptions := sdsaasService.NewHostOptions(
				hostIDLink,
			)

			host, response, err := sdsaasService.Host(hostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// end-host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
		It(`HostUpdate request example`, Label("block"), func() {
			fmt.Println("\nHostUpdate() result:")
			// begin-host_update

			hostUpdateOptions := sdsaasService.NewHostUpdateOptions(
				hostIDLink,
			)

			hostPatch := map[string]interface{}{
				"name": "my-host-updated",
			}

			hostUpdateOptions.SetHostPatch(hostPatch)

			host, response, err := sdsaasService.HostUpdate(hostUpdateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// end-host_update

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
		It(`HostMappings request example`, Label("block"), func() {
			fmt.Println("\nHostMappings() result:")
			// begin-host_mappings

			hostMappingsOptions := sdsaasService.NewHostMappingsOptions(
				hostIDLink,
			)

			volumeMappingCollection, response, err := sdsaasService.HostMappings(hostMappingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeMappingCollection, "", "  ")
			fmt.Println(string(b))

			// end-host_mappings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMappingCollection).ToNot(BeNil())
		})
		It(`HostMappingCreate request example`, Label("block"), func() {
			fmt.Println("\nHostMappingCreate() result:")
			// begin-host_mapping_create

			volumeIdentityModel := &sdsaasv1.VolumeIdentity{
				ID: core.StringPtr(volumeIDLinkOne),
			}
			volumeIdentityModelTwo := &sdsaasv1.VolumeIdentity{
				ID: core.StringPtr(volumeIDLinkTwo),
			}

			hostMappingCreateOptions := sdsaasService.NewHostMappingCreateOptions(
				hostIDLink,
				volumeIdentityModel,
			)
			hostMappingCreateOptionsTwo := sdsaasService.NewHostMappingCreateOptions(
				hostIDLink,
				volumeIdentityModelTwo,
			)

			volumeMapping, response, err := sdsaasService.HostMappingCreate(hostMappingCreateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeMapping, "", "  ")
			fmt.Println(string(b))

			volumeMappingTwo, responseTwo, errTwo := sdsaasService.HostMappingCreate(hostMappingCreateOptionsTwo)
			if errTwo != nil {
				panic(errTwo)
			}
			b, _ = json.MarshalIndent(volumeMappingTwo, "", "  ")
			fmt.Println(string(b))

			// end-host_mapping_create

			time.Sleep(5 * time.Second)

			volumeMappingIDLinkOne = *volumeMapping.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeMapping).ToNot(BeNil())

			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(202))
			Expect(volumeMappingTwo).ToNot(BeNil())
		})
		It(`HostMapping request example`, Label("block"), func() {
			fmt.Println("\nHostMapping() result:")
			// begin-host_mapping

			hostMappingOptions := sdsaasService.NewHostMappingOptions(
				hostIDLink,
				volumeMappingIDLinkOne,
			)

			volumeMapping, response, err := sdsaasService.HostMapping(hostMappingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeMapping, "", "  ")
			fmt.Println(string(b))

			// end-host_mapping

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMapping).ToNot(BeNil())
		})
		It(`HostMappingDelete request example`, Label("block"), func() {
			// begin-host_mapping_delete

			hostMappingDeleteOptions := sdsaasService.NewHostMappingDeleteOptions(
				hostIDLink,
				volumeMappingIDLinkOne,
			)

			response, err := sdsaasService.HostMappingDelete(hostMappingDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from HostMappingDelete(): %d\n", response.StatusCode)
			}

			// end-host_mapping_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`HostMappingDeleteAll request example`, Label("block"), func() {
			// begin-host_mapping_delete_all

			hostMappingDeleteAllOptions := sdsaasService.NewHostMappingDeleteAllOptions(
				hostIDLink,
			)

			response, err := sdsaasService.HostMappingDeleteAll(hostMappingDeleteAllOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from HostMappingDeleteAll(): %d\n", response.StatusCode)
			}

			// end-host_mapping_delete_all

			time.Sleep(5 * time.Second)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`VolumeSnapshotDelete request example`, Label("block"), func() {
			// begin-volume_snapshot_delete

			volumeSnapshotDeleteOptions := sdsaasService.NewVolumeSnapshotDeleteOptions(
				snapIDLinkOne,
			)

			response, err := sdsaasService.VolumeSnapshotDelete(volumeSnapshotDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from VolumeSnapshotDelete(): %d\n", response.StatusCode)
			}

			// end-volume_snapshot_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			time.Sleep(5 * time.Second) // VolumeSnapshots need to be deleted before volumes
		})
		It(`VolumeDelete request example`, Label("block"), func() {
			// begin-volume_delete

			volumeDeleteOptions := sdsaasService.NewVolumeDeleteOptions(
				volumeIDLinkOne,
			)

			response, err := sdsaasService.VolumeDelete(volumeDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from VolumeDelete(): %d\n", response.StatusCode)
			}

			// Delete the second volume
			volumeDeleteOptionsTwo := sdsaasService.NewVolumeDeleteOptions(
				volumeIDLinkTwo,
			)

			responseTwo, errTwo := sdsaasService.VolumeDelete(volumeDeleteOptionsTwo)
			if errTwo != nil {
				panic(errTwo)
			}
			if responseTwo.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from VolumeDelete(): %d\n", response.StatusCode)
			}

			// end-volume_delete

			time.Sleep(5 * time.Second)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(204))
		})

		It(`HostDelete request example`, Label("block"), func() {
			// begin-host_delete

			hostDeleteOptions := sdsaasService.NewHostDeleteOptions(
				hostIDLink,
			)

			response, err := sdsaasService.HostDelete(hostDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from HostDelete(): %d\n", response.StatusCode)
			}

			// end-host_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CredDelete request example`, Label("object"), func() {
			// begin-cred_delete

			credDeleteOptions := sdsaasService.NewCredDeleteOptions(
				"mytestkey",
			)

			response, err := sdsaasService.CredDelete(credDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CredDelete(): %d\n", response.StatusCode)
			}

			// end-cred_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CertDelete request example`, Pending, Label("object"), func() {
			// begin-cert_delete

			certDeleteOptions := sdsaasService.NewCertDeleteOptions(
				"s3",
			)

			response, err := sdsaasService.CertDelete(certDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CertDelete(): %d\n", response.StatusCode)
			}

			// end-cert_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
