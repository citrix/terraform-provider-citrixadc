package authenticationvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/state attributes that the resource
// deliberately omits. Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AuthenticationvserverDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Appflowlog           types.String `tfsdk:"appflowlog"`
	Authentication       types.String `tfsdk:"authentication"`
	Authenticationdomain types.String `tfsdk:"authenticationdomain"`
	Certkeynames         types.String `tfsdk:"certkeynames"`
	Comment              types.String `tfsdk:"comment"`
	Failedlogintimeout   types.Int64  `tfsdk:"failedlogintimeout"`
	Ipv46                types.String `tfsdk:"ipv46"`
	Maxloginattempts     types.Int64  `tfsdk:"maxloginattempts"`
	Name                 types.String `tfsdk:"name"`
	Newname              types.String `tfsdk:"newname"`
	Port                 types.Int64  `tfsdk:"port"`
	Range                types.Int64  `tfsdk:"range"`
	Samesite             types.String `tfsdk:"samesite"`
	Servicetype          types.String `tfsdk:"servicetype"`
	State                types.String `tfsdk:"state"`
	Td                   types.Int64  `tfsdk:"td"`
	Wasmmodule           types.String `tfsdk:"wasmmodule"`

	// Read-only (GET-only) runtime/state attributes from the NITRO doc read-only
	// set (zion73x_readonly/authenticationvserver.json). Never settable; populated
	// from GET.
	Ip                   types.String `tfsdk:"ip"`
	Value                types.String `tfsdk:"value"`
	Type                 types.String `tfsdk:"type"`
	Curstate             types.String `tfsdk:"curstate"`
	Status               types.Int64  `tfsdk:"status"`
	Cachetype            types.String `tfsdk:"cachetype"`
	Redirect             types.String `tfsdk:"redirect"`
	Precedence           types.String `tfsdk:"precedence"`
	Redirecturl          types.String `tfsdk:"redirecturl"`
	Curaaausers          types.Int64  `tfsdk:"curaaausers"`
	Policy               types.String `tfsdk:"policy"`
	Servicename          types.String `tfsdk:"servicename"`
	Weight               types.Int64  `tfsdk:"weight"`
	Cachevserver         types.String `tfsdk:"cachevserver"`
	Backupvserver        types.String `tfsdk:"backupvserver"`
	Clttimeout           types.Int64  `tfsdk:"clttimeout"`
	Somethod             types.String `tfsdk:"somethod"`
	Sothreshold          types.Int64  `tfsdk:"sothreshold"`
	Sopersistence        types.String `tfsdk:"sopersistence"`
	Sopersistencetimeout types.Int64  `tfsdk:"sopersistencetimeout"`
	Priority             types.Int64  `tfsdk:"priority"`
	Downstateflush       types.String `tfsdk:"downstateflush"`
	Bindpoint            types.String `tfsdk:"bindpoint"`
	Disableprimaryondown types.String `tfsdk:"disableprimaryondown"`
	Listenpolicy         types.String `tfsdk:"listenpolicy"`
	Listenpriority       types.Int64  `tfsdk:"listenpriority"`
	Tcpprofilename       types.String `tfsdk:"tcpprofilename"`
	Httpprofilename      types.String `tfsdk:"httpprofilename"`
	Vstype               types.Int64  `tfsdk:"vstype"`
	Ngname               types.String `tfsdk:"ngname"`
	Secondary            types.Bool   `tfsdk:"secondary"`
	Groupextraction      types.Bool   `tfsdk:"groupextraction"`
}

func AuthenticationvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log AppFlow flow information.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Require users to be authenticated before sending traffic through this virtual server.",
			},
			"authenticationdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The domain of the authentication cookie set by Authentication vserver",
			},
			"certkeynames": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with this virtual server.",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum Number of login Attempts",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authentication virtual server.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name of the authentication virtual server.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, 'my authentication policy' or \"my authentication policy\").",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port on which the virtual server accepts connections.",
			},
			"range": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range.\nThe new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol type of the authentication virtual server. Always SSL.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the new virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to assign to this virtual server.",
			},

			// Read-only (GET-only) runtime/state attributes surfaced by the data
			// source (these are intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The Virtual IP address of the authentication vserver.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates whether or not the certificate is bound or if SSL offload is disabled.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of Virtual Server, e.g. CONTENT based or ADDRESS based.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "The current state of the Virtual server, e.g. UP, DOWN, BUSY, etc.",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Whether or not this vserver responds to ARPs and whether or not round-robin selection is temporarily in effect.",
			},
			"cachetype": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server's cache type. The options are: TRANSPARENT, REVERSE and FORWARD.",
			},
			"redirect": schema.StringAttribute{
				Computed:    true,
				Description: "The cache redirect policy.",
			},
			"precedence": schema.StringAttribute{
				Computed:    true,
				Description: "The type of policy (URL or RULE) that takes precedence on the content switching virtual server.",
			},
			"redirecturl": schema.StringAttribute{
				Computed:    true,
				Description: "The URL where traffic is redirected if the virtual server becomes unavailable.",
			},
			"curaaausers": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of current users logged in to this vserver.",
			},
			"policy": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the policy, if any, bound to the authentication vserver.",
			},
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the service, if any, to which the vserver policy is bound.",
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight for this service, if any, used when the system performs load balancing.",
			},
			"cachevserver": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the default target cache virtual server, if any, to which requests are redirected.",
			},
			"backupvserver": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the backup vpn virtual server for this vpn virtual server.",
			},
			"clttimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "The idle time, if any, in seconds after which the client connection is terminated.",
			},
			"somethod": schema.StringAttribute{
				Computed:    true,
				Description: "The method used to determine whether or not a new connection will spillover the allocated block of Intranet IP addresses.",
			},
			"sothreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of client connections after which the Mapped IP address is used as the client source IP address.",
			},
			"sopersistence": schema.StringAttribute{
				Computed:    true,
				Description: "Whether or not cookie-based site persistance is enabled for this VPN vserver.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "The timeout, if any, for cookie-based site persistance of this VPN vserver.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "The priority, if any, of the vpn vserver policy.",
			},
			"downstateflush": schema.StringAttribute{
				Computed:    true,
				Description: "Perform delayed clean up of connections on this vserver.",
			},
			"bindpoint": schema.StringAttribute{
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Computed:    true,
				Description: "Tells whether traffic will continue reaching backup vservers even after primary comes UP from DOWN state.",
			},
			"listenpolicy": schema.StringAttribute{
				Computed:    true,
				Description: "Listenpolicy configured for authentication vserver.",
			},
			"listenpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "Priority of listen policy for authentication vserver.",
			},
			"tcpprofilename": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the TCP profile.",
			},
			"httpprofilename": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the HTTP profile.",
			},
			"vstype": schema.Int64Attribute{
				Computed:    true,
				Description: "Virtual Server Type, e.g. Load Balancing, Content Switch, Cache Redirection.",
			},
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Nodegroup devno to which this authentication vserver belongs to.",
			},
			"secondary": schema.BoolAttribute{
				Computed:    true,
				Description: "Bind the authentication policy to the secondary chain.",
			},
			"groupextraction": schema.BoolAttribute{
				Computed:    true,
				Description: "Bind the Authentication policy to a tertiary chain which will be used only for group extraction.",
			},
		},
	}
}

// authenticationvserverDataSourceSetAttrFromGet projects a NITRO
// authenticationvserver GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them) — no unknown->null resolution or
// plan preservation is required. The shared utils.MapGet* helpers implement that
// projection.
func authenticationvserverDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authenticationdomain = utils.MapGetString(g, "authenticationdomain")
	data.Certkeynames = utils.MapGetString(g, "certkeynames")
	data.Comment = utils.MapGetString(g, "comment")
	data.Failedlogintimeout = utils.MapGetInt64(g, "failedlogintimeout")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.Maxloginattempts = utils.MapGetInt64(g, "maxloginattempts")
	data.Port = utils.MapGetInt64(g, "port")
	data.Range = utils.MapGetInt64(g, "range")
	data.Samesite = utils.MapGetString(g, "samesite")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.State = utils.MapGetString(g, "state")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only runtime/state attributes.
	data.Ip = utils.MapGetString(g, "ip")
	data.Value = utils.MapGetString(g, "value")
	data.Type = utils.MapGetString(g, "type")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Redirect = utils.MapGetString(g, "redirect")
	data.Precedence = utils.MapGetString(g, "precedence")
	data.Redirecturl = utils.MapGetString(g, "redirecturl")
	data.Curaaausers = utils.MapGetInt64(g, "curaaausers")
	data.Policy = utils.MapGetString(g, "policy")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")
	data.Cachevserver = utils.MapGetString(g, "cachevserver")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Vstype = utils.MapGetInt64(g, "vstype")
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Secondary = utils.MapGetBool(g, "secondary")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
}
