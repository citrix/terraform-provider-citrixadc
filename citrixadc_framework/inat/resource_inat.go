package inat

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
var _ resource.Resource = &InatResource{}
var _ resource.ResourceWithConfigure = (*InatResource)(nil)
var _ resource.ResourceWithImportState = (*InatResource)(nil)

func NewInatResource() resource.Resource {
	return &InatResource{}
}

// InatResource defines the resource implementation.
type InatResource struct {
	client *service.NitroClient
}

func (r *InatResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *InatResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inat"
}

func (r *InatResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InatResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InatResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating inat resource")

	inat := inatGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource keyed on name
	inatName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Inat.Type(), inatName, &inat)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create inat, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created inat resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(inatName)

	// Read the updated state back
	if !r.readInatFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "inat not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InatResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InatResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading inat resource")

	found := r.readInatFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *InatResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state InatResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating inat resource")

	// Check if there are any changes in updateable attributes
	// (publicip, td and name are RequiresReplace and never reach Update)
	hasChange := false
	if !data.Connfailover.Equal(state.Connfailover) {
		tflog.Debug(ctx, "connfailover has changed for inat")
		hasChange = true
	}
	if !data.Ftp.Equal(state.Ftp) {
		tflog.Debug(ctx, "ftp has changed for inat")
		hasChange = true
	}
	if !data.Mode.Equal(state.Mode) {
		tflog.Debug(ctx, "mode has changed for inat")
		hasChange = true
	}
	if !data.Privateip.Equal(state.Privateip) {
		tflog.Debug(ctx, "privateip has changed for inat")
		hasChange = true
	}
	if !data.Proxyip.Equal(state.Proxyip) {
		tflog.Debug(ctx, "proxyip has changed for inat")
		hasChange = true
	}
	if !data.Tcpproxy.Equal(state.Tcpproxy) {
		tflog.Debug(ctx, "tcpproxy has changed for inat")
		hasChange = true
	}
	if !data.Tftp.Equal(state.Tftp) {
		tflog.Debug(ctx, "tftp has changed for inat")
		hasChange = true
	}
	if !data.Useproxyport.Equal(state.Useproxyport) {
		tflog.Debug(ctx, "useproxyport has changed for inat")
		hasChange = true
	}
	if !data.Usip.Equal(state.Usip) {
		tflog.Debug(ctx, "usip has changed for inat")
		hasChange = true
	}
	if !data.Usnip.Equal(state.Usnip) {
		tflog.Debug(ctx, "usnip has changed for inat")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		inat := inatGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource keyed on name
		inatName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Inat.Type(), inatName, &inat)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update inat, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated inat resource")
	} else {
		tflog.Debug(ctx, "No changes detected for inat resource, skipping update")
	}

	// Read the updated state back
	if !r.readInatFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "inat not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InatResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InatResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting inat resource")

	// Named resource - delete using DeleteResource keyed on the live name (id)
	inatName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Inat.Type(), inatName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete inat, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted inat resource")
}

// Helper function to read inat data from API. Returns false when the resource no
// longer exists on the ADC (so callers can drop it from state).
func (r *InatResource) readInatFromApi(ctx context.Context, data *InatResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	inatName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Inat.Type(), inatName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read inat, got error: %s", err))
		return false
	}

	inatSetAttrFromGet(ctx, data, getResponseData)

	return true
}
