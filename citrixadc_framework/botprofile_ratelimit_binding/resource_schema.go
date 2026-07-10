package botprofile_ratelimit_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileRatelimitBindingResourceModel describes the resource data model.
type BotprofileRatelimitBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	BotBindComment      types.String `tfsdk:"bot_bind_comment"`
	BotRateLimitAction  types.List   `tfsdk:"bot_rate_limit_action"`
	BotRateLimitEnabled types.String `tfsdk:"bot_rate_limit_enabled"`
	BotRateLimitType    types.String `tfsdk:"bot_rate_limit_type"`
	BotRateLimitUrl     types.String `tfsdk:"bot_rate_limit_url"`
	BotRatelimit        types.Bool   `tfsdk:"bot_ratelimit"`
	Condition           types.String `tfsdk:"condition"`
	Cookiename          types.String `tfsdk:"cookiename"`
	Countrycode         types.String `tfsdk:"countrycode"`
	Limittype           types.String `tfsdk:"limittype"`
	Logmessage          types.String `tfsdk:"logmessage"`
	Name                types.String `tfsdk:"name"`
	Rate                types.Int64  `tfsdk:"rate"`
	Timeslice           types.Int64  `tfsdk:"timeslice"`
}

func (r *BotprofileRatelimitBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_ratelimit_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about this binding.",
			},
			"bot_rate_limit_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "One or more actions to be taken when the current rate becomes more than the configured rate. Only LOG action can be combined with DROP, REDIRECT, RESPOND_STATUS_TOO_MANY_REQUESTS or RESET action.",
			},
			"bot_rate_limit_enabled": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enable or disable rate-limit binding.",
			},
			"bot_rate_limit_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Rate-limiting type Following rate-limiting types are allowed:\n*SOURCE_IP - Rate-limiting based on the client IP.\n*SESSION - Rate-limiting based on the configured cookie name.\n*URL - Rate-limiting based on the configured URL.\n*GEOLOCATION - Rate-limiting based on the configured country name.\n*JA3_FINGERPRINT - Rate-limiting based on client SSL JA3 fingerprint.",
			},
			"bot_rate_limit_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL for the resource based rate-limiting.",
			},
			"bot_ratelimit": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Rate-limit binding. Maximum 30 bindings can be configured per profile for rate-limit detection. For SOURCE_IP type, only one binding can be configured, and for URL type, only one binding is allowed per URL, and for SESSION type, only one binding is allowed for a cookie name. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.",
			},
			"condition": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression to be used in a rate-limiting condition. This expression result must be a boolean value.",
			},
			"cookiename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Cookie name which is used to identify the session for session rate-limiting.",
			},
			"countrycode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Country name which is used for geolocation rate-limiting.",
			},
			"limittype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Rate-Limiting traffic Type",
			},
			"logmessage": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Message to be logged for this binding.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"rate": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Maximum number of requests that are allowed in this session in the given period time.",
			},
			"timeslice": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time interval during which requests are tracked to check if they cross the given rate.",
			},
		},
	}
}

