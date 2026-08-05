package nsparam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsparamResourceModel describes the resource data model.
type NsparamResourceModel struct {
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
}

// SDK v2 baseline (citrixadc/resource_citrixadc_nsparam.go): every attribute is
// Optional + Computed + ForceNew, with NO Default (values are read from the ADC).
// To preserve that contract exactly:
//   - No Default is set (auto-gen wrongly added Defaults, several without Computed,
//     which is invalid) -> Optional + Computed, value read from ADC.
//   - ForceNew is reproduced with UseStateForUnknown() (avoid known-after-apply
//     churn on Computed attrs) followed by RequiresReplaceIfConfigured() (only a
//     user-configured change forces replacement; a Computed drift does not).
func (r *NsparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsparam resource.",
			},
			"advancedanalyticsstats": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Disable/Enable advanace analytics stats",
			},
			"aftpallowrandomsourceport": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Allow the FTP server to come from a random source port for active FTP data connections",
			},
			"cip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.\n* If the CIP header is specified, it will be used as the client IP header.\n* If the CIP header is not specified, the value that has been set will be used as the client IP header.",
			},
			"cipheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Text that will be used as the client IP address header.",
			},
			"cookieversion": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Version of the cookie inserted by the system.",
			},
			"crportrange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Port range for cache redirection services.",
			},
			"exclusivequotamaxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Percentage of maxClient threshold to be divided equally among PEs.",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Percentage of spillover threshold to be divided equally among PEs.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Minimum and maximum port (port range) that FTP services are allowed to use.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining maxclient quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured maxclient is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining spillover quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured spillover is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"httpport": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "HTTP ports on the web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"icaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"internaluserlogin": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be \"ns_comm_key\".",
			},
			"ipttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Set the IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.",
			},
			"maxconn": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Maximum number of connections that will be made from the appliance to the web server(s) attached to it. The value entered here is applied globally to all attached servers.",
			},
			"maxreq": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Maximum number of requests that the system can pass on a particular connection between the appliance and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed. This value is overridden by the maximum number of requests configured on the individual service.",
			},
			"mgmthttpport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "This allows the configuration of management HTTPS port.",
			},
			"pmtumin": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Interval, in minutes, for flushing the PMTU entries.",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Disable/Enable v1 or v2 proxy protocol header for client info insertion",
			},
			"securecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Enable or disable secure flag for persistence cookie.",
			},
			"secureicaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any\n            client request that has a destination port matching one of these configured ports.",
			},
			"servicepathingressvlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "VLAN on which the subscriber traffic arrives on the appliance.",
			},
			"tcpcip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.",
			},
			"timezone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Time zone for the Citrix ADC. Name of the time zone should be specified as argument.",
			},
			"useproxyport": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Enable/Disable use_proxy_port setting",
			},
		},
	}
}

