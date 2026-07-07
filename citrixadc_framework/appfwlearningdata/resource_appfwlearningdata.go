package appfwlearningdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// appfwlearningdata is a BEST-EFFORT model of the Application-Firewall learned-data
// table, which is NOT a normal CRUD object.
//
//   - NITRO exposes get(all), count, delete, and the reset/export actions only;
//     there is NO add/set endpoint, and no per-object identity to reconcile.
//   - This resource is therefore modeled as an ACTION resource: Create performs
//     the "reset" action (POST /nitro/v1/config/appfwlearningdata?action=reset),
//     which clears the learned data for the given profile/security check.
//   - Read and Update are no-ops (there is nothing to reconcile against a table of
//     learned entries). Delete is a state-only removal: the NITRO delete endpoint
//     (DELETE /appfwlearningdata) carries no ?args= key selector in the metadata,
//     so it does not map cleanly to DeleteResourceWithArgs; performing a bare
//     delete on Terraform destroy would be surprising, so Delete only drops the
//     resource from state.
//   - WARNING: applying this resource RESETS (clears) App-Firewall learned data.
//     The reset/delete semantics here are BEST-EFFORT and should be verified on a
//     live ADC before relying on them.
//   - The GET(all) side is exposed via the companion datasource.
var _ resource.Resource = &AppfwlearningdataResource{}
var _ resource.ResourceWithConfigure = (*AppfwlearningdataResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwlearningdataResource)(nil)

func NewAppfwlearningdataResource() resource.Resource {
	return &AppfwlearningdataResource{}
}

// AppfwlearningdataResource defines the resource implementation.
type AppfwlearningdataResource struct {
	client *service.NitroClient
}

func (r *AppfwlearningdataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwlearningdataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningdata"
}

func (r *AppfwlearningdataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwlearningdataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwlearningdataResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Resetting appfwlearningdata (reset action)")
	appfwlearningdata := appfwlearningdataGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: perform the "reset" action.
	// Verb "reset" comes from operation_details["reset"].action.
	err := r.client.ActOnResource(service.Appfwlearningdata.Type(), &appfwlearningdata, "reset")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reset appfwlearningdata, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered appfwlearningdata reset")

	// Synthetic ID - there is no per-object GET to read back an identity.
	data.Id = types.StringValue("appfwlearningdata-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. appfwlearningdata is a table of learned entries with no per-object
// identity to reconcile; the GET(all) view is exposed via the datasource instead.
func (r *AppfwlearningdataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwlearningdataResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for appfwlearningdata (action resource); preserving prior state")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes are RequiresReplace, so there is no in-place
// update to exercise; a changed input triggers a replace (a fresh reset action).
func (r *AppfwlearningdataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwlearningdataResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for appfwlearningdata; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a state-only removal. The NITRO delete endpoint for appfwlearningdata
// carries no ?args= key selector in the metadata, so it does not map cleanly to
// DeleteResourceWithArgs, and issuing a bare delete on destroy would be surprising.
// We therefore only drop the resource from Terraform state. If you need to clear
// learned data, apply the resource (reset action) rather than destroying it.
func (r *AppfwlearningdataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a state-only removal for appfwlearningdata (best-effort action model); no NITRO delete is issued")
}
