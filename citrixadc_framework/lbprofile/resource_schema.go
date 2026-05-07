package lbprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbprofileResourceModel describes the resource data model.
type LbprofileResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Computedadccookieattribute    types.String `tfsdk:"computedadccookieattribute"`
	Cookiepassphrase              types.String `tfsdk:"cookiepassphrase"`
	CookiepassphraseWo            types.String `tfsdk:"cookiepassphrase_wo"`
	CookiepassphraseWoVersion     types.Int64  `tfsdk:"cookiepassphrase_wo_version"`
	Dbslb                         types.String `tfsdk:"dbslb"`
	Httponlycookieflag            types.String `tfsdk:"httponlycookieflag"`
	Lbhashalgorithm               types.String `tfsdk:"lbhashalgorithm"`
	Lbhashfingers                 types.Int64  `tfsdk:"lbhashfingers"`
	Lbprofilename                 types.String `tfsdk:"lbprofilename"`
	Literaladccookieattribute     types.String `tfsdk:"literaladccookieattribute"`
	Processlocal                  types.String `tfsdk:"processlocal"`
	Proximityfromself             types.String `tfsdk:"proximityfromself"`
	Storemqttclientidandusername  types.String `tfsdk:"storemqttclientidandusername"`
	Useencryptedpersistencecookie types.String `tfsdk:"useencryptedpersistencecookie"`
	Usesecuredpersistencecookie   types.String `tfsdk:"usesecuredpersistencecookie"`
}

