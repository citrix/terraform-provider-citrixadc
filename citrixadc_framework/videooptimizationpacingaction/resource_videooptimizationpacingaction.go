package videooptimizationpacingaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VideooptimizationpacingactionResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationpacingactionResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationpacingactionResource)(nil)

func NewVideooptimizationpacingactionResource() resource.Resource {
	return &VideooptimizationpacingactionResource{}
}

// VideooptimizationpacingactionResource defines the resource implementation.
type VideooptimizationpacingactionResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationpacingactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationpacingactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationpacingaction"
}

func (r *VideooptimizationpacingactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationpacingactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationpacingaction resource")
	videooptimizationpacingaction := videooptimizationpacingactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationpacingaction.Type(), name_value, &videooptimizationpacingaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationpacingaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationpacingaction resource")

	// Set ID for the resource before reading state (single unique attribute - plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationpacingaction resource")

	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VideooptimizationpacingactionResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to distinguish "changed value" from "removed" (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating videooptimizationpacingaction resource")

	// Rename support: the primary key `name` is RequiresReplace, so a name change is
	// a recreate and never reaches Update. The in-place rename NITRO offers is the
	// `rename` action, triggered by a `newname` change. The rename SOURCE is the
	// CURRENT LIVE name, tracked by state.Id (NOT state.Name, which stays pinned to
	// the originally configured value and would be wrong on a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming videooptimizationpacingaction from %q to %q", oldName, newName))

		renamePayload := videooptimization.Videooptimizationpacingaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Videooptimizationpacingaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename videooptimizationpacingaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update /
		// read below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// In-place update of the updateable attributes (comment, rate). Mirrors the SDK v2
	// HasChange gate so a no-op plan does not issue a NITRO write.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for videooptimizationpacingaction")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Rate.Equal(state.Rate) {
		tflog.Debug(ctx, "rate has changed for videooptimizationpacingaction")
		hasChange = true
	}

	if hasChange {
		videooptimizationpacingaction := videooptimizationpacingactionGetThePayloadFromthePlan(ctx, &data)
		// Target the current live name (== newname after a rename in this same apply).
		liveName := data.Id.ValueString()
		videooptimizationpacingaction.Name = liveName
		_, err := r.client.UpdateResource(service.Videooptimizationpacingaction.Type(), liveName, &videooptimizationpacingaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationpacingaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated videooptimizationpacingaction resource")
	} else {
		tflog.Debug(ctx, "No updateable changes detected for videooptimizationpacingaction resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// defaults. Target the current live name (== newname after a rename this apply).
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Videooptimizationpacingaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset videooptimizationpacingaction attributes, got error: %s", err))
		return
	}

	// Read the current state back. Preserve the user-facing name / newname (the object
	// may be physically named newName now, but the plan keeps the configured values).
	planName := data.Name
	planNewname := data.Newname
	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationpacingaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would dangle the
	// object after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Videooptimizationpacingaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationpacingaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationpacingaction resource")
}

// Helper function to read videooptimizationpacingaction data from API
func (r *VideooptimizationpacingactionResource) readVideooptimizationpacingactionFromApi(ctx context.Context, data *VideooptimizationpacingactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Videooptimizationpacingaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationpacingaction, got error: %s", err))
		return
	}

	videooptimizationpacingactionSetAttrFromGet(ctx, data, getResponseData)
}
