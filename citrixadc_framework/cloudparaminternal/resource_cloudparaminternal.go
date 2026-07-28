package cloudparaminternal

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudparaminternalResource{}
var _ resource.ResourceWithConfigure = (*CloudparaminternalResource)(nil)
var _ resource.ResourceWithImportState = (*CloudparaminternalResource)(nil)

func NewCloudparaminternalResource() resource.Resource {
	return &CloudparaminternalResource{}
}

// CloudparaminternalResource defines the resource implementation.
type CloudparaminternalResource struct {
	client *service.NitroClient
}

func (r *CloudparaminternalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudparaminternalResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudparaminternal"
}

func (r *CloudparaminternalResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudparaminternalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudparaminternalResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudparaminternal resource")
	cloudparaminternal := cloudparaminternalGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudparaminternal.Type(), &cloudparaminternal)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudparaminternal, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudparaminternal resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("cloudparaminternal-config")

	// Read the updated state back
	r.readCloudparaminternalFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparaminternalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudparaminternalResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudparaminternal resource")

	r.readCloudparaminternalFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparaminternalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudparaminternalResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudparaminternal resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Nonftumode.Equal(state.Nonftumode) {
		tflog.Debug(ctx, fmt.Sprintf("nonftumode has changed for cloudparaminternal"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudparaminternal := cloudparaminternalGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudparaminternal.Type(), &cloudparaminternal)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudparaminternal, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudparaminternal resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudparaminternal resource, skipping update")
	}

	// Read the updated state back
	r.readCloudparaminternalFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudparaminternalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudparaminternalResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudparaminternal resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed cloudparaminternal from Terraform state")
}

// Helper function to read cloudparaminternal data from API
func (r *CloudparaminternalResource) readCloudparaminternalFromApi(ctx context.Context, data *CloudparaminternalResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudparaminternal.Type(), "")
	if err != nil {
		// cloudparaminternal GET/show is platform-gated: on unsupported platforms
		// NITRO returns "Operation not supported on this platform". Treat that as a
		// non-fatal read so create/apply is not broken; preserve the existing
		// plan/state value and just (re)affirm the static ID.
		if strings.Contains(err.Error(), "not supported on this platform") {
			tflog.Warn(ctx, "cloudparaminternal GET not supported on this platform; preserving state")
			data.Id = types.StringValue("cloudparaminternal-config")
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudparaminternal, got error: %s", err))
		return
	}

	cloudparaminternalSetAttrFromGet(ctx, data, getResponseData)

}
