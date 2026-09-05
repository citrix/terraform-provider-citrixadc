package mapbmr

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
var _ resource.Resource = &MapbmrResource{}
var _ resource.ResourceWithConfigure = (*MapbmrResource)(nil)
var _ resource.ResourceWithImportState = (*MapbmrResource)(nil)

func NewMapbmrResource() resource.Resource {
	return &MapbmrResource{}
}

// MapbmrResource defines the resource implementation.
type MapbmrResource struct {
	client *service.NitroClient
}

func (r *MapbmrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapbmrResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapbmr"
}

func (r *MapbmrResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapbmrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapbmrResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapbmr resource")

	// Create API request body from the model
	mapbmr := mapbmrGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Mapbmr.Type(), name_value, &mapbmr)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapbmr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created mapbmr resource")

	// Set ID for the resource before reading state (single unique attr: name)
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readMapbmrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapbmr not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapbmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapbmr resource")

	found := r.readMapbmrFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *MapbmrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MapbmrResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating mapbmr resource")

	// mapbmr has no updateable attributes (SDK v2: every attribute is ForceNew,
	// no UpdateContext). Any change forces recreation via RequiresReplace, so this
	// Update path is effectively never exercised - just re-read current state.
	tflog.Trace(ctx, "mapbmr has no updateable attributes; re-reading state")

	// Read the updated state back
	if !r.readMapbmrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapbmr not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapbmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapbmr resource")
	// Named resource - delete using DeleteResource by ID (the live name)
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Mapbmr.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete mapbmr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted mapbmr resource")
}

// Helper function to read mapbmr data from API. Returns false if the resource
// no longer exists on the ADC (so callers can remove it from state).
func (r *MapbmrResource) readMapbmrFromApi(ctx context.Context, data *MapbmrResourceModel, diags *diag.Diagnostics) bool {
	// Named resource: find by ID (the plain name value)
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Mapbmr.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapbmr, got error: %s", err))
		return false
	}

	mapbmrSetAttrFromGet(ctx, data, getResponseData)

	return true
}
