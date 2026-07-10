package systemuser

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemuserResource{}
var _ resource.ResourceWithConfigure = (*SystemuserResource)(nil)
var _ resource.ResourceWithImportState = (*SystemuserResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SystemuserResource)(nil)

func NewSystemuserResource() resource.Resource {
	return &SystemuserResource{}
}

// SystemuserResource defines the resource implementation.
type SystemuserResource struct {
	client *service.NitroClient
}

func (r *SystemuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemuser"
}

func (r *SystemuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemuserResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SystemuserResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// A secret was supplied via either path — configuration is valid.
	if !data.Password.IsNull() || !data.PasswordWo.IsNull() {
		return
	}

	// No local password was provided. That is legitimate for certain users, so do
	// not require one when:
	//  - the username is not yet known (computed/interpolated) — cannot judge here
	//  - the username is nsroot — its password must be managed via
	//    "citrixadc_change_password", and this resource rejects a password for it
	//  - the user authenticates externally (externalauth = ENABLED) — no local password
	//
	// Note: ValidateConfig also runs during offline "terraform validate", where the
	// provider client is not configured, so we rely only on config values here. The
	// non-nsroot provider login user is handled in Create/Update where the client is
	// available.
	if data.Username.IsUnknown() {
		return
	}
	if data.Username.ValueString() == "nsroot" {
		return
	}
	if data.Externalauth.IsUnknown() || strings.EqualFold(data.Externalauth.ValueString(), "ENABLED") {
		return
	}

	resp.Diagnostics.AddAttributeError(
		path.Root("password"),
		"Missing Required Attribute",
		"Either \"password\" or \"password_wo\" must be specified for a local (non-external-auth) system user.",
	)
}

// hasLocalPassword reports whether a non-empty local password was supplied via
// either the persisted "password" attribute or the write-only "password_wo".
func (r *SystemuserResource) hasLocalPassword(data *SystemuserResourceModel, config *SystemuserResourceModel) bool {
	if !data.Password.IsNull() && data.Password.ValueString() != "" {
		return true
	}
	if !config.PasswordWo.IsNull() && config.PasswordWo.ValueString() != "" {
		return true
	}
	return false
}

// isPasswordExempt reports whether the given user legitimately needs no local
// password: the provider's login user, nsroot, or an external-auth user. Unlike
// ValidateConfig, this can consult the configured client for the login user.
func (r *SystemuserResource) isPasswordExempt(username string, externalauth types.String) bool {
	if username == "nsroot" || username == r.client.GetUsername() {
		return true
	}
	if strings.EqualFold(externalauth.ValueString(), "ENABLED") {
		return true
	}
	return false
}

