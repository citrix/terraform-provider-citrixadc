package ping6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ping6 is an ACTION-ONLY diagnostic resource.
//
//   - NITRO exposes only the Ping6 action: POST /nitro/v1/config/ping6 (no
//     ?action= query param, empty verb), which runs an IPv6 ping from the ADC.
//   - There is NO add/set/get/delete endpoint, so Create performs the ping,
//     Read/Update are no-ops (nothing to reconcile), and Delete is a state-only
//     removal.
//   - Because there is no GET endpoint, there is NO datasource for ping6 and the
//     resource cannot be verified by reading it back.
var _ resource.Resource = &Ping6Resource{}
var _ resource.ResourceWithConfigure = (*Ping6Resource)(nil)

func NewPing6Resource() resource.Resource {
	return &Ping6Resource{}
}

// Ping6Resource defines the resource implementation.
type Ping6Resource struct {
	client *service.NitroClient
}

func (r *Ping6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ping6"
}

func (r *Ping6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Ping6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ping6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ping6 resource (ping6 action)")
	ping6 := ping6GetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the Ping6 action
	// (bare POST /config/ping6, empty action verb).
	err := r.client.ActOnResource(ping6ResourceType, &ping6, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to run ping6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Ran ping6")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("ping6-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. ping6 has no GET endpoint; there is nothing to reconcile.
func (r *Ping6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ping6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for ping6; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. Every ping6 attribute is RequiresReplace and NITRO exposes
// no set endpoint (action only).
func (r *Ping6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Ping6ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for ping6")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. ping6 has no delete endpoint; just remove from state.
func (r *Ping6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for ping6; removed from Terraform state")
}
