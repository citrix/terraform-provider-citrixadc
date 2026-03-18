package lsnappsprofile_lsnappsattributes_binding

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
var _ resource.Resource = &LsnappsprofileLsnappsattributesBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnappsprofileLsnappsattributesBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnappsprofileLsnappsattributesBindingResource)(nil)

func NewLsnappsprofileLsnappsattributesBindingResource() resource.Resource {
	return &LsnappsprofileLsnappsattributesBindingResource{}
}

// LsnappsprofileLsnappsattributesBindingResource defines the resource implementation.
type LsnappsprofileLsnappsattributesBindingResource struct {
	client *service.NitroClient
}

func (r *LsnappsprofileLsnappsattributesBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile_lsnappsattributes_binding"
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnappsprofileLsnappsattributesBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnappsprofile_lsnappsattributes_binding resource")

	// lsnappsprofile_lsnappsattributes_binding := lsnappsprofile_lsnappsattributes_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_lsnappsattributes_binding.Type(), &lsnappsprofile_lsnappsattributes_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnappsprofile_lsnappsattributes_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnappsprofile_lsnappsattributes_binding-config")

	tflog.Trace(ctx, "Created lsnappsprofile_lsnappsattributes_binding resource")

	// Read the updated state back
	r.readLsnappsprofileLsnappsattributesBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnappsprofileLsnappsattributesBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnappsprofile_lsnappsattributes_binding resource")

	r.readLsnappsprofileLsnappsattributesBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnappsprofileLsnappsattributesBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnappsprofile_lsnappsattributes_binding resource")

	// Create API request body from the model
	// lsnappsprofile_lsnappsattributes_binding := lsnappsprofile_lsnappsattributes_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_lsnappsattributes_binding.Type(), &lsnappsprofile_lsnappsattributes_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnappsprofile_lsnappsattributes_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnappsprofile_lsnappsattributes_binding resource")

	// Read the updated state back
	r.readLsnappsprofileLsnappsattributesBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnappsprofileLsnappsattributesBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnappsprofile_lsnappsattributes_binding resource")

	// For lsnappsprofile_lsnappsattributes_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnappsprofile_lsnappsattributes_binding resource from state")
}

// Helper function to read lsnappsprofile_lsnappsattributes_binding data from API
func (r *LsnappsprofileLsnappsattributesBindingResource) readLsnappsprofileLsnappsattributesBindingFromApi(ctx context.Context, data *LsnappsprofileLsnappsattributesBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnappsprofile_lsnappsattributes_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile_lsnappsattributes_binding, got error: %s", err))
		return
	}

	lsnappsprofile_lsnappsattributes_bindingSetAttrFromGet(ctx, data, getResponseData)

}
