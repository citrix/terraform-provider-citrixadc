package lsnsipalgprofile

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
var _ resource.Resource = &LsnsipalgprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnsipalgprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsipalgprofileResource)(nil)

func NewLsnsipalgprofileResource() resource.Resource {
	return &LsnsipalgprofileResource{}
}

// LsnsipalgprofileResource defines the resource implementation.
type LsnsipalgprofileResource struct {
	client *service.NitroClient
}

func (r *LsnsipalgprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsipalgprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsipalgprofile"
}

func (r *LsnsipalgprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnsipalgprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnsipalgprofile resource")

	// lsnsipalgprofile := lsnsipalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnsipalgprofile.Type(), &lsnsipalgprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnsipalgprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnsipalgprofile-config")

	tflog.Trace(ctx, "Created lsnsipalgprofile resource")

	// Read the updated state back
	r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsipalgprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnsipalgprofile resource")

	r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsipalgprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnsipalgprofile resource")

	// Create API request body from the model
	// lsnsipalgprofile := lsnsipalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnsipalgprofile.Type(), &lsnsipalgprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnsipalgprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnsipalgprofile resource")

	// Read the updated state back
	r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsipalgprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnsipalgprofile resource")

	// For lsnsipalgprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnsipalgprofile resource from state")
}

// Helper function to read lsnsipalgprofile data from API
func (r *LsnsipalgprofileResource) readLsnsipalgprofileFromApi(ctx context.Context, data *LsnsipalgprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnsipalgprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnsipalgprofile, got error: %s", err))
		return
	}

	lsnsipalgprofileSetAttrFromGet(ctx, data, getResponseData)

}
