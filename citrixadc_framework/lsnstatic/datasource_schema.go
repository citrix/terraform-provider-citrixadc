package lsnstatic

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnstaticDataSourceModel is the data-source-specific model, decoupled from
// LsnstaticResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as
// Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type LsnstaticDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Destip            types.String `tfsdk:"destip"`
	Dsttd             types.Int64  `tfsdk:"dsttd"`
	Name              types.String `tfsdk:"name"` // Required lookup key
	Natip             types.String `tfsdk:"natip"`
	Natport           types.Int64  `tfsdk:"natport"`
	Nattype           types.String `tfsdk:"nattype"`
	Network6          types.String `tfsdk:"network6"`
	Subscrip          types.String `tfsdk:"subscrip"`
	Subscrport        types.Int64  `tfsdk:"subscrport"`
	Td                types.Int64  `tfsdk:"td"`
	Transportprotocol types.String `tfsdk:"transportprotocol"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lsnstatic.json). Never settable; populated from GET.
	Status types.String `tfsdk:"status"`
}

func LsnstaticDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read a Large Scale NAT (LSN) static mapping entry.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address for the LSN mapping entry.",
			},
			"dsttd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain through which the destination IP address for this LSN mapping entry is reachable from the Citrix ADC.\n\nIf you do not specify an ID, the destination IP address is assumed to be reachable through the default traffic domain, which has an ID of 0.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN static mapping entry. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn static1\" or 'lsn static1').",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address, already existing on the Citrix ADC as type LSN, to be used as NAT IP address for this mapping entry.",
			},
			"natport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NAT port for this LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"nattype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of sessions to be displayed.",
			},
			"network6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "B4 address in DS-Lite setup",
			},
			"subscrip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4(NAT44 & DS-Lite)/IPv6(NAT64) address of an LSN subscriber for the LSN static mapping entry.",
			},
			"subscrport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of the LSN subscriber for the LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain to which the subscriber belongs. \n\nIf you do not specify an ID, the subscriber is assumed to be a part of the default traffic domain.",
			},
			"transportprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol for the LSN mapping entry.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The status of the mapping. Status could be Inactive, if mapping addition failed due to already existing dynamic/static mapping, or port allocation failure. Possible values: ACTIVE, INACTIVE.",
			},
		},
	}
}

// lsnstaticDataSourceSetAttrFromGet projects a NITRO lsnstatic GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func lsnstaticDataSourceSetAttrFromGet(ctx context.Context, data *LsnstaticDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lsnstaticDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Destip = utils.MapGetString(g, "destip")
	data.Dsttd = utils.MapGetInt64(g, "dsttd")
	data.Natip = utils.MapGetString(g, "natip")
	data.Natport = utils.MapGetInt64(g, "natport")
	data.Nattype = utils.MapGetString(g, "nattype")
	data.Network6 = utils.MapGetString(g, "network6")
	data.Subscrip = utils.MapGetString(g, "subscrip")
	data.Subscrport = utils.MapGetInt64(g, "subscrport")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Transportprotocol = utils.MapGetString(g, "transportprotocol")

	// Read-only attributes.
	data.Status = utils.MapGetString(g, "status")
}
