package nslicense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseDataSourceModel describes the DATASOURCE data model. It mirrors the
// fields the NITRO `nslicense` GET returns (the licensed-feature flags and the
// read-only license metadata), which is what a user queries via the datasource.
// This is intentionally distinct from the resource model
// (NslicenseResourceModel), because the resource is a custom SSH/SFTP
// license-upload resource, not a plain NITRO CRUD object.
type NslicenseDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Wl                      types.Bool   `tfsdk:"wl"`
	Sp                      types.Bool   `tfsdk:"sp"`
	Lb                      types.Bool   `tfsdk:"lb"`
	Cs                      types.Bool   `tfsdk:"cs"`
	Cr                      types.Bool   `tfsdk:"cr"`
	Cmp                     types.Bool   `tfsdk:"cmp"`
	Delta                   types.Bool   `tfsdk:"delta"`
	Ssl                     types.Bool   `tfsdk:"ssl"`
	Gslb                    types.Bool   `tfsdk:"gslb"`
	Gslbp                   types.Bool   `tfsdk:"gslbp"`
	Routing                 types.Bool   `tfsdk:"routing"`
	Cf                      types.Bool   `tfsdk:"cf"`
	Contentaccelerator      types.Bool   `tfsdk:"contentaccelerator"`
	Ic                      types.Bool   `tfsdk:"ic"`
	Sslvpn                  types.Bool   `tfsdk:"sslvpn"`
	F_sslvpn_users          types.String `tfsdk:"f_sslvpn_users"`
	F_ica_users             types.String `tfsdk:"f_ica_users"`
	Aaa                     types.Bool   `tfsdk:"aaa"`
	Ospf                    types.Bool   `tfsdk:"ospf"`
	Rip                     types.Bool   `tfsdk:"rip"`
	Bgp                     types.Bool   `tfsdk:"bgp"`
	Rewrite                 types.Bool   `tfsdk:"rewrite"`
	Ipv6pt                  types.Bool   `tfsdk:"ipv6pt"`
	Appfw                   types.Bool   `tfsdk:"appfw"`
	Responder               types.Bool   `tfsdk:"responder"`
	Agee                    types.Bool   `tfsdk:"agee"`
	Nsxn                    types.Bool   `tfsdk:"nsxn"`
	Modelid                 types.String `tfsdk:"modelid"`
	Push                    types.Bool   `tfsdk:"push"`
	Appflow                 types.Bool   `tfsdk:"appflow"`
	Cloudbridge             types.Bool   `tfsdk:"cloudbridge"`
	Cloudbridgeappliance    types.Bool   `tfsdk:"cloudbridgeappliance"`
	Cloudextenderappliance  types.Bool   `tfsdk:"cloudextenderappliance"`
	Isis                    types.Bool   `tfsdk:"isis"`
	Cluster                 types.Bool   `tfsdk:"cluster"`
	Ch                      types.Bool   `tfsdk:"ch"`
	Appqoe                  types.Bool   `tfsdk:"appqoe"`
	Appflowica              types.Bool   `tfsdk:"appflowica"`
	Isstandardlic           types.Bool   `tfsdk:"isstandardlic"`
	Isenterpriselic         types.Bool   `tfsdk:"isenterpriselic"`
	Isplatinumlic           types.Bool   `tfsdk:"isplatinumlic"`
	Issgwylic               types.Bool   `tfsdk:"issgwylic"`
	Isswglic                types.Bool   `tfsdk:"isswglic"`
	Feo                     types.Bool   `tfsdk:"feo"`
	Lsn                     types.Bool   `tfsdk:"lsn"`
	Licensingmode           types.String `tfsdk:"licensingmode"`
	Rdpproxy                types.Bool   `tfsdk:"rdpproxy"`
	Rep                     types.Bool   `tfsdk:"rep"`
	Urlfiltering            types.Bool   `tfsdk:"urlfiltering"`
	Videooptimization       types.Bool   `tfsdk:"videooptimization"`
	Forwardproxy            types.Bool   `tfsdk:"forwardproxy"`
	Sslinterception         types.Bool   `tfsdk:"sslinterception"`
	Remotecontentinspection types.Bool   `tfsdk:"remotecontentinspection"`
	Adaptivetcp             types.Bool   `tfsdk:"adaptivetcp"`
	Cqa                     types.Bool   `tfsdk:"cqa"`
	Bot                     types.Bool   `tfsdk:"bot"`
	Apigateway              types.Bool   `tfsdk:"apigateway"`

	// Read-only (GET-only) license metadata from the NITRO doc read-only set
	// (zion73x_readonly/nslicense.json). Never settable; populated from GET.
	Cloudsubscriptionimage types.String `tfsdk:"cloudsubscriptionimage"`
	Daystoexpiration       types.Int64  `tfsdk:"daystoexpiration"`
	Daystolasenforcement   types.Int64  `tfsdk:"daystolasenforcement"`
}

