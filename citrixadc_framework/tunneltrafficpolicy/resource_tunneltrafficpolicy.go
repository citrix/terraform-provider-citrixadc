package tunneltrafficpolicy

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
var _ resource.Resource = &TunneltrafficpolicyResource{}
var _ resource.ResourceWithConfigure = (*TunneltrafficpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TunneltrafficpolicyResource)(nil)

func NewTunneltrafficpolicyResource() resource.Resource {
	return &TunneltrafficpolicyResource{}
}

// TunneltrafficpolicyResource defines the resource implementation.
type TunneltrafficpolicyResource struct {
	client *service.NitroClient
}

func (r *TunneltrafficpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TunneltrafficpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunneltrafficpolicy"
}

func (r *TunneltrafficpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TunneltrafficpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tunneltrafficpolicy resource")

	tunneltrafficpolicy := tunneltrafficpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tunneltrafficpolicy.Type(), name_value, &tunneltrafficpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tunneltrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tunneltrafficpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tunneltrafficpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunneltrafficpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tunneltrafficpolicy resource")

	found := r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TunneltrafficpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TunneltrafficpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to distinguish an attribute removed from config (-> unset) from
	// one merely changed (-> update).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tunneltrafficpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for tunneltrafficpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for tunneltrafficpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for tunneltrafficpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for tunneltrafficpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		tunneltrafficpolicy := tunneltrafficpolicyGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Tunneltrafficpolicy.Type(), name_value, &tunneltrafficpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tunneltrafficpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated tunneltrafficpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tunneltrafficpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Tunneltrafficpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset tunneltrafficpolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tunneltrafficpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunneltrafficpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tunneltrafficpolicy resource")

	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Tunneltrafficpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tunneltrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tunneltrafficpolicy resource")
}

// Helper function to read tunneltrafficpolicy data from API
func (r *TunneltrafficpolicyResource) readTunneltrafficpolicyFromApi(ctx context.Context, data *TunneltrafficpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	tunneltrafficpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tunneltrafficpolicy.Type(), tunneltrafficpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tunneltrafficpolicy, got error: %s", err))
		return false
	}

	tunneltrafficpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
