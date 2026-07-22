package traceroute

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// tracerouteResourceType is the NITRO resource type string. There is no
// service.Traceroute enum constant in the vendored adc-nitro-go, so it is
// declared here (exact NITRO name).
const tracerouteResourceType = "traceroute"

// TracerouteResourceModel describes the resource data model.
//
// NOTE on single-letter attribute naming: NITRO exposes several parameters that
// differ only by case (e.g. "s" vs "S", "m" vs "M", "p" vs "P", "t" vs "T").
// Terraform attribute names must be lowercase and unique, so the upper-case
// variant is suffixed with "_upper". The payload sent to NITRO always uses the
// EXACT NITRO parameter name (handled by the vendored utility.Traceroute
// struct's json tags).
type TracerouteResourceModel struct {
	Id        types.String `tfsdk:"id"`
	SUpper    types.Bool   `tfsdk:"s_upper"`
	N         types.Bool   `tfsdk:"n"`
	R         types.Bool   `tfsdk:"r"`
	V         types.Bool   `tfsdk:"v"`
	MUpper    types.Int64  `tfsdk:"m_upper"`
	M         types.Int64  `tfsdk:"m"`
	PUpper    types.String `tfsdk:"p_upper"`
	P         types.Int64  `tfsdk:"p"`
	Q         types.Int64  `tfsdk:"q"`
	S         types.String `tfsdk:"s"`
	TUpper    types.Int64  `tfsdk:"t_upper"`
	T         types.Int64  `tfsdk:"t"`
	W         types.Int64  `tfsdk:"w"`
	Host      types.String `tfsdk:"host"`
	Packetlen types.Int64  `tfsdk:"packetlen"`
}

func (r *TracerouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the traceroute resource. Synthetic value (traceroute-config); the NITRO traceroute action exposes no readable object.",
			},
			"s_upper": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Print a summary of how many probes were not answered for each hop. (NITRO parameter: S)",
			},
			"n": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Print hop addresses numerically instead of symbolically and numerically.",
			},
			"r": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bypass normal routing tables and send directly to a host on an attached network. If the host is not on a directly attached network, an error is returned.",
			},
			"v": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Verbose output. List received ICMP packets other than TIME_EXCEEDED and UNREACHABLE.",
			},
			"m_upper": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Minimum TTL value used in outgoing probe packets. (NITRO parameter: M)",
			},
			"m": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Maximum TTL value used in outgoing probe packets. For Nitro API, default value is taken as 10. (NITRO parameter: m)",
			},
			"p_upper": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send packets of specified IP protocol. The currently supported protocols are UDP and ICMP. (NITRO parameter: P)",
			},
			"p": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Base port number used in probes. (NITRO parameter: p)",
			},
			"q": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of queries per hop. For Nitro API, defalut value is taken as 1.",
			},
			"s": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance, an error is returned and nothing is sent. (NITRO parameter: s)",
			},
			"t_upper": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Traffic Domain Id. (NITRO parameter: T)",
			},
			"t": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Type-of-service in query packets. (NITRO parameter: t)",
			},
			"w": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time (in seconds) to wait for a response to a query. For Nitro API, defalut value is set to 3.",
			},
			"host": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination host IP address or name.",
			},
			"packetlen": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Length (in bytes) of the query packets.",
			},
		},
	}
}

func tracerouteGetThePayloadFromthePlan(ctx context.Context, data *TracerouteResourceModel) utility.Traceroute {
	tflog.Debug(ctx, "In tracerouteGetThePayloadFromthePlan Function")

	traceroute := utility.Traceroute{}

	if !data.SUpper.IsNull() && !data.SUpper.IsUnknown() {
		traceroute.S = data.SUpper.ValueBool()
	}
	if !data.N.IsNull() && !data.N.IsUnknown() {
		traceroute.N = data.N.ValueBool()
	}
	if !data.R.IsNull() && !data.R.IsUnknown() {
		traceroute.R = data.R.ValueBool()
	}
	if !data.V.IsNull() && !data.V.IsUnknown() {
		traceroute.V = data.V.ValueBool()
	}
	if !data.MUpper.IsNull() && !data.MUpper.IsUnknown() {
		v := int(data.MUpper.ValueInt64())
		traceroute.M = &v
	}
	if !data.M.IsNull() && !data.M.IsUnknown() {
		v := int(data.M.ValueInt64())
		traceroute.M_ = &v
	}
	if !data.PUpper.IsNull() && !data.PUpper.IsUnknown() {
		traceroute.P = data.PUpper.ValueString()
	}
	if !data.P.IsNull() && !data.P.IsUnknown() {
		v := int(data.P.ValueInt64())
		traceroute.P_ = &v
	}
	if !data.Q.IsNull() && !data.Q.IsUnknown() {
		v := int(data.Q.ValueInt64())
		traceroute.Q = &v
	}
	if !data.S.IsNull() && !data.S.IsUnknown() {
		traceroute.S_ = data.S.ValueString()
	}
	if !data.TUpper.IsNull() && !data.TUpper.IsUnknown() {
		v := int(data.TUpper.ValueInt64())
		traceroute.T = &v
	}
	if !data.T.IsNull() && !data.T.IsUnknown() {
		v := int(data.T.ValueInt64())
		traceroute.T_ = &v
	}
	if !data.W.IsNull() && !data.W.IsUnknown() {
		v := int(data.W.ValueInt64())
		traceroute.W = &v
	}
	if !data.Host.IsNull() && !data.Host.IsUnknown() {
		traceroute.Host = data.Host.ValueString()
	}
	if !data.Packetlen.IsNull() && !data.Packetlen.IsUnknown() {
		v := int(data.Packetlen.ValueInt64())
		traceroute.Packetlen = &v
	}

	return traceroute
}
