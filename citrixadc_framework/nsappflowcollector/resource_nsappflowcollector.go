package nsappflowcollector

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
var _ resource.Resource = &NsappflowcollectorResource{}
var _ resource.ResourceWithConfigure = (*NsappflowcollectorResource)(nil)
var _ resource.ResourceWithImportState = (*NsappflowcollectorResource)(nil)

func NewNsappflowcollectorResource() resource.Resource {
	return &NsappflowcollectorResource{}
}

// NsappflowcollectorResource defines the resource implementation.
type NsappflowcollectorResource struct {
	client *service.NitroClient
}

func (r *NsappflowcollectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsappflowcollectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsappflowcollector"
}

func (r *NsappflowcollectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsappflowcollectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsappflowcollector resource")

	nsappflowcollector := nsappflowcollectorGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	nsappflowcollectorName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsappflowcollector.Type(), nsappflowcollectorName, &nsappflowcollector)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsappflowcollector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsappflowcollector resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(fmt.Sprintf("%v", nsappflowcollectorName))

	// Read the updated state back
	if !r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsappflowcollector not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowcollectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsappflowcollector resource")

	found := r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsappflowcollectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsappflowcollectorResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// nsappflowcollector has no NITRO-updatable attributes (all attributes are
	// ForceNew in SDK v2, so changes force replacement rather than an update).
	// This Update simply reads current state back.
	tflog.Debug(ctx, "Updating nsappflowcollector resource (no updatable attributes)")

	// Read the updated state back
	if !r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsappflowcollector not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowcollectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsappflowcollector resource")

	// Named resource - delete using DeleteResource
	nsappflowcollectorName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsappflowcollector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsappflowcollector resource")
}

// Helper function to read nsappflowcollector data from API.
// Returns false (without an error diagnostic) when the resource no longer exists.
func (r *NsappflowcollectorResource) readNsappflowcollectorFromApi(ctx context.Context, data *NsappflowcollectorResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nsappflowcollectorName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsappflowcollector, got error: %s", err))
		return false
	}

	nsappflowcollectorSetAttrFromGet(ctx, data, getResponseData)

	return true
}
