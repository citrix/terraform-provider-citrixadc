package citrixadc

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
// Framework-only test factory. The SDK v2 provider and the tf6 mux have been
// removed; the Framework provider serves every citrixadc_* type at protocol v6.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"citrixadc": providerserver.NewProtocol6WithError(provider.New("test")()),
}

var isCpxRun bool

var adcTestbed string

func init() {
	log.Printf("[DEBUG]  citrixadc-provider-test: In init")

	nsUrl := os.Getenv("NS_URL")
	isCpxRun = strings.Contains(nsUrl, "localhost")

	var exists bool
	adcTestbed, exists = os.LookupEnv("ADC_TESTBED")
	if !exists {
		adcTestbed = "UNSPECIFIED"
	}
}

func testAccPreCheck(t *testing.T) {
	// You can add any setup code here
	if v := os.Getenv("NS_URL"); v == "" {
		t.Fatal("NS_URL must be set for acceptance tests")
	}
}

func testAccGetFrameworkClient() (*service.NitroClient, error) {
	username := os.Getenv("NS_LOGIN")
	if username == "" {
		username = "nsroot"
	}

	password := os.Getenv("NS_PASSWORD")
	if password == "" {
		password = "nsroot"
	}

	endpoint := os.Getenv("NS_URL")
	if endpoint == "" {
		return nil, fmt.Errorf("NS_URL environment variable must be set")
	}

	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc",
	}

	params := service.NitroParams{
		Url:      endpoint,
		Username: username,
		Password: password,
		Headers:  userHeaders,
	}

	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	return client, nil
}
