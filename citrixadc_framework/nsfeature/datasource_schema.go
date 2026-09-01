package nsfeature

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsfeatureDataSourceModel is the data-source-specific model, decoupled from
// NsfeatureResourceModel.
//
// nsfeature is a global singleton (no lookup key): the appliance exposes the set
// of enabled features, and each feature is surfaced here as a Computed boolean.
// Every read-only attribute from the NITRO read-only set is already represented
// as a feature flag below, so there are no additional GET-only attributes to add.
type NsfeatureDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Wl                 types.Bool   `tfsdk:"wl"`
	Sp                 types.Bool   `tfsdk:"sp"`
	Lb                 types.Bool   `tfsdk:"lb"`
	Cs                 types.Bool   `tfsdk:"cs"`
	Cr                 types.Bool   `tfsdk:"cr"`
	Cmp                types.Bool   `tfsdk:"cmp"`
	Pq                 types.Bool   `tfsdk:"pq"`
	Ssl                types.Bool   `tfsdk:"ssl"`
	Gslb               types.Bool   `tfsdk:"gslb"`
	Hdosp              types.Bool   `tfsdk:"hdosp"`
	Cf                 types.Bool   `tfsdk:"cf"`
	Ic                 types.Bool   `tfsdk:"ic"`
	Sslvpn             types.Bool   `tfsdk:"sslvpn"`
	Aaa                types.Bool   `tfsdk:"aaa"`
	Ospf               types.Bool   `tfsdk:"ospf"`
	Rip                types.Bool   `tfsdk:"rip"`
	Bgp                types.Bool   `tfsdk:"bgp"`
	Rewrite            types.Bool   `tfsdk:"rewrite"`
	Ipv6pt             types.Bool   `tfsdk:"ipv6pt"`
	Appfw              types.Bool   `tfsdk:"appfw"`
	Responder          types.Bool   `tfsdk:"responder"`
	Htmlinjection      types.Bool   `tfsdk:"htmlinjection"`
	Push               types.Bool   `tfsdk:"push"`
	Appflow            types.Bool   `tfsdk:"appflow"`
	Cloudbridge        types.Bool   `tfsdk:"cloudbridge"`
	Isis               types.Bool   `tfsdk:"isis"`
	Ch                 types.Bool   `tfsdk:"ch"`
	Appqoe             types.Bool   `tfsdk:"appqoe"`
	Contentaccelerator types.Bool   `tfsdk:"contentaccelerator"`
	Rise               types.Bool   `tfsdk:"rise"`
	Feo                types.Bool   `tfsdk:"feo"`
	Lsn                types.Bool   `tfsdk:"lsn"`
	Rdpproxy           types.Bool   `tfsdk:"rdpproxy"`
	Rep                types.Bool   `tfsdk:"rep"`
	Urlfiltering       types.Bool   `tfsdk:"urlfiltering"`
	Videooptimization  types.Bool   `tfsdk:"videooptimization"`
	Forwardproxy       types.Bool   `tfsdk:"forwardproxy"`
	Sslinterception    types.Bool   `tfsdk:"sslinterception"`
	Adaptivetcp        types.Bool   `tfsdk:"adaptivetcp"`
	Cqa                types.Bool   `tfsdk:"cqa"`
	Ci                 types.Bool   `tfsdk:"ci"`
	Bot                types.Bool   `tfsdk:"bot"`
	Apigateway         types.Bool   `tfsdk:"apigateway"`
}

func NsfeatureDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsfeature datasource.",
			},
			"wl": schema.BoolAttribute{
				Computed:    true,
				Description: "Web Logging.",
			},
			"sp": schema.BoolAttribute{
				Computed:    true,
				Description: "Surge Protection.",
			},
			"lb": schema.BoolAttribute{
				Computed:    true,
				Description: "Load Balancing.",
			},
			"cs": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Switching.",
			},
			"cr": schema.BoolAttribute{
				Computed:    true,
				Description: "Cache Redirection.",
			},
			"cmp": schema.BoolAttribute{
				Computed:    true,
				Description: "Compression.",
			},
			"pq": schema.BoolAttribute{
				Computed:    true,
				Description: "Priority Queuing.",
			},
			"ssl": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Offloading.",
			},
			"gslb": schema.BoolAttribute{
				Computed:    true,
				Description: "Global Server Load Balancing.",
			},
			"hdosp": schema.BoolAttribute{
				Computed:    true,
				Description: "DoS Protection.",
			},
			"cf": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Filtering.",
			},
			"ic": schema.BoolAttribute{
				Computed:    true,
				Description: "Integrated Caching.",
			},
			"sslvpn": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL VPN.",
			},
			"aaa": schema.BoolAttribute{
				Computed:    true,
				Description: "AAA.",
			},
			"ospf": schema.BoolAttribute{
				Computed:    true,
				Description: "OSPF Routing.",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "RIP Routing.",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "BGP Routing.",
			},
			"rewrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Rewrite.",
			},
			"ipv6pt": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 Protocol Translation.",
			},
			"appfw": schema.BoolAttribute{
				Computed:    true,
				Description: "Application Firewall.",
			},
			"responder": schema.BoolAttribute{
				Computed:    true,
				Description: "Responder.",
			},
			"htmlinjection": schema.BoolAttribute{
				Computed:    true,
				Description: "HTML Injection.",
			},
			"push": schema.BoolAttribute{
				Computed:    true,
				Description: "Push.",
			},
			"appflow": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow.",
			},
			"cloudbridge": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge.",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "ISIS Routing.",
			},
			"ch": schema.BoolAttribute{
				Computed:    true,
				Description: "Call Home.",
			},
			"appqoe": schema.BoolAttribute{
				Computed:    true,
				Description: "AppQoE.",
			},
			"contentaccelerator": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Accelerator.",
			},
			"rise": schema.BoolAttribute{
				Computed:    true,
				Description: "RISE.",
			},
			"feo": schema.BoolAttribute{
				Computed:    true,
				Description: "Front End Optimization.",
			},
			"lsn": schema.BoolAttribute{
				Computed:    true,
				Description: "Large Scale NAT.",
			},
			"rdpproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "RDP Proxy.",
			},
			"rep": schema.BoolAttribute{
				Computed:    true,
				Description: "Reputation.",
			},
			"urlfiltering": schema.BoolAttribute{
				Computed:    true,
				Description: "URL Filtering.",
			},
			"videooptimization": schema.BoolAttribute{
				Computed:    true,
				Description: "Video Optimization.",
			},
			"forwardproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "Forward Proxy.",
			},
			"sslinterception": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Interception.",
			},
			"adaptivetcp": schema.BoolAttribute{
				Computed:    true,
				Description: "Adaptive TCP.",
			},
			"cqa": schema.BoolAttribute{
				Computed:    true,
				Description: "Connection Quality Analytics.",
			},
			"ci": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Inspection.",
			},
			"bot": schema.BoolAttribute{
				Computed:    true,
				Description: "Bot Management.",
			},
			"apigateway": schema.BoolAttribute{
				Computed:    true,
				Description: "API Gateway.",
			},
		},
	}
}

