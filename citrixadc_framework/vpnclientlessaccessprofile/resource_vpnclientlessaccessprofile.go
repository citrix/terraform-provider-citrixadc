package vpnclientlessaccessprofile

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
var _ resource.Resource = &VpnclientlessaccessprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnclientlessaccessprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnclientlessaccessprofileResource)(nil)

func NewVpnclientlessaccessprofileResource() resource.Resource {
	return &VpnclientlessaccessprofileResource{}
}

// VpnclientlessaccessprofileResource defines the resource implementation.
type VpnclientlessaccessprofileResource struct {
	client *service.NitroClient
}

func (r *VpnclientlessaccessprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnclientlessaccessprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnclientlessaccessprofile"
}

func (r *VpnclientlessaccessprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnclientlessaccessprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnclientlessaccessprofile resource")

	// Create API request body from the model
	vpnclientlessaccessprofile := vpnclientlessaccessprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	profilename_value := data.Profilename.ValueString()
	_, err := r.client.AddResource(service.Vpnclientlessaccessprofile.Type(), profilename_value, &vpnclientlessaccessprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnclientlessaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnclientlessaccessprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(profilename_value)

	// Read the updated state back
	if !r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnclientlessaccessprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccessprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnclientlessaccessprofile resource")

	found := r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnclientlessaccessprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnclientlessaccessprofileResourceModel

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

	tflog.Debug(ctx, "Updating vpnclientlessaccessprofile resource")

	// Determine whether a normal update is needed and which attributes were
	// removed from config and must be unset (reverted to NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Requirepersistentcookie.Equal(state.Requirepersistentcookie) {
		tflog.Debug(ctx, "requirepersistentcookie has changed for vpnclientlessaccessprofile")
		if config.Requirepersistentcookie.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "requirepersistentcookie")
		} else {
			hasChange = true
		}
	}
	// Other mutable attributes always go through UpdateResource when changed.
	if !data.Clientconsumedcookies.Equal(state.Clientconsumedcookies) ||
		!data.Javascriptrewritepolicylabel.Equal(state.Javascriptrewritepolicylabel) ||
		!data.Regexforfindingcustomurls.Equal(state.Regexforfindingcustomurls) ||
		!data.Regexforfindingurlincss.Equal(state.Regexforfindingurlincss) ||
		!data.Regexforfindingurlinjavascript.Equal(state.Regexforfindingurlinjavascript) ||
		!data.Regexforfindingurlinxcomponent.Equal(state.Regexforfindingurlinxcomponent) ||
		!data.Regexforfindingurlinxml.Equal(state.Regexforfindingurlinxml) ||
		!data.Reqhdrrewritepolicylabel.Equal(state.Reqhdrrewritepolicylabel) ||
		!data.Reshdrrewritepolicylabel.Equal(state.Reshdrrewritepolicylabel) ||
		!data.Urlrewritepolicylabel.Equal(state.Urlrewritepolicylabel) {
		hasChange = true
	}

	// Named resource - use UpdateResource
	profilename_value := data.Profilename.ValueString()
	if hasChange {
		// Create API request body from the model
		vpnclientlessaccessprofile := vpnclientlessaccessprofileGetThePayloadFromtheConfig(ctx, &data)
		_, err := r.client.UpdateResource(service.Vpnclientlessaccessprofile.Type(), profilename_value, &vpnclientlessaccessprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnclientlessaccessprofile, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpnclientlessaccessprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnclientlessaccessprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"profilename": data.Profilename.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnclientlessaccessprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnclientlessaccessprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnclientlessaccessprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccessprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnclientlessaccessprofile resource")

	// Named resource - delete using DeleteResource
	profilename_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnclientlessaccessprofile.Type(), profilename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnclientlessaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnclientlessaccessprofile resource")
}

// Helper function to read vpnclientlessaccessprofile data from API
func (r *VpnclientlessaccessprofileResource) readVpnclientlessaccessprofileFromApi(ctx context.Context, data *VpnclientlessaccessprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	profilename_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnclientlessaccessprofile.Type(), profilename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnclientlessaccessprofile, got error: %s", err))
		return false
	}

	vpnclientlessaccessprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
