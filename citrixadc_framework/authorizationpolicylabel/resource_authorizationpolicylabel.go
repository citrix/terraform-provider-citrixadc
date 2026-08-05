package authorizationpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthorizationpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*AuthorizationpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*AuthorizationpolicylabelResource)(nil)

func NewAuthorizationpolicylabelResource() resource.Resource {
	return &AuthorizationpolicylabelResource{}
}

// AuthorizationpolicylabelResource defines the resource implementation.
type AuthorizationpolicylabelResource struct {
	client *service.NitroClient
}

func (r *AuthorizationpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthorizationpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authorizationpolicylabel"
}

func (r *AuthorizationpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthorizationpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthorizationpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authorizationpolicylabel resource")
	authorizationpolicylabel := authorizationpolicylabelGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Authorizationpolicylabel.Type(), labelname_value, &authorizationpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authorizationpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authorizationpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readAuthorizationpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthorizationpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authorizationpolicylabel resource")

	r.readAuthorizationpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *AuthorizationpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthorizationpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authorizationpolicylabel resource")

	// Rename support: authorizationpolicylabel exposes NO set/update endpoint. The
	// only in-place mutation NITRO offers is the `rename` action. The primary key
	// (labelname) uses RequiresReplace, so Terraform recreates the resource on a
	// labelname change and never reaches here for it. The ONLY change that lands in
	// Update is `newname`.
	//
	// On a newname change, POST {labelname, newname} to ?action=rename, then point
	// the resource ID at the new name so subsequent reads address the live object.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID -
		// NOT state.Labelname. state.Labelname stays pinned to the originally
		// configured value, so on a SECOND rename it would point at the wrong (no
		// longer live) name. The live name is whatever the prior rename set the ID to
		// (== labelname before any rename, == the prior newname after one).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming authorizationpolicylabel from %q to %q", oldName, newName))

		renamePayload := authorization.Authorizationpolicylabel{
			Labelname: oldName,
			Newname:   newName,
		}
		if err := r.client.ActOnResource(service.Authorizationpolicylabel.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename authorizationpolicylabel, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read
		// below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	tflog.Trace(ctx, "Updated authorizationpolicylabel resource")

	// Read the current state back. Capture the plan's user-facing labelname/newname
	// and restore them after the read so a rename does not clobber the configured
	// values (avoids an inconsistent-result / perpetual diff).
	planLabelname := data.Labelname
	planNewname := data.Newname
	r.readAuthorizationpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
	data.Labelname = planLabelname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthorizationpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authorizationpolicylabel resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== labelname at create, == newname after a rename), so we must delete
	// by data.Id, NOT data.Labelname (which stays at the originally configured value
	// and would target a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authorizationpolicylabel.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authorizationpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authorizationpolicylabel resource")
}

// Helper function to read authorizationpolicylabel data from API
func (r *AuthorizationpolicylabelResource) readAuthorizationpolicylabelFromApi(ctx context.Context, data *AuthorizationpolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authorizationpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authorizationpolicylabel, got error: %s", err))
		return
	}

	authorizationpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
