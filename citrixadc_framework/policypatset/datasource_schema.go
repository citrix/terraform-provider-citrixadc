package policypatset

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicypatsetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this patset or a pattern bound to this patset.",
			},
			"dynamic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used to populate internal patset information so that the patset can also be used dynamically in an expression. Here dynamically means the patset name can also be derived using an expression. For example for a given patset name \"allow_test\" it can be used dynamically as http.req.url.contains_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default patsets.",
			},
			"dynamiconly": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shows only dynamic patsets when set true.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
			"patsetfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.",
			},
		},
	}
}
