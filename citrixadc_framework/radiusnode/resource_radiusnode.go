package radiusnode

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
var _ resource.Resource = &RadiusnodeResource{}
var _ resource.ResourceWithConfigure = (*RadiusnodeResource)(nil)
var _ resource.ResourceWithImportState = (*RadiusnodeResource)(nil)
var _ resource.ResourceWithValidateConfig = (*RadiusnodeResource)(nil)

func NewRadiusnodeResource() resource.Resource {
	return &RadiusnodeResource{}
}

// RadiusnodeResource defines the resource implementation.
type RadiusnodeResource struct {
	client *service.NitroClient
}

func (r *RadiusnodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RadiusnodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radiusnode"
}

func (r *RadiusnodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RadiusnodeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RadiusnodeResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either radkey or radkey_wo is specified
	if data.Radkey.IsNull() && data.RadkeyWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("radkey"),
			"Missing Required Attribute",
			"Either \"radkey\" or \"radkey_wo\" must be specified.",
		)
	}
}

func (r *RadiusnodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config RadiusnodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating radiusnode resource")
	// Get payload from plan (regular attributes)
	radiusnode := radiusnodeGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	radiusnodeGetThePayloadFromtheConfig(ctx, &config, &radiusnode)

	// Make API call
	// Named resource - use AddResource
	nodeprefix_value := data.Nodeprefix.ValueString()
	_, err := r.client.AddResource(service.Radiusnode.Type(), nodeprefix_value, &radiusnode)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create radiusnode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created radiusnode resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Nodeprefix.ValueString()))

	// Read the updated state back
	if !r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "radiusnode not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusnodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading radiusnode resource")

	found := r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RadiusnodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RadiusnodeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating radiusnode resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute radkey or its version tracker
	if !data.Radkey.Equal(state.Radkey) {
		tflog.Debug(ctx, fmt.Sprintf("radkey has changed for radiusnode"))
		hasChange = true
	} else if !data.RadkeyWoVersion.Equal(state.RadkeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("radkey_wo_version has changed for radiusnode"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		radiusnode := radiusnodeGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		radiusnodeGetThePayloadFromtheConfig(ctx, &config, &radiusnode)
		// Make API call
		// Named resource - use UpdateResource
		nodeprefix_value := data.Nodeprefix.ValueString()
		_, err := r.client.UpdateResource(service.Radiusnode.Type(), nodeprefix_value, &radiusnode)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update radiusnode, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated radiusnode resource")
	} else {
		tflog.Debug(ctx, "No changes detected for radiusnode resource, skipping update")
	}

	// Read the updated state back
	if !r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "radiusnode not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusnodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting radiusnode resource")
	// Named resource - delete using DeleteResource
	nodeprefix_value := data.Nodeprefix.ValueString()
	err := r.client.DeleteResource(service.Radiusnode.Type(), nodeprefix_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete radiusnode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted radiusnode resource")
}

// Helper function to read radiusnode data from API
func (r *RadiusnodeResource) readRadiusnodeFromApi(ctx context.Context, data *RadiusnodeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nodeprefix_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Radiusnode.Type(), nodeprefix_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read radiusnode, got error: %s", err))
		return false
	}

	radiusnodeSetAttrFromGet(ctx, data, getResponseData)

	return true

}
