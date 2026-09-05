package nsparam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsparamDataSourceModel describes the DATASOURCE data model. It mirrors the
// global parameters the NITRO `nsparam` GET returns, including the read-only
// fields the resource intentionally omits. It is decoupled from the resource
// model so the data source can expose the full GET projection.
type NsparamDataSourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Advancedanalyticsstats    types.String `tfsdk:"advancedanalyticsstats"`
	Aftpallowrandomsourceport types.String `tfsdk:"aftpallowrandomsourceport"`
	Cip                       types.String `tfsdk:"cip"`
	Cipheader                 types.String `tfsdk:"cipheader"`
	Cookieversion             types.String `tfsdk:"cookieversion"`
	Crportrange               types.String `tfsdk:"crportrange"`
	Exclusivequotamaxclient   types.Int64  `tfsdk:"exclusivequotamaxclient"`
	Exclusivequotaspillover   types.Int64  `tfsdk:"exclusivequotaspillover"`
	Ftpportrange              types.String `tfsdk:"ftpportrange"`
	Grantquotamaxclient       types.Int64  `tfsdk:"grantquotamaxclient"`
	Grantquotaspillover       types.Int64  `tfsdk:"grantquotaspillover"`
	Httpport                  types.List   `tfsdk:"httpport"`
	Icaports                  types.List   `tfsdk:"icaports"`
	Internaluserlogin         types.String `tfsdk:"internaluserlogin"`
	Ipttl                     types.Int64  `tfsdk:"ipttl"`
	Maxconn                   types.Int64  `tfsdk:"maxconn"`
	Maxreq                    types.Int64  `tfsdk:"maxreq"`
	Mgmthttpport              types.Int64  `tfsdk:"mgmthttpport"`
	Mgmthttpsport             types.Int64  `tfsdk:"mgmthttpsport"`
	Pmtumin                   types.Int64  `tfsdk:"pmtumin"`
	Pmtutimeout               types.Int64  `tfsdk:"pmtutimeout"`
	Proxyprotocol             types.String `tfsdk:"proxyprotocol"`
	Securecookie              types.String `tfsdk:"securecookie"`
	Secureicaports            types.List   `tfsdk:"secureicaports"`
	Servicepathingressvlan    types.Int64  `tfsdk:"servicepathingressvlan"`
	Tcpcip                    types.String `tfsdk:"tcpcip"`
	Timezone                  types.String `tfsdk:"timezone"`
	Useproxyport              types.String `tfsdk:"useproxyport"`

	// Read-only (GET-only) parameter from the NITRO doc read-only set
	// (zion73x_readonly/nsparam.json). Never settable; from GET.
	Autoscaleoption types.Int64 `tfsdk:"autoscaleoption"`
}

// nsparamDataSourceInt64List projects a NITRO array value onto a types.List of
// Int64 (matching the existing Int64-typed list attributes), returning a typed
// Null when the GET omits the key.
func nsparamDataSourceInt64List(ctx context.Context, g map[string]interface{}, key string) types.List {
	if val, ok := g[key]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			intList := utils.StringListToIntList(sliceVal)
			if listValue, d := types.ListValueFrom(ctx, types.Int64Type, intList); !d.HasError() {
				return listValue
			}
		}
	}
	return types.ListNull(types.Int64Type)
}

