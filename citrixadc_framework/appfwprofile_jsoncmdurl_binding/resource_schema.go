package appfwprofile_jsoncmdurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileJsoncmdurlBindingResourceModel describes the resource data model.
type AppfwprofileJsoncmdurlBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Alertonly           types.String `tfsdk:"alertonly"`
	AsValueExprJsonCmd  types.String `tfsdk:"as_value_expr_json_cmd"`
	AsValueTypeJsonCmd  types.String `tfsdk:"as_value_type_json_cmd"`
	Comment             types.String `tfsdk:"comment"`
	Isautodeployed      types.String `tfsdk:"isautodeployed"`
	IskeyregexJsonCmd   types.String `tfsdk:"iskeyregex_json_cmd"`
	IsvalueregexJsonCmd types.String `tfsdk:"isvalueregex_json_cmd"`
	Jsoncmdurl          types.String `tfsdk:"jsoncmdurl"`
	KeynameJsonCmd      types.String `tfsdk:"keyname_json_cmd"`
	Name                types.String `tfsdk:"name"`
	Resourceid          types.String `tfsdk:"resourceid"`
	Ruletype            types.String `tfsdk:"ruletype"`
	State               types.String `tfsdk:"state"`
}

func (r *AppfwprofileJsoncmdurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_jsoncmdurl_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_value_expr_json_cmd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The JSON CMD key value expression.",
			},
			"as_value_type_json_cmd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the relaxed JSON CMD key value",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"iskeyregex_json_cmd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the key name a regular expression?",
			},
			"isvalueregex_json_cmd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the JSON CMD key value a regular expression?",
			},
			"jsoncmdurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A regular expression that designates a URL on the Json CMD URL list for which Command injection violations are relaxed.\nEnclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.",
			},
			"keyname_json_cmd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "An expression that designates a keyname on the JSON CMD URL for which Command injection violations are relaxed.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_jsoncmdurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileJsoncmdurlBindingResourceModel) appfw.Appfwprofilejsoncmdurlbinding {
	tflog.Debug(ctx, "In appfwprofile_jsoncmdurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_jsoncmdurl_binding := appfw.Appfwprofilejsoncmdurlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsValueExprJsonCmd.IsNull() && !data.AsValueExprJsonCmd.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Asvalueexprjsoncmd = data.AsValueExprJsonCmd.ValueString()
	}
	if !data.AsValueTypeJsonCmd.IsNull() && !data.AsValueTypeJsonCmd.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Asvaluetypejsoncmd = data.AsValueTypeJsonCmd.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IskeyregexJsonCmd.IsNull() && !data.IskeyregexJsonCmd.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Iskeyregexjsoncmd = data.IskeyregexJsonCmd.ValueString()
	}
	if !data.IsvalueregexJsonCmd.IsNull() && !data.IsvalueregexJsonCmd.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Isvalueregexjsoncmd = data.IsvalueregexJsonCmd.ValueString()
	}
	if !data.Jsoncmdurl.IsNull() && !data.Jsoncmdurl.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Jsoncmdurl = data.Jsoncmdurl.ValueString()
	}
	if !data.KeynameJsonCmd.IsNull() && !data.KeynameJsonCmd.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Keynamejsoncmd = data.KeynameJsonCmd.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_jsoncmdurl_binding.State = data.State.ValueString()
	}

	return appfwprofile_jsoncmdurl_binding
}

// appfwprofile_jsoncmdurl_bindingSetAttrFromGet is the RESOURCE-side state setter.
// It preserves the user-configured values for server-overridden / non-echoed inputs
// (alertonly, isautodeployed) — mirroring the SDK v2 Read which deliberately did NOT
// copy these back (the ADC normalizes alertonly ON->OFF, etc.), avoiding an
// "inconsistent result after apply" error. It does NOT recompute the ID; the ID is
// set once in Create and parsed from state in Read.
func appfwprofile_jsoncmdurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileJsoncmdurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsoncmdurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsoncmdurl_bindingSetAttrFromGet Function")

	// alertonly and isautodeployed are intentionally NOT set from GET (server overrides
	// the configured value) — preserve the existing plan/state value.
	//
	// For all other (Optional+Computed) attributes, set the GET value when present,
	// otherwise set null — a Computed attribute must never be left unknown after apply.

	if val, ok := getResponseData["as_value_expr_json_cmd"]; ok && val != nil {
		data.AsValueExprJsonCmd = types.StringValue(val.(string))
	} else {
		data.AsValueExprJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_json_cmd"]; ok && val != nil {
		data.AsValueTypeJsonCmd = types.StringValue(val.(string))
	} else {
		data.AsValueTypeJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["iskeyregex_json_cmd"]; ok && val != nil {
		data.IskeyregexJsonCmd = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_json_cmd"]; ok && val != nil {
		data.IsvalueregexJsonCmd = types.StringValue(val.(string))
	} else {
		data.IsvalueregexJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["jsoncmdurl"]; ok && val != nil {
		data.Jsoncmdurl = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["keyname_json_cmd"]; ok && val != nil {
		data.KeynameJsonCmd = types.StringValue(val.(string))
	} else {
		data.KeynameJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(buildAppfwprofileJsoncmdurlBindingId(data))

	return data
}

// appfwprofile_jsoncmdurl_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior plan/state to preserve, so it faithfully copies
// every field from the GET response (including server-overridden alertonly /
// isautodeployed) and sets its own ID.
func appfwprofile_jsoncmdurl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileJsoncmdurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsoncmdurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsoncmdurl_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_json_cmd"]; ok && val != nil {
		data.AsValueExprJsonCmd = types.StringValue(val.(string))
	} else {
		data.AsValueExprJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_json_cmd"]; ok && val != nil {
		data.AsValueTypeJsonCmd = types.StringValue(val.(string))
	} else {
		data.AsValueTypeJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["iskeyregex_json_cmd"]; ok && val != nil {
		data.IskeyregexJsonCmd = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_json_cmd"]; ok && val != nil {
		data.IsvalueregexJsonCmd = types.StringValue(val.(string))
	} else {
		data.IsvalueregexJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["jsoncmdurl"]; ok && val != nil {
		data.Jsoncmdurl = types.StringValue(val.(string))
	} else {
		data.Jsoncmdurl = types.StringNull()
	}
	if val, ok := getResponseData["keyname_json_cmd"]; ok && val != nil {
		data.KeynameJsonCmd = types.StringValue(val.(string))
	} else {
		data.KeynameJsonCmd = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the datasource (no Create to set it)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_value_expr_json_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprJsonCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_json_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeJsonCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsoncmdurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsoncmdurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("keyname_json_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
