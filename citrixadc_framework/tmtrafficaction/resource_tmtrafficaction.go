package tmtrafficaction

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
var _ resource.Resource = &TmtrafficactionResource{}
var _ resource.ResourceWithConfigure = (*TmtrafficactionResource)(nil)
var _ resource.ResourceWithImportState = (*TmtrafficactionResource)(nil)

func NewTmtrafficactionResource() resource.Resource {
	return &TmtrafficactionResource{}
}

// TmtrafficactionResource defines the resource implementation.
type TmtrafficactionResource struct {
	client *service.NitroClient
}

func (r *TmtrafficactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmtrafficactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmtrafficaction"
}

func (r *TmtrafficactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmtrafficactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmtrafficaction resource")

	// tmtrafficaction := tmtrafficactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmtrafficaction.Type(), &tmtrafficaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmtrafficaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmtrafficaction-config")

	tflog.Trace(ctx, "Created tmtrafficaction resource")

	// Read the updated state back
	r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmtrafficaction resource")

	r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmtrafficaction resource")

	// Create API request body from the model
	// tmtrafficaction := tmtrafficactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmtrafficaction.Type(), &tmtrafficaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmtrafficaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmtrafficaction resource")

	// Read the updated state back
	r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmtrafficaction resource")

	// For tmtrafficaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmtrafficaction resource from state")
}

// Helper function to read tmtrafficaction data from API
func (r *TmtrafficactionResource) readTmtrafficactionFromApi(ctx context.Context, data *TmtrafficactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmtrafficaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmtrafficaction, got error: %s", err))
		return
	}

	tmtrafficactionSetAttrFromGet(ctx, data, getResponseData)

}
