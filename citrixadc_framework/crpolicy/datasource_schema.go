package crpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CrpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in cache redirection action: CACHE/ORIGIN.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the cache redirection policy",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the content switching policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}
