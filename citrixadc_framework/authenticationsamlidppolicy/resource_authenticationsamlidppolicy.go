package authenticationsamlidppolicy

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
var _ resource.Resource = &AuthenticationsamlidppolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationsamlidppolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationsamlidppolicyResource)(nil)

func NewAuthenticationsamlidppolicyResource() resource.Resource {
	return &AuthenticationsamlidppolicyResource{}
}

// AuthenticationsamlidppolicyResource defines the resource implementation.
type AuthenticationsamlidppolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationsamlidppolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationsamlidppolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationsamlidppolicy"
}

func (r *AuthenticationsamlidppolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationsamlidppolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationsamlidppolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationsamlidppolicy resource")
	authenticationsamlidppolicy := authenticationsamlidppolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationsamlidppolicy.Type(), name_value, &authenticationsamlidppolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationsamlidppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationsamlidppolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationsamlidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidppolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationsamlidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationsamlidppolicy resource")

	r.readAuthenticationsamlidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove it from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidppolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationsamlidppolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationsamlidppolicy resource")

	// Rename support: authenticationsamlidppolicy exposes a `rename` action
	// (?action=rename). The primary key `name` is RequiresReplace, so a name change
	// recreates the resource and never reaches here. The ONLY key mutation that lands
	// in Update is `newname`, which drives an in-place rename. Mirrors the SDK v2
	// convention (see citrixadc/resource_citrixadc_appfwpolicy.go).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID -
		// NOT state.Name. state.Name stays pinned to the originally configured value,
		// so on a SECOND rename it would point at the wrong (no longer live) name.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authenticationsamlidppolicy from %q to %q", oldName, newName))

		renamePayload := authentication.Authenticationsamlidppolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Authenticationsamlidppolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authenticationsamlidppolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update and
		// read below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Check if there are any changes in updateable attributes.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authenticationsamlidppolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for authenticationsamlidppolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for authenticationsamlidppolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationsamlidppolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for authenticationsamlidppolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to updatable fields.
		authenticationsamlidppolicy := authenticationsamlidppolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update payload must identify the object by its CURRENT LIVE name
		// (== configured name, or newname after a rename), which is tracked by the ID.
		liveName := data.Id.ValueString()
		authenticationsamlidppolicy.Name = liveName
		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Authenticationsamlidppolicy.Type(), liveName, &authenticationsamlidppolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationsamlidppolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationsamlidppolicy resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for authenticationsamlidppolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. The unset must identify the object by its CURRENT LIVE name
	// (tracked by the ID), matching the update path above.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationsamlidppolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationsamlidppolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back. Capture the plan key values and restore them after
	// the read so a rename (where GET returns the new live name) does not clobber the
	// user-facing name attribute and cause a spurious diff.
	planName := data.Name
	planNewname := data.Newname
	r.readAuthenticationsamlidppolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidppolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationsamlidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationsamlidppolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so we must delete by
	// data.Id, NOT data.Name (which stays at the originally configured value and would
	// target a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationsamlidppolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationsamlidppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationsamlidppolicy resource")
}

// Helper function to read authenticationsamlidppolicy data from API
func (r *AuthenticationsamlidppolicyResource) readAuthenticationsamlidppolicyFromApi(ctx context.Context, data *AuthenticationsamlidppolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (the policy name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationsamlidppolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationsamlidppolicy, got error: %s", err))
		return
	}

	authenticationsamlidppolicySetAttrFromGet(ctx, data, getResponseData)

}
