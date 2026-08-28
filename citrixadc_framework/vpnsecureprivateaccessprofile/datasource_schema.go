package vpnsecureprivateaccessprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnsecureprivateaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of Secure Private Access profile.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public URL for your Secure Private Access site or load balancing server.",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID of the citrix cloud customer.",
			},
			"chromeenterprisepremiummode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secure Private Access Chrome Enterprise Premium mode of operation. Possible values = OFF, WITH_PARTNER_CONNECTOR, WITHOUT_PARTNER_CONNECTOR",
			},
			"googlecustomerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Your organization's unique ID on Google's Admin console Profile settings.",
			},
			"googlesecuritygatewayid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the Google Secure Gateway.",
			},
			"forceclienttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Automatically configures the session for Citrix Secure Access client connectivity. Possible values = ON, OFF",
			},
			"sharedsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secure Private Access Shared Secret.",
			},
			"sharedsecret_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Secure Private Access Shared Secret.",
			},
			"sharedsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a sharedsecret_wo update.",
			},
		},
	}
}
