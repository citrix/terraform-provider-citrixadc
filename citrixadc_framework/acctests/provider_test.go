package acctests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	citrixadcProvider "github.com/citrix/terraform-provider-citrixadc/citrixadc"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"citrixadc": func() (tfprotov6.ProviderServer, error) {
		ctx := context.Background()

		// Create the SDK v2 provider and upgrade it to tf6
		sdkV2Provider := schema.NewGRPCProviderServer(citrixadcProvider.Provider())
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
	},
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

// func TestProviderNew(t *testing.T) {
// 	provider := provider.New("test")()

// 	if provider == nil {
// 		t.Fatal("Provider should not be nil")
// 	}

// 	// Test that the provider implements the required interface
// 	_, ok := provider.(*(provider.CitrixAdcFrameworkProvider))
// 	if !ok {
// 		t.Fatal("Provider should be of type *CitrixAdcFrameworkProvider")
// 	}
// }

// func TestProviderModel(t *testing.T) {
// 	model := provider.CitrixAdcFrameworkProviderModel{}

// 	// Test that all fields are properly defined
// 	if !model.Username.IsNull() {
// 		t.Error("New model Username should be null")
// 	}

// 	if !model.Password.IsNull() {
// 		t.Error("New model Password should be null")
// 	}

// 	if !model.Endpoint.IsNull() {
// 		t.Error("New model Endpoint should be null")
// 	}
// }
