package protocolhttpband

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProtocolhttpbandResource{}
var _ resource.ResourceWithConfigure = (*ProtocolhttpbandResource)(nil)

func NewProtocolhttpbandResource() resource.Resource {
	return &ProtocolhttpbandResource{}
}

// ProtocolhttpbandResource defines the resource implementation.
type ProtocolhttpbandResource struct {
	client *service.NitroClient
}

func (r *ProtocolhttpbandResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocolhttpband"
}

func (r *ProtocolhttpbandResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ProtocolhttpbandResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProtocolhttpbandResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating protocolhttpband resource")
	protocolhttpband := protocolhttpbandGetThePayloadFromthePlan(ctx, &data)

	// Singleton settings resource - create is an update/set (NITRO has no add verb).
	err := r.client.UpdateUnnamedResource(service.Protocolhttpband.Type(), &protocolhttpband)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create protocolhttpband, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created protocolhttpband resource")

	// Fixed synthetic ID for the singleton settings resource.
	data.Id = types.StringValue("protocolhttpband-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProtocolhttpbandResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for protocolhttpband: NITRO exposes only a stats-only show
	// (keyed by a mandatory `type` filter) and no readback of the configured
	// reqbandsize/respbandsize object state. Preserve the values from state.
	tflog.Debug(ctx, "Read is a no-op for protocolhttpband; no config readback endpoint on NITRO side, preserving state")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ProtocolhttpbandResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating protocolhttpband resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqbandsize.Equal(state.Reqbandsize) {
		tflog.Debug(ctx, "reqbandsize has changed for protocolhttpband")
		hasChange = true
	}
	if !data.Respbandsize.Equal(state.Respbandsize) {
		tflog.Debug(ctx, "respbandsize has changed for protocolhttpband")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		protocolhttpband := protocolhttpbandGetThePayloadFromthePlan(ctx, &data)
		// Singleton settings resource - update is an update/set.
		err := r.client.UpdateUnnamedResource(service.Protocolhttpband.Type(), &protocolhttpband)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update protocolhttpband, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated protocolhttpband resource")
	} else {
		tflog.Debug(ctx, "No changes detected for protocolhttpband resource, skipping update")
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProtocolhttpbandResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Singleton settings resource - NITRO exposes no rm/delete verb, just remove from state.
	tflog.Trace(ctx, "Removed protocolhttpband from Terraform state (no delete endpoint on NITRO side)")
}
