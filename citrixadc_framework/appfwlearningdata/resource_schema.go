package appfwlearningdata

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwlearningdataResourceModel describes the resource data model.
//
// appfwlearningdata is the Application-Firewall learned-data table. It has NO
// clean CRUD lifecycle: NITRO exposes only get(all), count, delete, and the
// reset/export actions. This is a BEST-EFFORT ACTION model — Create performs the
// "reset" action (clears learned data for the given profile/security check),
// Read/Update are no-ops, and Delete is a state-only removal. The reset/delete
// semantics should be verified on a live ADC before relying on them.
//
// The attributes below are the read/write inputs to the reset/export actions.
type AppfwlearningdataResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Profilename         types.String `tfsdk:"profilename"`
	Starturl            types.String `tfsdk:"starturl"`
	Cookieconsistency   types.String `tfsdk:"cookieconsistency"`
	Fieldconsistency    types.String `tfsdk:"fieldconsistency"`
	FormactionurlFfc    types.String `tfsdk:"formactionurl_ffc"`
	Contenttype         types.String `tfsdk:"contenttype"`
	Crosssitescripting  types.String `tfsdk:"crosssitescripting"`
	FormactionurlXss    types.String `tfsdk:"formactionurl_xss"`
	AsScanLocationXss   types.String `tfsdk:"as_scan_location_xss"`
	AsValueTypeXss      types.String `tfsdk:"as_value_type_xss"`
	AsValueExprXss      types.String `tfsdk:"as_value_expr_xss"`
	Sqlinjection        types.String `tfsdk:"sqlinjection"`
	FormactionurlSql    types.String `tfsdk:"formactionurl_sql"`
	AsScanLocationSql   types.String `tfsdk:"as_scan_location_sql"`
	AsValueTypeSql      types.String `tfsdk:"as_value_type_sql"`
	AsValueExprSql      types.String `tfsdk:"as_value_expr_sql"`
	Fieldformat         types.String `tfsdk:"fieldformat"`
	FormactionurlFf     types.String `tfsdk:"formactionurl_ff"`
	Csrftag             types.String `tfsdk:"csrftag"`
	Csrfformoriginurl   types.String `tfsdk:"csrfformoriginurl"`
	Creditcardnumber    types.String `tfsdk:"creditcardnumber"`
	Creditcardnumberurl types.String `tfsdk:"creditcardnumberurl"`
	Xmldoscheck         types.String `tfsdk:"xmldoscheck"`
	Xmlwsicheck         types.String `tfsdk:"xmlwsicheck"`
	Xmlattachmentcheck  types.String `tfsdk:"xmlattachmentcheck"`
	Totalxmlrequests    types.Bool   `tfsdk:"totalxmlrequests"`
	Securitycheck       types.String `tfsdk:"securitycheck"`
	Target              types.String `tfsdk:"target"`
}

func (r *AppfwlearningdataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	stringReplace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwlearningdata resource.",
			},
			"profilename": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Name of the profile.",
			},
			"starturl": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Start URL configuration.",
			},
			"cookieconsistency": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Cookie Name.",
			},
			"fieldconsistency": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form field name.",
			},
			"formactionurl_ffc": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form action URL.",
			},
			"contenttype": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Content Type Name.",
			},
			"crosssitescripting": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Cross-site scripting.",
			},
			"formactionurl_xss": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form action URL.",
			},
			"as_scan_location_xss": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Location of cross-site scripting exception - form field, header, cookie or url. Possible values = FORMFIELD, HEADER, COOKIE, URL.",
			},
			"as_value_type_xss": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "XSS value type. Possible values = Tag, Attribute, Pattern.",
			},
			"as_value_expr_xss": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "XSS value expressions consistituting expressions for Tag, Attribute or Pattern.",
			},
			"sqlinjection": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form field name.",
			},
			"formactionurl_sql": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form action URL.",
			},
			"as_scan_location_sql": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Location of sql injection exception - form field, header or cookie. Possible values = FORMFIELD, HEADER, COOKIE.",
			},
			"as_value_type_sql": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "SQL value type. Possible values = Keyword, SpecialString, Wildchar.",
			},
			"as_value_expr_sql": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "SQL value expressions consistituting expressions for Keyword, SpecialString or Wildchar.",
			},
			"fieldformat": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Field format name.",
			},
			"formactionurl_ff": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Form action URL.",
			},
			"csrftag": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "CSRF Form Action URL.",
			},
			"csrfformoriginurl": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "CSRF Form Origin URL.",
			},
			"creditcardnumber": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "The object expression that is to be excluded from safe commerce check.",
			},
			"creditcardnumberurl": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "The url for which the list of credit card numbers are needed to be bypassed from inspection.",
			},
			"xmldoscheck": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "XML Denial of Service check.",
			},
			"xmlwsicheck": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Web Services Interoperability Rule ID.",
			},
			"xmlattachmentcheck": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "XML Attachment Content-Type.",
			},
			"totalxmlrequests": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Total XML requests.",
			},
			"securitycheck": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Name of the security check. Possible values = startURL, cookieConsistency, fieldConsistency, crossSiteScripting, SQLInjection, fieldFormat, CSRFtag, XMLDoSCheck, XMLWSICheck, XMLAttachmentCheck, TotalXMLRequests, creditCardNumber, ContentType.",
			},
			"target": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Target filename for data to be exported.",
			},
		},
	}
}

