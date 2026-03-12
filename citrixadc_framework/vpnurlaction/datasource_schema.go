package vpnurlaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnurlactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actualurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address for the bookmark link.",
			},
			"applicationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN",
			},
			"clientlessaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on NetScaler Gateway for HTTPS resources.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the bookmark link.",
			},
			"iconurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to fetch icon file for displaying this resource.",
			},
			"linkname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description of the bookmark link. The description appears in the Access Interface.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bookmark link.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the vpn urlAction.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl action\" or 'my vpnurl action').",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO",
			},
			"ssotype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Single sign on type for unified gateway",
			},
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the associated vserver to handle selfAuth SSO",
			},
		},
	}
}