// nsfeatureDataSourceSetAttrFromGet projects the ADC's enabled-feature list onto
// the data-source model. nsfeature is a global singleton read via
// ListEnabledFeatures (a slice of feature tokens, not a keyed NITRO GET
// response), so this getter fills each Computed boolean from that list rather
// than from a map[string]interface{}. Any feature not present in the list is
// reported false.
func nsfeatureDataSourceSetAttrFromGet(ctx context.Context, data *NsfeatureDataSourceModel, enabledFeatures []string) {
	tflog.Debug(ctx, "In nsfeatureDataSourceSetAttrFromGet Function")

	enabledMap := make(map[string]bool)
	for _, feature := range enabledFeatures {
		enabledMap[feature] = true
	}

	data.Wl = types.BoolValue(enabledMap["wl"])
	data.Sp = types.BoolValue(enabledMap["sp"])
	data.Lb = types.BoolValue(enabledMap["lb"])
	data.Cs = types.BoolValue(enabledMap["cs"])
	data.Cr = types.BoolValue(enabledMap["cr"])
	data.Cmp = types.BoolValue(enabledMap["cmp"])
	data.Pq = types.BoolValue(enabledMap["pq"])
	data.Ssl = types.BoolValue(enabledMap["ssl"])
	data.Gslb = types.BoolValue(enabledMap["gslb"])
	data.Hdosp = types.BoolValue(enabledMap["hdosp"])
	data.Cf = types.BoolValue(enabledMap["cf"])
	data.Ic = types.BoolValue(enabledMap["ic"])
	data.Sslvpn = types.BoolValue(enabledMap["sslvpn"])
	data.Aaa = types.BoolValue(enabledMap["aaa"])
	data.Ospf = types.BoolValue(enabledMap["ospf"])
	data.Rip = types.BoolValue(enabledMap["rip"])
	data.Bgp = types.BoolValue(enabledMap["bgp"])
	data.Rewrite = types.BoolValue(enabledMap["rewrite"])
	data.Ipv6pt = types.BoolValue(enabledMap["ipv6pt"])
	data.Appfw = types.BoolValue(enabledMap["appfw"])
	data.Responder = types.BoolValue(enabledMap["responder"])
	data.Htmlinjection = types.BoolValue(enabledMap["htmlinjection"])
	data.Push = types.BoolValue(enabledMap["push"])
	data.Appflow = types.BoolValue(enabledMap["appflow"])
	data.Cloudbridge = types.BoolValue(enabledMap["cloudbridge"])
	data.Isis = types.BoolValue(enabledMap["isis"])
	data.Ch = types.BoolValue(enabledMap["ch"])
	data.Appqoe = types.BoolValue(enabledMap["appqoe"])
	data.Contentaccelerator = types.BoolValue(enabledMap["contentaccelerator"])
	data.Rise = types.BoolValue(enabledMap["rise"])
	data.Feo = types.BoolValue(enabledMap["feo"])
	data.Lsn = types.BoolValue(enabledMap["lsn"])
	data.Rdpproxy = types.BoolValue(enabledMap["rdpproxy"])
	data.Rep = types.BoolValue(enabledMap["rep"])
	data.Urlfiltering = types.BoolValue(enabledMap["urlfiltering"])
	data.Videooptimization = types.BoolValue(enabledMap["videooptimization"])
	data.Forwardproxy = types.BoolValue(enabledMap["forwardproxy"])
	data.Sslinterception = types.BoolValue(enabledMap["sslinterception"])
	data.Adaptivetcp = types.BoolValue(enabledMap["adaptivetcp"])
	data.Cqa = types.BoolValue(enabledMap["cqa"])
	data.Ci = types.BoolValue(enabledMap["ci"])
	data.Bot = types.BoolValue(enabledMap["bot"])
	data.Apigateway = types.BoolValue(enabledMap["apigateway"])

	// nsfeature has no unique lookup key; use a static ID.
	data.Id = types.StringValue("nsfeature-config")
}
