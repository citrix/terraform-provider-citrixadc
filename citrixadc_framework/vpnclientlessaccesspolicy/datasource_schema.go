package vpnclientlessaccesspolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnclientlessaccesspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the new clientless access policy.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile to invoke for the clientless access.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}
