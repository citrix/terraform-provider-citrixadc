package lbparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbparameterDataSourceModel is the data-source-specific model, decoupled from
// LbparameterResourceModel. A data source is a pure read surface (Read only), so
// it exposes the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (sessionsthreshold, builtin, feature, adccookieattributewarningmsg,
// lbhashalgowinsize, overridepersistencyfororder).
type LbparameterDataSourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Allowboundsvcremoval          types.String `tfsdk:"allowboundsvcremoval"`
	Computedadccookieattribute    types.String `tfsdk:"computedadccookieattribute"`
	Consolidatedlconn             types.String `tfsdk:"consolidatedlconn"`
	Cookiepassphrase              types.String `tfsdk:"cookiepassphrase"`
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
	Radiusmessageauthenticator    types.String `tfsdk:"radiusmessageauthenticator"`
	Retainservicestate            types.String `tfsdk:"retainservicestate"`
	Startuprrfactor               types.Int64  `tfsdk:"startuprrfactor"`
	Storemqttclientidandusername  types.String `tfsdk:"storemqttclientidandusername"`
	Undefaction                   types.String `tfsdk:"undefaction"`
	Useencryptedpersistencecookie types.String `tfsdk:"useencryptedpersistencecookie"`
	Useportforhashlb              types.String `tfsdk:"useportforhashlb"`
	Usesecuredpersistencecookie   types.String `tfsdk:"usesecuredpersistencecookie"`
	Vserverspecificmac            types.String `tfsdk:"vserverspecificmac"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/lbparameter.json). Never settable; populated from GET.
	Sessionsthreshold            types.Int64  `tfsdk:"sessionsthreshold"`
	Builtin                      types.List   `tfsdk:"builtin"`
	Feature                      types.String `tfsdk:"feature"`
	Adccookieattributewarningmsg types.String `tfsdk:"adccookieattributewarningmsg"`
	Lbhashalgowinsize            types.Int64  `tfsdk:"lbhashalgowinsize"`
	Overridepersistencyfororder  types.String `tfsdk:"overridepersistencyfororder"`
}

func LbparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.",
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
			"radiusmessageauthenticator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, NetScaler will verify the message authenticator and also generate message authenticator if not present.",
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
				Computed:    true,
				Description: "Encode persistence cookie values using SHA2 hash.",
			},
			"vserverspecificmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow a MAC-mode virtual server to accept traffic returned by an intermediary device, such as a firewall, to which the traffic was previously forwarded by another MAC-mode virtual server. The second virtual server can then distribute that traffic across the destination server farm. Also useful when load balancing Branch Repeater appliances.\nNote: The second virtual server can also send the traffic to another set of intermediary devices, such as another set of firewalls. If necessary, you can configure multiple MAC-mode virtual servers to pass traffic successively through multiple sets of intermediary devices.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"sessionsthreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Upper-limit on the number of persistent sessions set by the administrator for this system.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the lb parameter configuration is built-in. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]. A list of strings.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
			"adccookieattributewarningmsg": schema.StringAttribute{
				Computed:    true,
				Description: "Describes any configuration issue with respect to the ns variable configured as part of set lb parameter.",
			},
			"lbhashalgowinsize": schema.Int64Attribute{
				Computed:    true,
				Description: "Window size used in the LB hashing algorithm (DEFAULT). Default value: 16.",
			},
			"overridepersistencyfororder": schema.StringAttribute{
				Computed:    true,
				Description: "Whether persistency is overridden when order is configured for services or servicegroups. Possible values: [ YES, NO ]. Default value: NO.",
			},
		},
	}
}

// lbparameterDataSourceSetAttrFromGet projects a NITRO lbparameter GET response
// onto the data-source model. lbparameter is a singleton, so the ID is static.
// Attributes are simply filled from the GET (or left Null when the GET omits
// them) via the shared utils.MapGet* helpers.
func lbparameterDataSourceSetAttrFromGet(ctx context.Context, data *LbparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbparameterDataSourceSetAttrFromGet Function")

	// lbparameter is a singleton -> static ID.
	data.Id = types.StringValue("lbparameter-config")

	// Read/write attributes as read-back outputs.
	data.Allowboundsvcremoval = utils.MapGetString(g, "allowboundsvcremoval")
	data.Computedadccookieattribute = utils.MapGetString(g, "computedadccookieattribute")
	data.Consolidatedlconn = utils.MapGetString(g, "consolidatedlconn")
	data.Dbsttl = utils.MapGetInt64(g, "dbsttl")
	data.Dropmqttjumbomessage = utils.MapGetString(g, "dropmqttjumbomessage")
	data.Httponlycookieflag = utils.MapGetString(g, "httponlycookieflag")
	data.Lbhashalgorithm = utils.MapGetString(g, "lbhashalgorithm")
	data.Lbhashfingers = utils.MapGetInt64(g, "lbhashfingers")
	data.Literaladccookieattribute = utils.MapGetString(g, "literaladccookieattribute")
	data.Maxpipelinenat = utils.MapGetInt64(g, "maxpipelinenat")
	data.Monitorconnectionclose = utils.MapGetString(g, "monitorconnectionclose")
	data.Monitorskipmaxclient = utils.MapGetString(g, "monitorskipmaxclient")
	data.Preferdirectroute = utils.MapGetString(g, "preferdirectroute")
	data.Proximityfromself = utils.MapGetString(g, "proximityfromself")
	data.Radiusmessageauthenticator = utils.MapGetString(g, "radiusmessageauthenticator")
	data.Retainservicestate = utils.MapGetString(g, "retainservicestate")
	data.Startuprrfactor = utils.MapGetInt64(g, "startuprrfactor")
	data.Storemqttclientidandusername = utils.MapGetString(g, "storemqttclientidandusername")
	data.Undefaction = utils.MapGetString(g, "undefaction")
	data.Useencryptedpersistencecookie = utils.MapGetString(g, "useencryptedpersistencecookie")
	data.Useportforhashlb = utils.MapGetString(g, "useportforhashlb")
	data.Usesecuredpersistencecookie = utils.MapGetString(g, "usesecuredpersistencecookie")
	data.Vserverspecificmac = utils.MapGetString(g, "vserverspecificmac")

	// cookiepassphrase is a secret input the GET never returns -> Null.
	data.Cookiepassphrase = types.StringNull()

	// Read-only metadata.
	data.Sessionsthreshold = utils.MapGetInt64(g, "sessionsthreshold")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Adccookieattributewarningmsg = utils.MapGetString(g, "adccookieattributewarningmsg")
	data.Lbhashalgowinsize = utils.MapGetInt64(g, "lbhashalgowinsize")
	data.Overridepersistencyfororder = utils.MapGetString(g, "overridepersistencyfororder")
}
