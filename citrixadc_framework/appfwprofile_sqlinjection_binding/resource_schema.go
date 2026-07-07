package appfwprofile_sqlinjection_binding

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

// AppfwprofileSqlinjectionBindingResourceModel describes the resource data model.
type AppfwprofileSqlinjectionBindingResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Alertonly         types.String `tfsdk:"alertonly"`
	AsScanLocationSql types.String `tfsdk:"as_scan_location_sql"`
	AsValueExprSql    types.String `tfsdk:"as_value_expr_sql"`
	AsValueTypeSql    types.String `tfsdk:"as_value_type_sql"`
	Comment           types.String `tfsdk:"comment"`
	FormactionurlSql  types.String `tfsdk:"formactionurl_sql"`
	Isautodeployed    types.String `tfsdk:"isautodeployed"`
	IsregexSql        types.String `tfsdk:"isregex_sql"`
	IsvalueregexSql   types.String `tfsdk:"isvalueregex_sql"`
	Name              types.String `tfsdk:"name"`
	Resourceid        types.String `tfsdk:"resourceid"`
	Sqlinjection      types.String `tfsdk:"sqlinjection"`
	State             types.String `tfsdk:"state"`
	Ruletype          types.String `tfsdk:"ruletype"`
}

func (r *AppfwprofileSqlinjectionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_sqlinjection_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_scan_location_sql": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Location of SQL injection exception - form field, header or cookie.",
			},
			"as_value_expr_sql": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form value expression.",
			},
			"as_value_type_sql": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form value type.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"formactionurl_sql": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_sql": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the web form field name a regular expression?",
			},
			"isvalueregex_sql": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the web form field value a regular expression?",
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
			"sqlinjection": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form field name.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
		},
	}
}

func appfwprofile_sqlinjection_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel) appfw.Appfwprofilesqlinjectionbinding {
	tflog.Debug(ctx, "In appfwprofile_sqlinjection_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_sqlinjection_binding := appfw.Appfwprofilesqlinjectionbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_sqlinjection_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationSql.IsNull() && !data.AsScanLocationSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Asscanlocationsql = data.AsScanLocationSql.ValueString()
	}
	if !data.AsValueExprSql.IsNull() && !data.AsValueExprSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Asvalueexprsql = data.AsValueExprSql.ValueString()
	}
	if !data.AsValueTypeSql.IsNull() && !data.AsValueTypeSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Asvaluetypesql = data.AsValueTypeSql.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_sqlinjection_binding.Comment = data.Comment.ValueString()
	}
	if !data.FormactionurlSql.IsNull() && !data.FormactionurlSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Formactionurlsql = data.FormactionurlSql.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_sqlinjection_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexSql.IsNull() && !data.IsregexSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Isregexsql = data.IsregexSql.ValueString()
	}
	if !data.IsvalueregexSql.IsNull() && !data.IsvalueregexSql.IsUnknown() {
		appfwprofile_sqlinjection_binding.Isvalueregexsql = data.IsvalueregexSql.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_sqlinjection_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_sqlinjection_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Sqlinjection.IsNull() && !data.Sqlinjection.IsUnknown() {
		appfwprofile_sqlinjection_binding.Sqlinjection = data.Sqlinjection.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_sqlinjection_binding.State = data.State.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_sqlinjection_binding.Ruletype = data.Ruletype.ValueString()
	}

	return appfwprofile_sqlinjection_binding
}

// appfwprofile_sqlinjection_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly, isautodeployed,
// resourceid, ruletype. To avoid "inconsistent result after apply" we adopt the GET
// value only when the model field is currently null/unknown (e.g. import); otherwise
// we preserve the configured plan/state value. The ID is set once in Create and is
// preserved here.
func appfwprofile_sqlinjection_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileSqlinjectionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_sqlinjection_bindingSetAttrFromGet Function")

	adopt := func(cur types.String, key string) types.String {
		if !cur.IsNull() && !cur.IsUnknown() {
			return cur
		}
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = adopt(data.Alertonly, "alertonly")
	data.AsScanLocationSql = adopt(data.AsScanLocationSql, "as_scan_location_sql")
	data.AsValueExprSql = adopt(data.AsValueExprSql, "as_value_expr_sql")
	data.AsValueTypeSql = adopt(data.AsValueTypeSql, "as_value_type_sql")
	data.Comment = adopt(data.Comment, "comment")
	data.FormactionurlSql = adopt(data.FormactionurlSql, "formactionurl_sql")
	data.Isautodeployed = adopt(data.Isautodeployed, "isautodeployed")
	data.IsregexSql = adopt(data.IsregexSql, "isregex_sql")
	data.IsvalueregexSql = adopt(data.IsvalueregexSql, "isvalueregex_sql")
	data.Name = adopt(data.Name, "name")
	data.Resourceid = adopt(data.Resourceid, "resourceid")
	data.Sqlinjection = adopt(data.Sqlinjection, "sqlinjection")
	data.State = adopt(data.State, "state")
	data.Ruletype = adopt(data.Ruletype, "ruletype")

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sqlinjection:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sqlinjection.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_sqlinjection_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter: it faithfully copies every field from the GET response
// (the datasource has no prior plan/state to preserve) and sets the composite ID.
func appfwprofile_sqlinjection_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileSqlinjectionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_sqlinjection_bindingSetAttrFromGetForDatasource Function")

	copyField := func(key string) types.String {
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = copyField("alertonly")
	data.AsScanLocationSql = copyField("as_scan_location_sql")
	data.AsValueExprSql = copyField("as_value_expr_sql")
	data.AsValueTypeSql = copyField("as_value_type_sql")
	data.Comment = copyField("comment")
	data.FormactionurlSql = copyField("formactionurl_sql")
	data.Isautodeployed = copyField("isautodeployed")
	data.IsregexSql = copyField("isregex_sql")
	data.IsvalueregexSql = copyField("isvalueregex_sql")
	data.Name = copyField("name")
	data.Resourceid = copyField("resourceid")
	data.Sqlinjection = copyField("sqlinjection")
	data.State = copyField("state")
	data.Ruletype = copyField("ruletype")

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sqlinjection:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sqlinjection.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
