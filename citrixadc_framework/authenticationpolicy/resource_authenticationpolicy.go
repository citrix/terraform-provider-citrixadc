package authenticationpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthenticationpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationpolicyResource)(nil)

func NewAuthenticationpolicyResource() resource.Resource {
	return &AuthenticationpolicyResource{}
}

// AuthenticationpolicyResource defines the resource implementation.
type AuthenticationpolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationpolicy"
}

func (r *AuthenticationpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationpolicy resource")

	authenticationpolicy := authenticationpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationpolicy.Type(), authenticationpolicyName, &authenticationpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationpolicy resource")

	r.readAuthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove it from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationpolicyResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to distinguish an attribute removed from config (-> unset) from
	// one merely changed to a new value (-> update).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationpolicy resource")

	// Detect changes in the NITRO-updatable attributes (name is RequiresReplace and
	// never lands here; newname is handled separately via the rename action below).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authenticationpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for authenticationpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for authenticationpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for authenticationpolicy")
		hasChange = true
	}

	if hasChange {
		authenticationpolicy := authenticationpolicyGetThePayloadFromthePlan(ctx, &data)
		// Update targets the CURRENT LIVE name (tracked by the ID), which equals the
		// configured name unless a prior rename moved it.
		authenticationpolicy.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationpolicy.Type(), data.Id.ValueString(), &authenticationpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated authenticationpolicy resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for authenticationpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. The unset targets the CURRENT LIVE name tracked by the ID.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationpolicy attributes, got error: %s", err))
		return
	}

	// Rename support: authenticationpolicy exposes a NITRO `rename` action. A change
	// to newname drives an in-place rename (mirrors the SDK v2 appfwpolicy convention).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and would
		// point at the wrong name on a second rename).
		oldName := data.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authenticationpolicy from %q to %q", oldName, newName))

		renamePayload := authentication.Authenticationpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Authenticationpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authenticationpolicy, got error: %s", err))
			return
		}
		// The live object is now named newName. Point the ID at it so the read below
		// (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Read the current state back. Capture the plan-supplied name/newname and restore
	// them after the read: the object may now be physically named newName, but the
	// user-facing name attribute must keep its configured value to avoid a spurious
	// RequiresReplace diff / inconsistent-result error.
	planName := data.Name
	planNewname := data.Newname
	r.readAuthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationpolicy resource")
}

// Helper function to read authenticationpolicy data from API
func (r *AuthenticationpolicyResource) readAuthenticationpolicyFromApi(ctx context.Context, data *AuthenticationpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	authenticationpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationpolicy.Type(), authenticationpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationpolicy, got error: %s", err))
		return
	}

	authenticationpolicySetAttrFromGet(ctx, data, getResponseData)
}
