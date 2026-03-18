package vpnclientlessaccessprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnclientlessaccessprofileResourceModel describes the resource data model.
type VpnclientlessaccessprofileResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Clientconsumedcookies          types.String `tfsdk:"clientconsumedcookies"`
	Javascriptrewritepolicylabel   types.String `tfsdk:"javascriptrewritepolicylabel"`
	Profilename                    types.String `tfsdk:"profilename"`
	Regexforfindingcustomurls      types.String `tfsdk:"regexforfindingcustomurls"`
	Regexforfindingurlincss        types.String `tfsdk:"regexforfindingurlincss"`
	Regexforfindingurlinjavascript types.String `tfsdk:"regexforfindingurlinjavascript"`
	Regexforfindingurlinxcomponent types.String `tfsdk:"regexforfindingurlinxcomponent"`
	Regexforfindingurlinxml        types.String `tfsdk:"regexforfindingurlinxml"`
	Reqhdrrewritepolicylabel       types.String `tfsdk:"reqhdrrewritepolicylabel"`
	Requirepersistentcookie        types.String `tfsdk:"requirepersistentcookie"`
	Reshdrrewritepolicylabel       types.String `tfsdk:"reshdrrewritepolicylabel"`
	Urlrewritepolicylabel          types.String `tfsdk:"urlrewritepolicylabel"`
}

func (r *VpnclientlessaccessprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnclientlessaccessprofile resource.",
			},
			"clientconsumedcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the name of the pattern set containing the names of the cookies, which are allowed between the client and the server. If a pattern set is not specified, Citrix Gateway does not allow any cookies between the client and the server. A cookie that is not specified in the pattern set is handled by Citrix Gateway on behalf of the client.",
			},
			"javascriptrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured JavaScript rewrite policy label.  If you do not specify a policy label name, then JAVA scripts are not rewritten.",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Citrix Gateway clientless access profile. Must begin with an ASCII alphabetic or underscore (_) character, and must consist only of ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"regexforfindingcustomurls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URLs in the custom content type other than HTML, CSS, XML, XCOMP, and JavaScript. The custom content type should be included in the patset ns_cvpn_custom_content_types.",
			},
			"regexforfindingurlincss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in the CSS.",
			},
			"regexforfindingurlinjavascript": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in Java script.",
			},
			"regexforfindingurlinxcomponent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in X Component.",
			},
			"regexforfindingurlinxml": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in XML.",
			},
			"reqhdrrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured Request rewrite policy label.  If you do not specify a policy label name, then requests are not rewritten.",
			},
			"requirepersistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether a persistent session cookie is set and accepted for clientless access. If this parameter is set to ON, COM objects, such as MSOffice, which are invoked by the browser can access the files using clientless access. Use caution because the persistent cookie is stored on the disk.",
			},
			"reshdrrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured Response rewrite policy label.",
			},
			"urlrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured URL rewrite policy label. If you do not specify a policy label name, then URLs are not rewritten.",
			},
		},
	}
}

