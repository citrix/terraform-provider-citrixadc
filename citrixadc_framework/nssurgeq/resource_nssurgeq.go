package nssurgeq

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NssurgeqResource{}
var _ resource.ResourceWithConfigure = (*NssurgeqResource)(nil)
var _ resource.ResourceWithImportState = (*NssurgeqResource)(nil)
var _ resource.ResourceWithValidateConfig = (*NssurgeqResource)(nil)

func NewNssurgeqResource() resource.Resource {
	return &NssurgeqResource{}
}

// NssurgeqResource defines the resource implementation.
type NssurgeqResource struct {
	client *service.NitroClient
}

func (r *NssurgeqResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NssurgeqResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssurgeq"
}

func (r *NssurgeqResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the dependency hierarchy: servername requires name,
// port requires servername. flush may also be run with no arguments at all.
func (r *NssurgeqResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NssurgeqResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nameSet := !data.Name.IsNull() && !data.Name.IsUnknown()
	servernameSet := !data.Servername.IsNull() && !data.Servername.IsUnknown()
	portSet := !data.Port.IsNull() && !data.Port.IsUnknown()

	if servernameSet && !nameSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("servername"),
			"Missing required dependency",
			"servername requires name to be set.",
		)
	}
	if portSet && !servernameSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("port"),
			"Missing required dependency",
			"port requires servername to be set.",
		)
	}
}

func (r *NssurgeqResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssurgeqResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing nssurgeq (action-only resource)")
	payload := nssurgeqGetThePayloadFromthePlan(ctx, &data)

	// nssurgeq is an action-only resource (flush, no add/get).
	// The Create maps to the "flush" action.
	err := r.client.ActOnResource(service.Nssurgeq.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush nssurgeq, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed nssurgeq resource")

	// Generate a synthetic ID; nssurgeq has no GET endpoint.
	data.Id = types.StringValue("nssurgeq-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NssurgeqResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: nssurgeq is an action-only resource with no GET endpoint.
	tflog.Debug(ctx, "Read is a no-op for nssurgeq (no GET endpoint); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NssurgeqResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for nssurgeq; all attributes are RequiresReplace.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nssurgeq; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NssurgeqResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op: nssurgeq is action-only with no inverse API; just remove from state.
	tflog.Debug(ctx, "Delete is a no-op for nssurgeq; removing from Terraform state")
}
