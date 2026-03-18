package lsnclient_nsacl6_binding

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
var _ resource.Resource = &LsnclientNsacl6BindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNsacl6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNsacl6BindingResource)(nil)

func NewLsnclientNsacl6BindingResource() resource.Resource {
	return &LsnclientNsacl6BindingResource{}
}

// LsnclientNsacl6BindingResource defines the resource implementation.
type LsnclientNsacl6BindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNsacl6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNsacl6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_nsacl6_binding"
}

func (r *LsnclientNsacl6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNsacl6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_nsacl6_binding resource")

	// lsnclient_nsacl6_binding := lsnclient_nsacl6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl6_binding.Type(), &lsnclient_nsacl6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_nsacl6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnclient_nsacl6_binding-config")

	tflog.Trace(ctx, "Created lsnclient_nsacl6_binding resource")

	// Read the updated state back
	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_nsacl6_binding resource")

	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnclient_nsacl6_binding resource")

	// Create API request body from the model
	// lsnclient_nsacl6_binding := lsnclient_nsacl6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl6_binding.Type(), &lsnclient_nsacl6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_nsacl6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnclient_nsacl6_binding resource")

	// Read the updated state back
	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_nsacl6_binding resource")

	// For lsnclient_nsacl6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnclient_nsacl6_binding resource from state")
}

// Helper function to read lsnclient_nsacl6_binding data from API
func (r *LsnclientNsacl6BindingResource) readLsnclientNsacl6BindingFromApi(ctx context.Context, data *LsnclientNsacl6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnclient_nsacl6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_nsacl6_binding, got error: %s", err))
		return
	}

	lsnclient_nsacl6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
