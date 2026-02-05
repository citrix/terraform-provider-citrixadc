package linkset

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func LinksetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"linkset_id": schema.StringAttribute{
				Required:    true,
				Description: "Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.",
			},
			"interfacebinding": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Set of interface bindings for the linkset.",
			},
		},
	}
}
