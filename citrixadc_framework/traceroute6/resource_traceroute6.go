package traceroute6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// traceroute6 is an ACTION-ONLY diagnostic resource.
//
//   - NITRO exposes only the Traceroute6 action: POST
//     /nitro/v1/config/traceroute6 (no ?action= query param, empty verb), which
//     runs an IPv6 traceroute from the ADC.
//   - There is NO add/set/get/delete endpoint, so Create performs the
//     traceroute6, Read/Update are no-ops (nothing to reconcile), and Delete is a
//     state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for traceroute6
//     and the resource cannot be verified by reading it back.
var _ resource.Resource = &Traceroute6Resource{}
var _ resource.ResourceWithConfigure = (*Traceroute6Resource)(nil)

func NewTraceroute6Resource() resource.Resource {
	return &Traceroute6Resource{}
}

// Traceroute6Resource defines the resource implementation.
type Traceroute6Resource struct {
	client *service.NitroClient
}

func (r *Traceroute6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traceroute6"
}

func (r *Traceroute6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Traceroute6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Traceroute6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating traceroute6 resource (traceroute6 action)")
	traceroute6 := traceroute6GetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the Traceroute6 action
	// (bare POST /config/traceroute6, empty action verb).
	err := r.client.ActOnResource(traceroute6ResourceType, &traceroute6, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to run traceroute6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Ran traceroute6")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("traceroute6-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. traceroute6 has no GET endpoint; there is nothing to reconcile.
func (r *Traceroute6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Traceroute6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for traceroute6; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. Every traceroute6 attribute is RequiresReplace and NITRO
// exposes no set endpoint (action only).
func (r *Traceroute6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Traceroute6ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for traceroute6")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. traceroute6 has no delete endpoint; just remove from state.
func (r *Traceroute6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for traceroute6; removed from Terraform state")
}
