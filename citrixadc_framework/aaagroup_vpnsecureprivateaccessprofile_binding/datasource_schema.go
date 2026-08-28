package aaagroup_vpnsecureprivateaccessprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaagroupVpnsecureprivateaccessprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT, END, USE_INVOCATION_RESULT, or an expression that evaluates to a number.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Secure Private Access Profile bound to the group.",
			},
		},
	}
}
