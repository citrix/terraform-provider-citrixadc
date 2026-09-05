package appqoeaction

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
var _ resource.Resource = &AppqoeactionResource{}
var _ resource.ResourceWithConfigure = (*AppqoeactionResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoeactionResource)(nil)

func NewAppqoeactionResource() resource.Resource {
	return &AppqoeactionResource{}
}

// AppqoeactionResource defines the resource implementation.
type AppqoeactionResource struct {
	client *service.NitroClient
}

func (r *AppqoeactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoeactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoeaction"
}

func (r *AppqoeactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoeactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoeactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoeaction resource")
	// Get payload from plan
	appqoeaction := appqoeactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Appqoeaction.Type(), name_value, &appqoeaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoeaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appqoeaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppqoeactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoeaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoeactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoeaction resource")

	found := r.readAppqoeactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppqoeactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppqoeactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appqoeaction resource")

	// Check if there are any changes in NITRO-updatable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Altcontentpath.Equal(state.Altcontentpath) {
		tflog.Debug(ctx, "altcontentpath has changed for appqoeaction")
		hasChange = true
	}
	if !data.Altcontentsvcname.Equal(state.Altcontentsvcname) {
		tflog.Debug(ctx, "altcontentsvcname has changed for appqoeaction")
		hasChange = true
	}
	if !data.Delay.Equal(state.Delay) {
		tflog.Debug(ctx, "delay has changed for appqoeaction")
		hasChange = true
	}
	if !data.Dosaction.Equal(state.Dosaction) {
		tflog.Debug(ctx, "dosaction has changed for appqoeaction")
		hasChange = true
	}
	if !data.Dostrigexpression.Equal(state.Dostrigexpression) {
		tflog.Debug(ctx, "dostrigexpression has changed for appqoeaction")
		hasChange = true
	}
	if !data.Maxconn.Equal(state.Maxconn) {
		tflog.Debug(ctx, "maxconn has changed for appqoeaction")
		hasChange = true
	}
	if !data.Numretries.Equal(state.Numretries) {
		tflog.Debug(ctx, "numretries has changed for appqoeaction")
		if config.Numretries.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "numretries")
		} else {
			hasChange = true
		}
	}
	if !data.Polqdepth.Equal(state.Polqdepth) {
		tflog.Debug(ctx, "polqdepth has changed for appqoeaction")
		hasChange = true
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for appqoeaction")
		hasChange = true
	}
	if !data.Priqdepth.Equal(state.Priqdepth) {
		tflog.Debug(ctx, "priqdepth has changed for appqoeaction")
		hasChange = true
	}
	if !data.Retryonreset.Equal(state.Retryonreset) {
		tflog.Debug(ctx, "retryonreset has changed for appqoeaction")
		if config.Retryonreset.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "retryonreset")
		} else {
			hasChange = true
		}
	}
	if !data.Retryontimeout.Equal(state.Retryontimeout) {
		tflog.Debug(ctx, "retryontimeout has changed for appqoeaction")
		hasChange = true
	}
	if !data.Tcpprofile.Equal(state.Tcpprofile) {
		tflog.Debug(ctx, "tcpprofile has changed for appqoeaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		appqoeaction := appqoeactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Appqoeaction.Type(), name_value, &appqoeaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appqoeaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appqoeaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appqoeaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Appqoeaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset appqoeaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAppqoeactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoeaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoeactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoeaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appqoeaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appqoeaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appqoeaction resource")
}

// Helper function to read appqoeaction data from API
func (r *AppqoeactionResource) readAppqoeactionFromApi(ctx context.Context, data *AppqoeactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	appqoeaction_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appqoeaction.Type(), appqoeaction_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoeaction, got error: %s", err))
		return false
	}

	appqoeactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
