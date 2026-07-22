package systemcpuparam

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
var _ resource.Resource = &SystemcpuparamResource{}
var _ resource.ResourceWithConfigure = (*SystemcpuparamResource)(nil)
var _ resource.ResourceWithImportState = (*SystemcpuparamResource)(nil)

func NewSystemcpuparamResource() resource.Resource {
	return &SystemcpuparamResource{}
}

// SystemcpuparamResource defines the resource implementation.
type SystemcpuparamResource struct {
	client *service.NitroClient
}

func (r *SystemcpuparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemcpuparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemcpuparam"
}

func (r *SystemcpuparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemcpuparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemcpuparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemcpuparam resource")
	systemcpuparam := systemcpuparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Parameter singleton - NITRO has no "add" verb; create is the "set" verb (PUT).
	_, err := r.client.UpdateResource(service.Systemcpuparam.Type(), "", &systemcpuparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemcpuparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemcpuparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("systemcpuparam-config")

	// Read the updated state back
	r.readSystemcpuparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcpuparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemcpuparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemcpuparam resource")

	r.readSystemcpuparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcpuparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemcpuparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemcpuparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Pemode.Equal(state.Pemode) {
		tflog.Debug(ctx, fmt.Sprintf("pemode has changed for systemcpuparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		systemcpuparam := systemcpuparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Parameter singleton - update is the "set" verb (PUT).
		_, err := r.client.UpdateResource(service.Systemcpuparam.Type(), "", &systemcpuparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemcpuparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemcpuparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemcpuparam resource, skipping update")
	}

	// Read the updated state back
	r.readSystemcpuparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcpuparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemcpuparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemcpuparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed systemcpuparam from Terraform state")
}

// Helper function to read systemcpuparam data from API
func (r *SystemcpuparamResource) readSystemcpuparamFromApi(ctx context.Context, data *SystemcpuparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Systemcpuparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemcpuparam, got error: %s", err))
		return
	}

	systemcpuparamSetAttrFromGet(ctx, data, getResponseData)

}
