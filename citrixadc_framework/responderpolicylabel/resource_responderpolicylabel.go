package responderpolicylabel

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
var _ resource.Resource = &ResponderpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*ResponderpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderpolicylabelResource)(nil)

func NewResponderpolicylabelResource() resource.Resource {
	return &ResponderpolicylabelResource{}
}

// ResponderpolicylabelResource defines the resource implementation.
type ResponderpolicylabelResource struct {
	client *service.NitroClient
}

func (r *ResponderpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderpolicylabel"
}

func (r *ResponderpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderpolicylabel resource")

	// responderpolicylabel := responderpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderpolicylabel.Type(), &responderpolicylabel)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderpolicylabel, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("responderpolicylabel-config")

	tflog.Trace(ctx, "Created responderpolicylabel resource")

	// Read the updated state back
	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderpolicylabel resource")

	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating responderpolicylabel resource")

	// Create API request body from the model
	// responderpolicylabel := responderpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderpolicylabel.Type(), &responderpolicylabel)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderpolicylabel, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated responderpolicylabel resource")

	// Read the updated state back
	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderpolicylabel resource")

	// For responderpolicylabel, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted responderpolicylabel resource from state")
}

// Helper function to read responderpolicylabel data from API
func (r *ResponderpolicylabelResource) readResponderpolicylabelFromApi(ctx context.Context, data *ResponderpolicylabelResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Responderpolicylabel.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderpolicylabel, got error: %s", err))
		return
	}

	responderpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
