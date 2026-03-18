package lbprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
