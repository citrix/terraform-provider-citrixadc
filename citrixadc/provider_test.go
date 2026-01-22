package citrixadc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
var testAccProvider *schema.Provider

var isCpxRun bool

var adcTestbed string

func init() {
	log.Printf("[DEBUG]  citrixadc-provider-test: In init")
	testAccProvider = Provider()

	// For backward compatibility, keep the original provider factories
	// but also create muxed provider factories
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"citrixadc": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}

	// Muxed provider factories that include both SDK v2 and Framework providers
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"citrixadc": func() (tfprotov6.ProviderServer, error) {
			return createMuxedProvider()
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

// createMuxedProvider creates a muxed provider that includes both SDK v2 and Framework providers
func createMuxedProvider() (tfprotov6.ProviderServer, error) {
	ctx := context.Background()

	// Create the SDK v2 provider and upgrade it to tf6
	sdkV2Provider := schema.NewGRPCProviderServer(Provider())
	upgradedSDKProvider, err := tf5to6server.UpgradeServer(ctx, func() tfprotov5.ProviderServer {
		return sdkV2Provider
	})
	if err != nil {
		log.Fatalf("Failed to upgrade SDK v2 provider: %v", err)
		return nil, err
	}

	// Create the Framework provider (already tf6)
	frameworkProviderFunc := providerserver.NewProtocol6(provider.New("test")())

	// Create the mux server
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSDKProvider
		},
		frameworkProviderFunc,
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		return nil, err
	}

	return muxServer, nil
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

// testAccGetClient returns a configured NITRO client using environment variables
// This utility function can be reused across all test cases to avoid duplicating client creation logic
func testAccGetClient() (*service.NitroClient, error) {
	// Try to get client from SDK v2 provider first (for backward compatibility)
	if testAccProvider.Meta() != nil {
		return testAccProvider.Meta().(*NetScalerNitroClient).client, nil
	}

	// Fallback to creating client directly from environment variables
	// This is needed when using muxed providers where Meta() might be nil
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

	params := service.NitroParams{
		Url:      endpoint,
		Username: username,
		Password: password,
	}

	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	return client, nil
}
