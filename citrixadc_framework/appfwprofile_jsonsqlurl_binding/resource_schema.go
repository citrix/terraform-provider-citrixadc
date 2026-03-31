package appfwprofile_jsonsqlurl_binding

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

// AppfwprofileJsonsqlurlBindingResourceModel describes the resource data model.
type AppfwprofileJsonsqlurlBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Alertonly           types.String `tfsdk:"alertonly"`
	AsValueExprJsonSql  types.String `tfsdk:"as_value_expr_json_sql"`
	AsValueTypeJsonSql  types.String `tfsdk:"as_value_type_json_sql"`
	Comment             types.String `tfsdk:"comment"`
	Isautodeployed      types.String `tfsdk:"isautodeployed"`
	IskeyregexJsonSql   types.String `tfsdk:"iskeyregex_json_sql"`
	IsvalueregexJsonSql types.String `tfsdk:"isvalueregex_json_sql"`
	Jsonsqlurl          types.String `tfsdk:"jsonsqlurl"`
	KeynameJsonSql      types.String `tfsdk:"keyname_json_sql"`
	Name                types.String `tfsdk:"name"`
	Resourceid          types.String `tfsdk:"resourceid"`
	State               types.String `tfsdk:"state"`
}

func (r *AppfwprofileJsonsqlurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_jsonsqlurl_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_value_expr_json_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The JSON SQL key value expression.",
			},
			"as_value_type_json_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the relaxed JSON SQL key value",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"iskeyregex_json_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the key name a regular expression?",
			},
			"isvalueregex_json_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the JSON SQL key value a regular expression?",
			},
			"jsonsqlurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A regular expression that designates a URL on the Json SQL URL list for which SQL violations are relaxed.\nEnclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.",
			},
			"keyname_json_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An expression that designates a keyname on the JSON SQL URL for which SQL injection violations are relaxed.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_jsonsqlurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileJsonsqlurlBindingResourceModel) appfw.Appfwprofilejsonsqlurlbinding {
	tflog.Debug(ctx, "In appfwprofile_jsonsqlurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_jsonsqlurl_binding := appfw.Appfwprofilejsonsqlurlbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_jsonsqlurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsValueExprJsonSql.IsNull() {
		appfwprofile_jsonsqlurl_binding.Asvalueexprjsonsql = data.AsValueExprJsonSql.ValueString()
	}
	if !data.AsValueTypeJsonSql.IsNull() {
		appfwprofile_jsonsqlurl_binding.Asvaluetypejsonsql = data.AsValueTypeJsonSql.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_jsonsqlurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_jsonsqlurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IskeyregexJsonSql.IsNull() {
		appfwprofile_jsonsqlurl_binding.Iskeyregexjsonsql = data.IskeyregexJsonSql.ValueString()
	}
	if !data.IsvalueregexJsonSql.IsNull() {
		appfwprofile_jsonsqlurl_binding.Isvalueregexjsonsql = data.IsvalueregexJsonSql.ValueString()
	}
	if !data.Jsonsqlurl.IsNull() {
		appfwprofile_jsonsqlurl_binding.Jsonsqlurl = data.Jsonsqlurl.ValueString()
	}
	if !data.KeynameJsonSql.IsNull() {
		appfwprofile_jsonsqlurl_binding.Keynamejsonsql = data.KeynameJsonSql.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_jsonsqlurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_jsonsqlurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_jsonsqlurl_binding.State = data.State.ValueString()
	}

	return appfwprofile_jsonsqlurl_binding
}

func appfwprofile_jsonsqlurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileJsonsqlurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsonsqlurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsonsqlurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_json_sql"]; ok && val != nil {
		data.AsValueExprJsonSql = types.StringValue(val.(string))
	} else {
		data.AsValueExprJsonSql = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_json_sql"]; ok && val != nil {
		data.AsValueTypeJsonSql = types.StringValue(val.(string))
	} else {
		data.AsValueTypeJsonSql = types.StringNull()
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
	if val, ok := getResponseData["iskeyregex_json_sql"]; ok && val != nil {
		data.IskeyregexJsonSql = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonSql = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_json_sql"]; ok && val != nil {
		data.IsvalueregexJsonSql = types.StringValue(val.(string))
	} else {
		data.IsvalueregexJsonSql = types.StringNull()
	}
	if val, ok := getResponseData["jsonsqlurl"]; ok && val != nil {
		data.Jsonsqlurl = types.StringValue(val.(string))
	} else {
		data.Jsonsqlurl = types.StringNull()
	}
	if val, ok := getResponseData["keyname_json_sql"]; ok && val != nil {
		data.KeynameJsonSql = types.StringValue(val.(string))
	} else {
		data.KeynameJsonSql = types.StringNull()
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
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_value_expr_json_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprJsonSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_json_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeJsonSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonsqlurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonsqlurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("keyname_json_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
