package lsngroup_lsnlogprofile_binding

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
var _ resource.Resource = &LsngroupLsnlogprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnlogprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnlogprofileBindingResource)(nil)

func NewLsngroupLsnlogprofileBindingResource() resource.Resource {
	return &LsngroupLsnlogprofileBindingResource{}
}

// LsngroupLsnlogprofileBindingResource defines the resource implementation.
type LsngroupLsnlogprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnlogprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnlogprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnlogprofile_binding"
}

func (r *LsngroupLsnlogprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnlogprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnlogprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnlogprofile_binding resource")

	// lsngroup_lsnlogprofile_binding := lsngroup_lsnlogprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnlogprofile_binding.Type(), &lsngroup_lsnlogprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnlogprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsngroup_lsnlogprofile_binding-config")

	tflog.Trace(ctx, "Created lsngroup_lsnlogprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnlogprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnlogprofile_binding resource")

	r.readLsngroupLsnlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnlogprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsngroupLsnlogprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsngroup_lsnlogprofile_binding resource")

	// Create API request body from the model
	// lsngroup_lsnlogprofile_binding := lsngroup_lsnlogprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnlogprofile_binding.Type(), &lsngroup_lsnlogprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsnlogprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsngroup_lsnlogprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnlogprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnlogprofile_binding resource")

	// For lsngroup_lsnlogprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsngroup_lsnlogprofile_binding resource from state")
}

// Helper function to read lsngroup_lsnlogprofile_binding data from API
func (r *LsngroupLsnlogprofileBindingResource) readLsngroupLsnlogprofileBindingFromApi(ctx context.Context, data *LsngroupLsnlogprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsngroup_lsnlogprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnlogprofile_binding, got error: %s", err))
		return
	}

	lsngroup_lsnlogprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
