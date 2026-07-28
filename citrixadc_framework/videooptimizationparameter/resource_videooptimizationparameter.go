package videooptimizationparameter

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VideooptimizationparameterResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationparameterResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationparameterResource)(nil)

func NewVideooptimizationparameterResource() resource.Resource {
	return &VideooptimizationparameterResource{}
}

// VideooptimizationparameterResource defines the resource implementation.
type VideooptimizationparameterResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationparameter"
}

func (r *VideooptimizationparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationparameter resource")
	videooptimizationparameter := videooptimizationparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Videooptimizationparameter.Type(), &videooptimizationparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("videooptimizationparameter-config")

	// Read the updated state back
	r.readVideooptimizationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationparameter resource")

	r.readVideooptimizationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state, config VideooptimizationparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating videooptimizationparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Quicpacingrate.Equal(state.Quicpacingrate) {
		tflog.Debug(ctx, fmt.Sprintf("quicpacingrate has changed for videooptimizationparameter"))
		hasChange = true
	}
	if !data.Randomsamplingpercentage.Equal(state.Randomsamplingpercentage) {
		tflog.Debug(ctx, fmt.Sprintf("randomsamplingpercentage has changed for videooptimizationparameter"))
		if config.Randomsamplingpercentage.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "randomsamplingpercentage")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		videooptimizationparameter := videooptimizationparameterGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Videooptimizationparameter.Type(), &videooptimizationparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated videooptimizationparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for videooptimizationparameter resource, skipping update")
	}

	// Unset any attributes that were removed from configuration so the appliance
	// reverts them to their defaults. Singleton resource - no identity fields.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Videooptimizationparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset videooptimizationparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readVideooptimizationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed videooptimizationparameter from Terraform state")
}

// Helper function to read videooptimizationparameter data from API
func (r *VideooptimizationparameterResource) readVideooptimizationparameterFromApi(ctx context.Context, data *VideooptimizationparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Videooptimizationparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationparameter, got error: %s", err))
		return
	}

	videooptimizationparameterSetAttrFromGet(ctx, data, getResponseData)

}
