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
	Ruletype           types.String `tfsdk:"ruletype"`
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Location of cross-site scripting exception - form field, header, cookie or URL.",
			},
			"as_value_expr_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form value expression.",
			},
			"as_value_type_xss": schema.StringAttribute{
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
			"crosssitescripting": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form field name.",
			},
			"formactionurl_xss": schema.StringAttribute{
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
			"isregex_xss": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the web form field name a regular expression?",
			},
			"isvalueregex_xss": schema.StringAttribute{
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

func appfwprofile_crosssitescripting_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel) appfw.Appfwprofilecrosssitescriptingbinding {
	tflog.Debug(ctx, "In appfwprofile_crosssitescripting_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_crosssitescripting_binding := appfw.Appfwprofilecrosssitescriptingbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationXss.IsNull() && !data.AsScanLocationXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Asscanlocationxss = data.AsScanLocationXss.ValueString()
	}
	if !data.AsValueExprXss.IsNull() && !data.AsValueExprXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Asvalueexprxss = data.AsValueExprXss.ValueString()
	}
	if !data.AsValueTypeXss.IsNull() && !data.AsValueTypeXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Asvaluetypexss = data.AsValueTypeXss.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Comment = data.Comment.ValueString()
	}
	if !data.Crosssitescripting.IsNull() && !data.Crosssitescripting.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Crosssitescripting = data.Crosssitescripting.ValueString()
	}
	if !data.FormactionurlXss.IsNull() && !data.FormactionurlXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Formactionurlxss = data.FormactionurlXss.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexXss.IsNull() && !data.IsregexXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Isregexxss = data.IsregexXss.ValueString()
	}
	if !data.IsvalueregexXss.IsNull() && !data.IsvalueregexXss.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Isvalueregexxss = data.IsvalueregexXss.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_crosssitescripting_binding.State = data.State.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_crosssitescripting_binding.Ruletype = data.Ruletype.ValueString()
	}

	return appfwprofile_crosssitescripting_binding
}

// appfwprofile_crosssitescripting_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly, isautodeployed,
// resourceid, ruletype. To avoid "inconsistent result after apply" we adopt the GET
// value only when the model field is currently null/unknown (e.g. import); otherwise
// we preserve the configured plan/state value. The ID is set once in Create and is
// preserved here.
func appfwprofile_crosssitescripting_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCrosssitescriptingBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_crosssitescripting_bindingSetAttrFromGet Function")

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
	data.AsScanLocationXss = adopt(data.AsScanLocationXss, "as_scan_location_xss")
	data.AsValueExprXss = adopt(data.AsValueExprXss, "as_value_expr_xss")
	data.AsValueTypeXss = adopt(data.AsValueTypeXss, "as_value_type_xss")
	data.Comment = adopt(data.Comment, "comment")
	data.Crosssitescripting = adopt(data.Crosssitescripting, "crosssitescripting")
	data.FormactionurlXss = adopt(data.FormactionurlXss, "formactionurl_xss")
	data.Isautodeployed = adopt(data.Isautodeployed, "isautodeployed")
	data.IsregexXss = adopt(data.IsregexXss, "isregex_xss")
	data.IsvalueregexXss = adopt(data.IsvalueregexXss, "isvalueregex_xss")
	data.Name = adopt(data.Name, "name")
	data.Resourceid = adopt(data.Resourceid, "resourceid")
	data.State = adopt(data.State, "state")
	data.Ruletype = adopt(data.Ruletype, "ruletype")

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
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

// appfwprofile_crosssitescripting_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter: it faithfully copies every field from the GET response
// (the datasource has no prior plan/state to preserve) and sets the composite ID.
func appfwprofile_crosssitescripting_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCrosssitescriptingBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_crosssitescripting_bindingSetAttrFromGetForDatasource Function")

	copyField := func(key string) types.String {
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = copyField("alertonly")
	data.AsScanLocationXss = copyField("as_scan_location_xss")
	data.AsValueExprXss = copyField("as_value_expr_xss")
	data.AsValueTypeXss = copyField("as_value_type_xss")
	data.Comment = copyField("comment")
	data.Crosssitescripting = copyField("crosssitescripting")
	data.FormactionurlXss = copyField("formactionurl_xss")
	data.Isautodeployed = copyField("isautodeployed")
	data.IsregexXss = copyField("isregex_xss")
	data.IsvalueregexXss = copyField("isvalueregex_xss")
	data.Name = copyField("name")
	data.Resourceid = copyField("resourceid")
	data.State = copyField("state")
	data.Ruletype = copyField("ruletype")

	// Set ID for the datasource
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