func (r *LbprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbprofile resource.",
			},
			"computedadccookieattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.\n\nSample usage -\n             add ns variable lbvar -type TEXT(100) -scope Transaction\n             add ns assignment lbassign -variable $lbvar -set \"\\\\\";SameSite=Strict\\\\\"\"\n             add rewrite policy lbpol <valid policy expression> lbassign\n             bind rewrite global lbpol 100 next -type RES_OVERRIDE\n             add lb profile lbprof -ComputedADCCookieAttribute \"$lbvar\"\n             For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func lbprofileGetThePayloadFromthePlan(ctx context.Context, data *LbprofileResourceModel) lb.Lbprofile {
	tflog.Debug(ctx, "In lbprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbprofile := lb.Lbprofile{}
	if !data.Computedadccookieattribute.IsNull() && !data.Computedadccookieattribute.IsUnknown() {
		lbprofile.Computedadccookieattribute = data.Computedadccookieattribute.ValueString()
	}
	if !data.Cookiepassphrase.IsNull() && !data.Cookiepassphrase.IsUnknown() {
		lbprofile.Cookiepassphrase = data.Cookiepassphrase.ValueString()
	}
	// Skip write-only attribute: cookiepassphrase_wo
	// Skip version tracker attribute: cookiepassphrase_wo_version
	if !data.Dbslb.IsNull() && !data.Dbslb.IsUnknown() {
		lbprofile.Dbslb = data.Dbslb.ValueString()
	}
	if !data.Httponlycookieflag.IsNull() && !data.Httponlycookieflag.IsUnknown() {
		lbprofile.Httponlycookieflag = data.Httponlycookieflag.ValueString()
	}
	if !data.Lbhashalgorithm.IsNull() && !data.Lbhashalgorithm.IsUnknown() {
		lbprofile.Lbhashalgorithm = data.Lbhashalgorithm.ValueString()
	}
	if !data.Lbhashfingers.IsNull() && !data.Lbhashfingers.IsUnknown() {
		lbprofile.Lbhashfingers = utils.IntPtr(int(data.Lbhashfingers.ValueInt64()))
	}
	if !data.Lbprofilename.IsNull() && !data.Lbprofilename.IsUnknown() {
		lbprofile.Lbprofilename = data.Lbprofilename.ValueString()
	}
	if !data.Literaladccookieattribute.IsNull() && !data.Literaladccookieattribute.IsUnknown() {
		lbprofile.Literaladccookieattribute = data.Literaladccookieattribute.ValueString()
	}
	if !data.Processlocal.IsNull() && !data.Processlocal.IsUnknown() {
		lbprofile.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Proximityfromself.IsNull() && !data.Proximityfromself.IsUnknown() {
		lbprofile.Proximityfromself = data.Proximityfromself.ValueString()
	}
	if !data.Storemqttclientidandusername.IsNull() && !data.Storemqttclientidandusername.IsUnknown() {
		lbprofile.Storemqttclientidandusername = data.Storemqttclientidandusername.ValueString()
	}
	if !data.Useencryptedpersistencecookie.IsNull() && !data.Useencryptedpersistencecookie.IsUnknown() {
		lbprofile.Useencryptedpersistencecookie = data.Useencryptedpersistencecookie.ValueString()
	}
	if !data.Usesecuredpersistencecookie.IsNull() && !data.Usesecuredpersistencecookie.IsUnknown() {
		lbprofile.Usesecuredpersistencecookie = data.Usesecuredpersistencecookie.ValueString()
	}

	return lbprofile
}

func lbprofileGetThePayloadFromtheConfig(ctx context.Context, data *LbprofileResourceModel, payload *lb.Lbprofile) {
	tflog.Debug(ctx, "In lbprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: cookiepassphrase_wo -> cookiepassphrase
	if !data.CookiepassphraseWo.IsNull() {
		cookiepassphraseWo := data.CookiepassphraseWo.ValueString()
		if cookiepassphraseWo != "" {
			payload.Cookiepassphrase = cookiepassphraseWo
		}
	}
}

func lbprofileSetAttrFromGet(ctx context.Context, data *LbprofileResourceModel, getResponseData map[string]interface{}) *LbprofileResourceModel {
	tflog.Debug(ctx, "In lbprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["computedadccookieattribute"]; ok && val != nil {
		data.Computedadccookieattribute = types.StringValue(val.(string))
	} else {
		data.Computedadccookieattribute = types.StringNull()
	}
	// cookiepassphrase is not returned by NITRO API (secret/ephemeral) - retain from config
	// cookiepassphrase_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// cookiepassphrase_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["dbslb"]; ok && val != nil {
		data.Dbslb = types.StringValue(val.(string))
	} else {
		data.Dbslb = types.StringNull()
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
	if val, ok := getResponseData["lbprofilename"]; ok && val != nil {
		data.Lbprofilename = types.StringValue(val.(string))
	} else {
		data.Lbprofilename = types.StringNull()
	}
	if val, ok := getResponseData["literaladccookieattribute"]; ok && val != nil {
		data.Literaladccookieattribute = types.StringValue(val.(string))
	} else {
		data.Literaladccookieattribute = types.StringNull()
	}
	if val, ok := getResponseData["processlocal"]; ok && val != nil {
		data.Processlocal = types.StringValue(val.(string))
	} else {
		data.Processlocal = types.StringNull()
	}
	if val, ok := getResponseData["proximityfromself"]; ok && val != nil {
		data.Proximityfromself = types.StringValue(val.(string))
	} else {
		data.Proximityfromself = types.StringNull()
	}
	if val, ok := getResponseData["storemqttclientidandusername"]; ok && val != nil {
		data.Storemqttclientidandusername = types.StringValue(val.(string))
	} else {
		data.Storemqttclientidandusername = types.StringNull()
	}
	if val, ok := getResponseData["useencryptedpersistencecookie"]; ok && val != nil {
		data.Useencryptedpersistencecookie = types.StringValue(val.(string))
	} else {
		data.Useencryptedpersistencecookie = types.StringNull()
	}
	if val, ok := getResponseData["usesecuredpersistencecookie"]; ok && val != nil {
		data.Usesecuredpersistencecookie = types.StringValue(val.(string))
	} else {
		data.Usesecuredpersistencecookie = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Lbprofilename.ValueString()))

	return data
}
