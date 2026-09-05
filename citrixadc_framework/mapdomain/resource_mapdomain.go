package mapdomain

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
var _ resource.Resource = &MapdomainResource{}
var _ resource.ResourceWithConfigure = (*MapdomainResource)(nil)
var _ resource.ResourceWithImportState = (*MapdomainResource)(nil)

func NewMapdomainResource() resource.Resource {
	return &MapdomainResource{}
}

// MapdomainResource defines the resource implementation.
type MapdomainResource struct {
	client *service.NitroClient
}

func (r *MapdomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapdomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapdomain"
}

func (r *MapdomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapdomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapdomainResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapdomain resource")

	// Create API request body from the model
	mapdomain := mapdomainGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add is POST /config/mapdomain/{name})
	mapdomainName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Mapdomain.Type(), mapdomainName, &mapdomain)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapdomain, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created mapdomain resource")

	// Set ID for the resource before reading state (single unique attribute: name)
	data.Id = types.StringValue(mapdomainName)

	// Read the updated state back
	if !r.readMapdomainFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapdomain not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapdomain resource")

	found := r.readMapdomainFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *MapdomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MapdomainResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// mapdomain has no updateable attributes: both `name` and `mapdmrname` are
	// ForceNew (RequiresReplace) in the SDK v2 contract, so Terraform never
	// reaches Update for a value change (it destroys+recreates instead). There
	// is no NITRO update endpoint to call here (Pattern 5). Just read the live
	// state back to keep it consistent.
	tflog.Debug(ctx, "Updating mapdomain resource (no updateable attributes; refreshing state)")

	if !r.readMapdomainFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapdomain not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapdomain resource")

	// Named resource - delete using DeleteResource
	mapdomainName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Mapdomain.Type(), mapdomainName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete mapdomain, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted mapdomain resource")
}

// Helper function to read mapdomain data from API
func (r *MapdomainResource) readMapdomainFromApi(ctx context.Context, data *MapdomainResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	mapdomainName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Mapdomain.Type(), mapdomainName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapdomain, got error: %s", err))
		return false
	}

	mapdomainSetAttrFromGet(ctx, data, getResponseData)

	return true
}
