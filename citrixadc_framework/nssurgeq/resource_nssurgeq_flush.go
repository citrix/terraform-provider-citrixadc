package nssurgeq

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NssurgeqFlushResource{}
var _ resource.ResourceWithConfigure = (*NssurgeqFlushResource)(nil)
var _ resource.ResourceWithImportState = (*NssurgeqFlushResource)(nil)

func NewNssurgeqFlushResource() resource.Resource {
	return &NssurgeqFlushResource{}
}

// NssurgeqFlushResource defines the resource implementation.
type NssurgeqFlushResource struct {
	client *service.NitroClient
}

// NssurgeqFlushResourceModel describes the resource data model.
//
// This resource models the NITRO nssurgeq `?action=flush` action. flush is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The flush payload carries the optional filter
// attributes name, port and servername.
type NssurgeqFlushResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Port       types.Int64  `tfsdk:"port"`
	Servername types.String `tfsdk:"servername"`
}

func (r *NssurgeqFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NssurgeqFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssurgeq_flush"
}

func (r *NssurgeqFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NssurgeqFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssurgeq_flush resource.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a virtual server, service or service group for which the SurgeQ must be flushed.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "port on which server is bound to the entity(Servicegroup).",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a service group member. This argument is needed when you want to flush the SurgeQ of a service group.",
			},
		},
	}
}

func (r *NssurgeqFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssurgeqFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing nssurgeq (action-only resource)")
	payload := nssurgeq_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Nssurgeq.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush nssurgeq, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed nssurgeq")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nssurgeq_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data NssurgeqFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nssurgeq_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state NssurgeqFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nssurgeq_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssurgeqFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nssurgeq_flush; NITRO has no inverse of the flush action")
}

// nssurgeq_flushGetThePayloadFromthePlan builds the action payload, including only the set args.
func nssurgeq_flushGetThePayloadFromthePlan(ctx context.Context, data *NssurgeqFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nssurgeq_flushGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payload["name"] = data.Name.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		payload["servername"] = data.Servername.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		payload["port"] = int(data.Port.ValueInt64())
	}

	return payload
}
