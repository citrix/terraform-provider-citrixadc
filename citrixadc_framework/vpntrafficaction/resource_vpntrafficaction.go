package vpntrafficaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpntrafficactionResource{}
var _ resource.ResourceWithConfigure = (*VpntrafficactionResource)(nil)
var _ resource.ResourceWithImportState = (*VpntrafficactionResource)(nil)

func NewVpntrafficactionResource() resource.Resource {
	return &VpntrafficactionResource{}
}

// VpntrafficactionResource defines the resource implementation.
type VpntrafficactionResource struct {
	client *service.NitroClient
}

func (r *VpntrafficactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpntrafficactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpntrafficaction"
}

func (r *VpntrafficactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpntrafficactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpntrafficactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpntrafficaction resource")

	vpntrafficaction := vpntrafficactionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpntrafficactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpntrafficaction.Type(), vpntrafficactionName, &vpntrafficaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpntrafficaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpntrafficaction resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(vpntrafficactionName)

	// Read the updated state back
	if !r.readVpntrafficactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpntrafficaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpntrafficactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpntrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpntrafficaction resource")

	found := r.readVpntrafficactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpntrafficactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpntrafficactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpntrafficaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Apptimeout.Equal(state.Apptimeout) {
		tflog.Debug(ctx, "apptimeout has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Formssoaction.Equal(state.Formssoaction) {
		tflog.Debug(ctx, "formssoaction has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Fta.Equal(state.Fta) {
		tflog.Debug(ctx, "fta has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Hdx.Equal(state.Hdx) {
		tflog.Debug(ctx, "hdx has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		tflog.Debug(ctx, "kcdaccount has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Passwdexpression.Equal(state.Passwdexpression) {
		tflog.Debug(ctx, "passwdexpression has changed for vpntrafficaction")
		if config.Passwdexpression.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "passwdexpression")
		} else {
			hasChange = true
		}
	}
	if !data.Proxy.Equal(state.Proxy) {
		tflog.Debug(ctx, "proxy has changed for vpntrafficaction")
		if config.Proxy.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxy")
		} else {
			hasChange = true
		}
	}
	// qual is is_updateable:false (RequiresReplace) - a change to it forces
	// recreation and never reaches Update, so it is not checked here.
	if !data.Samlssoprofile.Equal(state.Samlssoprofile) {
		tflog.Debug(ctx, "samlssoprofile has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Sso.Equal(state.Sso) {
		tflog.Debug(ctx, "sso has changed for vpntrafficaction")
		hasChange = true
	}
	if !data.Userexpression.Equal(state.Userexpression) {
		tflog.Debug(ctx, "userexpression has changed for vpntrafficaction")
		if config.Userexpression.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "userexpression")
		} else {
			hasChange = true
		}
	}
	if !data.Wanscaler.Equal(state.Wanscaler) {
		tflog.Debug(ctx, "wanscaler has changed for vpntrafficaction")
		hasChange = true
	}

	if hasChange {
		// Use the updatable-only payload: qual is is_updateable:false and NITRO
		// rejects it on update (errorcode 278), so it must be excluded here.
		vpntrafficaction := vpntrafficactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		vpntrafficactionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpntrafficaction.Type(), vpntrafficactionName, &vpntrafficaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpntrafficaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpntrafficaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpntrafficaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpntrafficaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpntrafficaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpntrafficactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpntrafficaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpntrafficactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpntrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpntrafficaction resource")
	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Vpntrafficaction.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpntrafficaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpntrafficaction resource")
}

// Helper function to read vpntrafficaction data from API.
// Returns false when the resource no longer exists on the appliance.
func (r *VpntrafficactionResource) readVpntrafficactionFromApi(ctx context.Context, data *VpntrafficactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vpntrafficactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpntrafficaction.Type(), vpntrafficactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpntrafficaction, got error: %s", err))
		return false
	}

	vpntrafficactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
