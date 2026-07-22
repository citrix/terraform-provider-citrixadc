package traceroute

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// traceroute is an ACTION-ONLY diagnostic resource.
//
//   - NITRO exposes only the Traceroute action: POST /nitro/v1/config/traceroute
//     (no ?action= query param, empty verb), which runs a traceroute from the ADC.
//   - There is NO add/set/get/delete endpoint, so Create performs the traceroute,
//     Read/Update are no-ops (nothing to reconcile), and Delete is a state-only
//     removal.
//   - Because there is no GET endpoint, there is NO datasource for traceroute and
//     the resource cannot be verified by reading it back.
var _ resource.Resource = &TracerouteResource{}
var _ resource.ResourceWithConfigure = (*TracerouteResource)(nil)

func NewTracerouteResource() resource.Resource {
	return &TracerouteResource{}
}

// TracerouteResource defines the resource implementation.
type TracerouteResource struct {
	client *service.NitroClient
}

func (r *TracerouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traceroute"
}

func (r *TracerouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TracerouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TracerouteResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating traceroute resource (traceroute action)")
	traceroute := tracerouteGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the Traceroute action
	// (bare POST /config/traceroute, empty action verb).
	err := r.client.ActOnResource(tracerouteResourceType, &traceroute, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to run traceroute, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Ran traceroute")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("traceroute-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. traceroute has no GET endpoint; there is nothing to reconcile.
func (r *TracerouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TracerouteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for traceroute; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. Every traceroute attribute is RequiresReplace and NITRO
// exposes no set endpoint (action only).
func (r *TracerouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TracerouteResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for traceroute")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. traceroute has no delete endpoint; just remove from state.
func (r *TracerouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for traceroute; removed from Terraform state")
}
