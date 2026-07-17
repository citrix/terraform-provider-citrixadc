package endpointinfo

import (
	"context"
	"fmt"

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
var _ resource.Resource = &EndpointinfoClearResource{}
var _ resource.ResourceWithConfigure = (*EndpointinfoClearResource)(nil)
var _ resource.ResourceWithImportState = (*EndpointinfoClearResource)(nil)

func NewEndpointinfoClearResource() resource.Resource {
	return &EndpointinfoClearResource{}
}

// EndpointinfoClearResource defines the resource implementation.
type EndpointinfoClearResource struct {
	client *service.NitroClient
}

// EndpointinfoClearResourceModel describes the resource data model.
//
// This resource models the NITRO endpointinfo `?action=clear` action. clear is
// a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The clear payload carries only `endpointkind`.
type EndpointinfoClearResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Endpointkind types.String `tfsdk:"endpointkind"`
}

func (r *EndpointinfoClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *EndpointinfoClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpointinfo_clear"
}

func (r *EndpointinfoClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *EndpointinfoClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the endpointinfo_clear resource.",
			},
			// NITRO clear payload exposes only `endpointkind` (optional, enum=[IP], default=IP).
			"endpointkind": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Endpoint kind. Currently, IP endpoints are supported. Possible values: [ IP ]",
			},
		},
	}
}

func (r *EndpointinfoClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EndpointinfoClearResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating endpointinfo_clear resource")
	payload := endpointinfo_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Endpointinfo.Type(), &payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear endpointinfo, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared endpointinfo resource")

	// ID = the endpoint kind that was cleared; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("endpointinfo_clear-%v", data.Endpointkind.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data EndpointinfoClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for endpointinfo_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state EndpointinfoClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for endpointinfo_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for endpointinfo_clear; NITRO has no inverse of the clear action")
}

func endpointinfo_clearGetThePayloadFromthePlan(ctx context.Context, data *EndpointinfoClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In endpointinfo_clearGetThePayloadFromthePlan Function")

	// No vendored NITRO struct exists for endpointinfo (Pattern 3: missing
	// vendored struct -> use a map payload). NITRO `?action=clear` accepts only
	// `endpointkind`.
	endpointinfo := make(map[string]interface{})
	if !data.Endpointkind.IsNull() && !data.Endpointkind.IsUnknown() {
		endpointinfo["endpointkind"] = data.Endpointkind.ValueString()
	}

	return endpointinfo
}
