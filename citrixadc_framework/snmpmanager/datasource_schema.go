package snmpmanager

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpmanagerDataSourceModel is the data-source-specific model, decoupled from
// SnmpmanagerResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type SnmpmanagerDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Domainresolveretry types.Int64  `tfsdk:"domainresolveretry"`
	Ipaddress          types.String `tfsdk:"ipaddress"` // Required lookup key
	Netmask            types.String `tfsdk:"netmask"`   // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/snmpmanager.json). Never settable; populated from GET.
	Ip     types.String `tfsdk:"ip"`
	Domain types.String `tfsdk:"domain"`
}

func SnmpmanagerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the SNMP manager if the last query failed. This parameter is valid for host-name based SNMP managers only. After a query succeeds, the TTL determines the wait time. The minimum and default value is 5.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the SNMP manager. Can be an IPv4 or IPv6 address. You can instead specify an IPv4 network address or IPv6 network prefix if you want the Citrix ADC to respond to SNMP queries from any device on the specified network. Alternatively, instead of an IPv4 address, you can specify a host name that has been assigned to an SNMP manager. If you do so, you must add a DNS name server that resolves the host name of the SNMP manager to its IP address. \nNote: The Citrix ADC does not support host names for SNMP managers that have IPv6 addresses.",
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "Subnet mask associated with an IPv4 network address. If the IP address specifies the address or host name of a specific host, accept the default value of 255.255.255.255.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved IP address of the hostname manager.",
			},
			"domain": schema.StringAttribute{
				Computed:    true,
				Description: "IP address of manager. It will be zero for hostname manager.",
			},
		},
	}
}

// snmpmanagerDataSourceSetAttrFromGet projects a NITRO snmpmanager GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func snmpmanagerDataSourceSetAttrFromGet(ctx context.Context, data *SnmpmanagerDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpmanagerDataSourceSetAttrFromGet Function")

	if v, ok := g["ipaddress"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Ipaddress = types.StringValue(utils.AnyToString(v))
	}

	data.Domainresolveretry = utils.MapGetInt64(g, "domainresolveretry")
	data.Netmask = utils.MapGetString(g, "netmask")

	// Read-only (GET-only) metadata.
	data.Ip = utils.MapGetString(g, "ip")
	data.Domain = utils.MapGetString(g, "domain")
}
