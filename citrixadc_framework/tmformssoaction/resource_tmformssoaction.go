package tmformssoaction

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
var _ resource.Resource = &TmformssoactionResource{}
var _ resource.ResourceWithConfigure = (*TmformssoactionResource)(nil)
var _ resource.ResourceWithImportState = (*TmformssoactionResource)(nil)

func NewTmformssoactionResource() resource.Resource {
	return &TmformssoactionResource{}
}

// TmformssoactionResource defines the resource implementation.
type TmformssoactionResource struct {
	client *service.NitroClient
}

func (r *TmformssoactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmformssoactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmformssoaction"
}

func (r *TmformssoactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmformssoactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmformssoaction resource")

	// tmformssoaction := tmformssoactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmformssoaction.Type(), &tmformssoaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmformssoaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmformssoaction-config")

	tflog.Trace(ctx, "Created tmformssoaction resource")

	// Read the updated state back
	r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmformssoactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmformssoaction resource")

	r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmformssoactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmformssoaction resource")

	// Create API request body from the model
	// tmformssoaction := tmformssoactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmformssoaction.Type(), &tmformssoaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmformssoaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmformssoaction resource")

	// Read the updated state back
	r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmformssoactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmformssoaction resource")

	// For tmformssoaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmformssoaction resource from state")
}

// Helper function to read tmformssoaction data from API
func (r *TmformssoactionResource) readTmformssoactionFromApi(ctx context.Context, data *TmformssoactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmformssoaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmformssoaction, got error: %s", err))
		return
	}

	tmformssoactionSetAttrFromGet(ctx, data, getResponseData)

}
