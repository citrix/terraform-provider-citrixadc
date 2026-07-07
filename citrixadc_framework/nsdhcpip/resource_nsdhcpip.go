package nsdhcpip

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nsdhcpip is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the release action:
//     POST /nitro/v1/config/nsdhcpip?action=release, which releases the DHCP
//     lease for the appliance management IP.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the release action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for nsdhcpip.
var _ resource.Resource = &NsdhcpipResource{}
var _ resource.ResourceWithConfigure = (*NsdhcpipResource)(nil)
var _ resource.ResourceWithImportState = (*NsdhcpipResource)(nil)

func NewNsdhcpipResource() resource.Resource {
	return &NsdhcpipResource{}
}

// NsdhcpipResource defines the resource implementation.
type NsdhcpipResource struct {
	client *service.NitroClient
}

func (r *NsdhcpipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsdhcpipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsdhcpip"
}

func (r *NsdhcpipResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsdhcpipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsdhcpipResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsdhcpip resource (release action)")
	nsdhcpip := nsdhcpipGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=release
	err := r.client.ActOnResource(service.Nsdhcpip.Type(), &nsdhcpip, "release")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to release nsdhcpip, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Released nsdhcpip")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("nsdhcpip-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. nsdhcpip has no GET endpoint; there is nothing to reconcile.
func (r *NsdhcpipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsdhcpipResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsdhcpip; NITRO exposes no GET endpoint (action=release only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. nsdhcpip has no attributes and no set endpoint.
func (r *NsdhcpipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsdhcpipResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nsdhcpip; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. nsdhcpip has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *NsdhcpipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nsdhcpip; NITRO has no delete endpoint")
}
