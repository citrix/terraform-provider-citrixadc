package nslicense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseModel describes the DATASOURCE data model. It mirrors the fields the
// NITRO `nslicense` GET returns (the licensed-feature flags), which is what a
// user queries via the datasource. This is intentionally distinct from the
// resource model (NslicenseResourceModel), because the resource is a custom
// SSH/SFTP license-upload resource, not a plain NITRO CRUD object.
type NslicenseModel struct {
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
}

func nslicenseSetAttrFromGet(ctx context.Context, data *NslicenseModel, getResponseData map[string]interface{}) *NslicenseModel {
	tflog.Debug(ctx, "In nslicenseSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["wl"]; ok && val != nil {
		data.Wl = types.BoolValue(val.(bool))
	} else {
		data.Wl = types.BoolNull()
	}
	if val, ok := getResponseData["sp"]; ok && val != nil {
		data.Sp = types.BoolValue(val.(bool))
	} else {
		data.Sp = types.BoolNull()
	}
	if val, ok := getResponseData["lb"]; ok && val != nil {
		data.Lb = types.BoolValue(val.(bool))
	} else {
		data.Lb = types.BoolNull()
	}
	if val, ok := getResponseData["cs"]; ok && val != nil {
		data.Cs = types.BoolValue(val.(bool))
	} else {
		data.Cs = types.BoolNull()
	}
	if val, ok := getResponseData["cr"]; ok && val != nil {
		data.Cr = types.BoolValue(val.(bool))
	} else {
		data.Cr = types.BoolNull()
	}
	if val, ok := getResponseData["cmp"]; ok && val != nil {
		data.Cmp = types.BoolValue(val.(bool))
	} else {
		data.Cmp = types.BoolNull()
	}
	if val, ok := getResponseData["delta"]; ok && val != nil {
		data.Delta = types.BoolValue(val.(bool))
	} else {
		data.Delta = types.BoolNull()
	}
	if val, ok := getResponseData["ssl"]; ok && val != nil {
		data.Ssl = types.BoolValue(val.(bool))
	} else {
		data.Ssl = types.BoolNull()
	}
	if val, ok := getResponseData["gslb"]; ok && val != nil {
		data.Gslb = types.BoolValue(val.(bool))
	} else {
		data.Gslb = types.BoolNull()
	}
	if val, ok := getResponseData["gslbp"]; ok && val != nil {
		data.Gslbp = types.BoolValue(val.(bool))
	} else {
		data.Gslbp = types.BoolNull()
	}
	if val, ok := getResponseData["routing"]; ok && val != nil {
		data.Routing = types.BoolValue(val.(bool))
	} else {
		data.Routing = types.BoolNull()
	}
	if val, ok := getResponseData["cf"]; ok && val != nil {
		data.Cf = types.BoolValue(val.(bool))
	} else {
		data.Cf = types.BoolNull()
	}
	if val, ok := getResponseData["contentaccelerator"]; ok && val != nil {
		data.Contentaccelerator = types.BoolValue(val.(bool))
	} else {
		data.Contentaccelerator = types.BoolNull()
	}
	if val, ok := getResponseData["ic"]; ok && val != nil {
		data.Ic = types.BoolValue(val.(bool))
	} else {
		data.Ic = types.BoolNull()
	}
	if val, ok := getResponseData["sslvpn"]; ok && val != nil {
		data.Sslvpn = types.BoolValue(val.(bool))
	} else {
		data.Sslvpn = types.BoolNull()
	}
	if val, ok := getResponseData["f_sslvpn_users"]; ok && val != nil {
		data.F_sslvpn_users = types.StringValue(utils.ToString(val))
	} else {
		data.F_sslvpn_users = types.StringNull()
	}
	if val, ok := getResponseData["f_ica_users"]; ok && val != nil {
		data.F_ica_users = types.StringValue(utils.ToString(val))
	} else {
		data.F_ica_users = types.StringNull()
	}
	if val, ok := getResponseData["aaa"]; ok && val != nil {
		data.Aaa = types.BoolValue(val.(bool))
	} else {
		data.Aaa = types.BoolNull()
	}
	if val, ok := getResponseData["ospf"]; ok && val != nil {
		data.Ospf = types.BoolValue(val.(bool))
	} else {
		data.Ospf = types.BoolNull()
	}
	if val, ok := getResponseData["rip"]; ok && val != nil {
		data.Rip = types.BoolValue(val.(bool))
	} else {
		data.Rip = types.BoolNull()
	}
	if val, ok := getResponseData["bgp"]; ok && val != nil {
		data.Bgp = types.BoolValue(val.(bool))
	} else {
		data.Bgp = types.BoolNull()
	}
	if val, ok := getResponseData["rewrite"]; ok && val != nil {
		data.Rewrite = types.BoolValue(val.(bool))
	} else {
		data.Rewrite = types.BoolNull()
	}
	if val, ok := getResponseData["ipv6pt"]; ok && val != nil {
		data.Ipv6pt = types.BoolValue(val.(bool))
	} else {
		data.Ipv6pt = types.BoolNull()
	}
	if val, ok := getResponseData["appfw"]; ok && val != nil {
		data.Appfw = types.BoolValue(val.(bool))
	} else {
		data.Appfw = types.BoolNull()
	}
	if val, ok := getResponseData["responder"]; ok && val != nil {
		data.Responder = types.BoolValue(val.(bool))
	} else {
		data.Responder = types.BoolNull()
	}
	if val, ok := getResponseData["agee"]; ok && val != nil {
		data.Agee = types.BoolValue(val.(bool))
	} else {
		data.Agee = types.BoolNull()
	}
	if val, ok := getResponseData["nsxn"]; ok && val != nil {
		data.Nsxn = types.BoolValue(val.(bool))
	} else {
		data.Nsxn = types.BoolNull()
	}
	if val, ok := getResponseData["modelid"]; ok && val != nil {
		data.Modelid = types.StringValue(utils.ToString(val))
	} else {
		data.Modelid = types.StringNull()
	}
	if val, ok := getResponseData["push"]; ok && val != nil {
		data.Push = types.BoolValue(val.(bool))
	} else {
		data.Push = types.BoolNull()
	}
	if val, ok := getResponseData["appflow"]; ok && val != nil {
		data.Appflow = types.BoolValue(val.(bool))
	} else {
		data.Appflow = types.BoolNull()
	}
	if val, ok := getResponseData["cloudbridge"]; ok && val != nil {
		data.Cloudbridge = types.BoolValue(val.(bool))
	} else {
		data.Cloudbridge = types.BoolNull()
	}
	if val, ok := getResponseData["cloudbridgeappliance"]; ok && val != nil {
		data.Cloudbridgeappliance = types.BoolValue(val.(bool))
	} else {
		data.Cloudbridgeappliance = types.BoolNull()
	}
	if val, ok := getResponseData["cloudextenderappliance"]; ok && val != nil {
		data.Cloudextenderappliance = types.BoolValue(val.(bool))
	} else {
		data.Cloudextenderappliance = types.BoolNull()
	}
	if val, ok := getResponseData["isis"]; ok && val != nil {
		data.Isis = types.BoolValue(val.(bool))
	} else {
		data.Isis = types.BoolNull()
	}
	if val, ok := getResponseData["cluster"]; ok && val != nil {
		data.Cluster = types.BoolValue(val.(bool))
	} else {
		data.Cluster = types.BoolNull()
	}
	if val, ok := getResponseData["ch"]; ok && val != nil {
		data.Ch = types.BoolValue(val.(bool))
	} else {
		data.Ch = types.BoolNull()
	}
	if val, ok := getResponseData["appqoe"]; ok && val != nil {
		data.Appqoe = types.BoolValue(val.(bool))
	} else {
		data.Appqoe = types.BoolNull()
	}
	if val, ok := getResponseData["appflowica"]; ok && val != nil {
		data.Appflowica = types.BoolValue(val.(bool))
	} else {
		data.Appflowica = types.BoolNull()
	}
	if val, ok := getResponseData["isstandardlic"]; ok && val != nil {
		data.Isstandardlic = types.BoolValue(val.(bool))
	} else {
		data.Isstandardlic = types.BoolNull()
	}
	if val, ok := getResponseData["isenterpriselic"]; ok && val != nil {
		data.Isenterpriselic = types.BoolValue(val.(bool))
	} else {
		data.Isenterpriselic = types.BoolNull()
	}
	if val, ok := getResponseData["isplatinumlic"]; ok && val != nil {
		data.Isplatinumlic = types.BoolValue(val.(bool))
	} else {
		data.Isplatinumlic = types.BoolNull()
	}
	if val, ok := getResponseData["issgwylic"]; ok && val != nil {
		data.Issgwylic = types.BoolValue(val.(bool))
	} else {
		data.Issgwylic = types.BoolNull()
	}
	if val, ok := getResponseData["isswglic"]; ok && val != nil {
		data.Isswglic = types.BoolValue(val.(bool))
	} else {
		data.Isswglic = types.BoolNull()
	}
	if val, ok := getResponseData["feo"]; ok && val != nil {
		data.Feo = types.BoolValue(val.(bool))
	} else {
		data.Feo = types.BoolNull()
	}
	if val, ok := getResponseData["lsn"]; ok && val != nil {
		data.Lsn = types.BoolValue(val.(bool))
	} else {
		data.Lsn = types.BoolNull()
	}
	if val, ok := getResponseData["licensingmode"]; ok && val != nil {
		data.Licensingmode = types.StringValue(utils.ToString(val))
	} else {
		data.Licensingmode = types.StringNull()
	}
	if val, ok := getResponseData["rdpproxy"]; ok && val != nil {
		data.Rdpproxy = types.BoolValue(val.(bool))
	} else {
		data.Rdpproxy = types.BoolNull()
	}
	if val, ok := getResponseData["rep"]; ok && val != nil {
		data.Rep = types.BoolValue(val.(bool))
	} else {
		data.Rep = types.BoolNull()
	}
	if val, ok := getResponseData["urlfiltering"]; ok && val != nil {
		data.Urlfiltering = types.BoolValue(val.(bool))
	} else {
		data.Urlfiltering = types.BoolNull()
	}
	if val, ok := getResponseData["videooptimization"]; ok && val != nil {
		data.Videooptimization = types.BoolValue(val.(bool))
	} else {
		data.Videooptimization = types.BoolNull()
	}
	if val, ok := getResponseData["forwardproxy"]; ok && val != nil {
		data.Forwardproxy = types.BoolValue(val.(bool))
	} else {
		data.Forwardproxy = types.BoolNull()
	}
	if val, ok := getResponseData["sslinterception"]; ok && val != nil {
		data.Sslinterception = types.BoolValue(val.(bool))
	} else {
		data.Sslinterception = types.BoolNull()
	}
	if val, ok := getResponseData["remotecontentinspection"]; ok && val != nil {
		data.Remotecontentinspection = types.BoolValue(val.(bool))
	} else {
		data.Remotecontentinspection = types.BoolNull()
	}
	if val, ok := getResponseData["adaptivetcp"]; ok && val != nil {
		data.Adaptivetcp = types.BoolValue(val.(bool))
	} else {
		data.Adaptivetcp = types.BoolNull()
	}
	if val, ok := getResponseData["cqa"]; ok && val != nil {
		data.Cqa = types.BoolValue(val.(bool))
	} else {
		data.Cqa = types.BoolNull()
	}
	if val, ok := getResponseData["bot"]; ok && val != nil {
		data.Bot = types.BoolValue(val.(bool))
	} else {
		data.Bot = types.BoolNull()
	}
	if val, ok := getResponseData["apigateway"]; ok && val != nil {
		data.Apigateway = types.BoolValue(val.(bool))
	} else {
		data.Apigateway = types.BoolNull()
	}

	// Set ID for the datasource
	data.Id = types.StringValue("nslicense")

	return data
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
		},
	}
}
