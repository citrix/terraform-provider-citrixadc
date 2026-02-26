package sslvserver_sslciphersuite_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslvserverSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslciphersuiteBindingResource)(nil)

func NewSslvserverSslciphersuiteBindingResource() resource.Resource {
	return &SslvserverSslciphersuiteBindingResource{}
}

// SslvserverSslciphersuiteBindingResource defines the resource implementation.
type SslvserverSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslciphersuite_binding"
}

func (r *SslvserverSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslciphersuite_binding resource")

	// sslvserver_sslciphersuite_binding := sslvserver_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslvserver_sslciphersuite_binding.Type(), &sslvserver_sslciphersuite_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslvserver_sslciphersuite_binding-config")

	tflog.Trace(ctx, "Created sslvserver_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslciphersuite_binding resource")

	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslvserver_sslciphersuite_binding resource")

	// Create API request body from the model
	// sslvserver_sslciphersuite_binding := sslvserver_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslvserver_sslciphersuite_binding.Type(), &sslvserver_sslciphersuite_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslvserver_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslciphersuite_binding resource")

	// For sslvserver_sslciphersuite_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslvserver_sslciphersuite_binding resource from state")
}

// Helper function to read sslvserver_sslciphersuite_binding data from API
func (r *SslvserverSslciphersuiteBindingResource) readSslvserverSslciphersuiteBindingFromApi(ctx context.Context, data *SslvserverSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslvserver_sslciphersuite_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslciphersuite_binding, got error: %s", err))
		return
	}

	sslvserver_sslciphersuite_bindingSetAttrFromGet(ctx, data, getResponseData)

}
