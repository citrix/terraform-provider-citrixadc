package interfacepair

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
var _ resource.Resource = &InterfacepairResource{}
var _ resource.ResourceWithConfigure = (*InterfacepairResource)(nil)
var _ resource.ResourceWithImportState = (*InterfacepairResource)(nil)

func NewInterfacepairResource() resource.Resource {
	return &InterfacepairResource{}
}

// InterfacepairResource defines the resource implementation.
type InterfacepairResource struct {
	client *service.NitroClient
}

func (r *InterfacepairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *InterfacepairResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfacepair"
}

func (r *InterfacepairResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InterfacepairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InterfacepairResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating interfacepair resource")

	interfacepair := interfacepairGetThePayloadFromthePlan(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource - use AddResource. The resource name is the numeric interface id.
	interfacepairName := fmt.Sprintf("%d", data.Interfaceid.ValueInt64())
	_, err := r.client.AddResource(service.Interfacepair.Type(), interfacepairName, &interfacepair)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create interfacepair, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created interfacepair resource")

	// Set ID for the resource before reading state (single unique attribute: interface_id)
	data.Id = types.StringValue(interfacepairName)

	// Read the updated state back
	if !r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "interfacepair not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacepairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InterfacepairResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading interfacepair resource")

	found := r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *InterfacepairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state InterfacepairResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating interfacepair resource")

	// interfacepair exposes no NITRO "update" verb and every configurable attribute
	// (interface_id, ifnum) is RequiresReplace, so any attribute change triggers a
	// destroy/recreate rather than reaching this path. This method therefore performs
	// no API write and only refreshes state.
	if !r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "interfacepair not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacepairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InterfacepairResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting interfacepair resource")

	// Named resource - delete using DeleteResource keyed on the numeric id.
	err := r.client.DeleteResource(service.Interfacepair.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete interfacepair, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted interfacepair resource")
}

// Helper function to read interfacepair data from API. Returns false when the
// resource no longer exists on the ADC so callers can remove it from state.
func (r *InterfacepairResource) readInterfacepairFromApi(ctx context.Context, data *InterfacepairResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain numeric value
	interfacepairName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Interfacepair.Type(), interfacepairName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read interfacepair, got error: %s", err))
		return false
	}

	interfacepairSetAttrFromGet(ctx, data, getResponseData)

	return true
}
