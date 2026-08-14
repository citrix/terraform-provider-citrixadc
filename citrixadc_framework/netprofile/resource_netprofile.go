package netprofile

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

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NetprofileResource{}
var _ resource.ResourceWithConfigure = (*NetprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NetprofileResource)(nil)

func NewNetprofileResource() resource.Resource {
	return &NetprofileResource{}
}

// NetprofileResource defines the resource implementation.
type NetprofileResource struct {
	client *service.NitroClient
}

func (r *NetprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netprofile"
}

func (r *NetprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netprofile resource")

	// Backward-compatible with the SDK v2 resource: name is Optional+Computed.
	// When the user does not supply a name, generate a unique one (SDK v2 used
	// resource.PrefixedUniqueId("tf-netprofile-")).
	netprofileName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || netprofileName == "" {
		netprofileName = sdkid.PrefixedUniqueId("tf-netprofile-")
		data.Name = types.StringValue(netprofileName)
	}

	// Create API request body from the model
	netprofile := netprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Netprofile.Type(), netprofileName, &netprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created netprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(netprofileName)

	// Read the updated state back
	if !r.readNetprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "netprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netprofile resource")

	found := r.readNetprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NetprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NetprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset them)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating netprofile resource")

	// Check if there are any changes in updateable attributes. td is create-only
	// at the NITRO layer (excluded from the update payload) and is not ForceNew in
	// SDK v2, so it is intentionally not part of change detection.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Mbf.Equal(state.Mbf) {
		tflog.Debug(ctx, "mbf has changed for netprofile")
		hasChange = true
	}
	if !data.Overridelsn.Equal(state.Overridelsn) {
		tflog.Debug(ctx, "overridelsn has changed for netprofile")
		if config.Overridelsn.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "overridelsn")
		} else {
			hasChange = true
		}
	}
	if !data.Proxyprotocol.Equal(state.Proxyprotocol) {
		tflog.Debug(ctx, "proxyprotocol has changed for netprofile")
		if config.Proxyprotocol.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxyprotocol")
		} else {
			hasChange = true
		}
	}
	if !data.Proxyprotocolaftertlshandshake.Equal(state.Proxyprotocolaftertlshandshake) {
		tflog.Debug(ctx, "proxyprotocolaftertlshandshake has changed for netprofile")
		if config.Proxyprotocolaftertlshandshake.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxyprotocolaftertlshandshake")
		} else {
			hasChange = true
		}
	}
	if !data.Proxyprotocoltxversion.Equal(state.Proxyprotocoltxversion) {
		tflog.Debug(ctx, "proxyprotocoltxversion has changed for netprofile")
		if config.Proxyprotocoltxversion.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxyprotocoltxversion")
		} else {
			hasChange = true
		}
	}
	if !data.Srcip.Equal(state.Srcip) {
		tflog.Debug(ctx, "srcip has changed for netprofile")
		hasChange = true
	}
	if !data.Srcippersistency.Equal(state.Srcippersistency) {
		tflog.Debug(ctx, "srcippersistency has changed for netprofile")
		if config.Srcippersistency.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "srcippersistency")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body restricted to NITRO-updatable fields
		netprofile := netprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource (NITRO update is HTTP PUT)
		_, err := r.client.UpdateResource(service.Netprofile.Type(), data.Id.ValueString(), &netprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated netprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for netprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Netprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset netprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNetprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "netprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netprofile resource")

	// Named resource - delete using DeleteResource (NITRO delete is HTTP DELETE)
	err := r.client.DeleteResource(service.Netprofile.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete netprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted netprofile resource")
}

// Helper function to read netprofile data from API
func (r *NetprofileResource) readNetprofileFromApi(ctx context.Context, data *NetprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	netprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Netprofile.Type(), netprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netprofile, got error: %s", err))
		return false
	}

	netprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
