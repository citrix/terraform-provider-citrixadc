package vpnsecureprivateaccessprofile

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
var _ resource.Resource = &VpnsecureprivateaccessprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnsecureprivateaccessprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnsecureprivateaccessprofileResource)(nil)

func NewVpnsecureprivateaccessprofileResource() resource.Resource {
	return &VpnsecureprivateaccessprofileResource{}
}

// VpnsecureprivateaccessprofileResource defines the resource implementation.
type VpnsecureprivateaccessprofileResource struct {
	client *service.NitroClient
}

func (r *VpnsecureprivateaccessprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnsecureprivateaccessprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnsecureprivateaccessprofile"
}

func (r *VpnsecureprivateaccessprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnsecureprivateaccessprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config VpnsecureprivateaccessprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnsecureprivateaccessprofile resource")
	// Get payload from plan (regular attributes)
	vpnsecureprivateaccessprofile := vpnsecureprivateaccessprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	vpnsecureprivateaccessprofileGetThePayloadFromtheConfig(ctx, &config, &vpnsecureprivateaccessprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnsecureprivateaccessprofile.Type(), name_value, &vpnsecureprivateaccessprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnsecureprivateaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnsecureprivateaccessprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readVpnsecureprivateaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsecureprivateaccessprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsecureprivateaccessprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnsecureprivateaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnsecureprivateaccessprofile resource")

	found := r.readVpnsecureprivateaccessprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnsecureprivateaccessprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnsecureprivateaccessprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnsecureprivateaccessprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Collect eligible attributes that were removed from config so they can be unset on the appliance
	attributesToUnset := []string{}
	if !data.Url.Equal(state.Url) {
		tflog.Debug(ctx, fmt.Sprintf("url has changed for vpnsecureprivateaccessprofile"))
		hasChange = true
	}
	if !data.Customerid.Equal(state.Customerid) {
		tflog.Debug(ctx, fmt.Sprintf("customerid has changed for vpnsecureprivateaccessprofile"))
		if config.Customerid.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "customerid")
		} else {
			hasChange = true
		}
	}
	if !data.Chromeenterprisepremiummode.Equal(state.Chromeenterprisepremiummode) {
		tflog.Debug(ctx, fmt.Sprintf("chromeenterprisepremiummode has changed for vpnsecureprivateaccessprofile"))
		if config.Chromeenterprisepremiummode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "chromeenterprisepremiummode")
		} else {
			hasChange = true
		}
	}
	if !data.Googlecustomerid.Equal(state.Googlecustomerid) {
		tflog.Debug(ctx, fmt.Sprintf("googlecustomerid has changed for vpnsecureprivateaccessprofile"))
		if config.Googlecustomerid.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "googlecustomerid")
		} else {
			hasChange = true
		}
	}
	if !data.Googlesecuritygatewayid.Equal(state.Googlesecuritygatewayid) {
		tflog.Debug(ctx, fmt.Sprintf("googlesecuritygatewayid has changed for vpnsecureprivateaccessprofile"))
		if config.Googlesecuritygatewayid.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "googlesecuritygatewayid")
		} else {
			hasChange = true
		}
	}
	if !data.Forceclienttype.Equal(state.Forceclienttype) {
		tflog.Debug(ctx, fmt.Sprintf("forceclienttype has changed for vpnsecureprivateaccessprofile"))
		if config.Forceclienttype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "forceclienttype")
		} else {
			hasChange = true
		}
	}
	// Check secret attribute sharedsecret or its version tracker
	if !data.Sharedsecret.Equal(state.Sharedsecret) {
		tflog.Debug(ctx, fmt.Sprintf("sharedsecret has changed for vpnsecureprivateaccessprofile"))
		hasChange = true
	} else if !data.SharedsecretWoVersion.Equal(state.SharedsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sharedsecret_wo_version has changed for vpnsecureprivateaccessprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		vpnsecureprivateaccessprofile := vpnsecureprivateaccessprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		vpnsecureprivateaccessprofileGetThePayloadFromtheConfig(ctx, &config, &vpnsecureprivateaccessprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpnsecureprivateaccessprofile.Type(), name_value, &vpnsecureprivateaccessprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnsecureprivateaccessprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnsecureprivateaccessprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnsecureprivateaccessprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them to their defaults.
	// Update-then-unset ordering ensures any default carried in the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnsecureprivateaccessprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnsecureprivateaccessprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpnsecureprivateaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsecureprivateaccessprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsecureprivateaccessprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnsecureprivateaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnsecureprivateaccessprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Vpnsecureprivateaccessprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnsecureprivateaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnsecureprivateaccessprofile resource")
}

// Helper function to read vpnsecureprivateaccessprofile data from API
func (r *VpnsecureprivateaccessprofileResource) readVpnsecureprivateaccessprofileFromApi(ctx context.Context, data *VpnsecureprivateaccessprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Vpnsecureprivateaccessprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnsecureprivateaccessprofile, got error: %s", err))
		return false
	}

	vpnsecureprivateaccessprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// UpgradeState migrates pre-write-only state (GH #1441): it seeds the
// "sharedsecret_wo_version" tracker attribute to 1 when the stored state has no
// value for it, so the schema Default does not plan a spurious "null -> 1"
// update after upgrading the provider. Paired with the schema Version bump so the
// upgrade path actually runs. See utils.WoVersionUpgradeState.
