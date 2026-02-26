package lsngroup_lsnappsprofile_binding

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
var _ resource.Resource = &LsngroupLsnappsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnappsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnappsprofileBindingResource)(nil)

func NewLsngroupLsnappsprofileBindingResource() resource.Resource {
	return &LsngroupLsnappsprofileBindingResource{}
}

// LsngroupLsnappsprofileBindingResource defines the resource implementation.
type LsngroupLsnappsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnappsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnappsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnappsprofile_binding"
}

func (r *LsngroupLsnappsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnappsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnappsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnappsprofile_binding resource")

	// lsngroup_lsnappsprofile_binding := lsngroup_lsnappsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnappsprofile_binding.Type(), &lsngroup_lsnappsprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnappsprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsngroup_lsnappsprofile_binding-config")

	tflog.Trace(ctx, "Created lsngroup_lsnappsprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnappsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnappsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnappsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnappsprofile_binding resource")

	r.readLsngroupLsnappsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnappsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsngroupLsnappsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsngroup_lsnappsprofile_binding resource")

	// Create API request body from the model
	// lsngroup_lsnappsprofile_binding := lsngroup_lsnappsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnappsprofile_binding.Type(), &lsngroup_lsnappsprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsnappsprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsngroup_lsnappsprofile_binding resource")

	// Read the updated state back
	r.readLsngroupLsnappsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnappsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnappsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnappsprofile_binding resource")

	// For lsngroup_lsnappsprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsngroup_lsnappsprofile_binding resource from state")
}

// Helper function to read lsngroup_lsnappsprofile_binding data from API
func (r *LsngroupLsnappsprofileBindingResource) readLsngroupLsnappsprofileBindingFromApi(ctx context.Context, data *LsngroupLsnappsprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsngroup_lsnappsprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnappsprofile_binding, got error: %s", err))
		return
	}

	lsngroup_lsnappsprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
