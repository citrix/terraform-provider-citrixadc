package gslbvserver_gslbservice_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbvserverGslbserviceBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read
// surface, so it can expose the FULL GET projection: the read/write attributes
// (as Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type GslbvserverGslbserviceBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domainname  types.String `tfsdk:"domainname"`
	Name        types.String `tfsdk:"name"`
	Order       types.Int64  `tfsdk:"order"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbvserver_gslbservice_binding.json). Never settable;
	// populated from GET; null when the appliance omits them.
	Dynamicconfwt      types.Int64  `tfsdk:"dynamicconfwt"`
	Cumulativeweight   types.Int64  `tfsdk:"cumulativeweight"`
	Sitepersistcookie  types.String `tfsdk:"sitepersistcookie"`
	Orderstr           types.String `tfsdk:"orderstr"`
	Gslbthreshold      types.Int64  `tfsdk:"gslbthreshold"`
	Preferredlocation  types.String `tfsdk:"preferredlocation"`
	Svcsitepersistence types.String `tfsdk:"svcsitepersistence"`
	Gslbboundsvctype   types.String `tfsdk:"gslbboundsvctype"`
	Ipaddress          types.String `tfsdk:"ipaddress"`
	Iscname            types.String `tfsdk:"iscname"`
	Thresholdvalue     types.Int64  `tfsdk:"thresholdvalue"`
	Port               types.Int64  `tfsdk:"port"`
	Curstate           types.String `tfsdk:"curstate"`
	Svreffgslbstate    types.String `tfsdk:"svreffgslbstate"`
	Cnameentry         types.String `tfsdk:"cnameentry"`
}

func GslbvserverGslbserviceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service for which to change the weight.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"dynamicconfwt": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight obtained by the virtue of bound service count or weight.",
			},
			"cumulativeweight": schema.Int64Attribute{
				Computed:    true,
				Description: "Cumulative weight is the weight of the GSLB service considering both its configured weight and dynamic weight.",
			},
			"sitepersistcookie": schema.StringAttribute{
				Computed:    true,
				Description: "Cookie displayed for site persistence in a cluster setup.",
			},
			"orderstr": schema.StringAttribute{
				Computed:    true,
				Description: "Order number in string form assigned to the service when it is bound to the lb vserver.",
			},
			"gslbthreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates whether the GSLB service has reached its threshold.",
			},
			"preferredlocation": schema.StringAttribute{
				Computed:    true,
				Description: "The target site to be returned in the DNS response when a policy is successfully evaluated against the incoming DNS request.",
			},
			"svcsitepersistence": schema.StringAttribute{
				Computed:    true,
				Description: "Type of Site Persistence set on the bound service. Possible values = ConnectionProxy, HTTPRedirect, NONE.",
			},
			"gslbboundsvctype": schema.StringAttribute{
				Computed:    true,
				Description: "Protocol used by services bound to the GSLB virtual server.",
			},
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "IP address of the bound GSLB service.",
			},
			"iscname": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the cname feature is set on the vserver. Possible values = ENABLED, DISABLED.",
			},
			"thresholdvalue": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates whether the threshold has been exceeded for this service participating in CUSTOMLB.",
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: "Port number of the bound GSLB service.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the GSLB vserver.",
			},
			"svreffgslbstate": schema.StringAttribute{
				Computed:    true,
				Description: "Effective state of the GSLB service.",
			},
			"cnameentry": schema.StringAttribute{
				Computed:    true,
				Description: "The cname of the GSLB service.",
			},
		},
	}
}

// gslbvserver_gslbservice_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers. The composite ID is
// built exactly as the resource Create emits it.
func gslbvserver_gslbservice_bindingDataSourceSetAttrFromGet(ctx context.Context, data *GslbvserverGslbserviceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbvserver_gslbservice_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Domainname = utils.MapGetString(g, "domainname")
	data.Name = utils.MapGetString(g, "name")
	data.Order = utils.MapGetInt64(g, "order")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) attributes.
	data.Dynamicconfwt = utils.MapGetInt64(g, "dynamicconfwt")
	data.Cumulativeweight = utils.MapGetInt64(g, "cumulativeweight")
	data.Sitepersistcookie = utils.MapGetString(g, "sitepersistcookie")
	data.Orderstr = utils.MapGetString(g, "orderstr")
	data.Gslbthreshold = utils.MapGetInt64(g, "gslbthreshold")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")
	data.Svcsitepersistence = utils.MapGetString(g, "svcsitepersistence")
	data.Gslbboundsvctype = utils.MapGetString(g, "gslbboundsvctype")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Iscname = utils.MapGetString(g, "iscname")
	data.Thresholdvalue = utils.MapGetInt64(g, "thresholdvalue")
	data.Port = utils.MapGetInt64(g, "port")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Svreffgslbstate = utils.MapGetString(g, "svreffgslbstate")
	data.Cnameentry = utils.MapGetString(g, "cnameentry")

	// Composite ID: name,servicename (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
