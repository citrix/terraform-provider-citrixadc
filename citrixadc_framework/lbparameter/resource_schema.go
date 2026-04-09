package lbparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbparameterResourceModel describes the resource data model.
type LbparameterResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Allowboundsvcremoval          types.String `tfsdk:"allowboundsvcremoval"`
	Computedadccookieattribute    types.String `tfsdk:"computedadccookieattribute"`
	Consolidatedlconn             types.String `tfsdk:"consolidatedlconn"`
	Cookiepassphrase              types.String `tfsdk:"cookiepassphrase"`
	CookiepassphraseWo            types.String `tfsdk:"cookiepassphrase_wo"`
	CookiepassphraseWoVersion     types.Int64  `tfsdk:"cookiepassphrase_wo_version"`
	Dbsttl                        types.Int64  `tfsdk:"dbsttl"`
	Dropmqttjumbomessage          types.String `tfsdk:"dropmqttjumbomessage"`
	Httponlycookieflag            types.String `tfsdk:"httponlycookieflag"`
	Lbhashalgorithm               types.String `tfsdk:"lbhashalgorithm"`
	Lbhashfingers                 types.Int64  `tfsdk:"lbhashfingers"`
	Literaladccookieattribute     types.String `tfsdk:"literaladccookieattribute"`
	Maxpipelinenat                types.Int64  `tfsdk:"maxpipelinenat"`
	Monitorconnectionclose        types.String `tfsdk:"monitorconnectionclose"`
	Monitorskipmaxclient          types.String `tfsdk:"monitorskipmaxclient"`
	Preferdirectroute             types.String `tfsdk:"preferdirectroute"`
	Proximityfromself             types.String `tfsdk:"proximityfromself"`
	Retainservicestate            types.String `tfsdk:"retainservicestate"`
	Startuprrfactor               types.Int64  `tfsdk:"startuprrfactor"`
	Storemqttclientidandusername  types.String `tfsdk:"storemqttclientidandusername"`
	Undefaction                   types.String `tfsdk:"undefaction"`
	Useencryptedpersistencecookie types.String `tfsdk:"useencryptedpersistencecookie"`
	Useportforhashlb              types.String `tfsdk:"useportforhashlb"`
	Usesecuredpersistencecookie   types.String `tfsdk:"usesecuredpersistencecookie"`
	Vserverspecificmac            types.String `tfsdk:"vserverspecificmac"`
}

