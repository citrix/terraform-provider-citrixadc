package lsnclient_nsacl_binding

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
var _ resource.Resource = &LsnclientNsaclBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNsaclBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNsaclBindingResource)(nil)

func NewLsnclientNsaclBindingResource() resource.Resource {
	return &LsnclientNsaclBindingResource{}
}

// LsnclientNsaclBindingResource defines the resource implementation.
type LsnclientNsaclBindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNsaclBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNsaclBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_nsacl_binding"
}

func (r *LsnclientNsaclBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNsaclBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_nsacl_binding resource")

	// lsnclient_nsacl_binding := lsnclient_nsacl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl_binding.Type(), &lsnclient_nsacl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_nsacl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnclient_nsacl_binding-config")

	tflog.Trace(ctx, "Created lsnclient_nsacl_binding resource")

	// Read the updated state back
	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsaclBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_nsacl_binding resource")

	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsaclBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnclient_nsacl_binding resource")

	// Create API request body from the model
	// lsnclient_nsacl_binding := lsnclient_nsacl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl_binding.Type(), &lsnclient_nsacl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_nsacl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnclient_nsacl_binding resource")

	// Read the updated state back
	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsaclBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_nsacl_binding resource")

	// For lsnclient_nsacl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnclient_nsacl_binding resource from state")
}

// Helper function to read lsnclient_nsacl_binding data from API
func (r *LsnclientNsaclBindingResource) readLsnclientNsaclBindingFromApi(ctx context.Context, data *LsnclientNsaclBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnclient_nsacl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_nsacl_binding, got error: %s", err))
		return
	}

	lsnclient_nsacl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
