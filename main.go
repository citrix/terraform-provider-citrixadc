// main.go
package main

import (
	"context"
	"flag"
	"log"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider version. This is set via the -ldflags flag during build.
var version string = "dev"

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// Framework-only provider. Every citrixadc_* resource and data source has been
	// migrated to the Plugin Framework, so the SDK v2 provider and the tf6 mux are
	// no longer needed. The Framework natively speaks Terraform plugin protocol v6
	// (same protocol the mux served), so there is no Terraform CLI compatibility
	// change for users.
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/citrix/citrixadc",
		Debug:   debugMode,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
