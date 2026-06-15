package linkset_channel_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LinksetChannelBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"linkset_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the linkset to which to bind the interfaces.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "The interfaces to be bound to the linkset.",
			},
		},
	}
}
