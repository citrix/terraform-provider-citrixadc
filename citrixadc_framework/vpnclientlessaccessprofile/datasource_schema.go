package vpnclientlessaccessprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnclientlessaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
