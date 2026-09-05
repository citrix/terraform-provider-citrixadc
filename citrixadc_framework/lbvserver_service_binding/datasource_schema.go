package lbvserver_service_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbvserverServiceBindingDataSourceModel is the data-source-specific model,
// decoupled from LbvserverServiceBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the configurable
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type LbvserverServiceBindingDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"` // Required lookup key
	Order            types.Int64  `tfsdk:"order"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Servicename      types.String `tfsdk:"servicename"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/lbvserver_service_binding.json).
	Vserverid         types.String `tfsdk:"vserverid"`
	Vsvrbindsvcip     types.String `tfsdk:"vsvrbindsvcip"`
	Preferredlocation types.String `tfsdk:"preferredlocation"`
	Servicetype       types.String `tfsdk:"servicetype"`
	Dynamicweight     types.Int64  `tfsdk:"dynamicweight"`
	Orderstr          types.String `tfsdk:"orderstr"`
	Curstate          types.String `tfsdk:"curstate"`
	Port              types.Int64  `tfsdk:"port"`
	Cookieipport      types.String `tfsdk:"cookieipport"`
	Cookiename        types.String `tfsdk:"cookiename"`
	Vsvrbindsvcport   types.Int64  `tfsdk:"vsvrbindsvcport"`
	Ipv46             types.String `tfsdk:"ipv46"`
}

func LbvserverServiceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service group.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to bind to the virtual server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the specified service.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"vserverid": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver Id.",
			},
			"vsvrbindsvcip": schema.StringAttribute{
				Computed:    true,
				Description: "used for showing the ip of bound entities.",
			},
			"preferredlocation": schema.StringAttribute{
				Computed:    true,
				Description: "Used for displaying the location of bound services.",
			},
			"servicetype": schema.StringAttribute{
				Computed:    true,
				Description: "Protocol used by the service (also called the service type).",
			},
			"dynamicweight": schema.Int64Attribute{
				Computed:    true,
				Description: "Dynamic weight.",
			},
			"orderstr": schema.StringAttribute{
				Computed:    true,
				Description: "Order in string form assigned to the service when it is bound to the lb vserver.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "Current LB vserver state.",
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: "Port number for the virtual server.",
			},
			"cookieipport": schema.StringAttribute{
				Computed:    true,
				Description: "Encryped Ip address and port of the service that is inserted into the set-cookie http header.",
			},
			"cookiename": schema.StringAttribute{
				Computed:    true,
				Description: "Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"vsvrbindsvcport": schema.Int64Attribute{
				Computed:    true,
				Description: "used for showing ports of bound entities.",
			},
			"ipv46": schema.StringAttribute{
				Computed:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
		},
	}
}

// lbvserver_service_bindingDataSourceSetAttrFromGet projects a NITRO
// lbvserver_service_binding GET response onto the data-source model.
func lbvserver_service_bindingDataSourceSetAttrFromGet(ctx context.Context, data *LbvserverServiceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbvserver_service_bindingDataSourceSetAttrFromGet Function")

	// Preserve the config-provided lookup/filter keys when the GET omits them.
	name := data.Name
	servicename := data.Servicename
	servicegroupname := data.Servicegroupname

	if v, ok := g["name"]; ok && v != nil {
		data.Name = types.StringValue(utils.AnyToString(v))
	} else {
		data.Name = name
	}
	if v, ok := g["servicename"]; ok && v != nil {
		data.Servicename = types.StringValue(utils.AnyToString(v))
	} else {
		data.Servicename = servicename
	}
	if v, ok := g["servicegroupname"]; ok && v != nil {
		data.Servicegroupname = types.StringValue(utils.AnyToString(v))
	} else {
		data.Servicegroupname = servicegroupname
	}

	data.Order = utils.MapGetInt64(g, "order")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) metadata.
	data.Vserverid = utils.MapGetString(g, "vserverid")
	data.Vsvrbindsvcip = utils.MapGetString(g, "vsvrbindsvcip")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Dynamicweight = utils.MapGetInt64(g, "dynamicweight")
	data.Orderstr = utils.MapGetString(g, "orderstr")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Port = utils.MapGetInt64(g, "port")
	data.Cookieipport = utils.MapGetString(g, "cookieipport")
	data.Cookiename = utils.MapGetString(g, "cookiename")
	data.Vsvrbindsvcport = utils.MapGetInt64(g, "vsvrbindsvcport")
	data.Ipv46 = utils.MapGetString(g, "ipv46")

	// Set the composite ID (name:<v>,servicename:<v>).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(data.Servicename.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
