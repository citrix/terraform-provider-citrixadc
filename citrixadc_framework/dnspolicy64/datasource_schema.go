package dnspolicy64

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Dnspolicy64DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows:\n* A default dns64 action with prefix <default prefix> and mapped and exclude are any\nYou can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS64 policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which DNS traffic is evaluated.\nNote:\n* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.\nExample: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)",
			},
		},
	}
}
