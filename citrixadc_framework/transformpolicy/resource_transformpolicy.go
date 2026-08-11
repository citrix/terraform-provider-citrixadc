package transformpolicy

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
var _ resource.Resource = &TransformpolicyResource{}
var _ resource.ResourceWithConfigure = (*TransformpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TransformpolicyResource)(nil)

func NewTransformpolicyResource() resource.Resource {
	return &TransformpolicyResource{}
}

// TransformpolicyResource defines the resource implementation.
type TransformpolicyResource struct {
	client *service.NitroClient
}

func (r *TransformpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformpolicy"
}

func (r *TransformpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformpolicy resource")

	// Create API request body from the model
	transformpolicy := transformpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	transformpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Transformpolicy.Type(), transformpolicyName, &transformpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created transformpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(transformpolicyName)

	// Read the updated state back
	if !r.readTransformpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformpolicy resource")

	found := r.readTransformpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TransformpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TransformpolicyResourceModel

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

	tflog.Debug(ctx, "Updating transformpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for transformpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for transformpolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Profilename.Equal(state.Profilename) {
		tflog.Debug(ctx, "profilename has changed for transformpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for transformpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		transformpolicy := transformpolicyGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		transformpolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Transformpolicy.Type(), transformpolicyName, &transformpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated transformpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for transformpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Transformpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset transformpolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readTransformpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformpolicy resource")

	// Named resource - delete using DeleteResource
	transformpolicyName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Transformpolicy.Type(), transformpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete transformpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted transformpolicy resource")
}

// Helper function to read transformpolicy data from API
func (r *TransformpolicyResource) readTransformpolicyFromApi(ctx context.Context, data *TransformpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	transformpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Transformpolicy.Type(), transformpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformpolicy, got error: %s", err))
		return false
	}

	transformpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
