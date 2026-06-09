package hafiles

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
var _ resource.Resource = &HafilesResource{}
var _ resource.ResourceWithConfigure = (*HafilesResource)(nil)
var _ resource.ResourceWithImportState = (*HafilesResource)(nil)

func NewHafilesResource() resource.Resource {
	return &HafilesResource{}
}

// HafilesResource defines the resource implementation.
type HafilesResource struct {
	client *service.NitroClient
}

func (r *HafilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HafilesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hafiles"
}

func (r *HafilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HafilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HafilesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hafiles resource")

	// hafiles exposes only the POST ?action=sync action on NITRO.
	// There is no add/get/update/delete endpoint. Use ActOnResource with
	// the "sync" verb.
	payload := hafilesGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Hafiles.Type(), payload, "sync")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync hafiles, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Synced hafiles resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("hafiles")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HafilesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for hafiles: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for hafiles; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HafilesResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for hafiles; the only attribute (mode) is
	// RequiresReplace, so Terraform re-creates (re-syncs) on change instead.
	tflog.Debug(ctx, "Update is a no-op for hafiles; all attributes are RequiresReplace")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HafilesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for hafiles: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for hafiles; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed hafiles from Terraform state")
}