// nslicenseDataSourceSetAttrFromGet projects a NITRO nslicense GET response onto
// the data-source model using the shared utils.MapGet* helpers. Attributes the
// GET omits are left Null.
func nslicenseDataSourceSetAttrFromGet(ctx context.Context, data *NslicenseDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nslicenseDataSourceSetAttrFromGet Function")

	data.Wl = utils.MapGetBool(g, "wl")
	data.Sp = utils.MapGetBool(g, "sp")
	data.Lb = utils.MapGetBool(g, "lb")
	data.Cs = utils.MapGetBool(g, "cs")
	data.Cr = utils.MapGetBool(g, "cr")
	data.Cmp = utils.MapGetBool(g, "cmp")
	data.Delta = utils.MapGetBool(g, "delta")
	data.Ssl = utils.MapGetBool(g, "ssl")
	data.Gslb = utils.MapGetBool(g, "gslb")
	data.Gslbp = utils.MapGetBool(g, "gslbp")
	data.Routing = utils.MapGetBool(g, "routing")
	data.Cf = utils.MapGetBool(g, "cf")
	data.Contentaccelerator = utils.MapGetBool(g, "contentaccelerator")
	data.Ic = utils.MapGetBool(g, "ic")
	data.Sslvpn = utils.MapGetBool(g, "sslvpn")
	data.F_sslvpn_users = utils.MapGetString(g, "f_sslvpn_users")
	data.F_ica_users = utils.MapGetString(g, "f_ica_users")
	data.Aaa = utils.MapGetBool(g, "aaa")
	data.Ospf = utils.MapGetBool(g, "ospf")
	data.Rip = utils.MapGetBool(g, "rip")
	data.Bgp = utils.MapGetBool(g, "bgp")
	data.Rewrite = utils.MapGetBool(g, "rewrite")
	data.Ipv6pt = utils.MapGetBool(g, "ipv6pt")
	data.Appfw = utils.MapGetBool(g, "appfw")
	data.Responder = utils.MapGetBool(g, "responder")
	data.Agee = utils.MapGetBool(g, "agee")
	data.Nsxn = utils.MapGetBool(g, "nsxn")
	data.Modelid = utils.MapGetString(g, "modelid")
	data.Push = utils.MapGetBool(g, "push")
	data.Appflow = utils.MapGetBool(g, "appflow")
	data.Cloudbridge = utils.MapGetBool(g, "cloudbridge")
	data.Cloudbridgeappliance = utils.MapGetBool(g, "cloudbridgeappliance")
	data.Cloudextenderappliance = utils.MapGetBool(g, "cloudextenderappliance")
	data.Isis = utils.MapGetBool(g, "isis")
	data.Cluster = utils.MapGetBool(g, "cluster")
	data.Ch = utils.MapGetBool(g, "ch")
	data.Appqoe = utils.MapGetBool(g, "appqoe")
	data.Appflowica = utils.MapGetBool(g, "appflowica")
	data.Isstandardlic = utils.MapGetBool(g, "isstandardlic")
	data.Isenterpriselic = utils.MapGetBool(g, "isenterpriselic")
	data.Isplatinumlic = utils.MapGetBool(g, "isplatinumlic")
	data.Issgwylic = utils.MapGetBool(g, "issgwylic")
	data.Isswglic = utils.MapGetBool(g, "isswglic")
	data.Feo = utils.MapGetBool(g, "feo")
	data.Lsn = utils.MapGetBool(g, "lsn")
	data.Licensingmode = utils.MapGetString(g, "licensingmode")
	data.Rdpproxy = utils.MapGetBool(g, "rdpproxy")
	data.Rep = utils.MapGetBool(g, "rep")
	data.Urlfiltering = utils.MapGetBool(g, "urlfiltering")
	data.Videooptimization = utils.MapGetBool(g, "videooptimization")
	data.Forwardproxy = utils.MapGetBool(g, "forwardproxy")
	data.Sslinterception = utils.MapGetBool(g, "sslinterception")
	data.Remotecontentinspection = utils.MapGetBool(g, "remotecontentinspection")
	data.Adaptivetcp = utils.MapGetBool(g, "adaptivetcp")
	data.Cqa = utils.MapGetBool(g, "cqa")
	data.Bot = utils.MapGetBool(g, "bot")
	data.Apigateway = utils.MapGetBool(g, "apigateway")

	// Read-only license metadata.
	data.Cloudsubscriptionimage = utils.MapGetString(g, "cloudsubscriptionimage")
	data.Daystoexpiration = utils.MapGetInt64(g, "daystoexpiration")
	data.Daystolasenforcement = utils.MapGetInt64(g, "daystolasenforcement")

	// nslicense is a singleton; use a fixed ID.
	data.Id = types.StringValue("nslicense")
}

func NslicenseDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for querying Citrix ADC license information",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicense datasource (always 'nslicense')",
			},
			"wl": schema.BoolAttribute{
				Computed:    true,
				Description: "Web Logging feature is licensed",
			},
			"sp": schema.BoolAttribute{
				Computed:    true,
				Description: "Surge Protection feature is licensed",
			},
			"lb": schema.BoolAttribute{
				Computed:    true,
				Description: "Load Balancing feature is licensed",
			},
			"cs": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Switching feature is licensed",
			},
			"cr": schema.BoolAttribute{
				Computed:    true,
				Description: "Cache Redirection feature is licensed",
			},
			"cmp": schema.BoolAttribute{
				Computed:    true,
				Description: "Compression feature is licensed",
			},
			"delta": schema.BoolAttribute{
				Computed:    true,
				Description: "Delta Compression feature is licensed",
			},
			"ssl": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Offloading feature is licensed",
			},
			"gslb": schema.BoolAttribute{
				Computed:    true,
				Description: "Global Server Load Balancing feature is licensed",
			},
			"gslbp": schema.BoolAttribute{
				Computed:    true,
				Description: "GSLB Proximity feature is licensed",
			},
			"routing": schema.BoolAttribute{
				Computed:    true,
				Description: "Routing feature is licensed",
			},
			"cf": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Filtering feature is licensed",
			},
			"contentaccelerator": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Accelerator feature is licensed",
			},
			"ic": schema.BoolAttribute{
				Computed:    true,
				Description: "Integrated Caching feature is licensed",
			},
			"sslvpn": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL VPN feature is licensed",
			},
			"f_sslvpn_users": schema.StringAttribute{
				Computed:    true,
				Description: "Number of SSL VPN users licensed",
			},
			"f_ica_users": schema.StringAttribute{
				Computed:    true,
				Description: "Number of ICA users licensed",
			},
			"aaa": schema.BoolAttribute{
				Computed:    true,
				Description: "AAA (Authentication, Authorization, Accounting) feature is licensed",
			},
			"ospf": schema.BoolAttribute{
				Computed:    true,
				Description: "OSPF routing feature is licensed",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "RIP routing feature is licensed",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "BGP routing feature is licensed",
			},
			"rewrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Rewrite feature is licensed",
			},
			"ipv6pt": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 Protocol Translation feature is licensed",
			},
			"appfw": schema.BoolAttribute{
				Computed:    true,
				Description: "Application Firewall feature is licensed",
			},
			"responder": schema.BoolAttribute{
				Computed:    true,
				Description: "Responder feature is licensed",
			},
			"agee": schema.BoolAttribute{
				Computed:    true,
				Description: "AGEE feature is licensed",
			},
			"nsxn": schema.BoolAttribute{
				Computed:    true,
				Description: "NetScaler XN feature is licensed",
			},
			"modelid": schema.StringAttribute{
				Computed:    true,
				Description: "Model ID of the appliance",
			},
			"push": schema.BoolAttribute{
				Computed:    true,
				Description: "Push feature is licensed",
			},
			"appflow": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow feature is licensed",
			},
			"cloudbridge": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge feature is licensed",
			},
			"cloudbridgeappliance": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge Appliance feature is licensed",
			},
			"cloudextenderappliance": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudExtender Appliance feature is licensed",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "ISIS routing feature is licensed",
			},
			"cluster": schema.BoolAttribute{
				Computed:    true,
				Description: "Cluster feature is licensed",
			},
			"ch": schema.BoolAttribute{
				Computed:    true,
				Description: "Call Home feature is licensed",
			},
			"appqoe": schema.BoolAttribute{
				Computed:    true,
				Description: "AppQoE feature is licensed",
			},
			"appflowica": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow for ICA feature is licensed",
			},
			"isstandardlic": schema.BoolAttribute{
				Computed:    true,
				Description: "Standard license is applied",
			},
			"isenterpriselic": schema.BoolAttribute{
				Computed:    true,
				Description: "Enterprise license is applied",
			},
			"isplatinumlic": schema.BoolAttribute{
				Computed:    true,
				Description: "Platinum license is applied",
			},
			"issgwylic": schema.BoolAttribute{
				Computed:    true,
				Description: "Secure Gateway license is applied",
			},
			"isswglic": schema.BoolAttribute{
				Computed:    true,
				Description: "SWG license is applied",
			},
			"feo": schema.BoolAttribute{
				Computed:    true,
				Description: "Front End Optimization feature is licensed",
			},
			"lsn": schema.BoolAttribute{
				Computed:    true,
				Description: "Large Scale NAT feature is licensed",
			},
			"licensingmode": schema.StringAttribute{
				Computed:    true,
				Description: "Licensing mode (e.g., EXPRESS, POOLED)",
			},
			"rdpproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "RDP Proxy feature is licensed",
			},
			"rep": schema.BoolAttribute{
				Computed:    true,
				Description: "Reputation feature is licensed",
			},
			"urlfiltering": schema.BoolAttribute{
				Computed:    true,
				Description: "URL Filtering feature is licensed",
			},
			"videooptimization": schema.BoolAttribute{
				Computed:    true,
				Description: "Video Optimization feature is licensed",
			},
			"forwardproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "Forward Proxy feature is licensed",
			},
			"sslinterception": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Interception feature is licensed",
			},
			"remotecontentinspection": schema.BoolAttribute{
				Computed:    true,
				Description: "Remote Content Inspection feature is licensed",
			},
			"adaptivetcp": schema.BoolAttribute{
				Computed:    true,
				Description: "Adaptive TCP feature is licensed",
			},
			"cqa": schema.BoolAttribute{
				Computed:    true,
				Description: "CQA feature is licensed",
			},
			"bot": schema.BoolAttribute{
				Computed:    true,
				Description: "Bot Management feature is licensed",
			},
			"apigateway": schema.BoolAttribute{
				Computed:    true,
				Description: "API Gateway feature is licensed",
			},

			// Read-only (GET-only) license metadata surfaced by the data source.
			"cloudsubscriptionimage": schema.StringAttribute{
				Computed:    true,
				Description: "Cloud Subscription Image (YES/NO). Null when the appliance omits it.",
			},
			"daystoexpiration": schema.Int64Attribute{
				Computed:    true,
				Description: "Days to license expiration. Null when the appliance omits it.",
			},
			"daystolasenforcement": schema.Int64Attribute{
				Computed:    true,
				Description: "Days to expiration for LAS enforcement. Null when the appliance omits it.",
			},
		},
	}
}
