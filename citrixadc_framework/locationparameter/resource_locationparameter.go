package locationparameter

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
var _ resource.Resource = &LocationparameterResource{}
var _ resource.ResourceWithConfigure = (*LocationparameterResource)(nil)
var _ resource.ResourceWithImportState = (*LocationparameterResource)(nil)

func NewLocationparameterResource() resource.Resource {
	return &LocationparameterResource{}
}

// LocationparameterResource defines the resource implementation.
type LocationparameterResource struct {
	client *service.NitroClient
}

func (r *LocationparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LocationparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationparameter"
}

func (r *LocationparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationparameter resource")

	// Create API request body from the model
	locationparameter := locationparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Locationparameter.Type(), &locationparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create locationparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("locationparameter-config")

	tflog.Trace(ctx, "Created locationparameter resource")

	// Read the updated state back
	r.readLocationparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LocationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading locationparameter resource")

	found := r.readLocationparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LocationparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LocationparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating locationparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Context.Equal(state.Context) {
		tflog.Debug(ctx, "context has changed for locationparameter")
		hasChange = true
	}
	if !data.Matchwildcardtoany.Equal(state.Matchwildcardtoany) {
		tflog.Debug(ctx, "matchwildcardtoany has changed for locationparameter")
		hasChange = true
	}
	if !data.Q1label.Equal(state.Q1label) {
		tflog.Debug(ctx, "q1label has changed for locationparameter")
		hasChange = true
	}
	if !data.Q2label.Equal(state.Q2label) {
		tflog.Debug(ctx, "q2label has changed for locationparameter")
		hasChange = true
	}
	if !data.Q3label.Equal(state.Q3label) {
		tflog.Debug(ctx, "q3label has changed for locationparameter")
		hasChange = true
	}
	if !data.Q4label.Equal(state.Q4label) {
		tflog.Debug(ctx, "q4label has changed for locationparameter")
		hasChange = true
	}
	if !data.Q5label.Equal(state.Q5label) {
		tflog.Debug(ctx, "q5label has changed for locationparameter")
		hasChange = true
	}
	if !data.Q6label.Equal(state.Q6label) {
		tflog.Debug(ctx, "q6label has changed for locationparameter")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		locationparameter := locationparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Locationparameter.Type(), &locationparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update locationparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated locationparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for locationparameter resource, skipping update")
	}

	// Read the updated state back
	r.readLocationparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LocationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting locationparameter resource")

	// locationparameter is a singleton configuration resource with no DELETE
	// operation on the ADC (matches SDK v2 behavior). Just remove it from state.
	tflog.Trace(ctx, "Deleted locationparameter resource from state")
}

// Helper function to read locationparameter data from API
func (r *LocationparameterResource) readLocationparameterFromApi(ctx context.Context, data *LocationparameterResourceModel, diags *diag.Diagnostics) bool {

	// Case 1: Simple find without ID (singleton)
	getResponseData, err := r.client.FindResource(service.Locationparameter.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read locationparameter, got error: %s", err))
		return false
	}

	locationparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