func (r *SystemuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SystemuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemuser resource")

	username_value := data.Username.ValueString()
	loginUsername := r.client.GetUsername()

	// nsroot guard: block password changes for admin/login user via this resource
	if username_value == loginUsername || username_value == "nsroot" {
		hasPassword := (!data.Password.IsNull() && data.Password.ValueString() != "") ||
			(!config.PasswordWo.IsNull() && config.PasswordWo.ValueString() != "")
		if hasPassword {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"It seems you are trying to change the password of the Admin user. If so, please use the resource \"citrixadc_change_password\".",
			)
			return
		}
	}

	// Require a local password for regular local users. nsroot, the provider's
	// login user, and external-auth users legitimately have no local password.
	// This complements ValidateConfig, which cannot see the (dynamic) login user.
	if !r.hasLocalPassword(&data, &config) && !r.isPasswordExempt(username_value, data.Externalauth) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified for a local system user.",
		)
		return
	}

	// Get payload from plan (regular attributes)
	systemuser := systemuserGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	systemuserGetThePayloadFromtheConfig(ctx, &config, &systemuser)

	// Make API call
	if username_value == "nsroot" {
		// nsroot already exists — use UpdateResource instead of AddResource
		_, err := r.client.UpdateResource(service.Systemuser.Type(), username_value, &systemuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemuser %s, got error: %s", username_value, err))
			return
		}
	} else {
		_, err := r.client.AddResource(service.Systemuser.Type(), username_value, &systemuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemuser, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Created systemuser resource")

	// Handle inline cmdpolicybinding if configured
	if !data.Cmdpolicybinding.IsNull() {
		if err := r.updateCmdpolicyBindings(ctx, username_value, nil, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmdpolicybinding for systemuser, got error: %s", err))
			return
		}
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	// Read the updated state back
	if !r.readSystemuserFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemuser not found immediately after create/update")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemuser resource")

	found := r.readSystemuserFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SystemuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SystemuserResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	username_value := data.Username.ValueString()

	// nsroot guard: block password changes for admin/login user via this resource,
	// mirroring the same guard in Create so the admin password cannot be changed
	// through this resource on a subsequent apply.
	if username_value == r.client.GetUsername() || username_value == "nsroot" {
		hasPassword := (!data.Password.IsNull() && data.Password.ValueString() != "") ||
			(!config.PasswordWo.IsNull() && config.PasswordWo.ValueString() != "")
		if hasPassword {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"It seems you are trying to change the password of the Admin user. If so, please use the resource \"citrixadc_change_password\".",
			)
			return
		}
	}

	// Require a local password for regular local users. nsroot, the provider's
	// login user, and external-auth users legitimately have no local password.
	if !r.hasLocalPassword(&data, &config) && !r.isPasswordExempt(username_value, data.Externalauth) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified for a local system user.",
		)
		return
	}

	tflog.Debug(ctx, "Updating systemuser resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Allowedmanagementinterface.Equal(state.Allowedmanagementinterface) {
		tflog.Debug(ctx, fmt.Sprintf("allowedmanagementinterface has changed for systemuser"))
		hasChange = true
	}
	if !data.Externalauth.Equal(state.Externalauth) {
		tflog.Debug(ctx, fmt.Sprintf("externalauth has changed for systemuser"))
		hasChange = true
	}
	if !data.Logging.Equal(state.Logging) {
		tflog.Debug(ctx, fmt.Sprintf("logging has changed for systemuser"))
		hasChange = true
	}
	if !data.Maxsession.Equal(state.Maxsession) {
		tflog.Debug(ctx, fmt.Sprintf("maxsession has changed for systemuser"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for systemuser"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for systemuser"))
		hasChange = true
	}
	if !data.Promptstring.Equal(state.Promptstring) {
		tflog.Debug(ctx, fmt.Sprintf("promptstring has changed for systemuser"))
		hasChange = true
	}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, fmt.Sprintf("timeout has changed for systemuser"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		systemuser := systemuserGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		systemuserGetThePayloadFromtheConfig(ctx, &config, &systemuser)
		// Make API call
		// Named resource - use UpdateResource
		username_value := data.Username.ValueString()
		_, err := r.client.UpdateResource(service.Systemuser.Type(), username_value, &systemuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemuser, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemuser resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemuser resource, skipping update")
	}

	// Handle inline cmdpolicybinding changes
	if !data.Cmdpolicybinding.Equal(state.Cmdpolicybinding) {
		username_value := data.Username.ValueString()
		if err := r.updateCmdpolicyBindings(ctx, username_value, &state, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmdpolicybinding for systemuser, got error: %s", err))
			return
		}
	}

	// Read the updated state back
	if !r.readSystemuserFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemuser not found immediately after create/update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemuser resource")
	username_value := data.Username.ValueString()

	// nsroot guard: never delete the nsroot user, just remove from state
	if username_value == "nsroot" {
		tflog.Debug(ctx, "Skipping delete for nsroot user — removing from Terraform state only")
		return
	}

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Systemuser.Type(), username_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemuser resource")
}

// Helper function to read systemuser data from API
func (r *SystemuserResource) readSystemuserFromApi(ctx context.Context, data *SystemuserResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	username_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Systemuser.Type(), username_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemuser, got error: %s", err))
		return false
	}

	systemuserSetAttrFromGet(ctx, data, getResponseData)

	// Read cmdpolicybinding if it was configured
	if !data.Cmdpolicybinding.IsNull() {
		r.readCmdpolicyBindings(ctx, data, diags)
	}
	return true
}

