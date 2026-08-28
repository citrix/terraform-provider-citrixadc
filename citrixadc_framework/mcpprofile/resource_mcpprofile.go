package mcpprofile

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
var _ resource.Resource = &McpprofileResource{}
var _ resource.ResourceWithConfigure = (*McpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*McpprofileResource)(nil)

func NewMcpprofileResource() resource.Resource {
	return &McpprofileResource{}
}

// McpprofileResource defines the resource implementation.
type McpprofileResource struct {
	client *service.NitroClient
}

func (r *McpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *McpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mcpprofile"
}

func (r *McpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *McpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config McpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in the plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mcpprofile resource")
	// Get payload from plan (regular attributes)
	mcpprofile := mcpprofileGetThePayloadFromthePlan(ctx, &data)
	// Overlay write-only attributes from config (tokenorapi_wo -> tokenorapi)
	mcpprofileGetThePayloadFromtheConfig(ctx, &config, &mcpprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Mcpprofile.Type(), name_value, &mcpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created mcpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readMcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mcpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *McpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data McpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mcpprofile resource")

	found := r.readMcpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *McpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state McpprofileResourceModel

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

	tflog.Debug(ctx, "Updating mcpprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Collect eligible attributes that were removed from config so they can be unset on the appliance
	attributesToUnset := []string{}
	if !data.Proxymode.Equal(state.Proxymode) {
		tflog.Debug(ctx, fmt.Sprintf("proxymode has changed for mcpprofile"))
		if config.Proxymode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxymode")
		} else {
			hasChange = true
		}
	}
	if !data.Hostreplacement.Equal(state.Hostreplacement) {
		tflog.Debug(ctx, fmt.Sprintf("hostreplacement has changed for mcpprofile"))
		if config.Hostreplacement.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "hostreplacement")
		} else {
			hasChange = true
		}
	}
	if !data.Urlreplacement.Equal(state.Urlreplacement) {
		tflog.Debug(ctx, fmt.Sprintf("urlreplacement has changed for mcpprofile"))
		if config.Urlreplacement.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "urlreplacement")
		} else {
			hasChange = true
		}
	}
	if !data.Protocolversion.Equal(state.Protocolversion) {
		tflog.Debug(ctx, fmt.Sprintf("protocolversion has changed for mcpprofile"))
		if config.Protocolversion.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "protocolversion")
		} else {
			hasChange = true
		}
	}
	// tokenorapi is a secret: it is applied via the plaintext attribute or the
	// write-only tokenorapi_wo path (signalled by a tokenorapi_wo_version bump). It
	// is not part of the unset flow.
	if !data.Tokenorapi.Equal(state.Tokenorapi) {
		tflog.Debug(ctx, "tokenorapi has changed for mcpprofile")
		hasChange = true
	} else if !data.TokenorapiWoVersion.Equal(state.TokenorapiWoVersion) {
		tflog.Debug(ctx, "tokenorapi_wo_version has changed for mcpprofile")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, fmt.Sprintf("comment has changed for mcpprofile"))
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Insertheaderinclientrequest.Equal(state.Insertheaderinclientrequest) {
		tflog.Debug(ctx, fmt.Sprintf("insertheaderinclientrequest has changed for mcpprofile"))
		if config.Insertheaderinclientrequest.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "insertheaderinclientrequest")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		mcpprofile := mcpprofileGetThePayloadFromthePlan(ctx, &data)
		// Overlay write-only attributes from config (tokenorapi_wo -> tokenorapi)
		mcpprofileGetThePayloadFromtheConfig(ctx, &config, &mcpprofile)
		// profiletype is a create-only attribute (RequiresReplace) and is not a valid
		// member of the NITRO update payload, so drop it from the PUT request.
		mcpprofile.Profiletype = ""
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Mcpprofile.Type(), name_value, &mcpprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update mcpprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated mcpprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for mcpprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them to their defaults.
	// Update-then-unset ordering ensures any default carried in the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Mcpprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset mcpprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readMcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mcpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *McpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data McpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mcpprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Mcpprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete mcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted mcpprofile resource")
}

// Helper function to read mcpprofile data from API
func (r *McpprofileResource) readMcpprofileFromApi(ctx context.Context, data *McpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Mcpprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mcpprofile, got error: %s", err))
		return false
	}

	mcpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
