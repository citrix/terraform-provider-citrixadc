package cloudparameter

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
var _ resource.Resource = &CloudparameterResource{}
var _ resource.ResourceWithConfigure = (*CloudparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CloudparameterResource)(nil)

func NewCloudparameterResource() resource.Resource {
	return &CloudparameterResource{}
}

// CloudparameterResource defines the resource implementation.
type CloudparameterResource struct {
	client *service.NitroClient
}

func (r *CloudparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudparameter"
}

func (r *CloudparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudparameter resource")
	cloudparameter := cloudparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudparameter.Type(), &cloudparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("cloudparameter-config")

	// Read the updated state back
	r.readCloudparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudparameter resource")

	r.readCloudparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Activationcode.Equal(state.Activationcode) {
		tflog.Debug(ctx, fmt.Sprintf("activationcode has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Connectorresidence.Equal(state.Connectorresidence) {
		tflog.Debug(ctx, fmt.Sprintf("connectorresidence has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Controllerfqdn.Equal(state.Controllerfqdn) {
		tflog.Debug(ctx, fmt.Sprintf("controllerfqdn has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Controllerport.Equal(state.Controllerport) {
		tflog.Debug(ctx, fmt.Sprintf("controllerport has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Customerid.Equal(state.Customerid) {
		tflog.Debug(ctx, fmt.Sprintf("customerid has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Deployment.Equal(state.Deployment) {
		tflog.Debug(ctx, fmt.Sprintf("deployment has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Instanceid.Equal(state.Instanceid) {
		tflog.Debug(ctx, fmt.Sprintf("instanceid has changed for cloudparameter"))
		hasChange = true
	}
	if !data.Resourcelocation.Equal(state.Resourcelocation) {
		tflog.Debug(ctx, fmt.Sprintf("resourcelocation has changed for cloudparameter"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudparameter := cloudparameterGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudparameter.Type(), &cloudparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudparameter resource, skipping update")
	}

	// Read the updated state back
	r.readCloudparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed cloudparameter from Terraform state")
}

// Helper function to read cloudparameter data from API
func (r *CloudparameterResource) readCloudparameterFromApi(ctx context.Context, data *CloudparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudparameter, got error: %s", err))
		return
	}

	cloudparameterSetAttrFromGet(ctx, data, getResponseData)

}
