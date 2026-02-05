package aaapreauthenticationpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaapreauthenticationpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the preauthentication policy. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the preauthentication policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action that the policy is to invoke when a connection matches the policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, defining connections that match the policy.",
			},
		},
	}
}
