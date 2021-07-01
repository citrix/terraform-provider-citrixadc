package citrixadc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
)

func uploadTestdataFile(c *NetScalerNitroClient, t *testing.T, filename, targetDir string) error {
	nsClient := c.client

	// Get here path
	_, here_filename, _, _ := runtime.Caller(1)
	b, err := ioutil.ReadFile(path.Join(path.Dir(here_filename), "testdata", filename))

	if err != nil {
		return err
	}

	sf := system.Systemfile{
		Filename:     filename,
		Filecontent:  base64.StdEncoding.EncodeToString(b),
		Filelocation: targetDir,
	}
	_, err = nsClient.AddResource(service.Systemfile.Type(), filename, &sf)
	if err != nil && strings.Contains(err.Error(), "File already exists") {
		url_args := map[string]string{"filelocation": strings.Replace(targetDir, "/", "%2F", -1)}
		err := nsClient.DeleteResourceWithArgsMap(service.Systemfile.Type(), filename, url_args)
		if err != nil {
			return err
		}
		_, err = nsClient.AddResource(service.Systemfile.Type(), filename, &sf)
		if err != nil {
			return err
		}
	}
	return nil
}

var helperClient *NetScalerNitroClient

func testHelperInstantiateClient(nsUrl, username, password string, sslVerify bool) (*NetScalerNitroClient, error) {
	if helperClient != nil {
		log.Printf("Returning existing helper client\n")
		return helperClient, nil
	}

	if nsUrl == "" {
		if nsUrl = os.Getenv("NS_URL"); nsUrl == "" {
			return nil, errors.New("No nsUrl defined")
		}
	}

	if username == "" {
		if username = os.Getenv("NS_LOGIN"); username == "" {
			username = "nsroot"
		}
	}

	if password == "" {
		if password = os.Getenv("NS_PASSWORD"); password == "" {
			password = "nsroot"
		}
	}

	c := NetScalerNitroClient{
		Username: username,
		Password: password,
		Endpoint: nsUrl,
	}

	params := service.NitroParams{
		Url:      nsUrl,
		Username: username,
		Password: password,
		//ProxiedNs: d.Get("proxied_ns").(string),
		SslVerify: sslVerify,
	}
	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	c.client = client
	helperClient = &c
	log.Printf("Helper client instantiated\n")

	return helperClient, nil
}

func testHelperEnsureResourceDeletion(c *NetScalerNitroClient, t *testing.T, resourceType, resourceName string, deleteArgsMap map[string]string) {
	if _, err := c.client.FindResource(resourceType, resourceName); err != nil {
		targetSubstring := fmt.Sprintf("No resource %s of type %s found", resourceName, resourceType)
		actualError := err.Error()
		t.Logf("targetSubstring \"%s\"", targetSubstring)
		t.Logf("actualError \"%s\"", actualError)
		if strings.Contains(err.Error(), targetSubstring) {
			t.Logf("Ensure delete found no remaining resource %s", resourceName)
			return
		} else {
			t.Fatalf("Unexpected error while ensuring delete of resource %v. %v", resourceName, err)
			return
		}
	}

	// Fallthrough
	if deleteArgsMap == nil {
		if err := c.client.DeleteResource(resourceType, resourceName); err != nil {
			t.Logf("Ensuring delete failed for resource %s.", resourceName)
			t.Fatal(err)
			return
		} else {
			t.Logf("Ensuring deletion of %s successful", resourceName)
		}
	} else {
		if err := c.client.DeleteResourceWithArgsMap(resourceType, resourceName, deleteArgsMap); err != nil {
			t.Logf("Ensuring delete failed for resource %s with argsMap %v", resourceName, deleteArgsMap)
			t.Fatal(err)
			return
		} else {
			t.Logf("Ensuring deletion of %s successful", resourceName)
		}
	}

}
func testHelperVerifyImmutabilityFunc(c *NetScalerNitroClient, t *testing.T, resourceType, resourceName string, resourceInstance interface{}, attribute string) {
	if _, err := c.client.UpdateResource(resourceType, resourceName, resourceInstance); err != nil {
		r := regexp.MustCompile(fmt.Sprintf("errorcode.*278.*Invalid argument \\[%s\\]", attribute))

		if r.Match([]byte(err.Error())) {
			t.Logf("Succesfully verified immutability of attribute \"%s\"", attribute)
		} else {
			t.Errorf("Error while assesing immutability of attribute \"%s\"", attribute)
			t.Fatal(err)
		}
	} else {
		t.Fatalf("Error (no error) while assesing immutability of attribute \"%s\"", attribute)
	}
}

func testIsTargetAdcCluster() bool {
	log.Printf("[DEBUG]  citrixadc-provider-test: In isTargetAdcCluster")
	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		//lintignore:R009
		panic(err)
	}
	nsClient := c.client

	datalist, err := nsClient.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		//lintignore:R009
		panic(err)
	}

	if len(datalist) == 0 {
		return false
	} else {
		return true
	}
}
