package vxlanvlanmap

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
var _ resource.Resource = &VxlanvlanmapResource{}
var _ resource.ResourceWithConfigure = (*VxlanvlanmapResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanvlanmapResource)(nil)

func NewVxlanvlanmapResource() resource.Resource {
	return &VxlanvlanmapResource{}
}

// VxlanvlanmapResource defines the resource implementation.
type VxlanvlanmapResource struct {
	client *service.NitroClient
}

func (r *VxlanvlanmapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanvlanmapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlanvlanmap"
}

func (r *VxlanvlanmapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanvlanmapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanvlanmapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlanvlanmap resource")

	vxlanvlanmap := vxlanvlanmapGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	vxlanvlanmapName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vxlanvlanmap.Type(), vxlanvlanmapName, &vxlanvlanmap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlanvlanmap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vxlanvlanmap resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vxlanvlanmapName)

	// Read the updated state back
	if !r.readVxlanvlanmapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vxlanvlanmap not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanvlanmapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlanvlanmap resource")

	found := r.readVxlanvlanmapFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VxlanvlanmapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VxlanvlanmapResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vxlanvlanmap resource")

	// vxlanvlanmap has a single ForceNew attribute (name); any name change triggers
	// RequiresReplace (destroy/create), so Update is never invoked with a changed
	// attribute. There is nothing to push to NITRO here; just refresh state.

	if !r.readVxlanvlanmapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vxlanvlanmap not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanvlanmapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlanvlanmap resource")

	// Named resource - delete using DeleteResource
	vxlanvlanmapName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vxlanvlanmap.Type(), vxlanvlanmapName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vxlanvlanmap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vxlanvlanmap resource")
}

// Helper function to read vxlanvlanmap data from API. Returns false if the
// resource no longer exists on the ADC (so the caller can drop it from state).
func (r *VxlanvlanmapResource) readVxlanvlanmapFromApi(ctx context.Context, data *VxlanvlanmapResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vxlanvlanmapName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vxlanvlanmap.Type(), vxlanvlanmapName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlanvlanmap, got error: %s", err))
		return false
	}

	vxlanvlanmapSetAttrFromGet(ctx, data, getResponseData)

	return true
}
