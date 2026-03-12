package lsnappsprofile_port_binding

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
var _ resource.Resource = &LsnappsprofilePortBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnappsprofilePortBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnappsprofilePortBindingResource)(nil)

func NewLsnappsprofilePortBindingResource() resource.Resource {
	return &LsnappsprofilePortBindingResource{}
}

// LsnappsprofilePortBindingResource defines the resource implementation.
type LsnappsprofilePortBindingResource struct {
	client *service.NitroClient
}

func (r *LsnappsprofilePortBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnappsprofilePortBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile_port_binding"
}

func (r *LsnappsprofilePortBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnappsprofilePortBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnappsprofile_port_binding resource")

	// lsnappsprofile_port_binding := lsnappsprofile_port_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_port_binding.Type(), &lsnappsprofile_port_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnappsprofile_port_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnappsprofile_port_binding-config")

	tflog.Trace(ctx, "Created lsnappsprofile_port_binding resource")

	// Read the updated state back
	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnappsprofile_port_binding resource")

	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnappsprofile_port_binding resource")

	// Create API request body from the model
	// lsnappsprofile_port_binding := lsnappsprofile_port_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_port_binding.Type(), &lsnappsprofile_port_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnappsprofile_port_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnappsprofile_port_binding resource")

	// Read the updated state back
	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnappsprofile_port_binding resource")

	// For lsnappsprofile_port_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnappsprofile_port_binding resource from state")
}

// Helper function to read lsnappsprofile_port_binding data from API
func (r *LsnappsprofilePortBindingResource) readLsnappsprofilePortBindingFromApi(ctx context.Context, data *LsnappsprofilePortBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnappsprofile_port_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile_port_binding, got error: %s", err))
		return
	}

	lsnappsprofile_port_bindingSetAttrFromGet(ctx, data, getResponseData)

}
