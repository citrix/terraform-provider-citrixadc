package appfwprofile_xmlxss_binding

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

// AppfwprofileXmlxssBindingResourceModel describes the resource data model.
type AppfwprofileXmlxssBindingResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Alertonly            types.String `tfsdk:"alertonly"`
	AsScanLocationXmlxss types.String `tfsdk:"as_scan_location_xmlxss"`
	Comment              types.String `tfsdk:"comment"`
	Isautodeployed       types.String `tfsdk:"isautodeployed"`
	IsregexXmlxss        types.String `tfsdk:"isregex_xmlxss"`
	Name                 types.String `tfsdk:"name"`
	Resourceid           types.String `tfsdk:"resourceid"`
	State                types.String `tfsdk:"state"`
	Xmlxss               types.String `tfsdk:"xmlxss"`
}

func (r *AppfwprofileXmlxssBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_xmlxss_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_scan_location_xmlxss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location of XSS injection exception - XML Element or Attribute.",
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
			"isregex_xmlxss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the XML XSS exempted field name a regular expression?",
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
			"xmlxss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exempt the specified URL from the XML cross-site scripting (XSS) check.\nAn XML cross-site scripting exemption (relaxation) consists of the following items:\n* URL. URL to exempt, as a string or a PCRE-format regular expression.\n* ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string.\n* Location. ELEMENT if the attachment is located in an XML element, ATTRIBUTE if located in an XML attribute.",
			},
		},
	}
}

func appfwprofile_xmlxss_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileXmlxssBindingResourceModel) appfw.Appfwprofilexmlxssbinding {
	tflog.Debug(ctx, "In appfwprofile_xmlxss_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_xmlxss_binding := appfw.Appfwprofilexmlxssbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_xmlxss_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationXmlxss.IsNull() {
		appfwprofile_xmlxss_binding.Asscanlocationxmlxss = data.AsScanLocationXmlxss.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_xmlxss_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_xmlxss_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexXmlxss.IsNull() {
		appfwprofile_xmlxss_binding.Isregexxmlxss = data.IsregexXmlxss.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_xmlxss_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_xmlxss_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_xmlxss_binding.State = data.State.ValueString()
	}
	if !data.Xmlxss.IsNull() {
		appfwprofile_xmlxss_binding.Xmlxss = data.Xmlxss.ValueString()
	}

	return appfwprofile_xmlxss_binding
}

func appfwprofile_xmlxss_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmlxssBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlxssBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlxss_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_scan_location_xmlxss"]; ok && val != nil {
		data.AsScanLocationXmlxss = types.StringValue(val.(string))
	} else {
		data.AsScanLocationXmlxss = types.StringNull()
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
	if val, ok := getResponseData["isregex_xmlxss"]; ok && val != nil {
		data.IsregexXmlxss = types.StringValue(val.(string))
	} else {
		data.IsregexXmlxss = types.StringNull()
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
	if val, ok := getResponseData["xmlxss"]; ok && val != nil {
		data.Xmlxss = types.StringValue(val.(string))
	} else {
		data.Xmlxss = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_xmlxss:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsScanLocationXmlxss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlxss:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Xmlxss.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
