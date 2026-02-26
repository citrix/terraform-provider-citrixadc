package vpnglobal_domain_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnglobalDomainBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetdomain": schema.StringAttribute{
				Required:    true,
				Description: "The conflicting intranet domain name.",
			},
		},
	}
}
