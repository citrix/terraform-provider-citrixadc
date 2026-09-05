package cloudtunnelvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudtunnelvserverDataSourceModel is the data-source-specific model, decoupled
// from CloudtunnelvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attributes that the resource deliberately
// omits (state, effectivestate, ip, port, ...). Every non-key attribute is
// Computed.
type CloudtunnelvserverDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Listenpolicy   types.String `tfsdk:"listenpolicy"`
	Listenpriority types.Int64  `tfsdk:"listenpriority"`
	Name           types.String `tfsdk:"name"` // Required lookup key
	Servicetype    types.String `tfsdk:"servicetype"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cloudtunnelvserver.json). Never settable; populated from GET.
	State          types.String `tfsdk:"state"`
	Effectivestate types.String `tfsdk:"effectivestate"`
	Type           types.String `tfsdk:"type"`
	Ip             types.String `tfsdk:"ip"`
	Ipv46          types.String `tfsdk:"ipv46"`
	Ippattern      types.String `tfsdk:"ippattern"`
	Port           types.Int64  `tfsdk:"port"`
	Range          types.Int64  `tfsdk:"range"`
	Cachetype      types.String `tfsdk:"cachetype"`
}

func CloudtunnelvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the Cloud Tunnel virtual server. Can be either a named expression or an expression. The Cloud Tunnel virtual server processes only the traffic for which the expression evaluates to true.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Cloud Tunnel virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space,colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example,\n\"my server\" or 'my server').",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ServiceType of Listener using which traffic will be tunneled through cloud tunnel server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "The current state of the virtual server, as UP, DOWN, BUSY, and so on.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Effective state of the virtual server.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of virtual server; for example, CONTENT based or ADDRESS based.",
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The Virtual IP address of the Cloud Tunnel virtual server.",
			},
			"ipv46": schema.StringAttribute{
				Computed:    true,
				Description: "IPv4 or IPv6 address of the Cloud Tunnel virtual server.",
			},
			"ippattern": schema.StringAttribute{
				Computed:    true,
				Description: "The IP pattern of the virtual server.",
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: "Port on which the virtual server listens.",
			},
			"range": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of IP addresses that the appliance must generate and assign to the virtual server.",
			},
			"cachetype": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server cache type. The options are: TRANSPARENT, REVERSE, and FORWARD.",
			},
		},
	}
}

// cloudtunnelvserverDataSourceSetAttrFromGet projects a NITRO cloudtunnelvserver
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func cloudtunnelvserverDataSourceSetAttrFromGet(ctx context.Context, data *CloudtunnelvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cloudtunnelvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.Servicetype = utils.MapGetString(g, "servicetype")

	// Read-only attributes.
	data.State = utils.MapGetString(g, "state")
	data.Effectivestate = utils.MapGetString(g, "effectivestate")
	data.Type = utils.MapGetString(g, "type")
	data.Ip = utils.MapGetString(g, "ip")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.Ippattern = utils.MapGetString(g, "ippattern")
	data.Port = utils.MapGetInt64(g, "port")
	data.Range = utils.MapGetInt64(g, "range")
	data.Cachetype = utils.MapGetString(g, "cachetype")
}
