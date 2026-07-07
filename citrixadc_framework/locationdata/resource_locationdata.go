package locationdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// locationdata is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the clear action:
//     POST /nitro/v1/config/locationdata?action=clear, which clears the static
//     location (GSLB geo) database from memory.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the clear action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for locationdata.
var _ resource.Resource = &LocationdataResource{}
var _ resource.ResourceWithConfigure = (*LocationdataResource)(nil)
var _ resource.ResourceWithImportState = (*LocationdataResource)(nil)

func NewLocationdataResource() resource.Resource {
	return &LocationdataResource{}
}

// LocationdataResource defines the resource implementation.
type LocationdataResource struct {
	client *service.NitroClient
}

func (r *LocationdataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LocationdataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationdata"
}

func (r *LocationdataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationdataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationdataResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationdata resource (clear action)")
	locationdata := locationdataGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=clear
	err := r.client.ActOnResource(service.Locationdata.Type(), &locationdata, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear locationdata, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared locationdata")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("locationdata-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. locationdata has no GET endpoint; there is nothing to reconcile.
func (r *LocationdataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LocationdataResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for locationdata; NITRO exposes no GET endpoint (action=clear only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. locationdata has no attributes and no set endpoint.
func (r *LocationdataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LocationdataResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for locationdata; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. locationdata has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *LocationdataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for locationdata; NITRO has no delete endpoint")
}
