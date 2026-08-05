package vpnclientlessaccessprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
	if !data.Clientconsumedcookies.IsNull() && !data.Clientconsumedcookies.IsUnknown() {
		vpnclientlessaccessprofile.Clientconsumedcookies = data.Clientconsumedcookies.ValueString()
	}
	if !data.Javascriptrewritepolicylabel.IsNull() && !data.Javascriptrewritepolicylabel.IsUnknown() {
		vpnclientlessaccessprofile.Javascriptrewritepolicylabel = data.Javascriptrewritepolicylabel.ValueString()
	}
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		vpnclientlessaccessprofile.Profilename = data.Profilename.ValueString()
	}
	if !data.Regexforfindingcustomurls.IsNull() && !data.Regexforfindingcustomurls.IsUnknown() {
		vpnclientlessaccessprofile.Regexforfindingcustomurls = data.Regexforfindingcustomurls.ValueString()
	}
	if !data.Regexforfindingurlincss.IsNull() && !data.Regexforfindingurlincss.IsUnknown() {
		vpnclientlessaccessprofile.Regexforfindingurlincss = data.Regexforfindingurlincss.ValueString()
	}
	if !data.Regexforfindingurlinjavascript.IsNull() && !data.Regexforfindingurlinjavascript.IsUnknown() {
		vpnclientlessaccessprofile.Regexforfindingurlinjavascript = data.Regexforfindingurlinjavascript.ValueString()
	}
	if !data.Regexforfindingurlinxcomponent.IsNull() && !data.Regexforfindingurlinxcomponent.IsUnknown() {
		vpnclientlessaccessprofile.Regexforfindingurlinxcomponent = data.Regexforfindingurlinxcomponent.ValueString()
	}
	if !data.Regexforfindingurlinxml.IsNull() && !data.Regexforfindingurlinxml.IsUnknown() {
		vpnclientlessaccessprofile.Regexforfindingurlinxml = data.Regexforfindingurlinxml.ValueString()
	}
	if !data.Reqhdrrewritepolicylabel.IsNull() && !data.Reqhdrrewritepolicylabel.IsUnknown() {
		vpnclientlessaccessprofile.Reqhdrrewritepolicylabel = data.Reqhdrrewritepolicylabel.ValueString()
	}
	if !data.Requirepersistentcookie.IsNull() && !data.Requirepersistentcookie.IsUnknown() {
		vpnclientlessaccessprofile.Requirepersistentcookie = data.Requirepersistentcookie.ValueString()
	}
	if !data.Reshdrrewritepolicylabel.IsNull() && !data.Reshdrrewritepolicylabel.IsUnknown() {
		vpnclientlessaccessprofile.Reshdrrewritepolicylabel = data.Reshdrrewritepolicylabel.ValueString()
	}
	if !data.Urlrewritepolicylabel.IsNull() && !data.Urlrewritepolicylabel.IsUnknown() {
		vpnclientlessaccessprofile.Urlrewritepolicylabel = data.Urlrewritepolicylabel.ValueString()
	}

	return vpnclientlessaccessprofile
}

func vpnclientlessaccessprofileSetAttrFromGet(ctx context.Context, data *VpnclientlessaccessprofileResourceModel, getResponseData map[string]interface{}) *VpnclientlessaccessprofileResourceModel {
	tflog.Debug(ctx, "In vpnclientlessaccessprofileSetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches with IsUnknown() so a configured value that NITRO
	// omits from GET is not clobbered to null (omit-on-default trap).
	if val, ok := getResponseData["clientconsumedcookies"]; ok && val != nil {
		data.Clientconsumedcookies = types.StringValue(val.(string))
	} else if data.Clientconsumedcookies.IsUnknown() {
		data.Clientconsumedcookies = types.StringNull()
	}
	if val, ok := getResponseData["javascriptrewritepolicylabel"]; ok && val != nil {
		data.Javascriptrewritepolicylabel = types.StringValue(val.(string))
	} else if data.Javascriptrewritepolicylabel.IsUnknown() {
		data.Javascriptrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else if data.Profilename.IsUnknown() {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingcustomurls"]; ok && val != nil {
		data.Regexforfindingcustomurls = types.StringValue(val.(string))
	} else if data.Regexforfindingcustomurls.IsUnknown() {
		data.Regexforfindingcustomurls = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlincss"]; ok && val != nil {
		data.Regexforfindingurlincss = types.StringValue(val.(string))
	} else if data.Regexforfindingurlincss.IsUnknown() {
		data.Regexforfindingurlincss = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinjavascript"]; ok && val != nil {
		data.Regexforfindingurlinjavascript = types.StringValue(val.(string))
	} else if data.Regexforfindingurlinjavascript.IsUnknown() {
		data.Regexforfindingurlinjavascript = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinxcomponent"]; ok && val != nil {
		data.Regexforfindingurlinxcomponent = types.StringValue(val.(string))
	} else if data.Regexforfindingurlinxcomponent.IsUnknown() {
		data.Regexforfindingurlinxcomponent = types.StringNull()
	}
	if val, ok := getResponseData["regexforfindingurlinxml"]; ok && val != nil {
		data.Regexforfindingurlinxml = types.StringValue(val.(string))
	} else if data.Regexforfindingurlinxml.IsUnknown() {
		data.Regexforfindingurlinxml = types.StringNull()
	}
	if val, ok := getResponseData["reqhdrrewritepolicylabel"]; ok && val != nil {
		data.Reqhdrrewritepolicylabel = types.StringValue(val.(string))
	} else if data.Reqhdrrewritepolicylabel.IsUnknown() {
		data.Reqhdrrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["requirepersistentcookie"]; ok && val != nil {
		data.Requirepersistentcookie = types.StringValue(val.(string))
	} else if data.Requirepersistentcookie.IsUnknown() {
		data.Requirepersistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["reshdrrewritepolicylabel"]; ok && val != nil {
		data.Reshdrrewritepolicylabel = types.StringValue(val.(string))
	} else if data.Reshdrrewritepolicylabel.IsUnknown() {
		data.Reshdrrewritepolicylabel = types.StringNull()
	}
	if val, ok := getResponseData["urlrewritepolicylabel"]; ok && val != nil {
		data.Urlrewritepolicylabel = types.StringValue(val.(string))
	} else if data.Urlrewritepolicylabel.IsUnknown() {
		data.Urlrewritepolicylabel = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Profilename.ValueString())

	return data
}