func (r *LbparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbparameter resource.",
			},
			"allowboundsvcremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used, to enable/disable the option of svc/svcgroup removal, if it is bound to one or more vserver. If it is enabled, the svc/svcgroup can be removed, even if it bound to vservers. If disabled, an error will be thrown, when the user tries to remove a svc/svcgroup without unbinding from its vservers.",
			},
			"computedadccookieattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.\n\nSample usage -\n             add ns variable lbvar -type TEXT(100) -scope Transaction\n             add ns assignment lbassign -variable $lbvar -set \"\\\\\";SameSite=Strict\\\\\"\"\n             add rewrite policy lbpol <valid policy expression> lbassign\n             bind rewrite global lbpol 100 next -type RES_OVERRIDE\n             set lb param -ComputedADCCookieAttribute \"$lbvar\"\n             For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie",
			},
			"consolidatedlconn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To find the service with the fewest connections, the virtual server uses the consolidated connection statistics from all the packet engines. The NO setting allows consideration of only the number of connections on the packet engine that received the new connection.",
			},
			"cookiepassphrase": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.",
			},
			"cookiepassphrase_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.",
			},
			"cookiepassphrase_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a cookiepassphrase_wo update.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service. The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
			},
			"dropmqttjumbomessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When this option is enabled, MQTT messages of length greater than 64k will be dropped and the client/server connections will be reset.",
			},
			"httponlycookieflag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks.",
			},
			"lbhashalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH).",
			},
			"lbhashfingers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory",
			},
			"literaladccookieattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence).\n\nSample usage -\n             set lb parameter -LiteralADCCookieAttribute \";SameSite=None\"",
			},
			"maxpipelinenat": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. This parameter is applicable to ANY service type and all UDP service types (except DNS) and only when \"svrTimeout\" is set to zero. A value of 0 (zero) applies no limit to the number of concurrent requests allowed on a single client connection",
			},
			"monitorconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitorskipmaxclient": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When a monitor initiates a connection to a service, do not check to determine whether the number of connections to the service has reached the limit specified by the service's Max Clients setting. Enables monitoring to continue even if the service has reached its connection limit.",
			},
			"preferdirectroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform route lookup for traffic received by the Citrix ADC, and forward the traffic according to configured routes. Do not set this parameter if you want a wildcard virtual server to direct packets received by the appliance to an intermediary device, such as a firewall, even if their destination is directly connected to the appliance. Route lookup is performed after the packets have been processed and returned by the intermediary device.",
			},
			"proximityfromself": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the ADC location instead of client IP for static proximity LB or GSLB decision.",
			},
			"retainservicestate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to retain the original state of service or servicegroup member when an enable server command is issued.",
			},
			"startuprrfactor": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests, per service, for which to apply the round robin load balancing method before switching to the configured load balancing method, thus allowing services to ramp up gradually to full load. Until the specified number of requests is distributed, the Citrix ADC is said to be implementing the slow start mode (or startup round robin). Implemented for a virtual server when one of the following is true:\n* The virtual server is newly created.\n* One or more services are newly bound to the virtual server.\n* One or more services bound to the virtual server are enabled.\n* The load balancing method is changed.\nThis parameter applies to all the load balancing virtual servers configured on the Citrix ADC, except for those virtual servers for which the virtual server-level slow start parameters (New Service Startup Request Rate and Increment Interval) are configured. If the global slow start parameter and the slow start parameters for a given virtual server are not set, the appliance implements a default slow start for the virtual server, as follows:\n* For a newly configured virtual server, the appliance implements slow start for the first 100 requests received by the virtual server.\n* For an existing virtual server, if one or more services are newly bound or newly enabled, or if the load balancing method is changed, the appliance dynamically computes the number of requests for which to implement startup round robin. It obtains this number by multiplying the request rate by the number of bound services (it includes services that are marked as DOWN). For example, if the current request rate is 20 requests/s and ten services are bound to the virtual server, the appliance performs startup round robin for 200 requests.\nNot applicable to a virtual server for which a hash based load balancing method is configured.",
			},
			"storemqttclientidandusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option allows to store the MQTT clientid and username in transactional logs",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
			"useencryptedpersistencecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encode persistence cookie values using SHA2 hash.",
			},
			"useportforhashlb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the port number of the service when creating a hash for hash based load balancing methods. With the NO setting, only the IP address of the service is considered when creating a hash.",
			},
			"usesecuredpersistencecookie": schema.StringAttribute{
				Optional:    true,
				Description: "Encode persistence cookie values using SHA2 hash.",
			},
			"vserverspecificmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow a MAC-mode virtual server to accept traffic returned by an intermediary device, such as a firewall, to which the traffic was previously forwarded by another MAC-mode virtual server. The second virtual server can then distribute that traffic across the destination server farm. Also useful when load balancing Branch Repeater appliances.\nNote: The second virtual server can also send the traffic to another set of intermediary devices, such as another set of firewalls. If necessary, you can configure multiple MAC-mode virtual servers to pass traffic successively through multiple sets of intermediary devices.",
			},
		},
	}
}

