package iptunnel

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
var _ resource.Resource = &IptunnelResource{}
var _ resource.ResourceWithConfigure = (*IptunnelResource)(nil)
var _ resource.ResourceWithImportState = (*IptunnelResource)(nil)

func NewIptunnelResource() resource.Resource {
	return &IptunnelResource{}
}

// IptunnelResource defines the resource implementation.
type IptunnelResource struct {
	client *service.NitroClient
}

func (r *IptunnelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IptunnelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iptunnel"
}

func (r *IptunnelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IptunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IptunnelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating iptunnel resource")

	iptunnel := iptunnelGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	iptunnelName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Iptunnel.Type(), iptunnelName, &iptunnel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create iptunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created iptunnel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(iptunnelName)

	// Read the updated state back
	if !r.readIptunnelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "iptunnel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IptunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IptunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading iptunnel resource")

	found := r.readIptunnelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IptunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state IptunnelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating iptunnel resource")

	// Only destport, tosinherit and vlantagging are updatable in place (SDK v2 parity).
	hasChange := false
	if !data.Destport.Equal(state.Destport) {
		tflog.Debug(ctx, "destport has changed for iptunnel")
		hasChange = true
	}
	if !data.Tosinherit.Equal(state.Tosinherit) {
		tflog.Debug(ctx, "tosinherit has changed for iptunnel")
		hasChange = true
	}
	if !data.Vlantagging.Equal(state.Vlantagging) {
		tflog.Debug(ctx, "vlantagging has changed for iptunnel")
		hasChange = true
	}

	if hasChange {
		iptunnel := iptunnelGetTheUpdatablePayloadFromThePlan(ctx, &data)
		iptunnelName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Iptunnel.Type(), iptunnelName, &iptunnel)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update iptunnel, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated iptunnel resource")
	} else {
		tflog.Debug(ctx, "No changes detected for iptunnel resource, skipping update")
	}

	// Read the updated state back
	if !r.readIptunnelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "iptunnel not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IptunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IptunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting iptunnel resource")

	// Named resource - delete using DeleteResource
	iptunnelName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete iptunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted iptunnel resource")
}

// Helper function to read iptunnel data from API
func (r *IptunnelResource) readIptunnelFromApi(ctx context.Context, data *IptunnelResourceModel, diags *diag.Diagnostics) bool {
	// Single ID attribute (name) - ID is the plain value
	iptunnelName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read iptunnel, got error: %s", err))
		return false
	}

	iptunnelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