// nsparamGetThePayloadFromtheConfig builds the NITRO payload as a map so that ONLY
// user-configured attributes are sent (mirroring the SDK v2 GetOk/GetOkExists
// behaviour). This matters because several ns.Nsparam fields (maxconn, maxreq,
// cookieversion, the quota fields) are declared WITHOUT `omitempty` ("Zero is a
// valid value"); serialising the typed struct would leak their zero values on
// every write. Using a map avoids clobbering unrelated global parameters and lets
// a configured maxconn=0 be sent while an unconfigured maxconn is omitted.
func nsparamGetThePayloadFromtheConfig(ctx context.Context, data *NsparamResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nsparamGetThePayloadFromtheConfig Function")

	nsparam := make(map[string]interface{})

	if !data.Advancedanalyticsstats.IsNull() && !data.Advancedanalyticsstats.IsUnknown() {
		nsparam["advancedanalyticsstats"] = data.Advancedanalyticsstats.ValueString()
	}
	if !data.Aftpallowrandomsourceport.IsNull() && !data.Aftpallowrandomsourceport.IsUnknown() {
		nsparam["aftpallowrandomsourceport"] = data.Aftpallowrandomsourceport.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		nsparam["cip"] = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		nsparam["cipheader"] = data.Cipheader.ValueString()
	}
	if !data.Cookieversion.IsNull() && !data.Cookieversion.IsUnknown() {
		nsparam["cookieversion"] = data.Cookieversion.ValueString()
	}
	if !data.Crportrange.IsNull() && !data.Crportrange.IsUnknown() {
		nsparam["crportrange"] = data.Crportrange.ValueString()
	}
	if !data.Exclusivequotamaxclient.IsNull() && !data.Exclusivequotamaxclient.IsUnknown() {
		nsparam["exclusivequotamaxclient"] = int(data.Exclusivequotamaxclient.ValueInt64())
	}
	if !data.Exclusivequotaspillover.IsNull() && !data.Exclusivequotaspillover.IsUnknown() {
		nsparam["exclusivequotaspillover"] = int(data.Exclusivequotaspillover.ValueInt64())
	}
	if !data.Ftpportrange.IsNull() && !data.Ftpportrange.IsUnknown() {
		nsparam["ftpportrange"] = data.Ftpportrange.ValueString()
	}
	if !data.Grantquotamaxclient.IsNull() && !data.Grantquotamaxclient.IsUnknown() {
		nsparam["grantquotamaxclient"] = int(data.Grantquotamaxclient.ValueInt64())
	}
	if !data.Grantquotaspillover.IsNull() && !data.Grantquotaspillover.IsUnknown() {
		nsparam["grantquotaspillover"] = int(data.Grantquotaspillover.ValueInt64())
	}
	if !data.Httpport.IsNull() && !data.Httpport.IsUnknown() {
		var httpportList []int
		data.Httpport.ElementsAs(ctx, &httpportList, false)
		nsparam["httpport"] = httpportList
	}
	if !data.Icaports.IsNull() && !data.Icaports.IsUnknown() {
		var icaportsList []int
		data.Icaports.ElementsAs(ctx, &icaportsList, false)
		nsparam["icaports"] = icaportsList
	}
	if !data.Internaluserlogin.IsNull() && !data.Internaluserlogin.IsUnknown() {
		nsparam["internaluserlogin"] = data.Internaluserlogin.ValueString()
	}
	if !data.Ipttl.IsNull() && !data.Ipttl.IsUnknown() {
		nsparam["ipttl"] = int(data.Ipttl.ValueInt64())
	}
	if !data.Maxconn.IsNull() && !data.Maxconn.IsUnknown() {
		nsparam["maxconn"] = int(data.Maxconn.ValueInt64())
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		nsparam["maxreq"] = int(data.Maxreq.ValueInt64())
	}
	if !data.Mgmthttpport.IsNull() && !data.Mgmthttpport.IsUnknown() {
		nsparam["mgmthttpport"] = int(data.Mgmthttpport.ValueInt64())
	}
	if !data.Mgmthttpsport.IsNull() && !data.Mgmthttpsport.IsUnknown() {
		nsparam["mgmthttpsport"] = int(data.Mgmthttpsport.ValueInt64())
	}
	if !data.Pmtumin.IsNull() && !data.Pmtumin.IsUnknown() {
		nsparam["pmtumin"] = int(data.Pmtumin.ValueInt64())
	}
	if !data.Pmtutimeout.IsNull() && !data.Pmtutimeout.IsUnknown() {
		nsparam["pmtutimeout"] = int(data.Pmtutimeout.ValueInt64())
	}
	if !data.Proxyprotocol.IsNull() && !data.Proxyprotocol.IsUnknown() {
		nsparam["proxyprotocol"] = data.Proxyprotocol.ValueString()
	}
	if !data.Securecookie.IsNull() && !data.Securecookie.IsUnknown() {
		nsparam["securecookie"] = data.Securecookie.ValueString()
	}
	if !data.Secureicaports.IsNull() && !data.Secureicaports.IsUnknown() {
		var secureicaportsList []int
		data.Secureicaports.ElementsAs(ctx, &secureicaportsList, false)
		nsparam["secureicaports"] = secureicaportsList
	}
	if !data.Servicepathingressvlan.IsNull() && !data.Servicepathingressvlan.IsUnknown() {
		nsparam["servicepathingressvlan"] = int(data.Servicepathingressvlan.ValueInt64())
	}
	if !data.Tcpcip.IsNull() && !data.Tcpcip.IsUnknown() {
		nsparam["tcpcip"] = data.Tcpcip.ValueString()
	}
	if !data.Timezone.IsNull() && !data.Timezone.IsUnknown() {
		nsparam["timezone"] = data.Timezone.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		nsparam["useproxyport"] = data.Useproxyport.ValueString()
	}

	return nsparam
}

