package vpnsessionpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnsessionpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be applied by the new session policy if the rule criteria are met.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new session policy that is applied after the user logs on to Citrix Gateway.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}