func botprofile_ratelimit_bindingGetThePayloadFromthePlan(ctx context.Context, data *BotprofileRatelimitBindingResourceModel) bot.Botprofileratelimitbinding {
	tflog.Debug(ctx, "In botprofile_ratelimit_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botprofile_ratelimit_binding := bot.Botprofileratelimitbinding{}
	if !data.BotBindComment.IsNull() && !data.BotBindComment.IsUnknown() {
		botprofile_ratelimit_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotRateLimitAction.IsNull() && !data.BotRateLimitAction.IsUnknown() {
		var bot_rate_limit_actionList []string
		data.BotRateLimitAction.ElementsAs(ctx, &bot_rate_limit_actionList, false)
		botprofile_ratelimit_binding.Botratelimitaction = bot_rate_limit_actionList
	}
	if !data.BotRateLimitEnabled.IsNull() && !data.BotRateLimitEnabled.IsUnknown() {
		botprofile_ratelimit_binding.Botratelimitenabled = data.BotRateLimitEnabled.ValueString()
	}
	if !data.BotRateLimitType.IsNull() && !data.BotRateLimitType.IsUnknown() {
		botprofile_ratelimit_binding.Botratelimittype = data.BotRateLimitType.ValueString()
	}
	if !data.BotRateLimitUrl.IsNull() && !data.BotRateLimitUrl.IsUnknown() {
		botprofile_ratelimit_binding.Botratelimiturl = data.BotRateLimitUrl.ValueString()
	}
	if !data.BotRatelimit.IsNull() && !data.BotRatelimit.IsUnknown() {
		botprofile_ratelimit_binding.Botratelimit = data.BotRatelimit.ValueBool()
	}
	if !data.Condition.IsNull() && !data.Condition.IsUnknown() {
		botprofile_ratelimit_binding.Condition = data.Condition.ValueString()
	}
	if !data.Cookiename.IsNull() && !data.Cookiename.IsUnknown() {
		botprofile_ratelimit_binding.Cookiename = data.Cookiename.ValueString()
	}
	if !data.Countrycode.IsNull() && !data.Countrycode.IsUnknown() {
		botprofile_ratelimit_binding.Countrycode = data.Countrycode.ValueString()
	}
	if !data.Limittype.IsNull() && !data.Limittype.IsUnknown() {
		botprofile_ratelimit_binding.Limittype = data.Limittype.ValueString()
	}
	if !data.Logmessage.IsNull() && !data.Logmessage.IsUnknown() {
		botprofile_ratelimit_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botprofile_ratelimit_binding.Name = data.Name.ValueString()
	}
	if !data.Rate.IsNull() && !data.Rate.IsUnknown() {
		botprofile_ratelimit_binding.Rate = utils.IntPtr(int(data.Rate.ValueInt64()))
	}
	if !data.Timeslice.IsNull() && !data.Timeslice.IsUnknown() {
		botprofile_ratelimit_binding.Timeslice = utils.IntPtr(int(data.Timeslice.ValueInt64()))
	}

	return botprofile_ratelimit_binding
}

// botprofile_ratelimit_bindingSetAttrFromGet is the resource-side state setter.
// The live binding GET echoes back every configured field (verified against the ADC),
// so a field that is absent from the response is genuinely unset and is resolved to
// null. All schema attributes are Optional+Computed, so a null is a valid known value
// for the unset ones — this keeps Computed attributes known after apply (Pattern 13)
// while staying faithful to the SDK v2 d.Set(...) behavior. The ID is set once in
// Create (Pattern 6) and is NOT recomputed here.
func botprofile_ratelimit_bindingSetAttrFromGet(ctx context.Context, data *BotprofileRatelimitBindingResourceModel, getResponseData map[string]interface{}) *BotprofileRatelimitBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_ratelimit_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.BotRateLimitAction = listValue
		} else {
			data.BotRateLimitAction = types.ListNull(types.StringType)
		}
	} else {
		data.BotRateLimitAction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["bot_rate_limit_enabled"]; ok && val != nil {
		data.BotRateLimitEnabled = types.StringValue(val.(string))
	} else {
		data.BotRateLimitEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_type"]; ok && val != nil {
		data.BotRateLimitType = types.StringValue(val.(string))
	} else {
		data.BotRateLimitType = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_url"]; ok && val != nil {
		data.BotRateLimitUrl = types.StringValue(val.(string))
	} else {
		data.BotRateLimitUrl = types.StringNull()
	}
	if val, ok := getResponseData["bot_ratelimit"]; ok && val != nil {
		data.BotRatelimit = types.BoolValue(val.(bool))
	} else {
		data.BotRatelimit = types.BoolNull()
	}
	if val, ok := getResponseData["condition"]; ok && val != nil {
		data.Condition = types.StringValue(val.(string))
	} else {
		data.Condition = types.StringNull()
	}
	if val, ok := getResponseData["cookiename"]; ok && val != nil {
		data.Cookiename = types.StringValue(val.(string))
	} else {
		data.Cookiename = types.StringNull()
	}
	if val, ok := getResponseData["countrycode"]; ok && val != nil {
		data.Countrycode = types.StringValue(val.(string))
	} else {
		data.Countrycode = types.StringNull()
	}
	if val, ok := getResponseData["limittype"]; ok && val != nil {
		data.Limittype = types.StringValue(val.(string))
	} else {
		data.Limittype = types.StringNull()
	}
	if val, ok := getResponseData["logmessage"]; ok && val != nil {
		data.Logmessage = types.StringValue(val.(string))
	} else {
		data.Logmessage = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["rate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rate = types.Int64Value(intVal)
		}
	} else {
		data.Rate = types.Int64Null()
	}
	if val, ok := getResponseData["timeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeslice = types.Int64Value(intVal)
		}
	} else {
		data.Timeslice = types.Int64Null()
	}

	return data
}

