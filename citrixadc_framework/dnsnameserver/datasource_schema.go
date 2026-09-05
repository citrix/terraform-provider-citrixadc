package dnsnameserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsnameserverDataSourceModel is the data-source-specific model, decoupled from
// DnsnameserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (servicename, port, nameserverstate, clmonowner, clmonview). The Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares.
type DnsnameserverDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Dnsprofilename types.String `tfsdk:"dnsprofilename"`
	Dnsvservername types.String `tfsdk:"dnsvservername"`
	Ip             types.String `tfsdk:"ip"`
	Local          types.Bool   `tfsdk:"local"`
	State          types.String `tfsdk:"state"`
	Type           types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsnameserver.json). Never settable; populated from GET.
	Servicename     types.String `tfsdk:"servicename"`
	Port            types.Int64  `tfsdk:"port"`
	Nameserverstate types.String `tfsdk:"nameserverstate"`
	Clmonowner      types.Int64  `tfsdk:"clmonowner"`
	Clmonview       types.Int64  `tfsdk:"clmonview"`
}

func DnsnameserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the name server",
			},
			"dnsvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC. Either dnsvservername or ip must be specified.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS). Either dnsvservername or ip must be specified.",
			},
			"local": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set.\n\nIf no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Administrative state of the name server.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the dns vserver.",
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: "Port of the service. Range 1 - 65535 (* in CLI is represented as 65535 in NITRO API).",
			},
			"nameserverstate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the server. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
			"clmonowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the mon owner of the service.",
			},
			"clmonview": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the view id by which state of the service is updated.",
			},
		},
	}
}

// dnsnameserverDataSourceSetAttrFromGet projects a NITRO dnsnameserver GET
// response onto the data-source model and derives the composite id (name,type)
// from the response. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection; local is
// converted through the package-level interfaceToBool for robustness (NITRO may
// return it as a string).
func dnsnameserverDataSourceSetAttrFromGet(ctx context.Context, data *DnsnameserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsnameserverDataSourceSetAttrFromGet Function")

	data.Dnsprofilename = utils.MapGetString(g, "dnsprofilename")
	data.Ip = utils.MapGetString(g, "ip")
	data.Dnsvservername = utils.MapGetString(g, "dnsvservername")
	if val, ok := g["local"]; ok && val != nil {
		data.Local = types.BoolValue(interfaceToBool(val))
	} else {
		data.Local = types.BoolNull()
	}
	data.State = utils.MapGetString(g, "state")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Port = utils.MapGetInt64(g, "port")
	data.Nameserverstate = utils.MapGetString(g, "nameserverstate")
	data.Clmonowner = utils.MapGetInt64(g, "clmonowner")
	data.Clmonview = utils.MapGetInt64(g, "clmonview")

	// Derive the composite id in the SDK v2 format: "<name>,<type>".
	name := ""
	if !data.Ip.IsNull() && data.Ip.ValueString() != "" {
		name = data.Ip.ValueString()
	} else if !data.Dnsvservername.IsNull() && data.Dnsvservername.ValueString() != "" {
		name = data.Dnsvservername.ValueString()
	}
	data.Id = types.StringValue(name + "," + data.Type.ValueString())
}
