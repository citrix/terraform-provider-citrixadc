package nsstats

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsstatsClearResource{}
var _ resource.ResourceWithConfigure = (*NsstatsClearResource)(nil)
var _ resource.ResourceWithImportState = (*NsstatsClearResource)(nil)

func NewNsstatsClearResource() resource.Resource {
	return &NsstatsClearResource{}
}

// NsstatsClearResource defines the resource implementation.
type NsstatsClearResource struct {
	client *service.NitroClient
}

// NsstatsClearResourceModel describes the resource data model.
//
// This resource models the NITRO nsstats `?action=clear` action. clear is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The clear payload carries the optional
// cleanuplevel attribute.
type NsstatsClearResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Cleanuplevel types.String `tfsdk:"cleanuplevel"`
}

func (r *NsstatsClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsstatsClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsstats_clear"
}

func (r *NsstatsClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsstatsClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsstats_clear resource.",
			},
			"cleanuplevel": schema.StringAttribute{
				Required:    true,
				Description: "The level of stats to be cleared. 'global' option will clear global counters only, 'all' option will clear all device counters also along with global counters. For both the cases only 'ever incrementing counters' i.e. total counters will be cleared.\nPossible values = global, all",
			},
		},
	}
}

func (r *NsstatsClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsstatsClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing nsstats (action-only resource)")
	payload := nsstats_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Nsstats.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear nsstats, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared nsstats resource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nsstats_clear")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data NsstatsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsstats_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; clear is an action-only resource,
	// so Terraform never invokes Update for a real change.
	var data, state NsstatsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsstats_clear; NITRO has no update endpoint for the clear action")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nsstats_clear; NITRO has no inverse of the clear action")
}

// nsstats_clearGetThePayloadFromthePlan builds the action payload, including only the set args.
func nsstats_clearGetThePayloadFromthePlan(ctx context.Context, data *NsstatsClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nsstats_clearGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Cleanuplevel.IsNull() && !data.Cleanuplevel.IsUnknown() {
		payload["cleanuplevel"] = data.Cleanuplevel.ValueString()
	}

	return payload
}
