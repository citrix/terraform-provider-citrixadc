package systemnsbtracing

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
var _ resource.Resource = &SystemnsbtracingResource{}
var _ resource.ResourceWithConfigure = (*SystemnsbtracingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemnsbtracingResource)(nil)

func NewSystemnsbtracingResource() resource.Resource {
	return &SystemnsbtracingResource{}
}

// SystemnsbtracingResource defines the resource implementation.
type SystemnsbtracingResource struct {
	client *service.NitroClient
}

func (r *SystemnsbtracingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemnsbtracingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemnsbtracing"
}

func (r *SystemnsbtracingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// systemnsbtracing is an enable/disable toggle singleton. NITRO exposes only
// ?action=enable and ?action=disable (both with an EMPTY payload) plus get(all);
// there is no add/set/unset/delete verb and no writable desired-state attribute in
// the model. The toggle is therefore driven by the Terraform lifecycle:
//   Create -> enable, Delete -> disable, Update -> no-op (nodeid RequiresReplace).

func (r *SystemnsbtracingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemnsbtracingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (enabling) systemnsbtracing resource")

	// enable action takes an empty payload {"systemnsbtracing":{}}
	payload := map[string]interface{}{}
	err := r.client.ActOnResource(service.Systemnsbtracing.Type(), payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable systemnsbtracing, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled systemnsbtracing resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("systemnsbtracing-config")

	// Read the updated state back
	r.readSystemnsbtracingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemnsbtracingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemnsbtracingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemnsbtracing resource")

	r.readSystemnsbtracingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemnsbtracingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemnsbtracingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for systemnsbtracing: the only configurable attribute
	// (nodeid) is RequiresReplace, and the enable/disable verbs are driven by the
	// Create/Delete lifecycle. There is no NITRO set/update endpoint.
	tflog.Debug(ctx, "Update is a no-op for systemnsbtracing; enable/disable is driven by Create/Delete")

	r.readSystemnsbtracingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemnsbtracingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemnsbtracingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting (disabling) systemnsbtracing resource")

	// disable action takes an empty payload {"systemnsbtracing":{}}
	payload := map[string]interface{}{}
	err := r.client.ActOnResource(service.Systemnsbtracing.Type(), payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable systemnsbtracing, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Disabled systemnsbtracing resource")
}

// Helper function to read systemnsbtracing data from API
func (r *SystemnsbtracingResource) readSystemnsbtracingFromApi(ctx context.Context, data *SystemnsbtracingResourceModel, diags *diag.Diagnostics) {

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Systemnsbtracing.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemnsbtracing, got error: %s", err))
		return
	}

	systemnsbtracingSetAttrFromGet(ctx, data, getResponseData)

}
