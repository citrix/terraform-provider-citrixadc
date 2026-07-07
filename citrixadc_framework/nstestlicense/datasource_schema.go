package nstestlicense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstestlicenseDataSourceModel describes the datasource data model. It exposes
// the synthetic id and the read-only get(all) attributes.
//
// The NITRO key "_nextgenapiresource" cannot be used as a tfsdk tag verbatim
// (leading underscore), so its tag is "nextgenapiresource" while the exact NITRO
// key is read in SetAttrFromGet.
type NstestlicenseDataSourceModel struct {
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
	Fsslvpnusers            types.Int64  `tfsdk:"f_sslvpn_users"`
	Ficausers               types.Int64  `tfsdk:"f_ica_users"`
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
	Modelid                 types.Int64  `tfsdk:"modelid"`
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
	Daystoexpiration        types.Int64  `tfsdk:"daystoexpiration"`
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
	Nextgenapiresource      types.String `tfsdk:"nextgenapiresource"`
}

func boolAttr(desc string) schema.BoolAttribute {
	return schema.BoolAttribute{Computed: true, Description: desc}
}

func NstestlicenseDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstestlicense datasource.",
			},
			"wl":                      boolAttr("Web Logging."),
			"sp":                      boolAttr("Surge Protection."),
			"lb":                      boolAttr("Load Balancing."),
			"cs":                      boolAttr("Content Switching."),
			"cr":                      boolAttr("Cache Redirect."),
			"cmp":                     boolAttr("Compression."),
			"delta":                   boolAttr("Delta Compression."),
			"ssl":                     boolAttr("Secure Sockets Layer."),
			"gslb":                    boolAttr("Global Server Load Balancing."),
			"gslbp":                   boolAttr("GSLB Proximity."),
			"routing":                 boolAttr("Routing."),
			"cf":                      boolAttr("Content Filter."),
			"contentaccelerator":      boolAttr("Transparent Integrated Caching."),
			"ic":                      boolAttr("Integrated Caching."),
			"sslvpn":                  boolAttr("SSL VPN."),
			"f_sslvpn_users":          schema.Int64Attribute{Computed: true, Description: "Number of licensed users allowed by this license."},
			"f_ica_users":             schema.Int64Attribute{Computed: true, Description: "Number of licensed users allowed by ICAONLY license."},
			"aaa":                     boolAttr("AAA."),
			"ospf":                    boolAttr("OSPF Routing."),
			"rip":                     boolAttr("RIP Routing."),
			"bgp":                     boolAttr("BGP Routing."),
			"rewrite":                 boolAttr("Rewrite."),
			"ipv6pt":                  boolAttr("IPv6 protocol translation."),
			"appfw":                   boolAttr("Application Firewall."),
			"responder":               boolAttr("Responder."),
			"agee":                    boolAttr("Access Gateway Enterprise Edition."),
			"nsxn":                    boolAttr("NSXN."),
			"modelid":                 schema.Int64Attribute{Computed: true, Description: "Model Number ID."},
			"push":                    boolAttr("Citrix ADC Push."),
			"appflow":                 boolAttr("AppFlow."),
			"cloudbridge":             boolAttr("CloudBridge."),
			"cloudbridgeappliance":    boolAttr("CloudBridge Appliance."),
			"cloudextenderappliance":  boolAttr("Cloud Extender Appliance."),
			"isis":                    boolAttr("ISIS Routing."),
			"cluster":                 boolAttr("Clustering."),
			"ch":                      boolAttr("Call Home."),
			"appqoe":                  boolAttr("AppQoS."),
			"appflowica":              boolAttr("Appflow for ICA."),
			"isstandardlic":           boolAttr("Standard License."),
			"isenterpriselic":         boolAttr("Enterprise License."),
			"isplatinumlic":           boolAttr("Platinum License."),
			"issgwylic":               boolAttr("Simple Gateway License."),
			"isswglic":                boolAttr("Secure Web Gateway License."),
			"feo":                     boolAttr("Front End Optimization."),
			"lsn":                     boolAttr("Large Scale NAT."),
			"licensingmode":           schema.StringAttribute{Computed: true, Description: "Pooled Licensed. Default value: Local."},
			"daystoexpiration":        schema.Int64Attribute{Computed: true, Description: "Days to expire."},
			"rdpproxy":                boolAttr("RDPPROXY."),
			"rep":                     boolAttr("Reputation Services."),
			"urlfiltering":            boolAttr("URL Filtering."),
			"videooptimization":       boolAttr("Video Optimization."),
			"forwardproxy":            boolAttr("Forward Proxy."),
			"sslinterception":         boolAttr("SSL Interception."),
			"remotecontentinspection": boolAttr("Remote Content Inspection."),
			"adaptivetcp":             boolAttr("Adaptive TCP."),
			"cqa":                     boolAttr("Connection Quality Analytics."),
			"bot":                     boolAttr("Bot Management."),
			"apigateway":              boolAttr("API Gateway."),
			"nextgenapiresource":      schema.StringAttribute{Computed: true, Description: "Read-only attribute (_nextgenapiresource)."},
		},
	}
}

