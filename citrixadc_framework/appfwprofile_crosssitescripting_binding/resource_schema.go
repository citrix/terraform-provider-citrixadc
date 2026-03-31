package appfwprofile_crosssitescripting_binding

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

// AppfwprofileCrosssitescriptingBindingResourceModel describes the resource data model.
type AppfwprofileCrosssitescriptingBindingResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Alertonly          types.String `tfsdk:"alertonly"`
	AsScanLocationXss  types.String `tfsdk:"as_scan_location_xss"`
	AsValueExprXss     types.String `tfsdk:"as_value_expr_xss"`
	AsValueTypeXss     types.String `tfsdk:"as_value_type_xss"`
	Comment            types.String `tfsdk:"comment"`
	Crosssitescripting types.String `tfsdk:"crosssitescripting"`
	FormactionurlXss   types.String `tfsdk:"formactionurl_xss"`
	Isautodeployed     types.String `tfsdk:"isautodeployed"`
	IsregexXss         types.String `tfsdk:"isregex_xss"`
	IsvalueregexXss    types.String `tfsdk:"isvalueregex_xss"`
	Name               types.String `tfsdk:"name"`
	Resourceid         types.String `tfsdk:"resourceid"`
	State              types.String `tfsdk:"state"`
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_crosssitescripting_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_scan_location_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location of cross-site scripting exception - form field, header, cookie or URL.",
			},
			"as_value_expr_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value expression.",
			},
			"as_value_type_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value type.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"crosssitescripting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form field name.",
			},
			"formactionurl_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field name a regular expression?",
			},
			"isvalueregex_xss": schema.StringAttribute{
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
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_crosssitescripting_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel) appfw.Appfwprofilecrosssitescriptingbinding {
	tflog.Debug(ctx, "In appfwprofile_crosssitescripting_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_crosssitescripting_binding := appfw.Appfwprofilecrosssitescriptingbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_crosssitescripting_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Asscanlocationxss = data.AsScanLocationXss.ValueString()
	}
	if !data.AsValueExprXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Asvalueexprxss = data.AsValueExprXss.ValueString()
	}
	if !data.AsValueTypeXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Asvaluetypexss = data.AsValueTypeXss.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_crosssitescripting_binding.Comment = data.Comment.ValueString()
	}
	if !data.Crosssitescripting.IsNull() {
		appfwprofile_crosssitescripting_binding.Crosssitescripting = data.Crosssitescripting.ValueString()
	}
	if !data.FormactionurlXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Formactionurlxss = data.FormactionurlXss.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_crosssitescripting_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Isregexxss = data.IsregexXss.ValueString()
	}
	if !data.IsvalueregexXss.IsNull() {
		appfwprofile_crosssitescripting_binding.Isvalueregexxss = data.IsvalueregexXss.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_crosssitescripting_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_crosssitescripting_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_crosssitescripting_binding.State = data.State.ValueString()
	}

	return appfwprofile_crosssitescripting_binding
}

func appfwprofile_crosssitescripting_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCrosssitescriptingBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_crosssitescripting_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_scan_location_xss"]; ok && val != nil {
		data.AsScanLocationXss = types.StringValue(val.(string))
	} else {
		data.AsScanLocationXss = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_xss"]; ok && val != nil {
		data.AsValueExprXss = types.StringValue(val.(string))
	} else {
		data.AsValueExprXss = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_xss"]; ok && val != nil {
		data.AsValueTypeXss = types.StringValue(val.(string))
	} else {
		data.AsValueTypeXss = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["crosssitescripting"]; ok && val != nil {
		data.Crosssitescripting = types.StringValue(val.(string))
	} else {
		data.Crosssitescripting = types.StringNull()
	}
	if val, ok := getResponseData["formactionurl_xss"]; ok && val != nil {
		data.FormactionurlXss = types.StringValue(val.(string))
	} else {
		data.FormactionurlXss = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_xss"]; ok && val != nil {
		data.IsregexXss = types.StringValue(val.(string))
	} else {
		data.IsregexXss = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_xss"]; ok && val != nil {
		data.IsvalueregexXss = types.StringValue(val.(string))
	} else {
		data.IsvalueregexXss = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("as_scan_location_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crosssitescripting:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Crosssitescripting.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
