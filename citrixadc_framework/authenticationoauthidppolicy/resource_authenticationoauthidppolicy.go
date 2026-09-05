package authenticationoauthidppolicy

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
var _ resource.Resource = &AuthenticationoauthidppolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationoauthidppolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationoauthidppolicyResource)(nil)

func NewAuthenticationoauthidppolicyResource() resource.Resource {
	return &AuthenticationoauthidppolicyResource{}
}

// AuthenticationoauthidppolicyResource defines the resource implementation.
type AuthenticationoauthidppolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationoauthidppolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationoauthidppolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationoauthidppolicy"
}

func (r *AuthenticationoauthidppolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationoauthidppolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationoauthidppolicy resource")
	authenticationoauthidppolicy := authenticationoauthidppolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationoauthidppolicy.Type(), name_value, &authenticationoauthidppolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationoauthidppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationoauthidppolicy resource")

	// Set ID for the resource before reading state (single unique attribute -> plain name)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidppolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationoauthidppolicy resource")

	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *AuthenticationoauthidppolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationoauthidppolicyResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (holds the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationoauthidppolicy resource")

	// Rename branch: authenticationoauthidppolicy exposes a NITRO `rename` action. On a
	// newname change, POST {name, newname} to ?action=rename and point the resource ID
	// at the new name so subsequent reads/updates address the live object. (name itself
	// is RequiresReplace, so a name change never reaches here - only newname does.)
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name, which stays pinned to the originally configured value and would
		// be stale on a second rename.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authenticationoauthidppolicy from %q to %q", oldName, newName))

		renamePayload := authentication.Authenticationoauthidppolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Authenticationoauthidppolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authenticationoauthidppolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it.
		data.Id = types.StringValue(newName)
	}

	// Regular update branch: detect changes in the NITRO-updatable attributes.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authenticationoauthidppolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for authenticationoauthidppolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for authenticationoauthidppolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationoauthidppolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for authenticationoauthidppolicy")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		authenticationoauthidppolicy := authenticationoauthidppolicyGetThePayloadFromthePlan(ctx, &data)
		// Address the live object (== newname after a rename) as the update key.
		authenticationoauthidppolicy.Name = data.Id.ValueString()
		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Authenticationoauthidppolicy.Type(), data.Id.ValueString(), &authenticationoauthidppolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationoauthidppolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated authenticationoauthidppolicy resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for authenticationoauthidppolicy resource")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Addressed by the live object name (== data.Id).
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationoauthidppolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationoauthidppolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back. Preserve the plan's user-facing key attributes across
	// the read so a rename does not clobber the configured name / newname.
	planName := data.Name
	planNewname := data.Newname
	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidppolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationoauthidppolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationoauthidppolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationoauthidppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationoauthidppolicy resource")
}

// Helper function to read authenticationoauthidppolicy data from API
func (r *AuthenticationoauthidppolicyResource) readAuthenticationoauthidppolicyFromApi(ctx context.Context, data *AuthenticationoauthidppolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value.
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationoauthidppolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationoauthidppolicy, got error: %s", err))
		return
	}

	authenticationoauthidppolicySetAttrFromGet(ctx, data, getResponseData)
}
