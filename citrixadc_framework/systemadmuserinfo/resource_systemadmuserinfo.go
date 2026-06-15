package systemadmuserinfo

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemadmuserinfoResource{}
var _ resource.ResourceWithConfigure = (*SystemadmuserinfoResource)(nil)
var _ resource.ResourceWithImportState = (*SystemadmuserinfoResource)(nil)

func NewSystemadmuserinfoResource() resource.Resource {
	return &SystemadmuserinfoResource{}
}

// SystemadmuserinfoResource defines the resource implementation.
type SystemadmuserinfoResource struct {
	client *service.NitroClient
}

func (r *SystemadmuserinfoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemadmuserinfoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemadmuserinfo"
}

func (r *SystemadmuserinfoResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemadmuserinfoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemadmuserinfoResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemadmuserinfo resource")
	systemadmuserinfo := systemadmuserinfoGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Update-only resource - NITRO exposes only `update` (PUT), no `add`
	_, err := r.client.UpdateResource(service.Systemadmuserinfo.Type(), "", &systemadmuserinfo)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemadmuserinfo, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemadmuserinfo resource")

	// Set ID for the resource
	data.Id = types.StringValue("systemadmuserinfo-config")

	// No GET endpoint on NITRO side - state is taken directly from the plan

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemadmuserinfoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemadmuserinfoResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for systemadmuserinfo: NITRO exposes no GET endpoint
	// (only `update`), so drift detection is impossible. Preserve prior state.
	tflog.Debug(ctx, "Read is a no-op for systemadmuserinfo; NITRO has no GET endpoint")

	// Save prior state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemadmuserinfoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemadmuserinfoResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemadmuserinfo resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, "username has changed for systemadmuserinfo")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		systemadmuserinfo := systemadmuserinfoGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Update-only resource - NITRO exposes only `update` (PUT)
		_, err := r.client.UpdateResource(service.Systemadmuserinfo.Type(), "", &systemadmuserinfo)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemadmuserinfo, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemadmuserinfo resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemadmuserinfo resource, skipping update")
	}

	// No GET endpoint on NITRO side - state is taken directly from the plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemadmuserinfoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemadmuserinfoResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemadmuserinfo resource")
	// No delete operation on ADC - just remove from state
	tflog.Trace(ctx, "Removed systemadmuserinfo from Terraform state")
}
