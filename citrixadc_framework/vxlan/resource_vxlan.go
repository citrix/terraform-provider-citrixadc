package vxlan

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
var _ resource.Resource = &VxlanResource{}
var _ resource.ResourceWithConfigure = (*VxlanResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanResource)(nil)

func NewVxlanResource() resource.Resource {
	return &VxlanResource{}
}

// VxlanResource defines the resource implementation.
type VxlanResource struct {
	client *service.NitroClient
}

func (r *VxlanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan"
}

func (r *VxlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan resource")

	vxlan := vxlanGetThePayloadFromthePlan(ctx, &data)

	// Named resource keyed on vxlanid - use AddResource (SDK v2 parity)
	vxlanIdStr := fmt.Sprintf("%d", data.Vxlanid.ValueInt64())
	_, err := r.client.AddResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vxlan resource")

	// Set ID (SDK v2: d.SetId(vxlanIdStr)) before reading state back
	data.Id = types.StringValue(vxlanIdStr)

	// Read the updated state back
	if !r.readVxlanFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vxlan not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan resource")

	found := r.readVxlanFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VxlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VxlanResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vxlan resource")

	vxlan, hasChange := vxlanGetTheUpdatablePayloadFromThePlan(ctx, &data, &state)
	if hasChange {
		// Named resource - use UpdateResource keyed on the live ID
		vxlanIdStr := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlan, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vxlan resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vxlan resource, skipping update")
	}

	// Read the updated state back
	if !r.readVxlanFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vxlan not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan resource")

	// Named resource - delete using DeleteResource keyed on the live ID
	err := r.client.DeleteResource(service.Vxlan.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vxlan, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vxlan resource")
}

// Helper function to read vxlan data from API. Returns false when the resource
// no longer exists on the ADC.
func (r *VxlanResource) readVxlanFromApi(ctx context.Context, data *VxlanResourceModel, diags *diag.Diagnostics) bool {
	vxlanIdStr := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vxlan.Type(), vxlanIdStr)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan, got error: %s", err))
		return false
	}

	vxlanSetAttrFromGet(ctx, data, getResponseData)

	return true
}
