package authenticationpushservice

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationpushserviceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity for communicating with Citrix Push server in cloud.",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique secret for communicating with Citrix Push server in cloud.",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer id/name of the account in cloud that is used to create clientid/secret pair.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which certificates or idtoken is refreshed.",
			},
		},
	}
}
