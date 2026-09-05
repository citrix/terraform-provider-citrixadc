package lbprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbprofileDataSourceModel is the data-source-specific model, decoupled from
// LbprofileResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes that the resource deliberately omits.
type LbprofileDataSourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Lbprofilename                 types.String `tfsdk:"lbprofilename"` // Required lookup key
	Computedadccookieattribute    types.String `tfsdk:"computedadccookieattribute"`
	Cookiepassphrase              types.String `tfsdk:"cookiepassphrase"`
	Dbslb                         types.String `tfsdk:"dbslb"`
	Httponlycookieflag            types.String `tfsdk:"httponlycookieflag"`
	Lbhashalgorithm               types.String `tfsdk:"lbhashalgorithm"`
	Lbhashfingers                 types.Int64  `tfsdk:"lbhashfingers"`
	Literaladccookieattribute     types.String `tfsdk:"literaladccookieattribute"`
	Processlocal                  types.String `tfsdk:"processlocal"`
	Proximityfromself             types.String `tfsdk:"proximityfromself"`
	Storemqttclientidandusername  types.String `tfsdk:"storemqttclientidandusername"`
	Useencryptedpersistencecookie types.String `tfsdk:"useencryptedpersistencecookie"`
	Usesecuredpersistencecookie   types.String `tfsdk:"usesecuredpersistencecookie"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lbprofile.json). Never settable; populated from GET.
	Vsvrcount                    types.Int64  `tfsdk:"vsvrcount"`
	Adccookieattributewarningmsg types.String `tfsdk:"adccookieattributewarningmsg"`
	Lbhashalgowinsize            types.Int64  `tfsdk:"lbhashalgowinsize"`
}

func LbprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"computedadccookieattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.\n\nSample usage -\n             add ns variable lbvar -type TEXT(100) -scope Transaction\n             add ns assignment lbassign -variable $lbvar -set \"\\\\\";SameSite=Strict\\\\\"\"\n             add rewrite policy lbpol <valid policy expression> lbassign\n             bind rewrite global lbpol 100 next -type RES_OVERRIDE\n             add lb profile lbprof -ComputedADCCookieAttribute \"$lbvar\"\n             For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie",
			},
			"cookiepassphrase": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.",
			},
			"dbslb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable database specific load balancing for MySQL and MSSQL service types.",
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
				Description: "This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory.",
			},
			"lbprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LB profile.",
			},
			"literaladccookieattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence).\n\nSample usage -\n             add lb profile lbprof -LiteralADCCookieAttribute \";SameSite=None\"",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single pa\ncket request response mode or when the upstream device is performing a proper RSS for connection based distribution.",
			},
			"proximityfromself": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the ADC location instead of client IP for static proximity LB or GSLB decision.",
			},
			"storemqttclientidandusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option allows to store the MQTT clientid and username in transactional logs",
			},
			"useencryptedpersistencecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encode persistence cookie values using SHA2 hash.",
			},
			"usesecuredpersistencecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encode persistence cookie values using SHA2 hash.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"vsvrcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of vservers , the profile is bound to.",
			},
			"adccookieattributewarningmsg": schema.StringAttribute{
				Computed:    true,
				Description: "Used to describe any configuration issue with respect to ns variable configured as part of add/set lb profile.",
			},
			"lbhashalgowinsize": schema.Int64Attribute{
				Computed:    true,
				Description: "This options allows to increase window size used in LB hashing algorithm(DEFAULT).",
			},
		},
	}
}

// lbprofileDataSourceSetAttrFromGet projects a NITRO lbprofile GET response onto
// the data-source model. The shared utils.MapGet* helpers fill each attribute
// from the GET (or leave it Null when the GET omits it).
func lbprofileDataSourceSetAttrFromGet(ctx context.Context, data *LbprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["lbprofilename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Lbprofilename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Computedadccookieattribute = utils.MapGetString(g, "computedadccookieattribute")
	data.Dbslb = utils.MapGetString(g, "dbslb")
	data.Httponlycookieflag = utils.MapGetString(g, "httponlycookieflag")
	data.Lbhashalgorithm = utils.MapGetString(g, "lbhashalgorithm")
	data.Lbhashfingers = utils.MapGetInt64(g, "lbhashfingers")
	data.Literaladccookieattribute = utils.MapGetString(g, "literaladccookieattribute")
	data.Processlocal = utils.MapGetString(g, "processlocal")
	data.Proximityfromself = utils.MapGetString(g, "proximityfromself")
	data.Storemqttclientidandusername = utils.MapGetString(g, "storemqttclientidandusername")
	data.Useencryptedpersistencecookie = utils.MapGetString(g, "useencryptedpersistencecookie")
	data.Usesecuredpersistencecookie = utils.MapGetString(g, "usesecuredpersistencecookie")

	// cookiepassphrase is a secret input the GET never returns -> Null.
	data.Cookiepassphrase = types.StringNull()

	// Read-only attributes.
	data.Vsvrcount = utils.MapGetInt64(g, "vsvrcount")
	data.Adccookieattributewarningmsg = utils.MapGetString(g, "adccookieattributewarningmsg")
	data.Lbhashalgowinsize = utils.MapGetInt64(g, "lbhashalgowinsize")
}
