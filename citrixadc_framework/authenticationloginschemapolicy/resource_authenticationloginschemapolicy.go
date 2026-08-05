package authenticationloginschemapolicy

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
var _ resource.Resource = &AuthenticationloginschemapolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationloginschemapolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationloginschemapolicyResource)(nil)

func NewAuthenticationloginschemapolicyResource() resource.Resource {
	return &AuthenticationloginschemapolicyResource{}
}

// AuthenticationloginschemapolicyResource defines the resource implementation.
type AuthenticationloginschemapolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationloginschemapolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationloginschemapolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationloginschemapolicy"
}

func (r *AuthenticationloginschemapolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationloginschemapolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationloginschemapolicy resource")

	authenticationloginschemapolicy := authenticationloginschemapolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add is POST)
	authenticationloginschemapolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName, &authenticationloginschemapolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationloginschemapolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationloginschemapolicy resource")

	// Set ID for the resource before reading state (ID == name, matching SDK v2)
	data.Id = types.StringValue(fmt.Sprintf("%v", authenticationloginschemapolicyName))

	// Read the updated state back
	found := r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "authenticationloginschemapolicy not found immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemapolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationloginschemapolicy resource")

	found := r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationloginschemapolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationloginschemapolicyResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID (the CURRENT LIVE name) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationloginschemapolicy resource")

	// Rename support: NITRO exposes a `rename` action (?action=rename). A change to
	// newname drives an in-place rename rather than a recreate. name itself is
	// RequiresReplace, so a name change never reaches Update.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and would
		// be wrong on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authenticationloginschemapolicy from %q to %q", oldName, newName))

		renamePayload := authentication.Authenticationloginschemapolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Authenticationloginschemapolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authenticationloginschemapolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update and
		// read below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular update: NITRO update is PUT for the mutable attributes.
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authenticationloginschemapolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for authenticationloginschemapolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for authenticationloginschemapolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationloginschemapolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for authenticationloginschemapolicy")
		hasChange = true
	}

	if hasChange {
		authenticationloginschemapolicy := authenticationloginschemapolicyGetThePayloadFromthePlan(ctx, &data)
		// Target the CURRENT LIVE name (== newname after a rename, else == name).
		liveName := data.Id.ValueString()
		authenticationloginschemapolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Authenticationloginschemapolicy.Type(), liveName, &authenticationloginschemapolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationloginschemapolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated authenticationloginschemapolicy resource")
	} else {
		tflog.Debug(ctx, "No mutable changes detected for authenticationloginschemapolicy resource")
	}

	// Read the current state back. Capture the plan's user-facing key attributes and
	// restore them after the read so a rename does not clobber the configured values.
	planName := data.Name
	planNewname := data.Newname
	r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemapolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationloginschemapolicy resource")

	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationloginschemapolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationloginschemapolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationloginschemapolicy resource")
}

// Helper function to read authenticationloginschemapolicy data from API.
// Returns true if the resource was found, false if it no longer exists.
func (r *AuthenticationloginschemapolicyResource) readAuthenticationloginschemapolicyFromApi(ctx context.Context, data *AuthenticationloginschemapolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (the live name)
	authenticationloginschemapolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationloginschemapolicy, got error: %s", err))
		return false
	}

	authenticationloginschemapolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