// appfwlearningdataGetThePayloadFromthePlan builds the NITRO payload used for the
// reset/export actions from the Terraform model.
func appfwlearningdataGetThePayloadFromthePlan(ctx context.Context, data *AppfwlearningdataResourceModel) appfw.Appfwlearningdata {
	tflog.Debug(ctx, "In appfwlearningdataGetThePayloadFromthePlan Function")

	appfwlearningdata := appfw.Appfwlearningdata{}
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		appfwlearningdata.Profilename = data.Profilename.ValueString()
	}
	if !data.Starturl.IsNull() && !data.Starturl.IsUnknown() {
		appfwlearningdata.Starturl = data.Starturl.ValueString()
	}
	if !data.Cookieconsistency.IsNull() && !data.Cookieconsistency.IsUnknown() {
		appfwlearningdata.Cookieconsistency = data.Cookieconsistency.ValueString()
	}
	if !data.Fieldconsistency.IsNull() && !data.Fieldconsistency.IsUnknown() {
		appfwlearningdata.Fieldconsistency = data.Fieldconsistency.ValueString()
	}
	if !data.FormactionurlFfc.IsNull() && !data.FormactionurlFfc.IsUnknown() {
		appfwlearningdata.Formactionurlffc = data.FormactionurlFfc.ValueString()
	}
	if !data.Contenttype.IsNull() && !data.Contenttype.IsUnknown() {
		appfwlearningdata.Contenttype = data.Contenttype.ValueString()
	}
	if !data.Crosssitescripting.IsNull() && !data.Crosssitescripting.IsUnknown() {
		appfwlearningdata.Crosssitescripting = data.Crosssitescripting.ValueString()
	}
	if !data.FormactionurlXss.IsNull() && !data.FormactionurlXss.IsUnknown() {
		appfwlearningdata.Formactionurlxss = data.FormactionurlXss.ValueString()
	}
	if !data.AsScanLocationXss.IsNull() && !data.AsScanLocationXss.IsUnknown() {
		appfwlearningdata.Asscanlocationxss = data.AsScanLocationXss.ValueString()
	}
	if !data.AsValueTypeXss.IsNull() && !data.AsValueTypeXss.IsUnknown() {
		appfwlearningdata.Asvaluetypexss = data.AsValueTypeXss.ValueString()
	}
	if !data.AsValueExprXss.IsNull() && !data.AsValueExprXss.IsUnknown() {
		appfwlearningdata.Asvalueexprxss = data.AsValueExprXss.ValueString()
	}
	if !data.Sqlinjection.IsNull() && !data.Sqlinjection.IsUnknown() {
		appfwlearningdata.Sqlinjection = data.Sqlinjection.ValueString()
	}
	if !data.FormactionurlSql.IsNull() && !data.FormactionurlSql.IsUnknown() {
		appfwlearningdata.Formactionurlsql = data.FormactionurlSql.ValueString()
	}
	if !data.AsScanLocationSql.IsNull() && !data.AsScanLocationSql.IsUnknown() {
		appfwlearningdata.Asscanlocationsql = data.AsScanLocationSql.ValueString()
	}
	if !data.AsValueTypeSql.IsNull() && !data.AsValueTypeSql.IsUnknown() {
		appfwlearningdata.Asvaluetypesql = data.AsValueTypeSql.ValueString()
	}
	if !data.AsValueExprSql.IsNull() && !data.AsValueExprSql.IsUnknown() {
		appfwlearningdata.Asvalueexprsql = data.AsValueExprSql.ValueString()
	}
	if !data.Fieldformat.IsNull() && !data.Fieldformat.IsUnknown() {
		appfwlearningdata.Fieldformat = data.Fieldformat.ValueString()
	}
	if !data.FormactionurlFf.IsNull() && !data.FormactionurlFf.IsUnknown() {
		appfwlearningdata.Formactionurlff = data.FormactionurlFf.ValueString()
	}
	if !data.Csrftag.IsNull() && !data.Csrftag.IsUnknown() {
		appfwlearningdata.Csrftag = data.Csrftag.ValueString()
	}
	if !data.Csrfformoriginurl.IsNull() && !data.Csrfformoriginurl.IsUnknown() {
		appfwlearningdata.Csrfformoriginurl = data.Csrfformoriginurl.ValueString()
	}
	if !data.Creditcardnumber.IsNull() && !data.Creditcardnumber.IsUnknown() {
		appfwlearningdata.Creditcardnumber = data.Creditcardnumber.ValueString()
	}
	if !data.Creditcardnumberurl.IsNull() && !data.Creditcardnumberurl.IsUnknown() {
		appfwlearningdata.Creditcardnumberurl = data.Creditcardnumberurl.ValueString()
	}
	if !data.Xmldoscheck.IsNull() && !data.Xmldoscheck.IsUnknown() {
		appfwlearningdata.Xmldoscheck = data.Xmldoscheck.ValueString()
	}
	if !data.Xmlwsicheck.IsNull() && !data.Xmlwsicheck.IsUnknown() {
		appfwlearningdata.Xmlwsicheck = data.Xmlwsicheck.ValueString()
	}
	if !data.Xmlattachmentcheck.IsNull() && !data.Xmlattachmentcheck.IsUnknown() {
		appfwlearningdata.Xmlattachmentcheck = data.Xmlattachmentcheck.ValueString()
	}
	if !data.Totalxmlrequests.IsNull() && !data.Totalxmlrequests.IsUnknown() {
		appfwlearningdata.Totalxmlrequests = data.Totalxmlrequests.ValueBool()
	}
	if !data.Securitycheck.IsNull() && !data.Securitycheck.IsUnknown() {
		appfwlearningdata.Securitycheck = data.Securitycheck.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		appfwlearningdata.Target = data.Target.ValueString()
	}

	return appfwlearningdata
}
