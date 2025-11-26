// main.go
package main

import (
	"context"
	"flag"
	"log"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider version. This is set via the -ldflags flag during build.
var version string = "dev"

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()

	// Create the SDK v2 provider and upgrade it to tf6
	sdkV2Provider := schema.NewGRPCProviderServer(citrixadc.Provider())
	upgradedSDKProvider, err := tf5to6server.UpgradeServer(ctx, func() tfprotov5.ProviderServer {
		return sdkV2Provider
	})
	if err != nil {
		log.Fatalf("Failed to upgrade SDK v2 provider: %v", err)
	}

	// Create the Framework provider (already tf6)
	frameworkProviderFunc := providerserver.NewProtocol6(provider.New(version)())

	// Create the mux server
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSDKProvider
		},
		frameworkProviderFunc,
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatalf("Failed to create mux server: %v", err)
	}

	var serveOpts []tf6server.ServeOpt
	if debugMode {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/citrix/citrixadc",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
