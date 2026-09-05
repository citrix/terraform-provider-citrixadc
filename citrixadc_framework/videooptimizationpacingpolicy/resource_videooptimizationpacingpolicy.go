package videooptimizationpacingpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VideooptimizationpacingpolicyResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationpacingpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationpacingpolicyResource)(nil)

func NewVideooptimizationpacingpolicyResource() resource.Resource {
	return &VideooptimizationpacingpolicyResource{}
}

// VideooptimizationpacingpolicyResource defines the resource implementation.
type VideooptimizationpacingpolicyResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationpacingpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationpacingpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationpacingpolicy"
}

func (r *VideooptimizationpacingpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationpacingpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationpacingpolicy resource")

	videooptimizationpacingpolicy := videooptimizationpacingpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationpacingpolicy.Type(), name_value, &videooptimizationpacingpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationpacingpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationpacingpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "videooptimizationpacingpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationpacingpolicy resource")

	found := r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VideooptimizationpacingpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VideooptimizationpacingpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset candidates)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating videooptimizationpacingpolicy resource")

	// Rename support: a newname change is an in-place rename via ?action=rename, not
	// a destroy/recreate. The rename SOURCE is the current live name (tracked by the
	// ID), NOT state.Name, which stays pinned to the originally configured value.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming videooptimizationpacingpolicy from %q to %q", oldName, newName))

		renamePayload := videooptimization.Videooptimizationpacingpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Videooptimizationpacingpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename videooptimizationpacingpolicy, got error: %s", err))
			return
		}
		// The live object is now named newName. Point the ID at it.
		data.Id = types.StringValue(newName)
	}

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for videooptimizationpacingpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for videooptimizationpacingpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for videooptimizationpacingpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for videooptimizationpacingpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for videooptimizationpacingpolicy")
		hasChange = true
	}

	if hasChange {
		// Named resource - use UpdateResource against the current live name.
		videooptimizationpacingpolicy := videooptimizationpacingpolicyGetThePayloadFromtheConfig(ctx, &data)
		videooptimizationpacingpolicy.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Videooptimizationpacingpolicy.Type(), data.Id.ValueString(), &videooptimizationpacingpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationpacingpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated videooptimizationpacingpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for videooptimizationpacingpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Done against the current live name (ID) to survive renames.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Videooptimizationpacingpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset videooptimizationpacingpolicy attributes, got error: %s", err))
		return
	}

	// Preserve the user-facing name/newname across the read-back so a rename does not
	// produce a spurious diff on the configured key.
	planName := data.Name
	planNewname := data.Newname

	// Read the updated state back
	if !r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "videooptimizationpacingpolicy not found immediately after update")
		}
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationpacingpolicy resource")

	// Named resource - delete using DeleteResource by the live name (ID).
	err := r.client.DeleteResource(service.Videooptimizationpacingpolicy.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationpacingpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationpacingpolicy resource")
}

// Helper function to read videooptimizationpacingpolicy data from API
func (r *VideooptimizationpacingpolicyResource) readVideooptimizationpacingpolicyFromApi(ctx context.Context, data *VideooptimizationpacingpolicyResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain (live) name value
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Videooptimizationpacingpolicy.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationpacingpolicy, got error: %s", err))
		return false
	}

	videooptimizationpacingpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