// botprofile_ratelimit_bindingSetAttrFromGetForDatasource is the datasource-side setter.
// The datasource has no prior plan/state to preserve, so it faithfully copies every field
// from the GET response and explicitly sets the composite ID (datasource has no Create).
func botprofile_ratelimit_bindingSetAttrFromGetForDatasource(ctx context.Context, data *BotprofileRatelimitBindingResourceModel, getResponseData map[string]interface{}) *BotprofileRatelimitBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_ratelimit_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.BotRateLimitAction = listValue
		} else {
			data.BotRateLimitAction = types.ListNull(types.StringType)
		}
	} else {
		data.BotRateLimitAction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["bot_rate_limit_enabled"]; ok && val != nil {
		data.BotRateLimitEnabled = types.StringValue(val.(string))
	} else {
		data.BotRateLimitEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_type"]; ok && val != nil {
		data.BotRateLimitType = types.StringValue(val.(string))
	} else {
		data.BotRateLimitType = types.StringNull()
	}
	if val, ok := getResponseData["bot_rate_limit_url"]; ok && val != nil {
		data.BotRateLimitUrl = types.StringValue(val.(string))
	} else {
		data.BotRateLimitUrl = types.StringNull()
	}
	if val, ok := getResponseData["bot_ratelimit"]; ok && val != nil {
		data.BotRatelimit = types.BoolValue(val.(bool))
	} else {
		data.BotRatelimit = types.BoolNull()
	}
	if val, ok := getResponseData["condition"]; ok && val != nil {
		data.Condition = types.StringValue(val.(string))
	} else {
		data.Condition = types.StringNull()
	}
	if val, ok := getResponseData["cookiename"]; ok && val != nil {
		data.Cookiename = types.StringValue(val.(string))
	} else {
		data.Cookiename = types.StringNull()
	}
	if val, ok := getResponseData["countrycode"]; ok && val != nil {
		data.Countrycode = types.StringValue(val.(string))
	} else {
		data.Countrycode = types.StringNull()
	}
	if val, ok := getResponseData["limittype"]; ok && val != nil {
		data.Limittype = types.StringValue(val.(string))
	} else {
		data.Limittype = types.StringNull()
	}
	if val, ok := getResponseData["logmessage"]; ok && val != nil {
		data.Logmessage = types.StringValue(val.(string))
	} else {
		data.Logmessage = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["rate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rate = types.Int64Value(intVal)
		}
	} else {
		data.Rate = types.Int64Null()
	}
	if val, ok := getResponseData["timeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeslice = types.Int64Value(intVal)
		}
	} else {
		data.Timeslice = types.Int64Null()
	}

	// Set composite ID (datasource has no Create) - matches SDK v2 legacy order name,bot_rate_limit_type
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("bot_rate_limit_type:%s", utils.UrlEncode(data.BotRateLimitType.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