// nsparamDataSourceSetAttrFromGet projects a NITRO nsparam GET response onto the
// data-source model using the shared utils.MapGet* helpers. Attributes the GET
// omits are left Null.
func nsparamDataSourceSetAttrFromGet(ctx context.Context, data *NsparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsparamDataSourceSetAttrFromGet Function")

	data.Advancedanalyticsstats = utils.MapGetString(g, "advancedanalyticsstats")
	data.Aftpallowrandomsourceport = utils.MapGetString(g, "aftpallowrandomsourceport")
	data.Cip = utils.MapGetString(g, "cip")
	data.Cipheader = utils.MapGetString(g, "cipheader")
	data.Cookieversion = utils.MapGetString(g, "cookieversion")
	data.Crportrange = utils.MapGetString(g, "crportrange")
	data.Exclusivequotamaxclient = utils.MapGetInt64(g, "exclusivequotamaxclient")
	data.Exclusivequotaspillover = utils.MapGetInt64(g, "exclusivequotaspillover")
	data.Ftpportrange = utils.MapGetString(g, "ftpportrange")
	data.Grantquotamaxclient = utils.MapGetInt64(g, "grantquotamaxclient")
	data.Grantquotaspillover = utils.MapGetInt64(g, "grantquotaspillover")
	data.Httpport = nsparamDataSourceInt64List(ctx, g, "httpport")
	data.Icaports = nsparamDataSourceInt64List(ctx, g, "icaports")
	data.Internaluserlogin = utils.MapGetString(g, "internaluserlogin")
	data.Ipttl = utils.MapGetInt64(g, "ipttl")
	data.Maxconn = utils.MapGetInt64(g, "maxconn")
	data.Maxreq = utils.MapGetInt64(g, "maxreq")
	data.Mgmthttpport = utils.MapGetInt64(g, "mgmthttpport")
	data.Mgmthttpsport = utils.MapGetInt64(g, "mgmthttpsport")
	data.Pmtumin = utils.MapGetInt64(g, "pmtumin")
	data.Pmtutimeout = utils.MapGetInt64(g, "pmtutimeout")
	data.Proxyprotocol = utils.MapGetString(g, "proxyprotocol")
	data.Securecookie = utils.MapGetString(g, "securecookie")
	data.Secureicaports = nsparamDataSourceInt64List(ctx, g, "secureicaports")
	data.Servicepathingressvlan = utils.MapGetInt64(g, "servicepathingressvlan")
	data.Tcpcip = utils.MapGetString(g, "tcpcip")
	data.Timezone = utils.MapGetString(g, "timezone")
	data.Useproxyport = utils.MapGetString(g, "useproxyport")

	// Read-only parameter.
	data.Autoscaleoption = utils.MapGetInt64(g, "autoscaleoption")

	// nsparam is a singleton - static ID.
	data.Id = types.StringValue("nsparam-config")
}

func NsparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advancedanalyticsstats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disable/Enable advanace analytics stats",
			},
			"aftpallowrandomsourceport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow the FTP server to come from a random source port for active FTP data connections",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.\n* If the CIP header is specified, it will be used as the client IP header.\n* If the CIP header is not specified, the value that has been set will be used as the client IP header.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Text that will be used as the client IP address header.",
			},
			"cookieversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version of the cookie inserted by the system.",
			},
			"crportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port range for cache redirection services.",
			},
			"exclusivequotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of maxClient threshold to be divided equally among PEs.",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of spillover threshold to be divided equally among PEs.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum and maximum port (port range) that FTP services are allowed to use.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining maxclient quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured maxclient is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining spillover quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured spillover is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"httpport": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "HTTP ports on the web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"icaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"internaluserlogin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be \"ns_comm_key\".",
			},
			"ipttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of connections that will be made from the appliance to the web server(s) attached to it. The value entered here is applied globally to all attached servers.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that the system can pass on a particular connection between the appliance and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed. This value is overridden by the maximum number of requests configured on the individual service.",
			},
			"mgmthttpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allows the configuration of management HTTPS port.",
			},
			"pmtumin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in minutes, for flushing the PMTU entries.",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disable/Enable v1 or v2 proxy protocol header for client info insertion",
			},
			"securecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable secure flag for persistence cookie.",
			},
			"secureicaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any\n            client request that has a destination port matching one of these configured ports.",
			},
			"servicepathingressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN on which the subscriber traffic arrives on the appliance.",
			},
			"tcpcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone for the Citrix ADC. Name of the time zone should be specified as argument.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable use_proxy_port setting",
			},

			// Read-only (GET-only) parameter surfaced by the data source.
			"autoscaleoption": schema.Int64Attribute{
				Computed:    true,
				Description: "64 bits are provided for communication between ADM and ADC in cloud deployments. Currently LSB 3 bits are used (0x01=AWS, 0x02=Azure, 0x04=GCP). Null when the appliance omits it.",
			},
		},
	}
}
