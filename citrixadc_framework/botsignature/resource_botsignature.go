package botsignature

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BotsignatureResource{}
var _ resource.ResourceWithConfigure = (*BotsignatureResource)(nil)
var _ resource.ResourceWithImportState = (*BotsignatureResource)(nil)

func NewBotsignatureResource() resource.Resource {
	return &BotsignatureResource{}
}

// BotsignatureResource defines the resource implementation.
type BotsignatureResource struct {
	client *service.NitroClient
}

func (r *BotsignatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotsignatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botsignature"
}

func (r *BotsignatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotsignatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotsignatureResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botsignature resource")

	botsignature := botsignatureGetThePayloadFromthePlan(ctx, &data)

	// Named resource imported via the NITRO "Import" action (POST ?action=Import).
	// This mirrors the SDK v2 resource which called ActOnResource(..., "Import").
	err := r.client.ActOnResource(service.Botsignature.Type(), &botsignature, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botsignature, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botsignature resource")

	// Set ID for the resource before reading state.
	// Case 2: Single unique attribute - use plain name value as ID.
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botsignature not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotsignatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botsignature resource")

	found := r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotsignatureResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// botsignature has no NITRO-updatable attributes: every configurable attribute
	// (name, comment, overwrite, src) is ForceNew/RequiresReplace, exactly as in the
	// SDK v2 resource which defined no update. Terraform therefore recreates on any
	// change, so this path only refreshes state.
	tflog.Debug(ctx, "Updating botsignature resource - no updatable attributes, refreshing state")

	// Read the updated state back
	if !r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botsignature not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotsignatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botsignature resource")

	// Named resource - delete using DeleteResource keyed on the name (ID).
	name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Botsignature.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botsignature, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botsignature resource")
}

// Helper function to read botsignature data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *BotsignatureResource) readBotsignatureFromApi(ctx context.Context, data *BotsignatureResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Botsignature.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botsignature, got error: %s", err))
		return false
	}

	botsignatureSetAttrFromGet(ctx, data, getResponseData)

	return true
}