func vpnclientlessaccessprofileGetThePayloadFromtheConfig(ctx context.Context, data *VpnclientlessaccessprofileResourceModel) vpn.Vpnclientlessaccessprofile {
	tflog.Debug(ctx, "In vpnclientlessaccessprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnclientlessaccessprofile := vpn.Vpnclientlessaccessprofile{}
	if !data.Clientconsumedcookies.IsNull() {
		vpnclientlessaccessprofile.Clientconsumedcookies = data.Clientconsumedcookies.ValueString()
	}
	if !data.Javascriptrewritepolicylabel.IsNull() {
		vpnclientlessaccessprofile.Javascriptrewritepolicylabel = data.Javascriptrewritepolicylabel.ValueString()
	}
	if !data.Profilename.IsNull() {
		vpnclientlessaccessprofile.Profilename = data.Profilename.ValueString()
	}
	if !data.Regexforfindingcustomurls.IsNull() {
		vpnclientlessaccessprofile.Regexforfindingcustomurls = data.Regexforfindingcustomurls.ValueString()
	}
	if !data.Regexforfindingurlincss.IsNull() {
		vpnclientlessaccessprofile.Regexforfindingurlincss = data.Regexforfindingurlincss.ValueString()
	}
	if !data.Regexforfindingurlinjavascript.IsNull() {
		vpnclientlessaccessprofile.Regexforfindingurlinjavascript = data.Regexforfindingurlinjavascript.ValueString()
	}
	if !data.Regexforfindingurlinxcomponent.IsNull() {
		vpnclientlessaccessprofile.Regexforfindingurlinxcomponent = data.Regexforfindingurlinxcomponent.ValueString()
	}
	if !data.Regexforfindingurlinxml.IsNull() {
		vpnclientlessaccessprofile.Regexforfindingurlinxml = data.Regexforfindingurlinxml.ValueString()
	}
	if !data.Reqhdrrewritepolicylabel.IsNull() {
		vpnclientlessaccessprofile.Reqhdrrewritepolicylabel = data.Reqhdrrewritepolicylabel.ValueString()
	}
	if !data.Requirepersistentcookie.IsNull() {
		vpnclientlessaccessprofile.Requirepersistentcookie = data.Requirepersistentcookie.ValueString()
	}
	if !data.Reshdrrewritepolicylabel.IsNull() {
		vpnclientlessaccessprofile.Reshdrrewritepolicylabel = data.Reshdrrewritepolicylabel.ValueString()
	}
	if !data.Urlrewritepolicylabel.IsNull() {
		vpnclientlessaccessprofile.Urlrewritepolicylabel = data.Urlrewritepolicylabel.ValueString()
	}

	return vpnclientlessaccessprofile
}

func vpnclientlessaccessprofileSetAttrFromGet(ctx context.Context, data *VpnclientlessaccessprofileResourceModel, getResponseData map[string]interface{}) *VpnclientlessaccessprofileResourceModel {
	tflog.Debug(ctx, "In vpnclientlessaccessprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientconsumedcookies"]; ok && val != nil {
		data.Clientconsumedcookies = types.StringValue(val.(string))
	} else {
		data.Clientconsumedcookies = types.StringNull()
	}
	if val, ok := getResponseData["javascriptrewritepolicylabel"]; ok && val != nil {
		data.Javascriptrewritepolicylabel = types.StringValue(val.(string))
	} else {
		data.Javascriptrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingcustomurls"]; ok && val != nil {
		data.Regexforfindingcustomurls = types.StringValue(val.(string))
	} else {
		data.Regexforfindingcustomurls = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlincss"]; ok && val != nil {
		data.Regexforfindingurlincss = types.StringValue(val.(string))
	} else {
		data.Regexforfindingurlincss = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinjavascript"]; ok && val != nil {
		data.Regexforfindingurlinjavascript = types.StringValue(val.(string))
	} else {
		data.Regexforfindingurlinjavascript = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinxcomponent"]; ok && val != nil {
		data.Regexforfindingurlinxcomponent = types.StringValue(val.(string))
	} else {
		data.Regexforfindingurlinxcomponent = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinxml"]; ok && val != nil {
		data.Regexforfindingurlinxml = types.StringValue(val.(string))
	} else {
		data.Regexforfindingurlinxml = types.StringNull()
	}
	if val, ok := getResponseData["reqhdrrewritepolicylabel"]; ok && val != nil {
		data.Reqhdrrewritepolicylabel = types.StringValue(val.(string))
	} else {
		data.Reqhdrrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["requirepersistentcookie"]; ok && val != nil {
		data.Requirepersistentcookie = types.StringValue(val.(string))
	} else {
		data.Requirepersistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["reshdrrewritepolicylabel"]; ok && val != nil {
		data.Reshdrrewritepolicylabel = types.StringValue(val.(string))
	} else {
		data.Reshdrrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["urlrewritepolicylabel"]; ok && val != nil {
		data.Urlrewritepolicylabel = types.StringValue(val.(string))
	} else {
		data.Urlrewritepolicylabel = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Profilename.ValueString())

	return data
}
