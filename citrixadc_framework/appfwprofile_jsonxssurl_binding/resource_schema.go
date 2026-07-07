package appfwprofile_jsonxssurl_binding

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

// AppfwprofileJsonxssurlBindingResourceModel describes the resource data model.
type AppfwprofileJsonxssurlBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Alertonly           types.String `tfsdk:"alertonly"`
	AsValueExprJsonXss  types.String `tfsdk:"as_value_expr_json_xss"`
	AsValueTypeJsonXss  types.String `tfsdk:"as_value_type_json_xss"`
	Comment             types.String `tfsdk:"comment"`
	Isautodeployed      types.String `tfsdk:"isautodeployed"`
	IskeyregexJsonXss   types.String `tfsdk:"iskeyregex_json_xss"`
	IsvalueregexJsonXss types.String `tfsdk:"isvalueregex_json_xss"`
	Jsonxssurl          types.String `tfsdk:"jsonxssurl"`
	KeynameJsonXss      types.String `tfsdk:"keyname_json_xss"`
	Name                types.String `tfsdk:"name"`
	Resourceid          types.String `tfsdk:"resourceid"`
	Ruletype            types.String `tfsdk:"ruletype"`
	State               types.String `tfsdk:"state"`
}

func (r *AppfwprofileJsonxssurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_jsonxssurl_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_value_expr_json_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The JSON XSS key value expression.",
			},
			"as_value_type_json_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the relaxed JSON XSS key value",
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"iskeyregex_json_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the key name a regular expression?",
			},
			"isvalueregex_json_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the JSON XSS key value a regular expression?",
			},
			"jsonxssurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A regular expression that designates a URL on the Json XSS URL list for which XSS violations are relaxed.\nEnclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.",
			},
			"keyname_json_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "An expression that designates a keyname on the JSON XSS URL for which XSS injection violations are relaxed.",
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
				Description: "Specifies rule type of binding",
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

func appfwprofile_jsonxssurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileJsonxssurlBindingResourceModel) appfw.Appfwprofilejsonxssurlbinding {
	tflog.Debug(ctx, "In appfwprofile_jsonxssurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_jsonxssurl_binding := appfw.Appfwprofilejsonxssurlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsValueExprJsonXss.IsNull() && !data.AsValueExprJsonXss.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Asvalueexprjsonxss = data.AsValueExprJsonXss.ValueString()
	}
	if !data.AsValueTypeJsonXss.IsNull() && !data.AsValueTypeJsonXss.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Asvaluetypejsonxss = data.AsValueTypeJsonXss.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IskeyregexJsonXss.IsNull() && !data.IskeyregexJsonXss.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Iskeyregexjsonxss = data.IskeyregexJsonXss.ValueString()
	}
	if !data.IsvalueregexJsonXss.IsNull() && !data.IsvalueregexJsonXss.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Isvalueregexjsonxss = data.IsvalueregexJsonXss.ValueString()
	}
	if !data.Jsonxssurl.IsNull() && !data.Jsonxssurl.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Jsonxssurl = data.Jsonxssurl.ValueString()
	}
	if !data.KeynameJsonXss.IsNull() && !data.KeynameJsonXss.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Keynamejsonxss = data.KeynameJsonXss.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_jsonxssurl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_jsonxssurl_binding.State = data.State.ValueString()
	}

	return appfwprofile_jsonxssurl_binding
}

// appfwprofile_jsonxssurl_bindingSetAttrFromGet is the RESOURCE-side setter.
// It preserves user-supplied values for attributes the NITRO server overrides on
// GET (alertonly is returned as OFF, isautodeployed echoes a server-managed value)
// so that the post-apply state matches the user's config (mirrors SDK v2 Read which
// deliberately did NOT d.Set those two). It does NOT recompute the ID; the ID is set
// exactly once in Create / preserved in Read from prior state.
func appfwprofile_jsonxssurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileJsonxssurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsonxssurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsonxssurl_bindingSetAttrFromGet Function")

	// alertonly and isautodeployed are server-overridden on GET - preserve plan/state
	// value (do not copy from getResponseData), matching SDK v2 backward-compat behavior.
	if val, ok := getResponseData["as_value_expr_json_xss"]; ok && val != nil {
		data.AsValueExprJsonXss = types.StringValue(val.(string))
	} else {
		data.AsValueExprJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_json_xss"]; ok && val != nil {
		data.AsValueTypeJsonXss = types.StringValue(val.(string))
	} else {
		data.AsValueTypeJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["iskeyregex_json_xss"]; ok && val != nil {
		data.IskeyregexJsonXss = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_json_xss"]; ok && val != nil {
		data.IsvalueregexJsonXss = types.StringValue(val.(string))
	} else {
		data.IsvalueregexJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["jsonxssurl"]; ok && val != nil {
		data.Jsonxssurl = types.StringValue(val.(string))
	} else {
		data.Jsonxssurl = types.StringNull()
	}
	if val, ok := getResponseData["keyname_json_xss"]; ok && val != nil {
		data.KeynameJsonXss = types.StringValue(val.(string))
	} else {
		data.KeynameJsonXss = types.StringNull()
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

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonxssurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonxssurl.ValueString()))))
	if !data.KeynameJsonXss.IsNull() && data.KeynameJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("keyname_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonXss.ValueString()))))
	}
	if !data.AsValueTypeJsonXss.IsNull() && data.AsValueTypeJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_type_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeJsonXss.ValueString()))))
	}
	if !data.AsValueExprJsonXss.IsNull() && data.AsValueExprJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_expr_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprJsonXss.ValueString()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_jsonxssurl_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior plan/state to preserve, so it faithfully copies
// every attribute (including the server-overridden alertonly/isautodeployed) from the
// GET response and sets the composite ID itself (no Create runs for a datasource).
func appfwprofile_jsonxssurl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileJsonxssurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsonxssurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsonxssurl_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_json_xss"]; ok && val != nil {
		data.AsValueExprJsonXss = types.StringValue(val.(string))
	} else {
		data.AsValueExprJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_json_xss"]; ok && val != nil {
		data.AsValueTypeJsonXss = types.StringValue(val.(string))
	} else {
		data.AsValueTypeJsonXss = types.StringNull()
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
	if val, ok := getResponseData["iskeyregex_json_xss"]; ok && val != nil {
		data.IskeyregexJsonXss = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_json_xss"]; ok && val != nil {
		data.IsvalueregexJsonXss = types.StringValue(val.(string))
	} else {
		data.IsvalueregexJsonXss = types.StringNull()
	}
	if val, ok := getResponseData["jsonxssurl"]; ok && val != nil {
		data.Jsonxssurl = types.StringValue(val.(string))
	} else {
		data.Jsonxssurl = types.StringNull()
	}
	if val, ok := getResponseData["keyname_json_xss"]; ok && val != nil {
		data.KeynameJsonXss = types.StringValue(val.(string))
	} else {
		data.KeynameJsonXss = types.StringNull()
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

	// Set composite ID (datasource has no Create). Order mirrors resource_id_mapping.json:
	// name, jsonxssurl, then the optional keys that are actually present in the response.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonxssurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonxssurl.ValueString()))))
	if !data.KeynameJsonXss.IsNull() && data.KeynameJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("keyname_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonXss.ValueString()))))
	}
	if !data.AsValueTypeJsonXss.IsNull() && data.AsValueTypeJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_type_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeJsonXss.ValueString()))))
	}
	if !data.AsValueExprJsonXss.IsNull() && data.AsValueExprJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_expr_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprJsonXss.ValueString()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
