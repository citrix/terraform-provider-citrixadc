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
				Optional:    true,
				Computed:    true,
				Description: "Location of SQL injection exception - form field, header or cookie.",
			},
			"as_value_expr_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value expression.",
			},
			"as_value_type_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value type.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"formactionurl_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field name a regular expression?",
			},
			"isvalueregex_sql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field value a regular expression?",
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
			"sqlinjection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form field name.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
			"ruletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies rule type of binding.",
			},
		},
	}
}

func appfwprofile_sqlinjection_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel) appfw.Appfwprofilesqlinjectionbinding {
	tflog.Debug(ctx, "In appfwprofile_sqlinjection_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_sqlinjection_binding := appfw.Appfwprofilesqlinjectionbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_sqlinjection_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationSql.IsNull() {
		appfwprofile_sqlinjection_binding.Asscanlocationsql = data.AsScanLocationSql.ValueString()
	}
	if !data.AsValueExprSql.IsNull() {
		appfwprofile_sqlinjection_binding.Asvalueexprsql = data.AsValueExprSql.ValueString()
	}
	if !data.AsValueTypeSql.IsNull() {
		appfwprofile_sqlinjection_binding.Asvaluetypesql = data.AsValueTypeSql.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_sqlinjection_binding.Comment = data.Comment.ValueString()
	}
	if !data.FormactionurlSql.IsNull() {
		appfwprofile_sqlinjection_binding.Formactionurlsql = data.FormactionurlSql.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_sqlinjection_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexSql.IsNull() {
		appfwprofile_sqlinjection_binding.Isregexsql = data.IsregexSql.ValueString()
	}
	if !data.IsvalueregexSql.IsNull() {
		appfwprofile_sqlinjection_binding.Isvalueregexsql = data.IsvalueregexSql.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_sqlinjection_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_sqlinjection_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Sqlinjection.IsNull() {
		appfwprofile_sqlinjection_binding.Sqlinjection = data.Sqlinjection.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_sqlinjection_binding.State = data.State.ValueString()
	}
	if !data.Ruletype.IsNull() {
		appfwprofile_sqlinjection_binding.Ruletype = data.Ruletype.ValueString()
	}

	return appfwprofile_sqlinjection_binding
}

func appfwprofile_sqlinjection_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileSqlinjectionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_sqlinjection_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_scan_location_sql"]; ok && val != nil {
		data.AsScanLocationSql = types.StringValue(val.(string))
	} else {
		data.AsScanLocationSql = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_sql"]; ok && val != nil {
		data.AsValueExprSql = types.StringValue(val.(string))
	} else {
		data.AsValueExprSql = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_sql"]; ok && val != nil {
		data.AsValueTypeSql = types.StringValue(val.(string))
	} else {
		data.AsValueTypeSql = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["formactionurl_sql"]; ok && val != nil {
		data.FormactionurlSql = types.StringValue(val.(string))
	} else {
		data.FormactionurlSql = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_sql"]; ok && val != nil {
		data.IsregexSql = types.StringValue(val.(string))
	} else {
		data.IsregexSql = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_sql"]; ok && val != nil {
		data.IsvalueregexSql = types.StringValue(val.(string))
	} else {
		data.IsvalueregexSql = types.StringNull()
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
	if val, ok := getResponseData["sqlinjection"]; ok && val != nil {
		data.Sqlinjection = types.StringValue(val.(string))
	} else {
		data.Sqlinjection = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_sql:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsScanLocationSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_sql:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsValueExprSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_sql:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsValueTypeSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_sql:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.FormactionurlSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ruletype:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ruletype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sqlinjection:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Sqlinjection.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
