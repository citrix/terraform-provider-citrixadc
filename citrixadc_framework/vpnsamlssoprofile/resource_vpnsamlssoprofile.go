package vpnsamlssoprofile

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
var _ resource.Resource = &VpnsamlssoprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnsamlssoprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnsamlssoprofileResource)(nil)

func NewVpnsamlssoprofileResource() resource.Resource {
	return &VpnsamlssoprofileResource{}
}

// VpnsamlssoprofileResource defines the resource implementation.
type VpnsamlssoprofileResource struct {
	client *service.NitroClient
}

func (r *VpnsamlssoprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnsamlssoprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnsamlssoprofile"
}

func (r *VpnsamlssoprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnsamlssoprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnsamlssoprofile resource")

	// Create API request body from the plan
	vpnsamlssoprofile := vpnsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpnsamlssoprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName, &vpnsamlssoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnsamlssoprofile resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(vpnsamlssoprofileName)

	// Read the updated state back
	if !r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsamlssoprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsamlssoprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnsamlssoprofile resource")

	found := r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnsamlssoprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnsamlssoprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config so attributes removed from config can be detected (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnsamlssoprofile resource")

	// Determine which updateable attributes changed, and which unset-eligible
	// attributes were removed from config (so they can be reverted to defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Attribute1.Equal(state.Attribute1) {
		hasChange = true
	}
	if !data.Attribute1expr.Equal(state.Attribute1expr) {
		hasChange = true
	}
	if !data.Attribute1format.Equal(state.Attribute1format) {
		hasChange = true
	}
	if !data.Attribute1friendlyname.Equal(state.Attribute1friendlyname) {
		hasChange = true
	}
	if !data.Attribute2.Equal(state.Attribute2) {
		hasChange = true
	}
	if !data.Attribute2expr.Equal(state.Attribute2expr) {
		hasChange = true
	}
	if !data.Attribute2format.Equal(state.Attribute2format) {
		hasChange = true
	}
	if !data.Attribute2friendlyname.Equal(state.Attribute2friendlyname) {
		hasChange = true
	}
	if !data.Attribute3.Equal(state.Attribute3) {
		hasChange = true
	}
	if !data.Attribute3expr.Equal(state.Attribute3expr) {
		hasChange = true
	}
	if !data.Attribute3format.Equal(state.Attribute3format) {
		hasChange = true
	}
	if !data.Attribute3friendlyname.Equal(state.Attribute3friendlyname) {
		hasChange = true
	}
	if !data.Attribute4.Equal(state.Attribute4) {
		hasChange = true
	}
	if !data.Attribute4expr.Equal(state.Attribute4expr) {
		hasChange = true
	}
	if !data.Attribute4format.Equal(state.Attribute4format) {
		hasChange = true
	}
	if !data.Attribute4friendlyname.Equal(state.Attribute4friendlyname) {
		hasChange = true
	}
	if !data.Attribute5.Equal(state.Attribute5) {
		hasChange = true
	}
	if !data.Attribute5expr.Equal(state.Attribute5expr) {
		hasChange = true
	}
	if !data.Attribute5format.Equal(state.Attribute5format) {
		hasChange = true
	}
	if !data.Attribute5friendlyname.Equal(state.Attribute5friendlyname) {
		hasChange = true
	}
	if !data.Attribute6.Equal(state.Attribute6) {
		hasChange = true
	}
	if !data.Attribute6expr.Equal(state.Attribute6expr) {
		hasChange = true
	}
	if !data.Attribute6format.Equal(state.Attribute6format) {
		hasChange = true
	}
	if !data.Attribute6friendlyname.Equal(state.Attribute6friendlyname) {
		hasChange = true
	}
	if !data.Attribute7.Equal(state.Attribute7) {
		hasChange = true
	}
	if !data.Attribute7expr.Equal(state.Attribute7expr) {
		hasChange = true
	}
	if !data.Attribute7format.Equal(state.Attribute7format) {
		hasChange = true
	}
	if !data.Attribute7friendlyname.Equal(state.Attribute7friendlyname) {
		hasChange = true
	}
	if !data.Attribute8.Equal(state.Attribute8) {
		hasChange = true
	}
	if !data.Attribute8expr.Equal(state.Attribute8expr) {
		hasChange = true
	}
	if !data.Attribute8format.Equal(state.Attribute8format) {
		hasChange = true
	}
	if !data.Attribute8friendlyname.Equal(state.Attribute8friendlyname) {
		hasChange = true
	}
	if !data.Attribute9.Equal(state.Attribute9) {
		hasChange = true
	}
	if !data.Attribute9expr.Equal(state.Attribute9expr) {
		hasChange = true
	}
	if !data.Attribute9format.Equal(state.Attribute9format) {
		hasChange = true
	}
	if !data.Attribute9friendlyname.Equal(state.Attribute9friendlyname) {
		hasChange = true
	}
	if !data.Attribute10.Equal(state.Attribute10) {
		hasChange = true
	}
	if !data.Attribute10expr.Equal(state.Attribute10expr) {
		hasChange = true
	}
	if !data.Attribute10format.Equal(state.Attribute10format) {
		hasChange = true
	}
	if !data.Attribute10friendlyname.Equal(state.Attribute10friendlyname) {
		hasChange = true
	}
	if !data.Attribute11.Equal(state.Attribute11) {
		hasChange = true
	}
	if !data.Attribute11expr.Equal(state.Attribute11expr) {
		hasChange = true
	}
	if !data.Attribute11format.Equal(state.Attribute11format) {
		hasChange = true
	}
	if !data.Attribute11friendlyname.Equal(state.Attribute11friendlyname) {
		hasChange = true
	}
	if !data.Attribute12.Equal(state.Attribute12) {
		hasChange = true
	}
	if !data.Attribute12expr.Equal(state.Attribute12expr) {
		hasChange = true
	}
	if !data.Attribute12format.Equal(state.Attribute12format) {
		hasChange = true
	}
	if !data.Attribute12friendlyname.Equal(state.Attribute12friendlyname) {
		hasChange = true
	}
	if !data.Attribute13.Equal(state.Attribute13) {
		hasChange = true
	}
	if !data.Attribute13expr.Equal(state.Attribute13expr) {
		hasChange = true
	}
	if !data.Attribute13format.Equal(state.Attribute13format) {
		hasChange = true
	}
	if !data.Attribute13friendlyname.Equal(state.Attribute13friendlyname) {
		hasChange = true
	}
	if !data.Attribute14.Equal(state.Attribute14) {
		hasChange = true
	}
	if !data.Attribute14expr.Equal(state.Attribute14expr) {
		hasChange = true
	}
	if !data.Attribute14format.Equal(state.Attribute14format) {
		hasChange = true
	}
	if !data.Attribute14friendlyname.Equal(state.Attribute14friendlyname) {
		hasChange = true
	}
	if !data.Attribute15.Equal(state.Attribute15) {
		hasChange = true
	}
	if !data.Attribute15expr.Equal(state.Attribute15expr) {
		hasChange = true
	}
	if !data.Attribute15format.Equal(state.Attribute15format) {
		hasChange = true
	}
	if !data.Attribute15friendlyname.Equal(state.Attribute15friendlyname) {
		hasChange = true
	}
	if !data.Attribute16.Equal(state.Attribute16) {
		hasChange = true
	}
	if !data.Attribute16expr.Equal(state.Attribute16expr) {
		hasChange = true
	}
	if !data.Attribute16format.Equal(state.Attribute16format) {
		hasChange = true
	}
	if !data.Attribute16friendlyname.Equal(state.Attribute16friendlyname) {
		hasChange = true
	}
	if !data.Assertionconsumerserviceurl.Equal(state.Assertionconsumerserviceurl) {
		hasChange = true
	}
	if !data.Audience.Equal(state.Audience) {
		hasChange = true
	}
	if !data.Nameidexpr.Equal(state.Nameidexpr) {
		hasChange = true
	}
	if !data.Relaystaterule.Equal(state.Relaystaterule) {
		hasChange = true
	}
	if !data.Samlissuername.Equal(state.Samlissuername) {
		hasChange = true
	}
	if !data.Samlsigningcertname.Equal(state.Samlsigningcertname) {
		hasChange = true
	}
	if !data.Samlspcertname.Equal(state.Samlspcertname) {
		hasChange = true
	}
	if !data.Sendpassword.Equal(state.Sendpassword) {
		hasChange = true
	}
	if !data.Signatureservice.Equal(state.Signatureservice) {
		hasChange = true
	}
	if !data.Digestmethod.Equal(state.Digestmethod) {
		if config.Digestmethod.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "digestmethod")
		} else {
			hasChange = true
		}
	}
	if !data.Encryptassertion.Equal(state.Encryptassertion) {
		if config.Encryptassertion.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "encryptassertion")
		} else {
			hasChange = true
		}
	}
	if !data.Encryptionalgorithm.Equal(state.Encryptionalgorithm) {
		if config.Encryptionalgorithm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "encryptionalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Nameidformat.Equal(state.Nameidformat) {
		if config.Nameidformat.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "nameidformat")
		} else {
			hasChange = true
		}
	}
	if !data.Signassertion.Equal(state.Signassertion) {
		if config.Signassertion.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "signassertion")
		} else {
			hasChange = true
		}
	}
	if !data.Signaturealg.Equal(state.Signaturealg) {
		if config.Signaturealg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "signaturealg")
		} else {
			hasChange = true
		}
	}
	if !data.Skewtime.Equal(state.Skewtime) {
		if config.Skewtime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "skewtime")
		} else {
			hasChange = true
		}
	}

	vpnsamlssoprofileName := data.Name.ValueString()
	if hasChange {
		// Build the payload from the plan and push it.
		vpnsamlssoprofile := vpnsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName, &vpnsamlssoprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnsamlssoprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnsamlssoprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnsamlssoprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": vpnsamlssoprofileName,
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnsamlssoprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnsamlssoprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsamlssoprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsamlssoprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnsamlssoprofile resource")

	// Named resource - delete using DeleteResource
	vpnsamlssoprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnsamlssoprofile resource")
}

// Helper function to read vpnsamlssoprofile data from API
func (r *VpnsamlssoprofileResource) readVpnsamlssoprofileFromApi(ctx context.Context, data *VpnsamlssoprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	vpnsamlssoprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnsamlssoprofile, got error: %s", err))
		return false
	}

	vpnsamlssoprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
