package cloudtunnelparameter

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
var _ resource.Resource = &CloudtunnelparameterResource{}
var _ resource.ResourceWithConfigure = (*CloudtunnelparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CloudtunnelparameterResource)(nil)

func NewCloudtunnelparameterResource() resource.Resource {
	return &CloudtunnelparameterResource{}
}

// CloudtunnelparameterResource defines the resource implementation.
type CloudtunnelparameterResource struct {
	client *service.NitroClient
}

func (r *CloudtunnelparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudtunnelparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudtunnelparameter"
}

func (r *CloudtunnelparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudtunnelparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudtunnelparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudtunnelparameter resource")
	cloudtunnelparameter := cloudtunnelparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudtunnelparameter.Type(), &cloudtunnelparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudtunnelparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudtunnelparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("cloudtunnelparameter-config")

	// Read the updated state back
	r.readCloudtunnelparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudtunnelparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudtunnelparameter resource")

	r.readCloudtunnelparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudtunnelparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudtunnelparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Controllerfqdn.Equal(state.Controllerfqdn) {
		tflog.Debug(ctx, fmt.Sprintf("controllerfqdn has changed for cloudtunnelparameter"))
		hasChange = true
	}
	if !data.Fqdn.Equal(state.Fqdn) {
		tflog.Debug(ctx, fmt.Sprintf("fqdn has changed for cloudtunnelparameter"))
		hasChange = true
	}
	if !data.Resourcelocation.Equal(state.Resourcelocation) {
		tflog.Debug(ctx, fmt.Sprintf("resourcelocation has changed for cloudtunnelparameter"))
		hasChange = true
	}
	if !data.Subnetresourcelocationmappings.Equal(state.Subnetresourcelocationmappings) {
		tflog.Debug(ctx, fmt.Sprintf("subnetresourcelocationmappings has changed for cloudtunnelparameter"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudtunnelparameter := cloudtunnelparameterGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudtunnelparameter.Type(), &cloudtunnelparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudtunnelparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudtunnelparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudtunnelparameter resource, skipping update")
	}

	// Read the updated state back
	r.readCloudtunnelparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudtunnelparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudtunnelparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed cloudtunnelparameter from Terraform state")
}

// Helper function to read cloudtunnelparameter data from API
func (r *CloudtunnelparameterResource) readCloudtunnelparameterFromApi(ctx context.Context, data *CloudtunnelparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudtunnelparameter.Type(), "")
	if err != nil {
		// cloudtunnelparameter is feature-gated: on platforms/releases where the
		// feature is not enabled NITRO returns "Feature not supported in this release"
		// (and on some platforms "Operation not supported on this platform"). Treat
		// those as a non-fatal read so create/apply is not broken; preserve the
		// existing plan/state values and just (re)affirm the static ID.
		if isCloudtunnelparameterNotSupported(err) {
			tflog.Warn(ctx, "cloudtunnelparameter GET not supported on this platform/release; preserving state")
			data.Id = types.StringValue("cloudtunnelparameter-config")
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudtunnelparameter, got error: %s", err))
		return
	}

	cloudtunnelparameterSetAttrFromGet(ctx, data, getResponseData)

}

// isCloudtunnelparameterNotSupported reports whether the NITRO error indicates the
// cloudtunnelparameter feature is gated (not supported on this platform/release).
func isCloudtunnelparameterNotSupported(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "not supported on this platform") ||
		strings.Contains(msg, "Feature not supported")
}
