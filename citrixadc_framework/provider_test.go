package citrixadc_framework

import (
	"fmt"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"citrixadc": providerserver.NewProtocol6WithError(New("test")()),
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

func TestProviderNew(t *testing.T) {
	provider := New("test")()

	if provider == nil {
		t.Fatal("Provider should not be nil")
	}

	// Test that the provider implements the required interface
	_, ok := provider.(*CitrixAdcFrameworkProvider)
	if !ok {
		t.Fatal("Provider should be of type *CitrixAdcFrameworkProvider")
	}
}

func TestProviderModel(t *testing.T) {
	model := CitrixAdcFrameworkProviderModel{}

	// Test that all fields are properly defined
	if !model.Username.IsNull() {
		t.Error("New model Username should be null")
	}

	if !model.Password.IsNull() {
		t.Error("New model Password should be null")
	}

	if !model.Endpoint.IsNull() {
		t.Error("New model Endpoint should be null")
	}
}
