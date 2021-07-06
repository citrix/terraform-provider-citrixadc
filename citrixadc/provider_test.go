package citrixadc

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var isCpxRun bool
var isCluster bool
var adcTestbed string

func init() {
	log.Printf("[DEBUG]  citrixadc-provider-test: In init")
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"citrixadc": testAccProvider,
	}

	nsUrl := os.Getenv("NS_URL")
	isCpxRun = strings.Contains(nsUrl, "localhost")

	isCluster = testIsTargetAdcCluster()

	var exists bool
	adcTestbed, exists = os.LookupEnv("ADC_TESTBED")
	if !exists {
		adcTestbed = "UNSPECIFIED"
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NS_URL"); v == "" {
		t.Fatal("NS_URL must be set for acceptance tests")
	}
}
