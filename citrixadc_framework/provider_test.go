package citrixadc_framework

import (
	"testing"
)

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
