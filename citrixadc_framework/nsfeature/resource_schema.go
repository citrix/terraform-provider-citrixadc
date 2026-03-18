package nsfeature

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsfeatureResourceModel describes the resource data model.
type NsfeatureResourceModel struct {
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

var featureList = []string{
	"wl",
	"sp",
	"lb",
	"cs",
	"cr",
	"cmp",
	"pq",
	"ssl",
	"gslb",
	"hdosp",
	"cf",
	"ic",
	"sslvpn",
	"aaa",
	"ospf",
	"rip",
	"bgp",
	"rewrite",
	"ipv6pt",
	"appfw",
	"responder",
	"htmlinjection",
	"push",
	"appflow",
	"cloudbridge",
	"isis",
	"ch",
	"appqoe",
	"contentaccelerator",
	"rise",
	"feo",
	"lsn",
	"rdpproxy",
	"rep",
	"urlfiltering",
	"videooptimization",
	"forwardproxy",
	"sslinterception",
	"adaptivetcp",
	"cqa",
	"ci",
	"bot",
	"apigateway",
}

func (r *NsfeatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsfeature resource.",
			},
			"wl": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web Logging.",
			},
			"sp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Surge Protection.",
			},
			"lb": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Load Balancing.",
			},
			"cs": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content Switching.",
			},
			"cr": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache Redirection.",
			},
			"cmp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Compression.",
			},
			"pq": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority Queuing.",
			},
			"ssl": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Offloading.",
			},
			"gslb": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Global Server Load Balancing.",
			},
			"hdosp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DoS Protection.",
			},
			"cf": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content Filtering.",
			},
			"ic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Integrated Caching.",
			},
			"sslvpn": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL VPN.",
			},
			"aaa": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AAA.",
			},
			"ospf": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OSPF Routing.",
			},
			"rip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RIP Routing.",
			},
			"bgp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "BGP Routing.",
			},
			"rewrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rewrite.",
			},
			"ipv6pt": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 Protocol Translation.",
			},
			"appfw": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Application Firewall.",
			},
			"responder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Responder.",
			},
			"htmlinjection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTML Injection.",
			},
			"push": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Push.",
			},
			"appflow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppFlow.",
			},
			"cloudbridge": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CloudBridge.",
			},
			"isis": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ISIS Routing.",
			},
			"ch": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Call Home.",
			},
			"appqoe": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppQoE.",
			},
			"contentaccelerator": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content Accelerator.",
			},
			"rise": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RISE.",
			},
			"feo": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Front End Optimization.",
			},
			"lsn": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Large Scale NAT.",
			},
			"rdpproxy": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP Proxy.",
			},
			"rep": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reputation.",
			},
			"urlfiltering": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL Filtering.",
			},
			"videooptimization": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Video Optimization.",
			},
			"forwardproxy": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Forward Proxy.",
			},
			"sslinterception": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Interception.",
			},
			"adaptivetcp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adaptive TCP.",
			},
			"cqa": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection Quality Analytics.",
			},
			"ci": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content Inspection.",
			},
			"bot": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bot Management.",
			},
			"apigateway": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "API Gateway.",
			},
		},
	}
}

func nsfeatureGetThePayloadFromtheConfig(ctx context.Context, data *NsfeatureResourceModel) ns.Nsfeature {
	tflog.Debug(ctx, "In nsfeatureGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsfeature := ns.Nsfeature{}

	return nsfeature
}

func nsfeatureSetAttrFromGet(ctx context.Context, data *NsfeatureResourceModel, enabledFeatures []string) *NsfeatureResourceModel {
	tflog.Debug(ctx, "In nsfeatureSetAttrFromGet Function")

	// Convert enabled features list to map for quick lookup
	enabledMap := make(map[string]bool)
	for _, feature := range enabledFeatures {
		enabledMap[feature] = true
	}

	// Set each feature based on enabled list
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

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsfeature-config")

	return data
}
