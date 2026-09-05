package spilloverpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SpilloverpolicyResource{}
var _ resource.ResourceWithConfigure = (*SpilloverpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*SpilloverpolicyResource)(nil)

func NewSpilloverpolicyResource() resource.Resource {
	return &SpilloverpolicyResource{}
}

// SpilloverpolicyResource defines the resource implementation.
type SpilloverpolicyResource struct {
	client *service.NitroClient
}

func (r *SpilloverpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SpilloverpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spilloverpolicy"
}

func (r *SpilloverpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SpilloverpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SpilloverpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating spilloverpolicy resource")
	spilloverpolicy := spilloverpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Spilloverpolicy.Type(), name_value, &spilloverpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create spilloverpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created spilloverpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readSpilloverpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloverpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SpilloverpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading spilloverpolicy resource")

	r.readSpilloverpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SpilloverpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SpilloverpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID (the current live name) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating spilloverpolicy resource")

	// Rename support: on a newname change, POST {name, newname} to ?action=rename,
	// then point the resource ID at the new name so subsequent reads/updates address
	// the live object. name itself is RequiresReplace, so a name change never reaches
	// Update - only newname (and the updateable fields below) do.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and would
		// point at a no-longer-live name on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming spilloverpolicy from %q to %q", oldName, newName))

		renamePayload := spillover.Spilloverpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Spilloverpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename spilloverpolicy, got error: %s", err))
			return
		}
		// The live object is now named newName. Point the ID at it.
		data.Id = types.StringValue(newName)
	}

	// In-place update of the updateable fields (action, comment, rule). These are NOT
	// ForceNew in SDK v2, so a change must update in place rather than recreate.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for spilloverpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for spilloverpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for spilloverpolicy")
		hasChange = true
	}

	if hasChange {
		spilloverpolicy := spilloverpolicyGetThePayloadFromthePlan(ctx, &data)
		// Key the update on the CURRENT LIVE name (data.Id), which reflects any rename
		// performed above.
		liveName := data.Id.ValueString()
		spilloverpolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Spilloverpolicy.Type(), liveName, &spilloverpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update spilloverpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated spilloverpolicy resource")
	} else {
		tflog.Debug(ctx, "No updateable changes detected for spilloverpolicy resource")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// defaults. Keyed on the CURRENT LIVE name (data.Id), reflecting any rename above.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Spilloverpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset spilloverpolicy attributes, got error: %s", err))
		return
	}

	// Read the current state back. Preserve the plan's user-facing name/newname across
	// the read so a post-rename GET (which returns the live/new name) cannot clobber
	// the configured values and cause an inconsistent-result / perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readSpilloverpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloverpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SpilloverpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting spilloverpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which is stale after a rename and would dangle the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Spilloverpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete spilloverpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted spilloverpolicy resource")
}

// Helper function to read spilloverpolicy data from API
func (r *SpilloverpolicyResource) readSpilloverpolicyFromApi(ctx context.Context, data *SpilloverpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value.
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Spilloverpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read spilloverpolicy, got error: %s", err))
		return
	}

	spilloverpolicySetAttrFromGet(ctx, data, getResponseData)
}
