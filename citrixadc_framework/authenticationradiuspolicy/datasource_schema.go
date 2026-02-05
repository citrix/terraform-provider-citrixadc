package authenticationradiuspolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationradiuspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RADIUS authentication policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after RADIUS policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the RADIUS action to perform if the policy matches.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the RADIUS server.",
			},
		},
	}
}
