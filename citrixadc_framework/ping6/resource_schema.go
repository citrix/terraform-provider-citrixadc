package ping6

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

// ping6ResourceType is the NITRO resource type string. There is no
// service.Ping6 enum constant in the vendored adc-nitro-go, so it is declared
// here (exact NITRO name).
const ping6ResourceType = "ping6"

// Ping6ResourceModel describes the resource data model.
//
// NOTE on single-letter attribute naming: NITRO exposes several parameters that
// differ only by case (e.g. "i" vs "I", "s" vs "S", "t" vs "T"). Terraform
// attribute names must be lowercase and unique, so the upper-case variant is
// suffixed with "_upper". The payload sent to NITRO always uses the EXACT NITRO
// parameter name (handled by the vendored utility.Ping6 struct's json tags).
type Ping6ResourceModel struct {
	Id       types.String `tfsdk:"id"`
	B        types.Int64  `tfsdk:"b"`
	C        types.Int64  `tfsdk:"c"`
	I        types.Int64  `tfsdk:"i"`
	IUpper   types.String `tfsdk:"i_upper"`
	M        types.Bool   `tfsdk:"m"`
	N        types.Bool   `tfsdk:"n"`
	P        types.String `tfsdk:"p"`
	Q        types.Bool   `tfsdk:"q"`
	S        types.Int64  `tfsdk:"s"`
	V        types.Int64  `tfsdk:"v"`
	SUpper   types.String `tfsdk:"s_upper"`
	TUpper   types.Int64  `tfsdk:"t_upper"`
	T        types.Int64  `tfsdk:"t"`
	HostName types.String `tfsdk:"hostname"`
}

func (r *Ping6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ping6 resource. Synthetic value (ping6-config); the NITRO ping6 action exposes no readable object.",
			},
			"b": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Set socket buffer size. If used, should be used with roughly +100 then the datalen (-s option). The default value is 8192.",
			},
			"c": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of packets to send. The default value is infinite. For Nitro API, defalut value is taken as 5.",
			},
			"i": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Waiting time, in seconds. The default value is 1 second. (NITRO parameter: i)",
			},
			"i_upper": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Network interface on which to ping, if you have multiple interfaces. (NITRO parameter: I)",
			},
			"m": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "By default, ping6 asks the kernel to fragment packets to fit into the minimum IPv6 MTU. The -m option will suppress the behavior for unicast packets.",
			},
			"n": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Numeric output only. No name resolution.",
			},
			"p": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pattern to fill in packets. Can be up to 16 bytes, useful for diagnosing data-dependent problems.",
			},
			"q": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Quiet output. Only summary is printed. For Nitro API, this flag is set by default.",
			},
			"s": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Data size, in bytes. The default value is 32. (NITRO parameter: s)",
			},
			"v": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "VLAN ID for link local address. (NITRO parameter: V)",
			},
			"s_upper": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source IP address to be used in the outgoing query packets. (NITRO parameter: S)",
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
				Description: "Timeout in seconds before ping6 exits. (NITRO parameter: t)",
			},
			"hostname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Address of host to ping. (NITRO parameter: hostName)",
			},
		},
	}
}

func ping6GetThePayloadFromthePlan(ctx context.Context, data *Ping6ResourceModel) utility.Ping6 {
	tflog.Debug(ctx, "In ping6GetThePayloadFromthePlan Function")

	ping6 := utility.Ping6{}

	if !data.B.IsNull() && !data.B.IsUnknown() {
		v := int(data.B.ValueInt64())
		ping6.B = &v
	}
	if !data.C.IsNull() && !data.C.IsUnknown() {
		v := int(data.C.ValueInt64())
		ping6.C = &v
	}
	if !data.I.IsNull() && !data.I.IsUnknown() {
		v := int(data.I.ValueInt64())
		ping6.I = &v
	}
	if !data.IUpper.IsNull() && !data.IUpper.IsUnknown() {
		ping6.I_ = data.IUpper.ValueString()
	}
	if !data.M.IsNull() && !data.M.IsUnknown() {
		ping6.M = data.M.ValueBool()
	}
	if !data.N.IsNull() && !data.N.IsUnknown() {
		ping6.N = data.N.ValueBool()
	}
	if !data.P.IsNull() && !data.P.IsUnknown() {
		ping6.P = data.P.ValueString()
	}
	if !data.Q.IsNull() && !data.Q.IsUnknown() {
		ping6.Q = data.Q.ValueBool()
	}
	if !data.S.IsNull() && !data.S.IsUnknown() {
		v := int(data.S.ValueInt64())
		ping6.S = &v
	}
	if !data.V.IsNull() && !data.V.IsUnknown() {
		v := int(data.V.ValueInt64())
		ping6.V = &v
	}
	if !data.SUpper.IsNull() && !data.SUpper.IsUnknown() {
		ping6.S_ = data.SUpper.ValueString()
	}
	if !data.TUpper.IsNull() && !data.TUpper.IsUnknown() {
		v := int(data.TUpper.ValueInt64())
		ping6.T = &v
	}
	if !data.T.IsNull() && !data.T.IsUnknown() {
		v := int(data.T.ValueInt64())
		ping6.T_ = &v
	}
	if !data.HostName.IsNull() && !data.HostName.IsUnknown() {
		ping6.HostName = data.HostName.ValueString()
	}

	return ping6
}
