package spilloveraction

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
var _ resource.Resource = &SpilloveractionResource{}
var _ resource.ResourceWithConfigure = (*SpilloveractionResource)(nil)
var _ resource.ResourceWithImportState = (*SpilloveractionResource)(nil)

func NewSpilloveractionResource() resource.Resource {
	return &SpilloveractionResource{}
}

// SpilloveractionResource defines the resource implementation.
type SpilloveractionResource struct {
	client *service.NitroClient
}

func (r *SpilloveractionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SpilloveractionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spilloveraction"
}

func (r *SpilloveractionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SpilloveractionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating spilloveraction resource")

	// spilloveraction := spilloveractionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Spilloveraction.Type(), &spilloveraction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create spilloveraction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("spilloveraction-config")

	tflog.Trace(ctx, "Created spilloveraction resource")

	// Read the updated state back
	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading spilloveraction resource")

	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating spilloveraction resource")

	// Create API request body from the model
	// spilloveraction := spilloveractionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Spilloveraction.Type(), &spilloveraction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update spilloveraction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated spilloveraction resource")

	// Read the updated state back
	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting spilloveraction resource")

	// For spilloveraction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted spilloveraction resource from state")
}

// Helper function to read spilloveraction data from API
func (r *SpilloveractionResource) readSpilloveractionFromApi(ctx context.Context, data *SpilloveractionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Spilloveraction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read spilloveraction, got error: %s", err))
		return
	}

	spilloveractionSetAttrFromGet(ctx, data, getResponseData)

}
