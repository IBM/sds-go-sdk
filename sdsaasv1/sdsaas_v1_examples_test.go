//go:build examples

/**
 * (C) Copyright IBM Corp. 2024.
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

	"github.com/IBM/sds-go-sdk/sdsaasv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
		hostIDLink      string
		volumeIDLinkOne string
		volumeIDLinkTwo string
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
		It(`VolumeCreate request example`, func() {
			fmt.Println("\nVolumeCreate() result:")
			// begin-volume_create

			volumeCreateOptions := sdsaasService.NewVolumeCreateOptions(
				int64(10),
			)

			volumeCreateOptions.SetHostnqnstring("nqn.2014-06.org:9345")
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
			volumeCreateOptionsTwo.SetHostnqnstring("nqn.2014-06.org:9345")
			volumeCreateOptionsTwo.SetName("my-volume-two")

			volumeTwo, responseTwo, errTwo := sdsaasService.VolumeCreate(volumeCreateOptionsTwo)
			if errTwo != nil {
				panic(errTwo)
			}

			// end-volume_create

			time.Sleep(8 * time.Second)

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
		It(`HostCreate request example`, func() {
			fmt.Println("\nHostCreate() result:")
			// begin-host_create

			hostCreateOptions := sdsaasService.NewHostCreateOptions(
				"nqn.2014-06.org:9345",
			)

			hostCreateOptions.SetName("my-host")

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
		})
		It(`Volumes request example`, func() {
			fmt.Println("\nVolumes() result:")
			// begin-volumes

			volumesOptions := sdsaasService.NewVolumesOptions()
			volumesOptions.SetLimit(int64(10))
			volumesOptions.SetName("my-volume-one")

			volumeCollection, response, err := sdsaasService.Volumes(volumesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeCollection, "", "  ")
			fmt.Println(string(b))

			// end-volumes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeCollection).ToNot(BeNil())
		})
		It(`Volume request example`, func() {
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
		It(`VolumeUpdate request example`, func() {
			fmt.Println("\nVolumeUpdate() result:")
			// begin-volume_update

			volumeUpdateOptions := sdsaasService.NewVolumeUpdateOptions(
				volumeIDLinkOne,
			)

			volumePatch := map[string]interface{}{
				"name": "my-volume-updated",
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
		It(`Creds request example`, func() {
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
		It(`CredCreate request example`, func() {
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
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialsUpdated).ToNot(BeNil())
		})
		It(`Cert request example`, func() {
			fmt.Println("\nCert() result:")
			// begin-cert

			certOptions := sdsaasService.NewCertOptions()

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
		// It(`CertUpload request example`, func() {
		// 	fmt.Println("\nCertUpload() result:")
		// 	// begin-cert_upload

		// 	// Generate a temp cert
		// 	var tc, tk string
		// 	cert, key, _ := testcerts.GenerateCertsToTempFile("/tmp/")
		// 	c, _ := os.Open(cert)
		// 	k, _ := os.Open(key)
		// 	defer c.Close()
		// 	defer k.Close()
		// 	scanner1 := bufio.NewScanner(c)
		// 	for scanner1.Scan() {
		// 		tc = tc + scanner1.Text() + `\n`
		// 	}
		// 	scanner2 := bufio.NewScanner(k)
		// 	for scanner2.Scan() {
		// 		tk = tk + scanner2.Text() + `\n`
		// 	}
		// 	tempCert := tk + tc

		// 	certUploadOptions := sdsaasService.NewCertUploadOptions(
		// 		CreateMockReader(tempCert),
		// 	)

		// 	certificateUpdated, response, err := sdsaasService.CertUpload(certUploadOptions)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	b, _ := json.MarshalIndent(certificateUpdated, "", "  ")
		// 	fmt.Println(string(b))

		// 	// end-cert_upload

		// 	Expect(err).To(BeNil())
		// 	Expect(response.StatusCode).To(Equal(202))
		// 	Expect(certificateUpdated).ToNot(BeNil())
		// })

		It(`Hosts request example`, func() {
			fmt.Println("\nHosts() result:")
			// begin-hosts

			hostsOptions := sdsaasService.NewHostsOptions()
			hostsOptions.SetLimit(int64(10))
			hostsOptions.SetName("myhost1")

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
		It(`Host request example`, func() {
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
		It(`HostUpdate request example`, func() {
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
		It(`HostVolUpdate request example`, func() {
			fmt.Println("\nHostVolUpdate() result:")
			// begin-host_vol_update

			hostVolUpdateOptions := sdsaasService.NewHostVolUpdateOptions(
				hostIDLink,
				volumeIDLinkOne,
			)

			host, response, err := sdsaasService.HostVolUpdate(hostVolUpdateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(host, "", "  ")
			fmt.Println(string(b))

			// Assign the second volume to the host as well
			hostVolUpdateOptionsTwo := sdsaasService.NewHostVolUpdateOptions(
				hostIDLink,
				volumeIDLinkTwo,
			)

			hostTwo, responseTwo, errTwo := sdsaasService.HostVolUpdate(hostVolUpdateOptionsTwo)
			if err != nil {
				panic(err)
			}

			// end-host_vol_update

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(host).ToNot(BeNil())

			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(202))
			Expect(hostTwo).ToNot(BeNil())
		})

		It(`HostVolDelete request example`, func() {
			// begin-host_volid_delete

			hostVolDeleteOptions := sdsaasService.NewHostVolDeleteOptions(
				hostIDLink,
				volumeIDLinkTwo,
			)

			response, err := sdsaasService.HostVolDelete(hostVolDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from HostVolDelete(): %d\n", response.StatusCode)
			}

			// end-host_volid_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Make sure the volumes are no longer attached before deleting anything
			time.Sleep(15 * time.Second)
		})

		It(`HostVolDeleteall request example`, func() {
			// begin-host_vol_deleteall

			hostVolDeleteallOptions := sdsaasService.NewHostVolDeleteallOptions(
				hostIDLink,
			)

			response, err := sdsaasService.HostVolDeleteall(hostVolDeleteallOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from HostVolDeleteall(): %d\n", response.StatusCode)
			}

			// end-host_vol_deleteall

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Make sure the volumes are no longer attached before deleting anything
			time.Sleep(15 * time.Second)
		})

		It(`VolumeDelete request example`, func() {
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

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			Expect(errTwo).To(BeNil())
			Expect(responseTwo.StatusCode).To(Equal(204))
		})
		It(`CredDelete request example`, func() {
			// begin-cred_delete

			credDeleteOptions := sdsaasService.NewCredDeleteOptions(
				"mytestkey",
			)

			response, err := sdsaasService.CredDelete(credDeleteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from CredDelete(): %d\n", response.StatusCode)
			}

			// end-cred_delete

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`HostDelete request example`, func() {
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
	})
})
