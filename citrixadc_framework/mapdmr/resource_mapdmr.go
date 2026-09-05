package mapdmr

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
var _ resource.Resource = &MapdmrResource{}
var _ resource.ResourceWithConfigure = (*MapdmrResource)(nil)
var _ resource.ResourceWithImportState = (*MapdmrResource)(nil)

func NewMapdmrResource() resource.Resource {
	return &MapdmrResource{}
}

// MapdmrResource defines the resource implementation.
type MapdmrResource struct {
	client *service.NitroClient
}

func (r *MapdmrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapdmrResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapdmr"
}

func (r *MapdmrResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapdmrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapdmrResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapdmr resource")

	// Build the payload from the plan
	mapdmr := mapdmrGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	mapdmrName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Mapdmr.Type(), mapdmrName, &mapdmr)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapdmr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created mapdmr resource")

	// Set ID for the resource before reading state (Case 2: single unique attribute)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapdmr not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdmrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapdmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapdmr resource")

	found := r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *MapdmrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MapdmrResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating mapdmr resource")

	// mapdmr has no NITRO-updatable attributes: both name and bripv6prefix are
	// ForceNew (RequiresReplace), so any change forces a destroy/create instead of
	// an in-place update. There is nothing to push here; just refresh state from the ADC.
	if !r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "mapdmr not found immediately after update")
		}
		return
	}

	tflog.Trace(ctx, "Updated mapdmr resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdmrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapdmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapdmr resource")

	// Named resource - delete using DeleteResource
	mapdmrName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Mapdmr.Type(), mapdmrName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete mapdmr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted mapdmr resource")
}

// Helper function to read mapdmr data from API
func (r *MapdmrResource) readMapdmrFromApi(ctx context.Context, data *MapdmrResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	mapdmrName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Mapdmr.Type(), mapdmrName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapdmr, got error: %s", err))
		return false
	}

	mapdmrSetAttrFromGet(ctx, data, getResponseData)

	return true
}
