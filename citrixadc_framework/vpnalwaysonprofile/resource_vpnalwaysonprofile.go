package vpnalwaysonprofile

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
var _ resource.Resource = &VpnalwaysonprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnalwaysonprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnalwaysonprofileResource)(nil)

func NewVpnalwaysonprofileResource() resource.Resource {
	return &VpnalwaysonprofileResource{}
}

// VpnalwaysonprofileResource defines the resource implementation.
type VpnalwaysonprofileResource struct {
	client *service.NitroClient
}

func (r *VpnalwaysonprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnalwaysonprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnalwaysonprofile"
}

func (r *VpnalwaysonprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnalwaysonprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnalwaysonprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnalwaysonprofile resource")

	vpnalwaysonprofile := vpnalwaysonprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnalwaysonprofile.Type(), name_value, &vpnalwaysonprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnalwaysonprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnalwaysonprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readVpnalwaysonprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnalwaysonprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnalwaysonprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnalwaysonprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnalwaysonprofile resource")

	found := r.readVpnalwaysonprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnalwaysonprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnalwaysonprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnalwaysonprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Clientcontrol.Equal(state.Clientcontrol) {
		tflog.Debug(ctx, "clientcontrol has changed for vpnalwaysonprofile")
		if config.Clientcontrol.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientcontrol")
		} else {
			hasChange = true
		}
	}
	if !data.Locationbasedvpn.Equal(state.Locationbasedvpn) {
		tflog.Debug(ctx, "locationbasedvpn has changed for vpnalwaysonprofile")
		if config.Locationbasedvpn.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "locationbasedvpn")
		} else {
			hasChange = true
		}
	}
	if !data.Networkaccessonvpnfailure.Equal(state.Networkaccessonvpnfailure) {
		tflog.Debug(ctx, "networkaccessonvpnfailure has changed for vpnalwaysonprofile")
		if config.Networkaccessonvpnfailure.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "networkaccessonvpnfailure")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		vpnalwaysonprofile := vpnalwaysonprofileGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		name_value := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Vpnalwaysonprofile.Type(), name_value, &vpnalwaysonprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnalwaysonprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnalwaysonprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnalwaysonprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnalwaysonprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnalwaysonprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpnalwaysonprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnalwaysonprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnalwaysonprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnalwaysonprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnalwaysonprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnalwaysonprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnalwaysonprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnalwaysonprofile resource")
}

// Helper function to read vpnalwaysonprofile data from API
func (r *VpnalwaysonprofileResource) readVpnalwaysonprofileFromApi(ctx context.Context, data *VpnalwaysonprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnalwaysonprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnalwaysonprofile, got error: %s", err))
		return false
	}

	vpnalwaysonprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