// readCmdpolicyBindings reads the current cmdpolicy bindings from the API and sets them in state.
func (r *SystemuserResource) readCmdpolicyBindings(ctx context.Context, data *SystemuserResourceModel, diags *diag.Diagnostics) {
	username := data.Username.ValueString()
	bindings, _ := r.client.FindResourceArray("systemuser_systemcmdpolicy_binding", username)

	if len(bindings) == 0 {
		data.Cmdpolicybinding = types.SetNull(types.ObjectType{
			AttrTypes: cmdpolicyBindingAttrTypes(),
		})
		return
	}

	bindingModels := make([]CmdpolicyBindingModel, 0, len(bindings))
	for _, b := range bindings {
		model := CmdpolicyBindingModel{}
		if v, ok := b["policyname"].(string); ok {
			model.Policyname = types.StringValue(v)
		}
		if v, ok := b["priority"]; ok {
			if intVal, err := utils.ConvertToInt64(v); err == nil {
				model.Priority = types.Int64Value(intVal)
			}
		}
		bindingModels = append(bindingModels, model)
	}

	setValue, setDiags := types.SetValueFrom(ctx, types.ObjectType{
		AttrTypes: cmdpolicyBindingAttrTypes(),
	}, bindingModels)
	diags.Append(setDiags...)
	data.Cmdpolicybinding = setValue
}

// updateCmdpolicyBindings computes the diff between old and new bindings and applies changes.
func (r *SystemuserResource) updateCmdpolicyBindings(ctx context.Context, username string, oldState *SystemuserResourceModel, newState *SystemuserResourceModel) error {
	// Get old bindings
	var oldBindings []CmdpolicyBindingModel
	if oldState != nil && !oldState.Cmdpolicybinding.IsNull() {
		oldState.Cmdpolicybinding.ElementsAs(ctx, &oldBindings, false)
	}

	// Get new bindings
	var newBindings []CmdpolicyBindingModel
	if !newState.Cmdpolicybinding.IsNull() {
		newState.Cmdpolicybinding.ElementsAs(ctx, &newBindings, false)
	}

	// Build maps for diff
	oldMap := make(map[string]CmdpolicyBindingModel)
	for _, b := range oldBindings {
		oldMap[b.Policyname.ValueString()] = b
	}
	newMap := make(map[string]CmdpolicyBindingModel)
	for _, b := range newBindings {
		newMap[b.Policyname.ValueString()] = b
	}

	// Delete removed bindings
	for key, b := range oldMap {
		if _, exists := newMap[key]; !exists {
			if err := r.deleteSingleCmdpolicyBinding(username, b); err != nil {
				return err
			}
		}
	}

	// Add new or changed bindings
	for key, b := range newMap {
		if old, exists := oldMap[key]; !exists || !old.Priority.Equal(b.Priority) {
			// Delete old binding first if it exists with different priority
			if exists {
				if err := r.deleteSingleCmdpolicyBinding(username, old); err != nil {
					return err
				}
			}
			if err := r.addSingleCmdpolicyBinding(username, b); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *SystemuserResource) deleteSingleCmdpolicyBinding(username string, binding CmdpolicyBindingModel) error {
	args := []string{
		fmt.Sprintf("policyname:%s", binding.Policyname.ValueString()),
	}
	return r.client.DeleteResourceWithArgs("systemuser_systemcmdpolicy_binding", username, args)
}

func (r *SystemuserResource) addSingleCmdpolicyBinding(username string, binding CmdpolicyBindingModel) error {
	bindingStruct := system.Systemusercmdpolicybinding{
		Username:   username,
		Policyname: binding.Policyname.ValueString(),
		Priority:   uint32(binding.Priority.ValueInt64()),
	}
	_, err := r.client.UpdateResource("systemuser_systemcmdpolicy_binding", username, bindingStruct)
	return err
}

// cmdpolicyBindingAttrTypes returns the attribute types for the cmdpolicybinding set.
func cmdpolicyBindingAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"policyname": types.StringType,
		"priority":   types.Int64Type,
	}
}
