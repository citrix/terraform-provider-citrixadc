package videooptimizationdetectionpolicy

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
var _ resource.Resource = &VideooptimizationdetectionpolicyResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionpolicyResource)(nil)

func NewVideooptimizationdetectionpolicyResource() resource.Resource {
	return &VideooptimizationdetectionpolicyResource{}
}

// VideooptimizationdetectionpolicyResource defines the resource implementation.
type VideooptimizationdetectionpolicyResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicy"
}

func (r *VideooptimizationdetectionpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionpolicy resource")
	videooptimizationdetectionpolicy := videooptimizationdetectionpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationdetectionpolicy.Type(), name_value, &videooptimizationdetectionpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationdetectionpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "videooptimizationdetectionpolicy not found immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionpolicy resource")

	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VideooptimizationdetectionpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (== live name, tracked across renames)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating videooptimizationdetectionpolicy resource")

	// Rename support: on a newname change, POST {name, newname} to ?action=rename, then
	// point the resource ID at the new name so subsequent reads address the live object.
	// The rename SOURCE is the CURRENT LIVE name (state.Id), NOT state.Name (which stays
	// pinned to the originally configured value and would be stale on a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming videooptimizationdetectionpolicy from %q to %q", oldName, newName))

		renamePayload := videooptimization.Videooptimizationdetectionpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Videooptimizationdetectionpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename videooptimizationdetectionpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update/read
		// below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular in-place update of the mutable attributes. name uses RequiresReplace and
	// newname is handled above, so only action/rule/comment/logaction/undefaction land
	// here. NITRO exposes an update (set) endpoint for this resource, so use UpdateResource.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for videooptimizationdetectionpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for videooptimizationdetectionpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for videooptimizationdetectionpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for videooptimizationdetectionpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for videooptimizationdetectionpolicy")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		videooptimizationdetectionpolicy := videooptimizationdetectionpolicyGetThePayloadFromthePlan(ctx, &data)
		// The live object name is tracked by data.Id (== newname after a rename above).
		videooptimizationdetectionpolicy.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Videooptimizationdetectionpolicy.Type(), data.Id.ValueString(), &videooptimizationdetectionpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationdetectionpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated videooptimizationdetectionpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for videooptimizationdetectionpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. The live object name is tracked by data.Id
	// (== newname after a rename above), so address the unset by data.Id.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Videooptimizationdetectionpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset videooptimizationdetectionpolicy attributes, got error: %s", err))
		return
	}

	// Read the current state back. Capture the user-facing plan values (name/newname) and
	// restore them after the read to avoid an inconsistent-result / perpetual diff, since
	// after a rename GET returns the live (new) name while name stays the configured value.
	planName := data.Name
	planNewname := data.Newname
	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT data.Name
	// (which stays at the originally configured value and would dangle a renamed object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Videooptimizationdetectionpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationdetectionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationdetectionpolicy resource")
}

// Helper function to read videooptimizationdetectionpolicy data from API
func (r *VideooptimizationdetectionpolicyResource) readVideooptimizationdetectionpolicyFromApi(ctx context.Context, data *VideooptimizationdetectionpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Videooptimizationdetectionpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicy, got error: %s", err))
		return
	}

	videooptimizationdetectionpolicySetAttrFromGet(ctx, data, getResponseData)
}
