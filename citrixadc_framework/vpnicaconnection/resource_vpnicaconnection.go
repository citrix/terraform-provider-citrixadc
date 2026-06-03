package vpnicaconnection

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnicaconnectionResource{}
var _ resource.ResourceWithConfigure = (*VpnicaconnectionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnicaconnectionResource)(nil)

func NewVpnicaconnectionResource() resource.Resource {
	return &VpnicaconnectionResource{}
}

// VpnicaconnectionResource defines the resource implementation.
// vpnicaconnection is an ACTION-ONLY resource: NITRO exposes get(all), count and
// kill (POST ?action=kill). There is no add, update/set, or delete endpoint.
type VpnicaconnectionResource struct {
	client *service.NitroClient
}

func (r *VpnicaconnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnicaconnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnicaconnection"
}

func (r *VpnicaconnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// composeId builds the synthetic ID from the kill arguments. The killed
// connection is not a queryable managed object, so this ID is purely a Terraform
// state handle, not a NITRO lookup key.
func composeId(data *VpnicaconnectionResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("transproto:%s", utils.UrlEncode(data.Transproto.ValueString())))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(data.Username.ValueString())))
	return strings.Join(idParts, ",")
}

func (r *VpnicaconnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnicaconnectionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (kill action) vpnicaconnection resource")
	payload := vpnicaconnectionGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Vpnicaconnection.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill vpnicaconnection, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed vpnicaconnection")

	// Set synthetic ID once, here in Create.
	data.Id = types.StringValue(composeId(&data))

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnicaconnectionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op. vpnicaconnection.kill terminates active ICA
	// connections; the killed connection is not a persistent object keyed by
	// these inputs, so a GET cannot stably re-resolve "this" record. Re-fetching
	// would cause perpetual drift / spurious resource removal. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for vpnicaconnection (kill action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnicaconnectionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op: NITRO exposes no update/set endpoint for vpnicaconnection,
	// and every attribute is RequiresReplace, so Terraform never reaches here with
	// real diffs (Pattern 5).
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for vpnicaconnection; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnicaconnectionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "kill" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for vpnicaconnection. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for vpnicaconnection; kill has no inverse NITRO endpoint")
}
