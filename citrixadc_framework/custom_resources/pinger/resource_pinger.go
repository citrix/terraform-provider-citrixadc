package pinger

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/utility"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkresource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PingerResource{}
var _ resource.ResourceWithConfigure = (*PingerResource)(nil)

func NewPingerResource() resource.Resource {
	return &PingerResource{}
}

// PingerResource models the NITRO `ping` utility action (mirrors the legacy SDKv2
// citrixadc_pinger resource).
//
// It is a one-shot side-effect action: Create fires the ping via
// ActOnResource("ping", ...); there is no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. Every input attribute is RequiresReplace (the
// SDK v2 schema marked them all ForceNew), so re-running the ping happens through
// resource replacement, not Update. The ID is a synthetic PrefixedUniqueId handle,
// identical to the SDK v2 resource.
//
// The forcenew_id_set attribute carries no payload; it existed in the SDK v2
// schema purely as an extra ForceNew trigger that lets a user force a new ping by
// changing the set. It is preserved verbatim for backward compatibility.
type PingerResource struct {
	client *service.NitroClient
}

// PingerResourceModel describes the resource data model. Every schema attribute
// has a matching tfsdk field, with names/types identical to the SDK v2 schema.
type PingerResourceModel struct {
	Id            types.String `tfsdk:"id"`
	C             types.Int64  `tfsdk:"c"`
	Hostname      types.String `tfsdk:"hostname"`
	I             types.Int64  `tfsdk:"i"`
	N             types.Bool   `tfsdk:"n"`
	P             types.String `tfsdk:"p"`
	Q             types.Bool   `tfsdk:"q"`
	S             types.Int64  `tfsdk:"s"`
	T             types.Int64  `tfsdk:"t"`
	ForcenewIdSet types.Set    `tfsdk:"forcenew_id_set"`
}

func (r *PingerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pinger"
}

func (r *PingerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PingerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the pinger resource.",
			},
			// Every input attribute mirrors the SDK v2 schema exactly:
			// Optional + Computed + ForceNew (RequiresReplace). Create always
			// resolves an omitted (unknown) value to null, so Computed is safe.
			"c": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of packets to send.",
			},
			"hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Address of host to ping.",
			},
			"i": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Waiting time, in seconds. The default value is 1 second.",
			},
			"n": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Numeric output only. No name resolution.",
			},
			"p": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pattern to fill in packets. Can be up to 16 bytes, useful for diagnosing data-dependent problems.",
			},
			"q": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Quiet output. Only the summary is printed.",
			},
			"s": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Data size, in bytes. The default value is 56.",
			},
			"t": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Traffic Domain Id.",
			},
			"forcenew_id_set": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Description: "Helper set whose value forces a new ping (ForceNew) when changed. Not sent to NITRO.",
			},
		},
	}
}

func (r *PingerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PingerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating pinger resource")

	// Build the ping payload exactly as the SDK v2 resource did: hostname/n/p/q
	// are always set (their zero values when unset), while c/i/s/t are pointers
	// only populated when the attribute is present in the configuration.
	ping := pingerGetThePayloadFromthePlan(ctx, &data)

	// Fire the ping action. There is no GET endpoint, so no read-back.
	if err := r.client.ActOnResource("ping", &ping, ""); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to ping, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Performed ping action")

	// Synthetic ID for the action-only resource; identical to the SDK v2
	// resource.PrefixedUniqueId("tf-pinger-").
	data.Id = types.StringValue(sdkresource.PrefixedUniqueId("tf-pinger-"))

	// Resolve any omitted (unknown) Computed attributes to null so the applied
	// state is fully known (Read is a no-op and never populates them).
	pingerResolveUnknowns(&data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PingerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The ping is a one-shot action. NITRO has no GET endpoint that reports its
	// state, so Read is a pure preserve-state no-op (matches SDK v2 schema.Noop).
	var data PingerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for pinger; NITRO has no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PingerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the ping action; every schema attribute is
	// RequiresReplace (SDK v2 ForceNew), so Terraform never invokes Update for a
	// real change. Preserve prior state.
	var data, state PingerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	pingerResolveUnknowns(&data)

	tflog.Debug(ctx, "Update is a no-op for pinger; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PingerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// The ping is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete only removes the resource from Terraform state (matches SDK v2
	// schema.Noop).
	tflog.Debug(ctx, "Delete is a no-op for pinger; NITRO has no inverse of the ping action")
}

// pingerGetThePayloadFromthePlan builds the utility.Ping body from the plan,
// mirroring the SDK v2 createPingerFunc. hostname/n/p/q are set unconditionally
// (their zero values when unset); c/i/s/t are pointer fields populated only when
// the corresponding attribute is present in the configuration. The `t` attribute
// maps to Ping.T (NITRO "T", Traffic Domain Id) exactly as in SDK v2.
func pingerGetThePayloadFromthePlan(ctx context.Context, data *PingerResourceModel) utility.Ping {
	tflog.Debug(ctx, "In pingerGetThePayloadFromthePlan Function")

	ping := utility.Ping{
		HostName: data.Hostname.ValueString(),
		N:        data.N.ValueBool(),
		P:        data.P.ValueString(),
		Q:        data.Q.ValueBool(),
	}

	if !data.C.IsNull() && !data.C.IsUnknown() {
		ping.C = utils.IntPtr(int(data.C.ValueInt64()))
	}
	if !data.I.IsNull() && !data.I.IsUnknown() {
		ping.I = utils.IntPtr(int(data.I.ValueInt64()))
	}
	if !data.S.IsNull() && !data.S.IsUnknown() {
		ping.S = utils.IntPtr(int(data.S.ValueInt64()))
	}
	if !data.T.IsNull() && !data.T.IsUnknown() {
		ping.T = utils.IntPtr(int(data.T.ValueInt64()))
	}

	return ping
}

// pingerResolveUnknowns converts any Optional+Computed attribute left unknown
// (because it was omitted from configuration) into a null value of its type.
// Read never populates these attributes, so leaving them unknown would produce
// an "inconsistent result after apply" error (Pattern 13). Null is a known value
// and preserves the SDK v2 behavior of an unset attribute.
func pingerResolveUnknowns(data *PingerResourceModel) {
	if data.C.IsUnknown() {
		data.C = types.Int64Null()
	}
	if data.Hostname.IsUnknown() {
		data.Hostname = types.StringNull()
	}
	if data.I.IsUnknown() {
		data.I = types.Int64Null()
	}
	if data.N.IsUnknown() {
		data.N = types.BoolNull()
	}
	if data.P.IsUnknown() {
		data.P = types.StringNull()
	}
	if data.Q.IsUnknown() {
		data.Q = types.BoolNull()
	}
	if data.S.IsUnknown() {
		data.S = types.Int64Null()
	}
	if data.T.IsUnknown() {
		data.T = types.Int64Null()
	}
	if data.ForcenewIdSet.IsUnknown() {
		data.ForcenewIdSet = types.SetNull(types.StringType)
	}
}
