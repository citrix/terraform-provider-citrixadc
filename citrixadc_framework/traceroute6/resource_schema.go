package traceroute6

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

// traceroute6ResourceType is the NITRO resource type string. There is no
// service.Traceroute6 enum constant in the vendored adc-nitro-go, so it is
// declared here (exact NITRO name).
const traceroute6ResourceType = "traceroute6"

// Traceroute6ResourceModel describes the resource data model.
//
// The traceroute6 NITRO parameters do not collide when lowercased, so each
// Terraform attribute name is simply the lowercased NITRO parameter name. The
// payload sent to NITRO always uses the EXACT NITRO parameter name (handled by
// the vendored utility.Traceroute6 struct's json tags; e.g. tf "i" -> json "I",
// tf "t" -> json "T").
type Traceroute6ResourceModel struct {
	Id        types.String `tfsdk:"id"`
	N         types.Bool   `tfsdk:"n"`
	I         types.Bool   `tfsdk:"i"`
	R         types.Bool   `tfsdk:"r"`
	V         types.Bool   `tfsdk:"v"`
	M         types.Int64  `tfsdk:"m"`
	P         types.Int64  `tfsdk:"p"`
	Q         types.Int64  `tfsdk:"q"`
	S         types.String `tfsdk:"s"`
	T         types.Int64  `tfsdk:"t"`
	W         types.Int64  `tfsdk:"w"`
	Host      types.String `tfsdk:"host"`
	Packetlen types.Int64  `tfsdk:"packetlen"`
}

func (r *Traceroute6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the traceroute6 resource. Synthetic value (traceroute6-config); the NITRO traceroute6 action exposes no readable object.",
			},
			"n": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Print hop addresses numerically rather than symbolically and numerically.",
			},
			"i": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use ICMP ECHO for probes. (NITRO parameter: I)",
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
			"m": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Maximum hop value for outgoing probe packets. For Nitro API, default value is taken as 10.",
			},
			"p": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Base port number used in probes.",
			},
			"q": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of probes per hop. For Nitro API, default value is taken as 1.",
			},
			"s": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance, an error is returned and nothing is sent.",
			},
			"t": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Traffic Domain Id. (NITRO parameter: T)",
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

func traceroute6GetThePayloadFromthePlan(ctx context.Context, data *Traceroute6ResourceModel) utility.Traceroute6 {
	tflog.Debug(ctx, "In traceroute6GetThePayloadFromthePlan Function")

	traceroute6 := utility.Traceroute6{}

	if !data.N.IsNull() && !data.N.IsUnknown() {
		traceroute6.N = data.N.ValueBool()
	}
	if !data.I.IsNull() && !data.I.IsUnknown() {
		traceroute6.I = data.I.ValueBool()
	}
	if !data.R.IsNull() && !data.R.IsUnknown() {
		traceroute6.R = data.R.ValueBool()
	}
	if !data.V.IsNull() && !data.V.IsUnknown() {
		traceroute6.V = data.V.ValueBool()
	}
	if !data.M.IsNull() && !data.M.IsUnknown() {
		v := int(data.M.ValueInt64())
		traceroute6.M = &v
	}
	if !data.P.IsNull() && !data.P.IsUnknown() {
		v := int(data.P.ValueInt64())
		traceroute6.P = &v
	}
	if !data.Q.IsNull() && !data.Q.IsUnknown() {
		v := int(data.Q.ValueInt64())
		traceroute6.Q = &v
	}
	if !data.S.IsNull() && !data.S.IsUnknown() {
		traceroute6.S = data.S.ValueString()
	}
	if !data.T.IsNull() && !data.T.IsUnknown() {
		v := int(data.T.ValueInt64())
		traceroute6.T = &v
	}
	if !data.W.IsNull() && !data.W.IsUnknown() {
		v := int(data.W.ValueInt64())
		traceroute6.W = &v
	}
	if !data.Host.IsNull() && !data.Host.IsUnknown() {
		traceroute6.Host = data.Host.ValueString()
	}
	if !data.Packetlen.IsNull() && !data.Packetlen.IsUnknown() {
		v := int(data.Packetlen.ValueInt64())
		traceroute6.Packetlen = &v
	}

	return traceroute6
}
