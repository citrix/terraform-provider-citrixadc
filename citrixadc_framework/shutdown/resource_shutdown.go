package shutdown

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// shutdown is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the Shutdown action: POST /nitro/v1/config/shutdown
//     (no ?action= query param, empty body), which shuts down the appliance.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the shutdown action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for shutdown
//     (Pattern 13) and the resource cannot be verified by reading it back.
//   - WARNING: applying this resource SHUTS DOWN the Citrix ADC. It is intended
//     for deliberate, operator-initiated use only.
var _ resource.Resource = &ShutdownResource{}
var _ resource.ResourceWithConfigure = (*ShutdownResource)(nil)

func NewShutdownResource() resource.Resource {
	return &ShutdownResource{}
}

// ShutdownResource defines the resource implementation.
type ShutdownResource struct {
	client *service.NitroClient
}

func (r *ShutdownResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_shutdown"
}

func (r *ShutdownResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ShutdownResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ShutdownResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Shutting down the appliance (shutdown action)")
	shutdown := shutdownGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the Shutdown action
	// (bare POST /config/shutdown, empty action verb like reboot).
	err := r.client.ActOnResource(service.Shutdown.Type(), &shutdown, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to shut down the appliance, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered appliance shutdown")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("shutdown-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. shutdown has no GET endpoint; there is nothing to reconcile.
func (r *ShutdownResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ShutdownResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for shutdown; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. shutdown has no attributes and no set endpoint.
func (r *ShutdownResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ShutdownResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for shutdown; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. shutdown has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *ShutdownResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for shutdown; NITRO has no delete endpoint")
}
