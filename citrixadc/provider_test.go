package citrixadc

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)

var isCpxRun bool

var adcTestbed string

func init() {
	log.Printf("[DEBUG]  citrixadc-provider-test: In init")
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"citrixadc": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}

	nsUrl := os.Getenv("NS_URL")
	isCpxRun = strings.Contains(nsUrl, "localhost")

	var exists bool
	adcTestbed, exists = os.LookupEnv("ADC_TESTBED")
	if !exists {
		adcTestbed = "UNSPECIFIED"
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NS_URL"); v == "" {
		t.Fatal("NS_URL must be set for acceptance tests")
	}
}

// testAccGetClient creates and returns a configured NITRO client using environment variables
// This utility function can be reused across all test cases to avoid duplicating client creation logic
func testAccGetClient() (*service.NitroClient, error) {
	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc",
	}
	params := service.NitroParams{
		Url:       os.Getenv("NS_URL"),
		Username:  os.Getenv("NS_LOGIN"),
		Password:  os.Getenv("NS_PASSWORD"),
		SslVerify: false,
		Headers:   userHeaders,
	}

	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return client, nil
}
