package vpnurlpolicy

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
var _ resource.Resource = &VpnurlpolicyResource{}
var _ resource.ResourceWithConfigure = (*VpnurlpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VpnurlpolicyResource)(nil)

func NewVpnurlpolicyResource() resource.Resource {
	return &VpnurlpolicyResource{}
}

// VpnurlpolicyResource defines the resource implementation.
type VpnurlpolicyResource struct {
	client *service.NitroClient
}

func (r *VpnurlpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnurlpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnurlpolicy"
}

func (r *VpnurlpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnurlpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnurlpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnurlpolicy resource")

	vpnurlpolicy := vpnurlpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnurlpolicy.Type(), name_value, &vpnurlpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnurlpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnurlpolicy resource")

	// Set ID for the resource before reading state (single unique attribute - plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readVpnurlpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnurlpolicy not found immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnurlpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnurlpolicy resource")

	r.readVpnurlpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnurlpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnurlpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnurlpolicy resource")

	// Rename support: on a newname change, POST {name, newname} to ?action=rename,
	// then point the resource ID at the new name so subsequent reads/updates address
	// the live object. The rename SOURCE is the CURRENT LIVE name, tracked by the ID
	// (== name at create, == the prior newname after one rename) - NOT state.Name,
	// which stays pinned to the originally configured value.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming vpnurlpolicy from %q to %q", oldName, newName))

		renamePayload := vpn.Vpnurlpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Vpnurlpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename vpnurlpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it.
		data.Id = types.StringValue(newName)
	}

	// Regular in-place update of the mutable fields (action, comment, logaction, rule).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for vpnurlpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for vpnurlpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for vpnurlpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for vpnurlpolicy")
		hasChange = true
	}

	if hasChange {
		vpnurlpolicy := vpnurlpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Address the current live object (the ID reflects any rename above).
		vpnurlpolicy.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Vpnurlpolicy.Type(), data.Id.ValueString(), &vpnurlpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnurlpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpnurlpolicy resource")
	} else {
		tflog.Debug(ctx, "No mutable changes detected for vpnurlpolicy resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// defaults. Address the current live object (data.Id reflects any rename above).
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnurlpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnurlpolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back. SetAttrFromGet must not clobber the user-facing
	// name/newname (still the configured/plan values), so capture and restore them.
	planName := data.Name
	planNewname := data.Newname
	r.readVpnurlpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnurlpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnurlpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id,
	// NOT data.Name (which stays at the originally configured value and would target
	// a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnurlpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnurlpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnurlpolicy resource")
}

// Helper function to read vpnurlpolicy data from API
func (r *VpnurlpolicyResource) readVpnurlpolicyFromApi(ctx context.Context, data *VpnurlpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name.
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Vpnurlpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnurlpolicy, got error: %s", err))
		return
	}

	vpnurlpolicySetAttrFromGet(ctx, data, getResponseData)
}
