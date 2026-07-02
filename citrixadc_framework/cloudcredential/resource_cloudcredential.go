package cloudcredential

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudcredentialResource{}
var _ resource.ResourceWithConfigure = (*CloudcredentialResource)(nil)
var _ resource.ResourceWithImportState = (*CloudcredentialResource)(nil)
var _ resource.ResourceWithValidateConfig = (*CloudcredentialResource)(nil)

func NewCloudcredentialResource() resource.Resource {
	return &CloudcredentialResource{}
}

// CloudcredentialResource defines the resource implementation.
type CloudcredentialResource struct {
	client *service.NitroClient
}

func (r *CloudcredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudcredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudcredential"
}

func (r *CloudcredentialResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudcredentialResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data CloudcredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either applicationsecret or applicationsecret_wo is specified
	if data.Applicationsecret.IsNull() && data.ApplicationsecretWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("applicationsecret"),
			"Missing Required Attribute",
			"Either \"applicationsecret\" or \"applicationsecret_wo\" must be specified.",
		)
	}
}

func (r *CloudcredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config CloudcredentialResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudcredential resource")
	// Get payload from plan (regular attributes)
	cloudcredential := cloudcredentialGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	cloudcredentialGetThePayloadFromtheConfig(ctx, &config, &cloudcredential)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudcredential.Type(), &cloudcredential)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudcredential, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudcredential resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("cloudcredential-config")

	// Read the updated state back
	r.readCloudcredentialFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudcredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudcredentialResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudcredential resource")

	r.readCloudcredentialFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudcredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CloudcredentialResourceModel

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

	tflog.Debug(ctx, "Updating cloudcredential resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Applicationid.Equal(state.Applicationid) {
		tflog.Debug(ctx, fmt.Sprintf("applicationid has changed for cloudcredential"))
		hasChange = true
	}
	// Check secret attribute applicationsecret or its version tracker
	if !data.Applicationsecret.Equal(state.Applicationsecret) {
		tflog.Debug(ctx, fmt.Sprintf("applicationsecret has changed for cloudcredential"))
		hasChange = true
	} else if !data.ApplicationsecretWoVersion.Equal(state.ApplicationsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("applicationsecret_wo_version has changed for cloudcredential"))
		hasChange = true
	}
	if !data.Tenantidentifier.Equal(state.Tenantidentifier) {
		tflog.Debug(ctx, fmt.Sprintf("tenantidentifier has changed for cloudcredential"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		cloudcredential := cloudcredentialGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		cloudcredentialGetThePayloadFromtheConfig(ctx, &config, &cloudcredential)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudcredential.Type(), &cloudcredential)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudcredential, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudcredential resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudcredential resource, skipping update")
	}

	// Read the updated state back
	r.readCloudcredentialFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudcredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudcredentialResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudcredential resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed cloudcredential from Terraform state")
}

// Helper function to read cloudcredential data from API
func (r *CloudcredentialResource) readCloudcredentialFromApi(ctx context.Context, data *CloudcredentialResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudcredential.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudcredential, got error: %s", err))
		return
	}

	cloudcredentialSetAttrFromGet(ctx, data, getResponseData)

}