func lbparameterGetThePayloadFromthePlan(ctx context.Context, data *LbparameterResourceModel) lb.Lbparameter {
	tflog.Debug(ctx, "In lbparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbparameter := lb.Lbparameter{}
	if !data.Allowboundsvcremoval.IsNull() {
		lbparameter.Allowboundsvcremoval = data.Allowboundsvcremoval.ValueString()
	}
	if !data.Computedadccookieattribute.IsNull() {
		lbparameter.Computedadccookieattribute = data.Computedadccookieattribute.ValueString()
	}
	if !data.Consolidatedlconn.IsNull() {
		lbparameter.Consolidatedlconn = data.Consolidatedlconn.ValueString()
	}
	if !data.Cookiepassphrase.IsNull() {
		lbparameter.Cookiepassphrase = data.Cookiepassphrase.ValueString()
	}
	// Skip write-only attribute: cookiepassphrase_wo
	// Skip version tracker attribute: cookiepassphrase_wo_version
	if !data.Dbsttl.IsNull() {
		lbparameter.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	if !data.Dropmqttjumbomessage.IsNull() {
		lbparameter.Dropmqttjumbomessage = data.Dropmqttjumbomessage.ValueString()
	}
	if !data.Httponlycookieflag.IsNull() {
		lbparameter.Httponlycookieflag = data.Httponlycookieflag.ValueString()
	}
	if !data.Lbhashalgorithm.IsNull() {
		lbparameter.Lbhashalgorithm = data.Lbhashalgorithm.ValueString()
	}
	if !data.Lbhashfingers.IsNull() {
		lbparameter.Lbhashfingers = utils.IntPtr(int(data.Lbhashfingers.ValueInt64()))
	}
	if !data.Literaladccookieattribute.IsNull() {
		lbparameter.Literaladccookieattribute = data.Literaladccookieattribute.ValueString()
	}
	if !data.Maxpipelinenat.IsNull() {
		lbparameter.Maxpipelinenat = utils.IntPtr(int(data.Maxpipelinenat.ValueInt64()))
	}
	if !data.Monitorconnectionclose.IsNull() {
		lbparameter.Monitorconnectionclose = data.Monitorconnectionclose.ValueString()
	}
	if !data.Monitorskipmaxclient.IsNull() {
		lbparameter.Monitorskipmaxclient = data.Monitorskipmaxclient.ValueString()
	}
	if !data.Preferdirectroute.IsNull() {
		lbparameter.Preferdirectroute = data.Preferdirectroute.ValueString()
	}
	if !data.Proximityfromself.IsNull() {
		lbparameter.Proximityfromself = data.Proximityfromself.ValueString()
	}
	if !data.Retainservicestate.IsNull() {
		lbparameter.Retainservicestate = data.Retainservicestate.ValueString()
	}
	if !data.Startuprrfactor.IsNull() {
		lbparameter.Startuprrfactor = utils.IntPtr(int(data.Startuprrfactor.ValueInt64()))
	}
	if !data.Storemqttclientidandusername.IsNull() {
		lbparameter.Storemqttclientidandusername = data.Storemqttclientidandusername.ValueString()
	}
	if !data.Undefaction.IsNull() {
		lbparameter.Undefaction = data.Undefaction.ValueString()
	}
	if !data.Useencryptedpersistencecookie.IsNull() {
		lbparameter.Useencryptedpersistencecookie = data.Useencryptedpersistencecookie.ValueString()
	}
	if !data.Useportforhashlb.IsNull() {
		lbparameter.Useportforhashlb = data.Useportforhashlb.ValueString()
	}
	if !data.Usesecuredpersistencecookie.IsNull() {
		lbparameter.Usesecuredpersistencecookie = data.Usesecuredpersistencecookie.ValueString()
	}
	if !data.Vserverspecificmac.IsNull() {
		lbparameter.Vserverspecificmac = data.Vserverspecificmac.ValueString()
	}

	return lbparameter
}

func lbparameterGetThePayloadFromtheConfig(ctx context.Context, data *LbparameterResourceModel, payload *lb.Lbparameter) {
	tflog.Debug(ctx, "In lbparameterGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: cookiepassphrase_wo -> cookiepassphrase
	if !data.CookiepassphraseWo.IsNull() {
		cookiepassphraseWo := data.CookiepassphraseWo.ValueString()
		if cookiepassphraseWo != "" {
			payload.Cookiepassphrase = cookiepassphraseWo
		}
	}
}

func lbparameterSetAttrFromGet(ctx context.Context, data *LbparameterResourceModel, getResponseData map[string]interface{}) *LbparameterResourceModel {
	tflog.Debug(ctx, "In lbparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allowboundsvcremoval"]; ok && val != nil {
		data.Allowboundsvcremoval = types.StringValue(val.(string))
	} else {
		data.Allowboundsvcremoval = types.StringNull()
	}
	if val, ok := getResponseData["computedadccookieattribute"]; ok && val != nil {
		data.Computedadccookieattribute = types.StringValue(val.(string))
	} else {
		data.Computedadccookieattribute = types.StringNull()
	}
	if val, ok := getResponseData["consolidatedlconn"]; ok && val != nil {
		data.Consolidatedlconn = types.StringValue(val.(string))
	} else {
		data.Consolidatedlconn = types.StringNull()
	}
	// cookiepassphrase is not returned by NITRO API (secret/ephemeral) - retain from config
	// cookiepassphrase_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// cookiepassphrase_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dbsttl = types.Int64Value(intVal)
		}
	} else {
		data.Dbsttl = types.Int64Null()
	}
	if val, ok := getResponseData["dropmqttjumbomessage"]; ok && val != nil {
		data.Dropmqttjumbomessage = types.StringValue(val.(string))
	} else {
		data.Dropmqttjumbomessage = types.StringNull()
	}
	if val, ok := getResponseData["httponlycookieflag"]; ok && val != nil {
		data.Httponlycookieflag = types.StringValue(val.(string))
	} else {
		data.Httponlycookieflag = types.StringNull()
	}
	if val, ok := getResponseData["lbhashalgorithm"]; ok && val != nil {
		data.Lbhashalgorithm = types.StringValue(val.(string))
	} else {
		data.Lbhashalgorithm = types.StringNull()
	}
	if val, ok := getResponseData["lbhashfingers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lbhashfingers = types.Int64Value(intVal)
		}
	} else {
		data.Lbhashfingers = types.Int64Null()
	}
	if val, ok := getResponseData["literaladccookieattribute"]; ok && val != nil {
		data.Literaladccookieattribute = types.StringValue(val.(string))
	} else {
		data.Literaladccookieattribute = types.StringNull()
	}
	if val, ok := getResponseData["maxpipelinenat"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpipelinenat = types.Int64Value(intVal)
		}
	} else {
		data.Maxpipelinenat = types.Int64Null()
	}
	if val, ok := getResponseData["monitorconnectionclose"]; ok && val != nil {
		data.Monitorconnectionclose = types.StringValue(val.(string))
	} else {
		data.Monitorconnectionclose = types.StringNull()
	}
	if val, ok := getResponseData["monitorskipmaxclient"]; ok && val != nil {
		data.Monitorskipmaxclient = types.StringValue(val.(string))
	} else {
		data.Monitorskipmaxclient = types.StringNull()
	}
	if val, ok := getResponseData["preferdirectroute"]; ok && val != nil {
		data.Preferdirectroute = types.StringValue(val.(string))
	} else {
		data.Preferdirectroute = types.StringNull()
	}
	if val, ok := getResponseData["proximityfromself"]; ok && val != nil {
		data.Proximityfromself = types.StringValue(val.(string))
	} else {
		data.Proximityfromself = types.StringNull()
	}
	if val, ok := getResponseData["retainservicestate"]; ok && val != nil {
		data.Retainservicestate = types.StringValue(val.(string))
	} else {
		data.Retainservicestate = types.StringNull()
	}
	if val, ok := getResponseData["startuprrfactor"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Startuprrfactor = types.Int64Value(intVal)
		}
	} else {
		data.Startuprrfactor = types.Int64Null()
	}
	if val, ok := getResponseData["storemqttclientidandusername"]; ok && val != nil {
		data.Storemqttclientidandusername = types.StringValue(val.(string))
	} else {
		data.Storemqttclientidandusername = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}
	if val, ok := getResponseData["useencryptedpersistencecookie"]; ok && val != nil {
		data.Useencryptedpersistencecookie = types.StringValue(val.(string))
	} else {
		data.Useencryptedpersistencecookie = types.StringNull()
	}
	if val, ok := getResponseData["useportforhashlb"]; ok && val != nil {
		data.Useportforhashlb = types.StringValue(val.(string))
	} else {
		data.Useportforhashlb = types.StringNull()
	}
	if val, ok := getResponseData["usesecuredpersistencecookie"]; ok && val != nil {
		data.Usesecuredpersistencecookie = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["vserverspecificmac"]; ok && val != nil {
		data.Vserverspecificmac = types.StringValue(val.(string))
	} else {
		data.Vserverspecificmac = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("lbparameter-config")

	return data
}
