package authorizationpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthorizationpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthorizationpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthorizationpolicyResource)(nil)

func NewAuthorizationpolicyResource() resource.Resource {
	return &AuthorizationpolicyResource{}
}

// AuthorizationpolicyResource defines the resource implementation.
type AuthorizationpolicyResource struct {
	client *service.NitroClient
}

func (r *AuthorizationpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthorizationpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authorizationpolicy"
}

func (r *AuthorizationpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthorizationpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthorizationpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authorizationpolicy resource")

	authorizationpolicy := authorizationpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authorizationpolicy.Type(), name_value, &authorizationpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authorizationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authorizationpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthorizationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthorizationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authorizationpolicy resource")

	r.readAuthorizationpolicyFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *AuthorizationpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthorizationpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authorizationpolicy resource")

	// Rename support: authorizationpolicy exposes a NITRO `rename` action. `name` uses
	// RequiresReplace, so a name change recreates the resource and never reaches here.
	// The ONLY name-related change that lands in Update is `newname`, which drives an
	// in-place rename (mirrors SDK v2 appfwpolicy convention).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID - NOT
		// state.Name. state.Name stays pinned to the originally configured value, so on
		// a SECOND rename it would point at the wrong (no longer live) name.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authorizationpolicy from %q to %q", oldName, newName))

		renamePayload := authorization.Authorizationpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Authorizationpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authorizationpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update/read
		// below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Handle in-place attribute updates (rule/action) via NITRO update (PUT).
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authorizationpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authorizationpolicy")
		hasChange = true
	}

	if hasChange {
		authorizationpolicy := authorizationpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update must target the CURRENT LIVE name (post-rename), which is data.Id -
		// NOT data.Name (which stays pinned to the originally configured value).
		liveName := data.Id.ValueString()
		authorizationpolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Authorizationpolicy.Type(), liveName, &authorizationpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authorizationpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated authorizationpolicy resource")
	} else {
		tflog.Debug(ctx, "No attribute changes detected for authorizationpolicy resource, skipping update")
	}

	// Read the current state back. The resource may now be physically named newName,
	// so preserve the user-facing name/newname across the read-back to avoid an
	// inconsistent-result / perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readAuthorizationpolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthorizationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authorizationpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so we must delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authorizationpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authorizationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authorizationpolicy resource")
}

// Helper function to read authorizationpolicy data from API
func (r *AuthorizationpolicyResource) readAuthorizationpolicyFromApi(ctx context.Context, data *AuthorizationpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authorizationpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authorizationpolicy, got error: %s", err))
		return
	}

	authorizationpolicySetAttrFromGet(ctx, data, getResponseData)

}
