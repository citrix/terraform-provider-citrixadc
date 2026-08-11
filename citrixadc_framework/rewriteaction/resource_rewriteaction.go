package rewriteaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkv2resource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RewriteactionResource{}
var _ resource.ResourceWithConfigure = (*RewriteactionResource)(nil)
var _ resource.ResourceWithImportState = (*RewriteactionResource)(nil)

func NewRewriteactionResource() resource.Resource {
	return &RewriteactionResource{}
}

// RewriteactionResource defines the resource implementation.
type RewriteactionResource struct {
	client *service.NitroClient
}

func (r *RewriteactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewriteactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewriteaction"
}

func (r *RewriteactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewriteactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewriteactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewriteaction resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the user
	// does not supply a name, generate a unique one.
	rewriteactionName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || rewriteactionName == "" {
		rewriteactionName = sdkv2resource.PrefixedUniqueId("tf-rewriteaction-")
		data.Name = types.StringValue(rewriteactionName)
	}

	// Create API request body from the model
	rewriteaction := rewriteactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Rewriteaction.Type(), rewriteactionName, &rewriteaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewriteaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rewriteaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(rewriteactionName)

	// Read the updated state back
	if !r.readRewriteactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rewriteaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewriteactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewriteaction resource")

	found := r.readRewriteactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RewriteactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RewriteactionResourceModel

	// Read Terraform prior state to preserve ID / live name
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rewriteaction resource")

	// Rename support: NITRO exposes a `rename` action for rewriteaction. A newname
	// change drives an in-place rename (POST ?action=rename), never a destroy/recreate.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by data.Id (== name at
		// create, == the prior newname after one rename), NOT data.Name (which stays
		// pinned to the originally configured value).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming rewriteaction from %q to %q", oldName, newName))

		renamePayload := rewrite.Rewriteaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Rewriteaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename rewriteaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update /
		// read-back below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular (in-place) update of the NITRO-updatable attributes.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for rewriteaction")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Refinesearch.Equal(state.Refinesearch) {
		tflog.Debug(ctx, "refinesearch has changed for rewriteaction")
		if config.Refinesearch.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "refinesearch")
		} else {
			hasChange = true
		}
	}
	if !data.Search.Equal(state.Search) {
		tflog.Debug(ctx, "search has changed for rewriteaction")
		hasChange = true
	}
	if !data.Stringbuilderexpr.Equal(state.Stringbuilderexpr) {
		tflog.Debug(ctx, "stringbuilderexpr has changed for rewriteaction")
		hasChange = true
	}
	if !data.Target.Equal(state.Target) {
		tflog.Debug(ctx, "target has changed for rewriteaction")
		hasChange = true
	}

	if hasChange {
		rewriteaction := rewriteactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update key must be the CURRENT LIVE name (data.Id), which after a rename
		// is newName.
		rewriteaction.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Rewriteaction.Type(), data.Id.ValueString(), &rewriteaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewriteaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated rewriteaction resource")
	} else {
		tflog.Debug(ctx, "No in-place changes detected for rewriteaction resource")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// NITRO defaults. Keyed by the CURRENT LIVE name (data.Id), which after a rename
	// is newName.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Rewriteaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset rewriteaction attributes, got error: %s", err))
		return
	}

	// Read the current state back. Preserve the user-facing key/rename inputs across
	// the read so a rename (where the live name diverges from the configured name)
	// does not clobber them / produce a perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readRewriteactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewriteactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewriteaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Rewriteaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rewriteaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rewriteaction resource")
}

// Helper function to read rewriteaction data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *RewriteactionResource) readRewriteactionFromApi(ctx context.Context, data *RewriteactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value
	rewriteactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Rewriteaction.Type(), rewriteactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewriteaction, got error: %s", err))
		return false
	}

	rewriteactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
