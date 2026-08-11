package authenticationvserver

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
var _ resource.Resource = &AuthenticationvserverResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverResource)(nil)

func NewAuthenticationvserverResource() resource.Resource {
	return &AuthenticationvserverResource{}
}

// AuthenticationvserverResource defines the resource implementation.
type AuthenticationvserverResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver"
}

func (r *AuthenticationvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver resource")

	authenticationvserverName := data.Name.ValueString()
	authenticationvserver := authenticationvserverGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Authenticationvserver.Type(), authenticationvserverName, &authenticationvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver resource")

	// The ID is the resource name (single unique attribute).
	data.Id = types.StringValue(authenticationvserverName)

	// Read the updated state back
	if !r.readAuthenticationvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver resource")

	found := r.readAuthenticationvserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationvserverResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the live object name).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver resource")

	// Handle in-place rename first (newname change). The rename source must be the
	// current live name, which is held in state.Id (not the possibly-stale name attr).
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		renamePayload := authentication.Authenticationvserver{
			Name:    state.Id.ValueString(),
			Newname: newName,
		}
		err := r.client.ActOnResource(service.Authenticationvserver.Type(), &renamePayload, "rename")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authenticationvserver, got error: %s", err))
			return
		}
		// The ID now tracks the renamed (live) object.
		data.Id = types.StringValue(newName)
	}

	// Current live name for the subsequent set / enable-disable calls.
	authenticationvserverName := data.Id.ValueString()

	// Handle state (admin) change via the enable/disable actions (mirrors SDK v2).
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doAuthenticationvserverStateChange(ctx, authenticationvserverName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable/disable authenticationvserver, got error: %s", err))
			return
		}
	}

	// Detect changes in the NITRO-updatable ("set") attributes only. servicetype/port/
	// range/td are create-only (RequiresReplace) and never reach Update; state is handled
	// above via enable/disable.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Appflowlog.Equal(state.Appflowlog) {
		if config.Appflowlog.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "appflowlog")
		} else {
			hasChange = true
		}
	}
	if !data.Authentication.Equal(state.Authentication) {
		if config.Authentication.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "authentication")
		} else {
			hasChange = true
		}
	}
	if !data.Authenticationdomain.Equal(state.Authenticationdomain) {
		hasChange = true
	}
	if !data.Certkeynames.Equal(state.Certkeynames) {
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		hasChange = true
	}
	if !data.Failedlogintimeout.Equal(state.Failedlogintimeout) {
		hasChange = true
	}
	if !data.Ipv46.Equal(state.Ipv46) {
		hasChange = true
	}
	if !data.Maxloginattempts.Equal(state.Maxloginattempts) {
		hasChange = true
	}
	if !data.Samesite.Equal(state.Samesite) {
		hasChange = true
	}

	if hasChange {
		authenticationvserver := authenticationvserverGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Ensure the payload targets the current live name.
		authenticationvserver.Name = authenticationvserverName
		_, err := r.client.UpdateResource(service.Authenticationvserver.Type(), authenticationvserverName, &authenticationvserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated authenticationvserver resource")
	} else {
		tflog.Debug(ctx, "No set-attribute changes detected for authenticationvserver resource")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": authenticationvserverName,
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationvserver.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationvserver attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver resource")

	// Delete keyed on the ID (the live object name; may differ from the name attr after a rename).
	err := r.client.DeleteResource(service.Authenticationvserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver resource")
}

// doAuthenticationvserverStateChange enables/disables the virtual server via the
// NITRO action endpoints (mirrors the SDK v2 doAuthenticationvserverStateChange).
func (r *AuthenticationvserverResource) doAuthenticationvserverStateChange(ctx context.Context, name string, newstate string) error {
	tflog.Debug(ctx, "In doAuthenticationvserverStateChange")

	// A fresh struct with only the name is used: ActOnResource fails on superfluous attributes.
	authenticationvserver := authentication.Authenticationvserver{
		Name: name,
	}

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Authenticationvserver.Type(), &authenticationvserver, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Authenticationvserver.Type(), &authenticationvserver, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read authenticationvserver data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *AuthenticationvserverResource) readAuthenticationvserverFromApi(ctx context.Context, data *AuthenticationvserverResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (the live name).
	authenticationvserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationvserver.Type(), authenticationvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver, got error: %s", err))
		return false
	}

	authenticationvserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