// nsparamSetAttrFromGet maps the NITRO GET response onto the model. The else
// branches only null a value when it is still Unknown, so a user-configured value
// that NITRO omits from GET (the omit-on-default trap) is never clobbered while
// Computed attributes still resolve to a known value after apply.
func nsparamSetAttrFromGet(ctx context.Context, data *NsparamResourceModel, getResponseData map[string]interface{}) *NsparamResourceModel {
	tflog.Debug(ctx, "In nsparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["advancedanalyticsstats"]; ok && val != nil {
		data.Advancedanalyticsstats = types.StringValue(val.(string))
	} else if data.Advancedanalyticsstats.IsUnknown() {
		data.Advancedanalyticsstats = types.StringNull()
	}
	if val, ok := getResponseData["aftpallowrandomsourceport"]; ok && val != nil {
		data.Aftpallowrandomsourceport = types.StringValue(val.(string))
	} else if data.Aftpallowrandomsourceport.IsUnknown() {
		data.Aftpallowrandomsourceport = types.StringNull()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else if data.Cip.IsUnknown() {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else if data.Cipheader.IsUnknown() {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["cookieversion"]; ok && val != nil {
		data.Cookieversion = types.StringValue(val.(string))
	} else if data.Cookieversion.IsUnknown() {
		data.Cookieversion = types.StringNull()
	}
	if val, ok := getResponseData["crportrange"]; ok && val != nil {
		data.Crportrange = types.StringValue(val.(string))
	} else if data.Crportrange.IsUnknown() {
		data.Crportrange = types.StringNull()
	}
	if val, ok := getResponseData["exclusivequotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotamaxclient = types.Int64Value(intVal)
		}
	} else if data.Exclusivequotamaxclient.IsUnknown() {
		data.Exclusivequotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["exclusivequotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotaspillover = types.Int64Value(intVal)
		}
	} else if data.Exclusivequotaspillover.IsUnknown() {
		data.Exclusivequotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["ftpportrange"]; ok && val != nil {
		data.Ftpportrange = types.StringValue(val.(string))
	} else if data.Ftpportrange.IsUnknown() {
		data.Ftpportrange = types.StringNull()
	}
	if val, ok := getResponseData["grantquotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotamaxclient = types.Int64Value(intVal)
		}
	} else if data.Grantquotamaxclient.IsUnknown() {
		data.Grantquotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["grantquotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotaspillover = types.Int64Value(intVal)
		}
	} else if data.Grantquotaspillover.IsUnknown() {
		data.Grantquotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["httpport"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			intList := utils.StringListToIntList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, intList)
			data.Httpport = listValue
		} else if data.Httpport.IsUnknown() {
			data.Httpport = types.ListNull(types.Int64Type)
		}
	} else if data.Httpport.IsUnknown() {
		data.Httpport = types.ListNull(types.Int64Type)
	}
	if val, ok := getResponseData["icaports"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			intList := utils.StringListToIntList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, intList)
			data.Icaports = listValue
		} else if data.Icaports.IsUnknown() {
			data.Icaports = types.ListNull(types.Int64Type)
		}
	} else if data.Icaports.IsUnknown() {
		data.Icaports = types.ListNull(types.Int64Type)
	}
	if val, ok := getResponseData["internaluserlogin"]; ok && val != nil {
		data.Internaluserlogin = types.StringValue(val.(string))
	} else if data.Internaluserlogin.IsUnknown() {
		data.Internaluserlogin = types.StringNull()
	}
	if val, ok := getResponseData["ipttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ipttl = types.Int64Value(intVal)
		}
	} else if data.Ipttl.IsUnknown() {
		data.Ipttl = types.Int64Null()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else if data.Maxconn.IsUnknown() {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else if data.Maxreq.IsUnknown() {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpport = types.Int64Value(intVal)
		}
	} else if data.Mgmthttpport.IsUnknown() {
		data.Mgmthttpport = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpsport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpsport = types.Int64Value(intVal)
		}
	} else if data.Mgmthttpsport.IsUnknown() {
		data.Mgmthttpsport = types.Int64Null()
	}
	if val, ok := getResponseData["pmtumin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtumin = types.Int64Value(intVal)
		}
	} else if data.Pmtumin.IsUnknown() {
		data.Pmtumin = types.Int64Null()
	}
	if val, ok := getResponseData["pmtutimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtutimeout = types.Int64Value(intVal)
		}
	} else if data.Pmtutimeout.IsUnknown() {
		data.Pmtutimeout = types.Int64Null()
	}
	if val, ok := getResponseData["proxyprotocol"]; ok && val != nil {
		data.Proxyprotocol = types.StringValue(val.(string))
	} else if data.Proxyprotocol.IsUnknown() {
		data.Proxyprotocol = types.StringNull()
	}
	if val, ok := getResponseData["securecookie"]; ok && val != nil {
		data.Securecookie = types.StringValue(val.(string))
	} else if data.Securecookie.IsUnknown() {
		data.Securecookie = types.StringNull()
	}
	if val, ok := getResponseData["secureicaports"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			intList := utils.StringListToIntList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, intList)
			data.Secureicaports = listValue
		} else if data.Secureicaports.IsUnknown() {
			data.Secureicaports = types.ListNull(types.Int64Type)
		}
	} else if data.Secureicaports.IsUnknown() {
		data.Secureicaports = types.ListNull(types.Int64Type)
	}
	if val, ok := getResponseData["servicepathingressvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Servicepathingressvlan = types.Int64Value(intVal)
		}
	} else if data.Servicepathingressvlan.IsUnknown() {
		data.Servicepathingressvlan = types.Int64Null()
	}
	if val, ok := getResponseData["tcpcip"]; ok && val != nil {
		data.Tcpcip = types.StringValue(val.(string))
	} else if data.Tcpcip.IsUnknown() {
		data.Tcpcip = types.StringNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else if data.Timezone.IsUnknown() {
		data.Timezone = types.StringNull()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else if data.Useproxyport.IsUnknown() {
		data.Useproxyport = types.StringNull()
	}

	// Set ID for the resource (singleton - static ID)
	data.Id = types.StringValue("nsparam-config")

	return data
}
