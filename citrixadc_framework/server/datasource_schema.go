package server

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServerDataSourceModel is the data-source-specific model, decoupled from
// ServerResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (statechangetimesec, tickssincelaststatechange, autoscale, usip, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type ServerDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Internal           types.Bool   `tfsdk:"internal"`
	Comment            types.String `tfsdk:"comment"`
	Delay              types.Int64  `tfsdk:"delay"`
	Domain             types.String `tfsdk:"domain"`
	Domainresolvenow   types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry types.Int64  `tfsdk:"domainresolveretry"`
	Graceful           types.String `tfsdk:"graceful"`
	Ipaddress          types.String `tfsdk:"ipaddress"`
	Ipv6address        types.String `tfsdk:"ipv6address"`
	Name               types.String `tfsdk:"name"` // Required lookup key
	Newname            types.String `tfsdk:"newname"`
	Querytype          types.String `tfsdk:"querytype"`
	State              types.String `tfsdk:"state"`
	Td                 types.Int64  `tfsdk:"td"`
	Translationip      types.String `tfsdk:"translationip"`
	Translationmask    types.String `tfsdk:"translationmask"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/server.json). Never settable; populated from GET.
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Autoscale                 types.String `tfsdk:"autoscale"`
	Usip                      types.String `tfsdk:"usip"`
	Cka                       types.String `tfsdk:"cka"`
	Tcpb                      types.String `tfsdk:"tcpb"`
	Cmp                       types.String `tfsdk:"cmp"`
	Cacheable                 types.String `tfsdk:"cacheable"`
	Sp                        types.String `tfsdk:"sp"`
}

func ServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"internal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display names of the servers that have been created for internal use.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the server.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which all the services configured on the server are disabled.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the server. For a domain based configuration, you must create the server first.",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the NetScaler must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.",
			},
			"ipv6address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the server.\nMust begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the name is created.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"querytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"translationip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address used to transform the server's DNS-resolved IP address.",
			},
			"translationmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the translation ip",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
			"autoscale": schema.StringAttribute{
				Computed:    true,
				Description: "Auto scale option for a servicegroup. Possible values: DISABLED, DNS, POLICY, CLOUD, API.",
			},
			"usip": schema.StringAttribute{
				Computed:    true,
				Description: "Use the client's IP address as the source IP address when initiating a connection to the server. Possible values: YES, NO.",
			},
			"cka": schema.StringAttribute{
				Computed:    true,
				Description: "Enable client keep-alive for the service group. Possible values: YES, NO.",
			},
			"tcpb": schema.StringAttribute{
				Computed:    true,
				Description: "Enable TCP buffering for the service group. Possible values: YES, NO.",
			},
			"cmp": schema.StringAttribute{
				Computed:    true,
				Description: "Enable compression for the specified service. Possible values: YES, NO.",
			},
			"cacheable": schema.StringAttribute{
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward the request to the cache server. Possible values: YES, NO.",
			},
			"sp": schema.StringAttribute{
				Computed:    true,
				Description: "Enable surge protection for the service group. Possible values: ON, OFF.",
			},
		},
	}
}

// serverDataSourceSetAttrFromGet projects a NITRO server GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) — no unknown->null resolution or plan preservation is required. The
// shared utils.MapGet* helpers implement that projection.
func serverDataSourceSetAttrFromGet(ctx context.Context, data *ServerDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In serverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Internal = utils.MapGetBool(g, "Internal")
	data.Comment = utils.MapGetString(g, "comment")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Domain = utils.MapGetString(g, "domain")
	data.Domainresolvenow = utils.MapGetBool(g, "domainresolvenow")
	data.Domainresolveretry = utils.MapGetInt64(g, "domainresolveretry")
	data.Graceful = utils.MapGetString(g, "graceful")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Ipv6address = utils.MapGetString(g, "ipv6address")
	data.Querytype = utils.MapGetString(g, "querytype")
	data.State = utils.MapGetString(g, "state")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Translationip = utils.MapGetString(g, "translationip")
	data.Translationmask = utils.MapGetString(g, "translationmask")

	// newname is rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Autoscale = utils.MapGetString(g, "autoscale")
	data.Usip = utils.MapGetString(g, "usip")
	data.Cka = utils.MapGetString(g, "cka")
	data.Tcpb = utils.MapGetString(g, "tcpb")
	data.Cmp = utils.MapGetString(g, "cmp")
	data.Cacheable = utils.MapGetString(g, "cacheable")
	data.Sp = utils.MapGetString(g, "sp")
}
