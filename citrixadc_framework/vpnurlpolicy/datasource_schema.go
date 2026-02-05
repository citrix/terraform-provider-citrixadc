package vpnurlpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnurlpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be applied by the new urlPolicy if the rule criteria are met.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new urlPolicy.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the vpn urlPolicy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl policy\" or 'my vpnurl policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the NetScaler CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}
