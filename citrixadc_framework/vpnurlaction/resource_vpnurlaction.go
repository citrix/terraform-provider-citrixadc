package vpnurlaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnurlactionResource{}
var _ resource.ResourceWithConfigure = (*VpnurlactionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnurlactionResource)(nil)

func NewVpnurlactionResource() resource.Resource {
	return &VpnurlactionResource{}
}

// VpnurlactionResource defines the resource implementation.
type VpnurlactionResource struct {
	client *service.NitroClient
}

func (r *VpnurlactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnurlactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnurlaction"
}

func (r *VpnurlactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnurlactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnurlactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnurlaction resource")

	vpnurlaction := vpnurlactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnurlaction.Type(), name_value, &vpnurlaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnurlaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnurlaction resource")

	// Set ID for the resource before reading state (single_unique - plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readVpnurlactionFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnurlactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnurlaction resource")

	r.readVpnurlactionFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - drop it from state (mirrors SDK v2
	// clearing the ID when FindResource fails).
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnurlactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID (current live name) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnurlaction resource")

	// Rename support: NITRO exposes a `rename` action plus a `newname` attribute.
	// A newname change drives an in-place rename (NOT a destroy/recreate). The
	// rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT state.Name
	// (which stays pinned to the originally configured value and would be stale on
	// a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming vpnurlaction from %q to %q", oldName, newName))

		renamePayload := vpn.Vpnurlaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Vpnurlaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename vpnurlaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so subsequent
		// operations (update payload, read-back, future reads) address it.
		data.Id = types.StringValue(newName)
	}

	// Regular update of the mutable attributes. name is ForceNew (RequiresReplace)
	// so it never changes here; every other non-rename attribute is updateable.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Actualurl.Equal(state.Actualurl) {
		hasChange = true
	}
	if !data.Applicationtype.Equal(state.Applicationtype) {
		if config.Applicationtype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "applicationtype")
		} else {
			hasChange = true
		}
	}
	if !data.Clientlessaccess.Equal(state.Clientlessaccess) {
		if config.Clientlessaccess.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientlessaccess")
		} else {
			hasChange = true
		}
	}
	if !data.Comment.Equal(state.Comment) {
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Iconurl.Equal(state.Iconurl) {
		hasChange = true
	}
	if !data.Linkname.Equal(state.Linkname) {
		hasChange = true
	}
	if !data.Samlssoprofile.Equal(state.Samlssoprofile) {
		hasChange = true
	}
	if !data.Ssotype.Equal(state.Ssotype) {
		if config.Ssotype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssotype")
		} else {
			hasChange = true
		}
	}
	if !data.Vservername.Equal(state.Vservername) {
		if config.Vservername.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "vservername")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		vpnurlaction := vpnurlactionGetThePayloadFromthePlan(ctx, &data)
		// Address the current live name (== configured name, or the post-rename
		// name if a rename happened above).
		liveName := data.Id.ValueString()
		vpnurlaction.Name = liveName
		_, err := r.client.UpdateResource(service.Vpnurlaction.Type(), liveName, &vpnurlaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnurlaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpnurlaction resource")
	} else {
		tflog.Debug(ctx, "No mutable changes detected for vpnurlaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	// Address the current live name (post-rename if a rename happened above).
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnurlaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnurlaction attributes, got error: %s", err))
		return
	}

	// Read the current state back. Preserve the user-facing key (name) and the
	// rename-only newname across the read-back so a rename does not surface as a
	// spurious diff.
	planName := data.Name
	planNewname := data.Newname
	r.readVpnurlactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnurlactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnurlaction resource")

	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id,
	// NOT data.Name (which stays at the originally configured value and would target
	// a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnurlaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnurlaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnurlaction resource")
}

// Helper function to read vpnurlaction data from API
func (r *VpnurlactionResource) readVpnurlactionFromApi(ctx context.Context, data *VpnurlactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnurlaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnurlaction, got error: %s", err))
		return
	}

	vpnurlactionSetAttrFromGet(ctx, data, getResponseData)
}
