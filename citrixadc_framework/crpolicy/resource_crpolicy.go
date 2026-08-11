package crpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cr"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CrpolicyResource{}
var _ resource.ResourceWithConfigure = (*CrpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*CrpolicyResource)(nil)

func NewCrpolicyResource() resource.Resource {
	return &CrpolicyResource{}
}

// CrpolicyResource defines the resource implementation.
type CrpolicyResource struct {
	client *service.NitroClient
}

func (r *CrpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crpolicy"
}

func (r *CrpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crpolicy resource")

	crpolicy := crpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	policyname_value := data.Policyname.ValueString()
	_, err := r.client.AddResource(service.Crpolicy.Type(), policyname_value, &crpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created crpolicy resource")

	// Set ID for the resource before reading state (single unique attr -> plain value)
	data.Id = types.StringValue(data.Policyname.ValueString())

	// Read the updated state back
	if !r.readCrpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "crpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crpolicy resource")

	found := r.readCrpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CrpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CrpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to distinguish an attribute removed from config (-> unset) from
	// one merely changed (-> update).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating crpolicy resource")

	// Rename support (NITRO ?action=rename). policyname is RequiresReplace so a key
	// change recreates the resource and never reaches Update; the only change that
	// drives an in-place rename here is `newname`. Mirrors the SDK v2 convention
	// (see citrixadc/resource_citrixadc_appfwpolicy.go).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Policyname (which stays pinned to the originally configured value and
		// would point at the wrong name on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming crpolicy from %q to %q", oldName, newName))

		renamePayload := cr.Crpolicy{
			Policyname: oldName,
			Newname:    newName,
		}
		if err := r.client.ActOnResource(service.Crpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename crpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update and
		// read below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular update for the mutable, non-key attributes (PUT /crpolicy).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for crpolicy")
		hasChange = true
	}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for crpolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for crpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		crpolicy := crpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Key on the CURRENT LIVE name (data.Id) so an update after a rename targets
		// the renamed object rather than the old configured policyname.
		liveName := data.Id.ValueString()
		crpolicy.Policyname = liveName
		_, err := r.client.UpdateResource(service.Crpolicy.Type(), liveName, &crpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated crpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for crpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. data.Id holds the current live name (== policyname, or
	// newname after a rename), so the unset always targets the live object.
	unsetIdPayload := map[string]interface{}{
		"policyname": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Crpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset crpolicy attributes, got error: %s", err))
		return
	}

	// Read the current state back. The resource may now be physically named newName,
	// so preserve the user-facing key + newname across the read-back to avoid an
	// inconsistent-result / perpetual diff.
	planPolicyname := data.Policyname
	planNewname := data.Newname
	if !r.readCrpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "crpolicy not found immediately after update")
		}
		return
	}
	data.Policyname = planPolicyname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== policyname at create, == newname after a rename), so delete by
	// data.Id, NOT data.Policyname (which stays at the originally configured value
	// and would target a non-existent name after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Crpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete crpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted crpolicy resource")
}

// Helper function to read crpolicy data from API. Returns false if the resource
// no longer exists on the ADC.
func (r *CrpolicyResource) readCrpolicyFromApi(ctx context.Context, data *CrpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain (live) name
	policyname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Crpolicy.Type(), policyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crpolicy, got error: %s", err))
		return false
	}

	crpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
