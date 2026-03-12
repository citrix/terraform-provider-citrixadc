package vpnurl

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnurlDataSourceSchema() schema.Schema {
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
			"appjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To store the template details in the json format.",
			},
			"applicationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN",
			},
			"clientlessaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.",
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
			"urlname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bookmark link.",
			},
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the associated LB/CS vserver",
			},
		},
	}
}
