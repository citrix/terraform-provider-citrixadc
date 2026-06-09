package lbwlm_lbvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbwlmLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server which is to be bound to the WLM.",
			},
			"wlmname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the Work Load Manager.",
			},
		},
	}
}
