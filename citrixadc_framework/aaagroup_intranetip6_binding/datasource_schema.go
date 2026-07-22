package aaagroup_intranetip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaagroupIntranetip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"intranetip6": schema.StringAttribute{
				Required:    true,
				Description: "The Intranet IP6(s) bound to the group",
			},
			"numaddr": schema.Int64Attribute{
				Computed:    true,
				Description: "Numbers of ipv6 address bound starting with intranetip6",
			},
		},
	}
}
