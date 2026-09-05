package nssimpleacl6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nssimpleacl6DataSourceModel is the data-source-specific model, decoupled from
// Nssimpleacl6ResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attribute the resource deliberately omits
// (hits). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type Nssimpleacl6DataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Aclaction   types.String `tfsdk:"aclaction"`
	Aclname     types.String `tfsdk:"aclname"` // Required lookup key
	Destport    types.Int64  `tfsdk:"destport"`
	Estsessions types.Bool   `tfsdk:"estsessions"`
	Protocol    types.String `tfsdk:"protocol"`
	Srcipv6     types.String `tfsdk:"srcipv6"`
	Td          types.Int64  `tfsdk:"td"`
	Ttl         types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) runtime metadata from the NITRO doc read-only set
	// (zion73x_readonly/nssimpleacl6.json). Never settable; populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func Nssimpleacl6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aclaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop incoming IPv6 packets that match the simple ACL6 rule.",
			},
			"aclname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the simple ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL6 rule is created.",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to match against the destination port number of an incoming IPv6 packet.\n\nDestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocol simple ACL6 rule, which matches any port and any protocol. In that case, you cannot create another simple ACL6 rule specifying a specific port and the same source IPv6 address.",
			},
			"estsessions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol to match against the protocol of an incoming IPv6 packet. You must set this parameter if you set the Destination Port parameter.",
			},
			"srcipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to match against the source IP address of an incoming IPv6 packet.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds, in multiples of four, after which the simple ACL6 rule expires. If you do not want the simple ACL6 rule to expire, do not specify a TTL value.",
			},

			// Read-only (GET-only) runtime metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits for this SACL6 rule.",
			},
		},
	}
}

// nssimpleacl6DataSourceSetAttrFromGet projects a NITRO nssimpleacl6 GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nssimpleacl6DataSourceSetAttrFromGet(ctx context.Context, data *Nssimpleacl6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nssimpleacl6DataSourceSetAttrFromGet Function")

	if v, ok := g["aclname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Aclname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Aclaction = utils.MapGetString(g, "aclaction")
	data.Destport = utils.MapGetInt64(g, "destport")
	data.Estsessions = utils.MapGetBool(g, "estsessions")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Srcipv6 = utils.MapGetString(g, "srcipv6")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only runtime metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
}
