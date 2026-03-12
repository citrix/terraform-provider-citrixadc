package lsngroup_lsnhttphdrlogprofile_binding

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
var _ resource.Resource = &LsngroupLsnhttphdrlogprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnhttphdrlogprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnhttphdrlogprofileBindingResource)(nil)

func NewLsngroupLsnhttphdrlogprofileBindingResource() resource.Resource {
	return &LsngroupLsnhttphdrlogprofileBindingResource{}
}

// LsngroupLsnhttphdrlogprofileBindingResource defines the resource implementation.
type LsngroupLsnhttphdrlogprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnhttphdrlogprofile_binding"
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnhttphdrlogprofile_binding resource")

	// lsngroup_lsnhttphdrlogprofile_binding := lsngroup_lsnhttphdrlogprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), &lsngroup_lsnhttphdrlogprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsngroup_lsnhttphdrlogprofile_binding-config")

	tflog.Trace(ctx, "Created lsngroup_lsnhttphdrlogprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnhttphdrlogprofile_binding resource")

	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsngroup_lsnhttphdrlogprofile_binding resource")

	// Create API request body from the model
	// lsngroup_lsnhttphdrlogprofile_binding := lsngroup_lsnhttphdrlogprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), &lsngroup_lsnhttphdrlogprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsngroup_lsnhttphdrlogprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnhttphdrlogprofile_binding resource")

	// For lsngroup_lsnhttphdrlogprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsngroup_lsnhttphdrlogprofile_binding resource from state")
}

// Helper function to read lsngroup_lsnhttphdrlogprofile_binding data from API
func (r *LsngroupLsnhttphdrlogprofileBindingResource) readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx context.Context, data *LsngroupLsnhttphdrlogprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
		return
	}

	lsngroup_lsnhttphdrlogprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
