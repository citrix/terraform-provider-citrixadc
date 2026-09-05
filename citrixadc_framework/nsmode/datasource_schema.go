package nsmode

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsmodeDataSourceModel describes the DATASOURCE data model. It mirrors the mode
// flags the NITRO `nsmode` GET returns, including the read-only flags the
// resource intentionally omits. It is decoupled from the resource model so the
// data source can expose the full GET projection.
type NsmodeDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Fr                  types.Bool   `tfsdk:"fr"`
	L2                  types.Bool   `tfsdk:"l2"`
	Usip                types.Bool   `tfsdk:"usip"`
	Cka                 types.Bool   `tfsdk:"cka"`
	Tcpb                types.Bool   `tfsdk:"tcpb"`
	Mbf                 types.Bool   `tfsdk:"mbf"`
	Edge                types.Bool   `tfsdk:"edge"`
	Usnip               types.Bool   `tfsdk:"usnip"`
	L3                  types.Bool   `tfsdk:"l3"`
	Pmtud               types.Bool   `tfsdk:"pmtud"`
	Mediaclassification types.Bool   `tfsdk:"mediaclassification"`
	Sradv               types.Bool   `tfsdk:"sradv"`
	Dradv               types.Bool   `tfsdk:"dradv"`
	Iradv               types.Bool   `tfsdk:"iradv"`
	Sradv6              types.Bool   `tfsdk:"sradv6"`
	Dradv6              types.Bool   `tfsdk:"dradv6"`
	Bridgebpdus         types.Bool   `tfsdk:"bridgebpdus"`
	Ulfd                types.Bool   `tfsdk:"ulfd"`

	// Read-only (GET-only) mode flag from the NITRO doc read-only set
	// (zion73x_readonly/nsmode.json). Never settable; from GET.
	SingleIp types.Bool `tfsdk:"single_ip"`
}

// nsmodeDataSourceSetAttrFromGet projects a NITRO nsmode GET response onto the
// data-source model using the shared utils.MapGet* helpers. Flags the GET omits
// are left Null.
func nsmodeDataSourceSetAttrFromGet(ctx context.Context, data *NsmodeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsmodeDataSourceSetAttrFromGet Function")

	data.Fr = utils.MapGetBool(g, "fr")
	data.L2 = utils.MapGetBool(g, "l2")
	data.Usip = utils.MapGetBool(g, "usip")
	data.Cka = utils.MapGetBool(g, "cka")
	data.Tcpb = utils.MapGetBool(g, "tcpb")
	data.Mbf = utils.MapGetBool(g, "mbf")
	data.Edge = utils.MapGetBool(g, "edge")
	data.Usnip = utils.MapGetBool(g, "usnip")
	data.L3 = utils.MapGetBool(g, "l3")
	data.Pmtud = utils.MapGetBool(g, "pmtud")
	data.Mediaclassification = utils.MapGetBool(g, "mediaclassification")
	data.Sradv = utils.MapGetBool(g, "sradv")
	data.Dradv = utils.MapGetBool(g, "dradv")
	data.Iradv = utils.MapGetBool(g, "iradv")
	data.Sradv6 = utils.MapGetBool(g, "sradv6")
	data.Dradv6 = utils.MapGetBool(g, "dradv6")
	data.Bridgebpdus = utils.MapGetBool(g, "bridgebpdus")
	data.Ulfd = utils.MapGetBool(g, "ulfd")

	// Read-only mode flag.
	data.SingleIp = utils.MapGetBool(g, "single_ip")

	// nsmode is a singleton (no unique attributes) - static ID.
	data.Id = types.StringValue("nsmode-config")
}

func NsmodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmode datasource.",
			},
			"fr": schema.BoolAttribute{
				Computed:    true,
				Description: "Fast Ramp mode.",
			},
			"l2": schema.BoolAttribute{
				Computed:    true,
				Description: "Layer 2 mode.",
			},
			"usip": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Source IP mode.",
			},
			"cka": schema.BoolAttribute{
				Computed:    true,
				Description: "Client Keep-Alive mode.",
			},
			"tcpb": schema.BoolAttribute{
				Computed:    true,
				Description: "TCP Buffering mode.",
			},
			"mbf": schema.BoolAttribute{
				Computed:    true,
				Description: "MAC-based forwarding mode.",
			},
			"edge": schema.BoolAttribute{
				Computed:    true,
				Description: "Edge configuration mode.",
			},
			"usnip": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Subnet IP mode.",
			},
			"l3": schema.BoolAttribute{
				Computed:    true,
				Description: "Layer 3 mode.",
			},
			"pmtud": schema.BoolAttribute{
				Computed:    true,
				Description: "Path MTU Discovery mode.",
			},
			"mediaclassification": schema.BoolAttribute{
				Computed:    true,
				Description: "Media classification mode.",
			},
			"sradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Static route advertisement mode.",
			},
			"dradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Dynamic route advertisement mode.",
			},
			"iradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Intranet route advertisement mode.",
			},
			"sradv6": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 static route advertisement mode.",
			},
			"dradv6": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 dynamic route advertisement mode.",
			},
			"bridgebpdus": schema.BoolAttribute{
				Computed:    true,
				Description: "Bridge BPDUs mode.",
			},
			"ulfd": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Layer 2 mode for IPv4 packets.",
			},

			// Read-only (GET-only) mode flag surfaced by the data source.
			"single_ip": schema.BoolAttribute{
				Computed:    true,
				Description: "Single IP mode. Null when the appliance omits it.",
			},
		},
	}
}
