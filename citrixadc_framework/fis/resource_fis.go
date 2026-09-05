package fis

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
var _ resource.Resource = &FisResource{}
var _ resource.ResourceWithConfigure = (*FisResource)(nil)
var _ resource.ResourceWithImportState = (*FisResource)(nil)

func NewFisResource() resource.Resource {
	return &FisResource{}
}

// FisResource defines the resource implementation.
type FisResource struct {
	client *service.NitroClient
}

func (r *FisResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FisResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fis"
}

func (r *FisResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FisResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FisResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating fis resource")

	fis := fisGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	fisName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Fis.Type(), fisName, &fis)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create fis, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created fis resource")

	// Set ID for the resource before reading state.
	// Case 2: Single unique attribute - use plain value (name) as ID
	data.Id = types.StringValue(fisName)

	// Read the updated state back
	if !r.readFisFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "fis not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FisResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading fis resource")

	found := r.readFisFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *FisResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FisResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// fis exposes no NITRO "update" operation, and both name and ownernode are
	// ForceNew (RequiresReplace), so Terraform never reaches Update for a real
	// change. Read the current state back defensively.
	tflog.Debug(ctx, "Updating fis resource (no-op: all attributes require replacement)")

	if !r.readFisFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "fis not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FisResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting fis resource")

	// Named resource - delete using DeleteResource keyed on the name (ID)
	fisName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Fis.Type(), fisName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete fis, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted fis resource")
}

// Helper function to read fis data from API. Returns false when the resource no
// longer exists on the ADC (so callers can drop it from state).
func (r *FisResource) readFisFromApi(ctx context.Context, data *FisResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	fisName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Fis.Type(), fisName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read fis, got error: %s", err))
		return false
	}

	fisSetAttrFromGet(ctx, data, getResponseData)

	return true
}
