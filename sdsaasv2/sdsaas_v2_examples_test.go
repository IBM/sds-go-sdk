//go:build examples

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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/v2/sdsaasv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the sdsaas service.
//
// The following configuration properties are assumed to be defined:
// SDSAAS_URL=<service base url>
// SDSAAS_AUTH_TYPE=iam
// SDSAAS_APIKEY=<IAM apikey>
// SDSAAS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`SdsaasV2 Examples Tests`, func() {

	const externalConfigFile = "../sdsaas_v2.env"

	var (
		sdsaasService *sdsaasv2.SdsaasV2
		config        map[string]string

		// Variables to hold link values
		hostIDLink   string
		volumeIDLink string
		snapIDLink   string

		volumeMappingIDLink string
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
			config, err = core.GetServiceProperties(sdsaasv2.DefaultServiceName)
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

			sdsaasServiceOptions := &sdsaasv2.SdsaasV2Options{}

			sdsaasService, err = sdsaasv2.NewSdsaasV2UsingExternalConfig(sdsaasServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(sdsaasService).ToNot(BeNil())
		})
	})

	Describe(`SdsaasV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVolume request example`, Label("block"), func() {
			fmt.Println("\nCreateVolume() result:")
			// begin-create_volume

			createVolumeOptions := sdsaasService.NewCreateVolumeOptions(
				int64(1),
			)
			createVolumeOptions.SetName("my-volume-one")

			volumeSummary, response, err := sdsaasService.CreateVolume(createVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeSummary, "", "  ")
			fmt.Println(string(b))

			// end-create_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeSummary).ToNot(BeNil())

			volumeIDLink = *volumeSummary.ID
			fmt.Fprintf(GinkgoWriter, "Saved volumeIDLink value: %v\n", volumeIDLink)

			// Wait for volume to fully create
			time.Sleep(5 * time.Second)
		})
		It(`CreateHost request example`, Label("block"), func() {
			fmt.Println("\nCreateHost() result:")
			// begin-create_host

			createHostOptions := sdsaasService.NewCreateHostOptions(
				"nqn.2014-08.org.nvmexpress:uuid:12345678-4444-1234-1234-123456789abd",
			)
			createHostOptions.SetName("my-host")

			hostSummary, response, err := sdsaasService.CreateHost(createHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hostSummary, "", "  ")
			fmt.Println(string(b))

			// end-create_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hostSummary).ToNot(BeNil())

			hostIDLink = *hostSummary.ID
			fmt.Fprintf(GinkgoWriter, "Saved hostIDLink value: %v\n", hostIDLink)

			// Wait for host to fully create
			time.Sleep(5 * time.Second)
		})
		It(`ListVolumes request example`, Label("block"), func() {
			fmt.Println("\nListVolumes() result:")
			// begin-list_volumes
			listVolumesOptions := &sdsaasv2.ListVolumesOptions{
				Limit: core.Int64Ptr(int64(20)),
				Name:  core.StringPtr("my-volume-one"),
			}

			pager, err := sdsaasService.NewVolumesPager(listVolumesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []sdsaasv2.Volume
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_volumes
		})
		It(`GetVolume request example`, Label("block"), func() {
			fmt.Println("\nGetVolume() result:")
			// begin-get_volume

			getVolumeOptions := sdsaasService.NewGetVolumeOptions(
				volumeIDLink,
			)

			volume, response, err := sdsaasService.GetVolume(getVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-get_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
		It(`UpdateVolume request example`, Label("block"), func() {
			fmt.Println("\nUpdateVolume() result:")
			// begin-update_volume

			volumePatchModel := &sdsaasv2.VolumePatch{
				Name: core.StringPtr("my-volume-one-updated"),
			}
			volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVolumeOptions := sdsaasService.NewUpdateVolumeOptions(
				volumeIDLink,
				volumePatchModelAsPatch,
			)

			volume, response, err := sdsaasService.UpdateVolume(updateVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-update_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
		It(`ListHosts request example`, Label("block"), func() {
			fmt.Println("\nListHosts() result:")
			// begin-list_hosts
			listHostsOptions := &sdsaasv2.ListHostsOptions{
				Limit: core.Int64Ptr(int64(2)),
				Name:  core.StringPtr("my-host"),
			}

			pager, err := sdsaasService.NewHostsPager(listHostsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []sdsaasv2.Host
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_hosts
		})
		It(`GetHost request example`, Label("block"), func() {
			fmt.Println("\nGetHost() result:")
			// begin-get_host

			getHostOptions := sdsaasService.NewGetHostOptions(
				hostIDLink,
			)

			host, response, err := sdsaasService.GetHost(getHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// end-get_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
		It(`UpdateHost request example`, Label("block"), func() {
			fmt.Println("\nUpdateHost() result:")
			// begin-update_host

			hostPatchModel := &sdsaasv2.HostPatch{
				Name: core.StringPtr("my-host-updated"),
			}
			hostPatchModelAsPatch, asPatchErr := hostPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateHostOptions := sdsaasService.NewUpdateHostOptions(
				hostIDLink,
				hostPatchModelAsPatch,
			)

			host, response, err := sdsaasService.UpdateHost(updateHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// end-update_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(host).ToNot(BeNil())
		})
		It(`CreateVolumeMapping request example`, Label("block"), func() {
			fmt.Println("\nCreateVolumeMapping() result:")
			// begin-create_volume_mapping

			volumeIdentityModel := &sdsaasv2.VolumeIdentity{
				ID: core.StringPtr(volumeIDLink),
			}

			createVolumeMappingOptions := sdsaasService.NewCreateVolumeMappingOptions(
				hostIDLink,
				volumeIdentityModel,
			)

			volumeMappingReference, response, err := sdsaasService.CreateVolumeMapping(createVolumeMappingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeMappingReference, "", "  ")
			fmt.Println(string(b))

			// end-create_volume_mapping

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeMappingReference).ToNot(BeNil())

			volumeMappingIDLink = *volumeMappingReference.ID

			fmt.Fprintf(GinkgoWriter, "Saved volumeMappingIDLink value: %v\n", volumeMappingIDLink)

			// Wait for volume-mapping to fully create
			time.Sleep(5 * time.Second)
		})
		It(`ListVolumeMappings request example`, Label("block"), func() {
			fmt.Println("\nListVolumeMappings() result:")
			// begin-list_volume_mappings
			listVolumeMappingsOptions := &sdsaasv2.ListVolumeMappingsOptions{
				ID:    core.StringPtr(hostIDLink),
				Limit: core.Int64Ptr(int64(20)),
			}

			pager, err := sdsaasService.NewVolumeMappingsPager(listVolumeMappingsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []sdsaasv2.VolumeMapping
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_volume_mappings
		})
		It(`GetVolumeMapping request example`, Label("block"), func() {
			fmt.Println("\nGetVolumeMapping() result:")
			// begin-get_volume_mapping

			getVolumeMappingOptions := sdsaasService.NewGetVolumeMappingOptions(
				hostIDLink,
				volumeMappingIDLink,
			)

			volumeMapping, response, err := sdsaasService.GetVolumeMapping(getVolumeMappingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeMapping, "", "  ")
			fmt.Println(string(b))

			// end-get_volume_mapping

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeMapping).ToNot(BeNil())
		})
		It(`CreateHmacCredentials request example`, Label("object"), func() {
			fmt.Println("\nCreateHmacCredentials() result:")
			// begin-create_hmac_credentials

			createHmacCredentialsOptions := sdsaasService.NewCreateHmacCredentialsOptions(
				"mytestkey",
			)

			accessKeyResponse, response, err := sdsaasService.CreateHmacCredentials(createHmacCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_hmac_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessKeyResponse).ToNot(BeNil())
		})
		It(`ListHmacCredentials request example`, Label("object"), func() {
			fmt.Println("\nListHmacCredentials() result:")
			// begin-list_hmac_credentials

			listHmacCredentialsOptions := sdsaasService.NewListHmacCredentialsOptions()

			storageCredResponse, response, err := sdsaasService.ListHmacCredentials(listHmacCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(storageCredResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_hmac_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageCredResponse).ToNot(BeNil())
		})
		It(`ListCertificates request example`, Label("object"), func() {
			fmt.Println("\nListCertificates() result:")
			// begin-list_certificates

			listCertificatesOptions := sdsaasService.NewListCertificatesOptions()

			certListResponse, response, err := sdsaasService.ListCertificates(listCertificatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certListResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_certificates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(certListResponse).ToNot(BeNil())
		})
		It(`GetS3SslCertStatus request example`, Label("object"), func() {
			fmt.Println("\nGetS3SslCertStatus() result:")
			// begin-get_s3_ssl_cert_status

			getS3SslCertStatusOptions := sdsaasService.NewGetS3SslCertStatusOptions(
				"s3",
			)

			statusResponse, response, err := sdsaasService.GetS3SslCertStatus(getS3SslCertStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(statusResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_s3_ssl_cert_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statusResponse).ToNot(BeNil())
		})
		It(`CreateSslCert request example`, Pending, Label("object"), func() {
			fmt.Println("\nCreateSslCert() result:")
			// begin-create_ssl_cert

			createSslCertOptions := sdsaasService.NewCreateSslCertOptions(
				"s3",
			)

			certResponse, response, err := sdsaasService.CreateSslCert(createSslCertOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_ssl_cert

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certResponse).ToNot(BeNil())
		})
		It(`ReplaceSslCert request example`, Pending, Label("object"), func() {
			fmt.Println("\nReplaceSslCert() result:")
			// begin-replace_ssl_cert

			replaceSslCertOptions := sdsaasService.NewReplaceSslCertOptions(
				"s3",
			)

			certResponse, response, err := sdsaasService.ReplaceSslCert(replaceSslCertOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(certResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_ssl_cert

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(certResponse).ToNot(BeNil())
		})
		It(`CreateSnapshot request example`, Label("block"), func() {
			fmt.Println("\nCreateSnapshot() result:")
			// begin-create_snapshot

			sourceVolumePrototypeModel := &sdsaasv2.SourceVolumePrototype{
				ID: core.StringPtr(volumeIDLink),
			}

			createSnapshotOptions := &sdsaasv2.CreateSnapshotOptions{
				Name:         core.StringPtr("my-snapshot"),
				SourceVolume: sourceVolumePrototypeModel,
			}

			snapshot, response, err := sdsaasService.CreateSnapshot(createSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-create_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())

			snapIDLink = *snapshot.ID
			fmt.Fprintf(GinkgoWriter, "Saved snapIDLink value: %v\n", snapIDLink)

			// Wait for snapshot to fully create
			time.Sleep(5 * time.Second)
		})
		It(`ListSnapshots request example`, Label("block"), func() {
			fmt.Println("\nListSnapshots() result:")
			// begin-list_snapshots
			listSnapshotsOptions := &sdsaasv2.ListSnapshotsOptions{
				Limit:          core.Int64Ptr(int64(20)),
				SourceVolumeID: core.StringPtr(volumeIDLink),
			}

			pager, err := sdsaasService.NewSnapshotsPager(listSnapshotsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []sdsaasv2.Snapshot
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_snapshots
		})
		It(`GetSnapshot request example`, Label("block"), func() {
			fmt.Println("\nGetSnapshot() result:")
			// begin-get_snapshot

			getSnapshotOptions := sdsaasService.NewGetSnapshotOptions(
				snapIDLink,
			)

			snapshot, response, err := sdsaasService.GetSnapshot(getSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-get_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
		It(`UpdateSnapshot request example`, Label("block"), func() {
			fmt.Println("\nUpdateSnapshot() result:")
			// begin-update_snapshot

			snapshotPatchModel := &sdsaasv2.SnapshotPatch{
				Name: core.StringPtr("my-snapshot-updated"),
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSnapshotOptions := sdsaasService.NewUpdateSnapshotOptions(
				snapIDLink,
				snapshotPatchModelAsPatch,
			)

			snapshot, response, err := sdsaasService.UpdateSnapshot(updateSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-update_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
		It(`DeleteVolumeMapping request example`, Pending, Label("block"), func() {
			// begin-delete_volume_mapping

			deleteVolumeMappingOptions := sdsaasService.NewDeleteVolumeMappingOptions(
				hostIDLink,
				volumeMappingIDLink,
			)

			response, err := sdsaasService.DeleteVolumeMapping(deleteVolumeMappingOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteVolumeMapping(): %d\n", response.StatusCode)
			}

			// end-delete_volume_mapping

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Wait for the volume-mapping to delete successfully
			time.Sleep(5 * time.Second)
		})
		It(`DeleteVolumeMappings request example`, Label("block"), func() {
			// begin-delete_volume_mappings

			deleteVolumeMappingsOptions := sdsaasService.NewDeleteVolumeMappingsOptions(
				hostIDLink,
			)

			response, err := sdsaasService.DeleteVolumeMappings(deleteVolumeMappingsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteVolumeMappings(): %d\n", response.StatusCode)
			}

			// end-delete_volume_mappings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Wait for the volume-mappings to delete successfully
			time.Sleep(5 * time.Second)
		})
		It(`DeleteSnapshot request example`, Pending, Label("block"), func() {
			// begin-delete_snapshot

			deleteSnapshotOptions := sdsaasService.NewDeleteSnapshotOptions(
				snapIDLink,
			)

			response, err := sdsaasService.DeleteSnapshot(deleteSnapshotOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSnapshot(): %d\n", response.StatusCode)
			}

			// end-delete_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteSnapshots request example`, Label("block"), func() {
			// begin-delete_snapshots

			deleteSnapshotsOptions := sdsaasService.NewDeleteSnapshotsOptions(
				volumeIDLink,
			)

			response, err := sdsaasService.DeleteSnapshots(deleteSnapshotsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSnapshots(): %d\n", response.StatusCode)
			}

			// end-delete_snapshots

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Wait for resources to delete
			time.Sleep(5 * time.Second)
		})
		It(`DeleteVolume request example`, Label("block"), func() {
			// begin-delete_volume

			deleteVolumeOptions := sdsaasService.NewDeleteVolumeOptions(
				volumeIDLink,
			)

			response, err := sdsaasService.DeleteVolume(deleteVolumeOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteVolume(): %d\n", response.StatusCode)
			}

			// end-delete_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteHost request example`, Label("block"), func() {
			// begin-delete_host

			deleteHostOptions := sdsaasService.NewDeleteHostOptions(
				hostIDLink,
			)

			response, err := sdsaasService.DeleteHost(deleteHostOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteHost(): %d\n", response.StatusCode)
			}

			// end-delete_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteHmacCredentials request example`, Label("object"), func() {
			// begin-delete_hmac_credentials

			deleteHmacCredentialsOptions := sdsaasService.NewDeleteHmacCredentialsOptions(
				"mytestkey",
			)

			response, err := sdsaasService.DeleteHmacCredentials(deleteHmacCredentialsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteHmacCredentials(): %d\n", response.StatusCode)
			}

			// end-delete_hmac_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteSslCert request example`, Pending, Label("object"), func() {
			// begin-delete_ssl_cert

			deleteSslCertOptions := sdsaasService.NewDeleteSslCertOptions(
				"s3",
			)

			response, err := sdsaasService.DeleteSslCert(deleteSslCertOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSslCert(): %d\n", response.StatusCode)
			}

			// end-delete_ssl_cert

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
