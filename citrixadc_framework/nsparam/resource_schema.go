package nsparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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

func (r *NsparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsparam resource.",
			},
			"advancedanalyticsstats": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Disable/Enable advanace analytics stats",
			},
			"aftpallowrandomsourceport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     int64default.StaticInt64(80),
				Description: "Percentage of maxClient threshold to be divided equally among PEs.",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(80),
				Description: "Percentage of spillover threshold to be divided equally among PEs.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum and maximum port (port range) that FTP services are allowed to use.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining maxclient quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured maxclient is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be \"ns_comm_key\".",
			},
			"ipttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
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
				Default:     int64default.StaticInt64(80),
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(443),
				Description: "This allows the configuration of management HTTPS port.",
			},
			"pmtumin": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(576),
				Description: "Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Interval, in minutes, for flushing the PMTU entries.",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Disable/Enable v1 or v2 proxy protocol header for client info insertion",
			},
			"securecookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CoordinatedUniversalTime"),
				Description: "Time zone for the Citrix ADC. Name of the time zone should be specified as argument.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable/Disable use_proxy_port setting",
			},
		},
	}
}

func nsparamGetThePayloadFromtheConfig(ctx context.Context, data *NsparamResourceModel) ns.Nsparam {
	tflog.Debug(ctx, "In nsparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsparam := ns.Nsparam{}
	if !data.Advancedanalyticsstats.IsNull() {
		nsparam.Advancedanalyticsstats = data.Advancedanalyticsstats.ValueString()
	}
	if !data.Aftpallowrandomsourceport.IsNull() {
		nsparam.Aftpallowrandomsourceport = data.Aftpallowrandomsourceport.ValueString()
	}
	if !data.Cip.IsNull() {
		nsparam.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() {
		nsparam.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cookieversion.IsNull() {
		nsparam.Cookieversion = data.Cookieversion.ValueString()
	}
	if !data.Crportrange.IsNull() {
		nsparam.Crportrange = data.Crportrange.ValueString()
	}
	if !data.Exclusivequotamaxclient.IsNull() {
		nsparam.Exclusivequotamaxclient = utils.IntPtr(int(data.Exclusivequotamaxclient.ValueInt64()))
	}
	if !data.Exclusivequotaspillover.IsNull() {
		nsparam.Exclusivequotaspillover = utils.IntPtr(int(data.Exclusivequotaspillover.ValueInt64()))
	}
	if !data.Ftpportrange.IsNull() {
		nsparam.Ftpportrange = data.Ftpportrange.ValueString()
	}
	if !data.Grantquotamaxclient.IsNull() {
		nsparam.Grantquotamaxclient = utils.IntPtr(int(data.Grantquotamaxclient.ValueInt64()))
	}
	if !data.Grantquotaspillover.IsNull() {
		nsparam.Grantquotaspillover = utils.IntPtr(int(data.Grantquotaspillover.ValueInt64()))
	}
	if !data.Internaluserlogin.IsNull() {
		nsparam.Internaluserlogin = data.Internaluserlogin.ValueString()
	}
	if !data.Ipttl.IsNull() {
		nsparam.Ipttl = utils.IntPtr(int(data.Ipttl.ValueInt64()))
	}
	if !data.Maxconn.IsNull() {
		nsparam.Maxconn = utils.IntPtr(int(data.Maxconn.ValueInt64()))
	}
	if !data.Maxreq.IsNull() {
		nsparam.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Mgmthttpport.IsNull() {
		nsparam.Mgmthttpport = utils.IntPtr(int(data.Mgmthttpport.ValueInt64()))
	}
	if !data.Mgmthttpsport.IsNull() {
		nsparam.Mgmthttpsport = utils.IntPtr(int(data.Mgmthttpsport.ValueInt64()))
	}
	if !data.Pmtumin.IsNull() {
		nsparam.Pmtumin = utils.IntPtr(int(data.Pmtumin.ValueInt64()))
	}
	if !data.Pmtutimeout.IsNull() {
		nsparam.Pmtutimeout = utils.IntPtr(int(data.Pmtutimeout.ValueInt64()))
	}
	if !data.Proxyprotocol.IsNull() {
		nsparam.Proxyprotocol = data.Proxyprotocol.ValueString()
	}
	if !data.Securecookie.IsNull() {
		nsparam.Securecookie = data.Securecookie.ValueString()
	}
	if !data.Servicepathingressvlan.IsNull() {
		nsparam.Servicepathingressvlan = utils.IntPtr(int(data.Servicepathingressvlan.ValueInt64()))
	}
	if !data.Tcpcip.IsNull() {
		nsparam.Tcpcip = data.Tcpcip.ValueString()
	}
	if !data.Timezone.IsNull() {
		nsparam.Timezone = data.Timezone.ValueString()
	}
	if !data.Useproxyport.IsNull() {
		nsparam.Useproxyport = data.Useproxyport.ValueString()
	}

	return nsparam
}

func nsparamSetAttrFromGet(ctx context.Context, data *NsparamResourceModel, getResponseData map[string]interface{}) *NsparamResourceModel {
	tflog.Debug(ctx, "In nsparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["advancedanalyticsstats"]; ok && val != nil {
		data.Advancedanalyticsstats = types.StringValue(val.(string))
	} else {
		data.Advancedanalyticsstats = types.StringNull()
	}
	if val, ok := getResponseData["aftpallowrandomsourceport"]; ok && val != nil {
		data.Aftpallowrandomsourceport = types.StringValue(val.(string))
	} else {
		data.Aftpallowrandomsourceport = types.StringNull()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["cookieversion"]; ok && val != nil {
		data.Cookieversion = types.StringValue(val.(string))
	} else {
		data.Cookieversion = types.StringNull()
	}
	if val, ok := getResponseData["crportrange"]; ok && val != nil {
		data.Crportrange = types.StringValue(val.(string))
	} else {
		data.Crportrange = types.StringNull()
	}
	if val, ok := getResponseData["exclusivequotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["exclusivequotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["ftpportrange"]; ok && val != nil {
		data.Ftpportrange = types.StringValue(val.(string))
	} else {
		data.Ftpportrange = types.StringNull()
	}
	if val, ok := getResponseData["grantquotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["grantquotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["internaluserlogin"]; ok && val != nil {
		data.Internaluserlogin = types.StringValue(val.(string))
	} else {
		data.Internaluserlogin = types.StringNull()
	}
	if val, ok := getResponseData["ipttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ipttl = types.Int64Value(intVal)
		}
	} else {
		data.Ipttl = types.Int64Null()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpport = types.Int64Value(intVal)
		}
	} else {
		data.Mgmthttpport = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpsport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpsport = types.Int64Value(intVal)
		}
	} else {
		data.Mgmthttpsport = types.Int64Null()
	}
	if val, ok := getResponseData["pmtumin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtumin = types.Int64Value(intVal)
		}
	} else {
		data.Pmtumin = types.Int64Null()
	}
	if val, ok := getResponseData["pmtutimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtutimeout = types.Int64Value(intVal)
		}
	} else {
		data.Pmtutimeout = types.Int64Null()
	}
	if val, ok := getResponseData["proxyprotocol"]; ok && val != nil {
		data.Proxyprotocol = types.StringValue(val.(string))
	} else {
		data.Proxyprotocol = types.StringNull()
	}
	if val, ok := getResponseData["securecookie"]; ok && val != nil {
		data.Securecookie = types.StringValue(val.(string))
	} else {
		data.Securecookie = types.StringNull()
	}
	if val, ok := getResponseData["servicepathingressvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Servicepathingressvlan = types.Int64Value(intVal)
		}
	} else {
		data.Servicepathingressvlan = types.Int64Null()
	}
	if val, ok := getResponseData["tcpcip"]; ok && val != nil {
		data.Tcpcip = types.StringValue(val.(string))
	} else {
		data.Tcpcip = types.StringNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else {
		data.Timezone = types.StringNull()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else {
		data.Useproxyport = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsparam-config")

	return data
}
