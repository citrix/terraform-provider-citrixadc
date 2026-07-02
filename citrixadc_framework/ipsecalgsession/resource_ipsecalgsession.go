package ipsecalgsession

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
var _ resource.Resource = &IpsecalgsessionResource{}
var _ resource.ResourceWithConfigure = (*IpsecalgsessionResource)(nil)
var _ resource.ResourceWithImportState = (*IpsecalgsessionResource)(nil)

func NewIpsecalgsessionResource() resource.Resource {
	return &IpsecalgsessionResource{}
}

// IpsecalgsessionResource defines the resource implementation.
//
// ipsecalgsession is an ACTION-ONLY runtime object. NITRO exposes ONLY get(all),
// count, and the POST action ?action=flush. There is NO add, NO update/set, NO
// delete. The session table is populated by the IPSec ALG traffic engine, not by
// the config API. This resource therefore fires ?action=flush on Create (scoped
// by the optional sourceip/natip/destip fields; a bare flush flushes all) and
// treats Read/Update/Delete as no-ops.
type IpsecalgsessionResource struct {
	client *service.NitroClient
}

func (r *IpsecalgsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsecalgsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecalgsession"
}

func (r *IpsecalgsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsecalgsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpsecalgsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Firing ipsecalgsession flush action")

	// ipsecalgsession has no add/set/delete; the only write is POST ?action=flush.
	// Payload carries the optional scope fields sourceip/natip/destip.
	ipsecalgsession := ipsecalgsessionGetTheFlushPayload(ctx, &data)

	err := r.client.ActOnResource(service.Ipsecalgsession.Type(), &ipsecalgsession, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush ipsecalgsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed ipsecalgsession")

	// Synthetic ID: this is an action, not a persistent object. Compose it from the
	// supplied scope so distinct flush scopes produce distinct IDs; a bare flush
	// uses a static ID.
	id := "flush-all"
	if !data.Sourceip.IsNull() && !data.Sourceip.IsUnknown() {
		id = fmt.Sprintf("flush:sourceip:%s", data.Sourceip.ValueString())
	} else if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		id = fmt.Sprintf("flush:natip:%s", data.Natip.ValueString())
	} else if !data.Destip.IsNull() && !data.Destip.IsUnknown() {
		id = fmt.Sprintf("flush:destip:%s", data.Destip.ValueString())
	}
	data.Id = types.StringValue(id)

	// Read is a no-op (session table is ephemeral); persist the plan as-is.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpsecalgsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op. ipsecalgsession is an action-only runtime object: the traffic
	// engine owns the session table, and a fired flush leaves no re-findable object
	// keyed to our synthetic ID. Preserve prior state unchanged.
	tflog.Debug(ctx, "Read is a no-op for ipsecalgsession (action-only resource, ephemeral session table)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state IpsecalgsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op: ipsecalgsession has no set endpoint and every input
	// attribute is RequiresReplace, so a meaningful change forces recreation and
	// never reaches here.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for ipsecalgsession; all attributes are RequiresReplace and there is no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpsecalgsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op: ipsecalgsession exposes no delete operation. The prior
	// flush is not reversible via the config API; removing the resource simply drops
	// it from state.
	tflog.Debug(ctx, "Delete is a no-op for ipsecalgsession (no delete operation on NITRO side)")
}