// toBoolValue converts a NITRO get(all) value into a types.Bool, tolerating both
// native bool and string ("true"/"false") encodings.
func toBoolValue(val interface{}) types.Bool {
	switch v := val.(type) {
	case bool:
		return types.BoolValue(v)
	case string:
		return types.BoolValue(v == "true" || v == "True" || v == "TRUE")
	default:
		return types.BoolNull()
	}
}

func setBool(getResponseData map[string]interface{}, key string) types.Bool {
	if val, ok := getResponseData[key]; ok && val != nil {
		return toBoolValue(val)
	}
	return types.BoolNull()
}

func setInt(getResponseData map[string]interface{}, key string) types.Int64 {
	if val, ok := getResponseData[key]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			return types.Int64Value(intVal)
		}
	}
	return types.Int64Null()
}

func setString(getResponseData map[string]interface{}, key string) types.String {
	if val, ok := getResponseData[key]; ok && val != nil {
		if s, isStr := val.(string); isStr {
			return types.StringValue(s)
		}
	}
	return types.StringNull()
}

// nstestlicenseSetAttrFromGet maps the get(all) response (exact NITRO keys) into
// the datasource model.
func nstestlicenseSetAttrFromGet(ctx context.Context, data *NstestlicenseDataSourceModel, getResponseData map[string]interface{}) *NstestlicenseDataSourceModel {
	tflog.Debug(ctx, "In nstestlicenseSetAttrFromGet Function")

	data.Wl = setBool(getResponseData, "wl")
	data.Sp = setBool(getResponseData, "sp")
	data.Lb = setBool(getResponseData, "lb")
	data.Cs = setBool(getResponseData, "cs")
	data.Cr = setBool(getResponseData, "cr")
	data.Cmp = setBool(getResponseData, "cmp")
	data.Delta = setBool(getResponseData, "delta")
	data.Ssl = setBool(getResponseData, "ssl")
	data.Gslb = setBool(getResponseData, "gslb")
	data.Gslbp = setBool(getResponseData, "gslbp")
	data.Routing = setBool(getResponseData, "routing")
	data.Cf = setBool(getResponseData, "cf")
	data.Contentaccelerator = setBool(getResponseData, "contentaccelerator")
	data.Ic = setBool(getResponseData, "ic")
	data.Sslvpn = setBool(getResponseData, "sslvpn")
	data.Fsslvpnusers = setInt(getResponseData, "f_sslvpn_users")
	data.Ficausers = setInt(getResponseData, "f_ica_users")
	data.Aaa = setBool(getResponseData, "aaa")
	data.Ospf = setBool(getResponseData, "ospf")
	data.Rip = setBool(getResponseData, "rip")
	data.Bgp = setBool(getResponseData, "bgp")
	data.Rewrite = setBool(getResponseData, "rewrite")
	data.Ipv6pt = setBool(getResponseData, "ipv6pt")
	data.Appfw = setBool(getResponseData, "appfw")
	data.Responder = setBool(getResponseData, "responder")
	data.Agee = setBool(getResponseData, "agee")
	data.Nsxn = setBool(getResponseData, "nsxn")
	data.Modelid = setInt(getResponseData, "modelid")
	data.Push = setBool(getResponseData, "push")
	data.Appflow = setBool(getResponseData, "appflow")
	data.Cloudbridge = setBool(getResponseData, "cloudbridge")
	data.Cloudbridgeappliance = setBool(getResponseData, "cloudbridgeappliance")
	data.Cloudextenderappliance = setBool(getResponseData, "cloudextenderappliance")
	data.Isis = setBool(getResponseData, "isis")
	data.Cluster = setBool(getResponseData, "cluster")
	data.Ch = setBool(getResponseData, "ch")
	data.Appqoe = setBool(getResponseData, "appqoe")
	data.Appflowica = setBool(getResponseData, "appflowica")
	data.Isstandardlic = setBool(getResponseData, "isstandardlic")
	data.Isenterpriselic = setBool(getResponseData, "isenterpriselic")
	data.Isplatinumlic = setBool(getResponseData, "isplatinumlic")
	data.Issgwylic = setBool(getResponseData, "issgwylic")
	data.Isswglic = setBool(getResponseData, "isswglic")
	data.Feo = setBool(getResponseData, "feo")
	data.Lsn = setBool(getResponseData, "lsn")
	data.Licensingmode = setString(getResponseData, "licensingmode")
	data.Daystoexpiration = setInt(getResponseData, "daystoexpiration")
	data.Rdpproxy = setBool(getResponseData, "rdpproxy")
	data.Rep = setBool(getResponseData, "rep")
	data.Urlfiltering = setBool(getResponseData, "urlfiltering")
	data.Videooptimization = setBool(getResponseData, "videooptimization")
	data.Forwardproxy = setBool(getResponseData, "forwardproxy")
	data.Sslinterception = setBool(getResponseData, "sslinterception")
	data.Remotecontentinspection = setBool(getResponseData, "remotecontentinspection")
	data.Adaptivetcp = setBool(getResponseData, "adaptivetcp")
	data.Cqa = setBool(getResponseData, "cqa")
	data.Bot = setBool(getResponseData, "bot")
	data.Apigateway = setBool(getResponseData, "apigateway")
	data.Nextgenapiresource = setString(getResponseData, "_nextgenapiresource")

	data.Id = types.StringValue("nstestlicense-config")

	return data
}
