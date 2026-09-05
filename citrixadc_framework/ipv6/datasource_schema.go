package ipv6

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ipv6DataSourceModel is the data-source-specific model, decoupled from
// Ipv6ResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only ND6 timing attributes that the resource
// deliberately omits (basereachtime, reachtime, ndreachtime,
// retransmissiontime). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type Ipv6DataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Dodad                types.String `tfsdk:"dodad"`
	Natprefix            types.String `tfsdk:"natprefix"`
	Ndbasereachtime      types.Int64  `tfsdk:"ndbasereachtime"`
	Ndretransmissiontime types.Int64  `tfsdk:"ndretransmissiontime"`
	Ralearning           types.String `tfsdk:"ralearning"`
	Routerredirection    types.String `tfsdk:"routerredirection"`
	Td                   types.Int64  `tfsdk:"td"`
	Usipnatprefix        types.String `tfsdk:"usipnatprefix"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/ipv6.json). Never settable; populated from GET.
	Basereachtime      types.Int64 `tfsdk:"basereachtime"`
	Reachtime          types.Int64 `tfsdk:"reachtime"`
	Ndreachtime        types.Int64 `tfsdk:"ndreachtime"`
	Retransmissiontime types.Int64 `tfsdk:"retransmissiontime"`
}

func Ipv6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dodad": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Duplicate Address\nDetection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration.",
			},
			"natprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.",
			},
			"ndbasereachtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation.",
			},
			"ndretransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device.",
			},
			"ralearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers.",
			},
			"routerredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Router Redirection.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"usipnatprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"basereachtime": schema.Int64Attribute{
				Computed:    true,
				Description: "ND6 base reachable time (ms).",
			},
			"reachtime": schema.Int64Attribute{
				Computed:    true,
				Description: "ND6 computed reachable time (ms).",
			},
			"ndreachtime": schema.Int64Attribute{
				Computed:    true,
				Description: "ND6 computed reachable time (ms).",
			},
			"retransmissiontime": schema.Int64Attribute{
				Computed:    true,
				Description: "ND6 retransmission time (ms).",
			},
		},
	}
}

// ipv6DataSourceSetAttrFromGet projects a NITRO ipv6 GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func ipv6DataSourceSetAttrFromGet(ctx context.Context, data *Ipv6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In ipv6DataSourceSetAttrFromGet Function")

	// td is the config-supplied lookup key; keep it (and use it as the id) even
	// when the GET does not echo it back.
	if v := utils.MapGetInt64(g, "td"); !v.IsNull() {
		data.Td = v
	}
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	// Read/write attributes as read-back outputs.
	data.Dodad = utils.MapGetString(g, "dodad")
	data.Natprefix = utils.MapGetString(g, "natprefix")
	data.Ndbasereachtime = utils.MapGetInt64(g, "ndbasereachtime")
	data.Ndretransmissiontime = utils.MapGetInt64(g, "ndretransmissiontime")
	data.Ralearning = utils.MapGetString(g, "ralearning")
	data.Routerredirection = utils.MapGetString(g, "routerredirection")
	data.Usipnatprefix = utils.MapGetString(g, "usipnatprefix")

	// Read-only metadata.
	data.Basereachtime = utils.MapGetInt64(g, "basereachtime")
	data.Reachtime = utils.MapGetInt64(g, "reachtime")
	data.Ndreachtime = utils.MapGetInt64(g, "ndreachtime")
	data.Retransmissiontime = utils.MapGetInt64(g, "retransmissiontime")
}
