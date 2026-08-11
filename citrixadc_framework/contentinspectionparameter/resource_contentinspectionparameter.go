package contentinspectionparameter

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
var _ resource.Resource = &ContentinspectionparameterResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionparameterResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionparameterResource)(nil)

func NewContentinspectionparameterResource() resource.Resource {
	return &ContentinspectionparameterResource{}
}

// ContentinspectionparameterResource defines the resource implementation.
type ContentinspectionparameterResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionparameter"
}

func (r *ContentinspectionparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionparameter resource")

	// Create API request body from the model
	contentinspectionparameter := contentinspectionparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Contentinspectionparameter.Type(), &contentinspectionparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("contentinspectionparameter-config")

	tflog.Trace(ctx, "Created contentinspectionparameter resource")

	// Read the updated state back
	r.readContentinspectionparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionparameter resource")

	r.readContentinspectionparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ContentinspectionparameterResourceModel

	// Read Terraform prior state to preserve ID and detect changes
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

	tflog.Debug(ctx, "Updating contentinspectionparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for contentinspectionparameter")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		contentinspectionparameter := contentinspectionparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Contentinspectionparameter.Type(), &contentinspectionparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectionparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Contentinspectionparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset contentinspectionparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readContentinspectionparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionparameter resource")

	// For contentinspectionparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted contentinspectionparameter resource from state")
}

// Helper function to read contentinspectionparameter data from API
func (r *ContentinspectionparameterResource) readContentinspectionparameterFromApi(ctx context.Context, data *ContentinspectionparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Contentinspectionparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionparameter, got error: %s", err))
		return
	}

	contentinspectionparameterSetAttrFromGet(ctx, data, getResponseData)

}
