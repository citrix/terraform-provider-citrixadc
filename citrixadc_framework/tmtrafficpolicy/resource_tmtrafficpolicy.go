package tmtrafficpolicy

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
var _ resource.Resource = &TmtrafficpolicyResource{}
var _ resource.ResourceWithConfigure = (*TmtrafficpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TmtrafficpolicyResource)(nil)

func NewTmtrafficpolicyResource() resource.Resource {
	return &TmtrafficpolicyResource{}
}

// TmtrafficpolicyResource defines the resource implementation.
type TmtrafficpolicyResource struct {
	client *service.NitroClient
}

func (r *TmtrafficpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmtrafficpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmtrafficpolicy"
}

func (r *TmtrafficpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmtrafficpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmtrafficpolicy resource")

	// tmtrafficpolicy := tmtrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmtrafficpolicy.Type(), &tmtrafficpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmtrafficpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmtrafficpolicy-config")

	tflog.Trace(ctx, "Created tmtrafficpolicy resource")

	// Read the updated state back
	r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmtrafficpolicy resource")

	r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmtrafficpolicy resource")

	// Create API request body from the model
	// tmtrafficpolicy := tmtrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmtrafficpolicy.Type(), &tmtrafficpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmtrafficpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmtrafficpolicy resource")

	// Read the updated state back
	r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmtrafficpolicy resource")

	// For tmtrafficpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmtrafficpolicy resource from state")
}

// Helper function to read tmtrafficpolicy data from API
func (r *TmtrafficpolicyResource) readTmtrafficpolicyFromApi(ctx context.Context, data *TmtrafficpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmtrafficpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmtrafficpolicy, got error: %s", err))
		return
	}

	tmtrafficpolicySetAttrFromGet(ctx, data, getResponseData)

}
