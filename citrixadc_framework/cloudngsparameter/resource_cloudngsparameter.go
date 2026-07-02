package cloudngsparameter

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
var _ resource.Resource = &CloudngsparameterResource{}
var _ resource.ResourceWithConfigure = (*CloudngsparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CloudngsparameterResource)(nil)

func NewCloudngsparameterResource() resource.Resource {
	return &CloudngsparameterResource{}
}

// CloudngsparameterResource defines the resource implementation.
type CloudngsparameterResource struct {
	client *service.NitroClient
}

func (r *CloudngsparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudngsparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudngsparameter"
}

func (r *CloudngsparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudngsparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudngsparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudngsparameter resource")
	cloudngsparameter := cloudngsparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudngsparameter.Type(), &cloudngsparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudngsparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudngsparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("cloudngsparameter-config")

	// Read the updated state back
	r.readCloudngsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudngsparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudngsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudngsparameter resource")

	r.readCloudngsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudngsparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudngsparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudngsparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Allowdtls12.Equal(state.Allowdtls12) {
		tflog.Debug(ctx, fmt.Sprintf("allowdtls12 has changed for cloudngsparameter"))
		hasChange = true
	}
	if !data.Allowedudtversion.Equal(state.Allowedudtversion) {
		tflog.Debug(ctx, fmt.Sprintf("allowedudtversion has changed for cloudngsparameter"))
		hasChange = true
	}
	if !data.Blockonallowedngstktprof.Equal(state.Blockonallowedngstktprof) {
		tflog.Debug(ctx, fmt.Sprintf("blockonallowedngstktprof has changed for cloudngsparameter"))
		hasChange = true
	}
	if !data.Csvserverticketingdecouple.Equal(state.Csvserverticketingdecouple) {
		tflog.Debug(ctx, fmt.Sprintf("csvserverticketingdecouple has changed for cloudngsparameter"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudngsparameter := cloudngsparameterGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudngsparameter.Type(), &cloudngsparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudngsparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudngsparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudngsparameter resource, skipping update")
	}

	// Read the updated state back
	r.readCloudngsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudngsparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudngsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudngsparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed cloudngsparameter from Terraform state")
}

// Helper function to read cloudngsparameter data from API
func (r *CloudngsparameterResource) readCloudngsparameterFromApi(ctx context.Context, data *CloudngsparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudngsparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudngsparameter, got error: %s", err))
		return
	}

	cloudngsparameterSetAttrFromGet(ctx, data, getResponseData)

}