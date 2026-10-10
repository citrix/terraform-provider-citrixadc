package nsconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsconfigUnsetResource{}
var _ resource.ResourceWithConfigure = (*NsconfigUnsetResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigUnsetResource)(nil)

func NewNsconfigUnsetResource() resource.Resource {
	return &NsconfigUnsetResource{}
}

// NsconfigUnsetResource defines the resource implementation.
type NsconfigUnsetResource struct {
	client *service.NitroClient
}

// NsconfigUnsetResourceModel describes the resource data model.
// Action-only resource (NITRO nsconfig `?action=unset`): it resets the named
// settable nsconfig parameters to their appliance defaults. `attributes` is the
// set of nsconfig parameter names to unset; `timestamp` is a synthetic TF-only
// field used as the resource ID (re-running the action requires bumping it,
// RequiresReplace).
type NsconfigUnsetResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Attributes types.Set    `tfsdk:"attributes"`
	Timestamp  types.String `tfsdk:"timestamp"`
}

func (r *NsconfigUnsetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_unset"
}

func (r *NsconfigUnsetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconfig_unset resource (equals the configured timestamp).",
			},
			"attributes": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Description: "Set of nsconfig parameter names to reset to their appliance defaults via NITRO `unset ns config`. " +
					"Allowed values: nsvlan, securemanagementtraffic, ftpportrange, crportrange, timezone, ipaddress, " +
					"netmask, ifnum, tagged, httpport, maxconn, maxreq, cip, cipheader, cookieversion, securecookie, " +
					"pmtumin, pmtutimeout, grantquotamaxclient, exclusivequotamaxclient, grantquotaspillover, " +
					"exclusivequotaspillover. WARNING: unsetting ipaddress, netmask, ifnum, nsvlan or tagged can disrupt " +
					"appliance connectivity.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp marker used as the resource ID. Change it to re-run unset ns config.",
			},
		},
	}
}

func (r *NsconfigUnsetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigUnsetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsconfigUnsetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigUnsetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_unset resource (unset ns config)")

	var attrs []string
	resp.Diagnostics.Append(data.Attributes.ElementsAs(ctx, &attrs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// nsconfig is an unnamed singleton, so no identity fields are required for the
	// unset payload. utils.ExecuteUnset issues POST ?action=unset with each named
	// attribute set to "true". An empty set is a no-op.
	if err := utils.ExecuteUnset(r.client, service.Nsconfig.Type(), map[string]interface{}{}, attrs); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset ns config, got error: %s", err))
		return
	}

	// Synthetic ID equals the configured timestamp.
	data.Id = data.Timestamp

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUnsetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No GET endpoint for the unset action (Pattern 13): preserve state as-is.
	var data NsconfigUnsetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read is a no-op for nsconfig_unset (action-only, no GET endpoint)")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUnsetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; Update is never expected to run.
	var data, state NsconfigUnsetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsconfig_unset; all attributes are RequiresReplace")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUnsetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Action-only resource: there is no inverse API. Delete just removes from state.
	tflog.Debug(ctx, "Deleting nsconfig_unset: action-only, removing from state only")
}
