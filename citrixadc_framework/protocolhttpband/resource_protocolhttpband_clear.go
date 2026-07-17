package protocolhttpband

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/protocol"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProtocolhttpbandClearResource{}
var _ resource.ResourceWithConfigure = (*ProtocolhttpbandClearResource)(nil)
var _ resource.ResourceWithImportState = (*ProtocolhttpbandClearResource)(nil)

func NewProtocolhttpbandClearResource() resource.Resource {
	return &ProtocolhttpbandClearResource{}
}

// ProtocolhttpbandClearResource defines the resource implementation.
type ProtocolhttpbandClearResource struct {
	client *service.NitroClient
}

// ProtocolhttpbandClearResourceModel describes the resource data model.
//
// This resource models the NITRO protocolhttpband `?action=clear` action. clear
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The clear payload carries only `type`.
type ProtocolhttpbandClearResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (r *ProtocolhttpbandClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ProtocolhttpbandClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocolhttpband_clear"
}

func (r *ProtocolhttpbandClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ProtocolhttpbandClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the protocolhttpband_clear resource.",
			},
			// NITRO clear payload marks `type` as mandatory.
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of HTTP band statistics to clear. Possible values: [ REQUEST, RESPONSE, MQTT_JUMBO_REQ ]",
			},
		},
	}
}

func (r *ProtocolhttpbandClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProtocolhttpbandClearResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating protocolhttpband_clear resource")
	payload := protocolhttpband_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Protocolhttpband.Type(), &payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear protocolhttpband, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared protocolhttpband resource")

	// ID = the band type that was cleared; keeps Read/Delete no-ops addressable
	// by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("protocolhttpband_clear-%v", data.Type.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data ProtocolhttpbandClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for protocolhttpband_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state ProtocolhttpbandClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for protocolhttpband_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProtocolhttpbandClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for protocolhttpband_clear; NITRO has no inverse of the clear action")
}

func protocolhttpband_clearGetThePayloadFromthePlan(ctx context.Context, data *ProtocolhttpbandClearResourceModel) protocol.Protocolhttpband {
	tflog.Debug(ctx, "In protocolhttpband_clearGetThePayloadFromthePlan Function")

	// NITRO `?action=clear` accepts only `type`.
	protocolhttpband := protocol.Protocolhttpband{}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		protocolhttpband.Type = data.Type.ValueString()
	}

	return protocolhttpband
}
