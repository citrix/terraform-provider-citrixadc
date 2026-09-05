package responderaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ResponderactionResource{}
var _ resource.ResourceWithConfigure = (*ResponderactionResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderactionResource)(nil)

func NewResponderactionResource() resource.Resource {
	return &ResponderactionResource{}
}

// ResponderactionResource defines the resource implementation.
type ResponderactionResource struct {
	client *service.NitroClient
}

func (r *ResponderactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderaction"
}

func (r *ResponderactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderaction resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the user
	// does not supply a name, generate a unique one.
	responderactionName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || responderactionName == "" {
		responderactionName = sdkid.PrefixedUniqueId("tf-responderaction-")
		data.Name = types.StringValue(responderactionName)
	}

	// Create API request body from the model
	responderaction := responderactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Responderaction.Type(), responderactionName, &responderaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created responderaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(responderactionName)

	// Read the updated state back
	if !r.readResponderactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "responderaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderaction resource")

	found := r.readResponderactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ResponderactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ResponderactionResourceModel

	// Read Terraform prior state to preserve ID / live name
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating responderaction resource")

	// Rename support: NITRO exposes a `rename` action for responderaction. A newname
	// change drives an in-place rename (POST ?action=rename), never a destroy/recreate.
	// name and type are ForceNew (RequiresReplaceIfConfigured) so they never reach here.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by data.Id (== name at
		// create, == the prior newname after one rename), NOT data.Name (which stays
		// pinned to the originally configured value).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming responderaction from %q to %q", oldName, newName))

		renamePayload := responder.Responderaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Responderaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename responderaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update /
		// read-back below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular (in-place) update of the NITRO-updatable attributes.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Bypasssafetycheck.Equal(state.Bypasssafetycheck) {
		tflog.Debug(ctx, "bypasssafetycheck has changed for responderaction")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for responderaction")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Headers.Equal(state.Headers) {
		tflog.Debug(ctx, "headers has changed for responderaction")
		hasChange = true
	}
	if !data.Htmlpage.Equal(state.Htmlpage) {
		tflog.Debug(ctx, "htmlpage has changed for responderaction")
		hasChange = true
	}
	if !data.Reasonphrase.Equal(state.Reasonphrase) {
		tflog.Debug(ctx, "reasonphrase has changed for responderaction")
		if config.Reasonphrase.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "reasonphrase")
		} else {
			hasChange = true
		}
	}
	if !data.Responsestatuscode.Equal(state.Responsestatuscode) {
		tflog.Debug(ctx, "responsestatuscode has changed for responderaction")
		if config.Responsestatuscode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "responsestatuscode")
		} else {
			hasChange = true
		}
	}
	if !data.Target.Equal(state.Target) {
		tflog.Debug(ctx, "target has changed for responderaction")
		hasChange = true
	}

	if hasChange {
		responderaction := responderactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update key must be the CURRENT LIVE name (data.Id), which after a rename
		// is newName.
		responderaction.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Responderaction.Type(), data.Id.ValueString(), &responderaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated responderaction resource")
	} else {
		tflog.Debug(ctx, "No in-place changes detected for responderaction resource")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. The unset key is the CURRENT LIVE name (data.Id),
	// which after a rename is newName.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Responderaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset responderaction attributes, got error: %s", err))
		return
	}

	// Read the current state back. Preserve the user-facing key/rename inputs across
	// the read so a rename (where the live name diverges from the configured name)
	// does not clobber them / produce a perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readResponderactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Responderaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete responderaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted responderaction resource")
}

// Helper function to read responderaction data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *ResponderactionResource) readResponderactionFromApi(ctx context.Context, data *ResponderactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value
	responderactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Responderaction.Type(), responderactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderaction, got error: %s", err))
		return false
	}

	responderactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
