package locationfile

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
var _ resource.Resource = &LocationfileResource{}
var _ resource.ResourceWithConfigure = (*LocationfileResource)(nil)
var _ resource.ResourceWithImportState = (*LocationfileResource)(nil)

func NewLocationfileResource() resource.Resource {
	return &LocationfileResource{}
}

// LocationfileResource defines the resource implementation.
type LocationfileResource struct {
	client *service.NitroClient
}

func (r *LocationfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LocationfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationfile"
}

func (r *LocationfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationfile resource")

	// Build the `add` payload (Locationfile + format; src is not part of add).
	locationfile := locationfileGetThePayloadFromthePlan(ctx, &data)

	// Named-in-body resource: SDK v2 used AddResource with an empty name (the
	// location file name travels in the payload). Mirror that exactly.
	_, err := r.client.AddResource(service.Locationfile.Type(), "", &locationfile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create locationfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created locationfile resource")

	// SDK v2 ID scheme: d.SetId(locationfile name). Preserve it for backward compat.
	data.Id = types.StringValue(data.Locationfile.ValueString())

	// Read the updated state back
	if !r.readLocationfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "locationfile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LocationfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading locationfile resource")

	found := r.readLocationfileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LocationfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// The NITRO locationfile resource exposes no `update` operation and every
	// writable attribute is RequiresReplace / RequiresReplaceIfConfigured (matching
	// SDK v2, which had no Update and marked all attributes ForceNew). Terraform
	// therefore never invokes Update for a real change; this is a defensive
	// state-preserving read-back.
	var data, state LocationfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating locationfile resource (no NITRO update op; read-back only)")

	if !r.readLocationfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "locationfile not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LocationfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting locationfile resource")

	// SDK v2: DeleteResource(type, "") — DELETE /config/locationfile (no name).
	err := r.client.DeleteResource(service.Locationfile.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete locationfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted locationfile resource")
}

// readLocationfileFromApi reads locationfile data from the API into data. It
// returns false (without adding an error) when the resource no longer exists so
// callers can drop it from state.
func (r *LocationfileResource) readLocationfileFromApi(ctx context.Context, data *LocationfileResourceModel, diags *diag.Diagnostics) bool {
	// locationfile GET is a "get (all)"; SDK v2 read used an empty name.
	getResponseData, err := r.client.FindResource(service.Locationfile.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read locationfile, got error: %s", err))
		return false
	}
	if getResponseData == nil {
		return false
	}

	locationfileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
